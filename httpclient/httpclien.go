package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	//c := http.Client{Timeout: time.Microsecond}
	c := http.Client{}
	jsonvar := bytes.NewBuffer([]byte(`{"nome":"valdeiton"}`))
	resp, err := c.Post("http://www.google.com.br", "application/json", jsonvar)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
	// body, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// println(string(body))

}
