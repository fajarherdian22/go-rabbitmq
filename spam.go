package main

import (
	"fmt"
	"net/http"
)

func main() {
	for i := 0; i < 50; i++ {
		req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:3750/send?msg=%d", i), nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "text/plain")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
}
