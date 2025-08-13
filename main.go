package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golsp/lsp"
	"golsp/rpc"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	logger := getLogger("/home/allworldg/LearnLsp/log/log.txt")
	for scanner.Scan() {
		method, contents, err := rpc.DecodeMessage(scanner.Bytes())
		if err != nil {
			logger.Println(err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}
func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("Receive the method: %s\n", method)
	switch method {
	case "initialize":
		var initializeRequest lsp.InitializeRequest
		err := json.Unmarshal(content, &initializeRequest)
		if err != nil {
			logger.Printf("cannot unmarshal the content you give,%s\n", err)
		}
		version := initializeRequest.Params.ClientInfo.Version
		name := initializeRequest.Params.ClientInfo.Name
		logger.Printf("method is %s,the clientName is %s,clientVersion is %s\n", method, name, version)

		//reply to the client
		response := lsp.NewInitializeResponse(initializeRequest.Id)
		strResponse := rpc.EncodeMessage(response)
		writer := os.Stdout
		writer.Write([]byte(strResponse))
		logger.Printf("reply the request")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		err := json.Unmarshal(content, &request)
		if err != nil {
			logger.Printf("cannnot unmarshal the didOpen content: %s\n", err)
		}
		textDocumentItem :=request.Params.TextDocument 
		logger.Printf("uri is %s, languageId is %s, version is %d, text is %s\n",
			textDocumentItem.Uri, textDocumentItem.LanguageId,textDocumentItem.Version,textDocumentItem.Text)
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
