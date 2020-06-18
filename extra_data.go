package tgclient

import (
	"sync"

	"github.com/3bl3gamer/tgclient/mtproto"
)

type extraData struct {
	tg       *TGClient
	mutex    *sync.RWMutex
	users    map[int32]*mtproto.TL_user
	chats    map[int32]*mtproto.TL_chat
	channels map[int32]*mtproto.TL_channel
}

func newExtraData(tg *TGClient) *extraData {
	return &extraData{
		tg:       tg,
		mutex:    &sync.RWMutex{},
		users:    make(map[int32]*mtproto.TL_user),
		chats:    make(map[int32]*mtproto.TL_chat),
		channels: make(map[int32]*mtproto.TL_channel),
	}
}

func (e *extraData) rememberEventExtraData(objs []mtproto.TL) {
	if len(objs) == 0 {
		return
	}
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for _, obj := range objs {
		switch x := obj.(type) {
		case mtproto.TL_user:
			e.users[x.ID] = &x
			e.tg.log.Debug("extra: user: %d %s", x.ID, x.Username)
		case mtproto.TL_chat:
			e.chats[x.ID] = &x
			e.tg.log.Debug("extra: chat: %d %s", x.ID, x.Title)
		case mtproto.TL_channel:
			e.channels[x.ID] = &x
			e.tg.log.Debug("extra: channel: %d %s", x.ID, x.Username)
		default:
			e.tg.log.Warn(mtproto.UnexpectedTL("extra data", obj))
		}
	}
}

func (e *extraData) FindExtraUser(userID int32) *mtproto.TL_user {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.users[userID]
}
func (e *extraData) FindExtraChat(chatID int32) *mtproto.TL_chat {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.chats[chatID]
}
func (e *extraData) FindExtraChannel(channelID int32) *mtproto.TL_channel {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.channels[channelID]
}
