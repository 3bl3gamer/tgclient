package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"github.com/ansel1/merry/v2"
)

const ErrorBufStackKey = "mtproto_decode_err_buf_stack"

func init() {
	merry.RegisterDetail("Buffer stack", ErrorBufStackKey)
}

func notEnoughBytesErr(kind string, pos, inc, size int) error {
	err := fmt.Errorf("%s: not enough bytes in buffer: expected %d + %d = %d, got %d",
		kind, pos, inc, pos+inc, size)
	return merry.WrapSkipping(err, 1)
}

func unexpectedTypeErr[T TL](kind string, obj TL) error {
	var dummy T
	err := fmt.Errorf("%s: expected item of type %T, got %T", kind, dummy, obj)
	return merry.WrapSkipping(err, 1)
}

type DecodeBuf struct {
	buf  []byte
	off  int
	size int
	err  error
}

func NewDecodeBuf(b []byte) *DecodeBuf {
	return &DecodeBuf{b, 0, len(b), nil}
}

func (m *DecodeBuf) Err() error {
	return m.err
}

func (m *DecodeBuf) RemainingLen() int {
	return m.size - m.off
}

func (m *DecodeBuf) SeekBack(n int) {
	if m.off >= n {
		m.off -= n
	} else {
		m.err = merry.Errorf("tried to move back %d byte(s) from offset %d", n, m.off)
	}
}

func (m *DecodeBuf) Bool() bool {
	constructor := m.UInt()
	if m.err != nil {
		return false
	}
	switch constructor {
	case CRC_boolFalse:
		return false
	case CRC_boolTrue:
		return true
	}
	return false
}

func (m *DecodeBuf) Long() int64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = notEnoughBytesErr("DecodeLong", m.off, 8, m.size)
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Double() float64 {
	if m.err != nil {
		return 0
	}
	if m.off+8 > m.size {
		m.err = notEnoughBytesErr("DecodeDouble", m.off, 8, m.size)
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(m.buf[m.off : m.off+8]))
	m.off += 8
	return x
}

func (m *DecodeBuf) Int() int32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = notEnoughBytesErr("DecodeInt", m.off, 4, m.size)
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return int32(x)
}

func (m *DecodeBuf) UInt() uint32 {
	if m.err != nil {
		return 0
	}
	if m.off+4 > m.size {
		m.err = notEnoughBytesErr("DecodeUInt", m.off, 4, m.size)
		return 0
	}
	x := binary.LittleEndian.Uint32(m.buf[m.off : m.off+4])
	m.off += 4
	return x
}

