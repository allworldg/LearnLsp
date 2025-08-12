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
	var initializeRequest lsp.InitializeRequest
	switch method {
	case "initialize":
		err := json.Unmarshal(content, &initializeRequest)
		if err != nil {
			logger.Printf("cannot unmarshal the content you give,%s\n", err)
		}
		version := initializeRequest.Params.ClientInfo.Version
		name := initializeRequest.Params.ClientInfo.Name
		logger.Printf("method is %s,the clientName is %s,clientVersion is %s\n", method, name, version)
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
