package main

import (
	"io"
	"net/http"
)

// /hello controller
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello, World</title>
			</head>
			<body>
				Hello, World
			<body>
		</html>`,
	)
}

// error handling?
func main() {
	// register controller
	http.HandleFunc("/hello", hello)

	// serve static files using FileServer
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)

	// bind, listen, accept
	http.ListenAndServe("0.0.0.0:9000", nil)
}
