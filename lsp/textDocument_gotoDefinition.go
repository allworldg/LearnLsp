package lsp

type GotoDefinitionDocumentRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}
type DefinitionParams struct{
	TextDocumentPositionParams
}

type GotoDefinitionDocumentResponse struct{
	Response
	Result Location `json:"result"`
}

