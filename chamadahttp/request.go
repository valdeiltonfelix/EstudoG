package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://www.google.com/search?q=klbn4&oq=&gs_lcrp=EgZjaHJvbWUqCQgEEEUYOxjCAzIJCAAQRRg7GMIDMgkIARBFGDsYwgMyCQgCEEUYOxjCAzIJCAMQRRg7GMIDMgkIBBBFGDsYwgMyCQgFEEUYOxjCAzIJCAYQRRg7GMIDMgkIBxBFGDsYwgPSAQs3ODIxNTEwajBqN6gCCLACAQ&sourceid=chrome&ie=UTF-8")

	if err != nil {
		panic(err)
	}

	visualiza, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(visualiza))
	request.Body.Close()

}
