package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/flatcar-linux/container-linux-config-transpiler/config"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write([]byte("hello world"))
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		cfg, ast, report := config.Parse(body)
		if len(report.Entries) > 0 {
			w.Write([]byte(report.String()))
			return
		}
		ignCfg, report := config.Convert(cfg, "", ast)
		if len(report.Entries) > 0 {
			w.Write([]byte(report.String()))
			if report.IsFatal() {
				return
			}
		}
		dataOut, err := json.MarshalIndent(&ignCfg, "", "  ")
		dataOut = append(dataOut, '\n')
		if err != nil {
			log.Fatalf("Failed to marshal output: %v", err)
		}
		w.Write([]byte(dataOut))
	}
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/clctranspilerfunction", helloHandler)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
