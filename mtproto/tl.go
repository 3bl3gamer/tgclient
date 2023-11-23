package mtproto

type TL interface {
	encode() []byte
}

type TLReq interface {
	TL
	decodeResponse(*DecodeBuf) TL
}

type TL_msgContainer struct {
	Items []TL_mtMessage
}

func (e TL_msgContainer) encode() []byte { return nil }

type TL_mtMessage struct {
	MsgID int64
	SeqNo int32
	Size  int32
	Data  TL
}

type TL_rpcResult struct {
	reqMsgID int64
	obj      TL
}

func (e TL_rpcResult) encode() []byte { return nil }

type VectorInt []int32

func (e VectorInt) encode() []byte { return nil }

type VectorLong []int64

func (e VectorLong) encode() []byte { return nil }

type VectorObject []TL

func (e VectorObject) encode() []byte { return nil }
