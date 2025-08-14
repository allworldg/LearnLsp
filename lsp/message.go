package lsp

type Request struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
}

type Response struct {
	Rpc string `json:"jsonrpc"`
	Id  *int   `json:"id,omitempty"`
}

type Notification struct {
	Rpc    string `json:"jsonrpc"`
	Method string `json:"method"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			Rpc: "2.0",
			Id:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
				HoverProvider:true,
			},
			ServerInfo: ServerInfo{
				Name:    "learnLsp",
				Version: "0.0.0.1-beta",
			},
		},
	}
}
