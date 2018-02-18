package mtproto

type TL interface {
	encode() []byte
}

type TL_msg_container struct {
	Items []TL_MT_message
}

func (e TL_msg_container) encode() []byte { return nil }

type TL_MT_message struct {
	MsgID int64
	SeqNo int32
	Size  int32
	Data  TL
}

type TL_rpc_result struct {
	req_msg_id int64
	obj        TL
}

func (e TL_rpc_result) encode() []byte { return nil }
