package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/demo", demoHandler)

	log.Println("server is starting...")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal("server error", err)
		return
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("demo handler")
	log.Println("%+v", r)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{
		"message": "hello world",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//data, err := json.Marshal(response)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//code, err := w.Write(data)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//log.Println(code)

	json.NewEncoder(w).Encode(response)
}
