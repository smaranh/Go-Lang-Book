package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		errHandler(err, "fetch: %v\n")
		b, _ := ioutil.ReadAll(resp.Body)
		errHandler(err, "fetch: reading: "+url+" failed error: %v\n")
		resp.Body.Close()
		fmt.Printf("Body: %s", b)
	}
}

func errHandler(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg, err.Error())
		os.Exit(1)
	}
}
