package whatsapp

import "time"

type rawMessage struct {
	system bool
	body   string
}

type parsedRawMessage struct {
	date   string
	time   string
	ampm   *string
	author *string
	body   string
}

type Message struct {
	// The author of the message. Will be null for messages without an author (system messages).
	Author *string `json:"sender"`

	// The body of the message.
	Body string `json:"content"`

	// The bodyType of the message.
	BodyType string `json:"contentType"`

	// The date of the message.
	Date time.Time `json:"sentAt"`
}
