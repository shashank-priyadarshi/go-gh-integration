package gh

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func bearerAuthAPICall(reqURL, authToken string) []byte {
	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	request, err := http.NewRequest("GET", reqURL, bytes.NewBuffer([]byte("")))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authToken))
	if err != nil {
		log.Fatalln("err in creating new request: ", err)
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
	}

	if err != nil {
		log.Fatalln("err in making bearerAuth req: ", err)
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("err in reading req response: ", err)
	}

	return respBody
}
