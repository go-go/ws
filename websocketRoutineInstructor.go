package ws

import (
	"log"
	"strings"
)

type wsRoutineInstructor struct {
	routine WsRoutine
	context *WsContext
}

// this function starts the routine instructor
// specifically runs in its own routine
// TODO add waiter for managing the go routines
func (instructor *wsRoutineInstructor) start() {
	go func() {
		// send OnConnect event
		instructor.routine.OnConnect(instructor.context)
		// start reading up the data
		go instructor.read()
	}()
}

func (instructor *wsRoutineInstructor) read() {
	defer instructor.context.Close()
	for {
		log.Println("selecting a messsage")
		select {
		case message := <-instructor.context.reader:
			log.Println("got data", string(*message))
			instructor.routine.OnTextMessage(string(*message))
		case err := <-instructor.context.spoiler:
			if err == nil {
				continue
			}
			if strings.Contains(err.Error(), "close") {
				instructor.routine.OnClose()
				return
			}
			log.Println(err.Error())
			instructor.routine.OnError(err)
			return
		}
	}
}
