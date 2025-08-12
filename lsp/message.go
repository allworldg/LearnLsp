package lsp

type Request struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method string `json:"method"`
}

type Response struct{
	Rpc string `json:"jsonrpc"`
	Id *int `json:"id,omitemty"`
}

type Notification struct{
	Rpc string `json:"jsonrpc"`
	Method string `json:"method"`
}
