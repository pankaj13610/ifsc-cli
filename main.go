package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	codes := os.Args[1:]
	if len(codes) != 1 {
		fmt.Fprintf(os.Stderr, "Please give exactly 1 ifsc code as argument \n")
		os.Exit(1)
	}
	getData(codes[0])
}

func getData(ifsc string) {
	requestURL := fmt.Sprintf("https://ifsc.razorpay.com/%s", ifsc)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create request: %s \n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making http request %s \n", err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Invalid IFSC Code")
		os.Exit(1)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var j map[string]interface{}
	err = json.Unmarshal(resBody, &j)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to decode json %s\n", err)
		os.Exit(1)
	}

	fmt.Println(j["BRANCH"], j["CITY"])
}
