package lsp

type TextDocumentItem struct {
	Uri        string `json:"uri"`
	LanguageId string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type TextDocumentIdentifier struct {
	Uri string `json:"uri"`
}
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
type Location struct {
	Uri   string `json:"uri"`
	Range Range  `json:"range"`
}
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
type CodeActionContext struct {
	// Diagnostics []Diagnostic `json:"diagnostics"`
}

type WorkSpaceEdit struct {
	Changes map[string][]TextEdit `json:"changes"`
}
type TextEdit struct{
	Range Range `json:"range"` 
	NewText string `json:"newText"`
}
type Command struct {
	Title string `json:"title"`
	Command string `json:"command"`
	Arguments []any `json:"arguments,omitempty"`
}

type Diagnostic struct {
	Range    Range  `json:"range"`
	Severity int    `json:"severity"`
	Source   string `json:"source"`
	Message  string `json:"message"`
}
