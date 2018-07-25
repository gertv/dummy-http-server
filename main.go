package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/satori/go.uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.RequestURI)
	fmt.Println()

	fmt.Println("Headers:")
	for key, value := range r.Header {
		fmt.Printf(" %s -> %s\n", key, value)
	}
	fmt.Println()

	fmt.Println("Body:")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()

	var cookie http.Cookie
	cookie.Name = "my-cookie"
	cookie.Value = uuid.Must(uuid.NewV4()).String()
	http.SetCookie(w, &cookie)

	fmt.Println()
	fmt.Printf(" < response containing cookie %s=%s >", cookie.Name, cookie.Value)
	fmt.Println()
	fmt.Println("-----")
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dummy-http-server <port>")
		panic(" no port specified")
	}

	port := os.Args[1]

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
