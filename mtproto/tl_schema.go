package mtproto

import (
	"github.com/ansel1/merry"
)

const (
	TL_Layer                                                              = 114
	CRC_resPQ                                                             = 0x05162463
	CRC_p_q_inner_data                                                    = 0x83c95aec
	CRC_p_q_inner_data_dc                                                 = 0xa9f55f95
	CRC_p_q_inner_data_temp                                               = 0x3c6a84d4
	CRC_p_q_inner_data_temp_dc                                            = 0x56fddf88
	CRC_bind_auth_key_inner                                               = 0x75a3f765
	CRC_server_DH_params_fail                                             = 0x79cb045d
	CRC_server_DH_params_ok                                               = 0xd0e8075c
	CRC_server_DH_inner_data                                              = 0xb5890dba
	CRC_client_DH_inner_data                                              = 0x6643b654
	CRC_dh_gen_ok                                                         = 0x3bcbf734
	CRC_dh_gen_retry                                                      = 0x46dc1fb9
	CRC_dh_gen_fail                                                       = 0xa69dae02
	CRC_destroy_auth_key_ok                                               = 0xf660e1d4
	CRC_destroy_auth_key_none                                             = 0x0a9f2259
	CRC_destroy_auth_key_fail                                             = 0xea109b13
	CRC_req_pq                                                            = 0x60469778
	CRC_req_pq_multi                                                      = 0xbe7e8ef1
	CRC_req_DH_params                                                     = 0xd712e4be
	CRC_set_client_DH_params                                              = 0xf5045f1f
	CRC_destroy_auth_key                                                  = 0xd1435160
	CRC_msgs_ack                                                          = 0x62d6b459
	CRC_bad_msg_notification                                              = 0xa7eff811
	CRC_bad_server_salt                                                   = 0xedab447b
	CRC_msgs_state_req                                                    = 0xda69fb52
	CRC_msgs_state_info                                                   = 0x04deb57d
	CRC_msgs_all_info                                                     = 0x8cc0d131
	CRC_msg_detailed_info                                                 = 0x276d3ec6
	CRC_msg_new_detailed_info                                             = 0x809db6df
	CRC_msg_resend_req                                                    = 0x7d861a08
	CRC_rpc_error                                                         = 0x2144ca19
	CRC_rpc_answer_unknown                                                = 0x5e2ad36e
	CRC_rpc_answer_dropped_running                                        = 0xcd78e586
	CRC_rpc_answer_dropped                                                = 0xa43ad8b7
	CRC_future_salt                                                       = 0x0949d9dc
	CRC_future_salts                                                      = 0xae500895
	CRC_pong                                                              = 0x347773c5
	CRC_destroy_session_ok                                                = 0xe22045fc
	CRC_destroy_session_none                                              = 0x62d350c9
	CRC_new_session_created                                               = 0x9ec20908
	CRC_http_wait                                                         = 0x9299359f
	CRC_ipPort                                                            = 0xd433ad73
	CRC_ipPortSecret                                                      = 0x37982646
	CRC_accessPointRule                                                   = 0x4679b65f
	CRC_help_configSimple                                                 = 0x5a592a6c
	CRC_tlsClientHello                                                    = 0x6c52c484
	CRC_tlsBlockString                                                    = 0x4218a164
	CRC_tlsBlockRandom                                                    = 0x4d4dc41e
	CRC_tlsBlockZero                                                      = 0x09333afb
	CRC_tlsBlockDomain                                                    = 0x10e8636f
	CRC_tlsBlockGrease                                                    = 0xe675a1c1
	CRC_tlsBlockPublicKey                                                 = 0x9eb95b5c
	CRC_tlsBlockScope                                                     = 0xe725d44f
	CRC_rpc_drop_answer                                                   = 0x58e4a740
	CRC_get_future_salts                                                  = 0xb921bd04
	CRC_ping                                                              = 0x7abe77ec
	CRC_ping_delay_disconnect                                             = 0xf3427b8c
	CRC_destroy_session                                                   = 0xe7512126
	CRC_boolFalse                                                         = 0xbc799737
	CRC_boolTrue                                                          = 0x997275b5
	CRC_true                                                              = 0x3fedd339
	CRC_error                                                             = 0xc4b9f9bb
	CRC_null                                                              = 0x56730bcc
	CRC_inputPeerEmpty                                                    = 0x7f3b18ea
	CRC_inputPeerSelf                                                     = 0x7da07ec9
	CRC_inputPeerChat                                                     = 0x179be863
	CRC_inputPeerUser                                                     = 0x7b8e7de6
	CRC_inputPeerChannel                                                  = 0x20adaef8
	CRC_inputPeerUserFromMessage                                          = 0x17bae2e6
	CRC_inputPeerChannelFromMessage                                       = 0x9c95f7bb
	CRC_inputUserEmpty                                                    = 0xb98886cf
	CRC_inputUserSelf                                                     = 0xf7c1b13f
	CRC_inputUser                                                         = 0xd8292816
	CRC_inputUserFromMessage                                              = 0x2d117597
	CRC_inputPhoneContact                                                 = 0xf392b7f4
	CRC_inputFile                                                         = 0xf52ff27f
	CRC_inputFileBig                                                      = 0xfa4f0bb5
	CRC_inputMediaEmpty                                                   = 0x9664f57f
	CRC_inputMediaUploadedPhoto                                           = 0x1e287d04
	CRC_inputMediaPhoto                                                   = 0xb3ba0635
	CRC_inputMediaGeoPoint                                                = 0xf9c44144
	CRC_inputMediaContact                                                 = 0xf8ab7dfb
	CRC_inputMediaUploadedDocument                                        = 0x5b38c6c1
	CRC_inputMediaDocument                                                = 0x23ab23d2
	CRC_inputMediaVenue                                                   = 0xc13d1c11
	CRC_inputMediaGifExternal                                             = 0x4843b0fd
	CRC_inputMediaPhotoExternal                                           = 0xe5bbfe1a
	CRC_inputMediaDocumentExternal                                        = 0xfb52dc99
	CRC_inputMediaGame                                                    = 0xd33f43f3
	CRC_inputMediaInvoice                                                 = 0xf4e096c3
	CRC_inputMediaGeoLive                                                 = 0xce4e82fd
	CRC_inputMediaPoll                                                    = 0x0f94e5f1
	CRC_inputMediaDice                                                    = 0xe66fbf7b
	CRC_inputChatPhotoEmpty                                               = 0x1ca48f57
	CRC_inputChatUploadedPhoto                                            = 0x927c55b4
	CRC_inputChatPhoto                                                    = 0x8953ad37
	CRC_inputGeoPointEmpty                                                = 0xe4c123d6
	CRC_inputGeoPoint                                                     = 0xf3b7acc9
	CRC_inputPhotoEmpty                                                   = 0x1cd7bf0d
	CRC_inputPhoto                                                        = 0x3bb3b94a
	CRC_inputFileLocation                                                 = 0xdfdaabe1
	CRC_inputEncryptedFileLocation                                        = 0xf5235d55
	CRC_inputDocumentFileLocation                                         = 0xbad07584
	CRC_inputSecureFileLocation                                           = 0xcbc7ee28
	CRC_inputTakeoutFileLocation                                          = 0x29be5899
	CRC_inputPhotoFileLocation                                            = 0x40181ffe
	CRC_inputPhotoLegacyFileLocation                                      = 0xd83466f3
	CRC_inputPeerPhotoFileLocation                                        = 0x27d69997
	CRC_inputStickerSetThumb                                              = 0x0dbaeae9
	CRC_peerUser                                                          = 0x9db1bc6d
	CRC_peerChat                                                          = 0xbad0e5bb
	CRC_peerChannel                                                       = 0xbddde532
	CRC_storage_fileUnknown                                               = 0xaa963b05
	CRC_storage_filePartial                                               = 0x40bc6f52
	CRC_storage_fileJpeg                                                  = 0x007efe0e
	CRC_storage_fileGif                                                   = 0xcae1aadf
	CRC_storage_filePng                                                   = 0x0a4f63c0
	CRC_storage_filePdf                                                   = 0xae1e508d
	CRC_storage_fileMp3                                                   = 0x528a0677
	CRC_storage_fileMov                                                   = 0x4b09ebbc
	CRC_storage_fileMp4                                                   = 0xb3cea0e4
	CRC_storage_fileWebp                                                  = 0x1081464c
	CRC_userEmpty                                                         = 0x200250ba
	CRC_user                                                              = 0x938458c1
	CRC_userProfilePhotoEmpty                                             = 0x4f11bae1
	CRC_userProfilePhoto                                                  = 0xecd75d8c
	CRC_userStatusEmpty                                                   = 0x09d05049
	CRC_userStatusOnline                                                  = 0xedb93949
	CRC_userStatusOffline                                                 = 0x008c703f
	CRC_userStatusRecently                                                = 0xe26f42f1
	CRC_userStatusLastWeek                                                = 0x07bf09fc
	CRC_userStatusLastMonth                                               = 0x77ebc742
	CRC_chatEmpty                                                         = 0x9ba2d800
	CRC_chat                                                              = 0x3bda1bde
	CRC_chatForbidden                                                     = 0x07328bdb
	CRC_channel                                                           = 0xd31a961e
	CRC_channelForbidden                                                  = 0x289da732
	CRC_chatFull                                                          = 0x1b7c9db3
	CRC_channelFull                                                       = 0xf0e6672a
	CRC_chatParticipant                                                   = 0xc8d7493e
	CRC_chatParticipantCreator                                            = 0xda13538a
	CRC_chatParticipantAdmin                                              = 0xe2d6e436
	CRC_chatParticipantsForbidden                                         = 0xfc900c2b
	CRC_chatParticipants                                                  = 0x3f460fed
	CRC_chatPhotoEmpty                                                    = 0x37c1011c
	CRC_chatPhoto                                                         = 0x475cdbd5
	CRC_messageEmpty                                                      = 0x83e5de54
	CRC_message                                                           = 0x452c0e65
	CRC_messageService                                                    = 0x9e19a1f6
	CRC_messageMediaEmpty                                                 = 0x3ded6320
	CRC_messageMediaPhoto                                                 = 0x695150d7
	CRC_messageMediaGeo                                                   = 0x56e0d474
	CRC_messageMediaContact                                               = 0xcbf24940
	CRC_messageMediaUnsupported                                           = 0x9f84f49e
	CRC_messageMediaDocument                                              = 0x9cb070d7
	CRC_messageMediaWebPage                                               = 0xa32dd600
	CRC_messageMediaVenue                                                 = 0x2ec0533f
	CRC_messageMediaGame                                                  = 0xfdb19008
	CRC_messageMediaInvoice                                               = 0x84551347
	CRC_messageMediaGeoLive                                               = 0x7c3c2609
	CRC_messageMediaPoll                                                  = 0x4bd6e798
	CRC_messageMediaDice                                                  = 0x3f7ee58b
	CRC_messageActionEmpty                                                = 0xb6aef7b0
	CRC_messageActionChatCreate                                           = 0xa6638b9a
	CRC_messageActionChatEditTitle                                        = 0xb5a1ce5a
	CRC_messageActionChatEditPhoto                                        = 0x7fcb13a8
	CRC_messageActionChatDeletePhoto                                      = 0x95e3fbef
	CRC_messageActionChatAddUser                                          = 0x488a7337
	CRC_messageActionChatDeleteUser                                       = 0xb2ae9b0c
	CRC_messageActionChatJoinedByLink                                     = 0xf89cf5e8
	CRC_messageActionChannelCreate                                        = 0x95d2ac92
	CRC_messageActionChatMigrateTo                                        = 0x51bdb021
	CRC_messageActionChannelMigrateFrom                                   = 0xb055eaee
	CRC_messageActionPinMessage                                           = 0x94bd38ed
	CRC_messageActionHistoryClear                                         = 0x9fbab604
	CRC_messageActionGameScore                                            = 0x92a72876
	CRC_messageActionPaymentSentMe                                        = 0x8f31b327
	CRC_messageActionPaymentSent                                          = 0x40699cd0
	CRC_messageActionPhoneCall                                            = 0x80e11a7f
	CRC_messageActionScreenshotTaken                                      = 0x4792929b
	CRC_messageActionCustomAction                                         = 0xfae69f56
	CRC_messageActionBotAllowed                                           = 0xabe9affe
	CRC_messageActionSecureValuesSentMe                                   = 0x1b287353
	CRC_messageActionSecureValuesSent                                     = 0xd95c6154
	CRC_messageActionContactSignUp                                        = 0xf3f25f76
	CRC_dialog                                                            = 0x2c171f72
	CRC_dialogFolder                                                      = 0x71bd134c
	CRC_photoEmpty                                                        = 0x2331b22d
	CRC_photo                                                             = 0xd07504a5
	CRC_photoSizeEmpty                                                    = 0x0e17e23c
	CRC_photoSize                                                         = 0x77bfb61b
	CRC_photoCachedSize                                                   = 0xe9a734fa
	CRC_photoStrippedSize                                                 = 0xe0b0bc2e
	CRC_geoPointEmpty                                                     = 0x1117dd5f
	CRC_geoPoint                                                          = 0x0296f104
	CRC_auth_sentCode                                                     = 0x5e002502
	CRC_auth_authorization                                                = 0xcd050916
	CRC_auth_authorizationSignUpRequired                                  = 0x44747e9a
	CRC_auth_exportedAuthorization                                        = 0xdf969c2d
	CRC_inputNotifyPeer                                                   = 0xb8bc5b0c
	CRC_inputNotifyUsers                                                  = 0x193b4417
	CRC_inputNotifyChats                                                  = 0x4a95e84e
	CRC_inputNotifyBroadcasts                                             = 0xb1db7c7e
	CRC_inputPeerNotifySettings                                           = 0x9c3d198e
	CRC_peerNotifySettings                                                = 0xaf509d20
	CRC_peerSettings                                                      = 0x818426cd
	CRC_wallPaper                                                         = 0xa437c3ed
	CRC_wallPaperNoFile                                                   = 0x8af40b25
	CRC_inputReportReasonSpam                                             = 0x58dbcab8
	CRC_inputReportReasonViolence                                         = 0x1e22c78d
	CRC_inputReportReasonPornography                                      = 0x2e59d922
	CRC_inputReportReasonChildAbuse                                       = 0xadf44ee3
	CRC_inputReportReasonOther                                            = 0xe1746d0a
	CRC_inputReportReasonCopyright                                        = 0x9b89f93a
	CRC_inputReportReasonGeoIrrelevant                                    = 0xdbd4feed
	CRC_userFull                                                          = 0xedf17c12
	CRC_contact                                                           = 0xf911c994
	CRC_importedContact                                                   = 0xd0028438
	CRC_contactBlocked                                                    = 0x561bc879
	CRC_contactStatus                                                     = 0xd3680c61
	CRC_contacts_contactsNotModified                                      = 0xb74ba9d2
	CRC_contacts_contacts                                                 = 0xeae87e42
	CRC_contacts_importedContacts                                         = 0x77d01c3b
	CRC_contacts_blocked                                                  = 0x1c138d15
	CRC_contacts_blockedSlice                                             = 0x900802a1
	CRC_messages_dialogs                                                  = 0x15ba6c40
	CRC_messages_dialogsSlice                                             = 0x71e094f3
	CRC_messages_dialogsNotModified                                       = 0xf0e3e596
	CRC_messages_messages                                                 = 0x8c718e87
	CRC_messages_messagesSlice                                            = 0xc8edce1e
	CRC_messages_channelMessages                                          = 0x99262e37
	CRC_messages_messagesNotModified                                      = 0x74535f21
	CRC_messages_chats                                                    = 0x64ff9fd5
	CRC_messages_chatsSlice                                               = 0x9cd81144
	CRC_messages_chatFull                                                 = 0xe5d7d19c
	CRC_messages_affectedHistory                                          = 0xb45c69d1
	CRC_inputMessagesFilterEmpty                                          = 0x57e2f66c
	CRC_inputMessagesFilterPhotos                                         = 0x9609a51c
	CRC_inputMessagesFilterVideo                                          = 0x9fc00e65
	CRC_inputMessagesFilterPhotoVideo                                     = 0x56e9f0e4
	CRC_inputMessagesFilterDocument                                       = 0x9eddf188
	CRC_inputMessagesFilterUrl                                            = 0x7ef0dd87
	CRC_inputMessagesFilterGif                                            = 0xffc86587
	CRC_inputMessagesFilterVoice                                          = 0x50f5c392
	CRC_inputMessagesFilterMusic                                          = 0x3751b49e
	CRC_inputMessagesFilterChatPhotos                                     = 0x3a20ecb8
	CRC_inputMessagesFilterPhoneCalls                                     = 0x80c99768
	CRC_inputMessagesFilterRoundVoice                                     = 0x7a7c17a4
	CRC_inputMessagesFilterRoundVideo                                     = 0xb549da53
	CRC_inputMessagesFilterMyMentions                                     = 0xc1f8e69a
	CRC_inputMessagesFilterGeo                                            = 0xe7026d0d
	CRC_inputMessagesFilterContacts                                       = 0xe062db83
	CRC_updateNewMessage                                                  = 0x1f2b0afd
	CRC_updateMessageID                                                   = 0x4e90bfd6
	CRC_updateDeleteMessages                                              = 0xa20db0e5
	CRC_updateUserTyping                                                  = 0x5c486927
	CRC_updateChatUserTyping                                              = 0x9a65ea1f
	CRC_updateChatParticipants                                            = 0x07761198
	CRC_updateUserStatus                                                  = 0x1bfbd823
	CRC_updateUserName                                                    = 0xa7332b73
	CRC_updateUserPhoto                                                   = 0x95313b0c
	CRC_updateNewEncryptedMessage                                         = 0x12bcbd9a
	CRC_updateEncryptedChatTyping                                         = 0x1710f156
	CRC_updateEncryption                                                  = 0xb4a2e88d
	CRC_updateEncryptedMessagesRead                                       = 0x38fe25b7
	CRC_updateChatParticipantAdd                                          = 0xea4b0e5c
	CRC_updateChatParticipantDelete                                       = 0x6e5f8c22
	CRC_updateDcOptions                                                   = 0x8e5e9873
	CRC_updateUserBlocked                                                 = 0x80ece81a
	CRC_updateNotifySettings                                              = 0xbec268ef
	CRC_updateServiceNotification                                         = 0xebe46819
	CRC_updatePrivacy                                                     = 0xee3b272a
	CRC_updateUserPhone                                                   = 0x12b9417b
	CRC_updateReadHistoryInbox                                            = 0x9c974fdf
	CRC_updateReadHistoryOutbox                                           = 0x2f2f21bf
	CRC_updateWebPage                                                     = 0x7f891213
	CRC_updateReadMessagesContents                                        = 0x68c13933
	CRC_updateChannelTooLong                                              = 0xeb0467fb
	CRC_updateChannel                                                     = 0xb6d45656
	CRC_updateNewChannelMessage                                           = 0x62ba04d9
	CRC_updateReadChannelInbox                                            = 0x330b5424
	CRC_updateDeleteChannelMessages                                       = 0xc37521c9
	CRC_updateChannelMessageViews                                         = 0x98a12b4b
	CRC_updateChatParticipantAdmin                                        = 0xb6901959
	CRC_updateNewStickerSet                                               = 0x688a30aa
	CRC_updateStickerSetsOrder                                            = 0x0bb2d201
	CRC_updateStickerSets                                                 = 0x43ae3dec
	CRC_updateSavedGifs                                                   = 0x9375341e
	CRC_updateBotInlineQuery                                              = 0x54826690
	CRC_updateBotInlineSend                                               = 0x0e48f964
	CRC_updateEditChannelMessage                                          = 0x1b3f4df7
	CRC_updateChannelPinnedMessage                                        = 0x98592475
	CRC_updateBotCallbackQuery                                            = 0xe73547e1
	CRC_updateEditMessage                                                 = 0xe40370a3
	CRC_updateInlineBotCallbackQuery                                      = 0xf9d27a5a
	CRC_updateReadChannelOutbox                                           = 0x25d6c9c7
	CRC_updateDraftMessage                                                = 0xee2bb969
	CRC_updateReadFeaturedStickers                                        = 0x571d2742
	CRC_updateRecentStickers                                              = 0x9a422c20
	CRC_updateConfig                                                      = 0xa229dd06
	CRC_updatePtsChanged                                                  = 0x3354678f
	CRC_updateChannelWebPage                                              = 0x40771900
	CRC_updateDialogPinned                                                = 0x6e6fe51c
	CRC_updatePinnedDialogs                                               = 0xfa0f3ca2
	CRC_updateBotWebhookJSON                                              = 0x8317c0c3
	CRC_updateBotWebhookJSONQuery                                         = 0x9b9240a6
	CRC_updateBotShippingQuery                                            = 0xe0cdc940
	CRC_updateBotPrecheckoutQuery                                         = 0x5d2f3aa9
	CRC_updatePhoneCall                                                   = 0xab0f6b1e
	CRC_updateLangPackTooLong                                             = 0x46560264
	CRC_updateLangPack                                                    = 0x56022f4d
	CRC_updateFavedStickers                                               = 0xe511996d
	CRC_updateChannelReadMessagesContents                                 = 0x89893b45
	CRC_updateContactsReset                                               = 0x7084a7be
	CRC_updateChannelAvailableMessages                                    = 0x70db6837
	CRC_updateDialogUnreadMark                                            = 0xe16459c3
	CRC_updateUserPinnedMessage                                           = 0x4c43da18
	CRC_updateChatPinnedMessage                                           = 0xe10db349
	CRC_updateMessagePoll                                                 = 0xaca1657b
	CRC_updateChatDefaultBannedRights                                     = 0x54c01850
	CRC_updateFolderPeers                                                 = 0x19360dc0
	CRC_updatePeerSettings                                                = 0x6a7e7366
	CRC_updatePeerLocated                                                 = 0xb4afcfb0
	CRC_updateNewScheduledMessage                                         = 0x39a51dfb
	CRC_updateDeleteScheduledMessages                                     = 0x90866cee
	CRC_updateTheme                                                       = 0x8216fba3
	CRC_updateGeoLiveViewed                                               = 0x871fb939
	CRC_updateLoginToken                                                  = 0x564fe691
	CRC_updateMessagePollVote                                             = 0x42f88f2c
	CRC_updateDialogFilter                                                = 0x26ffde7d
	CRC_updateDialogFilterOrder                                           = 0xa5d72105
	CRC_updateDialogFilters                                               = 0x3504914f
	CRC_updatePhoneCallSignalingData                                      = 0x2661bf09
	CRC_updates_state                                                     = 0xa56c2a3e
	CRC_updates_differenceEmpty                                           = 0x5d75a138
	CRC_updates_difference                                                = 0x00f49ca0
	CRC_updates_differenceSlice                                           = 0xa8fb1981
	CRC_updates_differenceTooLong                                         = 0x4afe8f6d
	CRC_updatesTooLong                                                    = 0xe317af7e
	CRC_updateShortMessage                                                = 0x914fbf11
	CRC_updateShortChatMessage                                            = 0x16812688
	CRC_updateShort                                                       = 0x78d4dec1
	CRC_updatesCombined                                                   = 0x725b04c3
	CRC_updates                                                           = 0x74ae4240
	CRC_updateShortSentMessage                                            = 0x11f1331c
	CRC_photos_photos                                                     = 0x8dca6aa5
	CRC_photos_photosSlice                                                = 0x15051f54
	CRC_photos_photo                                                      = 0x20212ca8
	CRC_upload_file                                                       = 0x096a18d5
	CRC_upload_fileCdnRedirect                                            = 0xf18cda44
	CRC_dcOption                                                          = 0x18b7a10d
	CRC_config                                                            = 0x330b4067
	CRC_nearestDc                                                         = 0x8e1a1775
	CRC_help_appUpdate                                                    = 0x1da7158f
	CRC_help_noAppUpdate                                                  = 0xc45a6536
	CRC_help_inviteText                                                   = 0x18cb9f78
	CRC_encryptedChatEmpty                                                = 0xab7ec0a0
	CRC_encryptedChatWaiting                                              = 0x3bf703dc
	CRC_encryptedChatRequested                                            = 0xc878527e
	CRC_encryptedChat                                                     = 0xfa56ce36
	CRC_encryptedChatDiscarded                                            = 0x13d6dd27
	CRC_inputEncryptedChat                                                = 0xf141b5e1
	CRC_encryptedFileEmpty                                                = 0xc21f497e
	CRC_encryptedFile                                                     = 0x4a70994c
	CRC_inputEncryptedFileEmpty                                           = 0x1837c364
	CRC_inputEncryptedFileUploaded                                        = 0x64bd0306
	CRC_inputEncryptedFile                                                = 0x5a17b5e5
	CRC_inputEncryptedFileBigUploaded                                     = 0x2dc173c8
	CRC_encryptedMessage                                                  = 0xed18c118
	CRC_encryptedMessageService                                           = 0x23734b06
	CRC_messages_dhConfigNotModified                                      = 0xc0e24635
	CRC_messages_dhConfig                                                 = 0x2c221edd
	CRC_messages_sentEncryptedMessage                                     = 0x560f8935
	CRC_messages_sentEncryptedFile                                        = 0x9493ff32
	CRC_inputDocumentEmpty                                                = 0x72f0eaae
	CRC_inputDocument                                                     = 0x1abfb575
	CRC_documentEmpty                                                     = 0x36f8c871
	CRC_document                                                          = 0x1e87342b
	CRC_help_support                                                      = 0x17c6b5f6
	CRC_notifyPeer                                                        = 0x9fd40bd8
	CRC_notifyUsers                                                       = 0xb4c83b4c
	CRC_notifyChats                                                       = 0xc007cec3
	CRC_notifyBroadcasts                                                  = 0xd612e8ef
	CRC_sendMessageTypingAction                                           = 0x16bf744e
	CRC_sendMessageCancelAction                                           = 0xfd5ec8f5
	CRC_sendMessageRecordVideoAction                                      = 0xa187d66f
	CRC_sendMessageUploadVideoAction                                      = 0xe9763aec
	CRC_sendMessageRecordAudioAction                                      = 0xd52f73f7
	CRC_sendMessageUploadAudioAction                                      = 0xf351d7ab
	CRC_sendMessageUploadPhotoAction                                      = 0xd1d34a26
	CRC_sendMessageUploadDocumentAction                                   = 0xaa0cd9e4
	CRC_sendMessageGeoLocationAction                                      = 0x176f8ba1
	CRC_sendMessageChooseContactAction                                    = 0x628cbc6f
	CRC_sendMessageGamePlayAction                                         = 0xdd6a8f48
	CRC_sendMessageRecordRoundAction                                      = 0x88f27fbc
	CRC_sendMessageUploadRoundAction                                      = 0x243e1c66
	CRC_contacts_found                                                    = 0xb3134d9d
	CRC_inputPrivacyKeyStatusTimestamp                                    = 0x4f96cb18
	CRC_inputPrivacyKeyChatInvite                                         = 0xbdfb0426
	CRC_inputPrivacyKeyPhoneCall                                          = 0xfabadc5f
	CRC_inputPrivacyKeyPhoneP2P                                           = 0xdb9e70d2
	CRC_inputPrivacyKeyForwards                                           = 0xa4dd4c08
	CRC_inputPrivacyKeyProfilePhoto                                       = 0x5719bacc
	CRC_inputPrivacyKeyPhoneNumber                                        = 0x0352dafa
	CRC_inputPrivacyKeyAddedByPhone                                       = 0xd1219bdd
	CRC_privacyKeyStatusTimestamp                                         = 0xbc2eab30
	CRC_privacyKeyChatInvite                                              = 0x500e6dfa
	CRC_privacyKeyPhoneCall                                               = 0x3d662b7b
	CRC_privacyKeyPhoneP2P                                                = 0x39491cc8
	CRC_privacyKeyForwards                                                = 0x69ec56a3
	CRC_privacyKeyProfilePhoto                                            = 0x96151fed
	CRC_privacyKeyPhoneNumber                                             = 0xd19ae46d
	CRC_privacyKeyAddedByPhone                                            = 0x42ffd42b
	CRC_inputPrivacyValueAllowContacts                                    = 0x0d09e07b
	CRC_inputPrivacyValueAllowAll                                         = 0x184b35ce
	CRC_inputPrivacyValueAllowUsers                                       = 0x131cc67f
	CRC_inputPrivacyValueDisallowContacts                                 = 0x0ba52007
	CRC_inputPrivacyValueDisallowAll                                      = 0xd66b66c9
	CRC_inputPrivacyValueDisallowUsers                                    = 0x90110467
	CRC_inputPrivacyValueAllowChatParticipants                            = 0x4c81c1ba
	CRC_inputPrivacyValueDisallowChatParticipants                         = 0xd82363af
	CRC_privacyValueAllowContacts                                         = 0xfffe1bac
	CRC_privacyValueAllowAll                                              = 0x65427b82
	CRC_privacyValueAllowUsers                                            = 0x4d5bbe0c
	CRC_privacyValueDisallowContacts                                      = 0xf888fa1a
	CRC_privacyValueDisallowAll                                           = 0x8b73e763
	CRC_privacyValueDisallowUsers                                         = 0x0c7f49b7
	CRC_privacyValueAllowChatParticipants                                 = 0x18be796b
	CRC_privacyValueDisallowChatParticipants                              = 0xacae0690
	CRC_account_privacyRules                                              = 0x50a04e45
	CRC_accountDaysTTL                                                    = 0xb8d0afdf
	CRC_documentAttributeImageSize                                        = 0x6c37c15c
	CRC_documentAttributeAnimated                                         = 0x11b58939
	CRC_documentAttributeSticker                                          = 0x6319d612
	CRC_documentAttributeVideo                                            = 0x0ef02ce6
	CRC_documentAttributeAudio                                            = 0x9852f9c6
	CRC_documentAttributeFilename                                         = 0x15590068
	CRC_documentAttributeHasStickers                                      = 0x9801d2f7
	CRC_messages_stickersNotModified                                      = 0xf1749a22
	CRC_messages_stickers                                                 = 0xe4599bbd
	CRC_stickerPack                                                       = 0x12b299d4
	CRC_messages_allStickersNotModified                                   = 0xe86602c3
	CRC_messages_allStickers                                              = 0xedfd405f
	CRC_messages_affectedMessages                                         = 0x84d19185
	CRC_webPageEmpty                                                      = 0xeb1477e8
	CRC_webPagePending                                                    = 0xc586da1c
	CRC_webPage                                                           = 0xe89c45b2
	CRC_webPageNotModified                                                = 0x7311ca11
	CRC_authorization                                                     = 0xad01d61d
	CRC_account_authorizations                                            = 0x1250abde
	CRC_account_password                                                  = 0xad2641f8
	CRC_account_passwordSettings                                          = 0x9a5c33e5
	CRC_account_passwordInputSettings                                     = 0xc23727c9
	CRC_auth_passwordRecovery                                             = 0x137948a5
	CRC_receivedNotifyMessage                                             = 0xa384b779
	CRC_chatInviteEmpty                                                   = 0x69df3769
	CRC_chatInviteExported                                                = 0xfc2e05bc
	CRC_chatInviteAlready                                                 = 0x5a686d7c
	CRC_chatInvite                                                        = 0xdfc2f58e
	CRC_inputStickerSetEmpty                                              = 0xffb62b95
	CRC_inputStickerSetID                                                 = 0x9de7a269
	CRC_inputStickerSetShortName                                          = 0x861cc8a0
	CRC_inputStickerSetAnimatedEmoji                                      = 0x028703c8
	CRC_inputStickerSetDice                                               = 0xe67f520e
	CRC_stickerSet                                                        = 0xeeb46f27
	CRC_messages_stickerSet                                               = 0xb60a24a6
	CRC_botCommand                                                        = 0xc27ac8c7
	CRC_botInfo                                                           = 0x98e81d3a
	CRC_keyboardButton                                                    = 0xa2fa4880
	CRC_keyboardButtonUrl                                                 = 0x258aff05
	CRC_keyboardButtonCallback                                            = 0x683a5e46
	CRC_keyboardButtonRequestPhone                                        = 0xb16a6c29
	CRC_keyboardButtonRequestGeoLocation                                  = 0xfc796b3f
	CRC_keyboardButtonSwitchInline                                        = 0x0568a748
	CRC_keyboardButtonGame                                                = 0x50f41ccf
	CRC_keyboardButtonBuy                                                 = 0xafd93fbb
	CRC_keyboardButtonUrlAuth                                             = 0x10b78d29
	CRC_inputKeyboardButtonUrlAuth                                        = 0xd02e7fd4
	CRC_keyboardButtonRequestPoll                                         = 0xbbc7515d
	CRC_keyboardButtonRow                                                 = 0x77608b83
	CRC_replyKeyboardHide                                                 = 0xa03e5b85
	CRC_replyKeyboardForceReply                                           = 0xf4108aa0
	CRC_replyKeyboardMarkup                                               = 0x3502758c
	CRC_replyInlineMarkup                                                 = 0x48a30254
	CRC_messageEntityUnknown                                              = 0xbb92ba95
	CRC_messageEntityMention                                              = 0xfa04579d
	CRC_messageEntityHashtag                                              = 0x6f635b0d
	CRC_messageEntityBotCommand                                           = 0x6cef8ac7
	CRC_messageEntityUrl                                                  = 0x6ed02538
	CRC_messageEntityEmail                                                = 0x64e475c2
	CRC_messageEntityBold                                                 = 0xbd610bc9
	CRC_messageEntityItalic                                               = 0x826f8b60
	CRC_messageEntityCode                                                 = 0x28a20571
	CRC_messageEntityPre                                                  = 0x73924be0
	CRC_messageEntityTextUrl                                              = 0x76a6d327
	CRC_messageEntityMentionName                                          = 0x352dca58
	CRC_inputMessageEntityMentionName                                     = 0x208e68c9
	CRC_messageEntityPhone                                                = 0x9b69e34b
	CRC_messageEntityCashtag                                              = 0x4c4e743f
	CRC_messageEntityUnderline                                            = 0x9c4e7e8b
	CRC_messageEntityStrike                                               = 0xbf0693d4
	CRC_messageEntityBlockquote                                           = 0x020df5d0
	CRC_messageEntityBankCard                                             = 0x761e6af4
	CRC_inputChannelEmpty                                                 = 0xee8c1e86
	CRC_inputChannel                                                      = 0xafeb712e
	CRC_inputChannelFromMessage                                           = 0x2a286531
	CRC_contacts_resolvedPeer                                             = 0x7f077ad9
	CRC_messageRange                                                      = 0x0ae30253
	CRC_updates_channelDifferenceEmpty                                    = 0x3e11affb
	CRC_updates_channelDifferenceTooLong                                  = 0xa4bcc6fe
	CRC_updates_channelDifference                                         = 0x2064674e
	CRC_channelMessagesFilterEmpty                                        = 0x94d42ee7
	CRC_channelMessagesFilter                                             = 0xcd77d957
	CRC_channelParticipant                                                = 0x15ebac1d
	CRC_channelParticipantSelf                                            = 0xa3289a6d
	CRC_channelParticipantCreator                                         = 0x808d15a4
	CRC_channelParticipantAdmin                                           = 0xccbebbaf
	CRC_channelParticipantBanned                                          = 0x1c0facaf
	CRC_channelParticipantsRecent                                         = 0xde3f3c79
	CRC_channelParticipantsAdmins                                         = 0xb4608969
	CRC_channelParticipantsKicked                                         = 0xa3b54985
	CRC_channelParticipantsBots                                           = 0xb0d1865b
	CRC_channelParticipantsBanned                                         = 0x1427a5e1
	CRC_channelParticipantsSearch                                         = 0x0656ac4b
	CRC_channelParticipantsContacts                                       = 0xbb6ae88d
	CRC_channels_channelParticipants                                      = 0xf56ee2a8
	CRC_channels_channelParticipantsNotModified                           = 0xf0173fe9
	CRC_channels_channelParticipant                                       = 0xd0d9b163
	CRC_help_termsOfService                                               = 0x780a0310
	CRC_foundGif                                                          = 0x162ecc1f
	CRC_foundGifCached                                                    = 0x9c750409
	CRC_messages_foundGifs                                                = 0x450a1c0a
	CRC_messages_savedGifsNotModified                                     = 0xe8025ca2
	CRC_messages_savedGifs                                                = 0x2e0709a5
	CRC_inputBotInlineMessageMediaAuto                                    = 0x3380c786
	CRC_inputBotInlineMessageText                                         = 0x3dcd7a87
	CRC_inputBotInlineMessageMediaGeo                                     = 0xc1b15d65
	CRC_inputBotInlineMessageMediaVenue                                   = 0x417bbf11
	CRC_inputBotInlineMessageMediaContact                                 = 0xa6edbffd
	CRC_inputBotInlineMessageGame                                         = 0x4b425864
	CRC_inputBotInlineResult                                              = 0x88bf9319
	CRC_inputBotInlineResultPhoto                                         = 0xa8d864a7
	CRC_inputBotInlineResultDocument                                      = 0xfff8fdc4
	CRC_inputBotInlineResultGame                                          = 0x4fa417f2
	CRC_botInlineMessageMediaAuto                                         = 0x764cf810
	CRC_botInlineMessageText                                              = 0x8c7f65e2
	CRC_botInlineMessageMediaGeo                                          = 0xb722de65
	CRC_botInlineMessageMediaVenue                                        = 0x8a86659c
	CRC_botInlineMessageMediaContact                                      = 0x18d1cdc2
	CRC_botInlineResult                                                   = 0x11965f3a
	CRC_botInlineMediaResult                                              = 0x17db940b
	CRC_messages_botResults                                               = 0x947ca848
	CRC_exportedMessageLink                                               = 0x5dab1af4
	CRC_messageFwdHeader                                                  = 0x353a686b
	CRC_auth_codeTypeSms                                                  = 0x72a3158c
	CRC_auth_codeTypeCall                                                 = 0x741cd3e3
	CRC_auth_codeTypeFlashCall                                            = 0x226ccefb
	CRC_auth_sentCodeTypeApp                                              = 0x3dbb5986
	CRC_auth_sentCodeTypeSms                                              = 0xc000bba2
	CRC_auth_sentCodeTypeCall                                             = 0x5353e5a7
	CRC_auth_sentCodeTypeFlashCall                                        = 0xab03c6d9
	CRC_messages_botCallbackAnswer                                        = 0x36585ea4
	CRC_messages_messageEditData                                          = 0x26b5dde6
	CRC_inputBotInlineMessageID                                           = 0x890c3d89
	CRC_inlineBotSwitchPM                                                 = 0x3c20629f
	CRC_messages_peerDialogs                                              = 0x3371c354
	CRC_topPeer                                                           = 0xedcdc05b
	CRC_topPeerCategoryBotsPM                                             = 0xab661b5b
	CRC_topPeerCategoryBotsInline                                         = 0x148677e2
	CRC_topPeerCategoryCorrespondents                                     = 0x0637b7ed
	CRC_topPeerCategoryGroups                                             = 0xbd17a14a
	CRC_topPeerCategoryChannels                                           = 0x161d9628
	CRC_topPeerCategoryPhoneCalls                                         = 0x1e76a78c
	CRC_topPeerCategoryForwardUsers                                       = 0xa8406ca9
	CRC_topPeerCategoryForwardChats                                       = 0xfbeec0f0
	CRC_topPeerCategoryPeers                                              = 0xfb834291
	CRC_contacts_topPeersNotModified                                      = 0xde266ef5
	CRC_contacts_topPeers                                                 = 0x70b772a8
	CRC_contacts_topPeersDisabled                                         = 0xb52c939d
	CRC_draftMessageEmpty                                                 = 0x1b0c841a
	CRC_draftMessage                                                      = 0xfd8e711f
	CRC_messages_featuredStickersNotModified                              = 0xc6dc0c66
	CRC_messages_featuredStickers                                         = 0xb6abc341
	CRC_messages_recentStickersNotModified                                = 0x0b17f890
	CRC_messages_recentStickers                                           = 0x22f3afb3
	CRC_messages_archivedStickers                                         = 0x4fcba9c8
	CRC_messages_stickerSetInstallResultSuccess                           = 0x38641628
	CRC_messages_stickerSetInstallResultArchive                           = 0x35e410a8
	CRC_stickerSetCovered                                                 = 0x6410a5d2
	CRC_stickerSetMultiCovered                                            = 0x3407e51b
	CRC_maskCoords                                                        = 0xaed6dbb2
	CRC_inputStickeredMediaPhoto                                          = 0x4a992157
	CRC_inputStickeredMediaDocument                                       = 0x0438865b
	CRC_game                                                              = 0xbdf9653b
	CRC_inputGameID                                                       = 0x032c3e77
	CRC_inputGameShortName                                                = 0xc331e80a
	CRC_highScore                                                         = 0x58fffcd0
	CRC_messages_highScores                                               = 0x9a3bfd99
	CRC_textEmpty                                                         = 0xdc3d824f
	CRC_textPlain                                                         = 0x744694e0
	CRC_textBold                                                          = 0x6724abc4
	CRC_textItalic                                                        = 0xd912a59c
	CRC_textUnderline                                                     = 0xc12622c4
	CRC_textStrike                                                        = 0x9bf8bb95
	CRC_textFixed                                                         = 0x6c3f19b9
	CRC_textUrl                                                           = 0x3c2884c1
	CRC_textEmail                                                         = 0xde5a0dd6
	CRC_textConcat                                                        = 0x7e6260d7
	CRC_textSubscript                                                     = 0xed6a8504
	CRC_textSuperscript                                                   = 0xc7fb5e01
	CRC_textMarked                                                        = 0x034b8621
	CRC_textPhone                                                         = 0x1ccb966a
	CRC_textImage                                                         = 0x081ccf4f
	CRC_textAnchor                                                        = 0x35553762
	CRC_pageBlockUnsupported                                              = 0x13567e8a
	CRC_pageBlockTitle                                                    = 0x70abc3fd
	CRC_pageBlockSubtitle                                                 = 0x8ffa9a1f
	CRC_pageBlockAuthorDate                                               = 0xbaafe5e0
	CRC_pageBlockHeader                                                   = 0xbfd064ec
	CRC_pageBlockSubheader                                                = 0xf12bb6e1
	CRC_pageBlockParagraph                                                = 0x467a0766
	CRC_pageBlockPreformatted                                             = 0xc070d93e
	CRC_pageBlockFooter                                                   = 0x48870999
	CRC_pageBlockDivider                                                  = 0xdb20b188
	CRC_pageBlockAnchor                                                   = 0xce0d37b0
	CRC_pageBlockList                                                     = 0xe4e88011
	CRC_pageBlockBlockquote                                               = 0x263d7c26
	CRC_pageBlockPullquote                                                = 0x4f4456d3
	CRC_pageBlockPhoto                                                    = 0x1759c560
	CRC_pageBlockVideo                                                    = 0x7c8fe7b6
	CRC_pageBlockCover                                                    = 0x39f23300
	CRC_pageBlockEmbed                                                    = 0xa8718dc5
	CRC_pageBlockEmbedPost                                                = 0xf259a80b
	CRC_pageBlockCollage                                                  = 0x65a0fa4d
	CRC_pageBlockSlideshow                                                = 0x031f9590
	CRC_pageBlockChannel                                                  = 0xef1751b5
	CRC_pageBlockAudio                                                    = 0x804361ea
	CRC_pageBlockKicker                                                   = 0x1e148390
	CRC_pageBlockTable                                                    = 0xbf4dea82
	CRC_pageBlockOrderedList                                              = 0x9a8ae1e1
	CRC_pageBlockDetails                                                  = 0x76768bed
	CRC_pageBlockRelatedArticles                                          = 0x16115a96
	CRC_pageBlockMap                                                      = 0xa44f3ef6
	CRC_phoneCallDiscardReasonMissed                                      = 0x85e42301
	CRC_phoneCallDiscardReasonDisconnect                                  = 0xe095c1a0
	CRC_phoneCallDiscardReasonHangup                                      = 0x57adc690
	CRC_phoneCallDiscardReasonBusy                                        = 0xfaf7e8c9
	CRC_dataJSON                                                          = 0x7d748d04
	CRC_labeledPrice                                                      = 0xcb296bf8
	CRC_invoice                                                           = 0xc30aa358
	CRC_paymentCharge                                                     = 0xea02c27e
	CRC_postAddress                                                       = 0x1e8caaeb
	CRC_paymentRequestedInfo                                              = 0x909c3f94
	CRC_paymentSavedCredentialsCard                                       = 0xcdc27a1f
	CRC_webDocument                                                       = 0x1c570ed1
	CRC_webDocumentNoProxy                                                = 0xf9c8bcc6
	CRC_inputWebDocument                                                  = 0x9bed434d
	CRC_inputWebFileLocation                                              = 0xc239d686
	CRC_inputWebFileGeoPointLocation                                      = 0x9f2221c9
	CRC_upload_webFile                                                    = 0x21e753bc
	CRC_payments_paymentForm                                              = 0x3f56aea3
	CRC_payments_validatedRequestedInfo                                   = 0xd1451883
	CRC_payments_paymentResult                                            = 0x4e5f810d
	CRC_payments_paymentVerificationNeeded                                = 0xd8411139
	CRC_payments_paymentReceipt                                           = 0x500911e1
	CRC_payments_savedInfo                                                = 0xfb8fe43c
	CRC_inputPaymentCredentialsSaved                                      = 0xc10eb2cf
	CRC_inputPaymentCredentials                                           = 0x3417d728
	CRC_inputPaymentCredentialsApplePay                                   = 0x0aa1c39f
	CRC_inputPaymentCredentialsAndroidPay                                 = 0xca05d50e
	CRC_account_tmpPassword                                               = 0xdb64fd34
	CRC_shippingOption                                                    = 0xb6213cdf
	CRC_inputStickerSetItem                                               = 0xffa0a496
	CRC_inputPhoneCall                                                    = 0x1e36fded
	CRC_phoneCallEmpty                                                    = 0x5366c915
	CRC_phoneCallWaiting                                                  = 0x1b8f4ad1
	CRC_phoneCallRequested                                                = 0x87eabb53
	CRC_phoneCallAccepted                                                 = 0x997c454a
	CRC_phoneCall                                                         = 0x8742ae7f
	CRC_phoneCallDiscarded                                                = 0x50ca4de1
	CRC_phoneConnection                                                   = 0x9d4c17c0
	CRC_phoneCallProtocol                                                 = 0xfc878fc8
	CRC_phone_phoneCall                                                   = 0xec82e140
	CRC_upload_cdnFileReuploadNeeded                                      = 0xeea8e46e
	CRC_upload_cdnFile                                                    = 0xa99fca4f
	CRC_cdnPublicKey                                                      = 0xc982eaba
	CRC_cdnConfig                                                         = 0x5725e40a
	CRC_langPackString                                                    = 0xcad181f6
	CRC_langPackStringPluralized                                          = 0x6c47ac9f
	CRC_langPackStringDeleted                                             = 0x2979eeb2
	CRC_langPackDifference                                                = 0xf385c1f6
	CRC_langPackLanguage                                                  = 0xeeca5ce3
	CRC_channelAdminLogEventActionChangeTitle                             = 0xe6dfb825
	CRC_channelAdminLogEventActionChangeAbout                             = 0x55188a2e
	CRC_channelAdminLogEventActionChangeUsername                          = 0x6a4afc38
	CRC_channelAdminLogEventActionChangePhoto                             = 0x434bd2af
	CRC_channelAdminLogEventActionToggleInvites                           = 0x1b7907ae
	CRC_channelAdminLogEventActionToggleSignatures                        = 0x26ae0971
	CRC_channelAdminLogEventActionUpdatePinned                            = 0xe9e82c18
	CRC_channelAdminLogEventActionEditMessage                             = 0x709b2405
	CRC_channelAdminLogEventActionDeleteMessage                           = 0x42e047bb
	CRC_channelAdminLogEventActionParticipantJoin                         = 0x183040d3
	CRC_channelAdminLogEventActionParticipantLeave                        = 0xf89777f2
	CRC_channelAdminLogEventActionParticipantInvite                       = 0xe31c34d8
	CRC_channelAdminLogEventActionParticipantToggleBan                    = 0xe6d83d7e
	CRC_channelAdminLogEventActionParticipantToggleAdmin                  = 0xd5676710
	CRC_channelAdminLogEventActionChangeStickerSet                        = 0xb1c3caa7
	CRC_channelAdminLogEventActionTogglePreHistoryHidden                  = 0x5f5c95f1
	CRC_channelAdminLogEventActionDefaultBannedRights                     = 0x2df5fc0a
	CRC_channelAdminLogEventActionStopPoll                                = 0x8f079643
	CRC_channelAdminLogEventActionChangeLinkedChat                        = 0xa26f881b
	CRC_channelAdminLogEventActionChangeLocation                          = 0x0e6b76ae
	CRC_channelAdminLogEventActionToggleSlowMode                          = 0x53909779
	CRC_channelAdminLogEvent                                              = 0x3b5a3e40
	CRC_channels_adminLogResults                                          = 0xed8af74d
	CRC_channelAdminLogEventsFilter                                       = 0xea107ae4
	CRC_popularContact                                                    = 0x5ce14175
	CRC_messages_favedStickersNotModified                                 = 0x9e8fa6d3
	CRC_messages_favedStickers                                            = 0xf37f2f16
	CRC_recentMeUrlUnknown                                                = 0x46e1d13d
	CRC_recentMeUrlUser                                                   = 0x8dbc3336
	CRC_recentMeUrlChat                                                   = 0xa01b22f9
	CRC_recentMeUrlChatInvite                                             = 0xeb49081d
	CRC_recentMeUrlStickerSet                                             = 0xbc0a57dc
	CRC_help_recentMeUrls                                                 = 0x0e0310d7
	CRC_inputSingleMedia                                                  = 0x1cc6e91f
	CRC_webAuthorization                                                  = 0xcac943f2
	CRC_account_webAuthorizations                                         = 0xed56c9fc
	CRC_inputMessageID                                                    = 0xa676a322
	CRC_inputMessageReplyTo                                               = 0xbad88395
	CRC_inputMessagePinned                                                = 0x86872538
	CRC_inputDialogPeer                                                   = 0xfcaafeb7
	CRC_inputDialogPeerFolder                                             = 0x64600527
	CRC_dialogPeer                                                        = 0xe56dbf05
	CRC_dialogPeerFolder                                                  = 0x514519e2
	CRC_messages_foundStickerSetsNotModified                              = 0x0d54b65d
	CRC_messages_foundStickerSets                                         = 0x5108d648
	CRC_fileHash                                                          = 0x6242c773
	CRC_inputClientProxy                                                  = 0x75588b3f
	CRC_help_termsOfServiceUpdateEmpty                                    = 0xe3309f7f
	CRC_help_termsOfServiceUpdate                                         = 0x28ecf961
	CRC_inputSecureFileUploaded                                           = 0x3334b0f0
	CRC_inputSecureFile                                                   = 0x5367e5be
	CRC_secureFileEmpty                                                   = 0x64199744
	CRC_secureFile                                                        = 0xe0277a62
	CRC_secureData                                                        = 0x8aeabec3
	CRC_securePlainPhone                                                  = 0x7d6099dd
	CRC_securePlainEmail                                                  = 0x21ec5a5f
	CRC_secureValueTypePersonalDetails                                    = 0x9d2a81e3
	CRC_secureValueTypePassport                                           = 0x3dac6a00
	CRC_secureValueTypeDriverLicense                                      = 0x06e425c4
	CRC_secureValueTypeIdentityCard                                       = 0xa0d0744b
	CRC_secureValueTypeInternalPassport                                   = 0x99a48f23
	CRC_secureValueTypeAddress                                            = 0xcbe31e26
	CRC_secureValueTypeUtilityBill                                        = 0xfc36954e
	CRC_secureValueTypeBankStatement                                      = 0x89137c0d
	CRC_secureValueTypeRentalAgreement                                    = 0x8b883488
	CRC_secureValueTypePassportRegistration                               = 0x99e3806a
	CRC_secureValueTypeTemporaryRegistration                              = 0xea02ec33
	CRC_secureValueTypePhone                                              = 0xb320aadb
	CRC_secureValueTypeEmail                                              = 0x8e3ca7ee
	CRC_secureValue                                                       = 0x187fa0ca
	CRC_inputSecureValue                                                  = 0xdb21d0a7
	CRC_secureValueHash                                                   = 0xed1ecdb0
	CRC_secureValueErrorData                                              = 0xe8a40bd9
	CRC_secureValueErrorFrontSide                                         = 0x00be3dfa
	CRC_secureValueErrorReverseSide                                       = 0x868a2aa5
	CRC_secureValueErrorSelfie                                            = 0xe537ced6
	CRC_secureValueErrorFile                                              = 0x7a700873
	CRC_secureValueErrorFiles                                             = 0x666220e9
	CRC_secureValueError                                                  = 0x869d758f
	CRC_secureValueErrorTranslationFile                                   = 0xa1144770
	CRC_secureValueErrorTranslationFiles                                  = 0x34636dd8
	CRC_secureCredentialsEncrypted                                        = 0x33f0ea47
	CRC_account_authorizationForm                                         = 0xad2e1cd8
	CRC_account_sentEmailCode                                             = 0x811f854f
	CRC_help_deepLinkInfoEmpty                                            = 0x66afa166
	CRC_help_deepLinkInfo                                                 = 0x6a4ee832
	CRC_savedPhoneContact                                                 = 0x1142bd56
	CRC_account_takeout                                                   = 0x4dba4501
	CRC_passwordKdfAlgoUnknown                                            = 0xd45ab096
	CRC_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow = 0x3a912d4a
	CRC_securePasswordKdfAlgoUnknown                                      = 0x004a8537
	CRC_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000                   = 0xbbf2dda0
	CRC_securePasswordKdfAlgoSHA512                                       = 0x86471d92
	CRC_secureSecretSettings                                              = 0x1527bcac
	CRC_inputCheckPasswordEmpty                                           = 0x9880f658
	CRC_inputCheckPasswordSRP                                             = 0xd27ff082
	CRC_secureRequiredType                                                = 0x829d99da
	CRC_secureRequiredTypeOneOf                                           = 0x027477b4
	CRC_help_passportConfigNotModified                                    = 0xbfb9f457
	CRC_help_passportConfig                                               = 0xa098d6af
	CRC_inputAppEvent                                                     = 0x1d1b1245
	CRC_jsonObjectValue                                                   = 0xc0de1bd9
	CRC_jsonNull                                                          = 0x3f6d7b68
	CRC_jsonBool                                                          = 0xc7345e6a
	CRC_jsonNumber                                                        = 0x2be0dfa4
	CRC_jsonString                                                        = 0xb71e767a
	CRC_jsonArray                                                         = 0xf7444763
	CRC_jsonObject                                                        = 0x99c1d49d
	CRC_pageTableCell                                                     = 0x34566b6a
	CRC_pageTableRow                                                      = 0xe0c0c5e5
	CRC_pageCaption                                                       = 0x6f747657
	CRC_pageListItemText                                                  = 0xb92fb6cd
	CRC_pageListItemBlocks                                                = 0x25e073fc
	CRC_pageListOrderedItemText                                           = 0x5e068047
	CRC_pageListOrderedItemBlocks                                         = 0x98dd8936
	CRC_pageRelatedArticle                                                = 0xb390dc08
	CRC_page                                                              = 0x98657f0d
	CRC_help_supportName                                                  = 0x8c05f1c9
	CRC_help_userInfoEmpty                                                = 0xf3ae2eed
	CRC_help_userInfo                                                     = 0x01eb3758
	CRC_pollAnswer                                                        = 0x6ca9c2e9
	CRC_poll                                                              = 0x86e18161
	CRC_pollAnswerVoters                                                  = 0x3b6ddad2
	CRC_pollResults                                                       = 0xbadcc1a3
	CRC_chatOnlines                                                       = 0xf041e250
	CRC_statsURL                                                          = 0x47a971e0
	CRC_chatAdminRights                                                   = 0x5fb224d5
	CRC_chatBannedRights                                                  = 0x9f120418
	CRC_inputWallPaper                                                    = 0xe630b979
	CRC_inputWallPaperSlug                                                = 0x72091c80
	CRC_inputWallPaperNoFile                                              = 0x8427bbac
	CRC_account_wallPapersNotModified                                     = 0x1c199183
	CRC_account_wallPapers                                                = 0x702b65a9
	CRC_codeSettings                                                      = 0xdebebe83
	CRC_wallPaperSettings                                                 = 0x05086cf8
	CRC_autoDownloadSettings                                              = 0xe04232f3
	CRC_account_autoDownloadSettings                                      = 0x63cacf26
	CRC_emojiKeyword                                                      = 0xd5b3b9f9
	CRC_emojiKeywordDeleted                                               = 0x236df622
	CRC_emojiKeywordsDifference                                           = 0x5cc761bd
	CRC_emojiURL                                                          = 0xa575739d
	CRC_emojiLanguage                                                     = 0xb3fb5361
	CRC_fileLocationToBeDeprecated                                        = 0xbc7fc6cd
	CRC_folder                                                            = 0xff544e65
	CRC_inputFolderPeer                                                   = 0xfbd2c296
	CRC_folderPeer                                                        = 0xe9baa668
	CRC_messages_searchCounter                                            = 0xe844ebff
	CRC_urlAuthResultRequest                                              = 0x92d33a0e
	CRC_urlAuthResultAccepted                                             = 0x8f8c0e4e
	CRC_urlAuthResultDefault                                              = 0xa9d6db1f
	CRC_channelLocationEmpty                                              = 0xbfb5ad8b
	CRC_channelLocation                                                   = 0x209b82db
	CRC_peerLocated                                                       = 0xca461b5d
	CRC_peerSelfLocated                                                   = 0xf8ec284b
	CRC_restrictionReason                                                 = 0xd072acb4
	CRC_inputTheme                                                        = 0x3c5693e9
	CRC_inputThemeSlug                                                    = 0xf5890df1
	CRC_theme                                                             = 0x028f1114
	CRC_account_themesNotModified                                         = 0xf41eb622
	CRC_account_themes                                                    = 0x7f676421
	CRC_auth_loginToken                                                   = 0x629f1980
	CRC_auth_loginTokenMigrateTo                                          = 0x068e9916
	CRC_auth_loginTokenSuccess                                            = 0x390d5c5e
	CRC_account_contentSettings                                           = 0x57e28221
	CRC_messages_inactiveChats                                            = 0xa927fec5
	CRC_baseThemeClassic                                                  = 0xc3a12462
	CRC_baseThemeDay                                                      = 0xfbd81688
	CRC_baseThemeNight                                                    = 0xb7b31ea8
	CRC_baseThemeTinted                                                   = 0x6d5f77ee
	CRC_baseThemeArctic                                                   = 0x5b11125a
	CRC_inputThemeSettings                                                = 0xbd507cd1
	CRC_themeSettings                                                     = 0x9c14984a
	CRC_webPageAttributeTheme                                             = 0x54b56617
	CRC_messageUserVote                                                   = 0xa28e5559
	CRC_messageUserVoteInputOption                                        = 0x36377430
	CRC_messageUserVoteMultiple                                           = 0x0e8fe0de
	CRC_messages_votesList                                                = 0x0823f649
	CRC_bankCardOpenUrl                                                   = 0xf568028a
	CRC_payments_bankCardData                                             = 0x3e24e573
	CRC_dialogFilter                                                      = 0x7438f7e8
	CRC_dialogFilterSuggested                                             = 0x77744d4a
	CRC_statsDateRangeDays                                                = 0xb637edaf
	CRC_statsAbsValueAndPrev                                              = 0xcb43acde
	CRC_statsPercentValue                                                 = 0xcbce2fe0
	CRC_statsGraphAsync                                                   = 0x4a27eb2d
	CRC_statsGraphError                                                   = 0xbedc9822
	CRC_statsGraph                                                        = 0x8ea464b6
	CRC_messageInteractionCounters                                        = 0xad4fc9bd
	CRC_stats_broadcastStats                                              = 0xbdf78394
	CRC_help_promoDataEmpty                                               = 0x98f6ac75
	CRC_help_promoData                                                    = 0x8c39793f
	CRC_videoSize                                                         = 0x435bb987
	CRC_invokeAfterMsg                                                    = 0xcb9f372d
	CRC_invokeAfterMsgs                                                   = 0x3dc4b4f0
	CRC_initConnection                                                    = 0xc1cd5ea9
	CRC_invokeWithLayer                                                   = 0xda9b0d0d
	CRC_invokeWithoutUpdates                                              = 0xbf9459b7
	CRC_invokeWithMessagesRange                                           = 0x365275f2
	CRC_invokeWithTakeout                                                 = 0xaca9fd2e
	CRC_auth_sendCode                                                     = 0xa677244f
	CRC_auth_signUp                                                       = 0x80eee427
	CRC_auth_signIn                                                       = 0xbcd51581
	CRC_auth_logOut                                                       = 0x5717da40
	CRC_auth_resetAuthorizations                                          = 0x9fab0d1a
	CRC_auth_exportAuthorization                                          = 0xe5bfffcd
	CRC_auth_importAuthorization                                          = 0xe3ef9613
	CRC_auth_bindTempAuthKey                                              = 0xcdd42a05
	CRC_auth_importBotAuthorization                                       = 0x67a3ff2c
	CRC_auth_checkPassword                                                = 0xd18b4d16
	CRC_auth_requestPasswordRecovery                                      = 0xd897bc66
	CRC_auth_recoverPassword                                              = 0x4ea56e92
	CRC_auth_resendCode                                                   = 0x3ef1a9bf
	CRC_auth_cancelCode                                                   = 0x1f040578
	CRC_auth_dropTempAuthKeys                                             = 0x8e48a188
	CRC_auth_exportLoginToken                                             = 0xb1b41517
	CRC_auth_importLoginToken                                             = 0x95ac5ce4
	CRC_auth_acceptLoginToken                                             = 0xe894ad4d
	CRC_account_registerDevice                                            = 0x68976c6f
	CRC_account_unregisterDevice                                          = 0x3076c4bf
	CRC_account_updateNotifySettings                                      = 0x84be5b93
	CRC_account_getNotifySettings                                         = 0x12b3ad31
	CRC_account_resetNotifySettings                                       = 0xdb7e1747
	CRC_account_updateProfile                                             = 0x78515775
	CRC_account_updateStatus                                              = 0x6628562c
	CRC_account_getWallPapers                                             = 0xaabb1763
	CRC_account_reportPeer                                                = 0xae189d5f
	CRC_account_checkUsername                                             = 0x2714d86c
	CRC_account_updateUsername                                            = 0x3e0bdd7c
	CRC_account_getPrivacy                                                = 0xdadbc950
	CRC_account_setPrivacy                                                = 0xc9f81ce8
	CRC_account_deleteAccount                                             = 0x418d4e0b
	CRC_account_getAccountTTL                                             = 0x08fc711d
	CRC_account_setAccountTTL                                             = 0x2442485e
	CRC_account_sendChangePhoneCode                                       = 0x82574ae5
	CRC_account_changePhone                                               = 0x70c32edb
	CRC_account_updateDeviceLocked                                        = 0x38df3532
	CRC_account_getAuthorizations                                         = 0xe320c158
	CRC_account_resetAuthorization                                        = 0xdf77f3bc
	CRC_account_getPassword                                               = 0x548a30f5
	CRC_account_getPasswordSettings                                       = 0x9cd4eaf9
	CRC_account_updatePasswordSettings                                    = 0xa59b102f
	CRC_account_sendConfirmPhoneCode                                      = 0x1b3faa88
	CRC_account_confirmPhone                                              = 0x5f2178c3
	CRC_account_getTmpPassword                                            = 0x449e0b51
	CRC_account_getWebAuthorizations                                      = 0x182e6d6f
	CRC_account_resetWebAuthorization                                     = 0x2d01b9ef
	CRC_account_resetWebAuthorizations                                    = 0x682d2594
	CRC_account_getAllSecureValues                                        = 0xb288bc7d
	CRC_account_getSecureValue                                            = 0x73665bc2
	CRC_account_saveSecureValue                                           = 0x899fe31d
	CRC_account_deleteSecureValue                                         = 0xb880bc4b
	CRC_account_getAuthorizationForm                                      = 0xb86ba8e1
	CRC_account_acceptAuthorization                                       = 0xe7027c94
	CRC_account_sendVerifyPhoneCode                                       = 0xa5a356f9
	CRC_account_verifyPhone                                               = 0x4dd3a7f6
	CRC_account_sendVerifyEmailCode                                       = 0x7011509f
	CRC_account_verifyEmail                                               = 0xecba39db
	CRC_account_initTakeoutSession                                        = 0xf05b4804
	CRC_account_finishTakeoutSession                                      = 0x1d2652ee
	CRC_account_confirmPasswordEmail                                      = 0x8fdf1920
	CRC_account_resendPasswordEmail                                       = 0x7a7f2a15
	CRC_account_cancelPasswordEmail                                       = 0xc1cbd5b6
	CRC_account_getContactSignUpNotification                              = 0x9f07c728
	CRC_account_setContactSignUpNotification                              = 0xcff43f61
	CRC_account_getNotifyExceptions                                       = 0x53577479
	CRC_account_getWallPaper                                              = 0xfc8ddbea
	CRC_account_uploadWallPaper                                           = 0xdd853661
	CRC_account_saveWallPaper                                             = 0x6c5a5b37
	CRC_account_installWallPaper                                          = 0xfeed5769
	CRC_account_resetWallPapers                                           = 0xbb3b9804
	CRC_account_getAutoDownloadSettings                                   = 0x56da0b3f
	CRC_account_saveAutoDownloadSettings                                  = 0x76f36233
	CRC_account_uploadTheme                                               = 0x1c3db333
	CRC_account_createTheme                                               = 0x8432c21f
	CRC_account_updateTheme                                               = 0x5cb367d5
	CRC_account_saveTheme                                                 = 0xf257106c
	CRC_account_installTheme                                              = 0x7ae43737
	CRC_account_getTheme                                                  = 0x8d9d742b
	CRC_account_getThemes                                                 = 0x285946f8
	CRC_account_setContentSettings                                        = 0xb574b16b
	CRC_account_getContentSettings                                        = 0x8b9b4dae
	CRC_account_getMultiWallPapers                                        = 0x65ad71dc
	CRC_users_getUsers                                                    = 0x0d91a548
	CRC_users_getFullUser                                                 = 0xca30a5b1
	CRC_users_setSecureValueErrors                                        = 0x90c894b5
	CRC_contacts_getContactIDs                                            = 0x2caa4a42
	CRC_contacts_getStatuses                                              = 0xc4a353ee
	CRC_contacts_getContacts                                              = 0xc023849f
	CRC_contacts_importContacts                                           = 0x2c800be5
	CRC_contacts_deleteContacts                                           = 0x096a0e00
	CRC_contacts_deleteByPhones                                           = 0x1013fd9e
	CRC_contacts_block                                                    = 0x332b49fc
	CRC_contacts_unblock                                                  = 0xe54100bd
	CRC_contacts_getBlocked                                               = 0xf57c350f
	CRC_contacts_search                                                   = 0x11f812d8
	CRC_contacts_resolveUsername                                          = 0xf93ccba3
	CRC_contacts_getTopPeers                                              = 0xd4982db5
	CRC_contacts_resetTopPeerRating                                       = 0x1ae373ac
	CRC_contacts_resetSaved                                               = 0x879537f1
	CRC_contacts_getSaved                                                 = 0x82f1e39f
	CRC_contacts_toggleTopPeers                                           = 0x8514bdda
	CRC_contacts_addContact                                               = 0xe8f463d0
	CRC_contacts_acceptContact                                            = 0xf831a20f
	CRC_contacts_getLocated                                               = 0xd348bc44
	CRC_messages_getMessages                                              = 0x63c66506
	CRC_messages_getDialogs                                               = 0xa0ee3b73
	CRC_messages_getHistory                                               = 0xdcbb8260
	CRC_messages_search                                                   = 0x8614ef68
	CRC_messages_readHistory                                              = 0x0e306d3a
	CRC_messages_deleteHistory                                            = 0x1c015b09
	CRC_messages_deleteMessages                                           = 0xe58e95d2
	CRC_messages_receivedMessages                                         = 0x05a954c0
	CRC_messages_setTyping                                                = 0xa3825e50
	CRC_messages_sendMessage                                              = 0x520c3870
	CRC_messages_sendMedia                                                = 0x3491eba9
	CRC_messages_forwardMessages                                          = 0xd9fee60e
	CRC_messages_reportSpam                                               = 0xcf1592db
	CRC_messages_getPeerSettings                                          = 0x3672e09c
	CRC_messages_report                                                   = 0xbd82b658
	CRC_messages_getChats                                                 = 0x3c6aa187
	CRC_messages_getFullChat                                              = 0x3b831c66
	CRC_messages_editChatTitle                                            = 0xdc452855
	CRC_messages_editChatPhoto                                            = 0xca4c79d8
	CRC_messages_addChatUser                                              = 0xf9a0aa09
	CRC_messages_deleteChatUser                                           = 0xe0611f16
	CRC_messages_createChat                                               = 0x09cb126e
	CRC_messages_getDhConfig                                              = 0x26cf8950
	CRC_messages_requestEncryption                                        = 0xf64daf43
	CRC_messages_acceptEncryption                                         = 0x3dbc0415
	CRC_messages_discardEncryption                                        = 0xedd923c5
	CRC_messages_setEncryptedTyping                                       = 0x791451ed
	CRC_messages_readEncryptedHistory                                     = 0x7f4b690a
	CRC_messages_sendEncrypted                                            = 0xa9776773
	CRC_messages_sendEncryptedFile                                        = 0x9a901b66
	CRC_messages_sendEncryptedService                                     = 0x32d439a4
	CRC_messages_receivedQueue                                            = 0x55a5bb66
	CRC_messages_reportEncryptedSpam                                      = 0x4b0c8c0f
	CRC_messages_readMessageContents                                      = 0x36a73f77
	CRC_messages_getStickers                                              = 0x043d4f2c
	CRC_messages_getAllStickers                                           = 0x1c9618b1
	CRC_messages_getWebPagePreview                                        = 0x8b68b0cc
	CRC_messages_exportChatInvite                                         = 0x0df7534c
	CRC_messages_checkChatInvite                                          = 0x3eadb1bb
	CRC_messages_importChatInvite                                         = 0x6c50051c
	CRC_messages_getStickerSet                                            = 0x2619a90e
	CRC_messages_installStickerSet                                        = 0xc78fe460
	CRC_messages_uninstallStickerSet                                      = 0xf96e55de
	CRC_messages_startBot                                                 = 0xe6df7378
	CRC_messages_getMessagesViews                                         = 0xc4c8a55d
	CRC_messages_editChatAdmin                                            = 0xa9e69f2e
	CRC_messages_migrateChat                                              = 0x15a3b8e3
	CRC_messages_searchGlobal                                             = 0xbf7225a4
	CRC_messages_reorderStickerSets                                       = 0x78337739
	CRC_messages_getDocumentByHash                                        = 0x338e2464
	CRC_messages_searchGifs                                               = 0xbf9a776b
	CRC_messages_getSavedGifs                                             = 0x83bf3d52
	CRC_messages_saveGif                                                  = 0x327a30cb
	CRC_messages_getInlineBotResults                                      = 0x514e999d
	CRC_messages_setInlineBotResults                                      = 0xeb5ea206
	CRC_messages_sendInlineBotResult                                      = 0x220815b0
	CRC_messages_getMessageEditData                                       = 0xfda68d36
	CRC_messages_editMessage                                              = 0x48f71778
	CRC_messages_editInlineBotMessage                                     = 0x83557dba
	CRC_messages_getBotCallbackAnswer                                     = 0x810a9fec
	CRC_messages_setBotCallbackAnswer                                     = 0xd58f130a
	CRC_messages_getPeerDialogs                                           = 0xe470bcfd
	CRC_messages_saveDraft                                                = 0xbc39e14b
	CRC_messages_getAllDrafts                                             = 0x6a3f8d65
	CRC_messages_getFeaturedStickers                                      = 0x2dacca4f
	CRC_messages_readFeaturedStickers                                     = 0x5b118126
	CRC_messages_getRecentStickers                                        = 0x5ea192c9
	CRC_messages_saveRecentSticker                                        = 0x392718f8
	CRC_messages_clearRecentStickers                                      = 0x8999602d
	CRC_messages_getArchivedStickers                                      = 0x57f17692
	CRC_messages_getMaskStickers                                          = 0x65b8c79f
	CRC_messages_getAttachedStickers                                      = 0xcc5b67cc
	CRC_messages_setGameScore                                             = 0x8ef8ecc0
	CRC_messages_setInlineGameScore                                       = 0x15ad9f64
	CRC_messages_getGameHighScores                                        = 0xe822649d
	CRC_messages_getInlineGameHighScores                                  = 0x0f635e1b
	CRC_messages_getCommonChats                                           = 0x0d0a48c4
	CRC_messages_getAllChats                                              = 0xeba80ff0
	CRC_messages_getWebPage                                               = 0x32ca8f91
	CRC_messages_toggleDialogPin                                          = 0xa731e257
	CRC_messages_reorderPinnedDialogs                                     = 0x3b1adf37
	CRC_messages_getPinnedDialogs                                         = 0xd6b94df2
	CRC_messages_setBotShippingResults                                    = 0xe5f672fa
	CRC_messages_setBotPrecheckoutResults                                 = 0x09c2dd95
	CRC_messages_uploadMedia                                              = 0x519bc2b1
	CRC_messages_sendScreenshotNotification                               = 0xc97df020
	CRC_messages_getFavedStickers                                         = 0x21ce0b0e
	CRC_messages_faveSticker                                              = 0xb9ffc55b
	CRC_messages_getUnreadMentions                                        = 0x46578472
	CRC_messages_readMentions                                             = 0x0f0189d3
	CRC_messages_getRecentLocations                                       = 0xbbc45b09
	CRC_messages_sendMultiMedia                                           = 0xcc0110cb
	CRC_messages_uploadEncryptedFile                                      = 0x5057c497
	CRC_messages_searchStickerSets                                        = 0xc2b7d08b
	CRC_messages_getSplitRanges                                           = 0x1cff7e08
	CRC_messages_markDialogUnread                                         = 0xc286d98f
	CRC_messages_getDialogUnreadMarks                                     = 0x22e24e22
	CRC_messages_clearAllDrafts                                           = 0x7e58ee9c
	CRC_messages_updatePinnedMessage                                      = 0xd2aaf7ec
	CRC_messages_sendVote                                                 = 0x10ea6184
	CRC_messages_getPollResults                                           = 0x73bb643b
	CRC_messages_getOnlines                                               = 0x6e2be050
	CRC_messages_getStatsURL                                              = 0x812c2ae6
	CRC_messages_editChatAbout                                            = 0xdef60797
	CRC_messages_editChatDefaultBannedRights                              = 0xa5866b41
	CRC_messages_getEmojiKeywords                                         = 0x35a0e062
	CRC_messages_getEmojiKeywordsDifference                               = 0x1508b6af
	CRC_messages_getEmojiKeywordsLanguages                                = 0x4e9963b2
	CRC_messages_getEmojiURL                                              = 0xd5b10c26
	CRC_messages_getSearchCounters                                        = 0x732eef00
	CRC_messages_requestUrlAuth                                           = 0xe33f5613
	CRC_messages_acceptUrlAuth                                            = 0xf729ea98
	CRC_messages_hidePeerSettingsBar                                      = 0x4facb138
	CRC_messages_getScheduledHistory                                      = 0xe2c2685b
	CRC_messages_getScheduledMessages                                     = 0xbdbb0464
	CRC_messages_sendScheduledMessages                                    = 0xbd38850a
	CRC_messages_deleteScheduledMessages                                  = 0x59ae2b16
	CRC_messages_getPollVotes                                             = 0xb86e380e
	CRC_messages_toggleStickerSets                                        = 0xb5052fea
	CRC_messages_getDialogFilters                                         = 0xf19ed96d
	CRC_messages_getSuggestedDialogFilters                                = 0xa29cd42c
	CRC_messages_updateDialogFilter                                       = 0x1ad4a04a
	CRC_messages_updateDialogFiltersOrder                                 = 0xc563c1e4
	CRC_messages_getOldFeaturedStickers                                   = 0x5fe7025b
	CRC_updates_getState                                                  = 0xedd4882a
	CRC_updates_getDifference                                             = 0x25939651
	CRC_updates_getChannelDifference                                      = 0x03173d78
	CRC_photos_updateProfilePhoto                                         = 0xf0bb5152
	CRC_photos_uploadProfilePhoto                                         = 0x4f32c098
	CRC_photos_deletePhotos                                               = 0x87cf7f2f
	CRC_photos_getUserPhotos                                              = 0x91cd32a8
	CRC_upload_saveFilePart                                               = 0xb304a621
	CRC_upload_getFile                                                    = 0xb15a9afc
	CRC_upload_saveBigFilePart                                            = 0xde7b673d
	CRC_upload_getWebFile                                                 = 0x24e6818d
	CRC_upload_getCdnFile                                                 = 0x2000bcc3
	CRC_upload_reuploadCdnFile                                            = 0x9b2754a8
	CRC_upload_getCdnFileHashes                                           = 0x4da54231
	CRC_upload_getFileHashes                                              = 0xc7025931
	CRC_help_getConfig                                                    = 0xc4f9186b
	CRC_help_getNearestDc                                                 = 0x1fb33026
	CRC_help_getAppUpdate                                                 = 0x522d5a7d
	CRC_help_getInviteText                                                = 0x4d392343
	CRC_help_getSupport                                                   = 0x9cdf08cd
	CRC_help_getAppChangelog                                              = 0x9010ef6f
	CRC_help_setBotUpdatesStatus                                          = 0xec22cfcd
	CRC_help_getCdnConfig                                                 = 0x52029342
	CRC_help_getRecentMeUrls                                              = 0x3dc0f114
	CRC_help_getTermsOfServiceUpdate                                      = 0x2ca51fd1
	CRC_help_acceptTermsOfService                                         = 0xee72f79a
	CRC_help_getDeepLinkInfo                                              = 0x3fedc75f
	CRC_help_getAppConfig                                                 = 0x98914110
	CRC_help_saveAppLog                                                   = 0x6f02f748
	CRC_help_getPassportConfig                                            = 0xc661ad08
	CRC_help_getSupportName                                               = 0xd360e72c
	CRC_help_getUserInfo                                                  = 0x038a08d3
	CRC_help_editUserInfo                                                 = 0x66b91b70
	CRC_help_getPromoData                                                 = 0xc0977421
	CRC_help_hidePromoData                                                = 0x1e251c95
	CRC_channels_readHistory                                              = 0xcc104937
	CRC_channels_deleteMessages                                           = 0x84c1fd4e
	CRC_channels_deleteUserHistory                                        = 0xd10dd71b
	CRC_channels_reportSpam                                               = 0xfe087810
	CRC_channels_getMessages                                              = 0xad8c9a23
	CRC_channels_getParticipants                                          = 0x123e05e9
	CRC_channels_getParticipant                                           = 0x546dd7a6
	CRC_channels_getChannels                                              = 0x0a7f6bbb
	CRC_channels_getFullChannel                                           = 0x08736a09
	CRC_channels_createChannel                                            = 0x3d5fb10f
	CRC_channels_editAdmin                                                = 0xd33c8902
	CRC_channels_editTitle                                                = 0x566decd0
	CRC_channels_editPhoto                                                = 0xf12e57c9
	CRC_channels_checkUsername                                            = 0x10e6bd2c
	CRC_channels_updateUsername                                           = 0x3514b3de
	CRC_channels_joinChannel                                              = 0x24b524c5
	CRC_channels_leaveChannel                                             = 0xf836aa95
	CRC_channels_inviteToChannel                                          = 0x199f3a6c
	CRC_channels_deleteChannel                                            = 0xc0111fe3
	CRC_channels_exportMessageLink                                        = 0xceb77163
	CRC_channels_toggleSignatures                                         = 0x1f69b606
	CRC_channels_getAdminedPublicChannels                                 = 0xf8b036af
	CRC_channels_editBanned                                               = 0x72796912
	CRC_channels_getAdminLog                                              = 0x33ddf480
	CRC_channels_setStickers                                              = 0xea8ca4f9
	CRC_channels_readMessageContents                                      = 0xeab5dc38
	CRC_channels_deleteHistory                                            = 0xaf369d42
	CRC_channels_togglePreHistoryHidden                                   = 0xeabbb94c
	CRC_channels_getLeftChannels                                          = 0x8341ecc0
	CRC_channels_getGroupsForDiscussion                                   = 0xf5dad378
	CRC_channels_setDiscussionGroup                                       = 0x40582bb2
	CRC_channels_editCreator                                              = 0x8f38cd1f
	CRC_channels_editLocation                                             = 0x58e63f6d
	CRC_channels_toggleSlowMode                                           = 0xedd49ef0
	CRC_channels_getInactiveChannels                                      = 0x11e831ee
	CRC_bots_sendCustomRequest                                            = 0xaa2769ed
	CRC_bots_answerWebhookJSONQuery                                       = 0xe6213f4d
	CRC_bots_setBotCommands                                               = 0x805d46f6
	CRC_payments_getPaymentForm                                           = 0x99f09745
	CRC_payments_getPaymentReceipt                                        = 0xa092a980
	CRC_payments_validateRequestedInfo                                    = 0x770a8e74
	CRC_payments_sendPaymentForm                                          = 0x2b8879b3
	CRC_payments_getSavedInfo                                             = 0x227d824b
	CRC_payments_clearSavedInfo                                           = 0xd83d70c1
	CRC_payments_getBankCardData                                          = 0x2e79d779
	CRC_stickers_createStickerSet                                         = 0xf1036780
	CRC_stickers_removeStickerFromSet                                     = 0xf7760f51
	CRC_stickers_changeStickerPosition                                    = 0xffb6d4ca
	CRC_stickers_addStickerToSet                                          = 0x8653febe
	CRC_stickers_setStickerSetThumb                                       = 0x9a364e30
	CRC_phone_getCallConfig                                               = 0x55451fa9
	CRC_phone_requestCall                                                 = 0x42ff96ed
	CRC_phone_acceptCall                                                  = 0x3bd2b4a0
	CRC_phone_confirmCall                                                 = 0x2efe1722
	CRC_phone_receivedCall                                                = 0x17d54f61
	CRC_phone_discardCall                                                 = 0xb2cbc1c0
	CRC_phone_setCallRating                                               = 0x59ead627
	CRC_phone_saveCallDebug                                               = 0x277add7e
	CRC_phone_sendSignalingData                                           = 0xff7a9383
	CRC_langpack_getLangPack                                              = 0xf2f2330a
	CRC_langpack_getStrings                                               = 0xefea3803
	CRC_langpack_getDifference                                            = 0xcd984aa5
	CRC_langpack_getLanguages                                             = 0x42c6978f
	CRC_langpack_getLanguage                                              = 0x6a596502
	CRC_folders_editPeerFolders                                           = 0x6847d0ab
	CRC_folders_deleteFolder                                              = 0x1c295881
	CRC_stats_getBroadcastStats                                           = 0xab42441a
	CRC_stats_loadAsyncGraph                                              = 0x621d5fa0
)

type TL_resPQ struct {
	Nonce                       []byte
	ServerNonce                 []byte
	Pq                          string
	ServerPublicKeyFingerprints []int64
}

type TL_p_q_inner_data struct {
	Pq          string
	P           string
	Q           string
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
}

type TL_p_q_inner_data_dc struct {
	Pq          string
	P           string
	Q           string
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
	Dc          int32
}

type TL_p_q_inner_data_temp struct {
	Pq          string
	P           string
	Q           string
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
	ExpiresIn   int32
}

type TL_p_q_inner_data_temp_dc struct {
	Pq          string
	P           string
	Q           string
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
	Dc          int32
	ExpiresIn   int32
}

type TL_bind_auth_key_inner struct {
	Nonce         int64
	TempAuthKeyID int64
	PermAuthKeyID int64
	TempSessionID int64
	ExpiresAt     int32
}

type TL_server_DH_params_fail struct {
	Nonce        []byte
	ServerNonce  []byte
	NewNonceHash []byte
}

type TL_server_DH_params_ok struct {
	Nonce           []byte
	ServerNonce     []byte
	EncryptedAnswer string
}

type TL_server_DH_inner_data struct {
	Nonce       []byte
	ServerNonce []byte
	G           int32
	DhPrime     string
	GA          string
	ServerTime  int32
}

type TL_client_DH_inner_data struct {
	Nonce       []byte
	ServerNonce []byte
	RetryID     int64
	GB          string
}

type TL_dh_gen_ok struct {
	Nonce         []byte
	ServerNonce   []byte
	NewNonceHash1 []byte
}

type TL_dh_gen_retry struct {
	Nonce         []byte
	ServerNonce   []byte
	NewNonceHash2 []byte
}

type TL_dh_gen_fail struct {
	Nonce         []byte
	ServerNonce   []byte
	NewNonceHash3 []byte
}

type TL_destroy_auth_key_ok struct {
}

type TL_destroy_auth_key_none struct {
}

type TL_destroy_auth_key_fail struct {
}

type TL_req_pq struct {
	Nonce []byte
}

type TL_req_pq_multi struct {
	Nonce []byte
}

type TL_req_DH_params struct {
	Nonce                []byte
	ServerNonce          []byte
	P                    string
	Q                    string
	PublicKeyFingerprint int64
	EncryptedData        string
}

type TL_set_client_DH_params struct {
	Nonce         []byte
	ServerNonce   []byte
	EncryptedData string
}

type TL_destroy_auth_key struct {
}

type TL_msgs_ack struct {
	MsgIds []int64
}

type TL_bad_msg_notification struct {
	BadMsgID    int64
	BadMsgSeqno int32
	ErrorCode   int32
}

type TL_bad_server_salt struct {
	BadMsgID      int64
	BadMsgSeqno   int32
	ErrorCode     int32
	NewServerSalt int64
}

type TL_msgs_state_req struct {
	MsgIds []int64
}

type TL_msgs_state_info struct {
	ReqMsgID int64
	Info     string
}

type TL_msgs_all_info struct {
	MsgIds []int64
	Info   string
}

type TL_msg_detailed_info struct {
	MsgID       int64
	AnswerMsgID int64
	Bytes       int32
	Status      int32
}

type TL_msg_new_detailed_info struct {
	AnswerMsgID int64
	Bytes       int32
	Status      int32
}

type TL_msg_resend_req struct {
	MsgIds []int64
}

type TL_rpc_error struct {
	ErrorCode    int32
	ErrorMessage string
}

type TL_rpc_answer_unknown struct {
}

type TL_rpc_answer_dropped_running struct {
}

type TL_rpc_answer_dropped struct {
	MsgID int64
	SeqNo int32
	Bytes int32
}

type TL_future_salt struct {
	ValidSince int32
	ValidUntil int32
	Salt       int64
}

type TL_future_salts struct {
	ReqMsgID int64
	Now      int32
	Salts    TL // vector<future_salt>
}

type TL_pong struct {
	MsgID  int64
	PingID int64
}

type TL_destroy_session_ok struct {
	SessionID int64
}

type TL_destroy_session_none struct {
	SessionID int64
}

type TL_new_session_created struct {
	FirstMsgID int64
	UniqueID   int64
	ServerSalt int64
}

type TL_http_wait struct {
	MaxDelay  int32
	WaitAfter int32
	MaxWait   int32
}

type TL_ipPort struct {
	Ipv4 int32
	Port int32
}

type TL_ipPortSecret struct {
	Ipv4   int32
	Port   int32
	Secret []byte
}

type TL_accessPointRule struct {
	PhonePrefixRules string
	DcID             int32
	Ips              TL // vector<IpPort>
}

type TL_help_configSimple struct {
	Date    int32
	Expires int32
	Rules   TL // vector<AccessPointRule>
}

type TL_tlsClientHello struct {
	Blocks TL // vector<TlsBlock>
}

type TL_tlsBlockString struct {
	Data string
}

type TL_tlsBlockRandom struct {
	Length int32
}

type TL_tlsBlockZero struct {
	Length int32
}

type TL_tlsBlockDomain struct {
}

type TL_tlsBlockGrease struct {
	Seed int32
}

type TL_tlsBlockPublicKey struct {
}

type TL_tlsBlockScope struct {
	Entries []TL // TlsBlock
}

type TL_rpc_drop_answer struct {
	ReqMsgID int64
}

type TL_get_future_salts struct {
	Num int32
}

type TL_ping struct {
	PingID int64
}

type TL_ping_delay_disconnect struct {
	PingID          int64
	DisconnectDelay int32
}

type TL_destroy_session struct {
	SessionID int64
}

type TL_boolFalse struct {
}

type TL_boolTrue struct {
}

type TL_true struct {
}

type TL_error struct {
	Code int32
	Text string
}

type TL_null struct {
}

type TL_inputPeerEmpty struct {
}

type TL_inputPeerSelf struct {
}

type TL_inputPeerChat struct {
	ChatID int32
}

type TL_inputPeerUser struct {
	UserID     int32
	AccessHash int64
}

type TL_inputPeerChannel struct {
	ChannelID  int32
	AccessHash int64
}

type TL_inputPeerUserFromMessage struct {
	Peer   TL // InputPeer
	MsgID  int32
	UserID int32
}

type TL_inputPeerChannelFromMessage struct {
	Peer      TL // InputPeer
	MsgID     int32
	ChannelID int32
}

type TL_inputUserEmpty struct {
}

type TL_inputUserSelf struct {
}

type TL_inputUser struct {
	UserID     int32
	AccessHash int64
}

type TL_inputUserFromMessage struct {
	Peer   TL // InputPeer
	MsgID  int32
	UserID int32
}

type TL_inputPhoneContact struct {
	ClientID  int64
	Phone     string
	FirstName string
	LastName  string
}

type TL_inputFile struct {
	ID          int64
	Parts       int32
	Name        string
	Md5Checksum string
}

type TL_inputFileBig struct {
	ID    int64
	Parts int32
	Name  string
}

type TL_inputMediaEmpty struct {
}

type TL_inputMediaUploadedPhoto struct {
	Flags      int32
	File       TL    // InputFile
	Stickers   []TL  // InputDocument //flag
	TtlSeconds int32 //flag
}

type TL_inputMediaPhoto struct {
	Flags      int32
	ID         TL    // InputPhoto
	TtlSeconds int32 //flag
}

type TL_inputMediaGeoPoint struct {
	GeoPoint TL // InputGeoPoint
}

type TL_inputMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
}

type TL_inputMediaUploadedDocument struct {
	Flags        int32
	NosoundVideo bool //flag
	File         TL   // InputFile
	Thumb        TL   // InputFile //flag
	MimeType     string
	Attributes   []TL  // DocumentAttribute
	Stickers     []TL  // InputDocument //flag
	TtlSeconds   int32 //flag
}

type TL_inputMediaDocument struct {
	Flags      int32
	ID         TL    // InputDocument
	TtlSeconds int32 //flag
}

type TL_inputMediaVenue struct {
	GeoPoint  TL // InputGeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

type TL_inputMediaGifExternal struct {
	Url string
	Q   string
}

type TL_inputMediaPhotoExternal struct {
	Flags      int32
	Url        string
	TtlSeconds int32 //flag
}

type TL_inputMediaDocumentExternal struct {
	Flags      int32
	Url        string
	TtlSeconds int32 //flag
}

type TL_inputMediaGame struct {
	ID TL // InputGame
}

type TL_inputMediaInvoice struct {
	Flags        int32
	Title        string
	Description  string
	Photo        TL // InputWebDocument //flag
	Invoice      TL // Invoice
	Payload      []byte
	Provider     string
	ProviderData TL // DataJSON
	StartParam   string
}

type TL_inputMediaGeoLive struct {
	Flags    int32
	Stopped  bool  //flag
	GeoPoint TL    // InputGeoPoint
	Period   int32 //flag
}

type TL_inputMediaPoll struct {
	Flags            int32
	Poll             TL     // Poll
	CorrectAnswers   []TL   // bytes //flag
	Solution         string //flag
	SolutionEntities []TL   // MessageEntity //flag
}

type TL_inputMediaDice struct {
	Emoticon string
}

type TL_inputChatPhotoEmpty struct {
}

type TL_inputChatUploadedPhoto struct {
	File TL // InputFile
}

type TL_inputChatPhoto struct {
	ID TL // InputPhoto
}

type TL_inputGeoPointEmpty struct {
}

type TL_inputGeoPoint struct {
	Lat  float64
	Long float64
}

type TL_inputPhotoEmpty struct {
}

type TL_inputPhoto struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

type TL_inputFileLocation struct {
	VolumeID      int64
	LocalID       int32
	Secret        int64
	FileReference []byte
}

type TL_inputEncryptedFileLocation struct {
	ID         int64
	AccessHash int64
}

type TL_inputDocumentFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

type TL_inputSecureFileLocation struct {
	ID         int64
	AccessHash int64
}

type TL_inputTakeoutFileLocation struct {
}

type TL_inputPhotoFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

type TL_inputPhotoLegacyFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	VolumeID      int64
	LocalID       int32
	Secret        int64
}

type TL_inputPeerPhotoFileLocation struct {
	Flags    int32
	Big      bool //flag
	Peer     TL   // InputPeer
	VolumeID int64
	LocalID  int32
}

type TL_inputStickerSetThumb struct {
	Stickerset TL // InputStickerSet
	VolumeID   int64
	LocalID    int32
}

type TL_peerUser struct {
	UserID int32
}

type TL_peerChat struct {
	ChatID int32
}

type TL_peerChannel struct {
	ChannelID int32
}

type TL_storage_fileUnknown struct {
}

type TL_storage_filePartial struct {
}

type TL_storage_fileJpeg struct {
}

type TL_storage_fileGif struct {
}

type TL_storage_filePng struct {
}

type TL_storage_filePdf struct {
}

type TL_storage_fileMp3 struct {
}

type TL_storage_fileMov struct {
}

type TL_storage_fileMp4 struct {
}

type TL_storage_fileWebp struct {
}

type TL_userEmpty struct {
	ID int32
}

type TL_user struct {
	Flags                int32
	Self                 bool //flag
	Contact              bool //flag
	MutualContact        bool //flag
	Deleted              bool //flag
	Bot                  bool //flag
	BotChatHistory       bool //flag
	BotNochats           bool //flag
	Verified             bool //flag
	Restricted           bool //flag
	Min                  bool //flag
	BotInlineGeo         bool //flag
	Support              bool //flag
	Scam                 bool //flag
	ID                   int32
	AccessHash           int64  //flag
	FirstName            string //flag
	LastName             string //flag
	Username             string //flag
	Phone                string //flag
	Photo                TL     // UserProfilePhoto //flag
	Status               TL     // UserStatus //flag
	BotInfoVersion       int32  //flag
	RestrictionReason    []TL   // RestrictionReason //flag
	BotInlinePlaceholder string //flag
	LangCode             string //flag
}

type TL_userProfilePhotoEmpty struct {
}

type TL_userProfilePhoto struct {
	PhotoID    int64
	PhotoSmall TL // FileLocation
	PhotoBig   TL // FileLocation
	DcID       int32
}

type TL_userStatusEmpty struct {
}

type TL_userStatusOnline struct {
	Expires int32
}

type TL_userStatusOffline struct {
	WasOnline int32
}

type TL_userStatusRecently struct {
}

type TL_userStatusLastWeek struct {
}

type TL_userStatusLastMonth struct {
}

type TL_chatEmpty struct {
	ID int32
}

type TL_chat struct {
	Flags               int32
	Creator             bool //flag
	Kicked              bool //flag
	Left                bool //flag
	Deactivated         bool //flag
	ID                  int32
	Title               string
	Photo               TL // ChatPhoto
	ParticipantsCount   int32
	Date                int32
	Version             int32
	MigratedTo          TL // InputChannel //flag
	AdminRights         TL // ChatAdminRights //flag
	DefaultBannedRights TL // ChatBannedRights //flag
}

type TL_chatForbidden struct {
	ID    int32
	Title string
}

type TL_channel struct {
	Flags               int32
	Creator             bool //flag
	Left                bool //flag
	Broadcast           bool //flag
	Verified            bool //flag
	Megagroup           bool //flag
	Restricted          bool //flag
	Signatures          bool //flag
	Min                 bool //flag
	Scam                bool //flag
	HasLink             bool //flag
	HasGeo              bool //flag
	SlowmodeEnabled     bool //flag
	ID                  int32
	AccessHash          int64 //flag
	Title               string
	Username            string //flag
	Photo               TL     // ChatPhoto
	Date                int32
	Version             int32
	RestrictionReason   []TL  // RestrictionReason //flag
	AdminRights         TL    // ChatAdminRights //flag
	BannedRights        TL    // ChatBannedRights //flag
	DefaultBannedRights TL    // ChatBannedRights //flag
	ParticipantsCount   int32 //flag
}

type TL_channelForbidden struct {
	Flags      int32
	Broadcast  bool //flag
	Megagroup  bool //flag
	ID         int32
	AccessHash int64
	Title      string
	UntilDate  int32 //flag
}

type TL_chatFull struct {
	Flags          int32
	CanSetUsername bool //flag
	HasScheduled   bool //flag
	ID             int32
	About          string
	Participants   TL    // ChatParticipants
	ChatPhoto      TL    // Photo //flag
	NotifySettings TL    // PeerNotifySettings
	ExportedInvite TL    // ExportedChatInvite
	BotInfo        []TL  // BotInfo //flag
	PinnedMsgID    int32 //flag
	FolderID       int32 //flag
}

type TL_channelFull struct {
	Flags                int32
	CanViewParticipants  bool //flag
	CanSetUsername       bool //flag
	CanSetStickers       bool //flag
	HiddenPrehistory     bool //flag
	CanViewStats         bool //flag
	CanSetLocation       bool //flag
	HasScheduled         bool //flag
	ID                   int32
	About                string
	ParticipantsCount    int32 //flag
	AdminsCount          int32 //flag
	KickedCount          int32 //flag
	BannedCount          int32 //flag
	OnlineCount          int32 //flag
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	ChatPhoto            TL    // Photo
	NotifySettings       TL    // PeerNotifySettings
	ExportedInvite       TL    // ExportedChatInvite
	BotInfo              []TL  // BotInfo
	MigratedFromChatID   int32 //flag
	MigratedFromMaxID    int32 //flag
	PinnedMsgID          int32 //flag
	Stickerset           TL    // StickerSet //flag
	AvailableMinID       int32 //flag
	FolderID             int32 //flag
	LinkedChatID         int32 //flag
	Location             TL    // ChannelLocation //flag
	SlowmodeSeconds      int32 //flag
	SlowmodeNextSendDate int32 //flag
	StatsDc              int32 //flag
	Pts                  int32
}

type TL_chatParticipant struct {
	UserID    int32
	InviterID int32
	Date      int32
}

type TL_chatParticipantCreator struct {
	UserID int32
}

type TL_chatParticipantAdmin struct {
	UserID    int32
	InviterID int32
	Date      int32
}

type TL_chatParticipantsForbidden struct {
	Flags           int32
	ChatID          int32
	SelfParticipant TL // ChatParticipant //flag
}

type TL_chatParticipants struct {
	ChatID       int32
	Participants []TL // ChatParticipant
	Version      int32
}

type TL_chatPhotoEmpty struct {
}

type TL_chatPhoto struct {
	PhotoSmall TL // FileLocation
	PhotoBig   TL // FileLocation
	DcID       int32
}

type TL_messageEmpty struct {
	ID int32
}

type TL_message struct {
	Flags             int32
	Out               bool //flag
	Mentioned         bool //flag
	MediaUnread       bool //flag
	Silent            bool //flag
	Post              bool //flag
	FromScheduled     bool //flag
	Legacy            bool //flag
	EditHide          bool //flag
	ID                int32
	FromID            int32 //flag
	ToID              TL    // Peer
	FwdFrom           TL    // MessageFwdHeader //flag
	ViaBotID          int32 //flag
	ReplyToMsgID      int32 //flag
	Date              int32
	Message           string
	Media             TL     // MessageMedia //flag
	ReplyMarkup       TL     // ReplyMarkup //flag
	Entities          []TL   // MessageEntity //flag
	Views             int32  //flag
	EditDate          int32  //flag
	PostAuthor        string //flag
	GroupedID         int64  //flag
	RestrictionReason []TL   // RestrictionReason //flag
}

type TL_messageService struct {
	Flags        int32
	Out          bool //flag
	Mentioned    bool //flag
	MediaUnread  bool //flag
	Silent       bool //flag
	Post         bool //flag
	Legacy       bool //flag
	ID           int32
	FromID       int32 //flag
	ToID         TL    // Peer
	ReplyToMsgID int32 //flag
	Date         int32
	Action       TL // MessageAction
}

type TL_messageMediaEmpty struct {
}

type TL_messageMediaPhoto struct {
	Flags      int32
	Photo      TL    // Photo //flag
	TtlSeconds int32 //flag
}

type TL_messageMediaGeo struct {
	Geo TL // GeoPoint
}

type TL_messageMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	UserID      int32
}

type TL_messageMediaUnsupported struct {
}

type TL_messageMediaDocument struct {
	Flags      int32
	Document   TL    // Document //flag
	TtlSeconds int32 //flag
}

type TL_messageMediaWebPage struct {
	Webpage TL // WebPage
}

type TL_messageMediaVenue struct {
	Geo       TL // GeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

type TL_messageMediaGame struct {
	Game TL // Game
}

type TL_messageMediaInvoice struct {
	Flags                    int32
	ShippingAddressRequested bool //flag
	Test                     bool //flag
	Title                    string
	Description              string
	Photo                    TL    // WebDocument //flag
	ReceiptMsgID             int32 //flag
	Currency                 string
	TotalAmount              int64
	StartParam               string
}

type TL_messageMediaGeoLive struct {
	Geo    TL // GeoPoint
	Period int32
}

type TL_messageMediaPoll struct {
	Poll    TL // Poll
	Results TL // PollResults
}

type TL_messageMediaDice struct {
	Value    int32
	Emoticon string
}

type TL_messageActionEmpty struct {
}

type TL_messageActionChatCreate struct {
	Title string
	Users []int32
}

type TL_messageActionChatEditTitle struct {
	Title string
}

type TL_messageActionChatEditPhoto struct {
	Photo TL // Photo
}

type TL_messageActionChatDeletePhoto struct {
}

type TL_messageActionChatAddUser struct {
	Users []int32
}

type TL_messageActionChatDeleteUser struct {
	UserID int32
}

type TL_messageActionChatJoinedByLink struct {
	InviterID int32
}

type TL_messageActionChannelCreate struct {
	Title string
}

type TL_messageActionChatMigrateTo struct {
	ChannelID int32
}

type TL_messageActionChannelMigrateFrom struct {
	Title  string
	ChatID int32
}

type TL_messageActionPinMessage struct {
}

type TL_messageActionHistoryClear struct {
}

type TL_messageActionGameScore struct {
	GameID int64
	Score  int32
}

type TL_messageActionPaymentSentMe struct {
	Flags            int32
	Currency         string
	TotalAmount      int64
	Payload          []byte
	Info             TL     // PaymentRequestedInfo //flag
	ShippingOptionID string //flag
	Charge           TL     // PaymentCharge
}

type TL_messageActionPaymentSent struct {
	Currency    string
	TotalAmount int64
}

type TL_messageActionPhoneCall struct {
	Flags    int32
	Video    bool //flag
	CallID   int64
	Reason   TL    // PhoneCallDiscardReason //flag
	Duration int32 //flag
}

type TL_messageActionScreenshotTaken struct {
}

type TL_messageActionCustomAction struct {
	Message string
}

type TL_messageActionBotAllowed struct {
	Domain string
}

type TL_messageActionSecureValuesSentMe struct {
	Values      []TL // SecureValue
	Credentials TL   // SecureCredentialsEncrypted
}

type TL_messageActionSecureValuesSent struct {
	Types []TL // SecureValueType
}

type TL_messageActionContactSignUp struct {
}

type TL_dialog struct {
	Flags               int32
	Pinned              bool //flag
	UnreadMark          bool //flag
	Peer                TL   // Peer
	TopMessage          int32
	ReadInboxMaxID      int32
	ReadOutboxMaxID     int32
	UnreadCount         int32
	UnreadMentionsCount int32
	NotifySettings      TL    // PeerNotifySettings
	Pts                 int32 //flag
	Draft               TL    // DraftMessage //flag
	FolderID            int32 //flag
}

type TL_dialogFolder struct {
	Flags                      int32
	Pinned                     bool //flag
	Folder                     TL   // Folder
	Peer                       TL   // Peer
	TopMessage                 int32
	UnreadMutedPeersCount      int32
	UnreadUnmutedPeersCount    int32
	UnreadMutedMessagesCount   int32
	UnreadUnmutedMessagesCount int32
}

type TL_photoEmpty struct {
	ID int64
}

type TL_photo struct {
	Flags         int32
	HasStickers   bool //flag
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	Sizes         []TL // PhotoSize
	DcID          int32
}

type TL_photoSizeEmpty struct {
	Type string
}

type TL_photoSize struct {
	Type     string
	Location TL // FileLocation
	W        int32
	H        int32
	Size     int32
}

type TL_photoCachedSize struct {
	Type     string
	Location TL // FileLocation
	W        int32
	H        int32
	Bytes    []byte
}

type TL_photoStrippedSize struct {
	Type  string
	Bytes []byte
}

type TL_geoPointEmpty struct {
}

type TL_geoPoint struct {
	Long       float64
	Lat        float64
	AccessHash int64
}

type TL_auth_sentCode struct {
	Flags         int32
	Type          TL // auth_SentCodeType
	PhoneCodeHash string
	NextType      TL    // auth_CodeType //flag
	Timeout       int32 //flag
}

type TL_auth_authorization struct {
	Flags       int32
	TmpSessions int32 //flag
	User        TL    // User
}

type TL_auth_authorizationSignUpRequired struct {
	Flags          int32
	TermsOfService TL // help_TermsOfService //flag
}

type TL_auth_exportedAuthorization struct {
	ID    int32
	Bytes []byte
}

type TL_inputNotifyPeer struct {
	Peer TL // InputPeer
}

type TL_inputNotifyUsers struct {
}

type TL_inputNotifyChats struct {
}

type TL_inputNotifyBroadcasts struct {
}

type TL_inputPeerNotifySettings struct {
	Flags        int32
	ShowPreviews TL     // Bool //flag
	Silent       TL     // Bool //flag
	MuteUntil    int32  //flag
	Sound        string //flag
}

type TL_peerNotifySettings struct {
	Flags        int32
	ShowPreviews TL     // Bool //flag
	Silent       TL     // Bool //flag
	MuteUntil    int32  //flag
	Sound        string //flag
}

type TL_peerSettings struct {
	Flags                 int32
	ReportSpam            bool //flag
	AddContact            bool //flag
	BlockContact          bool //flag
	ShareContact          bool //flag
	NeedContactsException bool //flag
	ReportGeo             bool //flag
}

type TL_wallPaper struct {
	ID         int64
	Flags      int32
	Creator    bool //flag
	Default    bool //flag
	Pattern    bool //flag
	Dark       bool //flag
	AccessHash int64
	Slug       string
	Document   TL // Document
	Settings   TL // WallPaperSettings //flag
}

type TL_wallPaperNoFile struct {
	Flags    int32
	Default  bool //flag
	Dark     bool //flag
	Settings TL   // WallPaperSettings //flag
}

type TL_inputReportReasonSpam struct {
}

type TL_inputReportReasonViolence struct {
}

type TL_inputReportReasonPornography struct {
}

type TL_inputReportReasonChildAbuse struct {
}

type TL_inputReportReasonOther struct {
	Text string
}

type TL_inputReportReasonCopyright struct {
}

type TL_inputReportReasonGeoIrrelevant struct {
}

type TL_userFull struct {
	Flags               int32
	Blocked             bool   //flag
	PhoneCallsAvailable bool   //flag
	PhoneCallsPrivate   bool   //flag
	CanPinMessage       bool   //flag
	HasScheduled        bool   //flag
	User                TL     // User
	About               string //flag
	Settings            TL     // PeerSettings
	ProfilePhoto        TL     // Photo //flag
	NotifySettings      TL     // PeerNotifySettings
	BotInfo             TL     // BotInfo //flag
	PinnedMsgID         int32  //flag
	CommonChatsCount    int32
	FolderID            int32 //flag
}

type TL_contact struct {
	UserID int32
	Mutual TL // Bool
}

type TL_importedContact struct {
	UserID   int32
	ClientID int64
}

type TL_contactBlocked struct {
	UserID int32
	Date   int32
}

type TL_contactStatus struct {
	UserID int32
	Status TL // UserStatus
}

type TL_contacts_contactsNotModified struct {
}

type TL_contacts_contacts struct {
	Contacts   []TL // Contact
	SavedCount int32
	Users      []TL // User
}

type TL_contacts_importedContacts struct {
	Imported       []TL // ImportedContact
	PopularInvites []TL // PopularContact
	RetryContacts  []int64
	Users          []TL // User
}

type TL_contacts_blocked struct {
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TL_contacts_blockedSlice struct {
	Count   int32
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type TL_messages_dialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_dialogsSlice struct {
	Count    int32
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_dialogsNotModified struct {
	Count int32
}

type TL_messages_messages struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_messagesSlice struct {
	Flags    int32
	Inexact  bool //flag
	Count    int32
	NextRate int32 //flag
	Messages []TL  // Message
	Chats    []TL  // Chat
	Users    []TL  // User
}

type TL_messages_channelMessages struct {
	Flags    int32
	Inexact  bool //flag
	Pts      int32
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_messagesNotModified struct {
	Count int32
}

type TL_messages_chats struct {
	Chats []TL // Chat
}

type TL_messages_chatsSlice struct {
	Count int32
	Chats []TL // Chat
}

type TL_messages_chatFull struct {
	FullChat TL   // ChatFull
	Chats    []TL // Chat
	Users    []TL // User
}

type TL_messages_affectedHistory struct {
	Pts      int32
	PtsCount int32
	Offset   int32
}

type TL_inputMessagesFilterEmpty struct {
}

type TL_inputMessagesFilterPhotos struct {
}

type TL_inputMessagesFilterVideo struct {
}

type TL_inputMessagesFilterPhotoVideo struct {
}

type TL_inputMessagesFilterDocument struct {
}

type TL_inputMessagesFilterUrl struct {
}

type TL_inputMessagesFilterGif struct {
}

type TL_inputMessagesFilterVoice struct {
}

type TL_inputMessagesFilterMusic struct {
}

type TL_inputMessagesFilterChatPhotos struct {
}

type TL_inputMessagesFilterPhoneCalls struct {
	Flags  int32
	Missed bool //flag
}

type TL_inputMessagesFilterRoundVoice struct {
}

type TL_inputMessagesFilterRoundVideo struct {
}

type TL_inputMessagesFilterMyMentions struct {
}

type TL_inputMessagesFilterGeo struct {
}

type TL_inputMessagesFilterContacts struct {
}

type TL_updateNewMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type TL_updateMessageID struct {
	ID       int32
	RandomID int64
}

type TL_updateDeleteMessages struct {
	Messages []int32
	Pts      int32
	PtsCount int32
}

type TL_updateUserTyping struct {
	UserID int32
	Action TL // SendMessageAction
}

type TL_updateChatUserTyping struct {
	ChatID int32
	UserID int32
	Action TL // SendMessageAction
}

type TL_updateChatParticipants struct {
	Participants TL // ChatParticipants
}

type TL_updateUserStatus struct {
	UserID int32
	Status TL // UserStatus
}

type TL_updateUserName struct {
	UserID    int32
	FirstName string
	LastName  string
	Username  string
}

type TL_updateUserPhoto struct {
	UserID   int32
	Date     int32
	Photo    TL // UserProfilePhoto
	Previous TL // Bool
}

type TL_updateNewEncryptedMessage struct {
	Message TL // EncryptedMessage
	Qts     int32
}

type TL_updateEncryptedChatTyping struct {
	ChatID int32
}

type TL_updateEncryption struct {
	Chat TL // EncryptedChat
	Date int32
}

type TL_updateEncryptedMessagesRead struct {
	ChatID  int32
	MaxDate int32
	Date    int32
}

type TL_updateChatParticipantAdd struct {
	ChatID    int32
	UserID    int32
	InviterID int32
	Date      int32
	Version   int32
}

type TL_updateChatParticipantDelete struct {
	ChatID  int32
	UserID  int32
	Version int32
}

type TL_updateDcOptions struct {
	DcOptions []TL // DcOption
}

type TL_updateUserBlocked struct {
	UserID  int32
	Blocked TL // Bool
}

type TL_updateNotifySettings struct {
	Peer           TL // NotifyPeer
	NotifySettings TL // PeerNotifySettings
}

type TL_updateServiceNotification struct {
	Flags     int32
	Popup     bool  //flag
	InboxDate int32 //flag
	Type      string
	Message   string
	Media     TL   // MessageMedia
	Entities  []TL // MessageEntity
}

type TL_updatePrivacy struct {
	Key   TL   // PrivacyKey
	Rules []TL // PrivacyRule
}

type TL_updateUserPhone struct {
	UserID int32
	Phone  string
}

type TL_updateReadHistoryInbox struct {
	Flags            int32
	FolderID         int32 //flag
	Peer             TL    // Peer
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
	PtsCount         int32
}

type TL_updateReadHistoryOutbox struct {
	Peer     TL // Peer
	MaxID    int32
	Pts      int32
	PtsCount int32
}

type TL_updateWebPage struct {
	Webpage  TL // WebPage
	Pts      int32
	PtsCount int32
}

type TL_updateReadMessagesContents struct {
	Messages []int32
	Pts      int32
	PtsCount int32
}

type TL_updateChannelTooLong struct {
	Flags     int32
	ChannelID int32
	Pts       int32 //flag
}

type TL_updateChannel struct {
	ChannelID int32
}

type TL_updateNewChannelMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type TL_updateReadChannelInbox struct {
	Flags            int32
	FolderID         int32 //flag
	ChannelID        int32
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
}

type TL_updateDeleteChannelMessages struct {
	ChannelID int32
	Messages  []int32
	Pts       int32
	PtsCount  int32
}

type TL_updateChannelMessageViews struct {
	ChannelID int32
	ID        int32
	Views     int32
}

type TL_updateChatParticipantAdmin struct {
	ChatID  int32
	UserID  int32
	IsAdmin TL // Bool
	Version int32
}

type TL_updateNewStickerSet struct {
	Stickerset TL // messages_StickerSet
}

type TL_updateStickerSetsOrder struct {
	Flags int32
	Masks bool //flag
	Order []int64
}

type TL_updateStickerSets struct {
}

type TL_updateSavedGifs struct {
}

type TL_updateBotInlineQuery struct {
	Flags   int32
	QueryID int64
	UserID  int32
	Query   string
	Geo     TL // GeoPoint //flag
	Offset  string
}

type TL_updateBotInlineSend struct {
	Flags  int32
	UserID int32
	Query  string
	Geo    TL // GeoPoint //flag
	ID     string
	MsgID  TL // InputBotInlineMessageID //flag
}

type TL_updateEditChannelMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type TL_updateChannelPinnedMessage struct {
	ChannelID int32
	ID        int32
}

type TL_updateBotCallbackQuery struct {
	Flags         int32
	QueryID       int64
	UserID        int32
	Peer          TL // Peer
	MsgID         int32
	ChatInstance  int64
	Data          []byte //flag
	GameShortName string //flag
}

type TL_updateEditMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type TL_updateInlineBotCallbackQuery struct {
	Flags         int32
	QueryID       int64
	UserID        int32
	MsgID         TL // InputBotInlineMessageID
	ChatInstance  int64
	Data          []byte //flag
	GameShortName string //flag
}

type TL_updateReadChannelOutbox struct {
	ChannelID int32
	MaxID     int32
}

type TL_updateDraftMessage struct {
	Peer  TL // Peer
	Draft TL // DraftMessage
}

type TL_updateReadFeaturedStickers struct {
}

type TL_updateRecentStickers struct {
}

type TL_updateConfig struct {
}

type TL_updatePtsChanged struct {
}

type TL_updateChannelWebPage struct {
	ChannelID int32
	Webpage   TL // WebPage
	Pts       int32
	PtsCount  int32
}

type TL_updateDialogPinned struct {
	Flags    int32
	Pinned   bool  //flag
	FolderID int32 //flag
	Peer     TL    // DialogPeer
}

type TL_updatePinnedDialogs struct {
	Flags    int32
	FolderID int32 //flag
	Order    []TL  // DialogPeer //flag
}

type TL_updateBotWebhookJSON struct {
	Data TL // DataJSON
}

type TL_updateBotWebhookJSONQuery struct {
	QueryID int64
	Data    TL // DataJSON
	Timeout int32
}

type TL_updateBotShippingQuery struct {
	QueryID         int64
	UserID          int32
	Payload         []byte
	ShippingAddress TL // PostAddress
}

type TL_updateBotPrecheckoutQuery struct {
	Flags            int32
	QueryID          int64
	UserID           int32
	Payload          []byte
	Info             TL     // PaymentRequestedInfo //flag
	ShippingOptionID string //flag
	Currency         string
	TotalAmount      int64
}

type TL_updatePhoneCall struct {
	PhoneCall TL // PhoneCall
}

type TL_updateLangPackTooLong struct {
	LangCode string
}

type TL_updateLangPack struct {
	Difference TL // LangPackDifference
}

type TL_updateFavedStickers struct {
}

type TL_updateChannelReadMessagesContents struct {
	ChannelID int32
	Messages  []int32
}

type TL_updateContactsReset struct {
}

type TL_updateChannelAvailableMessages struct {
	ChannelID      int32
	AvailableMinID int32
}

type TL_updateDialogUnreadMark struct {
	Flags  int32
	Unread bool //flag
	Peer   TL   // DialogPeer
}

type TL_updateUserPinnedMessage struct {
	UserID int32
	ID     int32
}

type TL_updateChatPinnedMessage struct {
	ChatID  int32
	ID      int32
	Version int32
}

type TL_updateMessagePoll struct {
	Flags   int32
	PollID  int64
	Poll    TL // Poll //flag
	Results TL // PollResults
}

type TL_updateChatDefaultBannedRights struct {
	Peer                TL // Peer
	DefaultBannedRights TL // ChatBannedRights
	Version             int32
}

type TL_updateFolderPeers struct {
	FolderPeers []TL // FolderPeer
	Pts         int32
	PtsCount    int32
}

type TL_updatePeerSettings struct {
	Peer     TL // Peer
	Settings TL // PeerSettings
}

type TL_updatePeerLocated struct {
	Peers []TL // PeerLocated
}

type TL_updateNewScheduledMessage struct {
	Message TL // Message
}

type TL_updateDeleteScheduledMessages struct {
	Peer     TL // Peer
	Messages []int32
}

type TL_updateTheme struct {
	Theme TL // Theme
}

type TL_updateGeoLiveViewed struct {
	Peer  TL // Peer
	MsgID int32
}

type TL_updateLoginToken struct {
}

type TL_updateMessagePollVote struct {
	PollID  int64
	UserID  int32
	Options []TL // bytes
}

type TL_updateDialogFilter struct {
	Flags  int32
	ID     int32
	Filter TL // DialogFilter //flag
}

type TL_updateDialogFilterOrder struct {
	Order []int32
}

type TL_updateDialogFilters struct {
}

type TL_updatePhoneCallSignalingData struct {
	PhoneCallID int64
	Data        []byte
}

type TL_updates_state struct {
	Pts         int32
	Qts         int32
	Date        int32
	Seq         int32
	UnreadCount int32
}

type TL_updates_differenceEmpty struct {
	Date int32
	Seq  int32
}

type TL_updates_difference struct {
	NewMessages          []TL // Message
	NewEncryptedMessages []TL // EncryptedMessage
	OtherUpdates         []TL // Update
	Chats                []TL // Chat
	Users                []TL // User
	State                TL   // updates_State
}

type TL_updates_differenceSlice struct {
	NewMessages          []TL // Message
	NewEncryptedMessages []TL // EncryptedMessage
	OtherUpdates         []TL // Update
	Chats                []TL // Chat
	Users                []TL // User
	IntermediateState    TL   // updates_State
}

type TL_updates_differenceTooLong struct {
	Pts int32
}

type TL_updatesTooLong struct {
}

type TL_updateShortMessage struct {
	Flags        int32
	Out          bool //flag
	Mentioned    bool //flag
	MediaUnread  bool //flag
	Silent       bool //flag
	ID           int32
	UserID       int32
	Message      string
	Pts          int32
	PtsCount     int32
	Date         int32
	FwdFrom      TL    // MessageFwdHeader //flag
	ViaBotID     int32 //flag
	ReplyToMsgID int32 //flag
	Entities     []TL  // MessageEntity //flag
}

type TL_updateShortChatMessage struct {
	Flags        int32
	Out          bool //flag
	Mentioned    bool //flag
	MediaUnread  bool //flag
	Silent       bool //flag
	ID           int32
	FromID       int32
	ChatID       int32
	Message      string
	Pts          int32
	PtsCount     int32
	Date         int32
	FwdFrom      TL    // MessageFwdHeader //flag
	ViaBotID     int32 //flag
	ReplyToMsgID int32 //flag
	Entities     []TL  // MessageEntity //flag
}

type TL_updateShort struct {
	Update TL // Update
	Date   int32
}

type TL_updatesCombined struct {
	Updates  []TL // Update
	Users    []TL // User
	Chats    []TL // Chat
	Date     int32
	SeqStart int32
	Seq      int32
}

type TL_updates struct {
	Updates []TL // Update
	Users   []TL // User
	Chats   []TL // Chat
	Date    int32
	Seq     int32
}

type TL_updateShortSentMessage struct {
	Flags    int32
	Out      bool //flag
	ID       int32
	Pts      int32
	PtsCount int32
	Date     int32
	Media    TL   // MessageMedia //flag
	Entities []TL // MessageEntity //flag
}

type TL_photos_photos struct {
	Photos []TL // Photo
	Users  []TL // User
}

type TL_photos_photosSlice struct {
	Count  int32
	Photos []TL // Photo
	Users  []TL // User
}

type TL_photos_photo struct {
	Photo TL   // Photo
	Users []TL // User
}

type TL_upload_file struct {
	Type  TL // storage_FileType
	Mtime int32
	Bytes []byte
}

type TL_upload_fileCdnRedirect struct {
	DcID          int32
	FileToken     []byte
	EncryptionKey []byte
	EncryptionIv  []byte
	FileHashes    []TL // FileHash
}

type TL_dcOption struct {
	Flags     int32
	Ipv6      bool //flag
	MediaOnly bool //flag
	TcpoOnly  bool //flag
	Cdn       bool //flag
	Static    bool //flag
	ID        int32
	IpAddress string
	Port      int32
	Secret    []byte //flag
}

type TL_config struct {
	Flags                   int32
	PhonecallsEnabled       bool //flag
	DefaultP2pContacts      bool //flag
	PreloadFeaturedStickers bool //flag
	IgnorePhoneEntities     bool //flag
	RevokePmInbox           bool //flag
	BlockedMode             bool //flag
	PfsEnabled              bool //flag
	Date                    int32
	Expires                 int32
	TestMode                TL // Bool
	ThisDc                  int32
	DcOptions               []TL // DcOption
	DcTxtDomainName         string
	ChatSizeMax             int32
	MegagroupSizeMax        int32
	ForwardedCountMax       int32
	OnlineUpdatePeriodMs    int32
	OfflineBlurTimeoutMs    int32
	OfflineIdleTimeoutMs    int32
	OnlineCloudTimeoutMs    int32
	NotifyCloudDelayMs      int32
	NotifyDefaultDelayMs    int32
	PushChatPeriodMs        int32
	PushChatLimit           int32
	SavedGifsLimit          int32
	EditTimeLimit           int32
	RevokeTimeLimit         int32
	RevokePmTimeLimit       int32
	RatingEDecay            int32
	StickersRecentLimit     int32
	StickersFavedLimit      int32
	ChannelsReadMediaPeriod int32
	TmpSessions             int32 //flag
	PinnedDialogsCountMax   int32
	PinnedInfolderCountMax  int32
	CallReceiveTimeoutMs    int32
	CallRingTimeoutMs       int32
	CallConnectTimeoutMs    int32
	CallPacketTimeoutMs     int32
	MeUrlPrefix             string
	AutoupdateUrlPrefix     string //flag
	GifSearchUsername       string //flag
	VenueSearchUsername     string //flag
	ImgSearchUsername       string //flag
	StaticMapsProvider      string //flag
	CaptionLengthMax        int32
	MessageLengthMax        int32
	WebfileDcID             int32
	SuggestedLangCode       string //flag
	LangPackVersion         int32  //flag
	BaseLangPackVersion     int32  //flag
}

type TL_nearestDc struct {
	Country   string
	ThisDc    int32
	NearestDc int32
}

type TL_help_appUpdate struct {
	Flags      int32
	CanNotSkip bool //flag
	ID         int32
	Version    string
	Text       string
	Entities   []TL   // MessageEntity
	Document   TL     // Document //flag
	Url        string //flag
}

type TL_help_noAppUpdate struct {
}

type TL_help_inviteText struct {
	Message string
}

type TL_encryptedChatEmpty struct {
	ID int32
}

type TL_encryptedChatWaiting struct {
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
}

type TL_encryptedChatRequested struct {
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	GA            []byte
}

type TL_encryptedChat struct {
	ID             int32
	AccessHash     int64
	Date           int32
	AdminID        int32
	ParticipantID  int32
	GAOrB          []byte
	KeyFingerprint int64
}

type TL_encryptedChatDiscarded struct {
	ID int32
}

type TL_inputEncryptedChat struct {
	ChatID     int32
	AccessHash int64
}

type TL_encryptedFileEmpty struct {
}

type TL_encryptedFile struct {
	ID             int64
	AccessHash     int64
	Size           int32
	DcID           int32
	KeyFingerprint int32
}

type TL_inputEncryptedFileEmpty struct {
}

type TL_inputEncryptedFileUploaded struct {
	ID             int64
	Parts          int32
	Md5Checksum    string
	KeyFingerprint int32
}

type TL_inputEncryptedFile struct {
	ID         int64
	AccessHash int64
}

type TL_inputEncryptedFileBigUploaded struct {
	ID             int64
	Parts          int32
	KeyFingerprint int32
}

type TL_encryptedMessage struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
	File     TL // EncryptedFile
}

type TL_encryptedMessageService struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
}

type TL_messages_dhConfigNotModified struct {
	Random []byte
}

type TL_messages_dhConfig struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

type TL_messages_sentEncryptedMessage struct {
	Date int32
}

type TL_messages_sentEncryptedFile struct {
	Date int32
	File TL // EncryptedFile
}

type TL_inputDocumentEmpty struct {
}

type TL_inputDocument struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

type TL_documentEmpty struct {
	ID int64
}

type TL_document struct {
	Flags         int32
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	MimeType      string
	Size          int32
	Thumbs        []TL // PhotoSize //flag
	VideoThumbs   []TL // VideoSize //flag
	DcID          int32
	Attributes    []TL // DocumentAttribute
}

type TL_help_support struct {
	PhoneNumber string
	User        TL // User
}

type TL_notifyPeer struct {
	Peer TL // Peer
}

type TL_notifyUsers struct {
}

type TL_notifyChats struct {
}

type TL_notifyBroadcasts struct {
}

type TL_sendMessageTypingAction struct {
}

type TL_sendMessageCancelAction struct {
}

type TL_sendMessageRecordVideoAction struct {
}

type TL_sendMessageUploadVideoAction struct {
	Progress int32
}

type TL_sendMessageRecordAudioAction struct {
}

type TL_sendMessageUploadAudioAction struct {
	Progress int32
}

type TL_sendMessageUploadPhotoAction struct {
	Progress int32
}

type TL_sendMessageUploadDocumentAction struct {
	Progress int32
}

type TL_sendMessageGeoLocationAction struct {
}

type TL_sendMessageChooseContactAction struct {
}

type TL_sendMessageGamePlayAction struct {
}

type TL_sendMessageRecordRoundAction struct {
}

type TL_sendMessageUploadRoundAction struct {
	Progress int32
}

type TL_contacts_found struct {
	MyResults []TL // Peer
	Results   []TL // Peer
	Chats     []TL // Chat
	Users     []TL // User
}

type TL_inputPrivacyKeyStatusTimestamp struct {
}

type TL_inputPrivacyKeyChatInvite struct {
}

type TL_inputPrivacyKeyPhoneCall struct {
}

type TL_inputPrivacyKeyPhoneP2P struct {
}

type TL_inputPrivacyKeyForwards struct {
}

type TL_inputPrivacyKeyProfilePhoto struct {
}

type TL_inputPrivacyKeyPhoneNumber struct {
}

type TL_inputPrivacyKeyAddedByPhone struct {
}

type TL_privacyKeyStatusTimestamp struct {
}

type TL_privacyKeyChatInvite struct {
}

type TL_privacyKeyPhoneCall struct {
}

type TL_privacyKeyPhoneP2P struct {
}

type TL_privacyKeyForwards struct {
}

type TL_privacyKeyProfilePhoto struct {
}

type TL_privacyKeyPhoneNumber struct {
}

type TL_privacyKeyAddedByPhone struct {
}

type TL_inputPrivacyValueAllowContacts struct {
}

type TL_inputPrivacyValueAllowAll struct {
}

type TL_inputPrivacyValueAllowUsers struct {
	Users []TL // InputUser
}

type TL_inputPrivacyValueDisallowContacts struct {
}

type TL_inputPrivacyValueDisallowAll struct {
}

type TL_inputPrivacyValueDisallowUsers struct {
	Users []TL // InputUser
}

type TL_inputPrivacyValueAllowChatParticipants struct {
	Chats []int32
}

type TL_inputPrivacyValueDisallowChatParticipants struct {
	Chats []int32
}

type TL_privacyValueAllowContacts struct {
}

type TL_privacyValueAllowAll struct {
}

type TL_privacyValueAllowUsers struct {
	Users []int32
}

type TL_privacyValueDisallowContacts struct {
}

type TL_privacyValueDisallowAll struct {
}

type TL_privacyValueDisallowUsers struct {
	Users []int32
}

type TL_privacyValueAllowChatParticipants struct {
	Chats []int32
}

type TL_privacyValueDisallowChatParticipants struct {
	Chats []int32
}

type TL_account_privacyRules struct {
	Rules []TL // PrivacyRule
	Chats []TL // Chat
	Users []TL // User
}

type TL_accountDaysTTL struct {
	Days int32
}

type TL_documentAttributeImageSize struct {
	W int32
	H int32
}

type TL_documentAttributeAnimated struct {
}

type TL_documentAttributeSticker struct {
	Flags      int32
	Mask       bool //flag
	Alt        string
	Stickerset TL // InputStickerSet
	MaskCoords TL // MaskCoords //flag
}

type TL_documentAttributeVideo struct {
	Flags             int32
	RoundMessage      bool //flag
	SupportsStreaming bool //flag
	Duration          int32
	W                 int32
	H                 int32
}

type TL_documentAttributeAudio struct {
	Flags     int32
	Voice     bool //flag
	Duration  int32
	Title     string //flag
	Performer string //flag
	Waveform  []byte //flag
}

type TL_documentAttributeFilename struct {
	FileName string
}

type TL_documentAttributeHasStickers struct {
}

type TL_messages_stickersNotModified struct {
}

type TL_messages_stickers struct {
	Hash     int32
	Stickers []TL // Document
}

type TL_stickerPack struct {
	Emoticon  string
	Documents []int64
}

type TL_messages_allStickersNotModified struct {
}

type TL_messages_allStickers struct {
	Hash int32
	Sets []TL // StickerSet
}

type TL_messages_affectedMessages struct {
	Pts      int32
	PtsCount int32
}

type TL_webPageEmpty struct {
	ID int64
}

type TL_webPagePending struct {
	ID   int64
	Date int32
}

type TL_webPage struct {
	Flags       int32
	ID          int64
	Url         string
	DisplayUrl  string
	Hash        int32
	Type        string //flag
	SiteName    string //flag
	Title       string //flag
	Description string //flag
	Photo       TL     // Photo //flag
	EmbedUrl    string //flag
	EmbedType   string //flag
	EmbedWidth  int32  //flag
	EmbedHeight int32  //flag
	Duration    int32  //flag
	Author      string //flag
	Document    TL     // Document //flag
	CachedPage  TL     // Page //flag
	Attributes  []TL   // WebPageAttribute //flag
}

type TL_webPageNotModified struct {
	Flags           int32
	CachedPageViews int32 //flag
}

type TL_authorization struct {
	Flags           int32
	Current         bool //flag
	OfficialApp     bool //flag
	PasswordPending bool //flag
	Hash            int64
	DeviceModel     string
	Platform        string
	SystemVersion   string
	ApiID           int32
	AppName         string
	AppVersion      string
	DateCreated     int32
	DateActive      int32
	Ip              string
	Country         string
	Region          string
}

type TL_account_authorizations struct {
	Authorizations []TL // Authorization
}

type TL_account_password struct {
	Flags                   int32
	HasRecovery             bool   //flag
	HasSecureValues         bool   //flag
	HasPassword             bool   //flag
	CurrentAlgo             TL     // PasswordKdfAlgo //flag
	SrpB                    []byte //flag
	SrpID                   int64  //flag
	Hint                    string //flag
	EmailUnconfirmedPattern string //flag
	NewAlgo                 TL     // PasswordKdfAlgo
	NewSecureAlgo           TL     // SecurePasswordKdfAlgo
	SecureRandom            []byte
}

type TL_account_passwordSettings struct {
	Flags          int32
	Email          string //flag
	SecureSettings TL     // SecureSecretSettings //flag
}

type TL_account_passwordInputSettings struct {
	Flags             int32
	NewAlgo           TL     // PasswordKdfAlgo //flag
	NewPasswordHash   []byte //flag
	Hint              string //flag
	Email             string //flag
	NewSecureSettings TL     // SecureSecretSettings //flag
}

type TL_auth_passwordRecovery struct {
	EmailPattern string
}

type TL_receivedNotifyMessage struct {
	ID    int32
	Flags int32
}

type TL_chatInviteEmpty struct {
}

type TL_chatInviteExported struct {
	Link string
}

type TL_chatInviteAlready struct {
	Chat TL // Chat
}

type TL_chatInvite struct {
	Flags             int32
	Channel           bool //flag
	Broadcast         bool //flag
	Public            bool //flag
	Megagroup         bool //flag
	Title             string
	Photo             TL // Photo
	ParticipantsCount int32
	Participants      []TL // User //flag
}

type TL_inputStickerSetEmpty struct {
}

type TL_inputStickerSetID struct {
	ID         int64
	AccessHash int64
}

type TL_inputStickerSetShortName struct {
	ShortName string
}

type TL_inputStickerSetAnimatedEmoji struct {
}

type TL_inputStickerSetDice struct {
	Emoticon string
}

type TL_stickerSet struct {
	Flags         int32
	Archived      bool  //flag
	Official      bool  //flag
	Masks         bool  //flag
	Animated      bool  //flag
	InstalledDate int32 //flag
	ID            int64
	AccessHash    int64
	Title         string
	ShortName     string
	Thumb         TL    // PhotoSize //flag
	ThumbDcID     int32 //flag
	Count         int32
	Hash          int32
}

type TL_messages_stickerSet struct {
	Set       TL   // StickerSet
	Packs     []TL // StickerPack
	Documents []TL // Document
}

type TL_botCommand struct {
	Command     string
	Description string
}

type TL_botInfo struct {
	UserID      int32
	Description string
	Commands    []TL // BotCommand
}

type TL_keyboardButton struct {
	Text string
}

type TL_keyboardButtonUrl struct {
	Text string
	Url  string
}

type TL_keyboardButtonCallback struct {
	Text string
	Data []byte
}

type TL_keyboardButtonRequestPhone struct {
	Text string
}

type TL_keyboardButtonRequestGeoLocation struct {
	Text string
}

type TL_keyboardButtonSwitchInline struct {
	Flags    int32
	SamePeer bool //flag
	Text     string
	Query    string
}

type TL_keyboardButtonGame struct {
	Text string
}

type TL_keyboardButtonBuy struct {
	Text string
}

type TL_keyboardButtonUrlAuth struct {
	Flags    int32
	Text     string
	FwdText  string //flag
	Url      string
	ButtonID int32
}

type TL_inputKeyboardButtonUrlAuth struct {
	Flags              int32
	RequestWriteAccess bool //flag
	Text               string
	FwdText            string //flag
	Url                string
	Bot                TL // InputUser
}

type TL_keyboardButtonRequestPoll struct {
	Flags int32
	Quiz  TL // Bool //flag
	Text  string
}

type TL_keyboardButtonRow struct {
	Buttons []TL // KeyboardButton
}

type TL_replyKeyboardHide struct {
	Flags     int32
	Selective bool //flag
}

type TL_replyKeyboardForceReply struct {
	Flags     int32
	SingleUse bool //flag
	Selective bool //flag
}

type TL_replyKeyboardMarkup struct {
	Flags     int32
	Resize    bool //flag
	SingleUse bool //flag
	Selective bool //flag
	Rows      []TL // KeyboardButtonRow
}

type TL_replyInlineMarkup struct {
	Rows []TL // KeyboardButtonRow
}

type TL_messageEntityUnknown struct {
	Offset int32
	Length int32
}

type TL_messageEntityMention struct {
	Offset int32
	Length int32
}

type TL_messageEntityHashtag struct {
	Offset int32
	Length int32
}

type TL_messageEntityBotCommand struct {
	Offset int32
	Length int32
}

type TL_messageEntityUrl struct {
	Offset int32
	Length int32
}

type TL_messageEntityEmail struct {
	Offset int32
	Length int32
}

type TL_messageEntityBold struct {
	Offset int32
	Length int32
}

type TL_messageEntityItalic struct {
	Offset int32
	Length int32
}

type TL_messageEntityCode struct {
	Offset int32
	Length int32
}

type TL_messageEntityPre struct {
	Offset   int32
	Length   int32
	Language string
}

type TL_messageEntityTextUrl struct {
	Offset int32
	Length int32
	Url    string
}

type TL_messageEntityMentionName struct {
	Offset int32
	Length int32
	UserID int32
}

type TL_inputMessageEntityMentionName struct {
	Offset int32
	Length int32
	UserID TL // InputUser
}

type TL_messageEntityPhone struct {
	Offset int32
	Length int32
}

type TL_messageEntityCashtag struct {
	Offset int32
	Length int32
}

type TL_messageEntityUnderline struct {
	Offset int32
	Length int32
}

type TL_messageEntityStrike struct {
	Offset int32
	Length int32
}

type TL_messageEntityBlockquote struct {
	Offset int32
	Length int32
}

type TL_messageEntityBankCard struct {
	Offset int32
	Length int32
}

type TL_inputChannelEmpty struct {
}

type TL_inputChannel struct {
	ChannelID  int32
	AccessHash int64
}

type TL_inputChannelFromMessage struct {
	Peer      TL // InputPeer
	MsgID     int32
	ChannelID int32
}

type TL_contacts_resolvedPeer struct {
	Peer  TL   // Peer
	Chats []TL // Chat
	Users []TL // User
}

type TL_messageRange struct {
	MinID int32
	MaxID int32
}

type TL_updates_channelDifferenceEmpty struct {
	Flags   int32
	Final   bool //flag
	Pts     int32
	Timeout int32 //flag
}

type TL_updates_channelDifferenceTooLong struct {
	Flags    int32
	Final    bool  //flag
	Timeout  int32 //flag
	Dialog   TL    // Dialog
	Messages []TL  // Message
	Chats    []TL  // Chat
	Users    []TL  // User
}

type TL_updates_channelDifference struct {
	Flags        int32
	Final        bool //flag
	Pts          int32
	Timeout      int32 //flag
	NewMessages  []TL  // Message
	OtherUpdates []TL  // Update
	Chats        []TL  // Chat
	Users        []TL  // User
}

type TL_channelMessagesFilterEmpty struct {
}

type TL_channelMessagesFilter struct {
	Flags              int32
	ExcludeNewMessages bool //flag
	Ranges             []TL // MessageRange
}

type TL_channelParticipant struct {
	UserID int32
	Date   int32
}

type TL_channelParticipantSelf struct {
	UserID    int32
	InviterID int32
	Date      int32
}

type TL_channelParticipantCreator struct {
	Flags  int32
	UserID int32
	Rank   string //flag
}

type TL_channelParticipantAdmin struct {
	Flags       int32
	CanEdit     bool //flag
	Self        bool //flag
	UserID      int32
	InviterID   int32 //flag
	PromotedBy  int32
	Date        int32
	AdminRights TL     // ChatAdminRights
	Rank        string //flag
}

type TL_channelParticipantBanned struct {
	Flags        int32
	Left         bool //flag
	UserID       int32
	KickedBy     int32
	Date         int32
	BannedRights TL // ChatBannedRights
}

type TL_channelParticipantsRecent struct {
}

type TL_channelParticipantsAdmins struct {
}

type TL_channelParticipantsKicked struct {
	Q string
}

type TL_channelParticipantsBots struct {
}

type TL_channelParticipantsBanned struct {
	Q string
}

type TL_channelParticipantsSearch struct {
	Q string
}

type TL_channelParticipantsContacts struct {
	Q string
}

type TL_channels_channelParticipants struct {
	Count        int32
	Participants []TL // ChannelParticipant
	Users        []TL // User
}

type TL_channels_channelParticipantsNotModified struct {
}

type TL_channels_channelParticipant struct {
	Participant TL   // ChannelParticipant
	Users       []TL // User
}

type TL_help_termsOfService struct {
	Flags         int32
	Popup         bool //flag
	ID            TL   // DataJSON
	Text          string
	Entities      []TL  // MessageEntity
	MinAgeConfirm int32 //flag
}

type TL_foundGif struct {
	Url         string
	ThumbUrl    string
	ContentUrl  string
	ContentType string
	W           int32
	H           int32
}

type TL_foundGifCached struct {
	Url      string
	Photo    TL // Photo
	Document TL // Document
}

type TL_messages_foundGifs struct {
	NextOffset int32
	Results    []TL // FoundGif
}

type TL_messages_savedGifsNotModified struct {
}

type TL_messages_savedGifs struct {
	Hash int32
	Gifs []TL // Document
}

type TL_inputBotInlineMessageMediaAuto struct {
	Flags       int32
	Message     string
	Entities    []TL // MessageEntity //flag
	ReplyMarkup TL   // ReplyMarkup //flag
}

type TL_inputBotInlineMessageText struct {
	Flags       int32
	NoWebpage   bool //flag
	Message     string
	Entities    []TL // MessageEntity //flag
	ReplyMarkup TL   // ReplyMarkup //flag
}

type TL_inputBotInlineMessageMediaGeo struct {
	Flags       int32
	GeoPoint    TL // InputGeoPoint
	Period      int32
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_inputBotInlineMessageMediaVenue struct {
	Flags       int32
	GeoPoint    TL // InputGeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_inputBotInlineMessageMediaContact struct {
	Flags       int32
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_inputBotInlineMessageGame struct {
	Flags       int32
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_inputBotInlineResult struct {
	Flags       int32
	ID          string
	Type        string
	Title       string //flag
	Description string //flag
	Url         string //flag
	Thumb       TL     // InputWebDocument //flag
	Content     TL     // InputWebDocument //flag
	SendMessage TL     // InputBotInlineMessage
}

type TL_inputBotInlineResultPhoto struct {
	ID          string
	Type        string
	Photo       TL // InputPhoto
	SendMessage TL // InputBotInlineMessage
}

type TL_inputBotInlineResultDocument struct {
	Flags       int32
	ID          string
	Type        string
	Title       string //flag
	Description string //flag
	Document    TL     // InputDocument
	SendMessage TL     // InputBotInlineMessage
}

type TL_inputBotInlineResultGame struct {
	ID          string
	ShortName   string
	SendMessage TL // InputBotInlineMessage
}

type TL_botInlineMessageMediaAuto struct {
	Flags       int32
	Message     string
	Entities    []TL // MessageEntity //flag
	ReplyMarkup TL   // ReplyMarkup //flag
}

type TL_botInlineMessageText struct {
	Flags       int32
	NoWebpage   bool //flag
	Message     string
	Entities    []TL // MessageEntity //flag
	ReplyMarkup TL   // ReplyMarkup //flag
}

type TL_botInlineMessageMediaGeo struct {
	Flags       int32
	Geo         TL // GeoPoint
	Period      int32
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_botInlineMessageMediaVenue struct {
	Flags       int32
	Geo         TL // GeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_botInlineMessageMediaContact struct {
	Flags       int32
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup TL // ReplyMarkup //flag
}

type TL_botInlineResult struct {
	Flags       int32
	ID          string
	Type        string
	Title       string //flag
	Description string //flag
	Url         string //flag
	Thumb       TL     // WebDocument //flag
	Content     TL     // WebDocument //flag
	SendMessage TL     // BotInlineMessage
}

type TL_botInlineMediaResult struct {
	Flags       int32
	ID          string
	Type        string
	Photo       TL     // Photo //flag
	Document    TL     // Document //flag
	Title       string //flag
	Description string //flag
	SendMessage TL     // BotInlineMessage
}

type TL_messages_botResults struct {
	Flags      int32
	Gallery    bool //flag
	QueryID    int64
	NextOffset string //flag
	SwitchPm   TL     // InlineBotSwitchPM //flag
	Results    []TL   // BotInlineResult
	CacheTime  int32
	Users      []TL // User
}

type TL_exportedMessageLink struct {
	Link string
	Html string
}

type TL_messageFwdHeader struct {
	Flags          int32
	FromID         int32  //flag
	FromName       string //flag
	Date           int32
	ChannelID      int32  //flag
	ChannelPost    int32  //flag
	PostAuthor     string //flag
	SavedFromPeer  TL     // Peer //flag
	SavedFromMsgID int32  //flag
	PsaType        string //flag
}

type TL_auth_codeTypeSms struct {
}

type TL_auth_codeTypeCall struct {
}

type TL_auth_codeTypeFlashCall struct {
}

type TL_auth_sentCodeTypeApp struct {
	Length int32
}

type TL_auth_sentCodeTypeSms struct {
	Length int32
}

type TL_auth_sentCodeTypeCall struct {
	Length int32
}

type TL_auth_sentCodeTypeFlashCall struct {
	Pattern string
}

type TL_messages_botCallbackAnswer struct {
	Flags     int32
	Alert     bool   //flag
	HasUrl    bool   //flag
	NativeUi  bool   //flag
	Message   string //flag
	Url       string //flag
	CacheTime int32
}

type TL_messages_messageEditData struct {
	Flags   int32
	Caption bool //flag
}

type TL_inputBotInlineMessageID struct {
	DcID       int32
	ID         int64
	AccessHash int64
}

type TL_inlineBotSwitchPM struct {
	Text       string
	StartParam string
}

type TL_messages_peerDialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
	State    TL   // updates_State
}

type TL_topPeer struct {
	Peer   TL // Peer
	Rating float64
}

type TL_topPeerCategoryBotsPM struct {
}

type TL_topPeerCategoryBotsInline struct {
}

type TL_topPeerCategoryCorrespondents struct {
}

type TL_topPeerCategoryGroups struct {
}

type TL_topPeerCategoryChannels struct {
}

type TL_topPeerCategoryPhoneCalls struct {
}

type TL_topPeerCategoryForwardUsers struct {
}

type TL_topPeerCategoryForwardChats struct {
}

type TL_topPeerCategoryPeers struct {
	Category TL // TopPeerCategory
	Count    int32
	Peers    []TL // TopPeer
}

type TL_contacts_topPeersNotModified struct {
}

type TL_contacts_topPeers struct {
	Categories []TL // TopPeerCategoryPeers
	Chats      []TL // Chat
	Users      []TL // User
}

type TL_contacts_topPeersDisabled struct {
}

type TL_draftMessageEmpty struct {
	Flags int32
	Date  int32 //flag
}

type TL_draftMessage struct {
	Flags        int32
	NoWebpage    bool  //flag
	ReplyToMsgID int32 //flag
	Message      string
	Entities     []TL // MessageEntity //flag
	Date         int32
}

type TL_messages_featuredStickersNotModified struct {
	Count int32
}

type TL_messages_featuredStickers struct {
	Hash   int32
	Count  int32
	Sets   []TL // StickerSetCovered
	Unread []int64
}

type TL_messages_recentStickersNotModified struct {
}

type TL_messages_recentStickers struct {
	Hash     int32
	Packs    []TL // StickerPack
	Stickers []TL // Document
	Dates    []int32
}

type TL_messages_archivedStickers struct {
	Count int32
	Sets  []TL // StickerSetCovered
}

type TL_messages_stickerSetInstallResultSuccess struct {
}

type TL_messages_stickerSetInstallResultArchive struct {
	Sets []TL // StickerSetCovered
}

type TL_stickerSetCovered struct {
	Set   TL // StickerSet
	Cover TL // Document
}

type TL_stickerSetMultiCovered struct {
	Set    TL   // StickerSet
	Covers []TL // Document
}

type TL_maskCoords struct {
	N    int32
	X    float64
	Y    float64
	Zoom float64
}

type TL_inputStickeredMediaPhoto struct {
	ID TL // InputPhoto
}

type TL_inputStickeredMediaDocument struct {
	ID TL // InputDocument
}

type TL_game struct {
	Flags       int32
	ID          int64
	AccessHash  int64
	ShortName   string
	Title       string
	Description string
	Photo       TL // Photo
	Document    TL // Document //flag
}

type TL_inputGameID struct {
	ID         int64
	AccessHash int64
}

type TL_inputGameShortName struct {
	BotID     TL // InputUser
	ShortName string
}

type TL_highScore struct {
	Pos    int32
	UserID int32
	Score  int32
}

type TL_messages_highScores struct {
	Scores []TL // HighScore
	Users  []TL // User
}

type TL_textEmpty struct {
}

type TL_textPlain struct {
	Text string
}

type TL_textBold struct {
	Text TL // RichText
}

type TL_textItalic struct {
	Text TL // RichText
}

type TL_textUnderline struct {
	Text TL // RichText
}

type TL_textStrike struct {
	Text TL // RichText
}

type TL_textFixed struct {
	Text TL // RichText
}

type TL_textUrl struct {
	Text      TL // RichText
	Url       string
	WebpageID int64
}

type TL_textEmail struct {
	Text  TL // RichText
	Email string
}

type TL_textConcat struct {
	Texts []TL // RichText
}

type TL_textSubscript struct {
	Text TL // RichText
}

type TL_textSuperscript struct {
	Text TL // RichText
}

type TL_textMarked struct {
	Text TL // RichText
}

type TL_textPhone struct {
	Text  TL // RichText
	Phone string
}

type TL_textImage struct {
	DocumentID int64
	W          int32
	H          int32
}

type TL_textAnchor struct {
	Text TL // RichText
	Name string
}

type TL_pageBlockUnsupported struct {
}

type TL_pageBlockTitle struct {
	Text TL // RichText
}

type TL_pageBlockSubtitle struct {
	Text TL // RichText
}

type TL_pageBlockAuthorDate struct {
	Author        TL // RichText
	PublishedDate int32
}

type TL_pageBlockHeader struct {
	Text TL // RichText
}

type TL_pageBlockSubheader struct {
	Text TL // RichText
}

type TL_pageBlockParagraph struct {
	Text TL // RichText
}

type TL_pageBlockPreformatted struct {
	Text     TL // RichText
	Language string
}

type TL_pageBlockFooter struct {
	Text TL // RichText
}

type TL_pageBlockDivider struct {
}

type TL_pageBlockAnchor struct {
	Name string
}

type TL_pageBlockList struct {
	Items []TL // PageListItem
}

type TL_pageBlockBlockquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type TL_pageBlockPullquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type TL_pageBlockPhoto struct {
	Flags     int32
	PhotoID   int64
	Caption   TL     // PageCaption
	Url       string //flag
	WebpageID int64  //flag
}

type TL_pageBlockVideo struct {
	Flags    int32
	Autoplay bool //flag
	Loop     bool //flag
	VideoID  int64
	Caption  TL // PageCaption
}

type TL_pageBlockCover struct {
	Cover TL // PageBlock
}

type TL_pageBlockEmbed struct {
	Flags          int32
	FullWidth      bool   //flag
	AllowScrolling bool   //flag
	Url            string //flag
	Html           string //flag
	PosterPhotoID  int64  //flag
	W              int32  //flag
	H              int32  //flag
	Caption        TL     // PageCaption
}

type TL_pageBlockEmbedPost struct {
	Url           string
	WebpageID     int64
	AuthorPhotoID int64
	Author        string
	Date          int32
	Blocks        []TL // PageBlock
	Caption       TL   // PageCaption
}

type TL_pageBlockCollage struct {
	Items   []TL // PageBlock
	Caption TL   // PageCaption
}

type TL_pageBlockSlideshow struct {
	Items   []TL // PageBlock
	Caption TL   // PageCaption
}

type TL_pageBlockChannel struct {
	Channel TL // Chat
}

type TL_pageBlockAudio struct {
	AudioID int64
	Caption TL // PageCaption
}

type TL_pageBlockKicker struct {
	Text TL // RichText
}

type TL_pageBlockTable struct {
	Flags    int32
	Bordered bool //flag
	Striped  bool //flag
	Title    TL   // RichText
	Rows     []TL // PageTableRow
}

type TL_pageBlockOrderedList struct {
	Items []TL // PageListOrderedItem
}

type TL_pageBlockDetails struct {
	Flags  int32
	Open   bool //flag
	Blocks []TL // PageBlock
	Title  TL   // RichText
}

type TL_pageBlockRelatedArticles struct {
	Title    TL   // RichText
	Articles []TL // PageRelatedArticle
}

type TL_pageBlockMap struct {
	Geo     TL // GeoPoint
	Zoom    int32
	W       int32
	H       int32
	Caption TL // PageCaption
}

type TL_phoneCallDiscardReasonMissed struct {
}

type TL_phoneCallDiscardReasonDisconnect struct {
}

type TL_phoneCallDiscardReasonHangup struct {
}

type TL_phoneCallDiscardReasonBusy struct {
}

type TL_dataJSON struct {
	Data string
}

type TL_labeledPrice struct {
	Label  string
	Amount int64
}

type TL_invoice struct {
	Flags                    int32
	Test                     bool //flag
	NameRequested            bool //flag
	PhoneRequested           bool //flag
	EmailRequested           bool //flag
	ShippingAddressRequested bool //flag
	Flexible                 bool //flag
	PhoneToProvider          bool //flag
	EmailToProvider          bool //flag
	Currency                 string
	Prices                   []TL // LabeledPrice
}

type TL_paymentCharge struct {
	ID               string
	ProviderChargeID string
}

type TL_postAddress struct {
	StreetLine1 string
	StreetLine2 string
	City        string
	State       string
	CountryIso2 string
	PostCode    string
}

type TL_paymentRequestedInfo struct {
	Flags           int32
	Name            string //flag
	Phone           string //flag
	Email           string //flag
	ShippingAddress TL     // PostAddress //flag
}

type TL_paymentSavedCredentialsCard struct {
	ID    string
	Title string
}

type TL_webDocument struct {
	Url        string
	AccessHash int64
	Size       int32
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type TL_webDocumentNoProxy struct {
	Url        string
	Size       int32
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type TL_inputWebDocument struct {
	Url        string
	Size       int32
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type TL_inputWebFileLocation struct {
	Url        string
	AccessHash int64
}

type TL_inputWebFileGeoPointLocation struct {
	GeoPoint   TL // InputGeoPoint
	AccessHash int64
	W          int32
	H          int32
	Zoom       int32
	Scale      int32
}

type TL_upload_webFile struct {
	Size     int32
	MimeType string
	FileType TL // storage_FileType
	Mtime    int32
	Bytes    []byte
}

type TL_payments_paymentForm struct {
	Flags              int32
	CanSaveCredentials bool //flag
	PasswordMissing    bool //flag
	BotID              int32
	Invoice            TL // Invoice
	ProviderID         int32
	Url                string
	NativeProvider     string //flag
	NativeParams       TL     // DataJSON //flag
	SavedInfo          TL     // PaymentRequestedInfo //flag
	SavedCredentials   TL     // PaymentSavedCredentials //flag
	Users              []TL   // User
}

type TL_payments_validatedRequestedInfo struct {
	Flags           int32
	ID              string //flag
	ShippingOptions []TL   // ShippingOption //flag
}

type TL_payments_paymentResult struct {
	Updates TL // Updates
}

type TL_payments_paymentVerificationNeeded struct {
	Url string
}

type TL_payments_paymentReceipt struct {
	Flags            int32
	Date             int32
	BotID            int32
	Invoice          TL // Invoice
	ProviderID       int32
	Info             TL // PaymentRequestedInfo //flag
	Shipping         TL // ShippingOption //flag
	Currency         string
	TotalAmount      int64
	CredentialsTitle string
	Users            []TL // User
}

type TL_payments_savedInfo struct {
	Flags               int32
	HasSavedCredentials bool //flag
	SavedInfo           TL   // PaymentRequestedInfo //flag
}

type TL_inputPaymentCredentialsSaved struct {
	ID          string
	TmpPassword []byte
}

type TL_inputPaymentCredentials struct {
	Flags int32
	Save  bool //flag
	Data  TL   // DataJSON
}

type TL_inputPaymentCredentialsApplePay struct {
	PaymentData TL // DataJSON
}

type TL_inputPaymentCredentialsAndroidPay struct {
	PaymentToken        TL // DataJSON
	GoogleTransactionID string
}

type TL_account_tmpPassword struct {
	TmpPassword []byte
	ValidUntil  int32
}

type TL_shippingOption struct {
	ID     string
	Title  string
	Prices []TL // LabeledPrice
}

type TL_inputStickerSetItem struct {
	Flags      int32
	Document   TL // InputDocument
	Emoji      string
	MaskCoords TL // MaskCoords //flag
}

type TL_inputPhoneCall struct {
	ID         int64
	AccessHash int64
}

type TL_phoneCallEmpty struct {
	ID int64
}

type TL_phoneCallWaiting struct {
	Flags         int32
	Video         bool //flag
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	Protocol      TL    // PhoneCallProtocol
	ReceiveDate   int32 //flag
}

type TL_phoneCallRequested struct {
	Flags         int32
	Video         bool //flag
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	GAHash        []byte
	Protocol      TL // PhoneCallProtocol
}

type TL_phoneCallAccepted struct {
	Flags         int32
	Video         bool //flag
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	GB            []byte
	Protocol      TL // PhoneCallProtocol
}

type TL_phoneCall struct {
	Flags          int32
	P2pAllowed     bool //flag
	ID             int64
	AccessHash     int64
	Date           int32
	AdminID        int32
	ParticipantID  int32
	GAOrB          []byte
	KeyFingerprint int64
	Protocol       TL   // PhoneCallProtocol
	Connections    []TL // PhoneConnection
	StartDate      int32
}

type TL_phoneCallDiscarded struct {
	Flags      int32
	NeedRating bool //flag
	NeedDebug  bool //flag
	Video      bool //flag
	ID         int64
	Reason     TL    // PhoneCallDiscardReason //flag
	Duration   int32 //flag
}

type TL_phoneConnection struct {
	ID      int64
	Ip      string
	Ipv6    string
	Port    int32
	PeerTag []byte
}

type TL_phoneCallProtocol struct {
	Flags           int32
	UdpP2p          bool //flag
	UdpReflector    bool //flag
	MinLayer        int32
	MaxLayer        int32
	LibraryVersions []string
}

type TL_phone_phoneCall struct {
	PhoneCall TL   // PhoneCall
	Users     []TL // User
}

type TL_upload_cdnFileReuploadNeeded struct {
	RequestToken []byte
}

type TL_upload_cdnFile struct {
	Bytes []byte
}

type TL_cdnPublicKey struct {
	DcID      int32
	PublicKey string
}

type TL_cdnConfig struct {
	PublicKeys []TL // CdnPublicKey
}

type TL_langPackString struct {
	Key   string
	Value string
}

type TL_langPackStringPluralized struct {
	Flags      int32
	Key        string
	ZeroValue  string //flag
	OneValue   string //flag
	TwoValue   string //flag
	FewValue   string //flag
	ManyValue  string //flag
	OtherValue string
}

type TL_langPackStringDeleted struct {
	Key string
}

type TL_langPackDifference struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Strings     []TL // LangPackString
}

type TL_langPackLanguage struct {
	Flags           int32
	Official        bool //flag
	Rtl             bool //flag
	Beta            bool //flag
	Name            string
	NativeName      string
	LangCode        string
	BaseLangCode    string //flag
	PluralCode      string
	StringsCount    int32
	TranslatedCount int32
	TranslationsUrl string
}

type TL_channelAdminLogEventActionChangeTitle struct {
	PrevValue string
	NewValue  string
}

type TL_channelAdminLogEventActionChangeAbout struct {
	PrevValue string
	NewValue  string
}

type TL_channelAdminLogEventActionChangeUsername struct {
	PrevValue string
	NewValue  string
}

type TL_channelAdminLogEventActionChangePhoto struct {
	PrevPhoto TL // Photo
	NewPhoto  TL // Photo
}

type TL_channelAdminLogEventActionToggleInvites struct {
	NewValue TL // Bool
}

type TL_channelAdminLogEventActionToggleSignatures struct {
	NewValue TL // Bool
}

type TL_channelAdminLogEventActionUpdatePinned struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionEditMessage struct {
	PrevMessage TL // Message
	NewMessage  TL // Message
}

type TL_channelAdminLogEventActionDeleteMessage struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionParticipantJoin struct {
}

type TL_channelAdminLogEventActionParticipantLeave struct {
}

type TL_channelAdminLogEventActionParticipantInvite struct {
	Participant TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleBan struct {
	PrevParticipant TL // ChannelParticipant
	NewParticipant  TL // ChannelParticipant
}

type TL_channelAdminLogEventActionParticipantToggleAdmin struct {
	PrevParticipant TL // ChannelParticipant
	NewParticipant  TL // ChannelParticipant
}

type TL_channelAdminLogEventActionChangeStickerSet struct {
	PrevStickerset TL // InputStickerSet
	NewStickerset  TL // InputStickerSet
}

type TL_channelAdminLogEventActionTogglePreHistoryHidden struct {
	NewValue TL // Bool
}

type TL_channelAdminLogEventActionDefaultBannedRights struct {
	PrevBannedRights TL // ChatBannedRights
	NewBannedRights  TL // ChatBannedRights
}

type TL_channelAdminLogEventActionStopPoll struct {
	Message TL // Message
}

type TL_channelAdminLogEventActionChangeLinkedChat struct {
	PrevValue int32
	NewValue  int32
}

type TL_channelAdminLogEventActionChangeLocation struct {
	PrevValue TL // ChannelLocation
	NewValue  TL // ChannelLocation
}

type TL_channelAdminLogEventActionToggleSlowMode struct {
	PrevValue int32
	NewValue  int32
}

type TL_channelAdminLogEvent struct {
	ID     int64
	Date   int32
	UserID int32
	Action TL // ChannelAdminLogEventAction
}

type TL_channels_adminLogResults struct {
	Events []TL // ChannelAdminLogEvent
	Chats  []TL // Chat
	Users  []TL // User
}

type TL_channelAdminLogEventsFilter struct {
	Flags    int32
	Join     bool //flag
	Leave    bool //flag
	Invite   bool //flag
	Ban      bool //flag
	Unban    bool //flag
	Kick     bool //flag
	Unkick   bool //flag
	Promote  bool //flag
	Demote   bool //flag
	Info     bool //flag
	Settings bool //flag
	Pinned   bool //flag
	Edit     bool //flag
	Delete   bool //flag
}

type TL_popularContact struct {
	ClientID  int64
	Importers int32
}

type TL_messages_favedStickersNotModified struct {
}

type TL_messages_favedStickers struct {
	Hash     int32
	Packs    []TL // StickerPack
	Stickers []TL // Document
}

type TL_recentMeUrlUnknown struct {
	Url string
}

type TL_recentMeUrlUser struct {
	Url    string
	UserID int32
}

type TL_recentMeUrlChat struct {
	Url    string
	ChatID int32
}

type TL_recentMeUrlChatInvite struct {
	Url        string
	ChatInvite TL // ChatInvite
}

type TL_recentMeUrlStickerSet struct {
	Url string
	Set TL // StickerSetCovered
}

type TL_help_recentMeUrls struct {
	Urls  []TL // RecentMeUrl
	Chats []TL // Chat
	Users []TL // User
}

type TL_inputSingleMedia struct {
	Flags    int32
	Media    TL // InputMedia
	RandomID int64
	Message  string
	Entities []TL // MessageEntity //flag
}

type TL_webAuthorization struct {
	Hash        int64
	BotID       int32
	Domain      string
	Browser     string
	Platform    string
	DateCreated int32
	DateActive  int32
	Ip          string
	Region      string
}

type TL_account_webAuthorizations struct {
	Authorizations []TL // WebAuthorization
	Users          []TL // User
}

type TL_inputMessageID struct {
	ID int32
}

type TL_inputMessageReplyTo struct {
	ID int32
}

type TL_inputMessagePinned struct {
}

type TL_inputDialogPeer struct {
	Peer TL // InputPeer
}

type TL_inputDialogPeerFolder struct {
	FolderID int32
}

type TL_dialogPeer struct {
	Peer TL // Peer
}

type TL_dialogPeerFolder struct {
	FolderID int32
}

type TL_messages_foundStickerSetsNotModified struct {
}

type TL_messages_foundStickerSets struct {
	Hash int32
	Sets []TL // StickerSetCovered
}

type TL_fileHash struct {
	Offset int32
	Limit  int32
	Hash   []byte
}

type TL_inputClientProxy struct {
	Address string
	Port    int32
}

type TL_help_termsOfServiceUpdateEmpty struct {
	Expires int32
}

type TL_help_termsOfServiceUpdate struct {
	Expires        int32
	TermsOfService TL // help_TermsOfService
}

type TL_inputSecureFileUploaded struct {
	ID          int64
	Parts       int32
	Md5Checksum string
	FileHash    []byte
	Secret      []byte
}

type TL_inputSecureFile struct {
	ID         int64
	AccessHash int64
}

type TL_secureFileEmpty struct {
}

type TL_secureFile struct {
	ID         int64
	AccessHash int64
	Size       int32
	DcID       int32
	Date       int32
	FileHash   []byte
	Secret     []byte
}

type TL_secureData struct {
	Data     []byte
	DataHash []byte
	Secret   []byte
}

type TL_securePlainPhone struct {
	Phone string
}

type TL_securePlainEmail struct {
	Email string
}

type TL_secureValueTypePersonalDetails struct {
}

type TL_secureValueTypePassport struct {
}

type TL_secureValueTypeDriverLicense struct {
}

type TL_secureValueTypeIdentityCard struct {
}

type TL_secureValueTypeInternalPassport struct {
}

type TL_secureValueTypeAddress struct {
}

type TL_secureValueTypeUtilityBill struct {
}

type TL_secureValueTypeBankStatement struct {
}

type TL_secureValueTypeRentalAgreement struct {
}

type TL_secureValueTypePassportRegistration struct {
}

type TL_secureValueTypeTemporaryRegistration struct {
}

type TL_secureValueTypePhone struct {
}

type TL_secureValueTypeEmail struct {
}

type TL_secureValue struct {
	Flags       int32
	Type        TL   // SecureValueType
	Data        TL   // SecureData //flag
	FrontSide   TL   // SecureFile //flag
	ReverseSide TL   // SecureFile //flag
	Selfie      TL   // SecureFile //flag
	Translation []TL // SecureFile //flag
	Files       []TL // SecureFile //flag
	PlainData   TL   // SecurePlainData //flag
	Hash        []byte
}

type TL_inputSecureValue struct {
	Flags       int32
	Type        TL   // SecureValueType
	Data        TL   // SecureData //flag
	FrontSide   TL   // InputSecureFile //flag
	ReverseSide TL   // InputSecureFile //flag
	Selfie      TL   // InputSecureFile //flag
	Translation []TL // InputSecureFile //flag
	Files       []TL // InputSecureFile //flag
	PlainData   TL   // SecurePlainData //flag
}

type TL_secureValueHash struct {
	Type TL // SecureValueType
	Hash []byte
}

type TL_secureValueErrorData struct {
	Type     TL // SecureValueType
	DataHash []byte
	Field    string
	Text     string
}

type TL_secureValueErrorFrontSide struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type TL_secureValueErrorReverseSide struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type TL_secureValueErrorSelfie struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type TL_secureValueErrorFile struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type TL_secureValueErrorFiles struct {
	Type     TL   // SecureValueType
	FileHash []TL // bytes
	Text     string
}

type TL_secureValueError struct {
	Type TL // SecureValueType
	Hash []byte
	Text string
}

type TL_secureValueErrorTranslationFile struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type TL_secureValueErrorTranslationFiles struct {
	Type     TL   // SecureValueType
	FileHash []TL // bytes
	Text     string
}

type TL_secureCredentialsEncrypted struct {
	Data   []byte
	Hash   []byte
	Secret []byte
}

type TL_account_authorizationForm struct {
	Flags            int32
	RequiredTypes    []TL   // SecureRequiredType
	Values           []TL   // SecureValue
	Errors           []TL   // SecureValueError
	Users            []TL   // User
	PrivacyPolicyUrl string //flag
}

type TL_account_sentEmailCode struct {
	EmailPattern string
	Length       int32
}

type TL_help_deepLinkInfoEmpty struct {
}

type TL_help_deepLinkInfo struct {
	Flags     int32
	UpdateApp bool //flag
	Message   string
	Entities  []TL // MessageEntity //flag
}

type TL_savedPhoneContact struct {
	Phone     string
	FirstName string
	LastName  string
	Date      int32
}

type TL_account_takeout struct {
	ID int64
}

type TL_passwordKdfAlgoUnknown struct {
}

type TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

type TL_securePasswordKdfAlgoUnknown struct {
}

type TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000 struct {
	Salt []byte
}

type TL_securePasswordKdfAlgoSHA512 struct {
	Salt []byte
}

type TL_secureSecretSettings struct {
	SecureAlgo     TL // SecurePasswordKdfAlgo
	SecureSecret   []byte
	SecureSecretID int64
}

type TL_inputCheckPasswordEmpty struct {
}

type TL_inputCheckPasswordSRP struct {
	SrpID int64
	A     []byte
	M1    []byte
}

type TL_secureRequiredType struct {
	Flags               int32
	NativeNames         bool //flag
	SelfieRequired      bool //flag
	TranslationRequired bool //flag
	Type                TL   // SecureValueType
}

type TL_secureRequiredTypeOneOf struct {
	Types []TL // SecureRequiredType
}

type TL_help_passportConfigNotModified struct {
}

type TL_help_passportConfig struct {
	Hash           int32
	CountriesLangs TL // DataJSON
}

type TL_inputAppEvent struct {
	Time float64
	Type string
	Peer int64
	Data TL // JSONValue
}

type TL_jsonObjectValue struct {
	Key   string
	Value TL // JSONValue
}

type TL_jsonNull struct {
}

type TL_jsonBool struct {
	Value TL // Bool
}

type TL_jsonNumber struct {
	Value float64
}

type TL_jsonString struct {
	Value string
}

type TL_jsonArray struct {
	Value []TL // JSONValue
}

type TL_jsonObject struct {
	Value []TL // JSONObjectValue
}

type TL_pageTableCell struct {
	Flags        int32
	Header       bool  //flag
	AlignCenter  bool  //flag
	AlignRight   bool  //flag
	ValignMiddle bool  //flag
	ValignBottom bool  //flag
	Text         TL    // RichText //flag
	Colspan      int32 //flag
	Rowspan      int32 //flag
}

type TL_pageTableRow struct {
	Cells []TL // PageTableCell
}

type TL_pageCaption struct {
	Text   TL // RichText
	Credit TL // RichText
}

type TL_pageListItemText struct {
	Text TL // RichText
}

type TL_pageListItemBlocks struct {
	Blocks []TL // PageBlock
}

type TL_pageListOrderedItemText struct {
	Num  string
	Text TL // RichText
}

type TL_pageListOrderedItemBlocks struct {
	Num    string
	Blocks []TL // PageBlock
}

type TL_pageRelatedArticle struct {
	Flags         int32
	Url           string
	WebpageID     int64
	Title         string //flag
	Description   string //flag
	PhotoID       int64  //flag
	Author        string //flag
	PublishedDate int32  //flag
}

type TL_page struct {
	Flags     int32
	Part      bool //flag
	Rtl       bool //flag
	V2        bool //flag
	Url       string
	Blocks    []TL  // PageBlock
	Photos    []TL  // Photo
	Documents []TL  // Document
	Views     int32 //flag
}

type TL_help_supportName struct {
	Name string
}

type TL_help_userInfoEmpty struct {
}

type TL_help_userInfo struct {
	Message  string
	Entities []TL // MessageEntity
	Author   string
	Date     int32
}

type TL_pollAnswer struct {
	Text   string
	Option []byte
}

type TL_poll struct {
	ID             int64
	Flags          int32
	Closed         bool //flag
	PublicVoters   bool //flag
	MultipleChoice bool //flag
	Quiz           bool //flag
	Question       string
	Answers        []TL  // PollAnswer
	ClosePeriod    int32 //flag
	CloseDate      int32 //flag
}

type TL_pollAnswerVoters struct {
	Flags   int32
	Chosen  bool //flag
	Correct bool //flag
	Option  []byte
	Voters  int32
}

type TL_pollResults struct {
	Flags            int32
	Min              bool    //flag
	Results          []TL    // PollAnswerVoters //flag
	TotalVoters      int32   //flag
	RecentVoters     []int32 //flag
	Solution         string  //flag
	SolutionEntities []TL    // MessageEntity //flag
}

type TL_chatOnlines struct {
	Onlines int32
}

type TL_statsURL struct {
	Url string
}

type TL_chatAdminRights struct {
	Flags          int32
	ChangeInfo     bool //flag
	PostMessages   bool //flag
	EditMessages   bool //flag
	DeleteMessages bool //flag
	BanUsers       bool //flag
	InviteUsers    bool //flag
	PinMessages    bool //flag
	AddAdmins      bool //flag
}

type TL_chatBannedRights struct {
	Flags        int32
	ViewMessages bool //flag
	SendMessages bool //flag
	SendMedia    bool //flag
	SendStickers bool //flag
	SendGifs     bool //flag
	SendGames    bool //flag
	SendInline   bool //flag
	EmbedLinks   bool //flag
	SendPolls    bool //flag
	ChangeInfo   bool //flag
	InviteUsers  bool //flag
	PinMessages  bool //flag
	UntilDate    int32
}

type TL_inputWallPaper struct {
	ID         int64
	AccessHash int64
}

type TL_inputWallPaperSlug struct {
	Slug string
}

type TL_inputWallPaperNoFile struct {
}

type TL_account_wallPapersNotModified struct {
}

type TL_account_wallPapers struct {
	Hash       int32
	Wallpapers []TL // WallPaper
}

type TL_codeSettings struct {
	Flags          int32
	AllowFlashcall bool //flag
	CurrentNumber  bool //flag
	AllowAppHash   bool //flag
}

type TL_wallPaperSettings struct {
	Flags                 int32
	Blur                  bool  //flag
	Motion                bool  //flag
	BackgroundColor       int32 //flag
	SecondBackgroundColor int32 //flag
	Intensity             int32 //flag
	Rotation              int32 //flag
}

type TL_autoDownloadSettings struct {
	Flags                 int32
	Disabled              bool //flag
	VideoPreloadLarge     bool //flag
	AudioPreloadNext      bool //flag
	PhonecallsLessData    bool //flag
	PhotoSizeMax          int32
	VideoSizeMax          int32
	FileSizeMax           int32
	VideoUploadMaxbitrate int32
}

type TL_account_autoDownloadSettings struct {
	Low    TL // AutoDownloadSettings
	Medium TL // AutoDownloadSettings
	High   TL // AutoDownloadSettings
}

type TL_emojiKeyword struct {
	Keyword   string
	Emoticons []string
}

type TL_emojiKeywordDeleted struct {
	Keyword   string
	Emoticons []string
}

type TL_emojiKeywordsDifference struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Keywords    []TL // EmojiKeyword
}

type TL_emojiURL struct {
	Url string
}

type TL_emojiLanguage struct {
	LangCode string
}

type TL_fileLocationToBeDeprecated struct {
	VolumeID int64
	LocalID  int32
}

type TL_folder struct {
	Flags                     int32
	AutofillNewBroadcasts     bool //flag
	AutofillPublicGroups      bool //flag
	AutofillNewCorrespondents bool //flag
	ID                        int32
	Title                     string
	Photo                     TL // ChatPhoto //flag
}

type TL_inputFolderPeer struct {
	Peer     TL // InputPeer
	FolderID int32
}

type TL_folderPeer struct {
	Peer     TL // Peer
	FolderID int32
}

type TL_messages_searchCounter struct {
	Flags   int32
	Inexact bool //flag
	Filter  TL   // MessagesFilter
	Count   int32
}

type TL_urlAuthResultRequest struct {
	Flags              int32
	RequestWriteAccess bool //flag
	Bot                TL   // User
	Domain             string
}

type TL_urlAuthResultAccepted struct {
	Url string
}

type TL_urlAuthResultDefault struct {
}

type TL_channelLocationEmpty struct {
}

type TL_channelLocation struct {
	GeoPoint TL // GeoPoint
	Address  string
}

type TL_peerLocated struct {
	Peer     TL // Peer
	Expires  int32
	Distance int32
}

type TL_peerSelfLocated struct {
	Expires int32
}

type TL_restrictionReason struct {
	Platform string
	Reason   string
	Text     string
}

type TL_inputTheme struct {
	ID         int64
	AccessHash int64
}

type TL_inputThemeSlug struct {
	Slug string
}

type TL_theme struct {
	Flags         int32
	Creator       bool //flag
	Default       bool //flag
	ID            int64
	AccessHash    int64
	Slug          string
	Title         string
	Document      TL // Document //flag
	Settings      TL // ThemeSettings //flag
	InstallsCount int32
}

type TL_account_themesNotModified struct {
}

type TL_account_themes struct {
	Hash   int32
	Themes []TL // Theme
}

type TL_auth_loginToken struct {
	Expires int32
	Token   []byte
}

type TL_auth_loginTokenMigrateTo struct {
	DcID  int32
	Token []byte
}

type TL_auth_loginTokenSuccess struct {
	Authorization TL // auth_Authorization
}

type TL_account_contentSettings struct {
	Flags              int32
	SensitiveEnabled   bool //flag
	SensitiveCanChange bool //flag
}

type TL_messages_inactiveChats struct {
	Dates []int32
	Chats []TL // Chat
	Users []TL // User
}

type TL_baseThemeClassic struct {
}

type TL_baseThemeDay struct {
}

type TL_baseThemeNight struct {
}

type TL_baseThemeTinted struct {
}

type TL_baseThemeArctic struct {
}

type TL_inputThemeSettings struct {
	Flags              int32
	BaseTheme          TL // BaseTheme
	AccentColor        int32
	MessageTopColor    int32 //flag
	MessageBottomColor int32 //flag
	Wallpaper          TL    // InputWallPaper //flag
	WallpaperSettings  TL    // WallPaperSettings //flag
}

type TL_themeSettings struct {
	Flags              int32
	BaseTheme          TL // BaseTheme
	AccentColor        int32
	MessageTopColor    int32 //flag
	MessageBottomColor int32 //flag
	Wallpaper          TL    // WallPaper //flag
}

type TL_webPageAttributeTheme struct {
	Flags     int32
	Documents []TL // Document //flag
	Settings  TL   // ThemeSettings //flag
}

type TL_messageUserVote struct {
	UserID int32
	Option []byte
	Date   int32
}

type TL_messageUserVoteInputOption struct {
	UserID int32
	Date   int32
}

type TL_messageUserVoteMultiple struct {
	UserID  int32
	Options []TL // bytes
	Date    int32
}

type TL_messages_votesList struct {
	Flags      int32
	Count      int32
	Votes      []TL   // MessageUserVote
	Users      []TL   // User
	NextOffset string //flag
}

type TL_bankCardOpenUrl struct {
	Url  string
	Name string
}

type TL_payments_bankCardData struct {
	Title    string
	OpenUrls []TL // BankCardOpenUrl
}

type TL_dialogFilter struct {
	Flags           int32
	Contacts        bool //flag
	NonContacts     bool //flag
	Groups          bool //flag
	Broadcasts      bool //flag
	Bots            bool //flag
	ExcludeMuted    bool //flag
	ExcludeRead     bool //flag
	ExcludeArchived bool //flag
	ID              int32
	Title           string
	Emoticon        string //flag
	PinnedPeers     []TL   // InputPeer
	IncludePeers    []TL   // InputPeer
	ExcludePeers    []TL   // InputPeer
}

type TL_dialogFilterSuggested struct {
	Filter      TL // DialogFilter
	Description string
}

type TL_statsDateRangeDays struct {
	MinDate int32
	MaxDate int32
}

type TL_statsAbsValueAndPrev struct {
	Current  float64
	Previous float64
}

type TL_statsPercentValue struct {
	Part  float64
	Total float64
}

type TL_statsGraphAsync struct {
	Token string
}

type TL_statsGraphError struct {
	Error string
}

type TL_statsGraph struct {
	Flags     int32
	Json      TL     // DataJSON
	ZoomToken string //flag
}

type TL_messageInteractionCounters struct {
	MsgID    int32
	Views    int32
	Forwards int32
}

type TL_stats_broadcastStats struct {
	Period                    TL   // StatsDateRangeDays
	Followers                 TL   // StatsAbsValueAndPrev
	ViewsPerPost              TL   // StatsAbsValueAndPrev
	SharesPerPost             TL   // StatsAbsValueAndPrev
	EnabledNotifications      TL   // StatsPercentValue
	GrowthGraph               TL   // StatsGraph
	FollowersGraph            TL   // StatsGraph
	MuteGraph                 TL   // StatsGraph
	TopHoursGraph             TL   // StatsGraph
	InteractionsGraph         TL   // StatsGraph
	IvInteractionsGraph       TL   // StatsGraph
	ViewsBySourceGraph        TL   // StatsGraph
	NewFollowersBySourceGraph TL   // StatsGraph
	LanguagesGraph            TL   // StatsGraph
	RecentMessageInteractions []TL // MessageInteractionCounters
}

type TL_help_promoDataEmpty struct {
	Expires int32
}

type TL_help_promoData struct {
	Flags      int32
	Proxy      bool //flag
	Expires    int32
	Peer       TL     // Peer
	Chats      []TL   // Chat
	Users      []TL   // User
	PsaType    string //flag
	PsaMessage string //flag
}

type TL_videoSize struct {
	Type     string
	Location TL // FileLocation
	W        int32
	H        int32
	Size     int32
}

type TL_invokeAfterMsg struct {
	MsgID int64
	Query TL
}

type TL_invokeAfterMsgs struct {
	MsgIds []int64
	Query  TL
}

type TL_initConnection struct {
	Flags          int32
	ApiID          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          TL // InputClientProxy //flag
	Params         TL // JSONValue //flag
	Query          TL
}

type TL_invokeWithLayer struct {
	Layer int32
	Query TL
}

type TL_invokeWithoutUpdates struct {
	Query TL
}

type TL_invokeWithMessagesRange struct {
	Range TL // MessageRange
	Query TL
}

type TL_invokeWithTakeout struct {
	TakeoutID int64
	Query     TL
}

type TL_auth_sendCode struct {
	PhoneNumber string
	ApiID       int32
	ApiHash     string
	Settings    TL // CodeSettings
}

type TL_auth_signUp struct {
	PhoneNumber   string
	PhoneCodeHash string
	FirstName     string
	LastName      string
}

type TL_auth_signIn struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type TL_auth_logOut struct {
}

type TL_auth_resetAuthorizations struct {
}

type TL_auth_exportAuthorization struct {
	DcID int32
}

type TL_auth_importAuthorization struct {
	ID    int32
	Bytes []byte
}

type TL_auth_bindTempAuthKey struct {
	PermAuthKeyID    int64
	Nonce            int64
	ExpiresAt        int32
	EncryptedMessage []byte
}

type TL_auth_importBotAuthorization struct {
	Flags        int32
	ApiID        int32
	ApiHash      string
	BotAuthToken string
}

type TL_auth_checkPassword struct {
	Password TL // InputCheckPasswordSRP
}

type TL_auth_requestPasswordRecovery struct {
}

type TL_auth_recoverPassword struct {
	Code string
}

type TL_auth_resendCode struct {
	PhoneNumber   string
	PhoneCodeHash string
}

type TL_auth_cancelCode struct {
	PhoneNumber   string
	PhoneCodeHash string
}

type TL_auth_dropTempAuthKeys struct {
	ExceptAuthKeys []int64
}

type TL_auth_exportLoginToken struct {
	ApiID     int32
	ApiHash   string
	ExceptIds []int32
}

type TL_auth_importLoginToken struct {
	Token []byte
}

type TL_auth_acceptLoginToken struct {
	Token []byte
}

type TL_account_registerDevice struct {
	Flags      int32
	NoMuted    bool //flag
	TokenType  int32
	Token      string
	AppSandbox TL // Bool
	Secret     []byte
	OtherUids  []int32
}

type TL_account_unregisterDevice struct {
	TokenType int32
	Token     string
	OtherUids []int32
}

type TL_account_updateNotifySettings struct {
	Peer     TL // InputNotifyPeer
	Settings TL // InputPeerNotifySettings
}

type TL_account_getNotifySettings struct {
	Peer TL // InputNotifyPeer
}

type TL_account_resetNotifySettings struct {
}

type TL_account_updateProfile struct {
	Flags     int32
	FirstName string //flag
	LastName  string //flag
	About     string //flag
}

type TL_account_updateStatus struct {
	Offline TL // Bool
}

type TL_account_getWallPapers struct {
	Hash int32
}

type TL_account_reportPeer struct {
	Peer   TL // InputPeer
	Reason TL // ReportReason
}

type TL_account_checkUsername struct {
	Username string
}

type TL_account_updateUsername struct {
	Username string
}

type TL_account_getPrivacy struct {
	Key TL // InputPrivacyKey
}

type TL_account_setPrivacy struct {
	Key   TL   // InputPrivacyKey
	Rules []TL // InputPrivacyRule
}

type TL_account_deleteAccount struct {
	Reason string
}

type TL_account_getAccountTTL struct {
}

type TL_account_setAccountTTL struct {
	Ttl TL // AccountDaysTTL
}

type TL_account_sendChangePhoneCode struct {
	PhoneNumber string
	Settings    TL // CodeSettings
}

type TL_account_changePhone struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type TL_account_updateDeviceLocked struct {
	Period int32
}

type TL_account_getAuthorizations struct {
}

type TL_account_resetAuthorization struct {
	Hash int64
}

type TL_account_getPassword struct {
}

type TL_account_getPasswordSettings struct {
	Password TL // InputCheckPasswordSRP
}

type TL_account_updatePasswordSettings struct {
	Password    TL // InputCheckPasswordSRP
	NewSettings TL // account_PasswordInputSettings
}

type TL_account_sendConfirmPhoneCode struct {
	Hash     string
	Settings TL // CodeSettings
}

type TL_account_confirmPhone struct {
	PhoneCodeHash string
	PhoneCode     string
}

type TL_account_getTmpPassword struct {
	Password TL // InputCheckPasswordSRP
	Period   int32
}

type TL_account_getWebAuthorizations struct {
}

type TL_account_resetWebAuthorization struct {
	Hash int64
}

type TL_account_resetWebAuthorizations struct {
}

type TL_account_getAllSecureValues struct {
}

type TL_account_getSecureValue struct {
	Types []TL // SecureValueType
}

type TL_account_saveSecureValue struct {
	Value          TL // InputSecureValue
	SecureSecretID int64
}

type TL_account_deleteSecureValue struct {
	Types []TL // SecureValueType
}

type TL_account_getAuthorizationForm struct {
	BotID     int32
	Scope     string
	PublicKey string
}

type TL_account_acceptAuthorization struct {
	BotID       int32
	Scope       string
	PublicKey   string
	ValueHashes []TL // SecureValueHash
	Credentials TL   // SecureCredentialsEncrypted
}

type TL_account_sendVerifyPhoneCode struct {
	PhoneNumber string
	Settings    TL // CodeSettings
}

type TL_account_verifyPhone struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type TL_account_sendVerifyEmailCode struct {
	Email string
}

type TL_account_verifyEmail struct {
	Email string
	Code  string
}

type TL_account_initTakeoutSession struct {
	Flags             int32
	Contacts          bool  //flag
	MessageUsers      bool  //flag
	MessageChats      bool  //flag
	MessageMegagroups bool  //flag
	MessageChannels   bool  //flag
	Files             bool  //flag
	FileMaxSize       int32 //flag
}

type TL_account_finishTakeoutSession struct {
	Flags   int32
	Success bool //flag
}

type TL_account_confirmPasswordEmail struct {
	Code string
}

type TL_account_resendPasswordEmail struct {
}

type TL_account_cancelPasswordEmail struct {
}

type TL_account_getContactSignUpNotification struct {
}

type TL_account_setContactSignUpNotification struct {
	Silent TL // Bool
}

type TL_account_getNotifyExceptions struct {
	Flags        int32
	CompareSound bool //flag
	Peer         TL   // InputNotifyPeer //flag
}

type TL_account_getWallPaper struct {
	Wallpaper TL // InputWallPaper
}

type TL_account_uploadWallPaper struct {
	File     TL // InputFile
	MimeType string
	Settings TL // WallPaperSettings
}

type TL_account_saveWallPaper struct {
	Wallpaper TL // InputWallPaper
	Unsave    TL // Bool
	Settings  TL // WallPaperSettings
}

type TL_account_installWallPaper struct {
	Wallpaper TL // InputWallPaper
	Settings  TL // WallPaperSettings
}

type TL_account_resetWallPapers struct {
}

type TL_account_getAutoDownloadSettings struct {
}

type TL_account_saveAutoDownloadSettings struct {
	Flags    int32
	Low      bool //flag
	High     bool //flag
	Settings TL   // AutoDownloadSettings
}

type TL_account_uploadTheme struct {
	Flags    int32
	File     TL // InputFile
	Thumb    TL // InputFile //flag
	FileName string
	MimeType string
}

type TL_account_createTheme struct {
	Flags    int32
	Slug     string
	Title    string
	Document TL // InputDocument //flag
	Settings TL // InputThemeSettings //flag
}

type TL_account_updateTheme struct {
	Flags    int32
	Format   string
	Theme    TL     // InputTheme
	Slug     string //flag
	Title    string //flag
	Document TL     // InputDocument //flag
	Settings TL     // InputThemeSettings //flag
}

type TL_account_saveTheme struct {
	Theme  TL // InputTheme
	Unsave TL // Bool
}

type TL_account_installTheme struct {
	Flags  int32
	Dark   bool   //flag
	Format string //flag
	Theme  TL     // InputTheme //flag
}

type TL_account_getTheme struct {
	Format     string
	Theme      TL // InputTheme
	DocumentID int64
}

type TL_account_getThemes struct {
	Format string
	Hash   int32
}

type TL_account_setContentSettings struct {
	Flags            int32
	SensitiveEnabled bool //flag
}

type TL_account_getContentSettings struct {
}

type TL_account_getMultiWallPapers struct {
	Wallpapers []TL // InputWallPaper
}

type TL_users_getUsers struct {
	ID []TL // InputUser
}

type TL_users_getFullUser struct {
	ID TL // InputUser
}

type TL_users_setSecureValueErrors struct {
	ID     TL   // InputUser
	Errors []TL // SecureValueError
}

type TL_contacts_getContactIDs struct {
	Hash int32
}

type TL_contacts_getStatuses struct {
}

type TL_contacts_getContacts struct {
	Hash int32
}

type TL_contacts_importContacts struct {
	Contacts []TL // InputContact
}

type TL_contacts_deleteContacts struct {
	ID []TL // InputUser
}

type TL_contacts_deleteByPhones struct {
	Phones []string
}

type TL_contacts_block struct {
	ID TL // InputUser
}

type TL_contacts_unblock struct {
	ID TL // InputUser
}

type TL_contacts_getBlocked struct {
	Offset int32
	Limit  int32
}

type TL_contacts_search struct {
	Q     string
	Limit int32
}

type TL_contacts_resolveUsername struct {
	Username string
}

type TL_contacts_getTopPeers struct {
	Flags          int32
	Correspondents bool //flag
	BotsPm         bool //flag
	BotsInline     bool //flag
	PhoneCalls     bool //flag
	ForwardUsers   bool //flag
	ForwardChats   bool //flag
	Groups         bool //flag
	Channels       bool //flag
	Offset         int32
	Limit          int32
	Hash           int32
}

type TL_contacts_resetTopPeerRating struct {
	Category TL // TopPeerCategory
	Peer     TL // InputPeer
}

type TL_contacts_resetSaved struct {
}

type TL_contacts_getSaved struct {
}

type TL_contacts_toggleTopPeers struct {
	Enabled TL // Bool
}

type TL_contacts_addContact struct {
	Flags                    int32
	AddPhonePrivacyException bool //flag
	ID                       TL   // InputUser
	FirstName                string
	LastName                 string
	Phone                    string
}

type TL_contacts_acceptContact struct {
	ID TL // InputUser
}

type TL_contacts_getLocated struct {
	Flags       int32
	Background  bool  //flag
	GeoPoint    TL    // InputGeoPoint
	SelfExpires int32 //flag
}

type TL_messages_getMessages struct {
	ID []TL // InputMessage
}

type TL_messages_getDialogs struct {
	Flags         int32
	ExcludePinned bool  //flag
	FolderID      int32 //flag
	OffsetDate    int32
	OffsetID      int32
	OffsetPeer    TL // InputPeer
	Limit         int32
	Hash          int32
}

type TL_messages_getHistory struct {
	Peer       TL // InputPeer
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int32
}

type TL_messages_search struct {
	Flags     int32
	Peer      TL // InputPeer
	Q         string
	FromID    TL // InputUser //flag
	Filter    TL // MessagesFilter
	MinDate   int32
	MaxDate   int32
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
	Hash      int32
}

type TL_messages_readHistory struct {
	Peer  TL // InputPeer
	MaxID int32
}

type TL_messages_deleteHistory struct {
	Flags     int32
	JustClear bool //flag
	Revoke    bool //flag
	Peer      TL   // InputPeer
	MaxID     int32
}

type TL_messages_deleteMessages struct {
	Flags  int32
	Revoke bool //flag
	ID     []int32
}

type TL_messages_receivedMessages struct {
	MaxID int32
}

type TL_messages_setTyping struct {
	Peer   TL // InputPeer
	Action TL // SendMessageAction
}

type TL_messages_sendMessage struct {
	Flags        int32
	NoWebpage    bool  //flag
	Silent       bool  //flag
	Background   bool  //flag
	ClearDraft   bool  //flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 //flag
	Message      string
	RandomID     int64
	ReplyMarkup  TL    // ReplyMarkup //flag
	Entities     []TL  // MessageEntity //flag
	ScheduleDate int32 //flag
}

type TL_messages_sendMedia struct {
	Flags        int32
	Silent       bool  //flag
	Background   bool  //flag
	ClearDraft   bool  //flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 //flag
	Media        TL    // InputMedia
	Message      string
	RandomID     int64
	ReplyMarkup  TL    // ReplyMarkup //flag
	Entities     []TL  // MessageEntity //flag
	ScheduleDate int32 //flag
}

type TL_messages_forwardMessages struct {
	Flags        int32
	Silent       bool //flag
	Background   bool //flag
	WithMyScore  bool //flag
	Grouped      bool //flag
	FromPeer     TL   // InputPeer
	ID           []int32
	RandomID     []int64
	ToPeer       TL    // InputPeer
	ScheduleDate int32 //flag
}

type TL_messages_reportSpam struct {
	Peer TL // InputPeer
}

type TL_messages_getPeerSettings struct {
	Peer TL // InputPeer
}

type TL_messages_report struct {
	Peer   TL // InputPeer
	ID     []int32
	Reason TL // ReportReason
}

type TL_messages_getChats struct {
	ID []int32
}

type TL_messages_getFullChat struct {
	ChatID int32
}

type TL_messages_editChatTitle struct {
	ChatID int32
	Title  string
}

type TL_messages_editChatPhoto struct {
	ChatID int32
	Photo  TL // InputChatPhoto
}

type TL_messages_addChatUser struct {
	ChatID   int32
	UserID   TL // InputUser
	FwdLimit int32
}

type TL_messages_deleteChatUser struct {
	ChatID int32
	UserID TL // InputUser
}

type TL_messages_createChat struct {
	Users []TL // InputUser
	Title string
}

type TL_messages_getDhConfig struct {
	Version      int32
	RandomLength int32
}

type TL_messages_requestEncryption struct {
	UserID   TL // InputUser
	RandomID int32
	GA       []byte
}

type TL_messages_acceptEncryption struct {
	Peer           TL // InputEncryptedChat
	GB             []byte
	KeyFingerprint int64
}

type TL_messages_discardEncryption struct {
	ChatID int32
}

type TL_messages_setEncryptedTyping struct {
	Peer   TL // InputEncryptedChat
	Typing TL // Bool
}

type TL_messages_readEncryptedHistory struct {
	Peer    TL // InputEncryptedChat
	MaxDate int32
}

type TL_messages_sendEncrypted struct {
	Peer     TL // InputEncryptedChat
	RandomID int64
	Data     []byte
}

type TL_messages_sendEncryptedFile struct {
	Peer     TL // InputEncryptedChat
	RandomID int64
	Data     []byte
	File     TL // InputEncryptedFile
}

type TL_messages_sendEncryptedService struct {
	Peer     TL // InputEncryptedChat
	RandomID int64
	Data     []byte
}

type TL_messages_receivedQueue struct {
	MaxQts int32
}

type TL_messages_reportEncryptedSpam struct {
	Peer TL // InputEncryptedChat
}

type TL_messages_readMessageContents struct {
	ID []int32
}

type TL_messages_getStickers struct {
	Emoticon string
	Hash     int32
}

type TL_messages_getAllStickers struct {
	Hash int32
}

type TL_messages_getWebPagePreview struct {
	Flags    int32
	Message  string
	Entities []TL // MessageEntity //flag
}

type TL_messages_exportChatInvite struct {
	Peer TL // InputPeer
}

type TL_messages_checkChatInvite struct {
	Hash string
}

type TL_messages_importChatInvite struct {
	Hash string
}

type TL_messages_getStickerSet struct {
	Stickerset TL // InputStickerSet
}

type TL_messages_installStickerSet struct {
	Stickerset TL // InputStickerSet
	Archived   TL // Bool
}

type TL_messages_uninstallStickerSet struct {
	Stickerset TL // InputStickerSet
}

type TL_messages_startBot struct {
	Bot        TL // InputUser
	Peer       TL // InputPeer
	RandomID   int64
	StartParam string
}

type TL_messages_getMessagesViews struct {
	Peer      TL // InputPeer
	ID        []int32
	Increment TL // Bool
}

type TL_messages_editChatAdmin struct {
	ChatID  int32
	UserID  TL // InputUser
	IsAdmin TL // Bool
}

type TL_messages_migrateChat struct {
	ChatID int32
}

type TL_messages_searchGlobal struct {
	Flags      int32
	FolderID   int32 //flag
	Q          string
	OffsetRate int32
	OffsetPeer TL // InputPeer
	OffsetID   int32
	Limit      int32
}

type TL_messages_reorderStickerSets struct {
	Flags int32
	Masks bool //flag
	Order []int64
}

type TL_messages_getDocumentByHash struct {
	Sha256   []byte
	Size     int32
	MimeType string
}

type TL_messages_searchGifs struct {
	Q      string
	Offset int32
}

type TL_messages_getSavedGifs struct {
	Hash int32
}

type TL_messages_saveGif struct {
	ID     TL // InputDocument
	Unsave TL // Bool
}

type TL_messages_getInlineBotResults struct {
	Flags    int32
	Bot      TL // InputUser
	Peer     TL // InputPeer
	GeoPoint TL // InputGeoPoint //flag
	Query    string
	Offset   string
}

type TL_messages_setInlineBotResults struct {
	Flags      int32
	Gallery    bool //flag
	Private    bool //flag
	QueryID    int64
	Results    []TL // InputBotInlineResult
	CacheTime  int32
	NextOffset string //flag
	SwitchPm   TL     // InlineBotSwitchPM //flag
}

type TL_messages_sendInlineBotResult struct {
	Flags        int32
	Silent       bool  //flag
	Background   bool  //flag
	ClearDraft   bool  //flag
	HideVia      bool  //flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 //flag
	RandomID     int64
	QueryID      int64
	ID           string
	ScheduleDate int32 //flag
}

type TL_messages_getMessageEditData struct {
	Peer TL // InputPeer
	ID   int32
}

type TL_messages_editMessage struct {
	Flags        int32
	NoWebpage    bool //flag
	Peer         TL   // InputPeer
	ID           int32
	Message      string //flag
	Media        TL     // InputMedia //flag
	ReplyMarkup  TL     // ReplyMarkup //flag
	Entities     []TL   // MessageEntity //flag
	ScheduleDate int32  //flag
}

type TL_messages_editInlineBotMessage struct {
	Flags       int32
	NoWebpage   bool   //flag
	ID          TL     // InputBotInlineMessageID
	Message     string //flag
	Media       TL     // InputMedia //flag
	ReplyMarkup TL     // ReplyMarkup //flag
	Entities    []TL   // MessageEntity //flag
}

type TL_messages_getBotCallbackAnswer struct {
	Flags int32
	Game  bool //flag
	Peer  TL   // InputPeer
	MsgID int32
	Data  []byte //flag
}

type TL_messages_setBotCallbackAnswer struct {
	Flags     int32
	Alert     bool //flag
	QueryID   int64
	Message   string //flag
	Url       string //flag
	CacheTime int32
}

type TL_messages_getPeerDialogs struct {
	Peers []TL // InputDialogPeer
}

type TL_messages_saveDraft struct {
	Flags        int32
	NoWebpage    bool  //flag
	ReplyToMsgID int32 //flag
	Peer         TL    // InputPeer
	Message      string
	Entities     []TL // MessageEntity //flag
}

type TL_messages_getAllDrafts struct {
}

type TL_messages_getFeaturedStickers struct {
	Hash int32
}

type TL_messages_readFeaturedStickers struct {
	ID []int64
}

type TL_messages_getRecentStickers struct {
	Flags    int32
	Attached bool //flag
	Hash     int32
}

type TL_messages_saveRecentSticker struct {
	Flags    int32
	Attached bool //flag
	ID       TL   // InputDocument
	Unsave   TL   // Bool
}

type TL_messages_clearRecentStickers struct {
	Flags    int32
	Attached bool //flag
}

type TL_messages_getArchivedStickers struct {
	Flags    int32
	Masks    bool //flag
	OffsetID int64
	Limit    int32
}

type TL_messages_getMaskStickers struct {
	Hash int32
}

type TL_messages_getAttachedStickers struct {
	Media TL // InputStickeredMedia
}

type TL_messages_setGameScore struct {
	Flags       int32
	EditMessage bool //flag
	Force       bool //flag
	Peer        TL   // InputPeer
	ID          int32
	UserID      TL // InputUser
	Score       int32
}

type TL_messages_setInlineGameScore struct {
	Flags       int32
	EditMessage bool //flag
	Force       bool //flag
	ID          TL   // InputBotInlineMessageID
	UserID      TL   // InputUser
	Score       int32
}

type TL_messages_getGameHighScores struct {
	Peer   TL // InputPeer
	ID     int32
	UserID TL // InputUser
}

type TL_messages_getInlineGameHighScores struct {
	ID     TL // InputBotInlineMessageID
	UserID TL // InputUser
}

type TL_messages_getCommonChats struct {
	UserID TL // InputUser
	MaxID  int32
	Limit  int32
}

type TL_messages_getAllChats struct {
	ExceptIds []int32
}

type TL_messages_getWebPage struct {
	Url  string
	Hash int32
}

type TL_messages_toggleDialogPin struct {
	Flags  int32
	Pinned bool //flag
	Peer   TL   // InputDialogPeer
}

type TL_messages_reorderPinnedDialogs struct {
	Flags    int32
	Force    bool //flag
	FolderID int32
	Order    []TL // InputDialogPeer
}

type TL_messages_getPinnedDialogs struct {
	FolderID int32
}

type TL_messages_setBotShippingResults struct {
	Flags           int32
	QueryID         int64
	Error           string //flag
	ShippingOptions []TL   // ShippingOption //flag
}

type TL_messages_setBotPrecheckoutResults struct {
	Flags   int32
	Success bool //flag
	QueryID int64
	Error   string //flag
}

type TL_messages_uploadMedia struct {
	Peer  TL // InputPeer
	Media TL // InputMedia
}

type TL_messages_sendScreenshotNotification struct {
	Peer         TL // InputPeer
	ReplyToMsgID int32
	RandomID     int64
}

type TL_messages_getFavedStickers struct {
	Hash int32
}

type TL_messages_faveSticker struct {
	ID     TL // InputDocument
	Unfave TL // Bool
}

type TL_messages_getUnreadMentions struct {
	Peer      TL // InputPeer
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

type TL_messages_readMentions struct {
	Peer TL // InputPeer
}

type TL_messages_getRecentLocations struct {
	Peer  TL // InputPeer
	Limit int32
	Hash  int32
}

type TL_messages_sendMultiMedia struct {
	Flags        int32
	Silent       bool  //flag
	Background   bool  //flag
	ClearDraft   bool  //flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 //flag
	MultiMedia   []TL  // InputSingleMedia
	ScheduleDate int32 //flag
}

type TL_messages_uploadEncryptedFile struct {
	Peer TL // InputEncryptedChat
	File TL // InputEncryptedFile
}

type TL_messages_searchStickerSets struct {
	Flags           int32
	ExcludeFeatured bool //flag
	Q               string
	Hash            int32
}

type TL_messages_getSplitRanges struct {
}

type TL_messages_markDialogUnread struct {
	Flags  int32
	Unread bool //flag
	Peer   TL   // InputDialogPeer
}

type TL_messages_getDialogUnreadMarks struct {
}

type TL_messages_clearAllDrafts struct {
}

type TL_messages_updatePinnedMessage struct {
	Flags  int32
	Silent bool //flag
	Peer   TL   // InputPeer
	ID     int32
}

type TL_messages_sendVote struct {
	Peer    TL // InputPeer
	MsgID   int32
	Options []TL // bytes
}

type TL_messages_getPollResults struct {
	Peer  TL // InputPeer
	MsgID int32
}

type TL_messages_getOnlines struct {
	Peer TL // InputPeer
}

type TL_messages_getStatsURL struct {
	Flags  int32
	Dark   bool //flag
	Peer   TL   // InputPeer
	Params string
}

type TL_messages_editChatAbout struct {
	Peer  TL // InputPeer
	About string
}

type TL_messages_editChatDefaultBannedRights struct {
	Peer         TL // InputPeer
	BannedRights TL // ChatBannedRights
}

type TL_messages_getEmojiKeywords struct {
	LangCode string
}

type TL_messages_getEmojiKeywordsDifference struct {
	LangCode    string
	FromVersion int32
}

type TL_messages_getEmojiKeywordsLanguages struct {
	LangCodes []string
}

type TL_messages_getEmojiURL struct {
	LangCode string
}

type TL_messages_getSearchCounters struct {
	Peer    TL   // InputPeer
	Filters []TL // MessagesFilter
}

type TL_messages_requestUrlAuth struct {
	Peer     TL // InputPeer
	MsgID    int32
	ButtonID int32
}

type TL_messages_acceptUrlAuth struct {
	Flags        int32
	WriteAllowed bool //flag
	Peer         TL   // InputPeer
	MsgID        int32
	ButtonID     int32
}

type TL_messages_hidePeerSettingsBar struct {
	Peer TL // InputPeer
}

type TL_messages_getScheduledHistory struct {
	Peer TL // InputPeer
	Hash int32
}

type TL_messages_getScheduledMessages struct {
	Peer TL // InputPeer
	ID   []int32
}

type TL_messages_sendScheduledMessages struct {
	Peer TL // InputPeer
	ID   []int32
}

type TL_messages_deleteScheduledMessages struct {
	Peer TL // InputPeer
	ID   []int32
}

type TL_messages_getPollVotes struct {
	Flags  int32
	Peer   TL // InputPeer
	ID     int32
	Option []byte //flag
	Offset string //flag
	Limit  int32
}

type TL_messages_toggleStickerSets struct {
	Flags       int32
	Uninstall   bool //flag
	Archive     bool //flag
	Unarchive   bool //flag
	Stickersets []TL // InputStickerSet
}

type TL_messages_getDialogFilters struct {
}

type TL_messages_getSuggestedDialogFilters struct {
}

type TL_messages_updateDialogFilter struct {
	Flags  int32
	ID     int32
	Filter TL // DialogFilter //flag
}

type TL_messages_updateDialogFiltersOrder struct {
	Order []int32
}

type TL_messages_getOldFeaturedStickers struct {
	Offset int32
	Limit  int32
	Hash   int32
}

type TL_updates_getState struct {
}

type TL_updates_getDifference struct {
	Flags         int32
	Pts           int32
	PtsTotalLimit int32 //flag
	Date          int32
	Qts           int32
}

type TL_updates_getChannelDifference struct {
	Flags   int32
	Force   bool //flag
	Channel TL   // InputChannel
	Filter  TL   // ChannelMessagesFilter
	Pts     int32
	Limit   int32
}

type TL_photos_updateProfilePhoto struct {
	ID TL // InputPhoto
}

type TL_photos_uploadProfilePhoto struct {
	File TL // InputFile
}

type TL_photos_deletePhotos struct {
	ID []TL // InputPhoto
}

type TL_photos_getUserPhotos struct {
	UserID TL // InputUser
	Offset int32
	MaxID  int64
	Limit  int32
}

type TL_upload_saveFilePart struct {
	FileID   int64
	FilePart int32
	Bytes    []byte
}

type TL_upload_getFile struct {
	Flags        int32
	Precise      bool //flag
	CdnSupported bool //flag
	Location     TL   // InputFileLocation
	Offset       int32
	Limit        int32
}

type TL_upload_saveBigFilePart struct {
	FileID         int64
	FilePart       int32
	FileTotalParts int32
	Bytes          []byte
}

type TL_upload_getWebFile struct {
	Location TL // InputWebFileLocation
	Offset   int32
	Limit    int32
}

type TL_upload_getCdnFile struct {
	FileToken []byte
	Offset    int32
	Limit     int32
}

type TL_upload_reuploadCdnFile struct {
	FileToken    []byte
	RequestToken []byte
}

type TL_upload_getCdnFileHashes struct {
	FileToken []byte
	Offset    int32
}

type TL_upload_getFileHashes struct {
	Location TL // InputFileLocation
	Offset   int32
}

type TL_help_getConfig struct {
}

type TL_help_getNearestDc struct {
}

type TL_help_getAppUpdate struct {
	Source string
}

type TL_help_getInviteText struct {
}

type TL_help_getSupport struct {
}

type TL_help_getAppChangelog struct {
	PrevAppVersion string
}

type TL_help_setBotUpdatesStatus struct {
	PendingUpdatesCount int32
	Message             string
}

type TL_help_getCdnConfig struct {
}

type TL_help_getRecentMeUrls struct {
	Referer string
}

type TL_help_getTermsOfServiceUpdate struct {
}

type TL_help_acceptTermsOfService struct {
	ID TL // DataJSON
}

type TL_help_getDeepLinkInfo struct {
	Path string
}

type TL_help_getAppConfig struct {
}

type TL_help_saveAppLog struct {
	Events []TL // InputAppEvent
}

type TL_help_getPassportConfig struct {
	Hash int32
}

type TL_help_getSupportName struct {
}

type TL_help_getUserInfo struct {
	UserID TL // InputUser
}

type TL_help_editUserInfo struct {
	UserID   TL // InputUser
	Message  string
	Entities []TL // MessageEntity
}

type TL_help_getPromoData struct {
}

type TL_help_hidePromoData struct {
	Peer TL // InputPeer
}

type TL_channels_readHistory struct {
	Channel TL // InputChannel
	MaxID   int32
}

type TL_channels_deleteMessages struct {
	Channel TL // InputChannel
	ID      []int32
}

type TL_channels_deleteUserHistory struct {
	Channel TL // InputChannel
	UserID  TL // InputUser
}

type TL_channels_reportSpam struct {
	Channel TL // InputChannel
	UserID  TL // InputUser
	ID      []int32
}

type TL_channels_getMessages struct {
	Channel TL   // InputChannel
	ID      []TL // InputMessage
}

type TL_channels_getParticipants struct {
	Channel TL // InputChannel
	Filter  TL // ChannelParticipantsFilter
	Offset  int32
	Limit   int32
	Hash    int32
}

type TL_channels_getParticipant struct {
	Channel TL // InputChannel
	UserID  TL // InputUser
}

type TL_channels_getChannels struct {
	ID []TL // InputChannel
}

type TL_channels_getFullChannel struct {
	Channel TL // InputChannel
}

type TL_channels_createChannel struct {
	Flags     int32
	Broadcast bool //flag
	Megagroup bool //flag
	Title     string
	About     string
	GeoPoint  TL     // InputGeoPoint //flag
	Address   string //flag
}

type TL_channels_editAdmin struct {
	Channel     TL // InputChannel
	UserID      TL // InputUser
	AdminRights TL // ChatAdminRights
	Rank        string
}

type TL_channels_editTitle struct {
	Channel TL // InputChannel
	Title   string
}

type TL_channels_editPhoto struct {
	Channel TL // InputChannel
	Photo   TL // InputChatPhoto
}

type TL_channels_checkUsername struct {
	Channel  TL // InputChannel
	Username string
}

type TL_channels_updateUsername struct {
	Channel  TL // InputChannel
	Username string
}

type TL_channels_joinChannel struct {
	Channel TL // InputChannel
}

type TL_channels_leaveChannel struct {
	Channel TL // InputChannel
}

type TL_channels_inviteToChannel struct {
	Channel TL   // InputChannel
	Users   []TL // InputUser
}

type TL_channels_deleteChannel struct {
	Channel TL // InputChannel
}

type TL_channels_exportMessageLink struct {
	Channel TL // InputChannel
	ID      int32
	Grouped TL // Bool
}

type TL_channels_toggleSignatures struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type TL_channels_getAdminedPublicChannels struct {
	Flags      int32
	ByLocation bool //flag
	CheckLimit bool //flag
}

type TL_channels_editBanned struct {
	Channel      TL // InputChannel
	UserID       TL // InputUser
	BannedRights TL // ChatBannedRights
}

type TL_channels_getAdminLog struct {
	Flags        int32
	Channel      TL // InputChannel
	Q            string
	EventsFilter TL   // ChannelAdminLogEventsFilter //flag
	Admins       []TL // InputUser //flag
	MaxID        int64
	MinID        int64
	Limit        int32
}

type TL_channels_setStickers struct {
	Channel    TL // InputChannel
	Stickerset TL // InputStickerSet
}

type TL_channels_readMessageContents struct {
	Channel TL // InputChannel
	ID      []int32
}

type TL_channels_deleteHistory struct {
	Channel TL // InputChannel
	MaxID   int32
}

type TL_channels_togglePreHistoryHidden struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type TL_channels_getLeftChannels struct {
	Offset int32
}

type TL_channels_getGroupsForDiscussion struct {
}

type TL_channels_setDiscussionGroup struct {
	Broadcast TL // InputChannel
	Group     TL // InputChannel
}

type TL_channels_editCreator struct {
	Channel  TL // InputChannel
	UserID   TL // InputUser
	Password TL // InputCheckPasswordSRP
}

type TL_channels_editLocation struct {
	Channel  TL // InputChannel
	GeoPoint TL // InputGeoPoint
	Address  string
}

type TL_channels_toggleSlowMode struct {
	Channel TL // InputChannel
	Seconds int32
}

type TL_channels_getInactiveChannels struct {
}

type TL_bots_sendCustomRequest struct {
	CustomMethod string
	Params       TL // DataJSON
}

type TL_bots_answerWebhookJSONQuery struct {
	QueryID int64
	Data    TL // DataJSON
}

type TL_bots_setBotCommands struct {
	Commands []TL // BotCommand
}

type TL_payments_getPaymentForm struct {
	MsgID int32
}

type TL_payments_getPaymentReceipt struct {
	MsgID int32
}

type TL_payments_validateRequestedInfo struct {
	Flags int32
	Save  bool //flag
	MsgID int32
	Info  TL // PaymentRequestedInfo
}

type TL_payments_sendPaymentForm struct {
	Flags            int32
	MsgID            int32
	RequestedInfoID  string //flag
	ShippingOptionID string //flag
	Credentials      TL     // InputPaymentCredentials
}

type TL_payments_getSavedInfo struct {
}

type TL_payments_clearSavedInfo struct {
	Flags       int32
	Credentials bool //flag
	Info        bool //flag
}

type TL_payments_getBankCardData struct {
	Number string
}

type TL_stickers_createStickerSet struct {
	Flags     int32
	Masks     bool //flag
	Animated  bool //flag
	UserID    TL   // InputUser
	Title     string
	ShortName string
	Thumb     TL   // InputDocument //flag
	Stickers  []TL // InputStickerSetItem
}

type TL_stickers_removeStickerFromSet struct {
	Sticker TL // InputDocument
}

type TL_stickers_changeStickerPosition struct {
	Sticker  TL // InputDocument
	Position int32
}

type TL_stickers_addStickerToSet struct {
	Stickerset TL // InputStickerSet
	Sticker    TL // InputStickerSetItem
}

type TL_stickers_setStickerSetThumb struct {
	Stickerset TL // InputStickerSet
	Thumb      TL // InputDocument
}

type TL_phone_getCallConfig struct {
}

type TL_phone_requestCall struct {
	Flags    int32
	Video    bool //flag
	UserID   TL   // InputUser
	RandomID int32
	GAHash   []byte
	Protocol TL // PhoneCallProtocol
}

type TL_phone_acceptCall struct {
	Peer     TL // InputPhoneCall
	GB       []byte
	Protocol TL // PhoneCallProtocol
}

type TL_phone_confirmCall struct {
	Peer           TL // InputPhoneCall
	GA             []byte
	KeyFingerprint int64
	Protocol       TL // PhoneCallProtocol
}

type TL_phone_receivedCall struct {
	Peer TL // InputPhoneCall
}

type TL_phone_discardCall struct {
	Flags        int32
	Video        bool //flag
	Peer         TL   // InputPhoneCall
	Duration     int32
	Reason       TL // PhoneCallDiscardReason
	ConnectionID int64
}

type TL_phone_setCallRating struct {
	Flags          int32
	UserInitiative bool //flag
	Peer           TL   // InputPhoneCall
	Rating         int32
	Comment        string
}

type TL_phone_saveCallDebug struct {
	Peer  TL // InputPhoneCall
	Debug TL // DataJSON
}

type TL_phone_sendSignalingData struct {
	Peer TL // InputPhoneCall
	Data []byte
}

type TL_langpack_getLangPack struct {
	LangPack string
	LangCode string
}

type TL_langpack_getStrings struct {
	LangPack string
	LangCode string
	Keys     []string
}

type TL_langpack_getDifference struct {
	LangPack    string
	LangCode    string
	FromVersion int32
}

type TL_langpack_getLanguages struct {
	LangPack string
}

type TL_langpack_getLanguage struct {
	LangPack string
	LangCode string
}

type TL_folders_editPeerFolders struct {
	FolderPeers []TL // InputFolderPeer
}

type TL_folders_deleteFolder struct {
	FolderID int32
}

type TL_stats_getBroadcastStats struct {
	Flags   int32
	Dark    bool //flag
	Channel TL   // InputChannel
}

type TL_stats_loadAsyncGraph struct {
	Flags int32
	Token string
	X     int64 //flag
}

func (e TL_resPQ) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_resPQ)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.Pq)
	x.VectorLong(e.ServerPublicKeyFingerprints)
	return x.buf
}

func (e TL_p_q_inner_data) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_p_q_inner_data)
	x.String(e.Pq)
	x.String(e.P)
	x.String(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonce)
	return x.buf
}

func (e TL_p_q_inner_data_dc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_p_q_inner_data_dc)
	x.String(e.Pq)
	x.String(e.P)
	x.String(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonce)
	x.Int(e.Dc)
	return x.buf
}

func (e TL_p_q_inner_data_temp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_p_q_inner_data_temp)
	x.String(e.Pq)
	x.String(e.P)
	x.String(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonce)
	x.Int(e.ExpiresIn)
	return x.buf
}

func (e TL_p_q_inner_data_temp_dc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_p_q_inner_data_temp_dc)
	x.String(e.Pq)
	x.String(e.P)
	x.String(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonce)
	x.Int(e.Dc)
	x.Int(e.ExpiresIn)
	return x.buf
}

func (e TL_bind_auth_key_inner) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bind_auth_key_inner)
	x.Long(e.Nonce)
	x.Long(e.TempAuthKeyID)
	x.Long(e.PermAuthKeyID)
	x.Long(e.TempSessionID)
	x.Int(e.ExpiresAt)
	return x.buf
}

func (e TL_server_DH_params_fail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_server_DH_params_fail)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash)
	return x.buf
}

func (e TL_server_DH_params_ok) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_server_DH_params_ok)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.EncryptedAnswer)
	return x.buf
}

func (e TL_server_DH_inner_data) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_server_DH_inner_data)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Int(e.G)
	x.String(e.DhPrime)
	x.String(e.GA)
	x.Int(e.ServerTime)
	return x.buf
}

func (e TL_client_DH_inner_data) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_client_DH_inner_data)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Long(e.RetryID)
	x.String(e.GB)
	return x.buf
}

func (e TL_dh_gen_ok) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dh_gen_ok)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash1)
	return x.buf
}

func (e TL_dh_gen_retry) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dh_gen_retry)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash2)
	return x.buf
}

func (e TL_dh_gen_fail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dh_gen_fail)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash3)
	return x.buf
}

func (e TL_destroy_auth_key_ok) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_auth_key_ok)
	return x.buf
}

func (e TL_destroy_auth_key_none) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_auth_key_none)
	return x.buf
}

func (e TL_destroy_auth_key_fail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_auth_key_fail)
	return x.buf
}

func (e TL_req_pq) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_req_pq)
	x.Bytes(e.Nonce)
	return x.buf
}

func (e TL_req_pq_multi) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_req_pq_multi)
	x.Bytes(e.Nonce)
	return x.buf
}

func (e TL_req_DH_params) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_req_DH_params)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.P)
	x.String(e.Q)
	x.Long(e.PublicKeyFingerprint)
	x.String(e.EncryptedData)
	return x.buf
}

func (e TL_set_client_DH_params) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_set_client_DH_params)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.EncryptedData)
	return x.buf
}

func (e TL_destroy_auth_key) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_auth_key)
	return x.buf
}

func (e TL_msgs_ack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msgs_ack)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e TL_bad_msg_notification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bad_msg_notification)
	x.Long(e.BadMsgID)
	x.Int(e.BadMsgSeqno)
	x.Int(e.ErrorCode)
	return x.buf
}

func (e TL_bad_server_salt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bad_server_salt)
	x.Long(e.BadMsgID)
	x.Int(e.BadMsgSeqno)
	x.Int(e.ErrorCode)
	x.Long(e.NewServerSalt)
	return x.buf
}

func (e TL_msgs_state_req) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msgs_state_req)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e TL_msgs_state_info) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msgs_state_info)
	x.Long(e.ReqMsgID)
	x.String(e.Info)
	return x.buf
}

func (e TL_msgs_all_info) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msgs_all_info)
	x.VectorLong(e.MsgIds)
	x.String(e.Info)
	return x.buf
}

func (e TL_msg_detailed_info) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msg_detailed_info)
	x.Long(e.MsgID)
	x.Long(e.AnswerMsgID)
	x.Int(e.Bytes)
	x.Int(e.Status)
	return x.buf
}

func (e TL_msg_new_detailed_info) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msg_new_detailed_info)
	x.Long(e.AnswerMsgID)
	x.Int(e.Bytes)
	x.Int(e.Status)
	return x.buf
}

func (e TL_msg_resend_req) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_msg_resend_req)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e TL_rpc_error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_rpc_error)
	x.Int(e.ErrorCode)
	x.String(e.ErrorMessage)
	return x.buf
}

func (e TL_rpc_answer_unknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_rpc_answer_unknown)
	return x.buf
}

func (e TL_rpc_answer_dropped_running) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_rpc_answer_dropped_running)
	return x.buf
}

func (e TL_rpc_answer_dropped) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_rpc_answer_dropped)
	x.Long(e.MsgID)
	x.Int(e.SeqNo)
	x.Int(e.Bytes)
	return x.buf
}

func (e TL_future_salt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_future_salt)
	x.Int(e.ValidSince)
	x.Int(e.ValidUntil)
	x.Long(e.Salt)
	return x.buf
}

func (e TL_future_salts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_future_salts)
	x.Long(e.ReqMsgID)
	x.Int(e.Now)
	x.Bytes(e.Salts.encode())
	return x.buf
}

func (e TL_pong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pong)
	x.Long(e.MsgID)
	x.Long(e.PingID)
	return x.buf
}

func (e TL_destroy_session_ok) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_session_ok)
	x.Long(e.SessionID)
	return x.buf
}

func (e TL_destroy_session_none) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_session_none)
	x.Long(e.SessionID)
	return x.buf
}

func (e TL_new_session_created) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_new_session_created)
	x.Long(e.FirstMsgID)
	x.Long(e.UniqueID)
	x.Long(e.ServerSalt)
	return x.buf
}

func (e TL_http_wait) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_http_wait)
	x.Int(e.MaxDelay)
	x.Int(e.WaitAfter)
	x.Int(e.MaxWait)
	return x.buf
}

func (e TL_ipPort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_ipPort)
	x.Int(e.Ipv4)
	x.Int(e.Port)
	return x.buf
}

func (e TL_ipPortSecret) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_ipPortSecret)
	x.Int(e.Ipv4)
	x.Int(e.Port)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e TL_accessPointRule) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_accessPointRule)
	x.String(e.PhonePrefixRules)
	x.Int(e.DcID)
	x.Bytes(e.Ips.encode())
	return x.buf
}

func (e TL_help_configSimple) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_configSimple)
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.Rules.encode())
	return x.buf
}

func (e TL_tlsClientHello) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsClientHello)
	x.Bytes(e.Blocks.encode())
	return x.buf
}

func (e TL_tlsBlockString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockString)
	x.String(e.Data)
	return x.buf
}

func (e TL_tlsBlockRandom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockRandom)
	x.Int(e.Length)
	return x.buf
}

func (e TL_tlsBlockZero) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockZero)
	x.Int(e.Length)
	return x.buf
}

func (e TL_tlsBlockDomain) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockDomain)
	return x.buf
}

func (e TL_tlsBlockGrease) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockGrease)
	x.Int(e.Seed)
	return x.buf
}

func (e TL_tlsBlockPublicKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockPublicKey)
	return x.buf
}

func (e TL_tlsBlockScope) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_tlsBlockScope)
	x.Vector(e.Entries)
	return x.buf
}

func (e TL_rpc_drop_answer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_rpc_drop_answer)
	x.Long(e.ReqMsgID)
	return x.buf
}

func (e TL_get_future_salts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_get_future_salts)
	x.Int(e.Num)
	return x.buf
}

func (e TL_ping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_ping)
	x.Long(e.PingID)
	return x.buf
}

func (e TL_ping_delay_disconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_ping_delay_disconnect)
	x.Long(e.PingID)
	x.Int(e.DisconnectDelay)
	return x.buf
}

func (e TL_destroy_session) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_destroy_session)
	x.Long(e.SessionID)
	return x.buf
}

func (e TL_boolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_boolFalse)
	return x.buf
}

func (e TL_boolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_boolTrue)
	return x.buf
}

func (e TL_true) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_true)
	return x.buf
}

func (e TL_error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_error)
	x.Int(e.Code)
	x.String(e.Text)
	return x.buf
}

func (e TL_null) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_null)
	return x.buf
}

func (e TL_inputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerEmpty)
	return x.buf
}

func (e TL_inputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerSelf)
	return x.buf
}

func (e TL_inputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_inputPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerUser)
	x.Int(e.UserID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerChannel)
	x.Int(e.ChannelID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputPeerUserFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerUserFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_inputPeerChannelFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerChannelFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ChannelID)
	return x.buf
}

func (e TL_inputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputUserEmpty)
	return x.buf
}

func (e TL_inputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputUserSelf)
	return x.buf
}

func (e TL_inputUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputUser)
	x.Int(e.UserID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputUserFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputUserFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_inputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPhoneContact)
	x.Long(e.ClientID)
	x.String(e.Phone)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e TL_inputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputFile)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5Checksum)
	return x.buf
}

func (e TL_inputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputFileBig)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.buf
}

func (e TL_inputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaEmpty)
	return x.buf
}

func (e TL_inputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaUploadedPhoto)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	if e.Flags&1 != 0 {
		x.Vector(e.Stickers)
	}
	if e.Flags&2 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_inputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.ID.encode())
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_inputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaGeoPoint)
	x.Bytes(e.GeoPoint.encode())
	return x.buf
}

func (e TL_inputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	return x.buf
}

func (e TL_inputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaUploadedDocument)
	x.Int(e.Flags)
	//flag NosoundVideo
	x.Bytes(e.File.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	if e.Flags&1 != 0 {
		x.Vector(e.Stickers)
	}
	if e.Flags&2 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_inputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.ID.encode())
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_inputMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaVenue)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	return x.buf
}

func (e TL_inputMediaGifExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaGifExternal)
	x.String(e.Url)
	x.String(e.Q)
	return x.buf
}

func (e TL_inputMediaPhotoExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaPhotoExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_inputMediaDocumentExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaDocumentExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_inputMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaGame)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_inputMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	x.Bytes(e.Invoice.encode())
	x.StringBytes(e.Payload)
	x.String(e.Provider)
	x.Bytes(e.ProviderData.encode())
	x.String(e.StartParam)
	return x.buf
}

func (e TL_inputMediaGeoLive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaGeoLive)
	x.Int(e.Flags)
	//flag Stopped
	x.Bytes(e.GeoPoint.encode())
	if e.Flags&2 != 0 {
		x.Int(e.Period)
	}
	return x.buf
}

func (e TL_inputMediaPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaPoll)
	x.Int(e.Flags)
	x.Bytes(e.Poll.encode())
	if e.Flags&1 != 0 {
		x.Vector(e.CorrectAnswers)
	}
	if e.Flags&2 != 0 {
		x.String(e.Solution)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.SolutionEntities)
	}
	return x.buf
}

func (e TL_inputMediaDice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMediaDice)
	x.String(e.Emoticon)
	return x.buf
}

func (e TL_inputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputChatPhotoEmpty)
	return x.buf
}

func (e TL_inputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputChatUploadedPhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_inputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputChatPhoto)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_inputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputGeoPointEmpty)
	return x.buf
}

func (e TL_inputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.buf
}

func (e TL_inputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPhotoEmpty)
	return x.buf
}

func (e TL_inputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPhoto)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	return x.buf
}

func (e TL_inputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputFileLocation)
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	x.Long(e.Secret)
	x.StringBytes(e.FileReference)
	return x.buf
}

func (e TL_inputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputEncryptedFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputDocumentFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.String(e.ThumbSize)
	return x.buf
}

func (e TL_inputSecureFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputSecureFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputTakeoutFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputTakeoutFileLocation)
	return x.buf
}

func (e TL_inputPhotoFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPhotoFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.String(e.ThumbSize)
	return x.buf
}

func (e TL_inputPhotoLegacyFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPhotoLegacyFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	x.Long(e.Secret)
	return x.buf
}

func (e TL_inputPeerPhotoFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerPhotoFileLocation)
	x.Int(e.Flags)
	//flag Big
	x.Bytes(e.Peer.encode())
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	return x.buf
}

func (e TL_inputStickerSetThumb) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetThumb)
	x.Bytes(e.Stickerset.encode())
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	return x.buf
}

func (e TL_peerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerUser)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_peerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_peerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerChannel)
	x.Int(e.ChannelID)
	return x.buf
}

func (e TL_storage_fileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileUnknown)
	return x.buf
}

func (e TL_storage_filePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_filePartial)
	return x.buf
}

func (e TL_storage_fileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileJpeg)
	return x.buf
}

func (e TL_storage_fileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileGif)
	return x.buf
}

func (e TL_storage_filePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_filePng)
	return x.buf
}

func (e TL_storage_filePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_filePdf)
	return x.buf
}

func (e TL_storage_fileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileMp3)
	return x.buf
}

func (e TL_storage_fileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileMov)
	return x.buf
}

func (e TL_storage_fileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileMp4)
	return x.buf
}

func (e TL_storage_fileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_storage_fileWebp)
	return x.buf
}

func (e TL_userEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e TL_user) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_user)
	x.Int(e.Flags)
	//flag Self
	//flag Contact
	//flag MutualContact
	//flag Deleted
	//flag Bot
	//flag BotChatHistory
	//flag BotNochats
	//flag Verified
	//flag Restricted
	//flag Min
	//flag BotInlineGeo
	//flag Support
	//flag Scam
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.Long(e.AccessHash)
	}
	if e.Flags&2 != 0 {
		x.String(e.FirstName)
	}
	if e.Flags&4 != 0 {
		x.String(e.LastName)
	}
	if e.Flags&8 != 0 {
		x.String(e.Username)
	}
	if e.Flags&16 != 0 {
		x.String(e.Phone)
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&64 != 0 {
		x.Bytes(e.Status.encode())
	}
	if e.Flags&16384 != 0 {
		x.Int(e.BotInfoVersion)
	}
	if e.Flags&262144 != 0 {
		x.Vector(e.RestrictionReason)
	}
	if e.Flags&524288 != 0 {
		x.String(e.BotInlinePlaceholder)
	}
	if e.Flags&4194304 != 0 {
		x.String(e.LangCode)
	}
	return x.buf
}

func (e TL_userProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userProfilePhotoEmpty)
	return x.buf
}

func (e TL_userProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userProfilePhoto)
	x.Long(e.PhotoID)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	x.Int(e.DcID)
	return x.buf
}

func (e TL_userStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userStatusEmpty)
	return x.buf
}

func (e TL_userStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userStatusOnline)
	x.Int(e.Expires)
	return x.buf
}

func (e TL_userStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userStatusOffline)
	x.Int(e.WasOnline)
	return x.buf
}

func (e TL_userStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userStatusRecently)
	return x.buf
}

func (e TL_userStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userStatusLastWeek)
	return x.buf
}

func (e TL_userStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userStatusLastMonth)
	return x.buf
}

func (e TL_chatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e TL_chat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chat)
	x.Int(e.Flags)
	//flag Creator
	//flag Kicked
	//flag Left
	//flag Deactivated
	x.Int(e.ID)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	x.Int(e.Date)
	x.Int(e.Version)
	if e.Flags&64 != 0 {
		x.Bytes(e.MigratedTo.encode())
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.AdminRights.encode())
	}
	if e.Flags&262144 != 0 {
		x.Bytes(e.DefaultBannedRights.encode())
	}
	return x.buf
}

func (e TL_chatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatForbidden)
	x.Int(e.ID)
	x.String(e.Title)
	return x.buf
}

func (e TL_channel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channel)
	x.Int(e.Flags)
	//flag Creator
	//flag Left
	//flag Broadcast
	//flag Verified
	//flag Megagroup
	//flag Restricted
	//flag Signatures
	//flag Min
	//flag Scam
	//flag HasLink
	//flag HasGeo
	//flag SlowmodeEnabled
	x.Int(e.ID)
	if e.Flags&8192 != 0 {
		x.Long(e.AccessHash)
	}
	x.String(e.Title)
	if e.Flags&64 != 0 {
		x.String(e.Username)
	}
	x.Bytes(e.Photo.encode())
	x.Int(e.Date)
	x.Int(e.Version)
	if e.Flags&512 != 0 {
		x.Vector(e.RestrictionReason)
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.AdminRights.encode())
	}
	if e.Flags&32768 != 0 {
		x.Bytes(e.BannedRights.encode())
	}
	if e.Flags&262144 != 0 {
		x.Bytes(e.DefaultBannedRights.encode())
	}
	if e.Flags&131072 != 0 {
		x.Int(e.ParticipantsCount)
	}
	return x.buf
}

func (e TL_channelForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelForbidden)
	x.Int(e.Flags)
	//flag Broadcast
	//flag Megagroup
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.String(e.Title)
	if e.Flags&65536 != 0 {
		x.Int(e.UntilDate)
	}
	return x.buf
}

func (e TL_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatFull)
	x.Int(e.Flags)
	//flag CanSetUsername
	//flag HasScheduled
	x.Int(e.ID)
	x.String(e.About)
	x.Bytes(e.Participants.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.ChatPhoto.encode())
	}
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.ExportedInvite.encode())
	if e.Flags&8 != 0 {
		x.Vector(e.BotInfo)
	}
	if e.Flags&64 != 0 {
		x.Int(e.PinnedMsgID)
	}
	if e.Flags&2048 != 0 {
		x.Int(e.FolderID)
	}
	return x.buf
}

func (e TL_channelFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelFull)
	x.Int(e.Flags)
	//flag CanViewParticipants
	//flag CanSetUsername
	//flag CanSetStickers
	//flag HiddenPrehistory
	//flag CanViewStats
	//flag CanSetLocation
	//flag HasScheduled
	x.Int(e.ID)
	x.String(e.About)
	if e.Flags&1 != 0 {
		x.Int(e.ParticipantsCount)
	}
	if e.Flags&2 != 0 {
		x.Int(e.AdminsCount)
	}
	if e.Flags&4 != 0 {
		x.Int(e.KickedCount)
	}
	if e.Flags&4 != 0 {
		x.Int(e.BannedCount)
	}
	if e.Flags&8192 != 0 {
		x.Int(e.OnlineCount)
	}
	x.Int(e.ReadInboxMaxID)
	x.Int(e.ReadOutboxMaxID)
	x.Int(e.UnreadCount)
	x.Bytes(e.ChatPhoto.encode())
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.ExportedInvite.encode())
	x.Vector(e.BotInfo)
	if e.Flags&16 != 0 {
		x.Int(e.MigratedFromChatID)
	}
	if e.Flags&16 != 0 {
		x.Int(e.MigratedFromMaxID)
	}
	if e.Flags&32 != 0 {
		x.Int(e.PinnedMsgID)
	}
	if e.Flags&256 != 0 {
		x.Bytes(e.Stickerset.encode())
	}
	if e.Flags&512 != 0 {
		x.Int(e.AvailableMinID)
	}
	if e.Flags&2048 != 0 {
		x.Int(e.FolderID)
	}
	if e.Flags&16384 != 0 {
		x.Int(e.LinkedChatID)
	}
	if e.Flags&32768 != 0 {
		x.Bytes(e.Location.encode())
	}
	if e.Flags&131072 != 0 {
		x.Int(e.SlowmodeSeconds)
	}
	if e.Flags&262144 != 0 {
		x.Int(e.SlowmodeNextSendDate)
	}
	if e.Flags&4096 != 0 {
		x.Int(e.StatsDc)
	}
	x.Int(e.Pts)
	return x.buf
}

func (e TL_chatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatParticipant)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_chatParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatParticipantCreator)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_chatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatParticipantAdmin)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_chatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatParticipantsForbidden)
	x.Int(e.Flags)
	x.Int(e.ChatID)
	if e.Flags&1 != 0 {
		x.Bytes(e.SelfParticipant.encode())
	}
	return x.buf
}

func (e TL_chatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatParticipants)
	x.Int(e.ChatID)
	x.Vector(e.Participants)
	x.Int(e.Version)
	return x.buf
}

func (e TL_chatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatPhotoEmpty)
	return x.buf
}

func (e TL_chatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatPhoto)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	x.Int(e.DcID)
	return x.buf
}

func (e TL_messageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e TL_message) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_message)
	x.Int(e.Flags)
	//flag Out
	//flag Mentioned
	//flag MediaUnread
	//flag Silent
	//flag Post
	//flag FromScheduled
	//flag Legacy
	//flag EditHide
	x.Int(e.ID)
	if e.Flags&256 != 0 {
		x.Int(e.FromID)
	}
	x.Bytes(e.ToID.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.FwdFrom.encode())
	}
	if e.Flags&2048 != 0 {
		x.Int(e.ViaBotID)
	}
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Int(e.Date)
	x.String(e.Message)
	if e.Flags&512 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&64 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&1024 != 0 {
		x.Int(e.Views)
	}
	if e.Flags&32768 != 0 {
		x.Int(e.EditDate)
	}
	if e.Flags&65536 != 0 {
		x.String(e.PostAuthor)
	}
	if e.Flags&131072 != 0 {
		x.Long(e.GroupedID)
	}
	if e.Flags&4194304 != 0 {
		x.Vector(e.RestrictionReason)
	}
	return x.buf
}

func (e TL_messageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageService)
	x.Int(e.Flags)
	//flag Out
	//flag Mentioned
	//flag MediaUnread
	//flag Silent
	//flag Post
	//flag Legacy
	x.Int(e.ID)
	if e.Flags&256 != 0 {
		x.Int(e.FromID)
	}
	x.Bytes(e.ToID.encode())
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_messageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaEmpty)
	return x.buf
}

func (e TL_messageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaPhoto)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_messageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaGeo)
	x.Bytes(e.Geo.encode())
	return x.buf
}

func (e TL_messageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_messageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaUnsupported)
	return x.buf
}

func (e TL_messageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaDocument)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e TL_messageMediaWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaWebPage)
	x.Bytes(e.Webpage.encode())
	return x.buf
}

func (e TL_messageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaVenue)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	return x.buf
}

func (e TL_messageMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaGame)
	x.Bytes(e.Game.encode())
	return x.buf
}

func (e TL_messageMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaInvoice)
	x.Int(e.Flags)
	//flag ShippingAddressRequested
	//flag Test
	x.String(e.Title)
	x.String(e.Description)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.ReceiptMsgID)
	}
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.String(e.StartParam)
	return x.buf
}

func (e TL_messageMediaGeoLive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaGeoLive)
	x.Bytes(e.Geo.encode())
	x.Int(e.Period)
	return x.buf
}

func (e TL_messageMediaPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaPoll)
	x.Bytes(e.Poll.encode())
	x.Bytes(e.Results.encode())
	return x.buf
}

func (e TL_messageMediaDice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageMediaDice)
	x.Int(e.Value)
	x.String(e.Emoticon)
	return x.buf
}

func (e TL_messageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionEmpty)
	return x.buf
}

func (e TL_messageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_messageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatEditTitle)
	x.String(e.Title)
	return x.buf
}

func (e TL_messageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatEditPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TL_messageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatDeletePhoto)
	return x.buf
}

func (e TL_messageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatAddUser)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_messageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatDeleteUser)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_messageActionChatJoinedByLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatJoinedByLink)
	x.Int(e.InviterID)
	return x.buf
}

func (e TL_messageActionChannelCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChannelCreate)
	x.String(e.Title)
	return x.buf
}

func (e TL_messageActionChatMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChatMigrateTo)
	x.Int(e.ChannelID)
	return x.buf
}

func (e TL_messageActionChannelMigrateFrom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionChannelMigrateFrom)
	x.String(e.Title)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_messageActionPinMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionPinMessage)
	return x.buf
}

func (e TL_messageActionHistoryClear) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionHistoryClear)
	return x.buf
}

func (e TL_messageActionGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionGameScore)
	x.Long(e.GameID)
	x.Int(e.Score)
	return x.buf
}

func (e TL_messageActionPaymentSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionPaymentSentMe)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.StringBytes(e.Payload)
	if e.Flags&1 != 0 {
		x.Bytes(e.Info.encode())
	}
	if e.Flags&2 != 0 {
		x.String(e.ShippingOptionID)
	}
	x.Bytes(e.Charge.encode())
	return x.buf
}

func (e TL_messageActionPaymentSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionPaymentSent)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	return x.buf
}

func (e TL_messageActionPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionPhoneCall)
	x.Int(e.Flags)
	//flag Video
	x.Long(e.CallID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Reason.encode())
	}
	if e.Flags&2 != 0 {
		x.Int(e.Duration)
	}
	return x.buf
}

func (e TL_messageActionScreenshotTaken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionScreenshotTaken)
	return x.buf
}

func (e TL_messageActionCustomAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionCustomAction)
	x.String(e.Message)
	return x.buf
}

func (e TL_messageActionBotAllowed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionBotAllowed)
	x.String(e.Domain)
	return x.buf
}

func (e TL_messageActionSecureValuesSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionSecureValuesSentMe)
	x.Vector(e.Values)
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e TL_messageActionSecureValuesSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionSecureValuesSent)
	x.Vector(e.Types)
	return x.buf
}

func (e TL_messageActionContactSignUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageActionContactSignUp)
	return x.buf
}

func (e TL_dialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dialog)
	x.Int(e.Flags)
	//flag Pinned
	//flag UnreadMark
	x.Bytes(e.Peer.encode())
	x.Int(e.TopMessage)
	x.Int(e.ReadInboxMaxID)
	x.Int(e.ReadOutboxMaxID)
	x.Int(e.UnreadCount)
	x.Int(e.UnreadMentionsCount)
	x.Bytes(e.NotifySettings.encode())
	if e.Flags&1 != 0 {
		x.Int(e.Pts)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Draft.encode())
	}
	if e.Flags&16 != 0 {
		x.Int(e.FolderID)
	}
	return x.buf
}

func (e TL_dialogFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dialogFolder)
	x.Int(e.Flags)
	//flag Pinned
	x.Bytes(e.Folder.encode())
	x.Bytes(e.Peer.encode())
	x.Int(e.TopMessage)
	x.Int(e.UnreadMutedPeersCount)
	x.Int(e.UnreadUnmutedPeersCount)
	x.Int(e.UnreadMutedMessagesCount)
	x.Int(e.UnreadUnmutedMessagesCount)
	return x.buf
}

func (e TL_photoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photoEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e TL_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photo)
	x.Int(e.Flags)
	//flag HasStickers
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.Int(e.Date)
	x.Vector(e.Sizes)
	x.Int(e.DcID)
	return x.buf
}

func (e TL_photoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photoSizeEmpty)
	x.String(e.Type)
	return x.buf
}

func (e TL_photoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photoSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}

func (e TL_photoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photoCachedSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_photoStrippedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photoStrippedSize)
	x.String(e.Type)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_geoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_geoPointEmpty)
	return x.buf
}

func (e TL_geoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_geoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_auth_sentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_sentCode)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	x.String(e.PhoneCodeHash)
	if e.Flags&2 != 0 {
		x.Bytes(e.NextType.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.Timeout)
	}
	return x.buf
}

func (e TL_auth_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_authorization)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.TmpSessions)
	}
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TL_auth_authorizationSignUpRequired) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_authorizationSignUpRequired)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.TermsOfService.encode())
	}
	return x.buf
}

func (e TL_auth_exportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_exportedAuthorization)
	x.Int(e.ID)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_inputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_inputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputNotifyUsers)
	return x.buf
}

func (e TL_inputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputNotifyChats)
	return x.buf
}

func (e TL_inputNotifyBroadcasts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputNotifyBroadcasts)
	return x.buf
}

func (e TL_inputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPeerNotifySettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.ShowPreviews.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Silent.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.MuteUntil)
	}
	if e.Flags&8 != 0 {
		x.String(e.Sound)
	}
	return x.buf
}

func (e TL_peerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerNotifySettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.ShowPreviews.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Silent.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.MuteUntil)
	}
	if e.Flags&8 != 0 {
		x.String(e.Sound)
	}
	return x.buf
}

func (e TL_peerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerSettings)
	x.Int(e.Flags)
	//flag ReportSpam
	//flag AddContact
	//flag BlockContact
	//flag ShareContact
	//flag NeedContactsException
	//flag ReportGeo
	return x.buf
}

func (e TL_wallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_wallPaper)
	x.Long(e.ID)
	x.Int(e.Flags)
	//flag Creator
	//flag Default
	//flag Pattern
	//flag Dark
	x.Long(e.AccessHash)
	x.String(e.Slug)
	x.Bytes(e.Document.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e TL_wallPaperNoFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_wallPaperNoFile)
	x.Int(e.Flags)
	//flag Default
	//flag Dark
	if e.Flags&4 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e TL_inputReportReasonSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonSpam)
	return x.buf
}

func (e TL_inputReportReasonViolence) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonViolence)
	return x.buf
}

func (e TL_inputReportReasonPornography) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonPornography)
	return x.buf
}

func (e TL_inputReportReasonChildAbuse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonChildAbuse)
	return x.buf
}

func (e TL_inputReportReasonOther) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonOther)
	x.String(e.Text)
	return x.buf
}

func (e TL_inputReportReasonCopyright) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonCopyright)
	return x.buf
}

func (e TL_inputReportReasonGeoIrrelevant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputReportReasonGeoIrrelevant)
	return x.buf
}

func (e TL_userFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_userFull)
	x.Int(e.Flags)
	//flag Blocked
	//flag PhoneCallsAvailable
	//flag PhoneCallsPrivate
	//flag CanPinMessage
	//flag HasScheduled
	x.Bytes(e.User.encode())
	if e.Flags&2 != 0 {
		x.String(e.About)
	}
	x.Bytes(e.Settings.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.ProfilePhoto.encode())
	}
	x.Bytes(e.NotifySettings.encode())
	if e.Flags&8 != 0 {
		x.Bytes(e.BotInfo.encode())
	}
	if e.Flags&64 != 0 {
		x.Int(e.PinnedMsgID)
	}
	x.Int(e.CommonChatsCount)
	if e.Flags&2048 != 0 {
		x.Int(e.FolderID)
	}
	return x.buf
}

func (e TL_contact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contact)
	x.Int(e.UserID)
	x.Bytes(e.Mutual.encode())
	return x.buf
}

func (e TL_importedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_importedContact)
	x.Int(e.UserID)
	x.Long(e.ClientID)
	return x.buf
}

func (e TL_contactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contactBlocked)
	x.Int(e.UserID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_contactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contactStatus)
	x.Int(e.UserID)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TL_contacts_contactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_contactsNotModified)
	return x.buf
}

func (e TL_contacts_contacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_contacts)
	x.Vector(e.Contacts)
	x.Int(e.SavedCount)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_importedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_importedContacts)
	x.Vector(e.Imported)
	x.Vector(e.PopularInvites)
	x.VectorLong(e.RetryContacts)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_blocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_blocked)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_blockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_blockedSlice)
	x.Int(e.Count)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_dialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_dialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_dialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_dialogsSlice)
	x.Int(e.Count)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_dialogsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_dialogsNotModified)
	x.Int(e.Count)
	return x.buf
}

func (e TL_messages_messages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_messages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_messagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_messagesSlice)
	x.Int(e.Flags)
	//flag Inexact
	x.Int(e.Count)
	if e.Flags&1 != 0 {
		x.Int(e.NextRate)
	}
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_channelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_channelMessages)
	x.Int(e.Flags)
	//flag Inexact
	x.Int(e.Pts)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_messagesNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_messagesNotModified)
	x.Int(e.Count)
	return x.buf
}

func (e TL_messages_chats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_chats)
	x.Vector(e.Chats)
	return x.buf
}

func (e TL_messages_chatsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_chatsSlice)
	x.Int(e.Count)
	x.Vector(e.Chats)
	return x.buf
}

func (e TL_messages_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_chatFull)
	x.Bytes(e.FullChat.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messages_affectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_affectedHistory)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_inputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterEmpty)
	return x.buf
}

func (e TL_inputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterPhotos)
	return x.buf
}

func (e TL_inputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterPhotoVideo)
	return x.buf
}

func (e TL_inputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterDocument)
	return x.buf
}

func (e TL_inputMessagesFilterUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterUrl)
	return x.buf
}

func (e TL_inputMessagesFilterGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterGif)
	return x.buf
}

func (e TL_inputMessagesFilterVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterVoice)
	return x.buf
}

func (e TL_inputMessagesFilterMusic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterMusic)
	return x.buf
}

func (e TL_inputMessagesFilterChatPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterChatPhotos)
	return x.buf
}

func (e TL_inputMessagesFilterPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterPhoneCalls)
	x.Int(e.Flags)
	//flag Missed
	return x.buf
}

func (e TL_inputMessagesFilterRoundVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterRoundVoice)
	return x.buf
}

func (e TL_inputMessagesFilterRoundVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterRoundVideo)
	return x.buf
}

func (e TL_inputMessagesFilterMyMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterMyMentions)
	return x.buf
}

func (e TL_inputMessagesFilterGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterGeo)
	return x.buf
}

func (e TL_inputMessagesFilterContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagesFilterContacts)
	return x.buf
}

func (e TL_updateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateNewMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateMessageID)
	x.Int(e.ID)
	x.Long(e.RandomID)
	return x.buf
}

func (e TL_updateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserTyping)
	x.Int(e.UserID)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_updateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatUserTyping)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_updateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatParticipants)
	x.Bytes(e.Participants.encode())
	return x.buf
}

func (e TL_updateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserStatus)
	x.Int(e.UserID)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e TL_updateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserName)
	x.Int(e.UserID)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	return x.buf
}

func (e TL_updateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserPhoto)
	x.Int(e.UserID)
	x.Int(e.Date)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Previous.encode())
	return x.buf
}

func (e TL_updateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateNewEncryptedMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Qts)
	return x.buf
}

func (e TL_updateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateEncryptedChatTyping)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_updateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateEncryption)
	x.Bytes(e.Chat.encode())
	x.Int(e.Date)
	return x.buf
}

func (e TL_updateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateEncryptedMessagesRead)
	x.Int(e.ChatID)
	x.Int(e.MaxDate)
	x.Int(e.Date)
	return x.buf
}

func (e TL_updateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatParticipantAdd)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatParticipantDelete)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDcOptions)
	x.Vector(e.DcOptions)
	return x.buf
}

func (e TL_updateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserBlocked)
	x.Int(e.UserID)
	x.Bytes(e.Blocked.encode())
	return x.buf
}

func (e TL_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.NotifySettings.encode())
	return x.buf
}

func (e TL_updateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateServiceNotification)
	x.Int(e.Flags)
	//flag Popup
	if e.Flags&2 != 0 {
		x.Int(e.InboxDate)
	}
	x.String(e.Type)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_updatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e TL_updateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserPhone)
	x.Int(e.UserID)
	x.String(e.Phone)
	return x.buf
}

func (e TL_updateReadHistoryInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateReadHistoryInbox)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	x.Int(e.StillUnreadCount)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateReadHistoryOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateReadHistoryOutbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateWebPage)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateReadMessagesContents)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateChannelTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannelTooLong)
	x.Int(e.Flags)
	x.Int(e.ChannelID)
	if e.Flags&1 != 0 {
		x.Int(e.Pts)
	}
	return x.buf
}

func (e TL_updateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannel)
	x.Int(e.ChannelID)
	return x.buf
}

func (e TL_updateNewChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateNewChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateReadChannelInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateReadChannelInbox)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.Int(e.ChannelID)
	x.Int(e.MaxID)
	x.Int(e.StillUnreadCount)
	x.Int(e.Pts)
	return x.buf
}

func (e TL_updateDeleteChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDeleteChannelMessages)
	x.Int(e.ChannelID)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateChannelMessageViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannelMessageViews)
	x.Int(e.ChannelID)
	x.Int(e.ID)
	x.Int(e.Views)
	return x.buf
}

func (e TL_updateChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatParticipantAdmin)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Bytes(e.IsAdmin.encode())
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateNewStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateNewStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_updateStickerSetsOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateStickerSetsOrder)
	x.Int(e.Flags)
	//flag Masks
	x.VectorLong(e.Order)
	return x.buf
}

func (e TL_updateStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateStickerSets)
	return x.buf
}

func (e TL_updateSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateSavedGifs)
	return x.buf
}

func (e TL_updateBotInlineQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotInlineQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.String(e.Query)
	if e.Flags&1 != 0 {
		x.Bytes(e.Geo.encode())
	}
	x.String(e.Offset)
	return x.buf
}

func (e TL_updateBotInlineSend) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotInlineSend)
	x.Int(e.Flags)
	x.Int(e.UserID)
	x.String(e.Query)
	if e.Flags&1 != 0 {
		x.Bytes(e.Geo.encode())
	}
	x.String(e.ID)
	if e.Flags&2 != 0 {
		x.Bytes(e.MsgID.encode())
	}
	return x.buf
}

func (e TL_updateEditChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateEditChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateChannelPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannelPinnedMessage)
	x.Int(e.ChannelID)
	x.Int(e.ID)
	return x.buf
}

func (e TL_updateBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Long(e.ChatInstance)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Data)
	}
	if e.Flags&2 != 0 {
		x.String(e.GameShortName)
	}
	return x.buf
}

func (e TL_updateEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateEditMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateInlineBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateInlineBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.Bytes(e.MsgID.encode())
	x.Long(e.ChatInstance)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Data)
	}
	if e.Flags&2 != 0 {
		x.String(e.GameShortName)
	}
	return x.buf
}

func (e TL_updateReadChannelOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateReadChannelOutbox)
	x.Int(e.ChannelID)
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_updateDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDraftMessage)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Draft.encode())
	return x.buf
}

func (e TL_updateReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateReadFeaturedStickers)
	return x.buf
}

func (e TL_updateRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateRecentStickers)
	return x.buf
}

func (e TL_updateConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateConfig)
	return x.buf
}

func (e TL_updatePtsChanged) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePtsChanged)
	return x.buf
}

func (e TL_updateChannelWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannelWebPage)
	x.Int(e.ChannelID)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updateDialogPinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDialogPinned)
	x.Int(e.Flags)
	//flag Pinned
	if e.Flags&2 != 0 {
		x.Int(e.FolderID)
	}
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_updatePinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePinnedDialogs)
	x.Int(e.Flags)
	if e.Flags&2 != 0 {
		x.Int(e.FolderID)
	}
	if e.Flags&1 != 0 {
		x.Vector(e.Order)
	}
	return x.buf
}

func (e TL_updateBotWebhookJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotWebhookJSON)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_updateBotWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotWebhookJSONQuery)
	x.Long(e.QueryID)
	x.Bytes(e.Data.encode())
	x.Int(e.Timeout)
	return x.buf
}

func (e TL_updateBotShippingQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotShippingQuery)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.StringBytes(e.Payload)
	x.Bytes(e.ShippingAddress.encode())
	return x.buf
}

func (e TL_updateBotPrecheckoutQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateBotPrecheckoutQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.StringBytes(e.Payload)
	if e.Flags&1 != 0 {
		x.Bytes(e.Info.encode())
	}
	if e.Flags&2 != 0 {
		x.String(e.ShippingOptionID)
	}
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	return x.buf
}

func (e TL_updatePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePhoneCall)
	x.Bytes(e.PhoneCall.encode())
	return x.buf
}

func (e TL_updateLangPackTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateLangPackTooLong)
	x.String(e.LangCode)
	return x.buf
}

func (e TL_updateLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateLangPack)
	x.Bytes(e.Difference.encode())
	return x.buf
}

func (e TL_updateFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateFavedStickers)
	return x.buf
}

func (e TL_updateChannelReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannelReadMessagesContents)
	x.Int(e.ChannelID)
	x.VectorInt(e.Messages)
	return x.buf
}

func (e TL_updateContactsReset) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateContactsReset)
	return x.buf
}

func (e TL_updateChannelAvailableMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChannelAvailableMessages)
	x.Int(e.ChannelID)
	x.Int(e.AvailableMinID)
	return x.buf
}

func (e TL_updateDialogUnreadMark) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDialogUnreadMark)
	x.Int(e.Flags)
	//flag Unread
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_updateUserPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateUserPinnedMessage)
	x.Int(e.UserID)
	x.Int(e.ID)
	return x.buf
}

func (e TL_updateChatPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatPinnedMessage)
	x.Int(e.ChatID)
	x.Int(e.ID)
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateMessagePoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateMessagePoll)
	x.Int(e.Flags)
	x.Long(e.PollID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Poll.encode())
	}
	x.Bytes(e.Results.encode())
	return x.buf
}

func (e TL_updateChatDefaultBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateChatDefaultBannedRights)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.DefaultBannedRights.encode())
	x.Int(e.Version)
	return x.buf
}

func (e TL_updateFolderPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateFolderPeers)
	x.Vector(e.FolderPeers)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_updatePeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePeerSettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_updatePeerLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePeerLocated)
	x.Vector(e.Peers)
	return x.buf
}

func (e TL_updateNewScheduledMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateNewScheduledMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TL_updateDeleteScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDeleteScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.Messages)
	return x.buf
}

func (e TL_updateTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateTheme)
	x.Bytes(e.Theme.encode())
	return x.buf
}

func (e TL_updateGeoLiveViewed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateGeoLiveViewed)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	return x.buf
}

func (e TL_updateLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateLoginToken)
	return x.buf
}

func (e TL_updateMessagePollVote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateMessagePollVote)
	x.Long(e.PollID)
	x.Int(e.UserID)
	x.Vector(e.Options)
	return x.buf
}

func (e TL_updateDialogFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDialogFilter)
	x.Int(e.Flags)
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Filter.encode())
	}
	return x.buf
}

func (e TL_updateDialogFilterOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDialogFilterOrder)
	x.VectorInt(e.Order)
	return x.buf
}

func (e TL_updateDialogFilters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateDialogFilters)
	return x.buf
}

func (e TL_updatePhoneCallSignalingData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatePhoneCallSignalingData)
	x.Long(e.PhoneCallID)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_updates_state) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_state)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.UnreadCount)
	return x.buf
}

func (e TL_updates_differenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_differenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TL_updates_difference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_difference)
	x.Vector(e.NewMessages)
	x.Vector(e.NewEncryptedMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e TL_updates_differenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_differenceSlice)
	x.Vector(e.NewMessages)
	x.Vector(e.NewEncryptedMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.IntermediateState.encode())
	return x.buf
}

func (e TL_updates_differenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_differenceTooLong)
	x.Int(e.Pts)
	return x.buf
}

func (e TL_updatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatesTooLong)
	return x.buf
}

func (e TL_updateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateShortMessage)
	x.Int(e.Flags)
	//flag Out
	//flag Mentioned
	//flag MediaUnread
	//flag Silent
	x.Int(e.ID)
	x.Int(e.UserID)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	if e.Flags&4 != 0 {
		x.Bytes(e.FwdFrom.encode())
	}
	if e.Flags&2048 != 0 {
		x.Int(e.ViaBotID)
	}
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_updateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateShortChatMessage)
	x.Int(e.Flags)
	//flag Out
	//flag Mentioned
	//flag MediaUnread
	//flag Silent
	x.Int(e.ID)
	x.Int(e.FromID)
	x.Int(e.ChatID)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	if e.Flags&4 != 0 {
		x.Bytes(e.FwdFrom.encode())
	}
	if e.Flags&2048 != 0 {
		x.Int(e.ViaBotID)
	}
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_updateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateShort)
	x.Bytes(e.Update.encode())
	x.Int(e.Date)
	return x.buf
}

func (e TL_updatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updatesCombined)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.SeqStart)
	x.Int(e.Seq)
	return x.buf
}

func (e TL_updates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e TL_updateShortSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updateShortSentMessage)
	x.Int(e.Flags)
	//flag Out
	x.Int(e.ID)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	if e.Flags&512 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_photos_photos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_photos)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_photos_photosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_photosSlice)
	x.Int(e.Count)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_photos_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_photo)
	x.Bytes(e.Photo.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_upload_file) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_file)
	x.Bytes(e.Type.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_upload_fileCdnRedirect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_fileCdnRedirect)
	x.Int(e.DcID)
	x.StringBytes(e.FileToken)
	x.StringBytes(e.EncryptionKey)
	x.StringBytes(e.EncryptionIv)
	x.Vector(e.FileHashes)
	return x.buf
}

func (e TL_dcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dcOption)
	x.Int(e.Flags)
	//flag Ipv6
	//flag MediaOnly
	//flag TcpoOnly
	//flag Cdn
	//flag Static
	x.Int(e.ID)
	x.String(e.IpAddress)
	x.Int(e.Port)
	if e.Flags&1024 != 0 {
		x.StringBytes(e.Secret)
	}
	return x.buf
}

func (e TL_config) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_config)
	x.Int(e.Flags)
	//flag PhonecallsEnabled
	//flag DefaultP2pContacts
	//flag PreloadFeaturedStickers
	//flag IgnorePhoneEntities
	//flag RevokePmInbox
	//flag BlockedMode
	//flag PfsEnabled
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.TestMode.encode())
	x.Int(e.ThisDc)
	x.Vector(e.DcOptions)
	x.String(e.DcTxtDomainName)
	x.Int(e.ChatSizeMax)
	x.Int(e.MegagroupSizeMax)
	x.Int(e.ForwardedCountMax)
	x.Int(e.OnlineUpdatePeriodMs)
	x.Int(e.OfflineBlurTimeoutMs)
	x.Int(e.OfflineIdleTimeoutMs)
	x.Int(e.OnlineCloudTimeoutMs)
	x.Int(e.NotifyCloudDelayMs)
	x.Int(e.NotifyDefaultDelayMs)
	x.Int(e.PushChatPeriodMs)
	x.Int(e.PushChatLimit)
	x.Int(e.SavedGifsLimit)
	x.Int(e.EditTimeLimit)
	x.Int(e.RevokeTimeLimit)
	x.Int(e.RevokePmTimeLimit)
	x.Int(e.RatingEDecay)
	x.Int(e.StickersRecentLimit)
	x.Int(e.StickersFavedLimit)
	x.Int(e.ChannelsReadMediaPeriod)
	if e.Flags&1 != 0 {
		x.Int(e.TmpSessions)
	}
	x.Int(e.PinnedDialogsCountMax)
	x.Int(e.PinnedInfolderCountMax)
	x.Int(e.CallReceiveTimeoutMs)
	x.Int(e.CallRingTimeoutMs)
	x.Int(e.CallConnectTimeoutMs)
	x.Int(e.CallPacketTimeoutMs)
	x.String(e.MeUrlPrefix)
	if e.Flags&128 != 0 {
		x.String(e.AutoupdateUrlPrefix)
	}
	if e.Flags&512 != 0 {
		x.String(e.GifSearchUsername)
	}
	if e.Flags&1024 != 0 {
		x.String(e.VenueSearchUsername)
	}
	if e.Flags&2048 != 0 {
		x.String(e.ImgSearchUsername)
	}
	if e.Flags&4096 != 0 {
		x.String(e.StaticMapsProvider)
	}
	x.Int(e.CaptionLengthMax)
	x.Int(e.MessageLengthMax)
	x.Int(e.WebfileDcID)
	if e.Flags&4 != 0 {
		x.String(e.SuggestedLangCode)
	}
	if e.Flags&4 != 0 {
		x.Int(e.LangPackVersion)
	}
	if e.Flags&4 != 0 {
		x.Int(e.BaseLangPackVersion)
	}
	return x.buf
}

func (e TL_nearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_nearestDc)
	x.String(e.Country)
	x.Int(e.ThisDc)
	x.Int(e.NearestDc)
	return x.buf
}

func (e TL_help_appUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_appUpdate)
	x.Int(e.Flags)
	//flag CanNotSkip
	x.Int(e.ID)
	x.String(e.Version)
	x.String(e.Text)
	x.Vector(e.Entities)
	if e.Flags&2 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&4 != 0 {
		x.String(e.Url)
	}
	return x.buf
}

func (e TL_help_noAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_noAppUpdate)
	return x.buf
}

func (e TL_help_inviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_inviteText)
	x.String(e.Message)
	return x.buf
}

func (e TL_encryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedChatEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e TL_encryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedChatWaiting)
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	return x.buf
}

func (e TL_encryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedChatRequested)
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GA)
	return x.buf
}

func (e TL_encryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedChat)
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	return x.buf
}

func (e TL_encryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedChatDiscarded)
	x.Int(e.ID)
	return x.buf
}

func (e TL_inputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputEncryptedChat)
	x.Int(e.ChatID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_encryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedFileEmpty)
	return x.buf
}

func (e TL_encryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.Int(e.DcID)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e TL_inputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputEncryptedFileEmpty)
	return x.buf
}

func (e TL_inputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputEncryptedFileUploaded)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Md5Checksum)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e TL_inputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputEncryptedFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputEncryptedFileBigUploaded)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e TL_encryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedMessage)
	x.Long(e.RandomID)
	x.Int(e.ChatID)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_encryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_encryptedMessageService)
	x.Long(e.RandomID)
	x.Int(e.ChatID)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_messages_dhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_dhConfigNotModified)
	x.StringBytes(e.Random)
	return x.buf
}

func (e TL_messages_dhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_dhConfig)
	x.Int(e.G)
	x.StringBytes(e.P)
	x.Int(e.Version)
	x.StringBytes(e.Random)
	return x.buf
}

func (e TL_messages_sentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sentEncryptedMessage)
	x.Int(e.Date)
	return x.buf
}

func (e TL_messages_sentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_inputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputDocumentEmpty)
	return x.buf
}

func (e TL_inputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputDocument)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	return x.buf
}

func (e TL_documentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e TL_document) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_document)
	x.Int(e.Flags)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.Int(e.Date)
	x.String(e.MimeType)
	x.Int(e.Size)
	if e.Flags&1 != 0 {
		x.Vector(e.Thumbs)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.VideoThumbs)
	}
	x.Int(e.DcID)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TL_help_support) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_support)
	x.String(e.PhoneNumber)
	x.Bytes(e.User.encode())
	return x.buf
}

func (e TL_notifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_notifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_notifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_notifyUsers)
	return x.buf
}

func (e TL_notifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_notifyChats)
	return x.buf
}

func (e TL_notifyBroadcasts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_notifyBroadcasts)
	return x.buf
}

func (e TL_sendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageTypingAction)
	return x.buf
}

func (e TL_sendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageCancelAction)
	return x.buf
}

func (e TL_sendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageRecordVideoAction)
	return x.buf
}

func (e TL_sendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageUploadVideoAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageRecordAudioAction)
	return x.buf
}

func (e TL_sendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageUploadAudioAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageUploadPhotoAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageUploadDocumentAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_sendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageGeoLocationAction)
	return x.buf
}

func (e TL_sendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageChooseContactAction)
	return x.buf
}

func (e TL_sendMessageGamePlayAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageGamePlayAction)
	return x.buf
}

func (e TL_sendMessageRecordRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageRecordRoundAction)
	return x.buf
}

func (e TL_sendMessageUploadRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_sendMessageUploadRoundAction)
	x.Int(e.Progress)
	return x.buf
}

func (e TL_contacts_found) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_found)
	x.Vector(e.MyResults)
	x.Vector(e.Results)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_inputPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyChatInvite)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyPhoneCall)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneP2P) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyPhoneP2P)
	return x.buf
}

func (e TL_inputPrivacyKeyForwards) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyForwards)
	return x.buf
}

func (e TL_inputPrivacyKeyProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyProfilePhoto)
	return x.buf
}

func (e TL_inputPrivacyKeyPhoneNumber) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyPhoneNumber)
	return x.buf
}

func (e TL_inputPrivacyKeyAddedByPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyKeyAddedByPhone)
	return x.buf
}

func (e TL_privacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_privacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyChatInvite)
	return x.buf
}

func (e TL_privacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyPhoneCall)
	return x.buf
}

func (e TL_privacyKeyPhoneP2P) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyPhoneP2P)
	return x.buf
}

func (e TL_privacyKeyForwards) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyForwards)
	return x.buf
}

func (e TL_privacyKeyProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyProfilePhoto)
	return x.buf
}

func (e TL_privacyKeyPhoneNumber) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyPhoneNumber)
	return x.buf
}

func (e TL_privacyKeyAddedByPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyKeyAddedByPhone)
	return x.buf
}

func (e TL_inputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueAllowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueAllowAll)
	return x.buf
}

func (e TL_inputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueAllowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueDisallowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueDisallowAll)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueDisallowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputPrivacyValueAllowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueAllowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPrivacyValueDisallowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e TL_privacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueAllowContacts)
	return x.buf
}

func (e TL_privacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueAllowAll)
	return x.buf
}

func (e TL_privacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_privacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueDisallowContacts)
	return x.buf
}

func (e TL_privacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueDisallowAll)
	return x.buf
}

func (e TL_privacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e TL_privacyValueAllowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueAllowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e TL_privacyValueDisallowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_privacyValueDisallowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e TL_account_privacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_privacyRules)
	x.Vector(e.Rules)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_accountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_accountDaysTTL)
	x.Int(e.Days)
	return x.buf
}

func (e TL_documentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_documentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeAnimated)
	return x.buf
}

func (e TL_documentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeSticker)
	x.Int(e.Flags)
	//flag Mask
	x.String(e.Alt)
	x.Bytes(e.Stickerset.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.MaskCoords.encode())
	}
	return x.buf
}

func (e TL_documentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeVideo)
	x.Int(e.Flags)
	//flag RoundMessage
	//flag SupportsStreaming
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_documentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeAudio)
	x.Int(e.Flags)
	//flag Voice
	x.Int(e.Duration)
	if e.Flags&1 != 0 {
		x.String(e.Title)
	}
	if e.Flags&2 != 0 {
		x.String(e.Performer)
	}
	if e.Flags&4 != 0 {
		x.StringBytes(e.Waveform)
	}
	return x.buf
}

func (e TL_documentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeFilename)
	x.String(e.FileName)
	return x.buf
}

func (e TL_documentAttributeHasStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_documentAttributeHasStickers)
	return x.buf
}

func (e TL_messages_stickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_stickersNotModified)
	return x.buf
}

func (e TL_messages_stickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_stickers)
	x.Int(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_stickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.buf
}

func (e TL_messages_allStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_allStickersNotModified)
	return x.buf
}

func (e TL_messages_allStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_allStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_messages_affectedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_affectedMessages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e TL_webPageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webPageEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e TL_webPagePending) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webPagePending)
	x.Long(e.ID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_webPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webPage)
	x.Int(e.Flags)
	x.Long(e.ID)
	x.String(e.Url)
	x.String(e.DisplayUrl)
	x.Int(e.Hash)
	if e.Flags&1 != 0 {
		x.String(e.Type)
	}
	if e.Flags&2 != 0 {
		x.String(e.SiteName)
	}
	if e.Flags&4 != 0 {
		x.String(e.Title)
	}
	if e.Flags&8 != 0 {
		x.String(e.Description)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&32 != 0 {
		x.String(e.EmbedUrl)
	}
	if e.Flags&32 != 0 {
		x.String(e.EmbedType)
	}
	if e.Flags&64 != 0 {
		x.Int(e.EmbedWidth)
	}
	if e.Flags&64 != 0 {
		x.Int(e.EmbedHeight)
	}
	if e.Flags&128 != 0 {
		x.Int(e.Duration)
	}
	if e.Flags&256 != 0 {
		x.String(e.Author)
	}
	if e.Flags&512 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&1024 != 0 {
		x.Bytes(e.CachedPage.encode())
	}
	if e.Flags&4096 != 0 {
		x.Vector(e.Attributes)
	}
	return x.buf
}

func (e TL_webPageNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webPageNotModified)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.CachedPageViews)
	}
	return x.buf
}

func (e TL_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_authorization)
	x.Int(e.Flags)
	//flag Current
	//flag OfficialApp
	//flag PasswordPending
	x.Long(e.Hash)
	x.String(e.DeviceModel)
	x.String(e.Platform)
	x.String(e.SystemVersion)
	x.Int(e.ApiID)
	x.String(e.AppName)
	x.String(e.AppVersion)
	x.Int(e.DateCreated)
	x.Int(e.DateActive)
	x.String(e.Ip)
	x.String(e.Country)
	x.String(e.Region)
	return x.buf
}

func (e TL_account_authorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_authorizations)
	x.Vector(e.Authorizations)
	return x.buf
}

func (e TL_account_password) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_password)
	x.Int(e.Flags)
	//flag HasRecovery
	//flag HasSecureValues
	//flag HasPassword
	if e.Flags&4 != 0 {
		x.Bytes(e.CurrentAlgo.encode())
	}
	if e.Flags&4 != 0 {
		x.StringBytes(e.SrpB)
	}
	if e.Flags&4 != 0 {
		x.Long(e.SrpID)
	}
	if e.Flags&8 != 0 {
		x.String(e.Hint)
	}
	if e.Flags&16 != 0 {
		x.String(e.EmailUnconfirmedPattern)
	}
	x.Bytes(e.NewAlgo.encode())
	x.Bytes(e.NewSecureAlgo.encode())
	x.StringBytes(e.SecureRandom)
	return x.buf
}

func (e TL_account_passwordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_passwordSettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.Email)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.SecureSettings.encode())
	}
	return x.buf
}

func (e TL_account_passwordInputSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_passwordInputSettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.NewAlgo.encode())
	}
	if e.Flags&1 != 0 {
		x.StringBytes(e.NewPasswordHash)
	}
	if e.Flags&1 != 0 {
		x.String(e.Hint)
	}
	if e.Flags&2 != 0 {
		x.String(e.Email)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.NewSecureSettings.encode())
	}
	return x.buf
}

func (e TL_auth_passwordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_passwordRecovery)
	x.String(e.EmailPattern)
	return x.buf
}

func (e TL_receivedNotifyMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_receivedNotifyMessage)
	x.Int(e.ID)
	x.Int(e.Flags)
	return x.buf
}

func (e TL_chatInviteEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatInviteEmpty)
	return x.buf
}

func (e TL_chatInviteExported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatInviteExported)
	x.String(e.Link)
	return x.buf
}

func (e TL_chatInviteAlready) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatInviteAlready)
	x.Bytes(e.Chat.encode())
	return x.buf
}

func (e TL_chatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatInvite)
	x.Int(e.Flags)
	//flag Channel
	//flag Broadcast
	//flag Public
	//flag Megagroup
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	if e.Flags&16 != 0 {
		x.Vector(e.Participants)
	}
	return x.buf
}

func (e TL_inputStickerSetEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetEmpty)
	return x.buf
}

func (e TL_inputStickerSetID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetID)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputStickerSetShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetShortName)
	x.String(e.ShortName)
	return x.buf
}

func (e TL_inputStickerSetAnimatedEmoji) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetAnimatedEmoji)
	return x.buf
}

func (e TL_inputStickerSetDice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetDice)
	x.String(e.Emoticon)
	return x.buf
}

func (e TL_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickerSet)
	x.Int(e.Flags)
	//flag Archived
	//flag Official
	//flag Masks
	//flag Animated
	if e.Flags&1 != 0 {
		x.Int(e.InstalledDate)
	}
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.String(e.Title)
	x.String(e.ShortName)
	if e.Flags&16 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	if e.Flags&16 != 0 {
		x.Int(e.ThumbDcID)
	}
	x.Int(e.Count)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_stickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_stickerSet)
	x.Bytes(e.Set.encode())
	x.Vector(e.Packs)
	x.Vector(e.Documents)
	return x.buf
}

func (e TL_botCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botCommand)
	x.String(e.Command)
	x.String(e.Description)
	return x.buf
}

func (e TL_botInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInfo)
	x.Int(e.UserID)
	x.String(e.Description)
	x.Vector(e.Commands)
	return x.buf
}

func (e TL_keyboardButton) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButton)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonUrl)
	x.String(e.Text)
	x.String(e.Url)
	return x.buf
}

func (e TL_keyboardButtonCallback) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonCallback)
	x.String(e.Text)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_keyboardButtonRequestPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonRequestPhone)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonRequestGeoLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonRequestGeoLocation)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonSwitchInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonSwitchInline)
	x.Int(e.Flags)
	//flag SamePeer
	x.String(e.Text)
	x.String(e.Query)
	return x.buf
}

func (e TL_keyboardButtonGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonGame)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonBuy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonBuy)
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonUrlAuth)
	x.Int(e.Flags)
	x.String(e.Text)
	if e.Flags&1 != 0 {
		x.String(e.FwdText)
	}
	x.String(e.Url)
	x.Int(e.ButtonID)
	return x.buf
}

func (e TL_inputKeyboardButtonUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputKeyboardButtonUrlAuth)
	x.Int(e.Flags)
	//flag RequestWriteAccess
	x.String(e.Text)
	if e.Flags&2 != 0 {
		x.String(e.FwdText)
	}
	x.String(e.Url)
	x.Bytes(e.Bot.encode())
	return x.buf
}

func (e TL_keyboardButtonRequestPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonRequestPoll)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.Quiz.encode())
	}
	x.String(e.Text)
	return x.buf
}

func (e TL_keyboardButtonRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_keyboardButtonRow)
	x.Vector(e.Buttons)
	return x.buf
}

func (e TL_replyKeyboardHide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_replyKeyboardHide)
	x.Int(e.Flags)
	//flag Selective
	return x.buf
}

func (e TL_replyKeyboardForceReply) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_replyKeyboardForceReply)
	x.Int(e.Flags)
	//flag SingleUse
	//flag Selective
	return x.buf
}

func (e TL_replyKeyboardMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_replyKeyboardMarkup)
	x.Int(e.Flags)
	//flag Resize
	//flag SingleUse
	//flag Selective
	x.Vector(e.Rows)
	return x.buf
}

func (e TL_replyInlineMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_replyInlineMarkup)
	x.Vector(e.Rows)
	return x.buf
}

func (e TL_messageEntityUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityUnknown)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityMention) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityMention)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityHashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityHashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityBotCommand)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityEmail)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityBold)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityItalic)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityCode)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityPre) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityPre)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Language)
	return x.buf
}

func (e TL_messageEntityTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityTextUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Url)
	return x.buf
}

func (e TL_messageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_inputMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_messageEntityPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityPhone)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityCashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityCashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityUnderline)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityStrike)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityBlockquote)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_messageEntityBankCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageEntityBankCard)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e TL_inputChannelEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputChannelEmpty)
	return x.buf
}

func (e TL_inputChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputChannel)
	x.Int(e.ChannelID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputChannelFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputChannelFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ChannelID)
	return x.buf
}

func (e TL_contacts_resolvedPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_resolvedPeer)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_messageRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageRange)
	x.Int(e.MinID)
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_updates_channelDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_channelDifferenceEmpty)
	x.Int(e.Flags)
	//flag Final
	x.Int(e.Pts)
	if e.Flags&2 != 0 {
		x.Int(e.Timeout)
	}
	return x.buf
}

func (e TL_updates_channelDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_channelDifferenceTooLong)
	x.Int(e.Flags)
	//flag Final
	if e.Flags&2 != 0 {
		x.Int(e.Timeout)
	}
	x.Bytes(e.Dialog.encode())
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_updates_channelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_channelDifference)
	x.Int(e.Flags)
	//flag Final
	x.Int(e.Pts)
	if e.Flags&2 != 0 {
		x.Int(e.Timeout)
	}
	x.Vector(e.NewMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channelMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelMessagesFilterEmpty)
	return x.buf
}

func (e TL_channelMessagesFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelMessagesFilter)
	x.Int(e.Flags)
	//flag ExcludeNewMessages
	x.Vector(e.Ranges)
	return x.buf
}

func (e TL_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipant)
	x.Int(e.UserID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_channelParticipantSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantSelf)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_channelParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantCreator)
	x.Int(e.Flags)
	x.Int(e.UserID)
	if e.Flags&1 != 0 {
		x.String(e.Rank)
	}
	return x.buf
}

func (e TL_channelParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantAdmin)
	x.Int(e.Flags)
	//flag CanEdit
	//flag Self
	x.Int(e.UserID)
	if e.Flags&2 != 0 {
		x.Int(e.InviterID)
	}
	x.Int(e.PromotedBy)
	x.Int(e.Date)
	x.Bytes(e.AdminRights.encode())
	if e.Flags&4 != 0 {
		x.String(e.Rank)
	}
	return x.buf
}

func (e TL_channelParticipantBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantBanned)
	x.Int(e.Flags)
	//flag Left
	x.Int(e.UserID)
	x.Int(e.KickedBy)
	x.Int(e.Date)
	x.Bytes(e.BannedRights.encode())
	return x.buf
}

func (e TL_channelParticipantsRecent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsRecent)
	return x.buf
}

func (e TL_channelParticipantsAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsAdmins)
	return x.buf
}

func (e TL_channelParticipantsKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsKicked)
	x.String(e.Q)
	return x.buf
}

func (e TL_channelParticipantsBots) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsBots)
	return x.buf
}

func (e TL_channelParticipantsBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsBanned)
	x.String(e.Q)
	return x.buf
}

func (e TL_channelParticipantsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsSearch)
	x.String(e.Q)
	return x.buf
}

func (e TL_channelParticipantsContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelParticipantsContacts)
	x.String(e.Q)
	return x.buf
}

func (e TL_channels_channelParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_channelParticipants)
	x.Int(e.Count)
	x.Vector(e.Participants)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channels_channelParticipantsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_channelParticipantsNotModified)
	return x.buf
}

func (e TL_channels_channelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_channelParticipant)
	x.Bytes(e.Participant.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_help_termsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_termsOfService)
	x.Int(e.Flags)
	//flag Popup
	x.Bytes(e.ID.encode())
	x.String(e.Text)
	x.Vector(e.Entities)
	if e.Flags&2 != 0 {
		x.Int(e.MinAgeConfirm)
	}
	return x.buf
}

func (e TL_foundGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_foundGif)
	x.String(e.Url)
	x.String(e.ThumbUrl)
	x.String(e.ContentUrl)
	x.String(e.ContentType)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_foundGifCached) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_foundGifCached)
	x.String(e.Url)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Document.encode())
	return x.buf
}

func (e TL_messages_foundGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_foundGifs)
	x.Int(e.NextOffset)
	x.Vector(e.Results)
	return x.buf
}

func (e TL_messages_savedGifsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_savedGifsNotModified)
	return x.buf
}

func (e TL_messages_savedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_savedGifs)
	x.Int(e.Hash)
	x.Vector(e.Gifs)
	return x.buf
}

func (e TL_inputBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_inputBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageText)
	x.Int(e.Flags)
	//flag NoWebpage
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_inputBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.GeoPoint.encode())
	x.Int(e.Period)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_inputBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_inputBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_inputBotInlineMessageGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageGame)
	x.Int(e.Flags)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_inputBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineResult)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.String(e.Description)
	}
	if e.Flags&8 != 0 {
		x.String(e.Url)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.Content.encode())
	}
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e TL_inputBotInlineResultPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineResultPhoto)
	x.String(e.ID)
	x.String(e.Type)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e TL_inputBotInlineResultDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineResultDocument)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.String(e.Description)
	}
	x.Bytes(e.Document.encode())
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e TL_inputBotInlineResultGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineResultGame)
	x.String(e.ID)
	x.String(e.ShortName)
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e TL_botInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_botInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineMessageText)
	x.Int(e.Flags)
	//flag NoWebpage
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_botInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.Int(e.Period)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_botInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_botInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e TL_botInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineResult)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.String(e.Description)
	}
	if e.Flags&8 != 0 {
		x.String(e.Url)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.Content.encode())
	}
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e TL_botInlineMediaResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_botInlineMediaResult)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&4 != 0 {
		x.String(e.Title)
	}
	if e.Flags&8 != 0 {
		x.String(e.Description)
	}
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e TL_messages_botResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_botResults)
	x.Int(e.Flags)
	//flag Gallery
	x.Long(e.QueryID)
	if e.Flags&2 != 0 {
		x.String(e.NextOffset)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.SwitchPm.encode())
	}
	x.Vector(e.Results)
	x.Int(e.CacheTime)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_exportedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_exportedMessageLink)
	x.String(e.Link)
	x.String(e.Html)
	return x.buf
}

func (e TL_messageFwdHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageFwdHeader)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FromID)
	}
	if e.Flags&32 != 0 {
		x.String(e.FromName)
	}
	x.Int(e.Date)
	if e.Flags&2 != 0 {
		x.Int(e.ChannelID)
	}
	if e.Flags&4 != 0 {
		x.Int(e.ChannelPost)
	}
	if e.Flags&8 != 0 {
		x.String(e.PostAuthor)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.SavedFromPeer.encode())
	}
	if e.Flags&16 != 0 {
		x.Int(e.SavedFromMsgID)
	}
	if e.Flags&64 != 0 {
		x.String(e.PsaType)
	}
	return x.buf
}

func (e TL_auth_codeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_codeTypeSms)
	return x.buf
}

func (e TL_auth_codeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_codeTypeCall)
	return x.buf
}

func (e TL_auth_codeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_codeTypeFlashCall)
	return x.buf
}

func (e TL_auth_sentCodeTypeApp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_sentCodeTypeApp)
	x.Int(e.Length)
	return x.buf
}

func (e TL_auth_sentCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_sentCodeTypeSms)
	x.Int(e.Length)
	return x.buf
}

func (e TL_auth_sentCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_sentCodeTypeCall)
	x.Int(e.Length)
	return x.buf
}

func (e TL_auth_sentCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_sentCodeTypeFlashCall)
	x.String(e.Pattern)
	return x.buf
}

func (e TL_messages_botCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_botCallbackAnswer)
	x.Int(e.Flags)
	//flag Alert
	//flag HasUrl
	//flag NativeUi
	if e.Flags&1 != 0 {
		x.String(e.Message)
	}
	if e.Flags&4 != 0 {
		x.String(e.Url)
	}
	x.Int(e.CacheTime)
	return x.buf
}

func (e TL_messages_messageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_messageEditData)
	x.Int(e.Flags)
	//flag Caption
	return x.buf
}

func (e TL_inputBotInlineMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputBotInlineMessageID)
	x.Int(e.DcID)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inlineBotSwitchPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inlineBotSwitchPM)
	x.String(e.Text)
	x.String(e.StartParam)
	return x.buf
}

func (e TL_messages_peerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_peerDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e TL_topPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeer)
	x.Bytes(e.Peer.encode())
	x.Double(e.Rating)
	return x.buf
}

func (e TL_topPeerCategoryBotsPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryBotsPM)
	return x.buf
}

func (e TL_topPeerCategoryBotsInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryBotsInline)
	return x.buf
}

func (e TL_topPeerCategoryCorrespondents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryCorrespondents)
	return x.buf
}

func (e TL_topPeerCategoryGroups) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryGroups)
	return x.buf
}

func (e TL_topPeerCategoryChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryChannels)
	return x.buf
}

func (e TL_topPeerCategoryPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryPhoneCalls)
	return x.buf
}

func (e TL_topPeerCategoryForwardUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryForwardUsers)
	return x.buf
}

func (e TL_topPeerCategoryForwardChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryForwardChats)
	return x.buf
}

func (e TL_topPeerCategoryPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_topPeerCategoryPeers)
	x.Bytes(e.Category.encode())
	x.Int(e.Count)
	x.Vector(e.Peers)
	return x.buf
}

func (e TL_contacts_topPeersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_topPeersNotModified)
	return x.buf
}

func (e TL_contacts_topPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_topPeers)
	x.Vector(e.Categories)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_contacts_topPeersDisabled) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_topPeersDisabled)
	return x.buf
}

func (e TL_draftMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_draftMessageEmpty)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.Date)
	}
	return x.buf
}

func (e TL_draftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_draftMessage)
	x.Int(e.Flags)
	//flag NoWebpage
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.String(e.Message)
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	x.Int(e.Date)
	return x.buf
}

func (e TL_messages_featuredStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_featuredStickersNotModified)
	x.Int(e.Count)
	return x.buf
}

func (e TL_messages_featuredStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_featuredStickers)
	x.Int(e.Hash)
	x.Int(e.Count)
	x.Vector(e.Sets)
	x.VectorLong(e.Unread)
	return x.buf
}

func (e TL_messages_recentStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_recentStickersNotModified)
	return x.buf
}

func (e TL_messages_recentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_recentStickers)
	x.Int(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Stickers)
	x.VectorInt(e.Dates)
	return x.buf
}

func (e TL_messages_archivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_archivedStickers)
	x.Int(e.Count)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_stickerSetInstallResultSuccess)
	return x.buf
}

func (e TL_messages_stickerSetInstallResultArchive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_stickerSetInstallResultArchive)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_stickerSetCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickerSetCovered)
	x.Bytes(e.Set.encode())
	x.Bytes(e.Cover.encode())
	return x.buf
}

func (e TL_stickerSetMultiCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickerSetMultiCovered)
	x.Bytes(e.Set.encode())
	x.Vector(e.Covers)
	return x.buf
}

func (e TL_maskCoords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_maskCoords)
	x.Int(e.N)
	x.Double(e.X)
	x.Double(e.Y)
	x.Double(e.Zoom)
	return x.buf
}

func (e TL_inputStickeredMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickeredMediaPhoto)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_inputStickeredMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickeredMediaDocument)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_game) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_game)
	x.Int(e.Flags)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.String(e.ShortName)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Document.encode())
	}
	return x.buf
}

func (e TL_inputGameID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputGameID)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputGameShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputGameShortName)
	x.Bytes(e.BotID.encode())
	x.String(e.ShortName)
	return x.buf
}

func (e TL_highScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_highScore)
	x.Int(e.Pos)
	x.Int(e.UserID)
	x.Int(e.Score)
	return x.buf
}

func (e TL_messages_highScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_highScores)
	x.Vector(e.Scores)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_textEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textEmpty)
	return x.buf
}

func (e TL_textPlain) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textPlain)
	x.String(e.Text)
	return x.buf
}

func (e TL_textBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textBold)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textItalic)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textUnderline)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textStrike)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textFixed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textFixed)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textUrl)
	x.Bytes(e.Text.encode())
	x.String(e.Url)
	x.Long(e.WebpageID)
	return x.buf
}

func (e TL_textEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textEmail)
	x.Bytes(e.Text.encode())
	x.String(e.Email)
	return x.buf
}

func (e TL_textConcat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textConcat)
	x.Vector(e.Texts)
	return x.buf
}

func (e TL_textSubscript) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textSubscript)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textSuperscript) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textSuperscript)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textMarked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textMarked)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_textPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textPhone)
	x.Bytes(e.Text.encode())
	x.String(e.Phone)
	return x.buf
}

func (e TL_textImage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textImage)
	x.Long(e.DocumentID)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TL_textAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_textAnchor)
	x.Bytes(e.Text.encode())
	x.String(e.Name)
	return x.buf
}

func (e TL_pageBlockUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockUnsupported)
	return x.buf
}

func (e TL_pageBlockTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockTitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockSubtitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockSubtitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockAuthorDate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockAuthorDate)
	x.Bytes(e.Author.encode())
	x.Int(e.PublishedDate)
	return x.buf
}

func (e TL_pageBlockHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockHeader)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockSubheader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockSubheader)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockParagraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockParagraph)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockPreformatted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockPreformatted)
	x.Bytes(e.Text.encode())
	x.String(e.Language)
	return x.buf
}

func (e TL_pageBlockFooter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockFooter)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockDivider) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockDivider)
	return x.buf
}

func (e TL_pageBlockAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockAnchor)
	x.String(e.Name)
	return x.buf
}

func (e TL_pageBlockList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockList)
	x.Vector(e.Items)
	return x.buf
}

func (e TL_pageBlockBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockBlockquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockPullquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockPullquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockPhoto)
	x.Int(e.Flags)
	x.Long(e.PhotoID)
	x.Bytes(e.Caption.encode())
	if e.Flags&1 != 0 {
		x.String(e.Url)
	}
	if e.Flags&1 != 0 {
		x.Long(e.WebpageID)
	}
	return x.buf
}

func (e TL_pageBlockVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockVideo)
	x.Int(e.Flags)
	//flag Autoplay
	//flag Loop
	x.Long(e.VideoID)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockCover) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockCover)
	x.Bytes(e.Cover.encode())
	return x.buf
}

func (e TL_pageBlockEmbed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockEmbed)
	x.Int(e.Flags)
	//flag FullWidth
	//flag AllowScrolling
	if e.Flags&2 != 0 {
		x.String(e.Url)
	}
	if e.Flags&4 != 0 {
		x.String(e.Html)
	}
	if e.Flags&16 != 0 {
		x.Long(e.PosterPhotoID)
	}
	if e.Flags&32 != 0 {
		x.Int(e.W)
	}
	if e.Flags&32 != 0 {
		x.Int(e.H)
	}
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockEmbedPost) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockEmbedPost)
	x.String(e.Url)
	x.Long(e.WebpageID)
	x.Long(e.AuthorPhotoID)
	x.String(e.Author)
	x.Int(e.Date)
	x.Vector(e.Blocks)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockCollage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockCollage)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockSlideshow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockSlideshow)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_pageBlockAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockAudio)
	x.Long(e.AudioID)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_pageBlockKicker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockKicker)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageBlockTable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockTable)
	x.Int(e.Flags)
	//flag Bordered
	//flag Striped
	x.Bytes(e.Title.encode())
	x.Vector(e.Rows)
	return x.buf
}

func (e TL_pageBlockOrderedList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockOrderedList)
	x.Vector(e.Items)
	return x.buf
}

func (e TL_pageBlockDetails) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockDetails)
	x.Int(e.Flags)
	//flag Open
	x.Vector(e.Blocks)
	x.Bytes(e.Title.encode())
	return x.buf
}

func (e TL_pageBlockRelatedArticles) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockRelatedArticles)
	x.Bytes(e.Title.encode())
	x.Vector(e.Articles)
	return x.buf
}

func (e TL_pageBlockMap) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageBlockMap)
	x.Bytes(e.Geo.encode())
	x.Int(e.Zoom)
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e TL_phoneCallDiscardReasonMissed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallDiscardReasonMissed)
	return x.buf
}

func (e TL_phoneCallDiscardReasonDisconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallDiscardReasonDisconnect)
	return x.buf
}

func (e TL_phoneCallDiscardReasonHangup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallDiscardReasonHangup)
	return x.buf
}

func (e TL_phoneCallDiscardReasonBusy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallDiscardReasonBusy)
	return x.buf
}

func (e TL_dataJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dataJSON)
	x.String(e.Data)
	return x.buf
}

func (e TL_labeledPrice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_labeledPrice)
	x.String(e.Label)
	x.Long(e.Amount)
	return x.buf
}

func (e TL_invoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invoice)
	x.Int(e.Flags)
	//flag Test
	//flag NameRequested
	//flag PhoneRequested
	//flag EmailRequested
	//flag ShippingAddressRequested
	//flag Flexible
	//flag PhoneToProvider
	//flag EmailToProvider
	x.String(e.Currency)
	x.Vector(e.Prices)
	return x.buf
}

func (e TL_paymentCharge) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_paymentCharge)
	x.String(e.ID)
	x.String(e.ProviderChargeID)
	return x.buf
}

func (e TL_postAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_postAddress)
	x.String(e.StreetLine1)
	x.String(e.StreetLine2)
	x.String(e.City)
	x.String(e.State)
	x.String(e.CountryIso2)
	x.String(e.PostCode)
	return x.buf
}

func (e TL_paymentRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_paymentRequestedInfo)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.Name)
	}
	if e.Flags&2 != 0 {
		x.String(e.Phone)
	}
	if e.Flags&4 != 0 {
		x.String(e.Email)
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.ShippingAddress.encode())
	}
	return x.buf
}

func (e TL_paymentSavedCredentialsCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_paymentSavedCredentialsCard)
	x.String(e.ID)
	x.String(e.Title)
	return x.buf
}

func (e TL_webDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webDocument)
	x.String(e.Url)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TL_webDocumentNoProxy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webDocumentNoProxy)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TL_inputWebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputWebDocument)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e TL_inputWebFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputWebFileLocation)
	x.String(e.Url)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputWebFileGeoPointLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputWebFileGeoPointLocation)
	x.Bytes(e.GeoPoint.encode())
	x.Long(e.AccessHash)
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Zoom)
	x.Int(e.Scale)
	return x.buf
}

func (e TL_upload_webFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_webFile)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Bytes(e.FileType.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_payments_paymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_paymentForm)
	x.Int(e.Flags)
	//flag CanSaveCredentials
	//flag PasswordMissing
	x.Int(e.BotID)
	x.Bytes(e.Invoice.encode())
	x.Int(e.ProviderID)
	x.String(e.Url)
	if e.Flags&16 != 0 {
		x.String(e.NativeProvider)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.NativeParams.encode())
	}
	if e.Flags&1 != 0 {
		x.Bytes(e.SavedInfo.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.SavedCredentials.encode())
	}
	x.Vector(e.Users)
	return x.buf
}

func (e TL_payments_validatedRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_validatedRequestedInfo)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.ID)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.ShippingOptions)
	}
	return x.buf
}

func (e TL_payments_paymentResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_paymentResult)
	x.Bytes(e.Updates.encode())
	return x.buf
}

func (e TL_payments_paymentVerificationNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_paymentVerificationNeeded)
	x.String(e.Url)
	return x.buf
}

func (e TL_payments_paymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_paymentReceipt)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.BotID)
	x.Bytes(e.Invoice.encode())
	x.Int(e.ProviderID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Info.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Shipping.encode())
	}
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.String(e.CredentialsTitle)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_payments_savedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_savedInfo)
	x.Int(e.Flags)
	//flag HasSavedCredentials
	if e.Flags&1 != 0 {
		x.Bytes(e.SavedInfo.encode())
	}
	return x.buf
}

func (e TL_inputPaymentCredentialsSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPaymentCredentialsSaved)
	x.String(e.ID)
	x.StringBytes(e.TmpPassword)
	return x.buf
}

func (e TL_inputPaymentCredentials) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPaymentCredentials)
	x.Int(e.Flags)
	//flag Save
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_inputPaymentCredentialsApplePay) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPaymentCredentialsApplePay)
	x.Bytes(e.PaymentData.encode())
	return x.buf
}

func (e TL_inputPaymentCredentialsAndroidPay) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPaymentCredentialsAndroidPay)
	x.Bytes(e.PaymentToken.encode())
	x.String(e.GoogleTransactionID)
	return x.buf
}

func (e TL_account_tmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_tmpPassword)
	x.StringBytes(e.TmpPassword)
	x.Int(e.ValidUntil)
	return x.buf
}

func (e TL_shippingOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_shippingOption)
	x.String(e.ID)
	x.String(e.Title)
	x.Vector(e.Prices)
	return x.buf
}

func (e TL_inputStickerSetItem) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputStickerSetItem)
	x.Int(e.Flags)
	x.Bytes(e.Document.encode())
	x.String(e.Emoji)
	if e.Flags&1 != 0 {
		x.Bytes(e.MaskCoords.encode())
	}
	return x.buf
}

func (e TL_inputPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputPhoneCall)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_phoneCallEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e TL_phoneCallWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallWaiting)
	x.Int(e.Flags)
	//flag Video
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.Bytes(e.Protocol.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReceiveDate)
	}
	return x.buf
}

func (e TL_phoneCallRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallRequested)
	x.Int(e.Flags)
	//flag Video
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GAHash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phoneCallAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallAccepted)
	x.Int(e.Flags)
	//flag Video
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GB)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCall)
	x.Int(e.Flags)
	//flag P2pAllowed
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	x.Bytes(e.Protocol.encode())
	x.Vector(e.Connections)
	x.Int(e.StartDate)
	return x.buf
}

func (e TL_phoneCallDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallDiscarded)
	x.Int(e.Flags)
	//flag NeedRating
	//flag NeedDebug
	//flag Video
	x.Long(e.ID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Reason.encode())
	}
	if e.Flags&2 != 0 {
		x.Int(e.Duration)
	}
	return x.buf
}

func (e TL_phoneConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneConnection)
	x.Long(e.ID)
	x.String(e.Ip)
	x.String(e.Ipv6)
	x.Int(e.Port)
	x.StringBytes(e.PeerTag)
	return x.buf
}

func (e TL_phoneCallProtocol) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phoneCallProtocol)
	x.Int(e.Flags)
	//flag UdpP2p
	//flag UdpReflector
	x.Int(e.MinLayer)
	x.Int(e.MaxLayer)
	x.VectorString(e.LibraryVersions)
	return x.buf
}

func (e TL_phone_phoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_phoneCall)
	x.Bytes(e.PhoneCall.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_upload_cdnFileReuploadNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_cdnFileReuploadNeeded)
	x.StringBytes(e.RequestToken)
	return x.buf
}

func (e TL_upload_cdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_cdnFile)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_cdnPublicKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_cdnPublicKey)
	x.Int(e.DcID)
	x.String(e.PublicKey)
	return x.buf
}

func (e TL_cdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_cdnConfig)
	x.Vector(e.PublicKeys)
	return x.buf
}

func (e TL_langPackString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langPackString)
	x.String(e.Key)
	x.String(e.Value)
	return x.buf
}

func (e TL_langPackStringPluralized) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langPackStringPluralized)
	x.Int(e.Flags)
	x.String(e.Key)
	if e.Flags&1 != 0 {
		x.String(e.ZeroValue)
	}
	if e.Flags&2 != 0 {
		x.String(e.OneValue)
	}
	if e.Flags&4 != 0 {
		x.String(e.TwoValue)
	}
	if e.Flags&8 != 0 {
		x.String(e.FewValue)
	}
	if e.Flags&16 != 0 {
		x.String(e.ManyValue)
	}
	x.String(e.OtherValue)
	return x.buf
}

func (e TL_langPackStringDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langPackStringDeleted)
	x.String(e.Key)
	return x.buf
}

func (e TL_langPackDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langPackDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	x.Int(e.Version)
	x.Vector(e.Strings)
	return x.buf
}

func (e TL_langPackLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langPackLanguage)
	x.Int(e.Flags)
	//flag Official
	//flag Rtl
	//flag Beta
	x.String(e.Name)
	x.String(e.NativeName)
	x.String(e.LangCode)
	if e.Flags&2 != 0 {
		x.String(e.BaseLangCode)
	}
	x.String(e.PluralCode)
	x.Int(e.StringsCount)
	x.Int(e.TranslatedCount)
	x.String(e.TranslationsUrl)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangeTitle)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangeAbout)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangeUsername)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangePhoto)
	x.Bytes(e.PrevPhoto.encode())
	x.Bytes(e.NewPhoto.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionToggleInvites)
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionToggleSignatures)
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionUpdatePinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionUpdatePinned)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionEditMessage)
	x.Bytes(e.PrevMessage.encode())
	x.Bytes(e.NewMessage.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionDeleteMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionDeleteMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantJoin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionParticipantJoin)
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantLeave) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionParticipantLeave)
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionParticipantInvite)
	x.Bytes(e.Participant.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantToggleBan) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionParticipantToggleBan)
	x.Bytes(e.PrevParticipant.encode())
	x.Bytes(e.NewParticipant.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionParticipantToggleAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionParticipantToggleAdmin)
	x.Bytes(e.PrevParticipant.encode())
	x.Bytes(e.NewParticipant.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangeStickerSet)
	x.Bytes(e.PrevStickerset.encode())
	x.Bytes(e.NewStickerset.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionTogglePreHistoryHidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionTogglePreHistoryHidden)
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionDefaultBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionDefaultBannedRights)
	x.Bytes(e.PrevBannedRights.encode())
	x.Bytes(e.NewBannedRights.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionStopPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionStopPoll)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeLinkedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangeLinkedChat)
	x.Int(e.PrevValue)
	x.Int(e.NewValue)
	return x.buf
}

func (e TL_channelAdminLogEventActionChangeLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionChangeLocation)
	x.Bytes(e.PrevValue.encode())
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e TL_channelAdminLogEventActionToggleSlowMode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventActionToggleSlowMode)
	x.Int(e.PrevValue)
	x.Int(e.NewValue)
	return x.buf
}

func (e TL_channelAdminLogEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEvent)
	x.Long(e.ID)
	x.Int(e.Date)
	x.Int(e.UserID)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_channels_adminLogResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_adminLogResults)
	x.Vector(e.Events)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channelAdminLogEventsFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelAdminLogEventsFilter)
	x.Int(e.Flags)
	//flag Join
	//flag Leave
	//flag Invite
	//flag Ban
	//flag Unban
	//flag Kick
	//flag Unkick
	//flag Promote
	//flag Demote
	//flag Info
	//flag Settings
	//flag Pinned
	//flag Edit
	//flag Delete
	return x.buf
}

func (e TL_popularContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_popularContact)
	x.Long(e.ClientID)
	x.Int(e.Importers)
	return x.buf
}

func (e TL_messages_favedStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_favedStickersNotModified)
	return x.buf
}

func (e TL_messages_favedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_favedStickers)
	x.Int(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_recentMeUrlUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_recentMeUrlUnknown)
	x.String(e.Url)
	return x.buf
}

func (e TL_recentMeUrlUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_recentMeUrlUser)
	x.String(e.Url)
	x.Int(e.UserID)
	return x.buf
}

func (e TL_recentMeUrlChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_recentMeUrlChat)
	x.String(e.Url)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_recentMeUrlChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_recentMeUrlChatInvite)
	x.String(e.Url)
	x.Bytes(e.ChatInvite.encode())
	return x.buf
}

func (e TL_recentMeUrlStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_recentMeUrlStickerSet)
	x.String(e.Url)
	x.Bytes(e.Set.encode())
	return x.buf
}

func (e TL_help_recentMeUrls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_recentMeUrls)
	x.Vector(e.Urls)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputSingleMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputSingleMedia)
	x.Int(e.Flags)
	x.Bytes(e.Media.encode())
	x.Long(e.RandomID)
	x.String(e.Message)
	if e.Flags&1 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_webAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webAuthorization)
	x.Long(e.Hash)
	x.Int(e.BotID)
	x.String(e.Domain)
	x.String(e.Browser)
	x.String(e.Platform)
	x.Int(e.DateCreated)
	x.Int(e.DateActive)
	x.String(e.Ip)
	x.String(e.Region)
	return x.buf
}

func (e TL_account_webAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_webAuthorizations)
	x.Vector(e.Authorizations)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_inputMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessageID)
	x.Int(e.ID)
	return x.buf
}

func (e TL_inputMessageReplyTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessageReplyTo)
	x.Int(e.ID)
	return x.buf
}

func (e TL_inputMessagePinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputMessagePinned)
	return x.buf
}

func (e TL_inputDialogPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputDialogPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_inputDialogPeerFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputDialogPeerFolder)
	x.Int(e.FolderID)
	return x.buf
}

func (e TL_dialogPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dialogPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_dialogPeerFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dialogPeerFolder)
	x.Int(e.FolderID)
	return x.buf
}

func (e TL_messages_foundStickerSetsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_foundStickerSetsNotModified)
	return x.buf
}

func (e TL_messages_foundStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_foundStickerSets)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.buf
}

func (e TL_fileHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_fileHash)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.StringBytes(e.Hash)
	return x.buf
}

func (e TL_inputClientProxy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputClientProxy)
	x.String(e.Address)
	x.Int(e.Port)
	return x.buf
}

func (e TL_help_termsOfServiceUpdateEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_termsOfServiceUpdateEmpty)
	x.Int(e.Expires)
	return x.buf
}

func (e TL_help_termsOfServiceUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_termsOfServiceUpdate)
	x.Int(e.Expires)
	x.Bytes(e.TermsOfService.encode())
	return x.buf
}

func (e TL_inputSecureFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputSecureFileUploaded)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Md5Checksum)
	x.StringBytes(e.FileHash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e TL_inputSecureFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputSecureFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_secureFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureFileEmpty)
	return x.buf
}

func (e TL_secureFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.Int(e.DcID)
	x.Int(e.Date)
	x.StringBytes(e.FileHash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e TL_secureData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureData)
	x.StringBytes(e.Data)
	x.StringBytes(e.DataHash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e TL_securePlainPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_securePlainPhone)
	x.String(e.Phone)
	return x.buf
}

func (e TL_securePlainEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_securePlainEmail)
	x.String(e.Email)
	return x.buf
}

func (e TL_secureValueTypePersonalDetails) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypePersonalDetails)
	return x.buf
}

func (e TL_secureValueTypePassport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypePassport)
	return x.buf
}

func (e TL_secureValueTypeDriverLicense) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeDriverLicense)
	return x.buf
}

func (e TL_secureValueTypeIdentityCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeIdentityCard)
	return x.buf
}

func (e TL_secureValueTypeInternalPassport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeInternalPassport)
	return x.buf
}

func (e TL_secureValueTypeAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeAddress)
	return x.buf
}

func (e TL_secureValueTypeUtilityBill) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeUtilityBill)
	return x.buf
}

func (e TL_secureValueTypeBankStatement) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeBankStatement)
	return x.buf
}

func (e TL_secureValueTypeRentalAgreement) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeRentalAgreement)
	return x.buf
}

func (e TL_secureValueTypePassportRegistration) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypePassportRegistration)
	return x.buf
}

func (e TL_secureValueTypeTemporaryRegistration) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeTemporaryRegistration)
	return x.buf
}

func (e TL_secureValueTypePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypePhone)
	return x.buf
}

func (e TL_secureValueTypeEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueTypeEmail)
	return x.buf
}

func (e TL_secureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValue)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Data.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.FrontSide.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReverseSide.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Selfie.encode())
	}
	if e.Flags&64 != 0 {
		x.Vector(e.Translation)
	}
	if e.Flags&16 != 0 {
		x.Vector(e.Files)
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.PlainData.encode())
	}
	x.StringBytes(e.Hash)
	return x.buf
}

func (e TL_inputSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputSecureValue)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Data.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.FrontSide.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReverseSide.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Selfie.encode())
	}
	if e.Flags&64 != 0 {
		x.Vector(e.Translation)
	}
	if e.Flags&16 != 0 {
		x.Vector(e.Files)
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.PlainData.encode())
	}
	return x.buf
}

func (e TL_secureValueHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueHash)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.Hash)
	return x.buf
}

func (e TL_secureValueErrorData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorData)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.DataHash)
	x.String(e.Field)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorFrontSide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorFrontSide)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorReverseSide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorReverseSide)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorSelfie) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorSelfie)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorFile)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorFiles) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorFiles)
	x.Bytes(e.Type.encode())
	x.Vector(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueError)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.Hash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorTranslationFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorTranslationFile)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureValueErrorTranslationFiles) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureValueErrorTranslationFiles)
	x.Bytes(e.Type.encode())
	x.Vector(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e TL_secureCredentialsEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureCredentialsEncrypted)
	x.StringBytes(e.Data)
	x.StringBytes(e.Hash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e TL_account_authorizationForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_authorizationForm)
	x.Int(e.Flags)
	x.Vector(e.RequiredTypes)
	x.Vector(e.Values)
	x.Vector(e.Errors)
	x.Vector(e.Users)
	if e.Flags&1 != 0 {
		x.String(e.PrivacyPolicyUrl)
	}
	return x.buf
}

func (e TL_account_sentEmailCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_sentEmailCode)
	x.String(e.EmailPattern)
	x.Int(e.Length)
	return x.buf
}

func (e TL_help_deepLinkInfoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_deepLinkInfoEmpty)
	return x.buf
}

func (e TL_help_deepLinkInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_deepLinkInfo)
	x.Int(e.Flags)
	//flag UpdateApp
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_savedPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_savedPhoneContact)
	x.String(e.Phone)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.Int(e.Date)
	return x.buf
}

func (e TL_account_takeout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_takeout)
	x.Long(e.ID)
	return x.buf
}

func (e TL_passwordKdfAlgoUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_passwordKdfAlgoUnknown)
	return x.buf
}

func (e TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow)
	x.StringBytes(e.Salt1)
	x.StringBytes(e.Salt2)
	x.Int(e.G)
	x.StringBytes(e.P)
	return x.buf
}

func (e TL_securePasswordKdfAlgoUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_securePasswordKdfAlgoUnknown)
	return x.buf
}

func (e TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000)
	x.StringBytes(e.Salt)
	return x.buf
}

func (e TL_securePasswordKdfAlgoSHA512) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_securePasswordKdfAlgoSHA512)
	x.StringBytes(e.Salt)
	return x.buf
}

func (e TL_secureSecretSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureSecretSettings)
	x.Bytes(e.SecureAlgo.encode())
	x.StringBytes(e.SecureSecret)
	x.Long(e.SecureSecretID)
	return x.buf
}

func (e TL_inputCheckPasswordEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputCheckPasswordEmpty)
	return x.buf
}

func (e TL_inputCheckPasswordSRP) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputCheckPasswordSRP)
	x.Long(e.SrpID)
	x.StringBytes(e.A)
	x.StringBytes(e.M1)
	return x.buf
}

func (e TL_secureRequiredType) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureRequiredType)
	x.Int(e.Flags)
	//flag NativeNames
	//flag SelfieRequired
	//flag TranslationRequired
	x.Bytes(e.Type.encode())
	return x.buf
}

func (e TL_secureRequiredTypeOneOf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_secureRequiredTypeOneOf)
	x.Vector(e.Types)
	return x.buf
}

func (e TL_help_passportConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_passportConfigNotModified)
	return x.buf
}

func (e TL_help_passportConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_passportConfig)
	x.Int(e.Hash)
	x.Bytes(e.CountriesLangs.encode())
	return x.buf
}

func (e TL_inputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputAppEvent)
	x.Double(e.Time)
	x.String(e.Type)
	x.Long(e.Peer)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_jsonObjectValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonObjectValue)
	x.String(e.Key)
	x.Bytes(e.Value.encode())
	return x.buf
}

func (e TL_jsonNull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonNull)
	return x.buf
}

func (e TL_jsonBool) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonBool)
	x.Bytes(e.Value.encode())
	return x.buf
}

func (e TL_jsonNumber) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonNumber)
	x.Double(e.Value)
	return x.buf
}

func (e TL_jsonString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonString)
	x.String(e.Value)
	return x.buf
}

func (e TL_jsonArray) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonArray)
	x.Vector(e.Value)
	return x.buf
}

func (e TL_jsonObject) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_jsonObject)
	x.Vector(e.Value)
	return x.buf
}

func (e TL_pageTableCell) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageTableCell)
	x.Int(e.Flags)
	//flag Header
	//flag AlignCenter
	//flag AlignRight
	//flag ValignMiddle
	//flag ValignBottom
	if e.Flags&128 != 0 {
		x.Bytes(e.Text.encode())
	}
	if e.Flags&2 != 0 {
		x.Int(e.Colspan)
	}
	if e.Flags&4 != 0 {
		x.Int(e.Rowspan)
	}
	return x.buf
}

func (e TL_pageTableRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageTableRow)
	x.Vector(e.Cells)
	return x.buf
}

func (e TL_pageCaption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageCaption)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Credit.encode())
	return x.buf
}

func (e TL_pageListItemText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageListItemText)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageListItemBlocks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageListItemBlocks)
	x.Vector(e.Blocks)
	return x.buf
}

func (e TL_pageListOrderedItemText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageListOrderedItemText)
	x.String(e.Num)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TL_pageListOrderedItemBlocks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageListOrderedItemBlocks)
	x.String(e.Num)
	x.Vector(e.Blocks)
	return x.buf
}

func (e TL_pageRelatedArticle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pageRelatedArticle)
	x.Int(e.Flags)
	x.String(e.Url)
	x.Long(e.WebpageID)
	if e.Flags&1 != 0 {
		x.String(e.Title)
	}
	if e.Flags&2 != 0 {
		x.String(e.Description)
	}
	if e.Flags&4 != 0 {
		x.Long(e.PhotoID)
	}
	if e.Flags&8 != 0 {
		x.String(e.Author)
	}
	if e.Flags&16 != 0 {
		x.Int(e.PublishedDate)
	}
	return x.buf
}

func (e TL_page) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_page)
	x.Int(e.Flags)
	//flag Part
	//flag Rtl
	//flag V2
	x.String(e.Url)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Documents)
	if e.Flags&8 != 0 {
		x.Int(e.Views)
	}
	return x.buf
}

func (e TL_help_supportName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_supportName)
	x.String(e.Name)
	return x.buf
}

func (e TL_help_userInfoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_userInfoEmpty)
	return x.buf
}

func (e TL_help_userInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_userInfo)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.String(e.Author)
	x.Int(e.Date)
	return x.buf
}

func (e TL_pollAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pollAnswer)
	x.String(e.Text)
	x.StringBytes(e.Option)
	return x.buf
}

func (e TL_poll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_poll)
	x.Long(e.ID)
	x.Int(e.Flags)
	//flag Closed
	//flag PublicVoters
	//flag MultipleChoice
	//flag Quiz
	x.String(e.Question)
	x.Vector(e.Answers)
	if e.Flags&16 != 0 {
		x.Int(e.ClosePeriod)
	}
	if e.Flags&32 != 0 {
		x.Int(e.CloseDate)
	}
	return x.buf
}

func (e TL_pollAnswerVoters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pollAnswerVoters)
	x.Int(e.Flags)
	//flag Chosen
	//flag Correct
	x.StringBytes(e.Option)
	x.Int(e.Voters)
	return x.buf
}

func (e TL_pollResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_pollResults)
	x.Int(e.Flags)
	//flag Min
	if e.Flags&2 != 0 {
		x.Vector(e.Results)
	}
	if e.Flags&4 != 0 {
		x.Int(e.TotalVoters)
	}
	if e.Flags&8 != 0 {
		x.VectorInt(e.RecentVoters)
	}
	if e.Flags&16 != 0 {
		x.String(e.Solution)
	}
	if e.Flags&16 != 0 {
		x.Vector(e.SolutionEntities)
	}
	return x.buf
}

func (e TL_chatOnlines) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatOnlines)
	x.Int(e.Onlines)
	return x.buf
}

func (e TL_statsURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsURL)
	x.String(e.Url)
	return x.buf
}

func (e TL_chatAdminRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatAdminRights)
	x.Int(e.Flags)
	//flag ChangeInfo
	//flag PostMessages
	//flag EditMessages
	//flag DeleteMessages
	//flag BanUsers
	//flag InviteUsers
	//flag PinMessages
	//flag AddAdmins
	return x.buf
}

func (e TL_chatBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_chatBannedRights)
	x.Int(e.Flags)
	//flag ViewMessages
	//flag SendMessages
	//flag SendMedia
	//flag SendStickers
	//flag SendGifs
	//flag SendGames
	//flag SendInline
	//flag EmbedLinks
	//flag SendPolls
	//flag ChangeInfo
	//flag InviteUsers
	//flag PinMessages
	x.Int(e.UntilDate)
	return x.buf
}

func (e TL_inputWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputWallPaper)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputWallPaperSlug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputWallPaperSlug)
	x.String(e.Slug)
	return x.buf
}

func (e TL_inputWallPaperNoFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputWallPaperNoFile)
	return x.buf
}

func (e TL_account_wallPapersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_wallPapersNotModified)
	return x.buf
}

func (e TL_account_wallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_wallPapers)
	x.Int(e.Hash)
	x.Vector(e.Wallpapers)
	return x.buf
}

func (e TL_codeSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_codeSettings)
	x.Int(e.Flags)
	//flag AllowFlashcall
	//flag CurrentNumber
	//flag AllowAppHash
	return x.buf
}

func (e TL_wallPaperSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_wallPaperSettings)
	x.Int(e.Flags)
	//flag Blur
	//flag Motion
	if e.Flags&1 != 0 {
		x.Int(e.BackgroundColor)
	}
	if e.Flags&16 != 0 {
		x.Int(e.SecondBackgroundColor)
	}
	if e.Flags&8 != 0 {
		x.Int(e.Intensity)
	}
	if e.Flags&16 != 0 {
		x.Int(e.Rotation)
	}
	return x.buf
}

func (e TL_autoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_autoDownloadSettings)
	x.Int(e.Flags)
	//flag Disabled
	//flag VideoPreloadLarge
	//flag AudioPreloadNext
	//flag PhonecallsLessData
	x.Int(e.PhotoSizeMax)
	x.Int(e.VideoSizeMax)
	x.Int(e.FileSizeMax)
	x.Int(e.VideoUploadMaxbitrate)
	return x.buf
}

func (e TL_account_autoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_autoDownloadSettings)
	x.Bytes(e.Low.encode())
	x.Bytes(e.Medium.encode())
	x.Bytes(e.High.encode())
	return x.buf
}

func (e TL_emojiKeyword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_emojiKeyword)
	x.String(e.Keyword)
	x.VectorString(e.Emoticons)
	return x.buf
}

func (e TL_emojiKeywordDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_emojiKeywordDeleted)
	x.String(e.Keyword)
	x.VectorString(e.Emoticons)
	return x.buf
}

func (e TL_emojiKeywordsDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_emojiKeywordsDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	x.Int(e.Version)
	x.Vector(e.Keywords)
	return x.buf
}

func (e TL_emojiURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_emojiURL)
	x.String(e.Url)
	return x.buf
}

func (e TL_emojiLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_emojiLanguage)
	x.String(e.LangCode)
	return x.buf
}

func (e TL_fileLocationToBeDeprecated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_fileLocationToBeDeprecated)
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	return x.buf
}

func (e TL_folder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_folder)
	x.Int(e.Flags)
	//flag AutofillNewBroadcasts
	//flag AutofillPublicGroups
	//flag AutofillNewCorrespondents
	x.Int(e.ID)
	x.String(e.Title)
	if e.Flags&8 != 0 {
		x.Bytes(e.Photo.encode())
	}
	return x.buf
}

func (e TL_inputFolderPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputFolderPeer)
	x.Bytes(e.Peer.encode())
	x.Int(e.FolderID)
	return x.buf
}

func (e TL_folderPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_folderPeer)
	x.Bytes(e.Peer.encode())
	x.Int(e.FolderID)
	return x.buf
}

func (e TL_messages_searchCounter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_searchCounter)
	x.Int(e.Flags)
	//flag Inexact
	x.Bytes(e.Filter.encode())
	x.Int(e.Count)
	return x.buf
}

func (e TL_urlAuthResultRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_urlAuthResultRequest)
	x.Int(e.Flags)
	//flag RequestWriteAccess
	x.Bytes(e.Bot.encode())
	x.String(e.Domain)
	return x.buf
}

func (e TL_urlAuthResultAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_urlAuthResultAccepted)
	x.String(e.Url)
	return x.buf
}

func (e TL_urlAuthResultDefault) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_urlAuthResultDefault)
	return x.buf
}

func (e TL_channelLocationEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelLocationEmpty)
	return x.buf
}

func (e TL_channelLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channelLocation)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Address)
	return x.buf
}

func (e TL_peerLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerLocated)
	x.Bytes(e.Peer.encode())
	x.Int(e.Expires)
	x.Int(e.Distance)
	return x.buf
}

func (e TL_peerSelfLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_peerSelfLocated)
	x.Int(e.Expires)
	return x.buf
}

func (e TL_restrictionReason) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_restrictionReason)
	x.String(e.Platform)
	x.String(e.Reason)
	x.String(e.Text)
	return x.buf
}

func (e TL_inputTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputTheme)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e TL_inputThemeSlug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputThemeSlug)
	x.String(e.Slug)
	return x.buf
}

func (e TL_theme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_theme)
	x.Int(e.Flags)
	//flag Creator
	//flag Default
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.String(e.Slug)
	x.String(e.Title)
	if e.Flags&4 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Settings.encode())
	}
	x.Int(e.InstallsCount)
	return x.buf
}

func (e TL_account_themesNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_themesNotModified)
	return x.buf
}

func (e TL_account_themes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_themes)
	x.Int(e.Hash)
	x.Vector(e.Themes)
	return x.buf
}

func (e TL_auth_loginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_loginToken)
	x.Int(e.Expires)
	x.StringBytes(e.Token)
	return x.buf
}

func (e TL_auth_loginTokenMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_loginTokenMigrateTo)
	x.Int(e.DcID)
	x.StringBytes(e.Token)
	return x.buf
}

func (e TL_auth_loginTokenSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_loginTokenSuccess)
	x.Bytes(e.Authorization.encode())
	return x.buf
}

func (e TL_account_contentSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_contentSettings)
	x.Int(e.Flags)
	//flag SensitiveEnabled
	//flag SensitiveCanChange
	return x.buf
}

func (e TL_messages_inactiveChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_inactiveChats)
	x.VectorInt(e.Dates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e TL_baseThemeClassic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_baseThemeClassic)
	return x.buf
}

func (e TL_baseThemeDay) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_baseThemeDay)
	return x.buf
}

func (e TL_baseThemeNight) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_baseThemeNight)
	return x.buf
}

func (e TL_baseThemeTinted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_baseThemeTinted)
	return x.buf
}

func (e TL_baseThemeArctic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_baseThemeArctic)
	return x.buf
}

func (e TL_inputThemeSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_inputThemeSettings)
	x.Int(e.Flags)
	x.Bytes(e.BaseTheme.encode())
	x.Int(e.AccentColor)
	if e.Flags&1 != 0 {
		x.Int(e.MessageTopColor)
	}
	if e.Flags&1 != 0 {
		x.Int(e.MessageBottomColor)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Wallpaper.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.WallpaperSettings.encode())
	}
	return x.buf
}

func (e TL_themeSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_themeSettings)
	x.Int(e.Flags)
	x.Bytes(e.BaseTheme.encode())
	x.Int(e.AccentColor)
	if e.Flags&1 != 0 {
		x.Int(e.MessageTopColor)
	}
	if e.Flags&1 != 0 {
		x.Int(e.MessageBottomColor)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Wallpaper.encode())
	}
	return x.buf
}

func (e TL_webPageAttributeTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_webPageAttributeTheme)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Vector(e.Documents)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e TL_messageUserVote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageUserVote)
	x.Int(e.UserID)
	x.StringBytes(e.Option)
	x.Int(e.Date)
	return x.buf
}

func (e TL_messageUserVoteInputOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageUserVoteInputOption)
	x.Int(e.UserID)
	x.Int(e.Date)
	return x.buf
}

func (e TL_messageUserVoteMultiple) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageUserVoteMultiple)
	x.Int(e.UserID)
	x.Vector(e.Options)
	x.Int(e.Date)
	return x.buf
}

func (e TL_messages_votesList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_votesList)
	x.Int(e.Flags)
	x.Int(e.Count)
	x.Vector(e.Votes)
	x.Vector(e.Users)
	if e.Flags&1 != 0 {
		x.String(e.NextOffset)
	}
	return x.buf
}

func (e TL_bankCardOpenUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bankCardOpenUrl)
	x.String(e.Url)
	x.String(e.Name)
	return x.buf
}

func (e TL_payments_bankCardData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_bankCardData)
	x.String(e.Title)
	x.Vector(e.OpenUrls)
	return x.buf
}

func (e TL_dialogFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dialogFilter)
	x.Int(e.Flags)
	//flag Contacts
	//flag NonContacts
	//flag Groups
	//flag Broadcasts
	//flag Bots
	//flag ExcludeMuted
	//flag ExcludeRead
	//flag ExcludeArchived
	x.Int(e.ID)
	x.String(e.Title)
	if e.Flags&33554432 != 0 {
		x.String(e.Emoticon)
	}
	x.Vector(e.PinnedPeers)
	x.Vector(e.IncludePeers)
	x.Vector(e.ExcludePeers)
	return x.buf
}

func (e TL_dialogFilterSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_dialogFilterSuggested)
	x.Bytes(e.Filter.encode())
	x.String(e.Description)
	return x.buf
}

func (e TL_statsDateRangeDays) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsDateRangeDays)
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	return x.buf
}

func (e TL_statsAbsValueAndPrev) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsAbsValueAndPrev)
	x.Double(e.Current)
	x.Double(e.Previous)
	return x.buf
}

func (e TL_statsPercentValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsPercentValue)
	x.Double(e.Part)
	x.Double(e.Total)
	return x.buf
}

func (e TL_statsGraphAsync) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsGraphAsync)
	x.String(e.Token)
	return x.buf
}

func (e TL_statsGraphError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsGraphError)
	x.String(e.Error)
	return x.buf
}

func (e TL_statsGraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_statsGraph)
	x.Int(e.Flags)
	x.Bytes(e.Json.encode())
	if e.Flags&1 != 0 {
		x.String(e.ZoomToken)
	}
	return x.buf
}

func (e TL_messageInteractionCounters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messageInteractionCounters)
	x.Int(e.MsgID)
	x.Int(e.Views)
	x.Int(e.Forwards)
	return x.buf
}

func (e TL_stats_broadcastStats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stats_broadcastStats)
	x.Bytes(e.Period.encode())
	x.Bytes(e.Followers.encode())
	x.Bytes(e.ViewsPerPost.encode())
	x.Bytes(e.SharesPerPost.encode())
	x.Bytes(e.EnabledNotifications.encode())
	x.Bytes(e.GrowthGraph.encode())
	x.Bytes(e.FollowersGraph.encode())
	x.Bytes(e.MuteGraph.encode())
	x.Bytes(e.TopHoursGraph.encode())
	x.Bytes(e.InteractionsGraph.encode())
	x.Bytes(e.IvInteractionsGraph.encode())
	x.Bytes(e.ViewsBySourceGraph.encode())
	x.Bytes(e.NewFollowersBySourceGraph.encode())
	x.Bytes(e.LanguagesGraph.encode())
	x.Vector(e.RecentMessageInteractions)
	return x.buf
}

func (e TL_help_promoDataEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_promoDataEmpty)
	x.Int(e.Expires)
	return x.buf
}

func (e TL_help_promoData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_promoData)
	x.Int(e.Flags)
	//flag Proxy
	x.Int(e.Expires)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	if e.Flags&2 != 0 {
		x.String(e.PsaType)
	}
	if e.Flags&4 != 0 {
		x.String(e.PsaMessage)
	}
	return x.buf
}

func (e TL_videoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_videoSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}

func (e TL_invokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invokeAfterMsg)
	x.Long(e.MsgID)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_invokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invokeAfterMsgs)
	x.VectorLong(e.MsgIds)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_initConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_initConnection)
	x.Int(e.Flags)
	x.Int(e.ApiID)
	x.String(e.DeviceModel)
	x.String(e.SystemVersion)
	x.String(e.AppVersion)
	x.String(e.SystemLangCode)
	x.String(e.LangPack)
	x.String(e.LangCode)
	if e.Flags&1 != 0 {
		x.Bytes(e.Proxy.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Params.encode())
	}
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_invokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invokeWithLayer)
	x.Int(e.Layer)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_invokeWithoutUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invokeWithoutUpdates)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_invokeWithMessagesRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invokeWithMessagesRange)
	x.Bytes(e.Range.encode())
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_invokeWithTakeout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_invokeWithTakeout)
	x.Long(e.TakeoutID)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e TL_auth_sendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_sendCode)
	x.String(e.PhoneNumber)
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_auth_signUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_signUp)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e TL_auth_signIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_signIn)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e TL_auth_logOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_logOut)
	return x.buf
}

func (e TL_auth_resetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_resetAuthorizations)
	return x.buf
}

func (e TL_auth_exportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_exportAuthorization)
	x.Int(e.DcID)
	return x.buf
}

func (e TL_auth_importAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_importAuthorization)
	x.Int(e.ID)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_auth_bindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_bindTempAuthKey)
	x.Long(e.PermAuthKeyID)
	x.Long(e.Nonce)
	x.Int(e.ExpiresAt)
	x.StringBytes(e.EncryptedMessage)
	return x.buf
}

func (e TL_auth_importBotAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_importBotAuthorization)
	x.Int(e.Flags)
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	x.String(e.BotAuthToken)
	return x.buf
}

func (e TL_auth_checkPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_checkPassword)
	x.Bytes(e.Password.encode())
	return x.buf
}

func (e TL_auth_requestPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_requestPasswordRecovery)
	return x.buf
}

func (e TL_auth_recoverPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_recoverPassword)
	x.String(e.Code)
	return x.buf
}

func (e TL_auth_resendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_resendCode)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}

func (e TL_auth_cancelCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_cancelCode)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}

func (e TL_auth_dropTempAuthKeys) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_dropTempAuthKeys)
	x.VectorLong(e.ExceptAuthKeys)
	return x.buf
}

func (e TL_auth_exportLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_exportLoginToken)
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	x.VectorInt(e.ExceptIds)
	return x.buf
}

func (e TL_auth_importLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_importLoginToken)
	x.StringBytes(e.Token)
	return x.buf
}

func (e TL_auth_acceptLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_auth_acceptLoginToken)
	x.StringBytes(e.Token)
	return x.buf
}

func (e TL_account_registerDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_registerDevice)
	x.Int(e.Flags)
	//flag NoMuted
	x.Int(e.TokenType)
	x.String(e.Token)
	x.Bytes(e.AppSandbox.encode())
	x.StringBytes(e.Secret)
	x.VectorInt(e.OtherUids)
	return x.buf
}

func (e TL_account_unregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_unregisterDevice)
	x.Int(e.TokenType)
	x.String(e.Token)
	x.VectorInt(e.OtherUids)
	return x.buf
}

func (e TL_account_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_getNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getNotifySettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_account_resetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_resetNotifySettings)
	return x.buf
}

func (e TL_account_updateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updateProfile)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.FirstName)
	}
	if e.Flags&2 != 0 {
		x.String(e.LastName)
	}
	if e.Flags&4 != 0 {
		x.String(e.About)
	}
	return x.buf
}

func (e TL_account_updateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updateStatus)
	x.Bytes(e.Offline.encode())
	return x.buf
}

func (e TL_account_getWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getWallPapers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_account_reportPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_reportPeer)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Reason.encode())
	return x.buf
}

func (e TL_account_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_checkUsername)
	x.String(e.Username)
	return x.buf
}

func (e TL_account_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updateUsername)
	x.String(e.Username)
	return x.buf
}

func (e TL_account_getPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getPrivacy)
	x.Bytes(e.Key.encode())
	return x.buf
}

func (e TL_account_setPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_setPrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e TL_account_deleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_deleteAccount)
	x.String(e.Reason)
	return x.buf
}

func (e TL_account_getAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getAccountTTL)
	return x.buf
}

func (e TL_account_setAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_setAccountTTL)
	x.Bytes(e.Ttl.encode())
	return x.buf
}

func (e TL_account_sendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_sendChangePhoneCode)
	x.String(e.PhoneNumber)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_changePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_changePhone)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e TL_account_updateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updateDeviceLocked)
	x.Int(e.Period)
	return x.buf
}

func (e TL_account_getAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getAuthorizations)
	return x.buf
}

func (e TL_account_resetAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_resetAuthorization)
	x.Long(e.Hash)
	return x.buf
}

func (e TL_account_getPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getPassword)
	return x.buf
}

func (e TL_account_getPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getPasswordSettings)
	x.Bytes(e.Password.encode())
	return x.buf
}

func (e TL_account_updatePasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updatePasswordSettings)
	x.Bytes(e.Password.encode())
	x.Bytes(e.NewSettings.encode())
	return x.buf
}

func (e TL_account_sendConfirmPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_sendConfirmPhoneCode)
	x.String(e.Hash)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_confirmPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_confirmPhone)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e TL_account_getTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getTmpPassword)
	x.Bytes(e.Password.encode())
	x.Int(e.Period)
	return x.buf
}

func (e TL_account_getWebAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getWebAuthorizations)
	return x.buf
}

func (e TL_account_resetWebAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_resetWebAuthorization)
	x.Long(e.Hash)
	return x.buf
}

func (e TL_account_resetWebAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_resetWebAuthorizations)
	return x.buf
}

func (e TL_account_getAllSecureValues) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getAllSecureValues)
	return x.buf
}

func (e TL_account_getSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getSecureValue)
	x.Vector(e.Types)
	return x.buf
}

func (e TL_account_saveSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_saveSecureValue)
	x.Bytes(e.Value.encode())
	x.Long(e.SecureSecretID)
	return x.buf
}

func (e TL_account_deleteSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_deleteSecureValue)
	x.Vector(e.Types)
	return x.buf
}

func (e TL_account_getAuthorizationForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getAuthorizationForm)
	x.Int(e.BotID)
	x.String(e.Scope)
	x.String(e.PublicKey)
	return x.buf
}

func (e TL_account_acceptAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_acceptAuthorization)
	x.Int(e.BotID)
	x.String(e.Scope)
	x.String(e.PublicKey)
	x.Vector(e.ValueHashes)
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e TL_account_sendVerifyPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_sendVerifyPhoneCode)
	x.String(e.PhoneNumber)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_verifyPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_verifyPhone)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e TL_account_sendVerifyEmailCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_sendVerifyEmailCode)
	x.String(e.Email)
	return x.buf
}

func (e TL_account_verifyEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_verifyEmail)
	x.String(e.Email)
	x.String(e.Code)
	return x.buf
}

func (e TL_account_initTakeoutSession) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_initTakeoutSession)
	x.Int(e.Flags)
	//flag Contacts
	//flag MessageUsers
	//flag MessageChats
	//flag MessageMegagroups
	//flag MessageChannels
	//flag Files
	if e.Flags&32 != 0 {
		x.Int(e.FileMaxSize)
	}
	return x.buf
}

func (e TL_account_finishTakeoutSession) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_finishTakeoutSession)
	x.Int(e.Flags)
	//flag Success
	return x.buf
}

func (e TL_account_confirmPasswordEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_confirmPasswordEmail)
	x.String(e.Code)
	return x.buf
}

func (e TL_account_resendPasswordEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_resendPasswordEmail)
	return x.buf
}

func (e TL_account_cancelPasswordEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_cancelPasswordEmail)
	return x.buf
}

func (e TL_account_getContactSignUpNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getContactSignUpNotification)
	return x.buf
}

func (e TL_account_setContactSignUpNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_setContactSignUpNotification)
	x.Bytes(e.Silent.encode())
	return x.buf
}

func (e TL_account_getNotifyExceptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getNotifyExceptions)
	x.Int(e.Flags)
	//flag CompareSound
	if e.Flags&1 != 0 {
		x.Bytes(e.Peer.encode())
	}
	return x.buf
}

func (e TL_account_getWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getWallPaper)
	x.Bytes(e.Wallpaper.encode())
	return x.buf
}

func (e TL_account_uploadWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_uploadWallPaper)
	x.Bytes(e.File.encode())
	x.String(e.MimeType)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_saveWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_saveWallPaper)
	x.Bytes(e.Wallpaper.encode())
	x.Bytes(e.Unsave.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_installWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_installWallPaper)
	x.Bytes(e.Wallpaper.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_resetWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_resetWallPapers)
	return x.buf
}

func (e TL_account_getAutoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getAutoDownloadSettings)
	return x.buf
}

func (e TL_account_saveAutoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_saveAutoDownloadSettings)
	x.Int(e.Flags)
	//flag Low
	//flag High
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e TL_account_uploadTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_uploadTheme)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	x.String(e.FileName)
	x.String(e.MimeType)
	return x.buf
}

func (e TL_account_createTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_createTheme)
	x.Int(e.Flags)
	x.String(e.Slug)
	x.String(e.Title)
	if e.Flags&4 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e TL_account_updateTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_updateTheme)
	x.Int(e.Flags)
	x.String(e.Format)
	x.Bytes(e.Theme.encode())
	if e.Flags&1 != 0 {
		x.String(e.Slug)
	}
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e TL_account_saveTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_saveTheme)
	x.Bytes(e.Theme.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e TL_account_installTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_installTheme)
	x.Int(e.Flags)
	//flag Dark
	if e.Flags&2 != 0 {
		x.String(e.Format)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Theme.encode())
	}
	return x.buf
}

func (e TL_account_getTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getTheme)
	x.String(e.Format)
	x.Bytes(e.Theme.encode())
	x.Long(e.DocumentID)
	return x.buf
}

func (e TL_account_getThemes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getThemes)
	x.String(e.Format)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_account_setContentSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_setContentSettings)
	x.Int(e.Flags)
	//flag SensitiveEnabled
	return x.buf
}

func (e TL_account_getContentSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getContentSettings)
	return x.buf
}

func (e TL_account_getMultiWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_account_getMultiWallPapers)
	x.Vector(e.Wallpapers)
	return x.buf
}

func (e TL_users_getUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_users_getUsers)
	x.Vector(e.ID)
	return x.buf
}

func (e TL_users_getFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_users_getFullUser)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_users_setSecureValueErrors) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_users_setSecureValueErrors)
	x.Bytes(e.ID.encode())
	x.Vector(e.Errors)
	return x.buf
}

func (e TL_contacts_getContactIDs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getContactIDs)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_contacts_getStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getStatuses)
	return x.buf
}

func (e TL_contacts_getContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getContacts)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_contacts_importContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_importContacts)
	x.Vector(e.Contacts)
	return x.buf
}

func (e TL_contacts_deleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_deleteContacts)
	x.Vector(e.ID)
	return x.buf
}

func (e TL_contacts_deleteByPhones) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_deleteByPhones)
	x.VectorString(e.Phones)
	return x.buf
}

func (e TL_contacts_block) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_block)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_contacts_unblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_unblock)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_contacts_getBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_contacts_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_search)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_contacts_resolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_resolveUsername)
	x.String(e.Username)
	return x.buf
}

func (e TL_contacts_getTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getTopPeers)
	x.Int(e.Flags)
	//flag Correspondents
	//flag BotsPm
	//flag BotsInline
	//flag PhoneCalls
	//flag ForwardUsers
	//flag ForwardChats
	//flag Groups
	//flag Channels
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_contacts_resetTopPeerRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_resetTopPeerRating)
	x.Bytes(e.Category.encode())
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_contacts_resetSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_resetSaved)
	return x.buf
}

func (e TL_contacts_getSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getSaved)
	return x.buf
}

func (e TL_contacts_toggleTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_toggleTopPeers)
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e TL_contacts_addContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_addContact)
	x.Int(e.Flags)
	//flag AddPhonePrivacyException
	x.Bytes(e.ID.encode())
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Phone)
	return x.buf
}

func (e TL_contacts_acceptContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_acceptContact)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_contacts_getLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_contacts_getLocated)
	x.Int(e.Flags)
	//flag Background
	x.Bytes(e.GeoPoint.encode())
	if e.Flags&1 != 0 {
		x.Int(e.SelfExpires)
	}
	return x.buf
}

func (e TL_messages_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getMessages)
	x.Vector(e.ID)
	return x.buf
}

func (e TL_messages_getDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getDialogs)
	x.Int(e.Flags)
	//flag ExcludePinned
	if e.Flags&2 != 0 {
		x.Int(e.FolderID)
	}
	x.Int(e.OffsetDate)
	x.Int(e.OffsetID)
	x.Bytes(e.OffsetPeer.encode())
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.OffsetID)
	x.Int(e.OffsetDate)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxID)
	x.Int(e.MinID)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_search)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	if e.Flags&1 != 0 {
		x.Bytes(e.FromID.encode())
	}
	x.Bytes(e.Filter.encode())
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	x.Int(e.OffsetID)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxID)
	x.Int(e.MinID)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_readHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_messages_deleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_deleteHistory)
	x.Int(e.Flags)
	//flag JustClear
	//flag Revoke
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_messages_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_deleteMessages)
	x.Int(e.Flags)
	//flag Revoke
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_messages_receivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_receivedMessages)
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_messages_setTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e TL_messages_sendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendMessage)
	x.Int(e.Flags)
	//flag NoWebpage
	//flag Silent
	//flag Background
	//flag ClearDraft
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.String(e.Message)
	x.Long(e.RandomID)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e TL_messages_sendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendMedia)
	x.Int(e.Flags)
	//flag Silent
	//flag Background
	//flag ClearDraft
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Bytes(e.Media.encode())
	x.String(e.Message)
	x.Long(e.RandomID)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e TL_messages_forwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_forwardMessages)
	x.Int(e.Flags)
	//flag Silent
	//flag Background
	//flag WithMyScore
	//flag Grouped
	x.Bytes(e.FromPeer.encode())
	x.VectorInt(e.ID)
	x.VectorLong(e.RandomID)
	x.Bytes(e.ToPeer.encode())
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e TL_messages_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_reportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getPeerSettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_report) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_report)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	x.Bytes(e.Reason.encode())
	return x.buf
}

func (e TL_messages_getChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getChats)
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_messages_getFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getFullChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_messages_editChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editChatTitle)
	x.Int(e.ChatID)
	x.String(e.Title)
	return x.buf
}

func (e TL_messages_editChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editChatPhoto)
	x.Int(e.ChatID)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TL_messages_addChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_addChatUser)
	x.Int(e.ChatID)
	x.Bytes(e.UserID.encode())
	x.Int(e.FwdLimit)
	return x.buf
}

func (e TL_messages_deleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_deleteChatUser)
	x.Int(e.ChatID)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_messages_createChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_createChat)
	x.Vector(e.Users)
	x.String(e.Title)
	return x.buf
}

func (e TL_messages_getDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getDhConfig)
	x.Int(e.Version)
	x.Int(e.RandomLength)
	return x.buf
}

func (e TL_messages_requestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_requestEncryption)
	x.Bytes(e.UserID.encode())
	x.Int(e.RandomID)
	x.StringBytes(e.GA)
	return x.buf
}

func (e TL_messages_acceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_acceptEncryption)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Long(e.KeyFingerprint)
	return x.buf
}

func (e TL_messages_discardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_discardEncryption)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_messages_setEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setEncryptedTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}

func (e TL_messages_readEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_readEncryptedHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxDate)
	return x.buf
}

func (e TL_messages_sendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendEncrypted)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_messages_sendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.StringBytes(e.Data)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_messages_sendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendEncryptedService)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_messages_receivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_receivedQueue)
	x.Int(e.MaxQts)
	return x.buf
}

func (e TL_messages_reportEncryptedSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_reportEncryptedSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_readMessageContents)
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_messages_getStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getStickers)
	x.String(e.Emoticon)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getAllStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getWebPagePreview) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getWebPagePreview)
	x.Int(e.Flags)
	x.String(e.Message)
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_messages_exportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_exportChatInvite)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_checkChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_checkChatInvite)
	x.String(e.Hash)
	return x.buf
}

func (e TL_messages_importChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_importChatInvite)
	x.String(e.Hash)
	return x.buf
}

func (e TL_messages_getStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_messages_installStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_installStickerSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Archived.encode())
	return x.buf
}

func (e TL_messages_uninstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_uninstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_messages_startBot) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_startBot)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.String(e.StartParam)
	return x.buf
}

func (e TL_messages_getMessagesViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getMessagesViews)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	x.Bytes(e.Increment.encode())
	return x.buf
}

func (e TL_messages_editChatAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editChatAdmin)
	x.Int(e.ChatID)
	x.Bytes(e.UserID.encode())
	x.Bytes(e.IsAdmin.encode())
	return x.buf
}

func (e TL_messages_migrateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_migrateChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e TL_messages_searchGlobal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_searchGlobal)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.String(e.Q)
	x.Int(e.OffsetRate)
	x.Bytes(e.OffsetPeer.encode())
	x.Int(e.OffsetID)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_reorderStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_reorderStickerSets)
	x.Int(e.Flags)
	//flag Masks
	x.VectorLong(e.Order)
	return x.buf
}

func (e TL_messages_getDocumentByHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getDocumentByHash)
	x.StringBytes(e.Sha256)
	x.Int(e.Size)
	x.String(e.MimeType)
	return x.buf
}

func (e TL_messages_searchGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_searchGifs)
	x.String(e.Q)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_messages_getSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getSavedGifs)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_saveGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_saveGif)
	x.Bytes(e.ID.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e TL_messages_getInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getInlineBotResults)
	x.Int(e.Flags)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.GeoPoint.encode())
	}
	x.String(e.Query)
	x.String(e.Offset)
	return x.buf
}

func (e TL_messages_setInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setInlineBotResults)
	x.Int(e.Flags)
	//flag Gallery
	//flag Private
	x.Long(e.QueryID)
	x.Vector(e.Results)
	x.Int(e.CacheTime)
	if e.Flags&4 != 0 {
		x.String(e.NextOffset)
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.SwitchPm.encode())
	}
	return x.buf
}

func (e TL_messages_sendInlineBotResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendInlineBotResult)
	x.Int(e.Flags)
	//flag Silent
	//flag Background
	//flag ClearDraft
	//flag HideVia
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Long(e.RandomID)
	x.Long(e.QueryID)
	x.String(e.ID)
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e TL_messages_getMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getMessageEditData)
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	return x.buf
}

func (e TL_messages_editMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editMessage)
	x.Int(e.Flags)
	//flag NoWebpage
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	if e.Flags&2048 != 0 {
		x.String(e.Message)
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&32768 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e TL_messages_editInlineBotMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editInlineBotMessage)
	x.Int(e.Flags)
	//flag NoWebpage
	x.Bytes(e.ID.encode())
	if e.Flags&2048 != 0 {
		x.String(e.Message)
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_messages_getBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getBotCallbackAnswer)
	x.Int(e.Flags)
	//flag Game
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Data)
	}
	return x.buf
}

func (e TL_messages_setBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setBotCallbackAnswer)
	x.Int(e.Flags)
	//flag Alert
	x.Long(e.QueryID)
	if e.Flags&1 != 0 {
		x.String(e.Message)
	}
	if e.Flags&4 != 0 {
		x.String(e.Url)
	}
	x.Int(e.CacheTime)
	return x.buf
}

func (e TL_messages_getPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getPeerDialogs)
	x.Vector(e.Peers)
	return x.buf
}

func (e TL_messages_saveDraft) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_saveDraft)
	x.Int(e.Flags)
	//flag NoWebpage
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e TL_messages_getAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getAllDrafts)
	return x.buf
}

func (e TL_messages_getFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getFeaturedStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_readFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_readFeaturedStickers)
	x.VectorLong(e.ID)
	return x.buf
}

func (e TL_messages_getRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getRecentStickers)
	x.Int(e.Flags)
	//flag Attached
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_saveRecentSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_saveRecentSticker)
	x.Int(e.Flags)
	//flag Attached
	x.Bytes(e.ID.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e TL_messages_clearRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_clearRecentStickers)
	x.Int(e.Flags)
	//flag Attached
	return x.buf
}

func (e TL_messages_getArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getArchivedStickers)
	x.Int(e.Flags)
	//flag Masks
	x.Long(e.OffsetID)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_getMaskStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getMaskStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getAttachedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getAttachedStickers)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TL_messages_setGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setGameScore)
	x.Int(e.Flags)
	//flag EditMessage
	//flag Force
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	x.Bytes(e.UserID.encode())
	x.Int(e.Score)
	return x.buf
}

func (e TL_messages_setInlineGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setInlineGameScore)
	x.Int(e.Flags)
	//flag EditMessage
	//flag Force
	x.Bytes(e.ID.encode())
	x.Bytes(e.UserID.encode())
	x.Int(e.Score)
	return x.buf
}

func (e TL_messages_getGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getGameHighScores)
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_messages_getInlineGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getInlineGameHighScores)
	x.Bytes(e.ID.encode())
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_messages_getCommonChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getCommonChats)
	x.Bytes(e.UserID.encode())
	x.Int(e.MaxID)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_getAllChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getAllChats)
	x.VectorInt(e.ExceptIds)
	return x.buf
}

func (e TL_messages_getWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getWebPage)
	x.String(e.Url)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_toggleDialogPin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_toggleDialogPin)
	x.Int(e.Flags)
	//flag Pinned
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_reorderPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_reorderPinnedDialogs)
	x.Int(e.Flags)
	//flag Force
	x.Int(e.FolderID)
	x.Vector(e.Order)
	return x.buf
}

func (e TL_messages_getPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getPinnedDialogs)
	x.Int(e.FolderID)
	return x.buf
}

func (e TL_messages_setBotShippingResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setBotShippingResults)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	if e.Flags&1 != 0 {
		x.String(e.Error)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.ShippingOptions)
	}
	return x.buf
}

func (e TL_messages_setBotPrecheckoutResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_setBotPrecheckoutResults)
	x.Int(e.Flags)
	//flag Success
	x.Long(e.QueryID)
	if e.Flags&1 != 0 {
		x.String(e.Error)
	}
	return x.buf
}

func (e TL_messages_uploadMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_uploadMedia)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e TL_messages_sendScreenshotNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendScreenshotNotification)
	x.Bytes(e.Peer.encode())
	x.Int(e.ReplyToMsgID)
	x.Long(e.RandomID)
	return x.buf
}

func (e TL_messages_getFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getFavedStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_faveSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_faveSticker)
	x.Bytes(e.ID.encode())
	x.Bytes(e.Unfave.encode())
	return x.buf
}

func (e TL_messages_getUnreadMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getUnreadMentions)
	x.Bytes(e.Peer.encode())
	x.Int(e.OffsetID)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxID)
	x.Int(e.MinID)
	return x.buf
}

func (e TL_messages_readMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_readMentions)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getRecentLocations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getRecentLocations)
	x.Bytes(e.Peer.encode())
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_sendMultiMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendMultiMedia)
	x.Int(e.Flags)
	//flag Silent
	//flag Background
	//flag ClearDraft
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Vector(e.MultiMedia)
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e TL_messages_uploadEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_uploadEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_messages_searchStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_searchStickerSets)
	x.Int(e.Flags)
	//flag ExcludeFeatured
	x.String(e.Q)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getSplitRanges) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getSplitRanges)
	return x.buf
}

func (e TL_messages_markDialogUnread) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_markDialogUnread)
	x.Int(e.Flags)
	//flag Unread
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getDialogUnreadMarks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getDialogUnreadMarks)
	return x.buf
}

func (e TL_messages_clearAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_clearAllDrafts)
	return x.buf
}

func (e TL_messages_updatePinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_updatePinnedMessage)
	x.Int(e.Flags)
	//flag Silent
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	return x.buf
}

func (e TL_messages_sendVote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendVote)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Vector(e.Options)
	return x.buf
}

func (e TL_messages_getPollResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getPollResults)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	return x.buf
}

func (e TL_messages_getOnlines) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getOnlines)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getStatsURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getStatsURL)
	x.Int(e.Flags)
	//flag Dark
	x.Bytes(e.Peer.encode())
	x.String(e.Params)
	return x.buf
}

func (e TL_messages_editChatAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editChatAbout)
	x.Bytes(e.Peer.encode())
	x.String(e.About)
	return x.buf
}

func (e TL_messages_editChatDefaultBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_editChatDefaultBannedRights)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.BannedRights.encode())
	return x.buf
}

func (e TL_messages_getEmojiKeywords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getEmojiKeywords)
	x.String(e.LangCode)
	return x.buf
}

func (e TL_messages_getEmojiKeywordsDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getEmojiKeywordsDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	return x.buf
}

func (e TL_messages_getEmojiKeywordsLanguages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getEmojiKeywordsLanguages)
	x.VectorString(e.LangCodes)
	return x.buf
}

func (e TL_messages_getEmojiURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getEmojiURL)
	x.String(e.LangCode)
	return x.buf
}

func (e TL_messages_getSearchCounters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getSearchCounters)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Filters)
	return x.buf
}

func (e TL_messages_requestUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_requestUrlAuth)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ButtonID)
	return x.buf
}

func (e TL_messages_acceptUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_acceptUrlAuth)
	x.Int(e.Flags)
	//flag WriteAllowed
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ButtonID)
	return x.buf
}

func (e TL_messages_hidePeerSettingsBar) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_hidePeerSettingsBar)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_messages_getScheduledHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getScheduledHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Hash)
	return x.buf
}

func (e TL_messages_getScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_messages_sendScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_sendScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_messages_deleteScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_deleteScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_messages_getPollVotes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getPollVotes)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Option)
	}
	if e.Flags&2 != 0 {
		x.String(e.Offset)
	}
	x.Int(e.Limit)
	return x.buf
}

func (e TL_messages_toggleStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_toggleStickerSets)
	x.Int(e.Flags)
	//flag Uninstall
	//flag Archive
	//flag Unarchive
	x.Vector(e.Stickersets)
	return x.buf
}

func (e TL_messages_getDialogFilters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getDialogFilters)
	return x.buf
}

func (e TL_messages_getSuggestedDialogFilters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getSuggestedDialogFilters)
	return x.buf
}

func (e TL_messages_updateDialogFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_updateDialogFilter)
	x.Int(e.Flags)
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Filter.encode())
	}
	return x.buf
}

func (e TL_messages_updateDialogFiltersOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_updateDialogFiltersOrder)
	x.VectorInt(e.Order)
	return x.buf
}

func (e TL_messages_getOldFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_messages_getOldFeaturedStickers)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_updates_getState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_getState)
	return x.buf
}

func (e TL_updates_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_getDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	if e.Flags&1 != 0 {
		x.Int(e.PtsTotalLimit)
	}
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.buf
}

func (e TL_updates_getChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_updates_getChannelDifference)
	x.Int(e.Flags)
	//flag Force
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Pts)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_photos_updateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_updateProfilePhoto)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_photos_uploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_uploadProfilePhoto)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e TL_photos_deletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_deletePhotos)
	x.Vector(e.ID)
	return x.buf
}

func (e TL_photos_getUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_photos_getUserPhotos)
	x.Bytes(e.UserID.encode())
	x.Int(e.Offset)
	x.Long(e.MaxID)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_upload_saveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_saveFilePart)
	x.Long(e.FileID)
	x.Int(e.FilePart)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_upload_getFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_getFile)
	x.Int(e.Flags)
	//flag Precise
	//flag CdnSupported
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_upload_saveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_saveBigFilePart)
	x.Long(e.FileID)
	x.Int(e.FilePart)
	x.Int(e.FileTotalParts)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e TL_upload_getWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_getWebFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_upload_getCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_getCdnFile)
	x.StringBytes(e.FileToken)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_upload_reuploadCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_reuploadCdnFile)
	x.StringBytes(e.FileToken)
	x.StringBytes(e.RequestToken)
	return x.buf
}

func (e TL_upload_getCdnFileHashes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_getCdnFileHashes)
	x.StringBytes(e.FileToken)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_upload_getFileHashes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_upload_getFileHashes)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	return x.buf
}

func (e TL_help_getConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getConfig)
	return x.buf
}

func (e TL_help_getNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getNearestDc)
	return x.buf
}

func (e TL_help_getAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getAppUpdate)
	x.String(e.Source)
	return x.buf
}

func (e TL_help_getInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getInviteText)
	return x.buf
}

func (e TL_help_getSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getSupport)
	return x.buf
}

func (e TL_help_getAppChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getAppChangelog)
	x.String(e.PrevAppVersion)
	return x.buf
}

func (e TL_help_setBotUpdatesStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_setBotUpdatesStatus)
	x.Int(e.PendingUpdatesCount)
	x.String(e.Message)
	return x.buf
}

func (e TL_help_getCdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getCdnConfig)
	return x.buf
}

func (e TL_help_getRecentMeUrls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getRecentMeUrls)
	x.String(e.Referer)
	return x.buf
}

func (e TL_help_getTermsOfServiceUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getTermsOfServiceUpdate)
	return x.buf
}

func (e TL_help_acceptTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_acceptTermsOfService)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e TL_help_getDeepLinkInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getDeepLinkInfo)
	x.String(e.Path)
	return x.buf
}

func (e TL_help_getAppConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getAppConfig)
	return x.buf
}

func (e TL_help_saveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_saveAppLog)
	x.Vector(e.Events)
	return x.buf
}

func (e TL_help_getPassportConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getPassportConfig)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_help_getSupportName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getSupportName)
	return x.buf
}

func (e TL_help_getUserInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getUserInfo)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_help_editUserInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_editUserInfo)
	x.Bytes(e.UserID.encode())
	x.String(e.Message)
	x.Vector(e.Entities)
	return x.buf
}

func (e TL_help_getPromoData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_getPromoData)
	return x.buf
}

func (e TL_help_hidePromoData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_help_hidePromoData)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_channels_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_readHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_channels_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_deleteMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_channels_deleteUserHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_deleteUserHistory)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_channels_reportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_reportSpam)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_channels_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getMessages)
	x.Bytes(e.Channel.encode())
	x.Vector(e.ID)
	return x.buf
}

func (e TL_channels_getParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getParticipants)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e TL_channels_getParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getParticipant)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e TL_channels_getChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getChannels)
	x.Vector(e.ID)
	return x.buf
}

func (e TL_channels_getFullChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getFullChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_createChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_createChannel)
	x.Int(e.Flags)
	//flag Broadcast
	//flag Megagroup
	x.String(e.Title)
	x.String(e.About)
	if e.Flags&4 != 0 {
		x.Bytes(e.GeoPoint.encode())
	}
	if e.Flags&4 != 0 {
		x.String(e.Address)
	}
	return x.buf
}

func (e TL_channels_editAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_editAdmin)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.Bytes(e.AdminRights.encode())
	x.String(e.Rank)
	return x.buf
}

func (e TL_channels_editTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_editTitle)
	x.Bytes(e.Channel.encode())
	x.String(e.Title)
	return x.buf
}

func (e TL_channels_editPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_editPhoto)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e TL_channels_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_checkUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

func (e TL_channels_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_updateUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

func (e TL_channels_joinChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_joinChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_leaveChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_leaveChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_inviteToChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_inviteToChannel)
	x.Bytes(e.Channel.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e TL_channels_deleteChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_deleteChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_channels_exportMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_exportMessageLink)
	x.Bytes(e.Channel.encode())
	x.Int(e.ID)
	x.Bytes(e.Grouped.encode())
	return x.buf
}

func (e TL_channels_toggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_toggleSignatures)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e TL_channels_getAdminedPublicChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getAdminedPublicChannels)
	x.Int(e.Flags)
	//flag ByLocation
	//flag CheckLimit
	return x.buf
}

func (e TL_channels_editBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_editBanned)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.Bytes(e.BannedRights.encode())
	return x.buf
}

func (e TL_channels_getAdminLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getAdminLog)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.String(e.Q)
	if e.Flags&1 != 0 {
		x.Bytes(e.EventsFilter.encode())
	}
	if e.Flags&2 != 0 {
		x.Vector(e.Admins)
	}
	x.Long(e.MaxID)
	x.Long(e.MinID)
	x.Int(e.Limit)
	return x.buf
}

func (e TL_channels_setStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_setStickers)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e TL_channels_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_readMessageContents)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e TL_channels_deleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_deleteHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e TL_channels_togglePreHistoryHidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_togglePreHistoryHidden)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e TL_channels_getLeftChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getLeftChannels)
	x.Int(e.Offset)
	return x.buf
}

func (e TL_channels_getGroupsForDiscussion) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getGroupsForDiscussion)
	return x.buf
}

func (e TL_channels_setDiscussionGroup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_setDiscussionGroup)
	x.Bytes(e.Broadcast.encode())
	x.Bytes(e.Group.encode())
	return x.buf
}

func (e TL_channels_editCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_editCreator)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.Bytes(e.Password.encode())
	return x.buf
}

func (e TL_channels_editLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_editLocation)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Address)
	return x.buf
}

func (e TL_channels_toggleSlowMode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_toggleSlowMode)
	x.Bytes(e.Channel.encode())
	x.Int(e.Seconds)
	return x.buf
}

func (e TL_channels_getInactiveChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_channels_getInactiveChannels)
	return x.buf
}

func (e TL_bots_sendCustomRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bots_sendCustomRequest)
	x.String(e.CustomMethod)
	x.Bytes(e.Params.encode())
	return x.buf
}

func (e TL_bots_answerWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bots_answerWebhookJSONQuery)
	x.Long(e.QueryID)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e TL_bots_setBotCommands) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_bots_setBotCommands)
	x.Vector(e.Commands)
	return x.buf
}

func (e TL_payments_getPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_getPaymentForm)
	x.Int(e.MsgID)
	return x.buf
}

func (e TL_payments_getPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_getPaymentReceipt)
	x.Int(e.MsgID)
	return x.buf
}

func (e TL_payments_validateRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_validateRequestedInfo)
	x.Int(e.Flags)
	//flag Save
	x.Int(e.MsgID)
	x.Bytes(e.Info.encode())
	return x.buf
}

func (e TL_payments_sendPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_sendPaymentForm)
	x.Int(e.Flags)
	x.Int(e.MsgID)
	if e.Flags&1 != 0 {
		x.String(e.RequestedInfoID)
	}
	if e.Flags&2 != 0 {
		x.String(e.ShippingOptionID)
	}
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e TL_payments_getSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_getSavedInfo)
	return x.buf
}

func (e TL_payments_clearSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_clearSavedInfo)
	x.Int(e.Flags)
	//flag Credentials
	//flag Info
	return x.buf
}

func (e TL_payments_getBankCardData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_payments_getBankCardData)
	x.String(e.Number)
	return x.buf
}

func (e TL_stickers_createStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickers_createStickerSet)
	x.Int(e.Flags)
	//flag Masks
	//flag Animated
	x.Bytes(e.UserID.encode())
	x.String(e.Title)
	x.String(e.ShortName)
	if e.Flags&4 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	x.Vector(e.Stickers)
	return x.buf
}

func (e TL_stickers_removeStickerFromSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickers_removeStickerFromSet)
	x.Bytes(e.Sticker.encode())
	return x.buf
}

func (e TL_stickers_changeStickerPosition) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickers_changeStickerPosition)
	x.Bytes(e.Sticker.encode())
	x.Int(e.Position)
	return x.buf
}

func (e TL_stickers_addStickerToSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickers_addStickerToSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Sticker.encode())
	return x.buf
}

func (e TL_stickers_setStickerSetThumb) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stickers_setStickerSetThumb)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Thumb.encode())
	return x.buf
}

func (e TL_phone_getCallConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_getCallConfig)
	return x.buf
}

func (e TL_phone_requestCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_requestCall)
	x.Int(e.Flags)
	//flag Video
	x.Bytes(e.UserID.encode())
	x.Int(e.RandomID)
	x.StringBytes(e.GAHash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phone_acceptCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_acceptCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phone_confirmCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_confirmCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GA)
	x.Long(e.KeyFingerprint)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e TL_phone_receivedCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_receivedCall)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e TL_phone_discardCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_discardCall)
	x.Int(e.Flags)
	//flag Video
	x.Bytes(e.Peer.encode())
	x.Int(e.Duration)
	x.Bytes(e.Reason.encode())
	x.Long(e.ConnectionID)
	return x.buf
}

func (e TL_phone_setCallRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_setCallRating)
	x.Int(e.Flags)
	//flag UserInitiative
	x.Bytes(e.Peer.encode())
	x.Int(e.Rating)
	x.String(e.Comment)
	return x.buf
}

func (e TL_phone_saveCallDebug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_saveCallDebug)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Debug.encode())
	return x.buf
}

func (e TL_phone_sendSignalingData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_phone_sendSignalingData)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.Data)
	return x.buf
}

func (e TL_langpack_getLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langpack_getLangPack)
	x.String(e.LangPack)
	x.String(e.LangCode)
	return x.buf
}

func (e TL_langpack_getStrings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langpack_getStrings)
	x.String(e.LangPack)
	x.String(e.LangCode)
	x.VectorString(e.Keys)
	return x.buf
}

func (e TL_langpack_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langpack_getDifference)
	x.String(e.LangPack)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	return x.buf
}

func (e TL_langpack_getLanguages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langpack_getLanguages)
	x.String(e.LangPack)
	return x.buf
}

func (e TL_langpack_getLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_langpack_getLanguage)
	x.String(e.LangPack)
	x.String(e.LangCode)
	return x.buf
}

func (e TL_folders_editPeerFolders) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_folders_editPeerFolders)
	x.Vector(e.FolderPeers)
	return x.buf
}

func (e TL_folders_deleteFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_folders_deleteFolder)
	x.Int(e.FolderID)
	return x.buf
}

func (e TL_stats_getBroadcastStats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stats_getBroadcastStats)
	x.Int(e.Flags)
	//flag Dark
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e TL_stats_loadAsyncGraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC_stats_loadAsyncGraph)
	x.Int(e.Flags)
	x.String(e.Token)
	if e.Flags&1 != 0 {
		x.Long(e.X)
	}
	return x.buf
}

func (e TL_req_pq) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_req_pq_multi) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_req_DH_params) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_set_client_DH_params) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_destroy_auth_key) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_rpc_drop_answer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_get_future_salts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_ping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_ping_delay_disconnect) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_destroy_session) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_boolFalse) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_boolTrue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_true) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_error) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_null) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerSelf) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerUserFromMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerChannelFromMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputUserEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputUserSelf) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputUserFromMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPhoneContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputFileBig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaUploadedPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaGeoPoint) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaUploadedDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaVenue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaGifExternal) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaPhotoExternal) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaDocumentExternal) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaGame) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaInvoice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaGeoLive) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaPoll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMediaDice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputChatPhotoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputChatUploadedPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputChatPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputGeoPointEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputGeoPoint) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPhotoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputEncryptedFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputDocumentFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputSecureFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputTakeoutFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPhotoFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPhotoLegacyFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerPhotoFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetThumb) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileUnknown) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_filePartial) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileJpeg) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileGif) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_filePng) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_filePdf) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileMp3) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileMov) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileMp4) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_storage_fileWebp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_user) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userProfilePhotoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userStatusEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userStatusOnline) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userStatusOffline) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userStatusRecently) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userStatusLastWeek) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userStatusLastMonth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatForbidden) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelForbidden) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatFull) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelFull) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatParticipant) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatParticipantCreator) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatParticipantAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatParticipantsForbidden) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatPhotoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_message) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaGeo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaUnsupported) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaWebPage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaVenue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaGame) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaInvoice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaGeoLive) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaPoll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageMediaDice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatCreate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatEditTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatEditPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatDeletePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatAddUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatDeleteUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatJoinedByLink) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChannelCreate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChatMigrateTo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionChannelMigrateFrom) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionPinMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionHistoryClear) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionGameScore) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionPaymentSentMe) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionPaymentSent) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionPhoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionScreenshotTaken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionCustomAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionBotAllowed) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionSecureValuesSentMe) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionSecureValuesSent) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageActionContactSignUp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dialog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dialogFolder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photoSizeEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photoSize) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photoCachedSize) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photoStrippedSize) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_geoPointEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_geoPoint) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_sentCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_authorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_authorizationSignUpRequired) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_exportedAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputNotifyPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputNotifyUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputNotifyChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputNotifyBroadcasts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPeerNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_wallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_wallPaperNoFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonViolence) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonPornography) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonChildAbuse) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonOther) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonCopyright) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputReportReasonGeoIrrelevant) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_userFull) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_importedContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contactBlocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contactStatus) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_contactsNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_contacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_importedContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_blocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_blockedSlice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_dialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_dialogsSlice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_dialogsNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_messages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_messagesSlice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_channelMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_messagesNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_chats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_chatsSlice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_chatFull) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_affectedHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterPhotos) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterVideo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterPhotoVideo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterUrl) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterGif) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterVoice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterMusic) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterChatPhotos) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterPhoneCalls) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterRoundVoice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterRoundVideo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterMyMentions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterGeo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagesFilterContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateNewMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateMessageID) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDeleteMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatUserTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserStatus) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateNewEncryptedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateEncryptedChatTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateEncryptedMessagesRead) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatParticipantAdd) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatParticipantDelete) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDcOptions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserBlocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateServiceNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePrivacy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateReadHistoryInbox) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateReadHistoryOutbox) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateWebPage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateReadMessagesContents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannelTooLong) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateNewChannelMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateReadChannelInbox) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDeleteChannelMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannelMessageViews) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatParticipantAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateNewStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateStickerSetsOrder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateSavedGifs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotInlineQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotInlineSend) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateEditChannelMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannelPinnedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotCallbackQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateEditMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateInlineBotCallbackQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateReadChannelOutbox) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDraftMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateReadFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateRecentStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePtsChanged) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannelWebPage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDialogPinned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePinnedDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotWebhookJSON) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotWebhookJSONQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotShippingQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateBotPrecheckoutQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePhoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateLangPackTooLong) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateLangPack) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateFavedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannelReadMessagesContents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateContactsReset) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChannelAvailableMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDialogUnreadMark) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateUserPinnedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatPinnedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateMessagePoll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateChatDefaultBannedRights) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateFolderPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePeerSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePeerLocated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateNewScheduledMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDeleteScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateGeoLiveViewed) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateMessagePollVote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDialogFilter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDialogFilterOrder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateDialogFilters) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatePhoneCallSignalingData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_state) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_differenceEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_difference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_differenceSlice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_differenceTooLong) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatesTooLong) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateShortMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateShortChatMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateShort) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updatesCombined) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updateShortSentMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photos_photos) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photos_photosSlice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photos_photo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_file) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_fileCdnRedirect) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dcOption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_config) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_nearestDc) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_appUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_noAppUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_inviteText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedChatEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedChatWaiting) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedChatRequested) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedChatDiscarded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputEncryptedChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedFileEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputEncryptedFileEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputEncryptedFileUploaded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputEncryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputEncryptedFileBigUploaded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_encryptedMessageService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_dhConfigNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_dhConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sentEncryptedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sentEncryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputDocumentEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_document) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_support) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_notifyPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_notifyUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_notifyChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_notifyBroadcasts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageTypingAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageCancelAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageRecordVideoAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageUploadVideoAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageRecordAudioAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageUploadAudioAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageUploadPhotoAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageUploadDocumentAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageGeoLocationAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageChooseContactAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageGamePlayAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageRecordRoundAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_sendMessageUploadRoundAction) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_found) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyStatusTimestamp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyPhoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyPhoneP2P) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyForwards) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyPhoneNumber) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyKeyAddedByPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyStatusTimestamp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyPhoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyPhoneP2P) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyForwards) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyPhoneNumber) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyKeyAddedByPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueAllowContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueAllowAll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueAllowUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueDisallowContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueDisallowAll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueDisallowUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueAllowChatParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPrivacyValueDisallowChatParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueAllowContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueAllowAll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueAllowUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueDisallowContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueDisallowAll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueDisallowUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueAllowChatParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_privacyValueDisallowChatParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_privacyRules) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_accountDaysTTL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeImageSize) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeAnimated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeSticker) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeVideo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeAudio) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeFilename) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_documentAttributeHasStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_stickersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_stickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickerPack) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_allStickersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_allStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_affectedMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webPageEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webPagePending) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webPage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webPageNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_authorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_authorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_password) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_passwordSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_passwordInputSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_passwordRecovery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_receivedNotifyMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatInviteEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatInviteExported) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatInviteAlready) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetID) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetShortName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetAnimatedEmoji) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetDice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_stickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botCommand) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButton) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonUrl) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonCallback) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonRequestPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonRequestGeoLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonSwitchInline) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonGame) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonBuy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonUrlAuth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputKeyboardButtonUrlAuth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonRequestPoll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_keyboardButtonRow) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_replyKeyboardHide) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_replyKeyboardForceReply) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_replyKeyboardMarkup) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_replyInlineMarkup) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityUnknown) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityMention) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityHashtag) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityBotCommand) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityUrl) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityBold) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityItalic) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityPre) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityTextUrl) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityMentionName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessageEntityMentionName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityCashtag) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityUnderline) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityStrike) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityBlockquote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageEntityBankCard) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputChannelEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputChannelFromMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_resolvedPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageRange) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_channelDifferenceEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_channelDifferenceTooLong) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_channelDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelMessagesFilterEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelMessagesFilter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipant) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantSelf) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantCreator) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantBanned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsRecent) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsAdmins) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsKicked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsBots) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsBanned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsSearch) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelParticipantsContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_channelParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_channelParticipantsNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_channelParticipant) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_termsOfService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_foundGif) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_foundGifCached) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_foundGifs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_savedGifsNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_savedGifs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageMediaAuto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageMediaGeo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageMediaVenue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageMediaContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageGame) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineResult) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineResultPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineResultDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineResultGame) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineMessageMediaAuto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineMessageText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineMessageMediaGeo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineMessageMediaVenue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineMessageMediaContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineResult) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_botInlineMediaResult) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_botResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_exportedMessageLink) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageFwdHeader) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_codeTypeSms) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_codeTypeCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_codeTypeFlashCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_sentCodeTypeApp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_sentCodeTypeSms) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_sentCodeTypeCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_sentCodeTypeFlashCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_botCallbackAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_messageEditData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputBotInlineMessageID) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inlineBotSwitchPM) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_peerDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryBotsPM) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryBotsInline) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryCorrespondents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryGroups) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryPhoneCalls) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryForwardUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryForwardChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_topPeerCategoryPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_topPeersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_topPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_topPeersDisabled) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_draftMessageEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_draftMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_featuredStickersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_featuredStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_recentStickersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_recentStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_archivedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_stickerSetInstallResultSuccess) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_stickerSetInstallResultArchive) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickerSetCovered) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickerSetMultiCovered) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_maskCoords) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickeredMediaPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickeredMediaDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_game) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputGameID) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputGameShortName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_highScore) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_highScores) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textPlain) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textBold) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textItalic) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textUnderline) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textStrike) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textFixed) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textUrl) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textConcat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textSubscript) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textSuperscript) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textMarked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textImage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_textAnchor) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockUnsupported) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockSubtitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockAuthorDate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockHeader) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockSubheader) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockParagraph) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockPreformatted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockFooter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockDivider) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockAnchor) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockList) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockBlockquote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockPullquote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockVideo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockCover) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockEmbed) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockEmbedPost) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockCollage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockSlideshow) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockAudio) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockKicker) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockTable) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockOrderedList) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockDetails) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockRelatedArticles) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageBlockMap) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallDiscardReasonMissed) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallDiscardReasonDisconnect) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallDiscardReasonHangup) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallDiscardReasonBusy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dataJSON) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_labeledPrice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invoice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_paymentCharge) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_postAddress) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_paymentRequestedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_paymentSavedCredentialsCard) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webDocumentNoProxy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputWebDocument) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputWebFileLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputWebFileGeoPointLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_webFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_paymentForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_validatedRequestedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_paymentResult) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_paymentVerificationNeeded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_paymentReceipt) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_savedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPaymentCredentialsSaved) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPaymentCredentials) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPaymentCredentialsApplePay) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPaymentCredentialsAndroidPay) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_tmpPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_shippingOption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputStickerSetItem) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputPhoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallWaiting) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallRequested) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallAccepted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallDiscarded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneConnection) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phoneCallProtocol) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_phoneCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_cdnFileReuploadNeeded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_cdnFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_cdnPublicKey) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_cdnConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langPackString) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langPackStringPluralized) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langPackStringDeleted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langPackDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langPackLanguage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangeTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangeAbout) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangeUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionToggleInvites) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionToggleSignatures) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionUpdatePinned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionEditMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionDeleteMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionParticipantJoin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionParticipantLeave) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionParticipantInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionParticipantToggleBan) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionParticipantToggleAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangeStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionTogglePreHistoryHidden) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionDefaultBannedRights) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionStopPoll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangeLinkedChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionChangeLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventActionToggleSlowMode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEvent) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_adminLogResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelAdminLogEventsFilter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_popularContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_favedStickersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_favedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_recentMeUrlUnknown) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_recentMeUrlUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_recentMeUrlChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_recentMeUrlChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_recentMeUrlStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_recentMeUrls) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputSingleMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_webAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessageID) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessageReplyTo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputMessagePinned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputDialogPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputDialogPeerFolder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dialogPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dialogPeerFolder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_foundStickerSetsNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_foundStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_fileHash) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputClientProxy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_termsOfServiceUpdateEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_termsOfServiceUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputSecureFileUploaded) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputSecureFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureFileEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_securePlainPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_securePlainEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypePersonalDetails) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypePassport) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeDriverLicense) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeIdentityCard) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeInternalPassport) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeAddress) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeUtilityBill) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeBankStatement) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeRentalAgreement) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypePassportRegistration) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeTemporaryRegistration) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypePhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueTypeEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueHash) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorFrontSide) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorReverseSide) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorSelfie) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorFiles) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueError) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorTranslationFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureValueErrorTranslationFiles) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureCredentialsEncrypted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_authorizationForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_sentEmailCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_deepLinkInfoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_deepLinkInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_savedPhoneContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_takeout) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_passwordKdfAlgoUnknown) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_securePasswordKdfAlgoUnknown) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_securePasswordKdfAlgoSHA512) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureSecretSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputCheckPasswordEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputCheckPasswordSRP) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureRequiredType) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_secureRequiredTypeOneOf) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_passportConfigNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_passportConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputAppEvent) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonObjectValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonNull) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonBool) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonNumber) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonString) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonArray) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_jsonObject) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageTableCell) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageTableRow) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageCaption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageListItemText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageListItemBlocks) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageListOrderedItemText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageListOrderedItemBlocks) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pageRelatedArticle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_page) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_supportName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_userInfoEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_userInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pollAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_poll) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pollAnswerVoters) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_pollResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatOnlines) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsURL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatAdminRights) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_chatBannedRights) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputWallPaperSlug) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputWallPaperNoFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_wallPapersNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_wallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_codeSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_wallPaperSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_autoDownloadSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_autoDownloadSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_emojiKeyword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_emojiKeywordDeleted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_emojiKeywordsDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_emojiURL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_emojiLanguage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_fileLocationToBeDeprecated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_folder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputFolderPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_folderPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_searchCounter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_urlAuthResultRequest) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_urlAuthResultAccepted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_urlAuthResultDefault) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelLocationEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channelLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerLocated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_peerSelfLocated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_restrictionReason) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputThemeSlug) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_theme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_themesNotModified) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_themes) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_loginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_loginTokenMigrateTo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_loginTokenSuccess) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_contentSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_inactiveChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_baseThemeClassic) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_baseThemeDay) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_baseThemeNight) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_baseThemeTinted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_baseThemeArctic) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_inputThemeSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_themeSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_webPageAttributeTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageUserVote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageUserVoteInputOption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageUserVoteMultiple) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_votesList) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_bankCardOpenUrl) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_bankCardData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dialogFilter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_dialogFilterSuggested) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsDateRangeDays) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsAbsValueAndPrev) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsPercentValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsGraphAsync) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsGraphError) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_statsGraph) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messageInteractionCounters) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stats_broadcastStats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_promoDataEmpty) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_promoData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_videoSize) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invokeAfterMsg) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invokeAfterMsgs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_initConnection) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invokeWithLayer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invokeWithoutUpdates) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invokeWithMessagesRange) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_invokeWithTakeout) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_sendCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_signUp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_signIn) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_logOut) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_resetAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_exportAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_importAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_bindTempAuthKey) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_importBotAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_checkPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_requestPasswordRecovery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_recoverPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_resendCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_cancelCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_dropTempAuthKeys) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_exportLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_importLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_auth_acceptLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_registerDevice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_unregisterDevice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updateNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_resetNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updateProfile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updateStatus) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getWallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_reportPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_checkUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updateUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getPrivacy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_setPrivacy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_deleteAccount) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getAccountTTL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_setAccountTTL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_sendChangePhoneCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_changePhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updateDeviceLocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_resetAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getPasswordSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updatePasswordSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_sendConfirmPhoneCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_confirmPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getTmpPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getWebAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_resetWebAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_resetWebAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getAllSecureValues) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_account_getSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_account_saveSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_deleteSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getAuthorizationForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_acceptAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_sendVerifyPhoneCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_verifyPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_sendVerifyEmailCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_verifyEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_initTakeoutSession) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_finishTakeoutSession) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_confirmPasswordEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_resendPasswordEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_cancelPasswordEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getContactSignUpNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_setContactSignUpNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getNotifyExceptions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_uploadWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_saveWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_installWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_resetWallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getAutoDownloadSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_saveAutoDownloadSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_uploadTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_createTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_updateTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_saveTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_installTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getThemes) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_setContentSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getContentSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_account_getMultiWallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_users_getUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_users_getFullUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_users_setSecureValueErrors) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_getContactIDs) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorInt(dbuf.VectorInt())
}

func (e TL_contacts_getStatuses) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_contacts_getContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_importContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_deleteContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_deleteByPhones) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_block) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_unblock) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_getBlocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_search) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_resolveUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_getTopPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_resetTopPeerRating) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_resetSaved) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_getSaved) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_contacts_toggleTopPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_addContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_acceptContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_contacts_getLocated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_search) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_readHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_deleteHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_deleteMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_receivedMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_setTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_forwardMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_reportSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getPeerSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_report) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getFullChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_editChatTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_editChatPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_addChatUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_deleteChatUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_createChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getDhConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_requestEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_acceptEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_discardEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_setEncryptedTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_readEncryptedHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendEncrypted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendEncryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendEncryptedService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_receivedQueue) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorLong(dbuf.VectorLong())
}

func (e TL_messages_reportEncryptedSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_readMessageContents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getAllStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getWebPagePreview) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_exportChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_checkChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_importChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_installStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_uninstallStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_startBot) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getMessagesViews) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorInt(dbuf.VectorInt())
}

func (e TL_messages_editChatAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_migrateChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_searchGlobal) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_reorderStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getDocumentByHash) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_searchGifs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getSavedGifs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_saveGif) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getInlineBotResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_setInlineBotResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendInlineBotResult) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getMessageEditData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_editMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_editInlineBotMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getBotCallbackAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_setBotCallbackAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getPeerDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_saveDraft) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getAllDrafts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_readFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getRecentStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_saveRecentSticker) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_clearRecentStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getArchivedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getMaskStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getAttachedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_setGameScore) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_setInlineGameScore) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getGameHighScores) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getInlineGameHighScores) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getCommonChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getAllChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getWebPage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_toggleDialogPin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_reorderPinnedDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getPinnedDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_setBotShippingResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_setBotPrecheckoutResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_uploadMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendScreenshotNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getFavedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_faveSticker) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getUnreadMentions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_readMentions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getRecentLocations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendMultiMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_uploadEncryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_searchStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getSplitRanges) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_markDialogUnread) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getDialogUnreadMarks) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_clearAllDrafts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_updatePinnedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendVote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getPollResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getOnlines) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getStatsURL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_editChatAbout) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_editChatDefaultBannedRights) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getEmojiKeywords) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getEmojiKeywordsDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getEmojiKeywordsLanguages) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_getEmojiURL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getSearchCounters) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_requestUrlAuth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_acceptUrlAuth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_hidePeerSettingsBar) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getScheduledHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_sendScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_deleteScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getPollVotes) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_toggleStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getDialogFilters) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_getSuggestedDialogFilters) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_messages_updateDialogFilter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_updateDialogFiltersOrder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_messages_getOldFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_getState) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_getDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_updates_getChannelDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photos_updateProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photos_uploadProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_photos_deletePhotos) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorLong(dbuf.VectorLong())
}

func (e TL_photos_getUserPhotos) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_saveFilePart) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_getFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_saveBigFilePart) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_getWebFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_getCdnFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_upload_reuploadCdnFile) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_upload_getCdnFileHashes) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_upload_getFileHashes) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_help_getConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getNearestDc) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getAppUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getInviteText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getSupport) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getAppChangelog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_setBotUpdatesStatus) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getCdnConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getRecentMeUrls) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getTermsOfServiceUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_acceptTermsOfService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getDeepLinkInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getAppConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_saveAppLog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getPassportConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getSupportName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getUserInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_editUserInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_getPromoData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_help_hidePromoData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_readHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_deleteMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_deleteUserHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_reportSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getParticipant) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getFullChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_createChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_editAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_editTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_editPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_checkUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_updateUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_joinChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_leaveChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_inviteToChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_deleteChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_exportMessageLink) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_toggleSignatures) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getAdminedPublicChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_editBanned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getAdminLog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_setStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_readMessageContents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_deleteHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_togglePreHistoryHidden) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getLeftChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getGroupsForDiscussion) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_setDiscussionGroup) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_editCreator) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_editLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_toggleSlowMode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_channels_getInactiveChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_bots_sendCustomRequest) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_bots_answerWebhookJSONQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_bots_setBotCommands) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_getPaymentForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_getPaymentReceipt) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_validateRequestedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_sendPaymentForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_getSavedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_clearSavedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_payments_getBankCardData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickers_createStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickers_removeStickerFromSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickers_changeStickerPosition) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickers_addStickerToSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stickers_setStickerSetThumb) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_getCallConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_requestCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_acceptCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_confirmCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_receivedCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_discardCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_setCallRating) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_saveCallDebug) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_phone_sendSignalingData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langpack_getLangPack) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langpack_getStrings) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_langpack_getDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_langpack_getLanguages) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e TL_langpack_getLanguage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_folders_editPeerFolders) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_folders_deleteFolder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stats_getBroadcastStats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e TL_stats_loadAsyncGraph) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func readFlags(m *DecodeBuf, flagsPtr *int32) int32 {
	flags := m.Int()
	*flagsPtr = flags
	return flags
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case CRC_resPQ:
		r = TL_resPQ{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
			m.VectorLong(),
		}

	case CRC_p_q_inner_data:
		r = TL_p_q_inner_data{
			m.String(),
			m.String(),
			m.String(),
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(32),
		}

	case CRC_p_q_inner_data_dc:
		r = TL_p_q_inner_data_dc{
			m.String(),
			m.String(),
			m.String(),
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(32),
			m.Int(),
		}

	case CRC_p_q_inner_data_temp:
		r = TL_p_q_inner_data_temp{
			m.String(),
			m.String(),
			m.String(),
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(32),
			m.Int(),
		}

	case CRC_p_q_inner_data_temp_dc:
		r = TL_p_q_inner_data_temp_dc{
			m.String(),
			m.String(),
			m.String(),
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(32),
			m.Int(),
			m.Int(),
		}

	case CRC_bind_auth_key_inner:
		r = TL_bind_auth_key_inner{
			m.Long(),
			m.Long(),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case CRC_server_DH_params_fail:
		r = TL_server_DH_params_fail{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case CRC_server_DH_params_ok:
		r = TL_server_DH_params_ok{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
		}

	case CRC_server_DH_inner_data:
		r = TL_server_DH_inner_data{
			m.Bytes(16),
			m.Bytes(16),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case CRC_client_DH_inner_data:
		r = TL_client_DH_inner_data{
			m.Bytes(16),
			m.Bytes(16),
			m.Long(),
			m.String(),
		}

	case CRC_dh_gen_ok:
		r = TL_dh_gen_ok{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case CRC_dh_gen_retry:
		r = TL_dh_gen_retry{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case CRC_dh_gen_fail:
		r = TL_dh_gen_fail{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case CRC_destroy_auth_key_ok:
		r = TL_destroy_auth_key_ok{}

	case CRC_destroy_auth_key_none:
		r = TL_destroy_auth_key_none{}

	case CRC_destroy_auth_key_fail:
		r = TL_destroy_auth_key_fail{}

	case CRC_req_pq:
		r = TL_req_pq{
			m.Bytes(16),
		}

	case CRC_req_pq_multi:
		r = TL_req_pq_multi{
			m.Bytes(16),
		}

	case CRC_req_DH_params:
		r = TL_req_DH_params{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case CRC_set_client_DH_params:
		r = TL_set_client_DH_params{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
		}

	case CRC_destroy_auth_key:
		r = TL_destroy_auth_key{}

	case CRC_msgs_ack:
		r = TL_msgs_ack{
			m.VectorLong(),
		}

	case CRC_bad_msg_notification:
		r = TL_bad_msg_notification{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CRC_bad_server_salt:
		r = TL_bad_server_salt{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Long(),
		}

	case CRC_msgs_state_req:
		r = TL_msgs_state_req{
			m.VectorLong(),
		}

	case CRC_msgs_state_info:
		r = TL_msgs_state_info{
			m.Long(),
			m.String(),
		}

	case CRC_msgs_all_info:
		r = TL_msgs_all_info{
			m.VectorLong(),
			m.String(),
		}

	case CRC_msg_detailed_info:
		r = TL_msg_detailed_info{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CRC_msg_new_detailed_info:
		r = TL_msg_new_detailed_info{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CRC_msg_resend_req:
		r = TL_msg_resend_req{
			m.VectorLong(),
		}

	case CRC_rpc_error:
		r = TL_rpc_error{
			m.Int(),
			m.String(),
		}

	case CRC_rpc_answer_unknown:
		r = TL_rpc_answer_unknown{}

	case CRC_rpc_answer_dropped_running:
		r = TL_rpc_answer_dropped_running{}

	case CRC_rpc_answer_dropped:
		r = TL_rpc_answer_dropped{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CRC_future_salt:
		r = TL_future_salt{
			m.Int(),
			m.Int(),
			m.Long(),
		}

	case CRC_future_salts:
		r = TL_future_salts{
			m.Long(),
			m.Int(),
			m.Object(),
		}

	case CRC_pong:
		r = TL_pong{
			m.Long(),
			m.Long(),
		}

	case CRC_destroy_session_ok:
		r = TL_destroy_session_ok{
			m.Long(),
		}

	case CRC_destroy_session_none:
		r = TL_destroy_session_none{
			m.Long(),
		}

	case CRC_new_session_created:
		r = TL_new_session_created{
			m.Long(),
			m.Long(),
			m.Long(),
		}

	case CRC_http_wait:
		r = TL_http_wait{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_ipPort:
		r = TL_ipPort{
			m.Int(),
			m.Int(),
		}

	case CRC_ipPortSecret:
		r = TL_ipPortSecret{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_accessPointRule:
		r = TL_accessPointRule{
			m.String(),
			m.Int(),
			m.Object(),
		}

	case CRC_help_configSimple:
		r = TL_help_configSimple{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CRC_tlsClientHello:
		r = TL_tlsClientHello{
			m.Object(),
		}

	case CRC_tlsBlockString:
		r = TL_tlsBlockString{
			m.String(),
		}

	case CRC_tlsBlockRandom:
		r = TL_tlsBlockRandom{
			m.Int(),
		}

	case CRC_tlsBlockZero:
		r = TL_tlsBlockZero{
			m.Int(),
		}

	case CRC_tlsBlockDomain:
		r = TL_tlsBlockDomain{}

	case CRC_tlsBlockGrease:
		r = TL_tlsBlockGrease{
			m.Int(),
		}

	case CRC_tlsBlockPublicKey:
		r = TL_tlsBlockPublicKey{}

	case CRC_tlsBlockScope:
		r = TL_tlsBlockScope{
			m.Vector(),
		}

	case CRC_rpc_drop_answer:
		r = TL_rpc_drop_answer{
			m.Long(),
		}

	case CRC_get_future_salts:
		r = TL_get_future_salts{
			m.Int(),
		}

	case CRC_ping:
		r = TL_ping{
			m.Long(),
		}

	case CRC_ping_delay_disconnect:
		r = TL_ping_delay_disconnect{
			m.Long(),
			m.Int(),
		}

	case CRC_destroy_session:
		r = TL_destroy_session{
			m.Long(),
		}

	case CRC_boolFalse:
		r = TL_boolFalse{}

	case CRC_boolTrue:
		r = TL_boolTrue{}

	case CRC_true:
		r = TL_true{}

	case CRC_error:
		r = TL_error{
			m.Int(),
			m.String(),
		}

	case CRC_null:
		r = TL_null{}

	case CRC_inputPeerEmpty:
		r = TL_inputPeerEmpty{}

	case CRC_inputPeerSelf:
		r = TL_inputPeerSelf{}

	case CRC_inputPeerChat:
		r = TL_inputPeerChat{
			m.Int(),
		}

	case CRC_inputPeerUser:
		r = TL_inputPeerUser{
			m.Int(),
			m.Long(),
		}

	case CRC_inputPeerChannel:
		r = TL_inputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case CRC_inputPeerUserFromMessage:
		r = TL_inputPeerUserFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_inputPeerChannelFromMessage:
		r = TL_inputPeerChannelFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_inputUserEmpty:
		r = TL_inputUserEmpty{}

	case CRC_inputUserSelf:
		r = TL_inputUserSelf{}

	case CRC_inputUser:
		r = TL_inputUser{
			m.Int(),
			m.Long(),
		}

	case CRC_inputUserFromMessage:
		r = TL_inputUserFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_inputPhoneContact:
		r = TL_inputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_inputFile:
		r = TL_inputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case CRC_inputFileBig:
		r = TL_inputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case CRC_inputMediaEmpty:
		r = TL_inputMediaEmpty{}

	case CRC_inputMediaUploadedPhoto:
		var flags int32
		r = TL_inputMediaUploadedPhoto{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case CRC_inputMediaPhoto:
		var flags int32
		r = TL_inputMediaPhoto{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_inputMediaGeoPoint:
		r = TL_inputMediaGeoPoint{
			m.Object(),
		}

	case CRC_inputMediaContact:
		r = TL_inputMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_inputMediaUploadedDocument:
		var flags int32
		r = TL_inputMediaUploadedDocument{
			readFlags(m, &flags),
			flags&8 != 0, //flag #3
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.String(),
			m.Vector(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case CRC_inputMediaDocument:
		var flags int32
		r = TL_inputMediaDocument{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_inputMediaVenue:
		r = TL_inputMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_inputMediaGifExternal:
		r = TL_inputMediaGifExternal{
			m.String(),
			m.String(),
		}

	case CRC_inputMediaPhotoExternal:
		var flags int32
		r = TL_inputMediaPhotoExternal{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_inputMediaDocumentExternal:
		var flags int32
		r = TL_inputMediaDocumentExternal{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_inputMediaGame:
		r = TL_inputMediaGame{
			m.Object(),
		}

	case CRC_inputMediaInvoice:
		var flags int32
		r = TL_inputMediaInvoice{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case CRC_inputMediaGeoLive:
		var flags int32
		r = TL_inputMediaGeoLive{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.FlaggedInt(flags, 1),
		}

	case CRC_inputMediaPoll:
		var flags int32
		r = TL_inputMediaPoll{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedVector(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedVector(flags, 1),
		}

	case CRC_inputMediaDice:
		r = TL_inputMediaDice{
			m.String(),
		}

	case CRC_inputChatPhotoEmpty:
		r = TL_inputChatPhotoEmpty{}

	case CRC_inputChatUploadedPhoto:
		r = TL_inputChatUploadedPhoto{
			m.Object(),
		}

	case CRC_inputChatPhoto:
		r = TL_inputChatPhoto{
			m.Object(),
		}

	case CRC_inputGeoPointEmpty:
		r = TL_inputGeoPointEmpty{}

	case CRC_inputGeoPoint:
		r = TL_inputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case CRC_inputPhotoEmpty:
		r = TL_inputPhotoEmpty{}

	case CRC_inputPhoto:
		r = TL_inputPhoto{
			m.Long(),
			m.Long(),
			m.StringBytes(),
		}

	case CRC_inputFileLocation:
		r = TL_inputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
			m.StringBytes(),
		}

	case CRC_inputEncryptedFileLocation:
		r = TL_inputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case CRC_inputDocumentFileLocation:
		r = TL_inputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_inputSecureFileLocation:
		r = TL_inputSecureFileLocation{
			m.Long(),
			m.Long(),
		}

	case CRC_inputTakeoutFileLocation:
		r = TL_inputTakeoutFileLocation{}

	case CRC_inputPhotoFileLocation:
		r = TL_inputPhotoFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_inputPhotoLegacyFileLocation:
		r = TL_inputPhotoLegacyFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case CRC_inputPeerPhotoFileLocation:
		var flags int32
		r = TL_inputPeerPhotoFileLocation{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case CRC_inputStickerSetThumb:
		r = TL_inputStickerSetThumb{
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case CRC_peerUser:
		r = TL_peerUser{
			m.Int(),
		}

	case CRC_peerChat:
		r = TL_peerChat{
			m.Int(),
		}

	case CRC_peerChannel:
		r = TL_peerChannel{
			m.Int(),
		}

	case CRC_storage_fileUnknown:
		r = TL_storage_fileUnknown{}

	case CRC_storage_filePartial:
		r = TL_storage_filePartial{}

	case CRC_storage_fileJpeg:
		r = TL_storage_fileJpeg{}

	case CRC_storage_fileGif:
		r = TL_storage_fileGif{}

	case CRC_storage_filePng:
		r = TL_storage_filePng{}

	case CRC_storage_filePdf:
		r = TL_storage_filePdf{}

	case CRC_storage_fileMp3:
		r = TL_storage_fileMp3{}

	case CRC_storage_fileMov:
		r = TL_storage_fileMov{}

	case CRC_storage_fileMp4:
		r = TL_storage_fileMp4{}

	case CRC_storage_fileWebp:
		r = TL_storage_fileWebp{}

	case CRC_userEmpty:
		r = TL_userEmpty{
			m.Int(),
		}

	case CRC_user:
		var flags int32
		r = TL_user{
			readFlags(m, &flags),
			flags&1024 != 0,     //flag #10
			flags&2048 != 0,     //flag #11
			flags&4096 != 0,     //flag #12
			flags&8192 != 0,     //flag #13
			flags&16384 != 0,    //flag #14
			flags&32768 != 0,    //flag #15
			flags&65536 != 0,    //flag #16
			flags&131072 != 0,   //flag #17
			flags&262144 != 0,   //flag #18
			flags&1048576 != 0,  //flag #20
			flags&2097152 != 0,  //flag #21
			flags&8388608 != 0,  //flag #23
			flags&16777216 != 0, //flag #24
			m.Int(),
			m.FlaggedLong(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 5),
			m.FlaggedObject(flags, 6),
			m.FlaggedInt(flags, 14),
			m.FlaggedVector(flags, 18),
			m.FlaggedString(flags, 19),
			m.FlaggedString(flags, 22),
		}

	case CRC_userProfilePhotoEmpty:
		r = TL_userProfilePhotoEmpty{}

	case CRC_userProfilePhoto:
		r = TL_userProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case CRC_userStatusEmpty:
		r = TL_userStatusEmpty{}

	case CRC_userStatusOnline:
		r = TL_userStatusOnline{
			m.Int(),
		}

	case CRC_userStatusOffline:
		r = TL_userStatusOffline{
			m.Int(),
		}

	case CRC_userStatusRecently:
		r = TL_userStatusRecently{}

	case CRC_userStatusLastWeek:
		r = TL_userStatusLastWeek{}

	case CRC_userStatusLastMonth:
		r = TL_userStatusLastMonth{}

	case CRC_chatEmpty:
		r = TL_chatEmpty{
			m.Int(),
		}

	case CRC_chat:
		var flags int32
		r = TL_chat{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&32 != 0, //flag #5
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 6),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 18),
		}

	case CRC_chatForbidden:
		r = TL_chatForbidden{
			m.Int(),
			m.String(),
		}

	case CRC_channel:
		var flags int32
		r = TL_channel{
			readFlags(m, &flags),
			flags&1 != 0,       //flag #0
			flags&4 != 0,       //flag #2
			flags&32 != 0,      //flag #5
			flags&128 != 0,     //flag #7
			flags&256 != 0,     //flag #8
			flags&512 != 0,     //flag #9
			flags&2048 != 0,    //flag #11
			flags&4096 != 0,    //flag #12
			flags&524288 != 0,  //flag #19
			flags&1048576 != 0, //flag #20
			flags&2097152 != 0, //flag #21
			flags&4194304 != 0, //flag #22
			m.Int(),
			m.FlaggedLong(flags, 13),
			m.String(),
			m.FlaggedString(flags, 6),
			m.Object(),
			m.Int(),
			m.Int(),
			m.FlaggedVector(flags, 9),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 15),
			m.FlaggedObject(flags, 18),
			m.FlaggedInt(flags, 17),
		}

	case CRC_channelForbidden:
		var flags int32
		r = TL_channelForbidden{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&256 != 0, //flag #8
			m.Int(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 16),
		}

	case CRC_chatFull:
		var flags int32
		r = TL_chatFull{
			readFlags(m, &flags),
			flags&128 != 0, //flag #7
			flags&256 != 0, //flag #8
			m.Int(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.Object(),
			m.Object(),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 11),
		}

	case CRC_channelFull:
		var flags int32
		r = TL_channelFull{
			readFlags(m, &flags),
			flags&8 != 0,      //flag #3
			flags&64 != 0,     //flag #6
			flags&128 != 0,    //flag #7
			flags&1024 != 0,   //flag #10
			flags&4096 != 0,   //flag #12
			flags&65536 != 0,  //flag #16
			flags&524288 != 0, //flag #19
			m.Int(),
			m.String(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 13),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
			m.FlaggedObject(flags, 8),
			m.FlaggedInt(flags, 9),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 14),
			m.FlaggedObject(flags, 15),
			m.FlaggedInt(flags, 17),
			m.FlaggedInt(flags, 18),
			m.FlaggedInt(flags, 12),
			m.Int(),
		}

	case CRC_chatParticipant:
		r = TL_chatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_chatParticipantCreator:
		r = TL_chatParticipantCreator{
			m.Int(),
		}

	case CRC_chatParticipantAdmin:
		r = TL_chatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_chatParticipantsForbidden:
		var flags int32
		r = TL_chatParticipantsForbidden{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case CRC_chatParticipants:
		r = TL_chatParticipants{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case CRC_chatPhotoEmpty:
		r = TL_chatPhotoEmpty{}

	case CRC_chatPhoto:
		r = TL_chatPhoto{
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case CRC_messageEmpty:
		r = TL_messageEmpty{
			m.Int(),
		}

	case CRC_message:
		var flags int32
		r = TL_message{
			readFlags(m, &flags),
			flags&2 != 0,       //flag #1
			flags&16 != 0,      //flag #4
			flags&32 != 0,      //flag #5
			flags&8192 != 0,    //flag #13
			flags&16384 != 0,   //flag #14
			flags&262144 != 0,  //flag #18
			flags&524288 != 0,  //flag #19
			flags&2097152 != 0, //flag #21
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 6),
			m.FlaggedVector(flags, 7),
			m.FlaggedInt(flags, 10),
			m.FlaggedInt(flags, 15),
			m.FlaggedString(flags, 16),
			m.FlaggedLong(flags, 17),
			m.FlaggedVector(flags, 22),
		}

	case CRC_messageService:
		var flags int32
		r = TL_messageService{
			readFlags(m, &flags),
			flags&2 != 0,      //flag #1
			flags&16 != 0,     //flag #4
			flags&32 != 0,     //flag #5
			flags&8192 != 0,   //flag #13
			flags&16384 != 0,  //flag #14
			flags&524288 != 0, //flag #19
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.Object(),
		}

	case CRC_messageMediaEmpty:
		r = TL_messageMediaEmpty{}

	case CRC_messageMediaPhoto:
		var flags int32
		r = TL_messageMediaPhoto{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
		}

	case CRC_messageMediaGeo:
		r = TL_messageMediaGeo{
			m.Object(),
		}

	case CRC_messageMediaContact:
		r = TL_messageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case CRC_messageMediaUnsupported:
		r = TL_messageMediaUnsupported{}

	case CRC_messageMediaDocument:
		var flags int32
		r = TL_messageMediaDocument{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
		}

	case CRC_messageMediaWebPage:
		r = TL_messageMediaWebPage{
			m.Object(),
		}

	case CRC_messageMediaVenue:
		r = TL_messageMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_messageMediaGame:
		r = TL_messageMediaGame{
			m.Object(),
		}

	case CRC_messageMediaInvoice:
		var flags int32
		r = TL_messageMediaInvoice{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			flags&8 != 0, //flag #3
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
			m.String(),
			m.Long(),
			m.String(),
		}

	case CRC_messageMediaGeoLive:
		r = TL_messageMediaGeoLive{
			m.Object(),
			m.Int(),
		}

	case CRC_messageMediaPoll:
		r = TL_messageMediaPoll{
			m.Object(),
			m.Object(),
		}

	case CRC_messageMediaDice:
		r = TL_messageMediaDice{
			m.Int(),
			m.String(),
		}

	case CRC_messageActionEmpty:
		r = TL_messageActionEmpty{}

	case CRC_messageActionChatCreate:
		r = TL_messageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case CRC_messageActionChatEditTitle:
		r = TL_messageActionChatEditTitle{
			m.String(),
		}

	case CRC_messageActionChatEditPhoto:
		r = TL_messageActionChatEditPhoto{
			m.Object(),
		}

	case CRC_messageActionChatDeletePhoto:
		r = TL_messageActionChatDeletePhoto{}

	case CRC_messageActionChatAddUser:
		r = TL_messageActionChatAddUser{
			m.VectorInt(),
		}

	case CRC_messageActionChatDeleteUser:
		r = TL_messageActionChatDeleteUser{
			m.Int(),
		}

	case CRC_messageActionChatJoinedByLink:
		r = TL_messageActionChatJoinedByLink{
			m.Int(),
		}

	case CRC_messageActionChannelCreate:
		r = TL_messageActionChannelCreate{
			m.String(),
		}

	case CRC_messageActionChatMigrateTo:
		r = TL_messageActionChatMigrateTo{
			m.Int(),
		}

	case CRC_messageActionChannelMigrateFrom:
		r = TL_messageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case CRC_messageActionPinMessage:
		r = TL_messageActionPinMessage{}

	case CRC_messageActionHistoryClear:
		r = TL_messageActionHistoryClear{}

	case CRC_messageActionGameScore:
		r = TL_messageActionGameScore{
			m.Long(),
			m.Int(),
		}

	case CRC_messageActionPaymentSentMe:
		var flags int32
		r = TL_messageActionPaymentSentMe{
			readFlags(m, &flags),
			m.String(),
			m.Long(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case CRC_messageActionPaymentSent:
		r = TL_messageActionPaymentSent{
			m.String(),
			m.Long(),
		}

	case CRC_messageActionPhoneCall:
		var flags int32
		r = TL_messageActionPhoneCall{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case CRC_messageActionScreenshotTaken:
		r = TL_messageActionScreenshotTaken{}

	case CRC_messageActionCustomAction:
		r = TL_messageActionCustomAction{
			m.String(),
		}

	case CRC_messageActionBotAllowed:
		r = TL_messageActionBotAllowed{
			m.String(),
		}

	case CRC_messageActionSecureValuesSentMe:
		r = TL_messageActionSecureValuesSentMe{
			m.Vector(),
			m.Object(),
		}

	case CRC_messageActionSecureValuesSent:
		r = TL_messageActionSecureValuesSent{
			m.Vector(),
		}

	case CRC_messageActionContactSignUp:
		r = TL_messageActionContactSignUp{}

	case CRC_dialog:
		var flags int32
		r = TL_dialog{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 4),
		}

	case CRC_dialogFolder:
		var flags int32
		r = TL_dialogFolder{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_photoEmpty:
		r = TL_photoEmpty{
			m.Long(),
		}

	case CRC_photo:
		var flags int32
		r = TL_photo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case CRC_photoSizeEmpty:
		r = TL_photoSizeEmpty{
			m.String(),
		}

	case CRC_photoSize:
		r = TL_photoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_photoCachedSize:
		r = TL_photoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_photoStrippedSize:
		r = TL_photoStrippedSize{
			m.String(),
			m.StringBytes(),
		}

	case CRC_geoPointEmpty:
		r = TL_geoPointEmpty{}

	case CRC_geoPoint:
		r = TL_geoPoint{
			m.Double(),
			m.Double(),
			m.Long(),
		}

	case CRC_auth_sentCode:
		var flags int32
		r = TL_auth_sentCode{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case CRC_auth_authorization:
		var flags int32
		r = TL_auth_authorization{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Object(),
		}

	case CRC_auth_authorizationSignUpRequired:
		var flags int32
		r = TL_auth_authorizationSignUpRequired{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
		}

	case CRC_auth_exportedAuthorization:
		r = TL_auth_exportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case CRC_inputNotifyPeer:
		r = TL_inputNotifyPeer{
			m.Object(),
		}

	case CRC_inputNotifyUsers:
		r = TL_inputNotifyUsers{}

	case CRC_inputNotifyChats:
		r = TL_inputNotifyChats{}

	case CRC_inputNotifyBroadcasts:
		r = TL_inputNotifyBroadcasts{}

	case CRC_inputPeerNotifySettings:
		var flags int32
		r = TL_inputPeerNotifySettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case CRC_peerNotifySettings:
		var flags int32
		r = TL_peerNotifySettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case CRC_peerSettings:
		var flags int32
		r = TL_peerSettings{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			flags&32 != 0, //flag #5
		}

	case CRC_wallPaper:
		var flags int32
		r = TL_wallPaper{
			m.Long(),
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.Long(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_wallPaperNoFile:
		var flags int32
		r = TL_wallPaperNoFile{
			readFlags(m, &flags),
			flags&2 != 0,  //flag #1
			flags&16 != 0, //flag #4
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputReportReasonSpam:
		r = TL_inputReportReasonSpam{}

	case CRC_inputReportReasonViolence:
		r = TL_inputReportReasonViolence{}

	case CRC_inputReportReasonPornography:
		r = TL_inputReportReasonPornography{}

	case CRC_inputReportReasonChildAbuse:
		r = TL_inputReportReasonChildAbuse{}

	case CRC_inputReportReasonOther:
		r = TL_inputReportReasonOther{
			m.String(),
		}

	case CRC_inputReportReasonCopyright:
		r = TL_inputReportReasonCopyright{}

	case CRC_inputReportReasonGeoIrrelevant:
		r = TL_inputReportReasonGeoIrrelevant{}

	case CRC_userFull:
		var flags int32
		r = TL_userFull{
			readFlags(m, &flags),
			flags&1 != 0,    //flag #0
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&128 != 0,  //flag #7
			flags&4096 != 0, //flag #12
			m.Object(),
			m.FlaggedString(flags, 1),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.Object(),
			m.FlaggedObject(flags, 3),
			m.FlaggedInt(flags, 6),
			m.Int(),
			m.FlaggedInt(flags, 11),
		}

	case CRC_contact:
		r = TL_contact{
			m.Int(),
			m.Object(),
		}

	case CRC_importedContact:
		r = TL_importedContact{
			m.Int(),
			m.Long(),
		}

	case CRC_contactBlocked:
		r = TL_contactBlocked{
			m.Int(),
			m.Int(),
		}

	case CRC_contactStatus:
		r = TL_contactStatus{
			m.Int(),
			m.Object(),
		}

	case CRC_contacts_contactsNotModified:
		r = TL_contacts_contactsNotModified{}

	case CRC_contacts_contacts:
		r = TL_contacts_contacts{
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case CRC_contacts_importedContacts:
		r = TL_contacts_importedContacts{
			m.Vector(),
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case CRC_contacts_blocked:
		r = TL_contacts_blocked{
			m.Vector(),
			m.Vector(),
		}

	case CRC_contacts_blockedSlice:
		r = TL_contacts_blockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_dialogs:
		r = TL_messages_dialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_dialogsSlice:
		r = TL_messages_dialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_dialogsNotModified:
		r = TL_messages_dialogsNotModified{
			m.Int(),
		}

	case CRC_messages_messages:
		r = TL_messages_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_messagesSlice:
		var flags int32
		r = TL_messages_messagesSlice{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_channelMessages:
		var flags int32
		r = TL_messages_channelMessages{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_messagesNotModified:
		r = TL_messages_messagesNotModified{
			m.Int(),
		}

	case CRC_messages_chats:
		r = TL_messages_chats{
			m.Vector(),
		}

	case CRC_messages_chatsSlice:
		r = TL_messages_chatsSlice{
			m.Int(),
			m.Vector(),
		}

	case CRC_messages_chatFull:
		r = TL_messages_chatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messages_affectedHistory:
		r = TL_messages_affectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_inputMessagesFilterEmpty:
		r = TL_inputMessagesFilterEmpty{}

	case CRC_inputMessagesFilterPhotos:
		r = TL_inputMessagesFilterPhotos{}

	case CRC_inputMessagesFilterVideo:
		r = TL_inputMessagesFilterVideo{}

	case CRC_inputMessagesFilterPhotoVideo:
		r = TL_inputMessagesFilterPhotoVideo{}

	case CRC_inputMessagesFilterDocument:
		r = TL_inputMessagesFilterDocument{}

	case CRC_inputMessagesFilterUrl:
		r = TL_inputMessagesFilterUrl{}

	case CRC_inputMessagesFilterGif:
		r = TL_inputMessagesFilterGif{}

	case CRC_inputMessagesFilterVoice:
		r = TL_inputMessagesFilterVoice{}

	case CRC_inputMessagesFilterMusic:
		r = TL_inputMessagesFilterMusic{}

	case CRC_inputMessagesFilterChatPhotos:
		r = TL_inputMessagesFilterChatPhotos{}

	case CRC_inputMessagesFilterPhoneCalls:
		var flags int32
		r = TL_inputMessagesFilterPhoneCalls{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case CRC_inputMessagesFilterRoundVoice:
		r = TL_inputMessagesFilterRoundVoice{}

	case CRC_inputMessagesFilterRoundVideo:
		r = TL_inputMessagesFilterRoundVideo{}

	case CRC_inputMessagesFilterMyMentions:
		r = TL_inputMessagesFilterMyMentions{}

	case CRC_inputMessagesFilterGeo:
		r = TL_inputMessagesFilterGeo{}

	case CRC_inputMessagesFilterContacts:
		r = TL_inputMessagesFilterContacts{}

	case CRC_updateNewMessage:
		r = TL_updateNewMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateMessageID:
		r = TL_updateMessageID{
			m.Int(),
			m.Long(),
		}

	case CRC_updateDeleteMessages:
		r = TL_updateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateUserTyping:
		r = TL_updateUserTyping{
			m.Int(),
			m.Object(),
		}

	case CRC_updateChatUserTyping:
		r = TL_updateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CRC_updateChatParticipants:
		r = TL_updateChatParticipants{
			m.Object(),
		}

	case CRC_updateUserStatus:
		r = TL_updateUserStatus{
			m.Int(),
			m.Object(),
		}

	case CRC_updateUserName:
		r = TL_updateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_updateUserPhoto:
		r = TL_updateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case CRC_updateNewEncryptedMessage:
		r = TL_updateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case CRC_updateEncryptedChatTyping:
		r = TL_updateEncryptedChatTyping{
			m.Int(),
		}

	case CRC_updateEncryption:
		r = TL_updateEncryption{
			m.Object(),
			m.Int(),
		}

	case CRC_updateEncryptedMessagesRead:
		r = TL_updateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateChatParticipantAdd:
		r = TL_updateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateChatParticipantDelete:
		r = TL_updateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateDcOptions:
		r = TL_updateDcOptions{
			m.Vector(),
		}

	case CRC_updateUserBlocked:
		r = TL_updateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case CRC_updateNotifySettings:
		r = TL_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case CRC_updateServiceNotification:
		var flags int32
		r = TL_updateServiceNotification{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.String(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case CRC_updatePrivacy:
		r = TL_updatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case CRC_updateUserPhone:
		r = TL_updateUserPhone{
			m.Int(),
			m.String(),
		}

	case CRC_updateReadHistoryInbox:
		var flags int32
		r = TL_updateReadHistoryInbox{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateReadHistoryOutbox:
		r = TL_updateReadHistoryOutbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateWebPage:
		r = TL_updateWebPage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateReadMessagesContents:
		r = TL_updateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateChannelTooLong:
		var flags int32
		r = TL_updateChannelTooLong{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_updateChannel:
		r = TL_updateChannel{
			m.Int(),
		}

	case CRC_updateNewChannelMessage:
		r = TL_updateNewChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateReadChannelInbox:
		var flags int32
		r = TL_updateReadChannelInbox{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateDeleteChannelMessages:
		r = TL_updateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateChannelMessageViews:
		r = TL_updateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateChatParticipantAdmin:
		r = TL_updateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case CRC_updateNewStickerSet:
		r = TL_updateNewStickerSet{
			m.Object(),
		}

	case CRC_updateStickerSetsOrder:
		var flags int32
		r = TL_updateStickerSetsOrder{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.VectorLong(),
		}

	case CRC_updateStickerSets:
		r = TL_updateStickerSets{}

	case CRC_updateSavedGifs:
		r = TL_updateSavedGifs{}

	case CRC_updateBotInlineQuery:
		var flags int32
		r = TL_updateBotInlineQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
		}

	case CRC_updateBotInlineSend:
		var flags int32
		r = TL_updateBotInlineSend{
			readFlags(m, &flags),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.FlaggedObject(flags, 1),
		}

	case CRC_updateEditChannelMessage:
		r = TL_updateEditChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateChannelPinnedMessage:
		r = TL_updateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case CRC_updateBotCallbackQuery:
		var flags int32
		r = TL_updateBotCallbackQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case CRC_updateEditMessage:
		r = TL_updateEditMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateInlineBotCallbackQuery:
		var flags int32
		r = TL_updateInlineBotCallbackQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.Object(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case CRC_updateReadChannelOutbox:
		r = TL_updateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case CRC_updateDraftMessage:
		r = TL_updateDraftMessage{
			m.Object(),
			m.Object(),
		}

	case CRC_updateReadFeaturedStickers:
		r = TL_updateReadFeaturedStickers{}

	case CRC_updateRecentStickers:
		r = TL_updateRecentStickers{}

	case CRC_updateConfig:
		r = TL_updateConfig{}

	case CRC_updatePtsChanged:
		r = TL_updatePtsChanged{}

	case CRC_updateChannelWebPage:
		r = TL_updateChannelWebPage{
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateDialogPinned:
		var flags int32
		r = TL_updateDialogPinned{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.Object(),
		}

	case CRC_updatePinnedDialogs:
		var flags int32
		r = TL_updatePinnedDialogs{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 1),
			m.FlaggedVector(flags, 0),
		}

	case CRC_updateBotWebhookJSON:
		r = TL_updateBotWebhookJSON{
			m.Object(),
		}

	case CRC_updateBotWebhookJSONQuery:
		r = TL_updateBotWebhookJSONQuery{
			m.Long(),
			m.Object(),
			m.Int(),
		}

	case CRC_updateBotShippingQuery:
		r = TL_updateBotShippingQuery{
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_updateBotPrecheckoutQuery:
		var flags int32
		r = TL_updateBotPrecheckoutQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Long(),
		}

	case CRC_updatePhoneCall:
		r = TL_updatePhoneCall{
			m.Object(),
		}

	case CRC_updateLangPackTooLong:
		r = TL_updateLangPackTooLong{
			m.String(),
		}

	case CRC_updateLangPack:
		r = TL_updateLangPack{
			m.Object(),
		}

	case CRC_updateFavedStickers:
		r = TL_updateFavedStickers{}

	case CRC_updateChannelReadMessagesContents:
		r = TL_updateChannelReadMessagesContents{
			m.Int(),
			m.VectorInt(),
		}

	case CRC_updateContactsReset:
		r = TL_updateContactsReset{}

	case CRC_updateChannelAvailableMessages:
		r = TL_updateChannelAvailableMessages{
			m.Int(),
			m.Int(),
		}

	case CRC_updateDialogUnreadMark:
		var flags int32
		r = TL_updateDialogUnreadMark{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case CRC_updateUserPinnedMessage:
		r = TL_updateUserPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case CRC_updateChatPinnedMessage:
		r = TL_updateChatPinnedMessage{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateMessagePoll:
		var flags int32
		r = TL_updateMessagePoll{
			readFlags(m, &flags),
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.Object(),
		}

	case CRC_updateChatDefaultBannedRights:
		r = TL_updateChatDefaultBannedRights{
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case CRC_updateFolderPeers:
		r = TL_updateFolderPeers{
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CRC_updatePeerSettings:
		r = TL_updatePeerSettings{
			m.Object(),
			m.Object(),
		}

	case CRC_updatePeerLocated:
		r = TL_updatePeerLocated{
			m.Vector(),
		}

	case CRC_updateNewScheduledMessage:
		r = TL_updateNewScheduledMessage{
			m.Object(),
		}

	case CRC_updateDeleteScheduledMessages:
		r = TL_updateDeleteScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case CRC_updateTheme:
		r = TL_updateTheme{
			m.Object(),
		}

	case CRC_updateGeoLiveViewed:
		r = TL_updateGeoLiveViewed{
			m.Object(),
			m.Int(),
		}

	case CRC_updateLoginToken:
		r = TL_updateLoginToken{}

	case CRC_updateMessagePollVote:
		r = TL_updateMessagePollVote{
			m.Long(),
			m.Int(),
			m.Vector(),
		}

	case CRC_updateDialogFilter:
		var flags int32
		r = TL_updateDialogFilter{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case CRC_updateDialogFilterOrder:
		r = TL_updateDialogFilterOrder{
			m.VectorInt(),
		}

	case CRC_updateDialogFilters:
		r = TL_updateDialogFilters{}

	case CRC_updatePhoneCallSignalingData:
		r = TL_updatePhoneCallSignalingData{
			m.Long(),
			m.StringBytes(),
		}

	case CRC_updates_state:
		r = TL_updates_state{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updates_differenceEmpty:
		r = TL_updates_differenceEmpty{
			m.Int(),
			m.Int(),
		}

	case CRC_updates_difference:
		r = TL_updates_difference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case CRC_updates_differenceSlice:
		r = TL_updates_differenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case CRC_updates_differenceTooLong:
		r = TL_updates_differenceTooLong{
			m.Int(),
		}

	case CRC_updatesTooLong:
		r = TL_updatesTooLong{}

	case CRC_updateShortMessage:
		var flags int32
		r = TL_updateShortMessage{
			readFlags(m, &flags),
			flags&2 != 0,    //flag #1
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&8192 != 0, //flag #13
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case CRC_updateShortChatMessage:
		var flags int32
		r = TL_updateShortChatMessage{
			readFlags(m, &flags),
			flags&2 != 0,    //flag #1
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&8192 != 0, //flag #13
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case CRC_updateShort:
		r = TL_updateShort{
			m.Object(),
			m.Int(),
		}

	case CRC_updatesCombined:
		r = TL_updatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updates:
		r = TL_updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case CRC_updateShortSentMessage:
		var flags int32
		r = TL_updateShortSentMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 9),
			m.FlaggedVector(flags, 7),
		}

	case CRC_photos_photos:
		r = TL_photos_photos{
			m.Vector(),
			m.Vector(),
		}

	case CRC_photos_photosSlice:
		r = TL_photos_photosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_photos_photo:
		r = TL_photos_photo{
			m.Object(),
			m.Vector(),
		}

	case CRC_upload_file:
		r = TL_upload_file{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_upload_fileCdnRedirect:
		r = TL_upload_fileCdnRedirect{
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
			m.Vector(),
		}

	case CRC_dcOption:
		var flags int32
		r = TL_dcOption{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.Int(),
			m.String(),
			m.Int(),
			m.FlaggedStringBytes(flags, 10),
		}

	case CRC_config:
		var flags int32
		r = TL_config{
			readFlags(m, &flags),
			flags&2 != 0,    //flag #1
			flags&8 != 0,    //flag #3
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&64 != 0,   //flag #6
			flags&256 != 0,  //flag #8
			flags&8192 != 0, //flag #13
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 7),
			m.FlaggedString(flags, 9),
			m.FlaggedString(flags, 10),
			m.FlaggedString(flags, 11),
			m.FlaggedString(flags, 12),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedString(flags, 2),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 2),
		}

	case CRC_nearestDc:
		r = TL_nearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case CRC_help_appUpdate:
		var flags int32
		r = TL_help_appUpdate{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.String(),
			m.String(),
			m.Vector(),
			m.FlaggedObject(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case CRC_help_noAppUpdate:
		r = TL_help_noAppUpdate{}

	case CRC_help_inviteText:
		r = TL_help_inviteText{
			m.String(),
		}

	case CRC_encryptedChatEmpty:
		r = TL_encryptedChatEmpty{
			m.Int(),
		}

	case CRC_encryptedChatWaiting:
		r = TL_encryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_encryptedChatRequested:
		r = TL_encryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_encryptedChat:
		r = TL_encryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case CRC_encryptedChatDiscarded:
		r = TL_encryptedChatDiscarded{
			m.Int(),
		}

	case CRC_inputEncryptedChat:
		r = TL_inputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case CRC_encryptedFileEmpty:
		r = TL_encryptedFileEmpty{}

	case CRC_encryptedFile:
		r = TL_encryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_inputEncryptedFileEmpty:
		r = TL_inputEncryptedFileEmpty{}

	case CRC_inputEncryptedFileUploaded:
		r = TL_inputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case CRC_inputEncryptedFile:
		r = TL_inputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case CRC_inputEncryptedFileBigUploaded:
		r = TL_inputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CRC_encryptedMessage:
		r = TL_encryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_encryptedMessageService:
		r = TL_encryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_messages_dhConfigNotModified:
		r = TL_messages_dhConfigNotModified{
			m.StringBytes(),
		}

	case CRC_messages_dhConfig:
		r = TL_messages_dhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_messages_sentEncryptedMessage:
		r = TL_messages_sentEncryptedMessage{
			m.Int(),
		}

	case CRC_messages_sentEncryptedFile:
		r = TL_messages_sentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case CRC_inputDocumentEmpty:
		r = TL_inputDocumentEmpty{}

	case CRC_inputDocument:
		r = TL_inputDocument{
			m.Long(),
			m.Long(),
			m.StringBytes(),
		}

	case CRC_documentEmpty:
		r = TL_documentEmpty{
			m.Long(),
		}

	case CRC_document:
		var flags int32
		r = TL_document{
			readFlags(m, &flags),
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Int(),
			m.String(),
			m.Int(),
			m.FlaggedVector(flags, 0),
			m.FlaggedVector(flags, 1),
			m.Int(),
			m.Vector(),
		}

	case CRC_help_support:
		r = TL_help_support{
			m.String(),
			m.Object(),
		}

	case CRC_notifyPeer:
		r = TL_notifyPeer{
			m.Object(),
		}

	case CRC_notifyUsers:
		r = TL_notifyUsers{}

	case CRC_notifyChats:
		r = TL_notifyChats{}

	case CRC_notifyBroadcasts:
		r = TL_notifyBroadcasts{}

	case CRC_sendMessageTypingAction:
		r = TL_sendMessageTypingAction{}

	case CRC_sendMessageCancelAction:
		r = TL_sendMessageCancelAction{}

	case CRC_sendMessageRecordVideoAction:
		r = TL_sendMessageRecordVideoAction{}

	case CRC_sendMessageUploadVideoAction:
		r = TL_sendMessageUploadVideoAction{
			m.Int(),
		}

	case CRC_sendMessageRecordAudioAction:
		r = TL_sendMessageRecordAudioAction{}

	case CRC_sendMessageUploadAudioAction:
		r = TL_sendMessageUploadAudioAction{
			m.Int(),
		}

	case CRC_sendMessageUploadPhotoAction:
		r = TL_sendMessageUploadPhotoAction{
			m.Int(),
		}

	case CRC_sendMessageUploadDocumentAction:
		r = TL_sendMessageUploadDocumentAction{
			m.Int(),
		}

	case CRC_sendMessageGeoLocationAction:
		r = TL_sendMessageGeoLocationAction{}

	case CRC_sendMessageChooseContactAction:
		r = TL_sendMessageChooseContactAction{}

	case CRC_sendMessageGamePlayAction:
		r = TL_sendMessageGamePlayAction{}

	case CRC_sendMessageRecordRoundAction:
		r = TL_sendMessageRecordRoundAction{}

	case CRC_sendMessageUploadRoundAction:
		r = TL_sendMessageUploadRoundAction{
			m.Int(),
		}

	case CRC_contacts_found:
		r = TL_contacts_found{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_inputPrivacyKeyStatusTimestamp:
		r = TL_inputPrivacyKeyStatusTimestamp{}

	case CRC_inputPrivacyKeyChatInvite:
		r = TL_inputPrivacyKeyChatInvite{}

	case CRC_inputPrivacyKeyPhoneCall:
		r = TL_inputPrivacyKeyPhoneCall{}

	case CRC_inputPrivacyKeyPhoneP2P:
		r = TL_inputPrivacyKeyPhoneP2P{}

	case CRC_inputPrivacyKeyForwards:
		r = TL_inputPrivacyKeyForwards{}

	case CRC_inputPrivacyKeyProfilePhoto:
		r = TL_inputPrivacyKeyProfilePhoto{}

	case CRC_inputPrivacyKeyPhoneNumber:
		r = TL_inputPrivacyKeyPhoneNumber{}

	case CRC_inputPrivacyKeyAddedByPhone:
		r = TL_inputPrivacyKeyAddedByPhone{}

	case CRC_privacyKeyStatusTimestamp:
		r = TL_privacyKeyStatusTimestamp{}

	case CRC_privacyKeyChatInvite:
		r = TL_privacyKeyChatInvite{}

	case CRC_privacyKeyPhoneCall:
		r = TL_privacyKeyPhoneCall{}

	case CRC_privacyKeyPhoneP2P:
		r = TL_privacyKeyPhoneP2P{}

	case CRC_privacyKeyForwards:
		r = TL_privacyKeyForwards{}

	case CRC_privacyKeyProfilePhoto:
		r = TL_privacyKeyProfilePhoto{}

	case CRC_privacyKeyPhoneNumber:
		r = TL_privacyKeyPhoneNumber{}

	case CRC_privacyKeyAddedByPhone:
		r = TL_privacyKeyAddedByPhone{}

	case CRC_inputPrivacyValueAllowContacts:
		r = TL_inputPrivacyValueAllowContacts{}

	case CRC_inputPrivacyValueAllowAll:
		r = TL_inputPrivacyValueAllowAll{}

	case CRC_inputPrivacyValueAllowUsers:
		r = TL_inputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case CRC_inputPrivacyValueDisallowContacts:
		r = TL_inputPrivacyValueDisallowContacts{}

	case CRC_inputPrivacyValueDisallowAll:
		r = TL_inputPrivacyValueDisallowAll{}

	case CRC_inputPrivacyValueDisallowUsers:
		r = TL_inputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case CRC_inputPrivacyValueAllowChatParticipants:
		r = TL_inputPrivacyValueAllowChatParticipants{
			m.VectorInt(),
		}

	case CRC_inputPrivacyValueDisallowChatParticipants:
		r = TL_inputPrivacyValueDisallowChatParticipants{
			m.VectorInt(),
		}

	case CRC_privacyValueAllowContacts:
		r = TL_privacyValueAllowContacts{}

	case CRC_privacyValueAllowAll:
		r = TL_privacyValueAllowAll{}

	case CRC_privacyValueAllowUsers:
		r = TL_privacyValueAllowUsers{
			m.VectorInt(),
		}

	case CRC_privacyValueDisallowContacts:
		r = TL_privacyValueDisallowContacts{}

	case CRC_privacyValueDisallowAll:
		r = TL_privacyValueDisallowAll{}

	case CRC_privacyValueDisallowUsers:
		r = TL_privacyValueDisallowUsers{
			m.VectorInt(),
		}

	case CRC_privacyValueAllowChatParticipants:
		r = TL_privacyValueAllowChatParticipants{
			m.VectorInt(),
		}

	case CRC_privacyValueDisallowChatParticipants:
		r = TL_privacyValueDisallowChatParticipants{
			m.VectorInt(),
		}

	case CRC_account_privacyRules:
		r = TL_account_privacyRules{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_accountDaysTTL:
		r = TL_accountDaysTTL{
			m.Int(),
		}

	case CRC_documentAttributeImageSize:
		r = TL_documentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case CRC_documentAttributeAnimated:
		r = TL_documentAttributeAnimated{}

	case CRC_documentAttributeSticker:
		var flags int32
		r = TL_documentAttributeSticker{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case CRC_documentAttributeVideo:
		var flags int32
		r = TL_documentAttributeVideo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_documentAttributeAudio:
		var flags int32
		r = TL_documentAttributeAudio{
			readFlags(m, &flags),
			flags&1024 != 0, //flag #10
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedStringBytes(flags, 2),
		}

	case CRC_documentAttributeFilename:
		r = TL_documentAttributeFilename{
			m.String(),
		}

	case CRC_documentAttributeHasStickers:
		r = TL_documentAttributeHasStickers{}

	case CRC_messages_stickersNotModified:
		r = TL_messages_stickersNotModified{}

	case CRC_messages_stickers:
		r = TL_messages_stickers{
			m.Int(),
			m.Vector(),
		}

	case CRC_stickerPack:
		r = TL_stickerPack{
			m.String(),
			m.VectorLong(),
		}

	case CRC_messages_allStickersNotModified:
		r = TL_messages_allStickersNotModified{}

	case CRC_messages_allStickers:
		r = TL_messages_allStickers{
			m.Int(),
			m.Vector(),
		}

	case CRC_messages_affectedMessages:
		r = TL_messages_affectedMessages{
			m.Int(),
			m.Int(),
		}

	case CRC_webPageEmpty:
		r = TL_webPageEmpty{
			m.Long(),
		}

	case CRC_webPagePending:
		r = TL_webPagePending{
			m.Long(),
			m.Int(),
		}

	case CRC_webPage:
		var flags int32
		r = TL_webPage{
			readFlags(m, &flags),
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.FlaggedString(flags, 8),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 10),
			m.FlaggedVector(flags, 12),
		}

	case CRC_webPageNotModified:
		var flags int32
		r = TL_webPageNotModified{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
		}

	case CRC_authorization:
		var flags int32
		r = TL_authorization{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_account_authorizations:
		r = TL_account_authorizations{
			m.Vector(),
		}

	case CRC_account_password:
		var flags int32
		r = TL_account_password{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.FlaggedObject(flags, 2),
			m.FlaggedStringBytes(flags, 2),
			m.FlaggedLong(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.Object(),
			m.Object(),
			m.StringBytes(),
		}

	case CRC_account_passwordSettings:
		var flags int32
		r = TL_account_passwordSettings{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case CRC_account_passwordInputSettings:
		var flags int32
		r = TL_account_passwordInputSettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case CRC_auth_passwordRecovery:
		r = TL_auth_passwordRecovery{
			m.String(),
		}

	case CRC_receivedNotifyMessage:
		r = TL_receivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case CRC_chatInviteEmpty:
		r = TL_chatInviteEmpty{}

	case CRC_chatInviteExported:
		r = TL_chatInviteExported{
			m.String(),
		}

	case CRC_chatInviteAlready:
		r = TL_chatInviteAlready{
			m.Object(),
		}

	case CRC_chatInvite:
		var flags int32
		r = TL_chatInvite{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.String(),
			m.Object(),
			m.Int(),
			m.FlaggedVector(flags, 4),
		}

	case CRC_inputStickerSetEmpty:
		r = TL_inputStickerSetEmpty{}

	case CRC_inputStickerSetID:
		r = TL_inputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case CRC_inputStickerSetShortName:
		r = TL_inputStickerSetShortName{
			m.String(),
		}

	case CRC_inputStickerSetAnimatedEmoji:
		r = TL_inputStickerSetAnimatedEmoji{}

	case CRC_inputStickerSetDice:
		r = TL_inputStickerSetDice{
			m.String(),
		}

	case CRC_stickerSet:
		var flags int32
		r = TL_stickerSet{
			readFlags(m, &flags),
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&32 != 0, //flag #5
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 4),
			m.FlaggedInt(flags, 4),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_stickerSet:
		r = TL_messages_stickerSet{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_botCommand:
		r = TL_botCommand{
			m.String(),
			m.String(),
		}

	case CRC_botInfo:
		r = TL_botInfo{
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case CRC_keyboardButton:
		r = TL_keyboardButton{
			m.String(),
		}

	case CRC_keyboardButtonUrl:
		r = TL_keyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case CRC_keyboardButtonCallback:
		r = TL_keyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case CRC_keyboardButtonRequestPhone:
		r = TL_keyboardButtonRequestPhone{
			m.String(),
		}

	case CRC_keyboardButtonRequestGeoLocation:
		r = TL_keyboardButtonRequestGeoLocation{
			m.String(),
		}

	case CRC_keyboardButtonSwitchInline:
		var flags int32
		r = TL_keyboardButtonSwitchInline{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.String(),
		}

	case CRC_keyboardButtonGame:
		r = TL_keyboardButtonGame{
			m.String(),
		}

	case CRC_keyboardButtonBuy:
		r = TL_keyboardButtonBuy{
			m.String(),
		}

	case CRC_keyboardButtonUrlAuth:
		var flags int32
		r = TL_keyboardButtonUrlAuth{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedString(flags, 0),
			m.String(),
			m.Int(),
		}

	case CRC_inputKeyboardButtonUrlAuth:
		var flags int32
		r = TL_inputKeyboardButtonUrlAuth{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Object(),
		}

	case CRC_keyboardButtonRequestPoll:
		var flags int32
		r = TL_keyboardButtonRequestPoll{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.String(),
		}

	case CRC_keyboardButtonRow:
		r = TL_keyboardButtonRow{
			m.Vector(),
		}

	case CRC_replyKeyboardHide:
		var flags int32
		r = TL_replyKeyboardHide{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
		}

	case CRC_replyKeyboardForceReply:
		var flags int32
		r = TL_replyKeyboardForceReply{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
		}

	case CRC_replyKeyboardMarkup:
		var flags int32
		r = TL_replyKeyboardMarkup{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Vector(),
		}

	case CRC_replyInlineMarkup:
		r = TL_replyInlineMarkup{
			m.Vector(),
		}

	case CRC_messageEntityUnknown:
		r = TL_messageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityMention:
		r = TL_messageEntityMention{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityHashtag:
		r = TL_messageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityBotCommand:
		r = TL_messageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityUrl:
		r = TL_messageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityEmail:
		r = TL_messageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityBold:
		r = TL_messageEntityBold{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityItalic:
		r = TL_messageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityCode:
		r = TL_messageEntityCode{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityPre:
		r = TL_messageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case CRC_messageEntityTextUrl:
		r = TL_messageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case CRC_messageEntityMentionName:
		r = TL_messageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_inputMessageEntityMentionName:
		r = TL_inputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CRC_messageEntityPhone:
		r = TL_messageEntityPhone{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityCashtag:
		r = TL_messageEntityCashtag{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityUnderline:
		r = TL_messageEntityUnderline{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityStrike:
		r = TL_messageEntityStrike{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityBlockquote:
		r = TL_messageEntityBlockquote{
			m.Int(),
			m.Int(),
		}

	case CRC_messageEntityBankCard:
		r = TL_messageEntityBankCard{
			m.Int(),
			m.Int(),
		}

	case CRC_inputChannelEmpty:
		r = TL_inputChannelEmpty{}

	case CRC_inputChannel:
		r = TL_inputChannel{
			m.Int(),
			m.Long(),
		}

	case CRC_inputChannelFromMessage:
		r = TL_inputChannelFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_contacts_resolvedPeer:
		r = TL_contacts_resolvedPeer{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_messageRange:
		r = TL_messageRange{
			m.Int(),
			m.Int(),
		}

	case CRC_updates_channelDifferenceEmpty:
		var flags int32
		r = TL_updates_channelDifferenceEmpty{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.FlaggedInt(flags, 1),
		}

	case CRC_updates_channelDifferenceTooLong:
		var flags int32
		r = TL_updates_channelDifferenceTooLong{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_updates_channelDifference:
		var flags int32
		r = TL_updates_channelDifference{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_channelMessagesFilterEmpty:
		r = TL_channelMessagesFilterEmpty{}

	case CRC_channelMessagesFilter:
		var flags int32
		r = TL_channelMessagesFilter{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Vector(),
		}

	case CRC_channelParticipant:
		r = TL_channelParticipant{
			m.Int(),
			m.Int(),
		}

	case CRC_channelParticipantSelf:
		r = TL_channelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_channelParticipantCreator:
		var flags int32
		r = TL_channelParticipantCreator{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedString(flags, 0),
		}

	case CRC_channelParticipantAdmin:
		var flags int32
		r = TL_channelParticipantAdmin{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedString(flags, 2),
		}

	case CRC_channelParticipantBanned:
		var flags int32
		r = TL_channelParticipantBanned{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CRC_channelParticipantsRecent:
		r = TL_channelParticipantsRecent{}

	case CRC_channelParticipantsAdmins:
		r = TL_channelParticipantsAdmins{}

	case CRC_channelParticipantsKicked:
		r = TL_channelParticipantsKicked{
			m.String(),
		}

	case CRC_channelParticipantsBots:
		r = TL_channelParticipantsBots{}

	case CRC_channelParticipantsBanned:
		r = TL_channelParticipantsBanned{
			m.String(),
		}

	case CRC_channelParticipantsSearch:
		r = TL_channelParticipantsSearch{
			m.String(),
		}

	case CRC_channelParticipantsContacts:
		r = TL_channelParticipantsContacts{
			m.String(),
		}

	case CRC_channels_channelParticipants:
		r = TL_channels_channelParticipants{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_channels_channelParticipantsNotModified:
		r = TL_channels_channelParticipantsNotModified{}

	case CRC_channels_channelParticipant:
		r = TL_channels_channelParticipant{
			m.Object(),
			m.Vector(),
		}

	case CRC_help_termsOfService:
		var flags int32
		r = TL_help_termsOfService{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
			m.Vector(),
			m.FlaggedInt(flags, 1),
		}

	case CRC_foundGif:
		r = TL_foundGif{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case CRC_foundGifCached:
		r = TL_foundGifCached{
			m.String(),
			m.Object(),
			m.Object(),
		}

	case CRC_messages_foundGifs:
		r = TL_messages_foundGifs{
			m.Int(),
			m.Vector(),
		}

	case CRC_messages_savedGifsNotModified:
		r = TL_messages_savedGifsNotModified{}

	case CRC_messages_savedGifs:
		r = TL_messages_savedGifs{
			m.Int(),
			m.Vector(),
		}

	case CRC_inputBotInlineMessageMediaAuto:
		var flags int32
		r = TL_inputBotInlineMessageMediaAuto{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputBotInlineMessageText:
		var flags int32
		r = TL_inputBotInlineMessageText{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputBotInlineMessageMediaGeo:
		var flags int32
		r = TL_inputBotInlineMessageMediaGeo{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputBotInlineMessageMediaVenue:
		var flags int32
		r = TL_inputBotInlineMessageMediaVenue{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputBotInlineMessageMediaContact:
		var flags int32
		r = TL_inputBotInlineMessageMediaContact{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputBotInlineMessageGame:
		var flags int32
		r = TL_inputBotInlineMessageGame{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 2),
		}

	case CRC_inputBotInlineResult:
		var flags int32
		r = TL_inputBotInlineResult{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 5),
			m.Object(),
		}

	case CRC_inputBotInlineResultPhoto:
		r = TL_inputBotInlineResultPhoto{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case CRC_inputBotInlineResultDocument:
		var flags int32
		r = TL_inputBotInlineResultDocument{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.Object(),
			m.Object(),
		}

	case CRC_inputBotInlineResultGame:
		r = TL_inputBotInlineResultGame{
			m.String(),
			m.String(),
			m.Object(),
		}

	case CRC_botInlineMessageMediaAuto:
		var flags int32
		r = TL_botInlineMessageMediaAuto{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case CRC_botInlineMessageText:
		var flags int32
		r = TL_botInlineMessageText{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case CRC_botInlineMessageMediaGeo:
		var flags int32
		r = TL_botInlineMessageMediaGeo{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_botInlineMessageMediaVenue:
		var flags int32
		r = TL_botInlineMessageMediaVenue{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_botInlineMessageMediaContact:
		var flags int32
		r = TL_botInlineMessageMediaContact{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case CRC_botInlineResult:
		var flags int32
		r = TL_botInlineResult{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 5),
			m.Object(),
		}

	case CRC_botInlineMediaResult:
		var flags int32
		r = TL_botInlineMediaResult{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.Object(),
		}

	case CRC_messages_botResults:
		var flags int32
		r = TL_messages_botResults{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case CRC_exportedMessageLink:
		r = TL_exportedMessageLink{
			m.String(),
			m.String(),
		}

	case CRC_messageFwdHeader:
		var flags int32
		r = TL_messageFwdHeader{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.FlaggedString(flags, 5),
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedString(flags, 6),
		}

	case CRC_auth_codeTypeSms:
		r = TL_auth_codeTypeSms{}

	case CRC_auth_codeTypeCall:
		r = TL_auth_codeTypeCall{}

	case CRC_auth_codeTypeFlashCall:
		r = TL_auth_codeTypeFlashCall{}

	case CRC_auth_sentCodeTypeApp:
		r = TL_auth_sentCodeTypeApp{
			m.Int(),
		}

	case CRC_auth_sentCodeTypeSms:
		r = TL_auth_sentCodeTypeSms{
			m.Int(),
		}

	case CRC_auth_sentCodeTypeCall:
		r = TL_auth_sentCodeTypeCall{
			m.Int(),
		}

	case CRC_auth_sentCodeTypeFlashCall:
		r = TL_auth_sentCodeTypeFlashCall{
			m.String(),
		}

	case CRC_messages_botCallbackAnswer:
		var flags int32
		r = TL_messages_botCallbackAnswer{
			readFlags(m, &flags),
			flags&2 != 0,  //flag #1
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case CRC_messages_messageEditData:
		var flags int32
		r = TL_messages_messageEditData{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case CRC_inputBotInlineMessageID:
		r = TL_inputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case CRC_inlineBotSwitchPM:
		r = TL_inlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case CRC_messages_peerDialogs:
		r = TL_messages_peerDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case CRC_topPeer:
		r = TL_topPeer{
			m.Object(),
			m.Double(),
		}

	case CRC_topPeerCategoryBotsPM:
		r = TL_topPeerCategoryBotsPM{}

	case CRC_topPeerCategoryBotsInline:
		r = TL_topPeerCategoryBotsInline{}

	case CRC_topPeerCategoryCorrespondents:
		r = TL_topPeerCategoryCorrespondents{}

	case CRC_topPeerCategoryGroups:
		r = TL_topPeerCategoryGroups{}

	case CRC_topPeerCategoryChannels:
		r = TL_topPeerCategoryChannels{}

	case CRC_topPeerCategoryPhoneCalls:
		r = TL_topPeerCategoryPhoneCalls{}

	case CRC_topPeerCategoryForwardUsers:
		r = TL_topPeerCategoryForwardUsers{}

	case CRC_topPeerCategoryForwardChats:
		r = TL_topPeerCategoryForwardChats{}

	case CRC_topPeerCategoryPeers:
		r = TL_topPeerCategoryPeers{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case CRC_contacts_topPeersNotModified:
		r = TL_contacts_topPeersNotModified{}

	case CRC_contacts_topPeers:
		r = TL_contacts_topPeers{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_contacts_topPeersDisabled:
		r = TL_contacts_topPeersDisabled{}

	case CRC_draftMessageEmpty:
		var flags int32
		r = TL_draftMessageEmpty{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
		}

	case CRC_draftMessage:
		var flags int32
		r = TL_draftMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedInt(flags, 0),
			m.String(),
			m.FlaggedVector(flags, 3),
			m.Int(),
		}

	case CRC_messages_featuredStickersNotModified:
		r = TL_messages_featuredStickersNotModified{
			m.Int(),
		}

	case CRC_messages_featuredStickers:
		r = TL_messages_featuredStickers{
			m.Int(),
			m.Int(),
			m.Vector(),
			m.VectorLong(),
		}

	case CRC_messages_recentStickersNotModified:
		r = TL_messages_recentStickersNotModified{}

	case CRC_messages_recentStickers:
		r = TL_messages_recentStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.VectorInt(),
		}

	case CRC_messages_archivedStickers:
		r = TL_messages_archivedStickers{
			m.Int(),
			m.Vector(),
		}

	case CRC_messages_stickerSetInstallResultSuccess:
		r = TL_messages_stickerSetInstallResultSuccess{}

	case CRC_messages_stickerSetInstallResultArchive:
		r = TL_messages_stickerSetInstallResultArchive{
			m.Vector(),
		}

	case CRC_stickerSetCovered:
		r = TL_stickerSetCovered{
			m.Object(),
			m.Object(),
		}

	case CRC_stickerSetMultiCovered:
		r = TL_stickerSetMultiCovered{
			m.Object(),
			m.Vector(),
		}

	case CRC_maskCoords:
		r = TL_maskCoords{
			m.Int(),
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case CRC_inputStickeredMediaPhoto:
		r = TL_inputStickeredMediaPhoto{
			m.Object(),
		}

	case CRC_inputStickeredMediaDocument:
		r = TL_inputStickeredMediaDocument{
			m.Object(),
		}

	case CRC_game:
		var flags int32
		r = TL_game{
			readFlags(m, &flags),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case CRC_inputGameID:
		r = TL_inputGameID{
			m.Long(),
			m.Long(),
		}

	case CRC_inputGameShortName:
		r = TL_inputGameShortName{
			m.Object(),
			m.String(),
		}

	case CRC_highScore:
		r = TL_highScore{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_highScores:
		r = TL_messages_highScores{
			m.Vector(),
			m.Vector(),
		}

	case CRC_textEmpty:
		r = TL_textEmpty{}

	case CRC_textPlain:
		r = TL_textPlain{
			m.String(),
		}

	case CRC_textBold:
		r = TL_textBold{
			m.Object(),
		}

	case CRC_textItalic:
		r = TL_textItalic{
			m.Object(),
		}

	case CRC_textUnderline:
		r = TL_textUnderline{
			m.Object(),
		}

	case CRC_textStrike:
		r = TL_textStrike{
			m.Object(),
		}

	case CRC_textFixed:
		r = TL_textFixed{
			m.Object(),
		}

	case CRC_textUrl:
		r = TL_textUrl{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case CRC_textEmail:
		r = TL_textEmail{
			m.Object(),
			m.String(),
		}

	case CRC_textConcat:
		r = TL_textConcat{
			m.Vector(),
		}

	case CRC_textSubscript:
		r = TL_textSubscript{
			m.Object(),
		}

	case CRC_textSuperscript:
		r = TL_textSuperscript{
			m.Object(),
		}

	case CRC_textMarked:
		r = TL_textMarked{
			m.Object(),
		}

	case CRC_textPhone:
		r = TL_textPhone{
			m.Object(),
			m.String(),
		}

	case CRC_textImage:
		r = TL_textImage{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case CRC_textAnchor:
		r = TL_textAnchor{
			m.Object(),
			m.String(),
		}

	case CRC_pageBlockUnsupported:
		r = TL_pageBlockUnsupported{}

	case CRC_pageBlockTitle:
		r = TL_pageBlockTitle{
			m.Object(),
		}

	case CRC_pageBlockSubtitle:
		r = TL_pageBlockSubtitle{
			m.Object(),
		}

	case CRC_pageBlockAuthorDate:
		r = TL_pageBlockAuthorDate{
			m.Object(),
			m.Int(),
		}

	case CRC_pageBlockHeader:
		r = TL_pageBlockHeader{
			m.Object(),
		}

	case CRC_pageBlockSubheader:
		r = TL_pageBlockSubheader{
			m.Object(),
		}

	case CRC_pageBlockParagraph:
		r = TL_pageBlockParagraph{
			m.Object(),
		}

	case CRC_pageBlockPreformatted:
		r = TL_pageBlockPreformatted{
			m.Object(),
			m.String(),
		}

	case CRC_pageBlockFooter:
		r = TL_pageBlockFooter{
			m.Object(),
		}

	case CRC_pageBlockDivider:
		r = TL_pageBlockDivider{}

	case CRC_pageBlockAnchor:
		r = TL_pageBlockAnchor{
			m.String(),
		}

	case CRC_pageBlockList:
		r = TL_pageBlockList{
			m.Vector(),
		}

	case CRC_pageBlockBlockquote:
		r = TL_pageBlockBlockquote{
			m.Object(),
			m.Object(),
		}

	case CRC_pageBlockPullquote:
		r = TL_pageBlockPullquote{
			m.Object(),
			m.Object(),
		}

	case CRC_pageBlockPhoto:
		var flags int32
		r = TL_pageBlockPhoto{
			readFlags(m, &flags),
			m.Long(),
			m.Object(),
			m.FlaggedString(flags, 0),
			m.FlaggedLong(flags, 0),
		}

	case CRC_pageBlockVideo:
		var flags int32
		r = TL_pageBlockVideo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Long(),
			m.Object(),
		}

	case CRC_pageBlockCover:
		r = TL_pageBlockCover{
			m.Object(),
		}

	case CRC_pageBlockEmbed:
		var flags int32
		r = TL_pageBlockEmbed{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&8 != 0, //flag #3
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedLong(flags, 4),
			m.FlaggedInt(flags, 5),
			m.FlaggedInt(flags, 5),
			m.Object(),
		}

	case CRC_pageBlockEmbedPost:
		r = TL_pageBlockEmbedPost{
			m.String(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Int(),
			m.Vector(),
			m.Object(),
		}

	case CRC_pageBlockCollage:
		r = TL_pageBlockCollage{
			m.Vector(),
			m.Object(),
		}

	case CRC_pageBlockSlideshow:
		r = TL_pageBlockSlideshow{
			m.Vector(),
			m.Object(),
		}

	case CRC_pageBlockChannel:
		r = TL_pageBlockChannel{
			m.Object(),
		}

	case CRC_pageBlockAudio:
		r = TL_pageBlockAudio{
			m.Long(),
			m.Object(),
		}

	case CRC_pageBlockKicker:
		r = TL_pageBlockKicker{
			m.Object(),
		}

	case CRC_pageBlockTable:
		var flags int32
		r = TL_pageBlockTable{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Vector(),
		}

	case CRC_pageBlockOrderedList:
		r = TL_pageBlockOrderedList{
			m.Vector(),
		}

	case CRC_pageBlockDetails:
		var flags int32
		r = TL_pageBlockDetails{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Vector(),
			m.Object(),
		}

	case CRC_pageBlockRelatedArticles:
		r = TL_pageBlockRelatedArticles{
			m.Object(),
			m.Vector(),
		}

	case CRC_pageBlockMap:
		r = TL_pageBlockMap{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CRC_phoneCallDiscardReasonMissed:
		r = TL_phoneCallDiscardReasonMissed{}

	case CRC_phoneCallDiscardReasonDisconnect:
		r = TL_phoneCallDiscardReasonDisconnect{}

	case CRC_phoneCallDiscardReasonHangup:
		r = TL_phoneCallDiscardReasonHangup{}

	case CRC_phoneCallDiscardReasonBusy:
		r = TL_phoneCallDiscardReasonBusy{}

	case CRC_dataJSON:
		r = TL_dataJSON{
			m.String(),
		}

	case CRC_labeledPrice:
		r = TL_labeledPrice{
			m.String(),
			m.Long(),
		}

	case CRC_invoice:
		var flags int32
		r = TL_invoice{
			readFlags(m, &flags),
			flags&1 != 0,   //flag #0
			flags&2 != 0,   //flag #1
			flags&4 != 0,   //flag #2
			flags&8 != 0,   //flag #3
			flags&16 != 0,  //flag #4
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.String(),
			m.Vector(),
		}

	case CRC_paymentCharge:
		r = TL_paymentCharge{
			m.String(),
			m.String(),
		}

	case CRC_postAddress:
		r = TL_postAddress{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_paymentRequestedInfo:
		var flags int32
		r = TL_paymentRequestedInfo{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case CRC_paymentSavedCredentialsCard:
		r = TL_paymentSavedCredentialsCard{
			m.String(),
			m.String(),
		}

	case CRC_webDocument:
		r = TL_webDocument{
			m.String(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case CRC_webDocumentNoProxy:
		r = TL_webDocumentNoProxy{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case CRC_inputWebDocument:
		r = TL_inputWebDocument{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case CRC_inputWebFileLocation:
		r = TL_inputWebFileLocation{
			m.String(),
			m.Long(),
		}

	case CRC_inputWebFileGeoPointLocation:
		r = TL_inputWebFileGeoPointLocation{
			m.Object(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_upload_webFile:
		r = TL_upload_webFile{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_payments_paymentForm:
		var flags int32
		r = TL_payments_paymentForm{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Vector(),
		}

	case CRC_payments_validatedRequestedInfo:
		var flags int32
		r = TL_payments_validatedRequestedInfo{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case CRC_payments_paymentResult:
		r = TL_payments_paymentResult{
			m.Object(),
		}

	case CRC_payments_paymentVerificationNeeded:
		r = TL_payments_paymentVerificationNeeded{
			m.String(),
		}

	case CRC_payments_paymentReceipt:
		var flags int32
		r = TL_payments_paymentReceipt{
			readFlags(m, &flags),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.String(),
			m.Long(),
			m.String(),
			m.Vector(),
		}

	case CRC_payments_savedInfo:
		var flags int32
		r = TL_payments_savedInfo{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedObject(flags, 0),
		}

	case CRC_inputPaymentCredentialsSaved:
		r = TL_inputPaymentCredentialsSaved{
			m.String(),
			m.StringBytes(),
		}

	case CRC_inputPaymentCredentials:
		var flags int32
		r = TL_inputPaymentCredentials{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case CRC_inputPaymentCredentialsApplePay:
		r = TL_inputPaymentCredentialsApplePay{
			m.Object(),
		}

	case CRC_inputPaymentCredentialsAndroidPay:
		r = TL_inputPaymentCredentialsAndroidPay{
			m.Object(),
			m.String(),
		}

	case CRC_account_tmpPassword:
		r = TL_account_tmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case CRC_shippingOption:
		r = TL_shippingOption{
			m.String(),
			m.String(),
			m.Vector(),
		}

	case CRC_inputStickerSetItem:
		var flags int32
		r = TL_inputStickerSetItem{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case CRC_inputPhoneCall:
		r = TL_inputPhoneCall{
			m.Long(),
			m.Long(),
		}

	case CRC_phoneCallEmpty:
		r = TL_phoneCallEmpty{
			m.Long(),
		}

	case CRC_phoneCallWaiting:
		var flags int32
		r = TL_phoneCallWaiting{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_phoneCallRequested:
		var flags int32
		r = TL_phoneCallRequested{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_phoneCallAccepted:
		var flags int32
		r = TL_phoneCallAccepted{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_phoneCall:
		var flags int32
		r = TL_phoneCall{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
			m.Vector(),
			m.Int(),
		}

	case CRC_phoneCallDiscarded:
		var flags int32
		r = TL_phoneCallDiscarded{
			readFlags(m, &flags),
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&32 != 0, //flag #5
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case CRC_phoneConnection:
		r = TL_phoneConnection{
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_phoneCallProtocol:
		var flags int32
		r = TL_phoneCallProtocol{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.VectorString(),
		}

	case CRC_phone_phoneCall:
		r = TL_phone_phoneCall{
			m.Object(),
			m.Vector(),
		}

	case CRC_upload_cdnFileReuploadNeeded:
		r = TL_upload_cdnFileReuploadNeeded{
			m.StringBytes(),
		}

	case CRC_upload_cdnFile:
		r = TL_upload_cdnFile{
			m.StringBytes(),
		}

	case CRC_cdnPublicKey:
		r = TL_cdnPublicKey{
			m.Int(),
			m.String(),
		}

	case CRC_cdnConfig:
		r = TL_cdnConfig{
			m.Vector(),
		}

	case CRC_langPackString:
		r = TL_langPackString{
			m.String(),
			m.String(),
		}

	case CRC_langPackStringPluralized:
		var flags int32
		r = TL_langPackStringPluralized{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.String(),
		}

	case CRC_langPackStringDeleted:
		r = TL_langPackStringDeleted{
			m.String(),
		}

	case CRC_langPackDifference:
		r = TL_langPackDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case CRC_langPackLanguage:
		var flags int32
		r = TL_langPackLanguage{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case CRC_channelAdminLogEventActionChangeTitle:
		r = TL_channelAdminLogEventActionChangeTitle{
			m.String(),
			m.String(),
		}

	case CRC_channelAdminLogEventActionChangeAbout:
		r = TL_channelAdminLogEventActionChangeAbout{
			m.String(),
			m.String(),
		}

	case CRC_channelAdminLogEventActionChangeUsername:
		r = TL_channelAdminLogEventActionChangeUsername{
			m.String(),
			m.String(),
		}

	case CRC_channelAdminLogEventActionChangePhoto:
		r = TL_channelAdminLogEventActionChangePhoto{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionToggleInvites:
		r = TL_channelAdminLogEventActionToggleInvites{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionToggleSignatures:
		r = TL_channelAdminLogEventActionToggleSignatures{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionUpdatePinned:
		r = TL_channelAdminLogEventActionUpdatePinned{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionEditMessage:
		r = TL_channelAdminLogEventActionEditMessage{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionDeleteMessage:
		r = TL_channelAdminLogEventActionDeleteMessage{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionParticipantJoin:
		r = TL_channelAdminLogEventActionParticipantJoin{}

	case CRC_channelAdminLogEventActionParticipantLeave:
		r = TL_channelAdminLogEventActionParticipantLeave{}

	case CRC_channelAdminLogEventActionParticipantInvite:
		r = TL_channelAdminLogEventActionParticipantInvite{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionParticipantToggleBan:
		r = TL_channelAdminLogEventActionParticipantToggleBan{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionParticipantToggleAdmin:
		r = TL_channelAdminLogEventActionParticipantToggleAdmin{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionChangeStickerSet:
		r = TL_channelAdminLogEventActionChangeStickerSet{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionTogglePreHistoryHidden:
		r = TL_channelAdminLogEventActionTogglePreHistoryHidden{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionDefaultBannedRights:
		r = TL_channelAdminLogEventActionDefaultBannedRights{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionStopPoll:
		r = TL_channelAdminLogEventActionStopPoll{
			m.Object(),
		}

	case CRC_channelAdminLogEventActionChangeLinkedChat:
		r = TL_channelAdminLogEventActionChangeLinkedChat{
			m.Int(),
			m.Int(),
		}

	case CRC_channelAdminLogEventActionChangeLocation:
		r = TL_channelAdminLogEventActionChangeLocation{
			m.Object(),
			m.Object(),
		}

	case CRC_channelAdminLogEventActionToggleSlowMode:
		r = TL_channelAdminLogEventActionToggleSlowMode{
			m.Int(),
			m.Int(),
		}

	case CRC_channelAdminLogEvent:
		r = TL_channelAdminLogEvent{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case CRC_channels_adminLogResults:
		r = TL_channels_adminLogResults{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_channelAdminLogEventsFilter:
		var flags int32
		r = TL_channelAdminLogEventsFilter{
			readFlags(m, &flags),
			flags&1 != 0,    //flag #0
			flags&2 != 0,    //flag #1
			flags&4 != 0,    //flag #2
			flags&8 != 0,    //flag #3
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&64 != 0,   //flag #6
			flags&128 != 0,  //flag #7
			flags&256 != 0,  //flag #8
			flags&512 != 0,  //flag #9
			flags&1024 != 0, //flag #10
			flags&2048 != 0, //flag #11
			flags&4096 != 0, //flag #12
			flags&8192 != 0, //flag #13
		}

	case CRC_popularContact:
		r = TL_popularContact{
			m.Long(),
			m.Int(),
		}

	case CRC_messages_favedStickersNotModified:
		r = TL_messages_favedStickersNotModified{}

	case CRC_messages_favedStickers:
		r = TL_messages_favedStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_recentMeUrlUnknown:
		r = TL_recentMeUrlUnknown{
			m.String(),
		}

	case CRC_recentMeUrlUser:
		r = TL_recentMeUrlUser{
			m.String(),
			m.Int(),
		}

	case CRC_recentMeUrlChat:
		r = TL_recentMeUrlChat{
			m.String(),
			m.Int(),
		}

	case CRC_recentMeUrlChatInvite:
		r = TL_recentMeUrlChatInvite{
			m.String(),
			m.Object(),
		}

	case CRC_recentMeUrlStickerSet:
		r = TL_recentMeUrlStickerSet{
			m.String(),
			m.Object(),
		}

	case CRC_help_recentMeUrls:
		r = TL_help_recentMeUrls{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_inputSingleMedia:
		var flags int32
		r = TL_inputSingleMedia{
			readFlags(m, &flags),
			m.Object(),
			m.Long(),
			m.String(),
			m.FlaggedVector(flags, 0),
		}

	case CRC_webAuthorization:
		r = TL_webAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case CRC_account_webAuthorizations:
		r = TL_account_webAuthorizations{
			m.Vector(),
			m.Vector(),
		}

	case CRC_inputMessageID:
		r = TL_inputMessageID{
			m.Int(),
		}

	case CRC_inputMessageReplyTo:
		r = TL_inputMessageReplyTo{
			m.Int(),
		}

	case CRC_inputMessagePinned:
		r = TL_inputMessagePinned{}

	case CRC_inputDialogPeer:
		r = TL_inputDialogPeer{
			m.Object(),
		}

	case CRC_inputDialogPeerFolder:
		r = TL_inputDialogPeerFolder{
			m.Int(),
		}

	case CRC_dialogPeer:
		r = TL_dialogPeer{
			m.Object(),
		}

	case CRC_dialogPeerFolder:
		r = TL_dialogPeerFolder{
			m.Int(),
		}

	case CRC_messages_foundStickerSetsNotModified:
		r = TL_messages_foundStickerSetsNotModified{}

	case CRC_messages_foundStickerSets:
		r = TL_messages_foundStickerSets{
			m.Int(),
			m.Vector(),
		}

	case CRC_fileHash:
		r = TL_fileHash{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_inputClientProxy:
		r = TL_inputClientProxy{
			m.String(),
			m.Int(),
		}

	case CRC_help_termsOfServiceUpdateEmpty:
		r = TL_help_termsOfServiceUpdateEmpty{
			m.Int(),
		}

	case CRC_help_termsOfServiceUpdate:
		r = TL_help_termsOfServiceUpdate{
			m.Int(),
			m.Object(),
		}

	case CRC_inputSecureFileUploaded:
		r = TL_inputSecureFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case CRC_inputSecureFile:
		r = TL_inputSecureFile{
			m.Long(),
			m.Long(),
		}

	case CRC_secureFileEmpty:
		r = TL_secureFileEmpty{}

	case CRC_secureFile:
		r = TL_secureFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case CRC_secureData:
		r = TL_secureData{
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case CRC_securePlainPhone:
		r = TL_securePlainPhone{
			m.String(),
		}

	case CRC_securePlainEmail:
		r = TL_securePlainEmail{
			m.String(),
		}

	case CRC_secureValueTypePersonalDetails:
		r = TL_secureValueTypePersonalDetails{}

	case CRC_secureValueTypePassport:
		r = TL_secureValueTypePassport{}

	case CRC_secureValueTypeDriverLicense:
		r = TL_secureValueTypeDriverLicense{}

	case CRC_secureValueTypeIdentityCard:
		r = TL_secureValueTypeIdentityCard{}

	case CRC_secureValueTypeInternalPassport:
		r = TL_secureValueTypeInternalPassport{}

	case CRC_secureValueTypeAddress:
		r = TL_secureValueTypeAddress{}

	case CRC_secureValueTypeUtilityBill:
		r = TL_secureValueTypeUtilityBill{}

	case CRC_secureValueTypeBankStatement:
		r = TL_secureValueTypeBankStatement{}

	case CRC_secureValueTypeRentalAgreement:
		r = TL_secureValueTypeRentalAgreement{}

	case CRC_secureValueTypePassportRegistration:
		r = TL_secureValueTypePassportRegistration{}

	case CRC_secureValueTypeTemporaryRegistration:
		r = TL_secureValueTypeTemporaryRegistration{}

	case CRC_secureValueTypePhone:
		r = TL_secureValueTypePhone{}

	case CRC_secureValueTypeEmail:
		r = TL_secureValueTypeEmail{}

	case CRC_secureValue:
		var flags int32
		r = TL_secureValue{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
			m.FlaggedVector(flags, 6),
			m.FlaggedVector(flags, 4),
			m.FlaggedObject(flags, 5),
			m.StringBytes(),
		}

	case CRC_inputSecureValue:
		var flags int32
		r = TL_inputSecureValue{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
			m.FlaggedVector(flags, 6),
			m.FlaggedVector(flags, 4),
			m.FlaggedObject(flags, 5),
		}

	case CRC_secureValueHash:
		r = TL_secureValueHash{
			m.Object(),
			m.StringBytes(),
		}

	case CRC_secureValueErrorData:
		r = TL_secureValueErrorData{
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.String(),
		}

	case CRC_secureValueErrorFrontSide:
		r = TL_secureValueErrorFrontSide{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_secureValueErrorReverseSide:
		r = TL_secureValueErrorReverseSide{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_secureValueErrorSelfie:
		r = TL_secureValueErrorSelfie{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_secureValueErrorFile:
		r = TL_secureValueErrorFile{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_secureValueErrorFiles:
		r = TL_secureValueErrorFiles{
			m.Object(),
			m.Vector(),
			m.String(),
		}

	case CRC_secureValueError:
		r = TL_secureValueError{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_secureValueErrorTranslationFile:
		r = TL_secureValueErrorTranslationFile{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case CRC_secureValueErrorTranslationFiles:
		r = TL_secureValueErrorTranslationFiles{
			m.Object(),
			m.Vector(),
			m.String(),
		}

	case CRC_secureCredentialsEncrypted:
		r = TL_secureCredentialsEncrypted{
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case CRC_account_authorizationForm:
		var flags int32
		r = TL_account_authorizationForm{
			readFlags(m, &flags),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.FlaggedString(flags, 0),
		}

	case CRC_account_sentEmailCode:
		r = TL_account_sentEmailCode{
			m.String(),
			m.Int(),
		}

	case CRC_help_deepLinkInfoEmpty:
		r = TL_help_deepLinkInfoEmpty{}

	case CRC_help_deepLinkInfo:
		var flags int32
		r = TL_help_deepLinkInfo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedVector(flags, 1),
		}

	case CRC_savedPhoneContact:
		r = TL_savedPhoneContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case CRC_account_takeout:
		r = TL_account_takeout{
			m.Long(),
		}

	case CRC_passwordKdfAlgoUnknown:
		r = TL_passwordKdfAlgoUnknown{}

	case CRC_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow:
		r = TL_passwordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{
			m.StringBytes(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_securePasswordKdfAlgoUnknown:
		r = TL_securePasswordKdfAlgoUnknown{}

	case CRC_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000:
		r = TL_securePasswordKdfAlgoPBKDF2HMACSHA512iter100000{
			m.StringBytes(),
		}

	case CRC_securePasswordKdfAlgoSHA512:
		r = TL_securePasswordKdfAlgoSHA512{
			m.StringBytes(),
		}

	case CRC_secureSecretSettings:
		r = TL_secureSecretSettings{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case CRC_inputCheckPasswordEmpty:
		r = TL_inputCheckPasswordEmpty{}

	case CRC_inputCheckPasswordSRP:
		r = TL_inputCheckPasswordSRP{
			m.Long(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case CRC_secureRequiredType:
		var flags int32
		r = TL_secureRequiredType{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Object(),
		}

	case CRC_secureRequiredTypeOneOf:
		r = TL_secureRequiredTypeOneOf{
			m.Vector(),
		}

	case CRC_help_passportConfigNotModified:
		r = TL_help_passportConfigNotModified{}

	case CRC_help_passportConfig:
		r = TL_help_passportConfig{
			m.Int(),
			m.Object(),
		}

	case CRC_inputAppEvent:
		r = TL_inputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.Object(),
		}

	case CRC_jsonObjectValue:
		r = TL_jsonObjectValue{
			m.String(),
			m.Object(),
		}

	case CRC_jsonNull:
		r = TL_jsonNull{}

	case CRC_jsonBool:
		r = TL_jsonBool{
			m.Object(),
		}

	case CRC_jsonNumber:
		r = TL_jsonNumber{
			m.Double(),
		}

	case CRC_jsonString:
		r = TL_jsonString{
			m.String(),
		}

	case CRC_jsonArray:
		r = TL_jsonArray{
			m.Vector(),
		}

	case CRC_jsonObject:
		r = TL_jsonObject{
			m.Vector(),
		}

	case CRC_pageTableCell:
		var flags int32
		r = TL_pageTableCell{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			flags&32 != 0, //flag #5
			flags&64 != 0, //flag #6
			m.FlaggedObject(flags, 7),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case CRC_pageTableRow:
		r = TL_pageTableRow{
			m.Vector(),
		}

	case CRC_pageCaption:
		r = TL_pageCaption{
			m.Object(),
			m.Object(),
		}

	case CRC_pageListItemText:
		r = TL_pageListItemText{
			m.Object(),
		}

	case CRC_pageListItemBlocks:
		r = TL_pageListItemBlocks{
			m.Vector(),
		}

	case CRC_pageListOrderedItemText:
		r = TL_pageListOrderedItemText{
			m.String(),
			m.Object(),
		}

	case CRC_pageListOrderedItemBlocks:
		r = TL_pageListOrderedItemBlocks{
			m.String(),
			m.Vector(),
		}

	case CRC_pageRelatedArticle:
		var flags int32
		r = TL_pageRelatedArticle{
			readFlags(m, &flags),
			m.String(),
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedLong(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedInt(flags, 4),
		}

	case CRC_page:
		var flags int32
		r = TL_page{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.String(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.FlaggedInt(flags, 3),
		}

	case CRC_help_supportName:
		r = TL_help_supportName{
			m.String(),
		}

	case CRC_help_userInfoEmpty:
		r = TL_help_userInfoEmpty{}

	case CRC_help_userInfo:
		r = TL_help_userInfo{
			m.String(),
			m.Vector(),
			m.String(),
			m.Int(),
		}

	case CRC_pollAnswer:
		r = TL_pollAnswer{
			m.String(),
			m.StringBytes(),
		}

	case CRC_poll:
		var flags int32
		r = TL_poll{
			m.Long(),
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.String(),
			m.Vector(),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
		}

	case CRC_pollAnswerVoters:
		var flags int32
		r = TL_pollAnswerVoters{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.StringBytes(),
			m.Int(),
		}

	case CRC_pollResults:
		var flags int32
		r = TL_pollResults{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedVector(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedVectorInt(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedVector(flags, 4),
		}

	case CRC_chatOnlines:
		r = TL_chatOnlines{
			m.Int(),
		}

	case CRC_statsURL:
		r = TL_statsURL{
			m.String(),
		}

	case CRC_chatAdminRights:
		var flags int32
		r = TL_chatAdminRights{
			readFlags(m, &flags),
			flags&1 != 0,   //flag #0
			flags&2 != 0,   //flag #1
			flags&4 != 0,   //flag #2
			flags&8 != 0,   //flag #3
			flags&16 != 0,  //flag #4
			flags&32 != 0,  //flag #5
			flags&128 != 0, //flag #7
			flags&512 != 0, //flag #9
		}

	case CRC_chatBannedRights:
		var flags int32
		r = TL_chatBannedRights{
			readFlags(m, &flags),
			flags&1 != 0,      //flag #0
			flags&2 != 0,      //flag #1
			flags&4 != 0,      //flag #2
			flags&8 != 0,      //flag #3
			flags&16 != 0,     //flag #4
			flags&32 != 0,     //flag #5
			flags&64 != 0,     //flag #6
			flags&128 != 0,    //flag #7
			flags&256 != 0,    //flag #8
			flags&1024 != 0,   //flag #10
			flags&32768 != 0,  //flag #15
			flags&131072 != 0, //flag #17
			m.Int(),
		}

	case CRC_inputWallPaper:
		r = TL_inputWallPaper{
			m.Long(),
			m.Long(),
		}

	case CRC_inputWallPaperSlug:
		r = TL_inputWallPaperSlug{
			m.String(),
		}

	case CRC_inputWallPaperNoFile:
		r = TL_inputWallPaperNoFile{}

	case CRC_account_wallPapersNotModified:
		r = TL_account_wallPapersNotModified{}

	case CRC_account_wallPapers:
		r = TL_account_wallPapers{
			m.Int(),
			m.Vector(),
		}

	case CRC_codeSettings:
		var flags int32
		r = TL_codeSettings{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&16 != 0, //flag #4
		}

	case CRC_wallPaperSettings:
		var flags int32
		r = TL_wallPaperSettings{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 3),
			m.FlaggedInt(flags, 4),
		}

	case CRC_autoDownloadSettings:
		var flags int32
		r = TL_autoDownloadSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_account_autoDownloadSettings:
		r = TL_account_autoDownloadSettings{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CRC_emojiKeyword:
		r = TL_emojiKeyword{
			m.String(),
			m.VectorString(),
		}

	case CRC_emojiKeywordDeleted:
		r = TL_emojiKeywordDeleted{
			m.String(),
			m.VectorString(),
		}

	case CRC_emojiKeywordsDifference:
		r = TL_emojiKeywordsDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case CRC_emojiURL:
		r = TL_emojiURL{
			m.String(),
		}

	case CRC_emojiLanguage:
		r = TL_emojiLanguage{
			m.String(),
		}

	case CRC_fileLocationToBeDeprecated:
		r = TL_fileLocationToBeDeprecated{
			m.Long(),
			m.Int(),
		}

	case CRC_folder:
		var flags int32
		r = TL_folder{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 3),
		}

	case CRC_inputFolderPeer:
		r = TL_inputFolderPeer{
			m.Object(),
			m.Int(),
		}

	case CRC_folderPeer:
		r = TL_folderPeer{
			m.Object(),
			m.Int(),
		}

	case CRC_messages_searchCounter:
		var flags int32
		r = TL_messages_searchCounter{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
		}

	case CRC_urlAuthResultRequest:
		var flags int32
		r = TL_urlAuthResultRequest{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
		}

	case CRC_urlAuthResultAccepted:
		r = TL_urlAuthResultAccepted{
			m.String(),
		}

	case CRC_urlAuthResultDefault:
		r = TL_urlAuthResultDefault{}

	case CRC_channelLocationEmpty:
		r = TL_channelLocationEmpty{}

	case CRC_channelLocation:
		r = TL_channelLocation{
			m.Object(),
			m.String(),
		}

	case CRC_peerLocated:
		r = TL_peerLocated{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_peerSelfLocated:
		r = TL_peerSelfLocated{
			m.Int(),
		}

	case CRC_restrictionReason:
		r = TL_restrictionReason{
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_inputTheme:
		r = TL_inputTheme{
			m.Long(),
			m.Long(),
		}

	case CRC_inputThemeSlug:
		r = TL_inputThemeSlug{
			m.String(),
		}

	case CRC_theme:
		var flags int32
		r = TL_theme{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
			m.Int(),
		}

	case CRC_account_themesNotModified:
		r = TL_account_themesNotModified{}

	case CRC_account_themes:
		r = TL_account_themes{
			m.Int(),
			m.Vector(),
		}

	case CRC_auth_loginToken:
		r = TL_auth_loginToken{
			m.Int(),
			m.StringBytes(),
		}

	case CRC_auth_loginTokenMigrateTo:
		r = TL_auth_loginTokenMigrateTo{
			m.Int(),
			m.StringBytes(),
		}

	case CRC_auth_loginTokenSuccess:
		r = TL_auth_loginTokenSuccess{
			m.Object(),
		}

	case CRC_account_contentSettings:
		var flags int32
		r = TL_account_contentSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
		}

	case CRC_messages_inactiveChats:
		r = TL_messages_inactiveChats{
			m.VectorInt(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_baseThemeClassic:
		r = TL_baseThemeClassic{}

	case CRC_baseThemeDay:
		r = TL_baseThemeDay{}

	case CRC_baseThemeNight:
		r = TL_baseThemeNight{}

	case CRC_baseThemeTinted:
		r = TL_baseThemeTinted{}

	case CRC_baseThemeArctic:
		r = TL_baseThemeArctic{}

	case CRC_inputThemeSettings:
		var flags int32
		r = TL_inputThemeSettings{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedObject(flags, 1),
		}

	case CRC_themeSettings:
		var flags int32
		r = TL_themeSettings{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case CRC_webPageAttributeTheme:
		var flags int32
		r = TL_webPageAttributeTheme{
			readFlags(m, &flags),
			m.FlaggedVector(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case CRC_messageUserVote:
		r = TL_messageUserVote{
			m.Int(),
			m.StringBytes(),
			m.Int(),
		}

	case CRC_messageUserVoteInputOption:
		r = TL_messageUserVoteInputOption{
			m.Int(),
			m.Int(),
		}

	case CRC_messageUserVoteMultiple:
		r = TL_messageUserVoteMultiple{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case CRC_messages_votesList:
		var flags int32
		r = TL_messages_votesList{
			readFlags(m, &flags),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.FlaggedString(flags, 0),
		}

	case CRC_bankCardOpenUrl:
		r = TL_bankCardOpenUrl{
			m.String(),
			m.String(),
		}

	case CRC_payments_bankCardData:
		r = TL_payments_bankCardData{
			m.String(),
			m.Vector(),
		}

	case CRC_dialogFilter:
		var flags int32
		r = TL_dialogFilter{
			readFlags(m, &flags),
			flags&1 != 0,    //flag #0
			flags&2 != 0,    //flag #1
			flags&4 != 0,    //flag #2
			flags&8 != 0,    //flag #3
			flags&16 != 0,   //flag #4
			flags&2048 != 0, //flag #11
			flags&4096 != 0, //flag #12
			flags&8192 != 0, //flag #13
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 25),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case CRC_dialogFilterSuggested:
		r = TL_dialogFilterSuggested{
			m.Object(),
			m.String(),
		}

	case CRC_statsDateRangeDays:
		r = TL_statsDateRangeDays{
			m.Int(),
			m.Int(),
		}

	case CRC_statsAbsValueAndPrev:
		r = TL_statsAbsValueAndPrev{
			m.Double(),
			m.Double(),
		}

	case CRC_statsPercentValue:
		r = TL_statsPercentValue{
			m.Double(),
			m.Double(),
		}

	case CRC_statsGraphAsync:
		r = TL_statsGraphAsync{
			m.String(),
		}

	case CRC_statsGraphError:
		r = TL_statsGraphError{
			m.String(),
		}

	case CRC_statsGraph:
		var flags int32
		r = TL_statsGraph{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedString(flags, 0),
		}

	case CRC_messageInteractionCounters:
		r = TL_messageInteractionCounters{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_stats_broadcastStats:
		r = TL_stats_broadcastStats{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case CRC_help_promoDataEmpty:
		r = TL_help_promoDataEmpty{
			m.Int(),
		}

	case CRC_help_promoData:
		var flags int32
		r = TL_help_promoData{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case CRC_videoSize:
		r = TL_videoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_invokeAfterMsg:
		r = TL_invokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case CRC_invokeAfterMsgs:
		r = TL_invokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case CRC_initConnection:
		var flags int32
		r = TL_initConnection{
			readFlags(m, &flags),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Object(),
		}

	case CRC_invokeWithLayer:
		r = TL_invokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case CRC_invokeWithoutUpdates:
		r = TL_invokeWithoutUpdates{
			m.Object(),
		}

	case CRC_invokeWithMessagesRange:
		r = TL_invokeWithMessagesRange{
			m.Object(),
			m.Object(),
		}

	case CRC_invokeWithTakeout:
		r = TL_invokeWithTakeout{
			m.Long(),
			m.Object(),
		}

	case CRC_auth_sendCode:
		r = TL_auth_sendCode{
			m.String(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case CRC_auth_signUp:
		r = TL_auth_signUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_auth_signIn:
		r = TL_auth_signIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_auth_logOut:
		r = TL_auth_logOut{}

	case CRC_auth_resetAuthorizations:
		r = TL_auth_resetAuthorizations{}

	case CRC_auth_exportAuthorization:
		r = TL_auth_exportAuthorization{
			m.Int(),
		}

	case CRC_auth_importAuthorization:
		r = TL_auth_importAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case CRC_auth_bindTempAuthKey:
		r = TL_auth_bindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_auth_importBotAuthorization:
		r = TL_auth_importBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case CRC_auth_checkPassword:
		r = TL_auth_checkPassword{
			m.Object(),
		}

	case CRC_auth_requestPasswordRecovery:
		r = TL_auth_requestPasswordRecovery{}

	case CRC_auth_recoverPassword:
		r = TL_auth_recoverPassword{
			m.String(),
		}

	case CRC_auth_resendCode:
		r = TL_auth_resendCode{
			m.String(),
			m.String(),
		}

	case CRC_auth_cancelCode:
		r = TL_auth_cancelCode{
			m.String(),
			m.String(),
		}

	case CRC_auth_dropTempAuthKeys:
		r = TL_auth_dropTempAuthKeys{
			m.VectorLong(),
		}

	case CRC_auth_exportLoginToken:
		r = TL_auth_exportLoginToken{
			m.Int(),
			m.String(),
			m.VectorInt(),
		}

	case CRC_auth_importLoginToken:
		r = TL_auth_importLoginToken{
			m.StringBytes(),
		}

	case CRC_auth_acceptLoginToken:
		r = TL_auth_acceptLoginToken{
			m.StringBytes(),
		}

	case CRC_account_registerDevice:
		var flags int32
		r = TL_account_registerDevice{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.String(),
			m.Object(),
			m.StringBytes(),
			m.VectorInt(),
		}

	case CRC_account_unregisterDevice:
		r = TL_account_unregisterDevice{
			m.Int(),
			m.String(),
			m.VectorInt(),
		}

	case CRC_account_updateNotifySettings:
		r = TL_account_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case CRC_account_getNotifySettings:
		r = TL_account_getNotifySettings{
			m.Object(),
		}

	case CRC_account_resetNotifySettings:
		r = TL_account_resetNotifySettings{}

	case CRC_account_updateProfile:
		var flags int32
		r = TL_account_updateProfile{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case CRC_account_updateStatus:
		r = TL_account_updateStatus{
			m.Object(),
		}

	case CRC_account_getWallPapers:
		r = TL_account_getWallPapers{
			m.Int(),
		}

	case CRC_account_reportPeer:
		r = TL_account_reportPeer{
			m.Object(),
			m.Object(),
		}

	case CRC_account_checkUsername:
		r = TL_account_checkUsername{
			m.String(),
		}

	case CRC_account_updateUsername:
		r = TL_account_updateUsername{
			m.String(),
		}

	case CRC_account_getPrivacy:
		r = TL_account_getPrivacy{
			m.Object(),
		}

	case CRC_account_setPrivacy:
		r = TL_account_setPrivacy{
			m.Object(),
			m.Vector(),
		}

	case CRC_account_deleteAccount:
		r = TL_account_deleteAccount{
			m.String(),
		}

	case CRC_account_getAccountTTL:
		r = TL_account_getAccountTTL{}

	case CRC_account_setAccountTTL:
		r = TL_account_setAccountTTL{
			m.Object(),
		}

	case CRC_account_sendChangePhoneCode:
		r = TL_account_sendChangePhoneCode{
			m.String(),
			m.Object(),
		}

	case CRC_account_changePhone:
		r = TL_account_changePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_account_updateDeviceLocked:
		r = TL_account_updateDeviceLocked{
			m.Int(),
		}

	case CRC_account_getAuthorizations:
		r = TL_account_getAuthorizations{}

	case CRC_account_resetAuthorization:
		r = TL_account_resetAuthorization{
			m.Long(),
		}

	case CRC_account_getPassword:
		r = TL_account_getPassword{}

	case CRC_account_getPasswordSettings:
		r = TL_account_getPasswordSettings{
			m.Object(),
		}

	case CRC_account_updatePasswordSettings:
		r = TL_account_updatePasswordSettings{
			m.Object(),
			m.Object(),
		}

	case CRC_account_sendConfirmPhoneCode:
		r = TL_account_sendConfirmPhoneCode{
			m.String(),
			m.Object(),
		}

	case CRC_account_confirmPhone:
		r = TL_account_confirmPhone{
			m.String(),
			m.String(),
		}

	case CRC_account_getTmpPassword:
		r = TL_account_getTmpPassword{
			m.Object(),
			m.Int(),
		}

	case CRC_account_getWebAuthorizations:
		r = TL_account_getWebAuthorizations{}

	case CRC_account_resetWebAuthorization:
		r = TL_account_resetWebAuthorization{
			m.Long(),
		}

	case CRC_account_resetWebAuthorizations:
		r = TL_account_resetWebAuthorizations{}

	case CRC_account_getAllSecureValues:
		r = TL_account_getAllSecureValues{}

	case CRC_account_getSecureValue:
		r = TL_account_getSecureValue{
			m.Vector(),
		}

	case CRC_account_saveSecureValue:
		r = TL_account_saveSecureValue{
			m.Object(),
			m.Long(),
		}

	case CRC_account_deleteSecureValue:
		r = TL_account_deleteSecureValue{
			m.Vector(),
		}

	case CRC_account_getAuthorizationForm:
		r = TL_account_getAuthorizationForm{
			m.Int(),
			m.String(),
			m.String(),
		}

	case CRC_account_acceptAuthorization:
		r = TL_account_acceptAuthorization{
			m.Int(),
			m.String(),
			m.String(),
			m.Vector(),
			m.Object(),
		}

	case CRC_account_sendVerifyPhoneCode:
		r = TL_account_sendVerifyPhoneCode{
			m.String(),
			m.Object(),
		}

	case CRC_account_verifyPhone:
		r = TL_account_verifyPhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_account_sendVerifyEmailCode:
		r = TL_account_sendVerifyEmailCode{
			m.String(),
		}

	case CRC_account_verifyEmail:
		r = TL_account_verifyEmail{
			m.String(),
			m.String(),
		}

	case CRC_account_initTakeoutSession:
		var flags int32
		r = TL_account_initTakeoutSession{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			flags&32 != 0, //flag #5
			m.FlaggedInt(flags, 5),
		}

	case CRC_account_finishTakeoutSession:
		var flags int32
		r = TL_account_finishTakeoutSession{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case CRC_account_confirmPasswordEmail:
		r = TL_account_confirmPasswordEmail{
			m.String(),
		}

	case CRC_account_resendPasswordEmail:
		r = TL_account_resendPasswordEmail{}

	case CRC_account_cancelPasswordEmail:
		r = TL_account_cancelPasswordEmail{}

	case CRC_account_getContactSignUpNotification:
		r = TL_account_getContactSignUpNotification{}

	case CRC_account_setContactSignUpNotification:
		r = TL_account_setContactSignUpNotification{
			m.Object(),
		}

	case CRC_account_getNotifyExceptions:
		var flags int32
		r = TL_account_getNotifyExceptions{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedObject(flags, 0),
		}

	case CRC_account_getWallPaper:
		r = TL_account_getWallPaper{
			m.Object(),
		}

	case CRC_account_uploadWallPaper:
		r = TL_account_uploadWallPaper{
			m.Object(),
			m.String(),
			m.Object(),
		}

	case CRC_account_saveWallPaper:
		r = TL_account_saveWallPaper{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CRC_account_installWallPaper:
		r = TL_account_installWallPaper{
			m.Object(),
			m.Object(),
		}

	case CRC_account_resetWallPapers:
		r = TL_account_resetWallPapers{}

	case CRC_account_getAutoDownloadSettings:
		r = TL_account_getAutoDownloadSettings{}

	case CRC_account_saveAutoDownloadSettings:
		var flags int32
		r = TL_account_saveAutoDownloadSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
		}

	case CRC_account_uploadTheme:
		var flags int32
		r = TL_account_uploadTheme{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.String(),
		}

	case CRC_account_createTheme:
		var flags int32
		r = TL_account_createTheme{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case CRC_account_updateTheme:
		var flags int32
		r = TL_account_updateTheme{
			readFlags(m, &flags),
			m.String(),
			m.Object(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case CRC_account_saveTheme:
		r = TL_account_saveTheme{
			m.Object(),
			m.Object(),
		}

	case CRC_account_installTheme:
		var flags int32
		r = TL_account_installTheme{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 1),
		}

	case CRC_account_getTheme:
		r = TL_account_getTheme{
			m.String(),
			m.Object(),
			m.Long(),
		}

	case CRC_account_getThemes:
		r = TL_account_getThemes{
			m.String(),
			m.Int(),
		}

	case CRC_account_setContentSettings:
		var flags int32
		r = TL_account_setContentSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case CRC_account_getContentSettings:
		r = TL_account_getContentSettings{}

	case CRC_account_getMultiWallPapers:
		r = TL_account_getMultiWallPapers{
			m.Vector(),
		}

	case CRC_users_getUsers:
		r = TL_users_getUsers{
			m.Vector(),
		}

	case CRC_users_getFullUser:
		r = TL_users_getFullUser{
			m.Object(),
		}

	case CRC_users_setSecureValueErrors:
		r = TL_users_setSecureValueErrors{
			m.Object(),
			m.Vector(),
		}

	case CRC_contacts_getContactIDs:
		r = TL_contacts_getContactIDs{
			m.Int(),
		}

	case CRC_contacts_getStatuses:
		r = TL_contacts_getStatuses{}

	case CRC_contacts_getContacts:
		r = TL_contacts_getContacts{
			m.Int(),
		}

	case CRC_contacts_importContacts:
		r = TL_contacts_importContacts{
			m.Vector(),
		}

	case CRC_contacts_deleteContacts:
		r = TL_contacts_deleteContacts{
			m.Vector(),
		}

	case CRC_contacts_deleteByPhones:
		r = TL_contacts_deleteByPhones{
			m.VectorString(),
		}

	case CRC_contacts_block:
		r = TL_contacts_block{
			m.Object(),
		}

	case CRC_contacts_unblock:
		r = TL_contacts_unblock{
			m.Object(),
		}

	case CRC_contacts_getBlocked:
		r = TL_contacts_getBlocked{
			m.Int(),
			m.Int(),
		}

	case CRC_contacts_search:
		r = TL_contacts_search{
			m.String(),
			m.Int(),
		}

	case CRC_contacts_resolveUsername:
		r = TL_contacts_resolveUsername{
			m.String(),
		}

	case CRC_contacts_getTopPeers:
		var flags int32
		r = TL_contacts_getTopPeers{
			readFlags(m, &flags),
			flags&1 != 0,     //flag #0
			flags&2 != 0,     //flag #1
			flags&4 != 0,     //flag #2
			flags&8 != 0,     //flag #3
			flags&16 != 0,    //flag #4
			flags&32 != 0,    //flag #5
			flags&1024 != 0,  //flag #10
			flags&32768 != 0, //flag #15
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_contacts_resetTopPeerRating:
		r = TL_contacts_resetTopPeerRating{
			m.Object(),
			m.Object(),
		}

	case CRC_contacts_resetSaved:
		r = TL_contacts_resetSaved{}

	case CRC_contacts_getSaved:
		r = TL_contacts_getSaved{}

	case CRC_contacts_toggleTopPeers:
		r = TL_contacts_toggleTopPeers{
			m.Object(),
		}

	case CRC_contacts_addContact:
		var flags int32
		r = TL_contacts_addContact{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
		}

	case CRC_contacts_acceptContact:
		r = TL_contacts_acceptContact{
			m.Object(),
		}

	case CRC_contacts_getLocated:
		var flags int32
		r = TL_contacts_getLocated{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case CRC_messages_getMessages:
		r = TL_messages_getMessages{
			m.Vector(),
		}

	case CRC_messages_getDialogs:
		var flags int32
		r = TL_messages_getDialogs{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_getHistory:
		r = TL_messages_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_search:
		var flags int32
		r = TL_messages_search{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_readHistory:
		r = TL_messages_readHistory{
			m.Object(),
			m.Int(),
		}

	case CRC_messages_deleteHistory:
		var flags int32
		r = TL_messages_deleteHistory{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
		}

	case CRC_messages_deleteMessages:
		var flags int32
		r = TL_messages_deleteMessages{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.VectorInt(),
		}

	case CRC_messages_receivedMessages:
		r = TL_messages_receivedMessages{
			m.Int(),
		}

	case CRC_messages_setTyping:
		r = TL_messages_setTyping{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_sendMessage:
		var flags int32
		r = TL_messages_sendMessage{
			readFlags(m, &flags),
			flags&2 != 0,   //flag #1
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Long(),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 10),
		}

	case CRC_messages_sendMedia:
		var flags int32
		r = TL_messages_sendMedia{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.String(),
			m.Long(),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 10),
		}

	case CRC_messages_forwardMessages:
		var flags int32
		r = TL_messages_forwardMessages{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&256 != 0, //flag #8
			flags&512 != 0, //flag #9
			m.Object(),
			m.VectorInt(),
			m.VectorLong(),
			m.Object(),
			m.FlaggedInt(flags, 10),
		}

	case CRC_messages_reportSpam:
		r = TL_messages_reportSpam{
			m.Object(),
		}

	case CRC_messages_getPeerSettings:
		r = TL_messages_getPeerSettings{
			m.Object(),
		}

	case CRC_messages_report:
		r = TL_messages_report{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case CRC_messages_getChats:
		r = TL_messages_getChats{
			m.VectorInt(),
		}

	case CRC_messages_getFullChat:
		r = TL_messages_getFullChat{
			m.Int(),
		}

	case CRC_messages_editChatTitle:
		r = TL_messages_editChatTitle{
			m.Int(),
			m.String(),
		}

	case CRC_messages_editChatPhoto:
		r = TL_messages_editChatPhoto{
			m.Int(),
			m.Object(),
		}

	case CRC_messages_addChatUser:
		r = TL_messages_addChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case CRC_messages_deleteChatUser:
		r = TL_messages_deleteChatUser{
			m.Int(),
			m.Object(),
		}

	case CRC_messages_createChat:
		r = TL_messages_createChat{
			m.Vector(),
			m.String(),
		}

	case CRC_messages_getDhConfig:
		r = TL_messages_getDhConfig{
			m.Int(),
			m.Int(),
		}

	case CRC_messages_requestEncryption:
		r = TL_messages_requestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_messages_acceptEncryption:
		r = TL_messages_acceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case CRC_messages_discardEncryption:
		r = TL_messages_discardEncryption{
			m.Int(),
		}

	case CRC_messages_setEncryptedTyping:
		r = TL_messages_setEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_readEncryptedHistory:
		r = TL_messages_readEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case CRC_messages_sendEncrypted:
		r = TL_messages_sendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case CRC_messages_sendEncryptedFile:
		r = TL_messages_sendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_messages_sendEncryptedService:
		r = TL_messages_sendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case CRC_messages_receivedQueue:
		r = TL_messages_receivedQueue{
			m.Int(),
		}

	case CRC_messages_reportEncryptedSpam:
		r = TL_messages_reportEncryptedSpam{
			m.Object(),
		}

	case CRC_messages_readMessageContents:
		r = TL_messages_readMessageContents{
			m.VectorInt(),
		}

	case CRC_messages_getStickers:
		r = TL_messages_getStickers{
			m.String(),
			m.Int(),
		}

	case CRC_messages_getAllStickers:
		r = TL_messages_getAllStickers{
			m.Int(),
		}

	case CRC_messages_getWebPagePreview:
		var flags int32
		r = TL_messages_getWebPagePreview{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedVector(flags, 3),
		}

	case CRC_messages_exportChatInvite:
		r = TL_messages_exportChatInvite{
			m.Object(),
		}

	case CRC_messages_checkChatInvite:
		r = TL_messages_checkChatInvite{
			m.String(),
		}

	case CRC_messages_importChatInvite:
		r = TL_messages_importChatInvite{
			m.String(),
		}

	case CRC_messages_getStickerSet:
		r = TL_messages_getStickerSet{
			m.Object(),
		}

	case CRC_messages_installStickerSet:
		r = TL_messages_installStickerSet{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_uninstallStickerSet:
		r = TL_messages_uninstallStickerSet{
			m.Object(),
		}

	case CRC_messages_startBot:
		r = TL_messages_startBot{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
		}

	case CRC_messages_getMessagesViews:
		r = TL_messages_getMessagesViews{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case CRC_messages_editChatAdmin:
		r = TL_messages_editChatAdmin{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case CRC_messages_migrateChat:
		r = TL_messages_migrateChat{
			m.Int(),
		}

	case CRC_messages_searchGlobal:
		var flags int32
		r = TL_messages_searchGlobal{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_reorderStickerSets:
		var flags int32
		r = TL_messages_reorderStickerSets{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.VectorLong(),
		}

	case CRC_messages_getDocumentByHash:
		r = TL_messages_getDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case CRC_messages_searchGifs:
		r = TL_messages_searchGifs{
			m.String(),
			m.Int(),
		}

	case CRC_messages_getSavedGifs:
		r = TL_messages_getSavedGifs{
			m.Int(),
		}

	case CRC_messages_saveGif:
		r = TL_messages_saveGif{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_getInlineBotResults:
		var flags int32
		r = TL_messages_getInlineBotResults{
			readFlags(m, &flags),
			m.Object(),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.String(),
		}

	case CRC_messages_setInlineBotResults:
		var flags int32
		r = TL_messages_setInlineBotResults{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Long(),
			m.Vector(),
			m.Int(),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case CRC_messages_sendInlineBotResult:
		var flags int32
		r = TL_messages_sendInlineBotResult{
			readFlags(m, &flags),
			flags&32 != 0,   //flag #5
			flags&64 != 0,   //flag #6
			flags&128 != 0,  //flag #7
			flags&2048 != 0, //flag #11
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 10),
		}

	case CRC_messages_getMessageEditData:
		r = TL_messages_getMessageEditData{
			m.Object(),
			m.Int(),
		}

	case CRC_messages_editMessage:
		var flags int32
		r = TL_messages_editMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 15),
		}

	case CRC_messages_editInlineBotMessage:
		var flags int32
		r = TL_messages_editInlineBotMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case CRC_messages_getBotCallbackAnswer:
		var flags int32
		r = TL_messages_getBotCallbackAnswer{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
		}

	case CRC_messages_setBotCallbackAnswer:
		var flags int32
		r = TL_messages_setBotCallbackAnswer{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case CRC_messages_getPeerDialogs:
		r = TL_messages_getPeerDialogs{
			m.Vector(),
		}

	case CRC_messages_saveDraft:
		var flags int32
		r = TL_messages_saveDraft{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.String(),
			m.FlaggedVector(flags, 3),
		}

	case CRC_messages_getAllDrafts:
		r = TL_messages_getAllDrafts{}

	case CRC_messages_getFeaturedStickers:
		r = TL_messages_getFeaturedStickers{
			m.Int(),
		}

	case CRC_messages_readFeaturedStickers:
		r = TL_messages_readFeaturedStickers{
			m.VectorLong(),
		}

	case CRC_messages_getRecentStickers:
		var flags int32
		r = TL_messages_getRecentStickers{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
		}

	case CRC_messages_saveRecentSticker:
		var flags int32
		r = TL_messages_saveRecentSticker{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Object(),
		}

	case CRC_messages_clearRecentStickers:
		var flags int32
		r = TL_messages_clearRecentStickers{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case CRC_messages_getArchivedStickers:
		var flags int32
		r = TL_messages_getArchivedStickers{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.Int(),
		}

	case CRC_messages_getMaskStickers:
		r = TL_messages_getMaskStickers{
			m.Int(),
		}

	case CRC_messages_getAttachedStickers:
		r = TL_messages_getAttachedStickers{
			m.Object(),
		}

	case CRC_messages_setGameScore:
		var flags int32
		r = TL_messages_setGameScore{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case CRC_messages_setInlineGameScore:
		var flags int32
		r = TL_messages_setInlineGameScore{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case CRC_messages_getGameHighScores:
		r = TL_messages_getGameHighScores{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case CRC_messages_getInlineGameHighScores:
		r = TL_messages_getInlineGameHighScores{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_getCommonChats:
		r = TL_messages_getCommonChats{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_getAllChats:
		r = TL_messages_getAllChats{
			m.VectorInt(),
		}

	case CRC_messages_getWebPage:
		r = TL_messages_getWebPage{
			m.String(),
			m.Int(),
		}

	case CRC_messages_toggleDialogPin:
		var flags int32
		r = TL_messages_toggleDialogPin{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case CRC_messages_reorderPinnedDialogs:
		var flags int32
		r = TL_messages_reorderPinnedDialogs{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Vector(),
		}

	case CRC_messages_getPinnedDialogs:
		r = TL_messages_getPinnedDialogs{
			m.Int(),
		}

	case CRC_messages_setBotShippingResults:
		var flags int32
		r = TL_messages_setBotShippingResults{
			readFlags(m, &flags),
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case CRC_messages_setBotPrecheckoutResults:
		var flags int32
		r = TL_messages_setBotPrecheckoutResults{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Long(),
			m.FlaggedString(flags, 0),
		}

	case CRC_messages_uploadMedia:
		r = TL_messages_uploadMedia{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_sendScreenshotNotification:
		r = TL_messages_sendScreenshotNotification{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case CRC_messages_getFavedStickers:
		r = TL_messages_getFavedStickers{
			m.Int(),
		}

	case CRC_messages_faveSticker:
		r = TL_messages_faveSticker{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_getUnreadMentions:
		r = TL_messages_getUnreadMentions{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_readMentions:
		r = TL_messages_readMentions{
			m.Object(),
		}

	case CRC_messages_getRecentLocations:
		r = TL_messages_getRecentLocations{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_sendMultiMedia:
		var flags int32
		r = TL_messages_sendMultiMedia{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Vector(),
			m.FlaggedInt(flags, 10),
		}

	case CRC_messages_uploadEncryptedFile:
		r = TL_messages_uploadEncryptedFile{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_searchStickerSets:
		var flags int32
		r = TL_messages_searchStickerSets{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.Int(),
		}

	case CRC_messages_getSplitRanges:
		r = TL_messages_getSplitRanges{}

	case CRC_messages_markDialogUnread:
		var flags int32
		r = TL_messages_markDialogUnread{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case CRC_messages_getDialogUnreadMarks:
		r = TL_messages_getDialogUnreadMarks{}

	case CRC_messages_clearAllDrafts:
		r = TL_messages_clearAllDrafts{}

	case CRC_messages_updatePinnedMessage:
		var flags int32
		r = TL_messages_updatePinnedMessage{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
		}

	case CRC_messages_sendVote:
		r = TL_messages_sendVote{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case CRC_messages_getPollResults:
		r = TL_messages_getPollResults{
			m.Object(),
			m.Int(),
		}

	case CRC_messages_getOnlines:
		r = TL_messages_getOnlines{
			m.Object(),
		}

	case CRC_messages_getStatsURL:
		var flags int32
		r = TL_messages_getStatsURL{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
		}

	case CRC_messages_editChatAbout:
		r = TL_messages_editChatAbout{
			m.Object(),
			m.String(),
		}

	case CRC_messages_editChatDefaultBannedRights:
		r = TL_messages_editChatDefaultBannedRights{
			m.Object(),
			m.Object(),
		}

	case CRC_messages_getEmojiKeywords:
		r = TL_messages_getEmojiKeywords{
			m.String(),
		}

	case CRC_messages_getEmojiKeywordsDifference:
		r = TL_messages_getEmojiKeywordsDifference{
			m.String(),
			m.Int(),
		}

	case CRC_messages_getEmojiKeywordsLanguages:
		r = TL_messages_getEmojiKeywordsLanguages{
			m.VectorString(),
		}

	case CRC_messages_getEmojiURL:
		r = TL_messages_getEmojiURL{
			m.String(),
		}

	case CRC_messages_getSearchCounters:
		r = TL_messages_getSearchCounters{
			m.Object(),
			m.Vector(),
		}

	case CRC_messages_requestUrlAuth:
		r = TL_messages_requestUrlAuth{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_acceptUrlAuth:
		var flags int32
		r = TL_messages_acceptUrlAuth{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_messages_hidePeerSettingsBar:
		r = TL_messages_hidePeerSettingsBar{
			m.Object(),
		}

	case CRC_messages_getScheduledHistory:
		r = TL_messages_getScheduledHistory{
			m.Object(),
			m.Int(),
		}

	case CRC_messages_getScheduledMessages:
		r = TL_messages_getScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case CRC_messages_sendScheduledMessages:
		r = TL_messages_sendScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case CRC_messages_deleteScheduledMessages:
		r = TL_messages_deleteScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case CRC_messages_getPollVotes:
		var flags int32
		r = TL_messages_getPollVotes{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
			m.Int(),
		}

	case CRC_messages_toggleStickerSets:
		var flags int32
		r = TL_messages_toggleStickerSets{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Vector(),
		}

	case CRC_messages_getDialogFilters:
		r = TL_messages_getDialogFilters{}

	case CRC_messages_getSuggestedDialogFilters:
		r = TL_messages_getSuggestedDialogFilters{}

	case CRC_messages_updateDialogFilter:
		var flags int32
		r = TL_messages_updateDialogFilter{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case CRC_messages_updateDialogFiltersOrder:
		r = TL_messages_updateDialogFiltersOrder{
			m.VectorInt(),
		}

	case CRC_messages_getOldFeaturedStickers:
		r = TL_messages_getOldFeaturedStickers{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_updates_getState:
		r = TL_updates_getState{}

	case CRC_updates_getDifference:
		var flags int32
		r = TL_updates_getDifference{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
		}

	case CRC_updates_getChannelDifference:
		var flags int32
		r = TL_updates_getChannelDifference{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_photos_updateProfilePhoto:
		r = TL_photos_updateProfilePhoto{
			m.Object(),
		}

	case CRC_photos_uploadProfilePhoto:
		r = TL_photos_uploadProfilePhoto{
			m.Object(),
		}

	case CRC_photos_deletePhotos:
		r = TL_photos_deletePhotos{
			m.Vector(),
		}

	case CRC_photos_getUserPhotos:
		r = TL_photos_getUserPhotos{
			m.Object(),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case CRC_upload_saveFilePart:
		r = TL_upload_saveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_upload_getFile:
		var flags int32
		r = TL_upload_getFile{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_upload_saveBigFilePart:
		r = TL_upload_saveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case CRC_upload_getWebFile:
		r = TL_upload_getWebFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case CRC_upload_getCdnFile:
		r = TL_upload_getCdnFile{
			m.StringBytes(),
			m.Int(),
			m.Int(),
		}

	case CRC_upload_reuploadCdnFile:
		r = TL_upload_reuploadCdnFile{
			m.StringBytes(),
			m.StringBytes(),
		}

	case CRC_upload_getCdnFileHashes:
		r = TL_upload_getCdnFileHashes{
			m.StringBytes(),
			m.Int(),
		}

	case CRC_upload_getFileHashes:
		r = TL_upload_getFileHashes{
			m.Object(),
			m.Int(),
		}

	case CRC_help_getConfig:
		r = TL_help_getConfig{}

	case CRC_help_getNearestDc:
		r = TL_help_getNearestDc{}

	case CRC_help_getAppUpdate:
		r = TL_help_getAppUpdate{
			m.String(),
		}

	case CRC_help_getInviteText:
		r = TL_help_getInviteText{}

	case CRC_help_getSupport:
		r = TL_help_getSupport{}

	case CRC_help_getAppChangelog:
		r = TL_help_getAppChangelog{
			m.String(),
		}

	case CRC_help_setBotUpdatesStatus:
		r = TL_help_setBotUpdatesStatus{
			m.Int(),
			m.String(),
		}

	case CRC_help_getCdnConfig:
		r = TL_help_getCdnConfig{}

	case CRC_help_getRecentMeUrls:
		r = TL_help_getRecentMeUrls{
			m.String(),
		}

	case CRC_help_getTermsOfServiceUpdate:
		r = TL_help_getTermsOfServiceUpdate{}

	case CRC_help_acceptTermsOfService:
		r = TL_help_acceptTermsOfService{
			m.Object(),
		}

	case CRC_help_getDeepLinkInfo:
		r = TL_help_getDeepLinkInfo{
			m.String(),
		}

	case CRC_help_getAppConfig:
		r = TL_help_getAppConfig{}

	case CRC_help_saveAppLog:
		r = TL_help_saveAppLog{
			m.Vector(),
		}

	case CRC_help_getPassportConfig:
		r = TL_help_getPassportConfig{
			m.Int(),
		}

	case CRC_help_getSupportName:
		r = TL_help_getSupportName{}

	case CRC_help_getUserInfo:
		r = TL_help_getUserInfo{
			m.Object(),
		}

	case CRC_help_editUserInfo:
		r = TL_help_editUserInfo{
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case CRC_help_getPromoData:
		r = TL_help_getPromoData{}

	case CRC_help_hidePromoData:
		r = TL_help_hidePromoData{
			m.Object(),
		}

	case CRC_channels_readHistory:
		r = TL_channels_readHistory{
			m.Object(),
			m.Int(),
		}

	case CRC_channels_deleteMessages:
		r = TL_channels_deleteMessages{
			m.Object(),
			m.VectorInt(),
		}

	case CRC_channels_deleteUserHistory:
		r = TL_channels_deleteUserHistory{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_reportSpam:
		r = TL_channels_reportSpam{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case CRC_channels_getMessages:
		r = TL_channels_getMessages{
			m.Object(),
			m.Vector(),
		}

	case CRC_channels_getParticipants:
		r = TL_channels_getParticipants{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case CRC_channels_getParticipant:
		r = TL_channels_getParticipant{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_getChannels:
		r = TL_channels_getChannels{
			m.Vector(),
		}

	case CRC_channels_getFullChannel:
		r = TL_channels_getFullChannel{
			m.Object(),
		}

	case CRC_channels_createChannel:
		var flags int32
		r = TL_channels_createChannel{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.FlaggedString(flags, 2),
		}

	case CRC_channels_editAdmin:
		r = TL_channels_editAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
		}

	case CRC_channels_editTitle:
		r = TL_channels_editTitle{
			m.Object(),
			m.String(),
		}

	case CRC_channels_editPhoto:
		r = TL_channels_editPhoto{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_checkUsername:
		r = TL_channels_checkUsername{
			m.Object(),
			m.String(),
		}

	case CRC_channels_updateUsername:
		r = TL_channels_updateUsername{
			m.Object(),
			m.String(),
		}

	case CRC_channels_joinChannel:
		r = TL_channels_joinChannel{
			m.Object(),
		}

	case CRC_channels_leaveChannel:
		r = TL_channels_leaveChannel{
			m.Object(),
		}

	case CRC_channels_inviteToChannel:
		r = TL_channels_inviteToChannel{
			m.Object(),
			m.Vector(),
		}

	case CRC_channels_deleteChannel:
		r = TL_channels_deleteChannel{
			m.Object(),
		}

	case CRC_channels_exportMessageLink:
		r = TL_channels_exportMessageLink{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case CRC_channels_toggleSignatures:
		r = TL_channels_toggleSignatures{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_getAdminedPublicChannels:
		var flags int32
		r = TL_channels_getAdminedPublicChannels{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
		}

	case CRC_channels_editBanned:
		r = TL_channels_editBanned{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CRC_channels_getAdminLog:
		var flags int32
		r = TL_channels_getAdminLog{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedVector(flags, 1),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case CRC_channels_setStickers:
		r = TL_channels_setStickers{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_readMessageContents:
		r = TL_channels_readMessageContents{
			m.Object(),
			m.VectorInt(),
		}

	case CRC_channels_deleteHistory:
		r = TL_channels_deleteHistory{
			m.Object(),
			m.Int(),
		}

	case CRC_channels_togglePreHistoryHidden:
		r = TL_channels_togglePreHistoryHidden{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_getLeftChannels:
		r = TL_channels_getLeftChannels{
			m.Int(),
		}

	case CRC_channels_getGroupsForDiscussion:
		r = TL_channels_getGroupsForDiscussion{}

	case CRC_channels_setDiscussionGroup:
		r = TL_channels_setDiscussionGroup{
			m.Object(),
			m.Object(),
		}

	case CRC_channels_editCreator:
		r = TL_channels_editCreator{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case CRC_channels_editLocation:
		r = TL_channels_editLocation{
			m.Object(),
			m.Object(),
			m.String(),
		}

	case CRC_channels_toggleSlowMode:
		r = TL_channels_toggleSlowMode{
			m.Object(),
			m.Int(),
		}

	case CRC_channels_getInactiveChannels:
		r = TL_channels_getInactiveChannels{}

	case CRC_bots_sendCustomRequest:
		r = TL_bots_sendCustomRequest{
			m.String(),
			m.Object(),
		}

	case CRC_bots_answerWebhookJSONQuery:
		r = TL_bots_answerWebhookJSONQuery{
			m.Long(),
			m.Object(),
		}

	case CRC_bots_setBotCommands:
		r = TL_bots_setBotCommands{
			m.Vector(),
		}

	case CRC_payments_getPaymentForm:
		r = TL_payments_getPaymentForm{
			m.Int(),
		}

	case CRC_payments_getPaymentReceipt:
		r = TL_payments_getPaymentReceipt{
			m.Int(),
		}

	case CRC_payments_validateRequestedInfo:
		var flags int32
		r = TL_payments_validateRequestedInfo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Object(),
		}

	case CRC_payments_sendPaymentForm:
		var flags int32
		r = TL_payments_sendPaymentForm{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case CRC_payments_getSavedInfo:
		r = TL_payments_getSavedInfo{}

	case CRC_payments_clearSavedInfo:
		var flags int32
		r = TL_payments_clearSavedInfo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
		}

	case CRC_payments_getBankCardData:
		r = TL_payments_getBankCardData{
			m.String(),
		}

	case CRC_stickers_createStickerSet:
		var flags int32
		r = TL_stickers_createStickerSet{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.Vector(),
		}

	case CRC_stickers_removeStickerFromSet:
		r = TL_stickers_removeStickerFromSet{
			m.Object(),
		}

	case CRC_stickers_changeStickerPosition:
		r = TL_stickers_changeStickerPosition{
			m.Object(),
			m.Int(),
		}

	case CRC_stickers_addStickerToSet:
		r = TL_stickers_addStickerToSet{
			m.Object(),
			m.Object(),
		}

	case CRC_stickers_setStickerSetThumb:
		r = TL_stickers_setStickerSetThumb{
			m.Object(),
			m.Object(),
		}

	case CRC_phone_getCallConfig:
		r = TL_phone_getCallConfig{}

	case CRC_phone_requestCall:
		var flags int32
		r = TL_phone_requestCall{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_phone_acceptCall:
		r = TL_phone_acceptCall{
			m.Object(),
			m.StringBytes(),
			m.Object(),
		}

	case CRC_phone_confirmCall:
		r = TL_phone_confirmCall{
			m.Object(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
		}

	case CRC_phone_receivedCall:
		r = TL_phone_receivedCall{
			m.Object(),
		}

	case CRC_phone_discardCall:
		var flags int32
		r = TL_phone_discardCall{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.Object(),
			m.Long(),
		}

	case CRC_phone_setCallRating:
		var flags int32
		r = TL_phone_setCallRating{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.String(),
		}

	case CRC_phone_saveCallDebug:
		r = TL_phone_saveCallDebug{
			m.Object(),
			m.Object(),
		}

	case CRC_phone_sendSignalingData:
		r = TL_phone_sendSignalingData{
			m.Object(),
			m.StringBytes(),
		}

	case CRC_langpack_getLangPack:
		r = TL_langpack_getLangPack{
			m.String(),
			m.String(),
		}

	case CRC_langpack_getStrings:
		r = TL_langpack_getStrings{
			m.String(),
			m.String(),
			m.VectorString(),
		}

	case CRC_langpack_getDifference:
		r = TL_langpack_getDifference{
			m.String(),
			m.String(),
			m.Int(),
		}

	case CRC_langpack_getLanguages:
		r = TL_langpack_getLanguages{
			m.String(),
		}

	case CRC_langpack_getLanguage:
		r = TL_langpack_getLanguage{
			m.String(),
			m.String(),
		}

	case CRC_folders_editPeerFolders:
		r = TL_folders_editPeerFolders{
			m.Vector(),
		}

	case CRC_folders_deleteFolder:
		r = TL_folders_deleteFolder{
			m.Int(),
		}

	case CRC_stats_getBroadcastStats:
		var flags int32
		r = TL_stats_getBroadcastStats{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case CRC_stats_loadAsyncGraph:
		var flags int32
		r = TL_stats_loadAsyncGraph{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedLong(flags, 0),
		}

	default:
		m.err = merry.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}
