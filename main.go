package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golsp/analysis"
	"golsp/lsp"
	"golsp/rpc"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	logger := getLogger("/home/allworldg/LearnLsp/log/log.txt")
	state := analysis.NewState()
	for scanner.Scan() {
		method, contents, err := rpc.DecodeMessage(scanner.Bytes())
		if err != nil {
			logger.Println(err)
			continue
		}
		handleMessage(logger, state, method, contents)
	}
}
func handleMessage(logger *log.Logger, state analysis.State, method string, content []byte) {
	logger.Printf("Receive the method: %s\n", method)
	switch method {
	case "initialize":
		var initializeRequest lsp.InitializeRequest
		err := json.Unmarshal(content, &initializeRequest)
		if err != nil {
			logger.Printf("cannot unmarshal the content you give,%s\n", err)
			return
		}
		version := initializeRequest.Params.ClientInfo.Version
		name := initializeRequest.Params.ClientInfo.Name
		logger.Printf("method is %s,the clientName is %s,clientVersion is %s\n", method, name, version)

		//reply to the client
		response := lsp.NewInitializeResponse(initializeRequest.Id)
		strResponse := rpc.EncodeMessage(response)
		writer := os.Stdout
		writer.Write([]byte(strResponse))
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("cannnot unmarshal the didOpen content: %s\n", err)
			return
		}
		state.UpdateDocument(request.Params.TextDocument.Uri, request.Params.TextDocument.Text)
		logger.Printf("uri is %s, languageId is %s, version is %d\n",
			request.Params.TextDocument.Uri,
			request.Params.TextDocument.LanguageId,
			request.Params.TextDocument.Version)
	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("cannot unmarshal the didChange content:%s\n", err)
			return
		}
		for _, v := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.Uri, v.Text)
			logger.Printf("didChange uri is %s,text is %s", request.Params.TextDocument.Uri, v.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("cannot unmarshal the hover request:%s\n", err)
			return
		}
		logger.Printf("hover position line  is %d, and character is %d", request.Params.Position.Line, request.Params.Position.Character)
		writer := os.Stdout
		writer.Write([]byte(rpc.EncodeMessage(lsp.HoverResponse{
			Response: lsp.Response{
				Rpc: "2.0",
				Id:  &request.Id,
			},
			Result: lsp.HoverResult{
				Contents: "this is myhover",
			},
		})))
	case "textDocument/definition":
		var request lsp.GotoDefinitionDocumentRequest
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("cannot unmarshal goto definition: %s\n", err)
		}
		logger.Printf("gotodefinition position line is %d, character is %d", request.Params.Position.Line, request.Params.Position.Character)
		writer := os.Stdout
		writer.Write([]byte(rpc.EncodeMessage(lsp.GotoDefinitionDocumentResponse{
			Result: lsp.Location{
				Uri: request.Params.TextDocument.Uri,
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      request.Params.Position.Line - 1,
						Character: request.Params.Position.Character,
					},
					End: lsp.Position{
						Line:      request.Params.Position.Line - 1,
						Character: request.Params.Position.Character,
					},
				},
			},
			Response: lsp.Response{
				Rpc: "2.0",
				Id:  &request.Id,
			},
		})))
	}

}
func getLogger(filePath string) *log.Logger {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		panic("u do not give a right filePath")
	}
	logger := log.New(file, "golsp ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
