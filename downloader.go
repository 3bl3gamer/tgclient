package tgclient

import (
	"io"
	"mtproto"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/ansel1/merry"
)

func clampI(a, v, b int) int {
	if v < a {
		v = a
	}
	if v > b {
		v = b
	}
	return v
}

type FileResponse struct {
	DcID int32
	Data []byte
	Err  error
}

type filePart struct {
	dcID          int32
	location      mtproto.TL
	outChan       chan *FileResponse
	offset, limit int32
}

type FilePartsResult struct {
	Finished      bool
	ActualDcID    int32
	BytesRecieved int
	BytesWritten  int
}

type Downloader struct {
	tg             *TGClient
	fileMTs        map[int32]*mtproto.MTProto
	fileMTsMutex   *sync.Mutex
	filePartsQueue chan *filePart
	log            mtproto.Logger
}

func NewDownloader(tg *TGClient) *Downloader {
	return &Downloader{
		tg:             tg,
		fileMTs:        make(map[int32]*mtproto.MTProto),
		fileMTsMutex:   &sync.Mutex{},
		filePartsQueue: make(chan *filePart, 4),
		log:            tg.log,
	}
}

func (d *Downloader) DownloadFileToPath(
	fpath string, fileLocation mtproto.TL, dcID int32, size int64,
) (*FilePartsResult, error) {
	partSize := int64(512 * 1024)
	tempFpath := fpath + ".temp"
	if err := os.MkdirAll(filepath.Dir(tempFpath), os.ModePerm); err != nil {
		return nil, merry.Wrap(err)
	}

	fd, err := os.OpenFile(tempFpath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	defer fd.Close()

	offset, err := fd.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	if offset%int64(partSize) != 0 {
		d.log.Warn("file '%s' exists but size is not multiple of block size (%d % %d != 0), moving to start",
			tempFpath, offset, partSize)
		offset, err = fd.Seek(0, io.SeekStart)
		if err != nil {
			return nil, merry.Wrap(err)
		}
		if err := fd.Truncate(0); err != nil {
			return nil, merry.Wrap(err)
		}
	}

	partsRes, err := d.DownloadFileParts(fd, fileLocation, dcID, size, partSize, offset)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	if err := fd.Close(); err != nil {
		return nil, merry.Wrap(err)
	}
	if err := os.Rename(tempFpath, fpath); err != nil {
		return nil, merry.Wrap(err)
	}
	return partsRes, nil
}

func (d *Downloader) DownloadFileParts(
	file io.Writer, fileLocation mtproto.TL, dcID int32, size, partSize, offset int64,
) (*FilePartsResult, error) {
	partsRes := &FilePartsResult{ActualDcID: dcID}

	partsCount := int((size - offset + partSize - 1) / partSize)
	resChans := make([]chan *FileResponse, clampI(1, partsCount, 4))

	for i := 0; i < len(resChans); i++ {
		resChans[i] = d.ReqestFilePart(dcID, fileLocation,
			offset+partSize*int64(i), partSize)
	}

	for {
		res := <-resChans[0]
		if res.Err != nil {
			return nil, merry.Wrap(res.Err)
		}

		offset += partSize
		for i := 1; i < len(resChans); i++ {
			resChans[i-1] = resChans[i]
		}
		newPartOffset := offset + partSize*int64(len(resChans)-1)
		if newPartOffset < size || len(resChans) == 1 {
			resChans[len(resChans)-1] = d.ReqestFilePart(dcID, fileLocation, newPartOffset, partSize)
		} else {
			resChans = resChans[:len(resChans)-1]
		}

		partsRes.ActualDcID = res.DcID
		partsRes.BytesRecieved += len(res.Data)

		n, err := file.Write(res.Data)
		if err != nil {
			return nil, merry.Wrap(err)
		}
		partsRes.BytesWritten += n

		if len(res.Data) < int(partSize) {
			partsRes.Finished = true
			break
		}
	}
	return partsRes, nil
}

func (d *Downloader) ReqestFilePart(dcID int32, fileLocation mtproto.TL, offset, limit int64) chan *FileResponse {
	part := &filePart{
		dcID:     dcID,
		location: fileLocation,
		outChan:  make(chan *FileResponse, 1),
		limit:    int32(limit),
		offset:   int32(offset),
	}
	d.filePartsQueue <- part
	return part.outChan
}

func (d *Downloader) partsDownloadRoutine() {
	for part := range d.filePartsQueue {
		fileResp := FileResponse{DcID: part.dcID}

		mt, err := d.getFileMT(part.dcID)
		if err != nil {
			fileResp.Err = merry.Wrap(err)
			part.outChan <- &fileResp
			continue
		}

		resTL := mt.SendSync(mtproto.TL_upload_getFile{
			Location: part.location,
			Offset:   part.offset,
			Limit:    part.limit,
		})

		switch res := resTL.(type) {
		case mtproto.TL_upload_file:
			fileResp.Data = res.Bytes
		case mtproto.TL_upload_fileCdnRedirect:
			fileResp.Err = merry.New("cdn redirect: " + mtproto.Sprint(res))
		case mtproto.TL_rpc_error:
			if strings.HasPrefix(res.ErrorMessage, "FILE_MIGRATE_") {
				d.log.Warn("got %s, part DC is %d", res.ErrorMessage, part.dcID)
				id, _ := strconv.Atoi(res.ErrorMessage[13:])
				part.dcID = int32(id)
				select {
				case d.filePartsQueue <- part:
					continue
				default:
					fileResp.Err = merry.New("file queue overflow while handling DC migration error")
				}
			} else {
				fileResp.Err = merry.New(mtproto.UnexpectedTL("file part", resTL))
			}
		default:
			fileResp.Err = merry.New(mtproto.UnexpectedTL("file part", resTL))
		}

		part.outChan <- &fileResp
		close(part.outChan)
	}
}

func (d *Downloader) getFileMT(dcID int32) (*mtproto.MTProto, error) {
	d.fileMTsMutex.Lock()
	defer d.fileMTsMutex.Unlock()

	mt, _ := d.fileMTs[dcID]
	if mt != nil {
		return mt, nil
	}

	var err error
	mt, err = d.tg.mt.NewConnection(dcID)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	d.log.Info("connected to file DC %d", dcID)
	d.fileMTs[dcID] = mt
	return mt, nil
}
