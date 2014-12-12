package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	//TODO change this
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Handler(routine WsRoutine) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// upgrading the websocket connection
		conn, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			log.Println("Error while upgrading the connection", err.Error())
			conn.Close()
			return
		}
		// now everything looks good
		// creating context
		context := &WsContext{conn: conn}
		context.start()
		// kick off the instructor
		instructor := wsRoutineInstructor{routine: routine, context: context}
		instructor.start()
	}
}
