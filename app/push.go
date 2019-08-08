package app

import (
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func channelName(token uuid.UUID) string {
	return "/push/" + token.String()
}

func (sess *session) pushMessage(event string) {
	channel := channelName(sess.pushToken)
	log.Println(channel)
	if !sess.server.sse.HasChannel(channel) {
		// no need to send message if client is not connected
		return
	}

	msgID, err := uuid.NewRandom()
	if err != nil {
		log.Println("could not generate msgID")
		return
	}
	sess.server.sse.SendMessage(channel, sse.NewMessage(msgID.String(), time.Now().String(), event))
}

// Setups push messages via server-sent events (SSE).
// This is used for events which need a notification of a client by the server.
// It is only used to alert the client that is has to do some action. e.g. "load new messages"
// the actual action is still done by REST.
// It is optional to use the push messages. A client could also Pull the data from the server frequently.
// A push message consists only of a event name and a timestamp
func (s *Server) setupPushRegister(router chi.Router) {
	router.Get("/push/register", func(w http.ResponseWriter, r *http.Request) {
		token, ok := token(w, r)
		if !ok {
			return
		}

		pushToken, err := uuid.NewRandom()
		if err != nil {
			log.Println("could not generate pushToken")
			return
		}

		sess, ok := s.sessions[token]
		if !ok {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		sess.pushToken = pushToken
		w.Write([]byte(pushToken.String()))
	})
}
