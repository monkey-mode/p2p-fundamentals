package peer

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Peer struct {
	ListenAddr string
}

func New(listenAddr string) *Peer {
	return &Peer{ListenAddr: listenAddr}
}

// Start begins listening and blocks
func (p *Peer) Start() error {
	ln, err := net.Listen("tcp", p.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %w", err)
	}
	fmt.Printf("[*] Listening on %s\n", p.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go p.handleIncoming(conn)
	}
}

// handleIncoming reads messages from a connected peer
func (p *Peer) handleIncoming(conn net.Conn) {
	defer conn.Close()
	remote := conn.RemoteAddr().String()
	fmt.Printf("[+] Peer connected: %s\n", remote)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("[%s] %s\n", remote, msg)
	}

	fmt.Printf("[-] Peer disconnected: %s\n", remote)
}

// ConnectAndChat dials a remote peer and lets you type messages
func (p *Peer) ConnectAndChat(remoteAddr string) error {
	conn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		return fmt.Errorf("dial error: %w", err)
	}
	defer conn.Close()
	fmt.Printf("[*] Connected to %s — type messages below\n", remoteAddr)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fmt.Fprintln(conn, line) // \n is the message delimiter
	}
	return nil
}