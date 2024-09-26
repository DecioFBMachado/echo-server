package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const defaultPort = 8080

var statusCode *int
var responseBody string

func main() {
	port := flag.Int("port", defaultPort, "Define port to be used")
	cert := flag.String("c", "", "Path to certificate")
	keyFile := flag.String("key", "", "")
	statusCode = flag.Int("status", http.StatusOK, " Status Code Response")
	responseFile := flag.String("rf", "", "Path to response body file")
	help := flag.Bool("h", false, "Show available options")
	flag.Parse()

	if *help {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if flag.Lookup("rf").Value.String() != "" {
		file, err := os.ReadFile(*responseFile)
		if err != nil {
			log.Fatalf("Failed to open response file. %s", err)
		}
		if len(file) > 0 {
			responseBody = string(file)
		}
	}

	http.HandleFunc("/", handleRequest)
	if len(*cert) > 0 && len(*keyFile) > 0 {
		err := http.ListenAndServeTLS(fmt.Sprintf(":%d", *port), *cert, *keyFile, nil)
		if err != nil {
			log.Fatalf("Unable to bind to port. %s", err.Error())
		}
	} else {
		err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
		if err != nil {
			log.Fatalf("Unable to bind to port. %s", err.Error())
		}
	}

	log.Printf("Listening on 0.0.0.0:%d \n", *port)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.RequestURI, req.Method)

	fmt.Println("<<<<<<<<<<<<<<< User Agent >>>>>>>>>>>>>>>")
	fmt.Printf("%s\n", req.UserAgent())

	fmt.Println("<<<<<<<<<<<<<<< Headers >>>>>>>>>>>>>>>")
	parseHeaders(req)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Print("Unbale to parse Body")
	}

	if len(body) > 0 {
		fmt.Println("<<<<<<<<<<<<<<< Body >>>>>>>>>>>>>>>")
		responseBody = string(body)
		fmt.Println(string(body))
	}

	w.WriteHeader(*statusCode)

	fmt.Fprint(w, responseBody)
}

func parseHeaders(req *http.Request) {
	for name, headers := range req.Header {
		for header := range headers {
			fmt.Printf("\t%v: %v\n", name, header)
			headers = append(headers, fmt.Sprintf("\t%v: %v\n", name, header))
		}
	}
}
