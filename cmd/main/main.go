package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// Alias package name to be used with different name on import
	restPackage "github.com/ik5/go-into/rest"
	"github.com/ik5/go-into/signals"
)

func handleSignals(quit chan<- bool) {
	quitSigs := make(chan os.Signal, 1)
	hupSig := make(chan os.Signal, 1)
	infoSig := make(chan os.Signal, 1)

	defer close(quitSigs)
	defer close(hupSig)
	defer close(infoSig)

	signal.Notify(hupSig, syscall.SIGHUP)
	signal.Notify(infoSig, signals.SIGINFO)
	signal.Notify(quitSigs, syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)

	for {
		select {
		case <-hupSig:
			// TODO: rotate logs
			fmt.Println("Going to rotate logs... ")
		case <-infoSig:
			// TODO: print debug info...
			fmt.Println("Debug information: ")
		case sig := <-quitSigs:
			if sig == syscall.SIGINT {
				fmt.Println("You shell not pass!")
				continue
			}
			fmt.Printf("Going to terminate application due to signal %d\n", sig)
			quit <- true
		}
	}
}

func initialize() {
	// TODO: Add settings, initialize of logging systems etc...
}

func main() {
	initialize()

	rest := restPackage.InitREST("", uint16(3000))
	rest.RegisterUserRoute("/", "GET", indexPage)
	rest.SetUserRouting()
	defer rest.Stop()

	quit := make(chan bool, 1)
	defer close(quit)

	go handleSignals(quit)

	go rest.Serve()

	<-quit
}
