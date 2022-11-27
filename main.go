package main

import (
	"net/http"
	"encoding/json"
	//"image/jpeg"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		data, err := getRandomDog()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := http.Get(data.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(b)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello bro!"))
}

func getRandomDog() (dogData, error) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		return dogData{}, err
	}

	defer resp.Body.Close()

	var d dogData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return dogData{}, err
	}

	return d, nil
}

type dogData struct {
	Url string `json:"message"`
	Status string `json:"status"`
}
