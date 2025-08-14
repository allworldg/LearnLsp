package lsp

type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentParam `json:"params"`
}

type DidOpenTextDocumentParam struct {
	TextDocument   TextDocumentItem `json:"textDocument"`
}

