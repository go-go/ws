# Ws

Gorilla has written an amazing library for managing websockets with go. But we feel that there are some little gaps needs to be filled so as to make it easier and quicker to integrate. And thats why this library, ws.

This library aims to handle all the routine and channels related functionality of websockets and you just have to worry about the business logic over it. So this library is designed by keeping evented io in reference (for example the way you work with netty).

You have to write routines that implement <pre>WsRoutine</pre> interface and just pass it to the handler like 
```
http.HandleFunc("/ws", ws.Handler(ws.WsRoutine(&yourRoutine)))
```

## WsRoutine Interface
```
type WsRoutine interface {
  OnConnect(context *WsContext)
  OnTextMessage(message string) error
  OnBinaryMessage(message []byte) error
  OnError(err error)
  OnClose()
}
```

WsContext gives you a way to write data onto the connection. 
```
context.WriteMessage("data")
```
or
```
context.WriteBinaryMessage([]byte{})
```

## It's still in beta
We are building this as a part of one of our product, but its still baking and not complete yet. But we are open to get your thoughts about it.

contact: akshay@rainingclouds.com
site: http://rainingclouds.com



