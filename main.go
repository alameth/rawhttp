package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"rawhttp/rawhttp"
)

// rawhttp is a test client for HTTP and HTTPS. It's a YAFIYGI client (you asked
// for it, you got it), differing from curl and wget in that it does what it's
// told and stays out of the way.

// General Options
var optVerbose = flag.Bool("v", false, "Verbose output")
var optOutputFile = flag.String("o", "-", "Output file")
var optRequestFile = flag.String("m", "", "Read entire message from file")
var optHeaderFile = flag.String("h", "", "Read request header from file")
var optBody = flag.String("b", "", "Read message body from file")

// TLS Options
var optServerName = flag.String("S", "-", "SNI Name")
var optUseTLS = flag.Bool("T", false, "Force use of TLS")

func main() {
	// Check command-line options
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "Need exactly one argument\n")
		os.Exit(1)
	}

	// Get request from stdin
	request := rawhttp.NewRawHTTP()
	err := request.Read(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	conn, err := rawhttp.Dial(flag.Arg(0), *optUseTLS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	request.Write(conn)
	response := rawhttp.NewRawHeader()
	err = response.ReadHeader(bufio.NewScanner(conn))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if *optOutputFile == "-" {
		response.WriteHeader(os.Stdout)
	} else {
		w, err := os.Create(*optOutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr,"%v\n", err)
			os.Exit(2)
		}
		defer w.Close()
		response.WriteHeader(w)
	}
}
