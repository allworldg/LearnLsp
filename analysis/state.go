package analysis

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (state *State) UpdateDocument(uri string, text string) {
	state.Documents[uri] = text
}
