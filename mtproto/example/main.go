package main

import (
	"flag"
	"log"
	"math/rand"
	"os"

	"github.com/3bl3gamer/tgclient/mtproto"
	"github.com/ansel1/merry/v2"
)

func main() {
	appID := flag.Int("app_id", 0, "app id")
	appHash := flag.String("app_hash", "", "app hash")
	flag.Parse()

	if *appID == 0 || *appHash == "" {
		println("App ID and hash are required!")
		flag.Usage()
		os.Exit(2)
	}

	if err := start(int32(*appID), *appHash); err != nil {
		log.Fatal(merry.Details(err))
	}
}

func start(appID int32, appHash string) error {
	// simple:
	m := mtproto.NewMTProto(appID, appHash)

	// with proxy (golang.org/x/net/proxy):
	// dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	// if err != nil {
	// 	return merry.Wrap(err)
	// }
	// m := mtproto.NewMTProtoExt(mtproto.MTParams{AppID: appID, AppHash: appHash, ConnDialer: dialer})

	if err := m.InitSessAndConnect(); err != nil {
		return merry.Wrap(err)
	}

	for {
		res := m.SendSync(mtproto.TL_updates_getState{})
		if mtproto.IsErrorType(res, mtproto.TL_ErrUnauthorized) { //AUTH_KEY_UNREGISTERED SESSION_REVOKED SESSION_EXPIRED
			if err := m.Auth(mtproto.ScanfAuthDataProvider{}); err != nil {
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
	log.Println("Seems authed.")

	if err := m.GetContacts(); err != nil {
		return merry.Wrap(err)
	}

	{
		log.Println("Requesting self via TL_users_getUsers (vector response test)")
		res := m.SendSync(mtproto.TL_users_getUsers{
			ID: []mtproto.TL{mtproto.TL_inputUserSelf{}},
		})
		self := res.(mtproto.VectorObject)[0].(mtproto.TL_user)
		log.Printf("done: #%d %s", self.ID, mtproto.DerefOr(self.FirstName, ""))
	}

	{
		log.Println("Requesting self via TL_users_getUsers and TL_invokeWithLayer (polymorphic vector response test)")
		res := m.SendSync(mtproto.TL_invokeWithLayer{
			Layer: mtproto.TL_Layer,
			Query: mtproto.TL_users_getUsers{
				ID: []mtproto.TL{mtproto.TL_inputUserSelf{}},
			},
		})
		self := res.(mtproto.VectorObject)[0].(mtproto.TL_user)
		log.Printf("done: #%d %s", self.ID, mtproto.DerefOr(self.FirstName, ""))
	}

	{
		log.Println("Sending self message")
		randomID := rand.Int63()
		res := m.SendSync(mtproto.TL_messages_sendMessage{
			Peer:     mtproto.TL_inputPeerSelf{},
			Message:  "tgclient test message",
			RandomID: randomID,
		})
		updates := res.(mtproto.TL_updates)
		var msgID int32 = 0
		for _, updTL := range updates.Updates {
			if upd, ok := updTL.(mtproto.TL_updateMessageID); ok {
				if upd.RandomID == randomID {
					msgID = upd.ID
					break
				}
			}
		}
		if msgID == 0 {
			return merry.Errorf("new message ID not found in: %#v", updates.Updates)
		}
		log.Printf("done: msg ID = %d, removing now", msgID)
		res = m.SendSync(mtproto.TL_messages_deleteMessages{
			ID: []int32{msgID},
		})
		_ = res.(mtproto.TL_messages_affectedMessages)
		log.Printf("done: message #%d removed", msgID)
	}

	{
		log.Println("Reconnecting")
		if err := m.Reconnect(); err != nil {
			return merry.Wrap(err)
		}
		log.Println("done")
	}

	{
		curDC := m.CopySession().DCID
		newDC := int32(2)
		if curDC == newDC {
			newDC = 1
		}
		log.Printf("Connecting to another DC (%d -> %d)", curDC, newDC)
		m1, err := m.NewConnection(newDC)
		if err != nil {
			return merry.Wrap(err)
		}
		res := m1.SendSync(mtproto.TL_help_getConfig{})
		log.Printf("done, date from TL_config: %d", res.(mtproto.TL_config).Date)
	}

	log.Print("example finished, pausing forever")
	<-chan bool(nil)
	return nil
}
