package staticserver

import (
	"bufio"
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

const defaultStartupDelay = 1800 // milliseconds

type staticServer struct {
	svrPort string
	server  *http.Server
	consoleCmdChan chan struct{}
	osSigChan      chan os.Signal
}

func SetupEmbeddingSvr(gctx context.Context, gwg *sync.WaitGroup, svr_PORT, svr_Path, exitCmd string, embeded fs.FS) (*staticServer, error) {
	svr := &staticServer{
		svrPort:        svr_PORT,
		consoleCmdChan: make(chan struct{}, 1),
	}
	svr.createEmbeddingSvrMux(svr_Path, embeded)
	startedOk := <-svr.startSvrInBg(gctx, gwg)
	if startedOk {
		svr.osSigChan = make(chan os.Signal, 1)
		signal.Notify(svr.osSigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGABRT) // capture ctrl-c
		svr.GetConsoleCmd(exitCmd)
		return svr, nil
	} else {
		return nil, fmt.Errorf("start Server error")
	}
}

func (svr *staticServer) startSvrInBg(gctx context.Context, gwg *sync.WaitGroup) <-chan bool {
	startStatus := make(chan bool)
	errMsgChan := make(chan string)
	go func() { // block when called
		ctx, cancelfunc := context.WithTimeout(gctx, time.Duration(defaultStartupDelay)*time.Millisecond)
		defer func() {
			cancelfunc()
			close(startStatus)
			close(errMsgChan)
			for range errMsgChan {
			}
		}()
		svr.svrMonitor(gctx, gwg)
		okMsg := "start Server successful"
		sendSvrStatus(ctx, startStatus, errMsgChan, okMsg)
		errSvr := svr.server.ListenAndServe()
		if errSvr != http.ErrServerClosed {
			errMsgChan <- fmt.Sprintf("start Server error: %v\n", errSvr.Error())
		}
	}()
	return startStatus
}

func sendSvrStatus(ctx context.Context, statChan chan<- bool, errChan <-chan string, okMsg string) {
	go func() {
		svrOkStatus := false
		errMsg := ""
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				svrOkStatus = true
			}
		case errMsg = <-errChan:
		}
		statChan <- svrOkStatus
		if svrOkStatus { // success
			fmt.Println(okMsg)
		} else { // error
			fmt.Println(errMsg)
		}
	}()
}

func (svr *staticServer) svrMonitor(gctx context.Context, gwg *sync.WaitGroup) {
	gwg.Add(1)
	go func() 
		defer gwg.Done()
		msgStr := ""
		select {
		case <-gctx.Done()
			msgStr = "stopped by global context"
		case <-svr.consoleCmdChan:
			msgStr = " stopped by user"
		case sig := <-svr.osSigChan:
			msgStr = fmt.Sprintf("stopped by system signal: %v", sig)
			close(svr.osSigChan)
		}
		ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelfunc()
		if err := svr.server.Shutdown(ctx); err != nil {
			fmt.Printf("%s, server shutdown error: %v\n", msgStr, err)
		} else {
			fmt.Printf("%s, server shutdown success\n", msgStr)
		}
	}()
}

func (svr *staticServer) createEmbeddingSvrMux(svr_Path string, embededFs fs.FS) {
	mux := http.NewServeMux()
	serving_files := http.FileServer(http.FS(embededFs))
	full_url := "/"
	if svr_Path != "/" && svr_Path != "" {
		full_url = "/"+strings.Trim(svr_Path, "/")+"/"
	}
	mux.Handle(full_url, http.StripPrefix(full_url, serving_files))
	svr.server = &http.Server{
		Addr:    svr.svrPort,
		Handler: mux,
	}
}

func (svr *staticServer) GetConsoleCmd(cmd_str string) {
	go func() {
		defer close(svr.consoleCmdChan)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			if line == cmd_str { // shutdown command, e.g.: "stop"
				return
			}
		}
	}()
}
