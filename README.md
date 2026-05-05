# p2p-fundamentals

A hands-on Go exercise exploring P2P networking fundamentals: TCP sockets, bidirectional peers, and line-delimited messaging.

## What it does

Each process acts as both a server and a client simultaneously:

- **Listens** on a TCP port for incoming peer connections (goroutine in background)
- **Dials** a remote peer and enters an interactive chat session (foreground)

Messages are newline-delimited and sent over raw TCP. Incoming messages are printed to stdout with the sender's address as a prefix.

## Project structure

```
.
├── main.go        # CLI entrypoint — parses flags, wires up the peer
└── peer/
    └── peer.go    # Peer type: Start (listener), ConnectAndChat (dialer)
```

## Usage

### Run Peer A (listen only, no remote yet)

```bash
go run . -listen :3000
```

### Run Peer B (listen + connect to A)

```bash
go run . -listen :3001 -connect localhost:3000
```

Peer B connects to Peer A and enters interactive mode. Type a message and press Enter — it appears on Peer A's terminal prefixed with Peer B's address.

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-listen` | `:3000` | Local address to accept incoming connections on |
| `-connect` | *(empty)* | Remote peer address to dial; omit to wait for inbound connections |

## How it works

1. `peer.New(addr)` creates a `Peer` with the given listen address.
2. `p.Start()` opens a TCP listener and spawns a goroutine per accepted connection that reads lines and prints them.
3. `p.ConnectAndChat(addr)` dials the remote address, then reads stdin line-by-line, writing each line to the connection.
4. `\n` is the message delimiter — `fmt.Fprintln` writes it on send, `bufio.Scanner` splits on it on receive.

## Requirements

- Go 1.21+
