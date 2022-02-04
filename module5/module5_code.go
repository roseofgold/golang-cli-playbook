package module5

import (
	"fmt"
	"io"
	"net/http"
)

// GetExampleDotCom uses the "net/http" package to send a GET request to example.com
func GetExampleDotCom() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("something went wrong")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}
