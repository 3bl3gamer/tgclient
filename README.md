# TG Client
It is a pure-Golang MTProto client for Telegram API.

Initially based on https://github.com/sdidyk/mtproto and https://github.com/ronaksoft/mtproto.

TL layer: 114

Consists of two parts:
  * mtproto — core part for API interation;
  * tgclient — wrapper over mtproto with some useful logic like file download.

## API
### NewMTProtoExt
`dialer` will open connection(s) to TG server(s).
May be just `&net.Dialer{}` for bare TCP connection or (for example)
`dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)`
for socks5 proxy.

## TODO
* if error occures while performing request to `TL_invokeWithLayer` in `Connect()`, two `TL_invokeWithLayer` may be sent. Nothing bad happens though.