func (m *DecodeBuf) Bytes(size int) []byte {
	if m.err != nil {
		return nil
	}
	if m.off+size > m.size {
		m.err = notEnoughBytesErr("DecodeBytes", m.off, size, m.size)
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size
	return x
}

func (m *DecodeBuf) Bytes16() [16]byte {
	return [16]byte(m.Bytes(16))
}

func (m *DecodeBuf) Bytes32() [32]byte {
	return [32]byte(m.Bytes(32))
}

func (m *DecodeBuf) StringBytes() []byte {
	if m.err != nil {
		return nil
	}
	var size, padding int

	if m.off+1 > m.size {
		m.err = notEnoughBytesErr("DecodeStringBytes", m.off, 1, m.size)
		return nil
	}
	size = int(m.buf[m.off])
	m.off++
	padding = (4 - ((size + 1) % 4)) & 3
	if size == 254 {
		if m.off+3 > m.size {
			m.err = notEnoughBytesErr("DecodeStringBytes", m.off, 3, m.size)
			return nil
		}
		size = int(m.buf[m.off]) | int(m.buf[m.off+1])<<8 | int(m.buf[m.off+2])<<16
		m.off += 3
		padding = (4 - size%4) & 3
	}

	if m.off+size > m.size {
		m.err = notEnoughBytesErr("DecodeStringBytes", m.off, size, m.size)
		return nil
	}
	x := make([]byte, size)
	copy(x, m.buf[m.off:m.off+size])
	m.off += size

	if m.off+padding > m.size {
		m.err = notEnoughBytesErr("DecodeStringBytes: padding", m.off, padding, m.size)
		return nil
	}
	m.off += padding

	return x
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *DecodeBuf) BigInt() *big.Int {
	b := m.StringBytes()
	if m.err != nil {
		return nil
	}
	y := make([]byte, len(b)+1)
	y[0] = 0
	copy(y[1:], b)
	x := new(big.Int).SetBytes(y)
	return x
}

func (m *DecodeBuf) VectorInt() []int32 {
	size := m.vectorHeader("DecodeVectorInt")
	if m.err != nil {
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := m.Int()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorLong() []int64 {
	size := m.vectorHeader("DecodeVectorLong")
	if m.err != nil {
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := m.Long()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorString() []string {
	size := m.vectorHeader("DecodeVectorString")
	if m.err != nil {
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := m.String()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) VectorBytes() [][]byte {
	size := m.vectorHeader("DecodeVectorBytes")
	if m.err != nil {
		return nil
	}
	x := make([][]byte, size)
	i := int32(0)
	for i < size {
		y := m.StringBytes()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func DecodeBuf_GenericVector[T TL](m *DecodeBuf) []T {
	size := m.vectorHeader("DecodeVector")
	if m.err != nil {
		return nil
	}
	x := make([]T, size)
	i := int32(0)
	for i < size {
		objTL := m.Object()
		if m.err != nil {
			return nil
		}
		obj, ok := objTL.(T)
		if !ok {
			m.err = unexpectedTypeErr[T]("GenericVector", objTL)
		}
		if m.err != nil {
			return nil
		}
		x[i] = obj
		i++
	}
	return x
}

func (m *DecodeBuf) Vector() []TL {
	return DecodeBuf_GenericVector[TL](m)
}

func (m *DecodeBuf) Vector2d() [][]TL {
	size := m.vectorHeader("DecodeVector2d")
	if m.err != nil {
		return nil
	}
	x := make([][]TL, size)
	i := int32(0)
	for i < size {
		y := m.Vector()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func DecodeBuf_GenericObject[T TL](m *DecodeBuf) T {
	constructor := m.UInt()
	if m.err != nil {
		var blank T
		return blank
	}
	objTL := m.ObjectGenerated(constructor)
	if m.err != nil {
		var blank T
		return blank
	}
	obj, ok := objTL.(T)
	if !ok {
		m.err = unexpectedTypeErr[T]("GenericVector", objTL)
		var blank T
		return blank
	}
	return obj
}

func (m *DecodeBuf) Object() TL {
	return DecodeBuf_GenericObject[TL](m)
}

func (m *DecodeBuf) vectorHeader(errLabel string) int32 {
	constructor := m.UInt()
	if m.err != nil {
		return 0
	}
	if constructor != CRC_vector {
		m.err = merry.Errorf("%s: wrong constructor (0x%08x)", errLabel, constructor)
		return 0
	}
	size := m.Int()
	if m.err != nil {
		return 0
	}
	if size < 0 {
		m.err = merry.Errorf("%s: negative size: %d", errLabel, size)
		return 0
	}
	return size
}

func (d *DecodeBuf) dump() {
	fmt.Println(hex.Dump(d.buf[d.off:d.size]))
}

func (d *DecodeBuf) pushToErrBufStack(objStartOffset int, constructor uint32) {
	if d.err == nil {
		return
	}

	constructorBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(constructorBytes, constructor)

	bufStack, _ := merry.Value(d.err, ErrorBufStackKey).(string)
	startOffset := 0
	if objStartOffset-64 > 0 {
		startOffset = objStartOffset - 32
	}
	endOffset := d.size
	if d.off+64 < endOffset {
		endOffset = d.off + 32
	}
	ellipsis := ""
	if endOffset < d.size {
		ellipsis = "... ... ...\n"
	}

	bufStack += fmt.Sprintf(
		`
constructor:        %08x (%02x %02x %02x %02x)
obj decoded from:   %d 0x%x (absolute: %d 0x%x)
offset after error: %d 0x%x (absolute: %d 0x%x)
this slice offset:  %d 0x%x
full buffer size:   %d 0x%x
%s%s
`,
		constructor, constructorBytes[0], constructorBytes[1], constructorBytes[2], constructorBytes[3],
		objStartOffset-startOffset, objStartOffset-startOffset, objStartOffset, objStartOffset,
		d.off-startOffset, d.off-startOffset, d.off, d.off,
		startOffset, startOffset,
		d.size, d.size,
		hex.Dump(d.buf[startOffset:endOffset]), ellipsis,
	)
	d.err = merry.Wrap(d.err, merry.WithValue(ErrorBufStackKey, bufStack))
}

func str2big(str string) *big.Int {
	return new(big.Int).SetBytes([]byte(str))
}

func big2str(val *big.Int) string {
	return string(val.Bytes())
}

// decodeMessage decodes types before actual RPC response by itself,
// then decodes remaining data by DecodeBuf or RLReq.decodeResponse().
// This all could be done in DecodeBuf, but...
// Some TL methods return bare vectors:
//
//	account.getWallPapers#c04cfac2 = Vector<WallPaper>
//	messages.getMessagesViews#c4c8a55d ... = Vector<int>
//	photos.deletePhotos#87cf7f2f ... = Vector<long>
//
// All this vectors will be received with same constructor ID: 0x1cb5c415.
// To find out correct vector type, we need no find a request for this response.
// Request message ID is in `rpc_result.reqMsgID`. So we have to decode everything
// until `rpc_result`, find request and use it (reqMsg arg) for further decoding.
func (m *MTProto) decodeMessage(dbuf *DecodeBuf, reqMsg TLReq) (r TL) {
	constructor := dbuf.UInt()
	if dbuf.err != nil {
		return nil
	}

	switch constructor {
	case CRC_msg_container:
		size := dbuf.Int()
		arr := make([]TL_mtMessage, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_mtMessage{dbuf.Long(), dbuf.Int(), dbuf.Int(), m.decodeMessage(dbuf, reqMsg)}
			if dbuf.err != nil {
				return nil
			}
		}
		r = TL_msgContainer{arr}

	case CRC_rpc_result:
		requestID := dbuf.Long()
		m.mutex.Lock()
		packet, ok := m.msgsByID[requestID]
		m.mutex.Unlock()
		if ok {
			if req, ok := packet.msg.(TLReq); ok {
				//r = req.decodeResponse(dbuf)
				r = m.decodeMessage(dbuf, req)
			} else {
				r = m.decodeMessage(dbuf, nil)
				m.log.Warn("got RPC result (%T) not-a-request message #%d #T", r, requestID, packet.msg)
			}
		} else {
			r = m.decodeMessage(dbuf, nil)
			m.log.Warn("got RPC result (%T) for unknown message #%d", r, requestID)
		}
		r = TL_rpcResult{requestID, r}

	case CRC_gzip_packed:
		obj := make([]byte, 0, 4096)

		var buf bytes.Buffer
		_, _ = buf.Write(dbuf.StringBytes())
		gz, _ := gzip.NewReader(&buf)

		b := make([]byte, 4096)
		for true {
			n, _ := gz.Read(b)
			obj = append(obj, b[0:n]...)
			if n <= 0 {
				break
			}
		}
		d := NewDecodeBuf(obj)
		r = m.decodeMessage(d, reqMsg)
		dbuf.err = d.err

	default:
		dbuf.SeekBack(4) //returning constructor ID
		if reqMsg == nil || constructor == CRC_rpcError {
			r = dbuf.Object()
		} else {
			r = reqMsg.decodeResponse(dbuf)
		}
	}

	if dbuf.err != nil {
		return nil
	}
	return r
}
