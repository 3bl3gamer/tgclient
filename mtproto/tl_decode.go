package mtproto

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"

	"github.com/ansel1/merry"
)

func notEnoughBytesErr(kind string, pos, inc, size int) error {
	err := merry.Errorf("%s: not enough bytes in buffer: expected %d + %d = %d, got %d",
		kind, pos, inc, pos+inc, size)
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

func (m *DecodeBuf) SeekBack(n int) {
	if m.off >= n {
		m.off -= n
	} else {
		m.err = merry.Errorf("tried to move back %d byte(s) from offset %d", n, m.off)
	}
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

func (m *DecodeBuf) FlaggedLong(flags, num int32) int64 {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return 0
	}
	return m.Long()
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

func (m *DecodeBuf) FlaggedDouble(flags, num int32) float64 {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return 0
	}
	return m.Double()
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

func (m *DecodeBuf) FlaggedInt(flags, num int32) int32 {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return 0
	}
	return m.Int()
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

func (m *DecodeBuf) FlaggedUInt(flags, num int32) uint32 {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return 0
	}
	return m.UInt()
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

func (m *DecodeBuf) FlaggedStringBytes(flags, num int32) []byte {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.StringBytes()
}

func (m *DecodeBuf) String() string {
	b := m.StringBytes()
	if m.err != nil {
		return ""
	}
	x := string(b)
	return x
}

func (m *DecodeBuf) FlaggedString(flags, num int32) string {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return ""
	}
	return m.String()
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

func (m *DecodeBuf) FlaggedBigInt(flags, num int32) *big.Int {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.BigInt()
}

func (m *DecodeBuf) VectorInt() []int32 {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != CRC_vector {
		m.err = merry.Errorf("DecodeVectorInt: wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = merry.Errorf("DecodeVectorInt: negative size: %d", size)
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

func (m *DecodeBuf) FlaggedVectorInt(flags, num int32) []int32 {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.VectorInt()
}

func (m *DecodeBuf) VectorLong() []int64 {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != CRC_vector {
		m.err = merry.Errorf("DecodeVectorLong: wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = merry.Errorf("DecodeVectorLong: negative size: %d", size)
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

func (m *DecodeBuf) FlaggedVectorLong(flags, num int32) []int64 {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.VectorLong()
}

func (m *DecodeBuf) VectorString() []string {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != CRC_vector {
		m.err = merry.Errorf("DecodeVectorString: wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = merry.Errorf("DecodeVectorString: negative size: %d", size)
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

func (m *DecodeBuf) FlaggedVectorString(flags, num int32) []string {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.VectorString()
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

func (m *DecodeBuf) Vector() []TL {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	if constructor != CRC_vector {
		m.err = merry.Errorf("DecodeVector: wrong constructor (0x%08x)", constructor)
		return nil
	}
	size := m.Int()
	if m.err != nil {
		return nil
	}
	if size < 0 {
		m.err = merry.Errorf("DecodeVector: negative size: %d", size)
		return nil
	}
	x := make([]TL, size)
	i := int32(0)
	for i < size {
		y := m.Object()
		if m.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (m *DecodeBuf) FlaggedVector(flags, num int32) []TL {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.Vector()
}

func (m *DecodeBuf) Object() TL {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}
	return m.ObjectGenerated(constructor)
}

func (m *DecodeBuf) FlaggedObject(flags, num int32) TL {
	bit := int32(1 << uint(num))
	if flags&bit == 0 {
		return nil
	}
	return m.Object()
}

func (d *DecodeBuf) dump() {
	fmt.Println(hex.Dump(d.buf[d.off:d.size]))
}

func toBool(x TL) bool {
	_, ok := x.(TL_boolTrue)
	return ok
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
//   account.getWallPapers#c04cfac2 = Vector<WallPaper>
//   messages.getMessagesViews#c4c8a55d ... = Vector<int>
//   photos.deletePhotos#87cf7f2f ... = Vector<long>
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
		arr := make([]TL_MT_message, size)
		for i := int32(0); i < size; i++ {
			arr[i] = TL_MT_message{dbuf.Long(), dbuf.Int(), dbuf.Int(), m.decodeMessage(dbuf, reqMsg)}
			if dbuf.err != nil {
				return nil
			}
		}
		r = TL_msg_container{arr}

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
		r = TL_rpc_result{requestID, r}

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
		if reqMsg == nil || constructor == CRC_rpc_error {
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
