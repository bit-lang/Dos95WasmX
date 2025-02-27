package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"sync"
	"github.com/bit-lang/Dos95WasmX/staticserver"
)

const (
	server_port  = ":8088"
	prefixStr    = "$>"
	exit_command = "stop"
)

//go:embed all:assets
var allEmbeds embed.FS

func getAssets(folder string) (fs.FS, error) {
	return fs.Sub(allEmbeds, folder)
}

func main() {
	fmt.Println("starting dos95")
	embeddedFiles, err := getAssets("assets")
	if err != nil {
		fmt.Printf("init getAssets() error: %v\n", err)
		return
	}
	gtx, globalcancelfunc := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	_, svrErr := staticserver.SetupEmbeddingSvr(gtx, &wg, server_port, "/", exit_command, embeddedFiles)
	if svrErr != nil {
		globalcancelfunc()
		fmt.Printf("init error: %v\n", svrErr)
		return
	}
	fmt.Println(" based on doswasmx")
	fmt.Printf("\n%s usage: browser (Chrome or Edge) to open http://localhost%s \n", prefixStr, server_port)
	fmt.Printf("%s to stop, type %s after the following prompt \n\n%s", prefixStr, exit_command, prefixStr)
	wg.Wait() // wait for server to complete
	fmt.Println("dos95 stopped")
}
