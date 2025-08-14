package lsp

type TextDocumentItem struct{
	Uri string `json:"uri"`
	LanguageId string `json:"languageId"`
	Version int `json:"version"` 
	Text string `json:"text"`
}

type TextDocumentIdentifier struct{
	Uri string `json:"uri"`
}
type VersionedTextDocumentIdentifier struct{
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position Position `json:"position"`
}
type Position struct {
	Line int `json:"line"`
	Character int `json:"character"`
}

