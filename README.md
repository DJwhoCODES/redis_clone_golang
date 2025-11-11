# GoRedis — A Minimal Redis Clone in Go

## 🚀 Overview

GoRedis is a simplified Redis-like in-memory key-value database built entirely in **Go**.  
It demonstrates core concepts like **network programming**, **concurrency**, **protocol design**, and **data synchronization** — the same building blocks real-world systems like Redis use.

---

## 🧱 Features

- TCP server built using Go's `net` package.
- RESP-like protocol for client-server communication.
- Thread-safe in-memory key-value store.
- Supports basic Redis commands:
  - `PING` → `PONG`
  - `SET key value`
  - `GET key`
  - `DEL key`
- Handles multiple concurrent client connections.
- Clean, production-grade folder structure.

---

## 🗂️ Folder Structure

```
goredis/
│
├── cmd/
│   └── goredis/
│       └── main.go         # Entry point of the application
│
├── internal/
│   ├── server/
│   │   └── server.go       # Core server logic (TCP, connection handling)
│   ├── store/
│   │   └── store.go        # In-memory key-value database
│   ├── protocol/
│   │   └── protocol.go     # RESP-like protocol parsing and formatting
│   └── peer/
│       └── peer.go         # Manages client connection lifecycle
│
└── go.mod
```

---

## ⚙️ Installation

### 1. Clone the repository

```bash
git clone https://github.com/djwhocodes/goredis.git
cd goredis
```

### 2. Run the server

```bash
go run ./cmd/goredis
```

By default, the server listens on port **5000**.

---

## 💬 Usage

### Option 1: Using Redis CLI

If you already have `redis-cli` installed:

```bash
redis-cli -p 5000
```

Now try:

```bash
PING
SET name devanshu
GET name
DEL name
```

### Option 2: Using Telnet

If you don’t have `redis-cli`, use Telnet:

```bash
telnet localhost 5000
```

Then type:

```
PING
SET city indore
GET city
```

---

## 🧠 Concepts You’ll Learn

### Networking

- How TCP servers and clients communicate over sockets.
- How to handle multiple clients using goroutines.

### Concurrency

- How to use channels, mutexes, and goroutines for thread-safe operations.

### Protocols

- How Redis uses a serialization protocol (RESP) for structured command exchange.

### Data Structures

- Building an in-memory key-value store using Go maps.

---

## 👨‍💻 Author

**Devanshu Jain - DJwhoCodes**  
Built while learning Go systems programming and Redis internals.

---

## 🧾 License

MIT License © 2025 Devanshu Jain
