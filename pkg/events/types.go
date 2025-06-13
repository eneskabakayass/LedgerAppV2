package events

import "time"

type Event interface {
	Name() string
	Timestamp() time.Time
}

type BaseEvent struct {
	EventName string
	Occurred  time.Time
}

func (e BaseEvent) Name() string         { return e.EventName }
func (e BaseEvent) Timestamp() time.Time { return e.Occurred }

type TransactionCredited struct {
	BaseEvent
	UserID string
	Amount float64
}

func NewTransactionCredited(userID string, amount float64) TransactionCredited {
	return TransactionCredited{
		BaseEvent: BaseEvent{
			EventName: "TransactionCredited",
			Occurred:  time.Now(),
		},
		UserID: userID,
		Amount: amount,
	}
}
