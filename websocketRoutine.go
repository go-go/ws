package ws

type WsRoutine interface {
	OnConnect(context *WsContext)
	OnTextMessage(message string) error
	OnBinaryMessage(message []byte) error
	OnError(err error)
	OnClose()
}
