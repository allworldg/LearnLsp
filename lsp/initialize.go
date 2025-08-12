package lsp

type InitializeRequest struct{
	Request
	Params InitialRequestParams `json:"params"`
}
type InitialRequestParams struct{
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct{
	Name string `json:"name"`
	Version string `json:"version"`
}

type InitilizeResult struct{
	Capabilities ServerCapabilities `json:""`
} 
