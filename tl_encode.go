package mtproto

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
	"time"
)

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	return ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
}

type EncodeBuf struct {
	buf []byte
}

func NewEncodeBuf(cap int) *EncodeBuf {
	return &EncodeBuf{make([]byte, 0, cap)}
}

func (e *EncodeBuf) Int(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *EncodeBuf) UInt(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *EncodeBuf) Long(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *EncodeBuf) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
}

func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

func (e *EncodeBuf) StringBytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)

	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)

	}
	e.buf = append(e.buf, res...)
}

func (e *EncodeBuf) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *EncodeBuf) VectorInt(v []int32) {
	x := make([]byte, 4+4+len(v)*4)
	binary.LittleEndian.PutUint32(x, CRC_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint32(x[i:], uint32(v))
		i += 4
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorLong(v []int64) {
	x := make([]byte, 4+4+len(v)*8)
	binary.LittleEndian.PutUint32(x, CRC_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	i := 8
	for _, v := range v {
		binary.LittleEndian.PutUint64(x[i:], uint64(v))
		i += 8
	}
	e.buf = append(e.buf, x...)
}

func (e *EncodeBuf) VectorString(v []string) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, CRC_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.String(v)
	}
}

func (e *EncodeBuf) Vector(v []TL) {
	x := make([]byte, 8)
	binary.LittleEndian.PutUint32(x, CRC_vector)
	binary.LittleEndian.PutUint32(x[4:], uint32(len(v)))
	e.buf = append(e.buf, x...)
	for _, v := range v {
		e.buf = append(e.buf, v.encode()...)
	}
}
