package analysis

import (
	"golsp/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (state *State) UpdateDocument(uri string, text string) {
	state.Documents[uri] = text
}

func (state *State) GetCodeActionResult(uri string, actionRange lsp.Range) []lsp.CodeAction {
	var codeAction []lsp.CodeAction
	addFirstLineAction := lsp.CodeAction{
		Title: "addTextInFistLine",
		Edit: &lsp.WorkSpaceEdit{
			Changes: map[string][]lsp.TextEdit{
				uri: {
					{
						Range: lsp.Range{
							Start: lsp.Position{
								Line:      0,
								Character: 0,
							},
							End: lsp.Position{
								Line:      0,
								Character: 0,
							},
						},
						NewText: "this is firstlineText with codeAction\n",
					},
				},
			},
		},
	}
	addWordInCurrentPosition := lsp.CodeAction{
		Title: "addWordInCurrentPosition ",
		Edit: &lsp.WorkSpaceEdit{
			Changes: map[string][]lsp.TextEdit{
				uri: {
					{
						Range: lsp.Range{
							Start: lsp.Position{
								Line:      actionRange.Start.Line,
								Character: actionRange.Start.Character,
							},
							End: lsp.Position{
								Line:      actionRange.Start.Line,
								Character: actionRange.End.Character,
							},
						},
						NewText: " I love world ",
					},
				},
			},
		},
	}
	codeAction = append(codeAction, addFirstLineAction,addWordInCurrentPosition)

	return codeAction
}
