package tgclient

import (
	"fmt"
	"mtproto"
	"reflect"
	"strconv"
	"strings"

	"github.com/ansel1/merry"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type TGClient struct {
	mt                   *mtproto.MTProto
	fileMTs              map[int32]*mtproto.MTProto
	updatesState         *mtproto.TL_updates_state
	filePartsQueue       chan *filePart
	handleUpdateExternal UpdateHandler
	log                  Logger
}

type UpdateHandler func(mtproto.TL)

func NewTGClient(appID int32, appHash string, handleUpdate UpdateHandler, log Logger) (*TGClient, error) {
	mt, err := mtproto.NewMTProto(appID, appHash)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	if err := mt.Connect(); err != nil {
		return nil, merry.Wrap(err)
	}

	client := &TGClient{
		mt:                   mt,
		fileMTs:              make(map[int32]*mtproto.MTProto),
		updatesState:         &mtproto.TL_updates_state{},
		filePartsQueue:       make(chan *filePart, 4),
		handleUpdateExternal: handleUpdate,
		log:                  log,
	}
	mt.SetEventsHandler(client.handleEvent)
	for i := 0; i < 4; i++ {
		go client.filesRoutine()
	}
	return client, nil
}

func (c *TGClient) handleEvent(eventObj mtproto.TL) {
	switch event := eventObj.(type) {
	case mtproto.TL_updatesTooLong:
		//TODO: what?
		// Too many updates, it is necessary to execute updates.getDifference.
		// https://core.telegram.org/constructor/updatesTooLong
		fmt.Println("[WARN] updates too long")
	case mtproto.TL_updateShort:
		c.updatesState.Date = event.Date
		c.handleUpdate(event.Update)
	case mtproto.TL_updates:
		c.updatesState.Date = event.Date
		c.updatesState.Seq = event.Seq
		c.rememberEventExtraData(event.Users)
		c.rememberEventExtraData(event.Chats)
		for _, u := range event.Updates {
			c.handleUpdate(u)
		}
	case mtproto.TL_updateShortMessage:
		c.updatesState.Date = event.Date
		c.updatesState.Pts = event.Pts
		// update.PtsCount
		c.handleUpdate(event)
	case mtproto.TL_updateShortChatMessage:
		c.updatesState.Date = event.Date
		c.updatesState.Pts = event.Pts
		// update.PtsCount
		c.handleUpdate(event)
	case mtproto.TL_updatesCombined:
		c.updatesState.Date = event.Date
		c.updatesState.Seq = event.Seq
		c.rememberEventExtraData(event.Users)
		c.rememberEventExtraData(event.Chats)
		// update.SeqStart
		for _, u := range event.Updates {
			c.handleUpdate(u)
		}
	case mtproto.TL_updateShortSentMessage:
		c.updatesState.Pts = event.Pts
		// update.PtsCount
		c.handleUpdate(event)
	case mtproto.TL_pong:
	case mtproto.TL_rpc_result:
		// do nothing
	default:
		c.log.Warning(mtproto.UnexpectedTL("event", eventObj))
	}
}

func (e *TGClient) handleUpdate(obj mtproto.TL) {
	value := reflect.ValueOf(obj).FieldByName("Pts")
	if value != (reflect.Value{}) {
		e.updatesState.Pts = int32(value.Int())
	}
	e.handleUpdateExternal(obj)
}

func (c *TGClient) rememberEventExtraData(objs []mtproto.TL) {
	for _, obj := range objs {
		switch x := obj.(type) {
		case mtproto.TL_user:
			c.log.Info("extra: user: ", x.ID, x.Username)
		case mtproto.TL_chat:
			c.log.Info("extra: chat: ", x.ID, x.Title)
		case mtproto.TL_channel:
			c.log.Info("extra: channel: ", x.ID, x.Username)
		default:
			c.log.Warning(mtproto.UnexpectedTL("extra data", obj))
		}
	}
}

func (c *TGClient) AuthAndInitEvents() error {
	for {
		res := c.mt.SendSync(mtproto.TL_updates_getState{})
		if mtproto.IsErrorType(res, mtproto.TL_ErrUnauthorized) { //AUTH_KEY_UNREGISTERED SESSION_REVOKED SESSION_EXPIRED
			if err := c.mt.Auth(); err != nil {
				return merry.Wrap(err)
			}
			continue
		}
		_, ok := res.(mtproto.TL_updates_state)
		if !ok {
			return mtproto.WrongRespError(res)
		}
		break
	}
	c.log.Info("Seems authed.")
	return nil
}

func (c *TGClient) SendSync(msg mtproto.TL) mtproto.TL {
	return c.mt.SendSync(msg)
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

func (c *TGClient) filesRoutine() {
	for part := range c.filePartsQueue {
		fileResp := FileResponse{DcID: part.dcID}

		mt, err := c.getFileMT(part.dcID)
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
				c.log.Infof("got %s, part DC is %d", res.ErrorMessage, part.dcID)
				id, _ := strconv.Atoi(res.ErrorMessage[13:])
				part.dcID = int32(id)
				select {
				case c.filePartsQueue <- part:
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

func (c *TGClient) getFileMT(dcID int32) (*mtproto.MTProto, error) {
	mt, _ := c.fileMTs[dcID]
	if mt != nil {
		return mt, nil
	}
	session := c.mt.SessionCopy()
	session.Addr = c.mt.DCAddr(dcID)
	mt, err := mtproto.NewMTProtoExt(c.mt.AppConfig(), &mtproto.SessNoopStore{}, session)
	if err != nil {
		return nil, merry.Wrap(err)
	}
	if err := mt.Connect(); err != nil {
		return nil, merry.Wrap(err)
	}
	return mt, nil
}

func (c *TGClient) ReqestFilePart(dcID int32, fileLocation mtproto.TL, offset, limit int32) chan *FileResponse {
	part := &filePart{
		dcID:     dcID,
		location: fileLocation,
		outChan:  make(chan *FileResponse, 1),
		limit:    limit,
		offset:   offset,
	}
	c.filePartsQueue <- part
	return part.outChan
}
