package types

type State struct {
	Handler   string `json:"channel,omitempty"`
	Question  string `json:"user,omitempty"`
	ID        int    `json:"ts,omitempty"`
	Count     int    `json:"thread_ts,omitempty"`
	Reply     string
	DbType    string
	Artifact  string
	Cluster   string
	Sentiment string
	Action    string
}

func NewState() State {
	return State{
		ID:        0,
		Count:     0,
		Artifact:  "unknown",
		Cluster:   "unknown",
		Sentiment: "positive",
		Action:    "ask",
		Question:  "What do you want to do today?",
	}
}

// Apply changes the State according to the msg recieved
func (s *State) Apply(msg string) error {
	// s.Count = s.Count + 1
	s.Reply = "Duck Duck "
	return nil
}

// Apply changes the State according to the msg recieved
func (s *State) ApplyYes() error {
	// s.Count = s.Count + 1
	s.Reply = "Duck Duck "
	return nil
}
