package tgclient

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"

	"github.com/3bl3gamer/tgclient/mtproto"
	"github.com/ansel1/merry/v2"
	"golang.org/x/net/proxy"
)

type TGClient struct {
	mt                   *mtproto.MTProto
	updatesState         *mtproto.TL_updates_state
	handleUpdateExternal UpdateHandler
	log                  mtproto.Logger
	extraData
	Downloader
}

type UpdateHandler func(mtproto.TL)

func NewTGClient(appID int32, appHash string, logHnd mtproto.LogHandler) *TGClient {
	var exPath string
	ex, err := os.Executable()
	if err != nil {
		mtproto.Logger{logHnd}.Error(err, "failed to get executable file path")
		exPath = "."
	} else {
		exPath = filepath.Dir(ex)
	}
	sessStore := &mtproto.SessFileStore{exPath + "/tg.session"}

	cfg := &mtproto.AppConfig{
		AppID:          appID,
		AppHash:        appHash,
		AppVersion:     "0.0.1",
		DeviceModel:    "Unknown",
		SystemVersion:  runtime.GOOS + "/" + runtime.GOARCH,
		SystemLangCode: "en",
		LangPack:       "",
		LangCode:       "en",
	}
	return NewTGClientExt(cfg, sessStore, logHnd, nil)
}

func NewTGClientExt(cfg *mtproto.AppConfig, sessStore mtproto.SessionStore, logHnd mtproto.LogHandler, daler proxy.Dialer) *TGClient {
	mt := mtproto.NewMTProtoExt(mtproto.MTParams{
		AppConfig:  cfg,
		SessStore:  sessStore,
		LogHandler: logHnd,
		ConnDialer: daler,
	})

	client := &TGClient{
		mt:           mt,
		updatesState: &mtproto.TL_updates_state{},
		log:          mtproto.Logger{logHnd},
	}
	client.extraData = *newExtraData(client)

	mt.SetEventsHandler(client.handleEvent)
	return client
}

func (c *TGClient) SetUpdateHandler(handleUpdate UpdateHandler) {
	c.handleUpdateExternal = handleUpdate
}

func (c *TGClient) InitAndConnect() error {
	c.Downloader.Start(c)
	return merry.Wrap(c.mt.InitSessAndConnect())
}

func (c *TGClient) Disconnect() error {
	err := c.Downloader.Stop()
	err = c.mt.Disconnect()
	return merry.Wrap(err)
}

func (c *TGClient) handleEvent(eventObj mtproto.TL) {
	switch event := eventObj.(type) {
	case mtproto.TL_updatesTooLong:
		//TODO: what?
		// Too many updates, it is necessary to execute updates.getDifference.
		// https://core.telegram.org/constructor/updatesTooLong
		c.log.Warn("updates too long")
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
		c.log.Warn(mtproto.UnexpectedTL("event", eventObj))
	}
}

func (e *TGClient) handleUpdate(obj mtproto.TL) {
	value := reflect.ValueOf(obj).FieldByName("Pts")
	if value != (reflect.Value{}) {
		e.updatesState.Pts = int32(value.Int())
	}
	if e.handleUpdateExternal != nil {
		e.handleUpdateExternal(obj)
	}
}

func (c *TGClient) AuthExt(authData mtproto.AuthDataProvider, message mtproto.TLReq) (mtproto.TL, error) {
	for {
		res := c.mt.SendSync(message)
		if mtproto.IsErrorType(res, mtproto.TL_ErrUnauthorized) { //AUTH_KEY_UNREGISTERED SESSION_REVOKED SESSION_EXPIRED
			if err := c.mt.Auth(authData); err != nil {
				return nil, merry.Wrap(err)
			}
			continue
		}
		c.log.Info("Seems authed.")
		return res, nil
	}
}

func (c *TGClient) AuthAndInitEvents(authData mtproto.AuthDataProvider) error {
	// after reconnection TG *sometimes* stops sending updates
	c.mt.SetReconnectionHandler(func() error {
		res := c.SendSync(mtproto.TL_updates_getState{})
		if _, ok := res.(mtproto.TL_updates_state); !ok {
			return mtproto.WrongRespError(res)
		}
		return nil
	})

	res, err := c.AuthExt(authData, mtproto.TL_updates_getState{})
	if err != nil {
		return merry.Wrap(err)
	}
	if _, ok := res.(mtproto.TL_updates_state); !ok {
		return mtproto.WrongRespError(res)
	}
	return nil
}

func (c *TGClient) SendSync(msg mtproto.TLReq) mtproto.TL {
	return c.mt.SendSync(msg)
}

func (c *TGClient) SendSyncRetry(
	msg mtproto.TLReq, failRetryInterval time.Duration,
	floodNumShortRetries int, floodMaxWait time.Duration,
) mtproto.TL {
	return c.mt.SendSyncRetry(msg, failRetryInterval, floodNumShortRetries, floodMaxWait)
}
