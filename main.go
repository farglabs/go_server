package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"os"
	"strconv"
	// Some useful packages for go newbies looking to for ideas:
	/*
		"bufio"
		"html"
		"regexp"
		"strings"
		"time"
	*/
)

func main() {
	config, ini_err := ini.Load("settings.ini")
	if ini_err != nil {
		fmt.Printf("Failed to read configuration file: %v", ini_err)
		os.Exit(1)
	}
	tls_cert := config.Section("server").Key("tls_cert").String()
	tls_key := config.Section("server").Key("tls_key").String()
	// TODO: Should there be an error check for these? ^^

	port, port_err := config.Section("server").Key("http_port").Int()

	if port_err != nil {
		fmt.Printf("Invalid value for 'http_port' in 'server'.\nError: ", port_err)
		fmt.Printf("\nNo port set. We're using 8086 as a default.\n")
		port = 8086
	} else {
		fmt.Printf("Server listening on port: " + strconv.Itoa(port) + "\n")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		//// Enable this to output the request URL to the console:
		//fmt.Printf(r.URL.String())
		
		if r.URL.String() == "/" {
	/* EXAMPLE SECTION 1: This is what's returned at the root of your web server */
			/* See note after this default response */
			response := "<!DOCTYPE html>\n<html>\n<head>\n<script src="https://unpkg.com/htmx.org@1.8.5" integrity="sha384-7aHh9lqPYGYZ7sTHvzP1t3BAfLhYSTy9ArHdP3Xsr9/3TlGurYgcPBoFmXX2TX/w" crossorigin="anonymous"></script>\n</head>\n<body>\nHey there! You're running in the following mode:\n<br>\n"
			/* Note that a specific version of HTMX is linked here; this can be copied to other responses. For documentation, see: https://htmx.org/ */
			response += config.Section("").Key("app_mode").String()
			response += "\n<hr>\n<bold>Feel free to change that in your settings.ini</bold>\n</body>\n</html>"
			w.Write([]byte(response))
	/* End of SECTION 1 (Make sure to set a response variable and then w.Write([]byte(response)) ) */
	/* EXAMPLE SECTION 2+: This can be duplicated as many times as needed (or even removed) */
		} else if r.URL.String() == "/your-url-example/here" {
			response := "your html-formatted response would go here"
			w.Write([]byte(response))
	/* End of SECTION 2+ (Make sure to change "/your-url-example/here" to a valid web path string) */
		} else {
			http.NotFound(w, r)
		}

	})

	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(port), tls_cert, tls_key, nil))

}
