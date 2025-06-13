package events

type State struct {
	UserID  string
	Balance float64
}

func RebuildState(events []Event) map[string]*State {
	stateMap := make(map[string]*State)

	for _, e := range events {
		switch evt := e.(type) {
		case TransactionCredited:
			s, ok := stateMap[evt.UserID]
			if !ok {
				s = &State{UserID: evt.UserID}
				stateMap[evt.UserID] = s
			}
			s.Balance += evt.Amount
		}
	}

	return stateMap
}
