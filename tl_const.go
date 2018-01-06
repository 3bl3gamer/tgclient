package mtproto

const (
	layer = 71

	// https://core.telegram.org/schema/mtproto
	CRC_vector                     = 0x1cb5c415
	CRC_resPQ                      = 0x05162463
	CRC_p_q_inner_data             = 0x83c95aec
	CRC_server_DH_params_fail      = 0x79cb045d
	CRC_server_DH_params_ok        = 0xd0e8075c
	CRC_server_DH_inner_data       = 0xb5890dba
	CRC_client_DH_inner_data       = 0x6643b654
	CRC_dh_gen_ok                  = 0x3bcbf734
	CRC_dh_gen_retry               = 0x46dc1fb9
	CRC_dh_gen_fail                = 0xa69dae02
	CRC_rpc_result                 = 0xf35c6d01
	CRC_rpc_error                  = 0x2144ca19
	CRC_rpc_answer_unknown         = 0x5e2ad36e
	CRC_rpc_answer_dropped_running = 0xcd78e586
	CRC_rpc_answer_dropped         = 0xa43ad8b7
	CRC_future_salt                = 0x0949d9dc
	CRC_future_salts               = 0xae500895
	CRC_pong                       = 0x347773c5
	CRC_destroy_session_ok         = 0xe22045fc
	CRC_destroy_session_none       = 0x62d350c9
	CRC_new_session_created        = 0x9ec20908
	CRC_msg_container              = 0x73f1f8dc
	CRC_msg_copy                   = 0xe06046b2
	CRC_gzip_packed                = 0x3072cfa1
	CRC_msgs_ack                   = 0x62d6b459
	CRC_bad_msg_notification       = 0xa7eff811
	CRC_bad_server_salt            = 0xedab447b
	CRC_msg_resend_req             = 0x7d861a08
	CRC_msgs_state_req             = 0xda69fb52
	CRC_msgs_state_info            = 0x04deb57d
	CRC_msgs_all_info              = 0x8cc0d131
	CRC_msg_detailed_info          = 0x276d3ec6
	CRC_msg_new_detailed_info      = 0x809db6df
	CRC_req_pq                     = 0x60469778
	CRC_req_DH_params              = 0xd712e4be
	CRC_set_client_DH_params       = 0xf5045f1f
	CRC_rpc_drop_answer            = 0x58e4a740
	CRC_get_future_salts           = 0xb921bd04
	CRC_ping                       = 0x7abe77ec
	CRC_ping_delay_disconnect      = 0xf3427b8c
	CRC_destroy_session            = 0xe7512126
	CRC_http_wait                  = 0x9299359f
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
