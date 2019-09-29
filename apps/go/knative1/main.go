package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")

	defer r.Body.Close()

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Print("Request body", string(reqBody))

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", os.Getenv("KNATIVE_GATEWAY"), nil)

	req.Host = os.Getenv("NEXT_APP")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Error Occurred  %s!\n",err.Error())
	}else{
		bodyStr, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Fprintf(w, "Error Occurred  %s!\n",err.Error())
		}else{
			fmt.Fprintf(w, "%s!\n", bodyStr)
		}
	}
}

func randoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s!\n", "No so random...")
}


func main() {
	log.Print("Hello world sample started.")

	http.HandleFunc("/", Handler)
	http.HandleFunc("/random", randoHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
