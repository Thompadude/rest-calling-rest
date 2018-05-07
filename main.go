package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func getJson(w http.ResponseWriter, r *http.Request) {
	getUrl := r.URL.Query().Get("getUrl")
	validateGetUrl(&getUrl)

	defer r.Body.Close()

	resp := getResponse(&getUrl)
	bytes := readResponse(resp)
	data := unmarshalResponse(&bytes)

	writeHeaders(w)

	json.NewEncoder(w).Encode(data)
}

func writeHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTeapot)
}

func getResponse(getUrl *string) (resp *http.Response) {
	resp, err := http.Get(*getUrl)
	if err != nil {
		panic("ERROR SENDING GET REQUEST")
	}
	return
}

func readResponse(resp *http.Response) (bytes []byte) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("ERROR READING RESPONSE")
	}
	return
}

func unmarshalResponse(bytes *[]byte) (data map[string]interface{}) {
	if err := json.Unmarshal(*bytes, &data); err != nil {
		panic("ERROR UNMARSHMALLOW")
	}
	return data
}

func validateGetUrl(getUrl *string) {
	if len(*getUrl) < 1 {
		*getUrl = "https://jsonplaceholder.typicode.com/posts/1"
	}
}

func main() {
	http.HandleFunc("/get-json/", getJson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
