package main

import (
	"flag"
	"fmt"
	"os"
	"p2p/peer"
)

func main() {
	listen := flag.String("listen", ":3000", "address to listen on")
	connect := flag.String("connect", "", "remote peer address to connect to")
	flag.Parse()

	p := peer.New(*listen)

	// Start listener in background
	go func() {
		if err := p.Start(); err != nil {
			fmt.Println("server error:", err)
			os.Exit(1)
		}
	}()

	// If a remote address is given, connect and chat
	if *connect != "" {
		if err := p.ConnectAndChat(*connect); err != nil {
			fmt.Println("connect error:", err)
			os.Exit(1)
		}
	} else {
		// No remote — just wait for incoming connections
		fmt.Println("[*] Waiting for peers... (use -connect to dial someone)")
		select {} // block forever
	}
}