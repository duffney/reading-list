```go
func home(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:4000/v1/books")
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bodyBytes)
}

```