package tgclient

import (
	"mtproto"
	"reflect"

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
	updatesState         *mtproto.TL_updates_state
	handleUpdateExternal UpdateHandler
	log                  Logger
	Downloader
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
		updatesState:         &mtproto.TL_updates_state{},
		handleUpdateExternal: handleUpdate,
		log:                  log,
	}
	client.Downloader = *NewDownloader(client)

	mt.SetEventsHandler(client.handleEvent)
	for i := 0; i < 4; i++ {
		go client.partsDownloadRoutine()
	}
	return client, nil
}

func (c *TGClient) handleEvent(eventObj mtproto.TL) {
	switch event := eventObj.(type) {
	case mtproto.TL_updatesTooLong:
		//TODO: what?
		// Too many updates, it is necessary to execute updates.getDifference.
		// https://core.telegram.org/constructor/updatesTooLong
		c.log.Warning("[WARN] updates too long")
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
