package model

type EventType int

const (
	GameStartedEventType    EventType = 1
	GameEndedEventType      EventType = 2
	MeetingStartedEventType EventType = 3
	MeetingEndedEventType   EventType = 4
)

type Event struct {
	Type    EventType
	GuildID string
}
