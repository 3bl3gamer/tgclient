**Attention!** v0.166.2 has some important changes, [check them out](https://github.com/3bl3gamer/tgclient/releases/tag/v0.166.2).

# TG Client

It is a pure-Golang MTProto client for Telegram API.

Initially based on https://github.com/sdidyk/mtproto and https://github.com/ronaksoft/mtproto.

TL layer: 183

Consists of two parts:
  * mtproto — core part for API interaction;
  * tgclient — wrapper over mtproto with some useful logic like file download.


## How to use it

You may use [tg-history-dumper](https://github.com/3bl3gamer/tg-history-dumper) repo as an example, in particular this file: https://github.com/3bl3gamer/tg-history-dumper/blob/master/tg.go#L67

### Create

First, `TGClient` object should be created with `tgclient.NewTGClient` or `tgclient.NewTGClientExt` for example:

```go
tg := tgclient.NewTGClient(appID, appHash, &mtproto.SimpleLogHandler{})
```
And if you need more customization:

```go
// Telegram client info
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

// store that will save/read session data
sessStore := &mtproto.SessFileStore{FPath: "tg_session.json"}

// optional dialer (may be nil) for proxying requests
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)

tg := tgclient.NewTGClientExt(cfg, sessStore, logHandler, dialer)
```

### Connect

Then, the connection should be opened:

```go
err := tg.InitAndConnect()
```

### Auth

And client should authenticate with `AuthAndInitEvents` or `AuthExt`, for example:

```go
// object that will... provide auth data:
// phone number, confirmation code and optionally password.
// ScanfAuthDataProvider will just prompt all this from terminal.
authDataProvider := mtproto.ScanfAuthDataProvider{}

err := tgclient.AuthAndInitEvents(authDataProvider)
```

While authing, `AuthAndInitEvents` sends `mtproto.TL_updates_getState` request. Same request will also be sent after each reconnection. It makes TG server send updates to client (like new incoming messages). If you do not need those (maybe you just want to dump your chats history), you may send something different:

```go
authDataProvider := mtproto.ScanfAuthDataProvider{}
payload := mtproto.TL_users_getUsers{ID: []mtproto.TL{mtproto.TL_inputUserSelf{}}}

resp, err := tg.AuthExt(authDataProvider, payload)
// here resp should be a mtproto.VectorObject with one mtproto.TL_user item.
```

### Communicate

With `tg.SendSync` or `tg.SendSyncRetry`. First one just sends request and returns response whatever it will be, so you generally should check if you got what you expected. For example:

```go
res := c.tg.SendSync(mtproto.TL_contacts_resolveUsername{Username: "some chat name"})
peer, ok := res.(mtproto.TL_contacts_resolvedPeer)
if !ok {
  // res will likely be TL_rpc_error with message USERNAME_NOT_OCCUPIED or RPC_CALL_FAIL or other
  return mtproto.WrongRespError(res)
}
for _, chat := range peer.Chats {
    fmt.Printf("%#v\n", chat)
}
```

Often you will receive temporary errors like `RPC_CALL_FAIL` of `FOOLD_WAIT_123` and want to re-send same request after little delay. This is done by
```go
res := tg.SendSyncRetry(request, time.Second, 0, 30*time.Second)
```

### Updates

Pass callback func:

```go
tg.SetUpdateHandler(func(updateTL mtproto.TL) {
    switch update := updateTL.(type) {
        case mtproto.TL_updateUserStatus:
            fmt.Printf("U#%d %T\n", update.UserID, update.Status)
        case mtproto.TL_updateNewChannelMessage:
            ...
        case mtproto.TL_updateEditChannelMessage:
            ...
        ...
        default:
            return mtproto.WrongRespError(updateTL)
        }
    }
})
```


## Updating API schema version (aka layer)

Get new schema from https://core.telegram.org/schema (remove definitions for `int`, `long`, `double`, `string`, `vector`, `int128` and `int256`: they are hard-coded and must not be generated). If it is ~~still~~ outdated check other repos (like official ones), some useful links are at the top of [generate_tl_schema.go](https://github.com/3bl3gamer/tgclient/blob/master/mtproto/scheme/generate_tl_schema.go).

Place new `.tl` file to `mtproto/scheme` folder.

Update `//go:generate` command in `mtproto/mtproto.go`. It should be

```go
//go:generate go run scheme/generate_tl_schema.go <layer> scheme/<file>.tl tl_schema.go
```

Run `go generate` in `mtproto` folder.

Update this README.


## TODO
* if error occures while performing request to `TL_invokeWithLayer` in `Connect()`, two `TL_invokeWithLayer` may be sent. Nothing bad happens though.
