package mtproto

const (
	CRC_vector        = 0x1cb5c415
	CRC_rpc_result    = 0xf35c6d01
	CRC_msg_container = 0x73f1f8dc
	CRC_gzip_packed   = 0x3072cfa1
)

const (
	TL_ErrSeeOther     = int32(303)
	TL_ErrBadRequest   = int32(400)
	TL_ErrUnauthorized = int32(401)
	TL_ErrForbidden    = int32(403)
	TL_ErrNotFound     = int32(404)
	TL_ErrFlood        = int32(420)
	TL_ErrInterbal     = int32(500)
)
