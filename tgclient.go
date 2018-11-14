package tgclient

import (
	"mtproto"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/ansel1/merry"
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
	client.Downloader = *NewDownloader(client)
	client.extraData = *newExtraData(client)

	mt.SetEventsHandler(client.handleEvent)
	for i := 0; i < 4; i++ {
		go client.partsDownloadRoutine()
	}
	return client
}

func (c *TGClient) SetUpdateHandler(handleUpdate UpdateHandler) {
	c.handleUpdateExternal = handleUpdate
}

func (c *TGClient) InitAndConnect() error {
	return merry.Wrap(c.mt.InitSessAndConnect())
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

func (c *TGClient) AuthAndInitEvents(authData mtproto.AuthDataProvider) error {
	for {
		res := c.mt.SendSync(mtproto.TL_updates_getState{})
		if mtproto.IsErrorType(res, mtproto.TL_ErrUnauthorized) { //AUTH_KEY_UNREGISTERED SESSION_REVOKED SESSION_EXPIRED
			if err := c.mt.Auth(authData); err != nil {
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

func (c *TGClient) SendSync(msg mtproto.TLReq) mtproto.TL {
	return c.mt.SendSync(msg)
}
