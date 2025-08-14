package lsp

type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParam `json:"params"`
}

type DidChangeTextDocumentParam struct {
	TextDocument   VersionedTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []ContentChanges                `json:"contentChanges"`
}

type ContentChanges struct {
	Text string `json:"text"`
}
