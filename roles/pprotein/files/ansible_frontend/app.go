package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type RequestData struct {
	Param string `json:"param"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Stdout  string `json:"stdout"`
	Stderr  string `json:"stderr"`
	Message string `json:"message,omitempty"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/isucon/ansible_frontend/static/index.html")
}

func runScriptHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if data.Param == "" {
		data.Param = "main" // デフォルト値
	}
	fmt.Printf("Received parameter: %s\n", data.Param)

	cmd := exec.Command("/home/isucon/ansible_frontend/script.sh", data.Param)
	cmdOutput, err := cmd.CombinedOutput()

	response := ResponseData{
		Status: "success",
		Stdout: string(cmdOutput),
	}

	if err != nil {
		response.Status = "error"
		response.Message = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/run-script", runScriptHandler)

	port := ":5000"
	fmt.Printf("Server running on http://0.0.0.0%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
