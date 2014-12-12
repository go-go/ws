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

## Licence

If you are gonna fork it then please do read the licence from Gorilla Developers for the websocket project.

Gorilla:

Copyright (c) 2013 The Gorilla WebSocket Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

  Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

  Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

