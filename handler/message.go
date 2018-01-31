package handler

import "time"

// Message is just what is is
type Message struct {
	messenger Messenger

	payload struct {
		encryptedBody string

		sentAt time.Time
	}
}
