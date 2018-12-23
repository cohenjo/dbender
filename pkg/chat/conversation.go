package chat

import (
	"github.com/cohenjo/dbender/pkg/types"
)

// Converse tries to carry a simple conversation
/**
* conversation is a state machine where every sentence takes us from state to state
 */
func Converse(session *types.UserSession, message string) (string, error) {
	state := session.CurrentState
	err := state.Apply(message)
	if err != nil {
		return "all failed", err
	}
	return "ok: " + state.Reply, nil

}
