package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

func pluginManifest(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./.well-known/ai-plugin.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteValue)
}

func openApiSpec(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("openapi.yaml")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	w.Header().Set("Content-Type", "text/yaml")
	w.Write(byteValue)
}

func pluginLogo(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("logo.png")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	w.Header().Set("Content-Type", "image/png")
	w.Write(byteValue)
}

type requestBody struct {
	City string `json:"city"`
}

func getCityWttr(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	var rb requestBody
	err = json.Unmarshal(body, &rb)
	if err != nil {
		http.Error(w, "Error unmarshalling json", http.StatusInternalServerError)
	}

	// send http call to api to get weather data
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + strings.ToLower(rb.City) + "&appid=TOKEN"
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error getting weather data", http.StatusInternalServerError)
	}

	// convert resp to json
	var cityweather interface{}
	err = json.NewDecoder(resp.Body).Decode(&cityweather)
	if err != nil {
		http.Error(w, "Error decoding json", http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(cityweather)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	fmt.Println("Server running on port 5004")

	mux := http.NewServeMux()

	mux.HandleFunc("/openapi.yaml", openApiSpec)
	mux.HandleFunc("/.well-known/ai-plugin.json", pluginManifest)
	mux.HandleFunc("/logo.png", pluginLogo)
	mux.HandleFunc("/wttr", getCityWttr)

	handler := cors.New(cors.Options{
		AllowedOrigins:      []string{"https://chat.openai.com"},
		AllowCredentials:    true,
		AllowedHeaders:      []string{"*"},
		AllowedMethods:      []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowPrivateNetwork: true,
		Debug:               true,
	}).Handler(mux)

	http.ListenAndServe(":5004", handler)
}
