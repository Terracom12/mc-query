package main

import (
	"fmt"
	"log"
	"mc-query/mcproto/json"
	"mc-query/mcproto/packets"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultPort = 25565

const usageString = `USAGE:
	mc-query <server-ip> [<port>]

ARGS:
	server-ip   valid server ip address
	port        port number to use (default = %d)`

func main() {
	var port int = defaultPort

	if len(os.Args) == 1 {
		cliErr("Not enough arguments!")
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage()
	}

	if len(os.Args) > 2 {
		var err error
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			cliErr("Bad port: %s (%v)", os.Args[2], err)
		}
	}

	devNull, err := os.Open(os.DevNull)
	if err != nil {
		log.Panic("Failed to open /dev/null")
	}
	defer func() { devNull.Close() }()
	log.SetOutput(devNull)

	for _, s := range os.Environ() {
		splitKV := strings.Split(s, "=")

		if splitKV[0] == "LOG_LEVEL" {
			log.SetOutput(os.Stderr)
		}
	}

	log.Printf("Attempting to connect to IP %s:%d\n", os.Args[1], port)

	qualAddr := fmt.Sprintf("%s:%d", os.Args[1], port)
	conn, err := net.Dial("tcp", qualAddr)

	if err != nil {
		log.Fatal(err)
	}

	conn.SetDeadline(time.Time{})

	defer func() { conn.Close() }()

	// See Protocol_version_numbers page on minecraft.wiki
	const protocolVersion = 772 // For 1.21.8
	handshake := packets.MakeHandshake(protocolVersion, "", 0, packets.IntentStatus)
	err = handshake.Send(conn)
	log.Printf("Sending Handshake packet... (err=%v)\n", err)
	log.Println(handshake)

	statusReq := packets.StatusRequest{}
	err = statusReq.Send(conn)
	log.Printf("Sending StatusRequest packet... (err=%v)\n", err)

	log.Printf("Receiving a StatusResponse packet... (err=%v)\n", err)

	statusResp := packets.StatusResponse{}
	numRead, err := statusResp.Receive(conn)

	log.Printf("Response %d (err=%v): %s\n", numRead, err, statusResp.JsonResponse.Str)

	status, err := json.DeserializeStatus(statusResp.JsonResponse.Str)

	if err != nil {
		log.Fatal("Error parsing json: ", statusResp.JsonResponse.Str)
	}

	fmt.Println("Online players: ", status.Players.Online)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, usageString, defaultPort)
	fmt.Fprintf(os.Stderr, "\n")
}

func cliErr(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Fprintf(os.Stderr, "\n")
	printUsage()
	os.Exit(1)
}
