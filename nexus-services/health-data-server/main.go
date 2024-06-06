package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const defaultDataDir = "./data"

// Item represents the structure of the JSON data.
type Item struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func main() {
	// Load environment variables
	port := getEnv("PORT", "8080")
	dataDir := getEnv("DATA_DIR", defaultDataDir)

	// Ensure data directory exists
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// HTTP handlers
	http.HandleFunc("/create", createHandler(dataDir))
	http.HandleFunc("/read", readHandler(dataDir))
	http.HandleFunc("/update", updateHandler(dataDir))
	http.HandleFunc("/delete", deleteHandler(dataDir))

	// Start server
	log.Printf("Server started at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// createHandler handles the creation of a new item.
func createHandler(dataDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Printf("Invalid input: %v", err)
			return
		}
		item.ID = generateID()
		filePath := filepath.Join(dataDir, item.ID+".json")

		file, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			log.Printf("Failed to marshal JSON: %v", err)
			return
		}

		if err := os.WriteFile(filePath, file, 0644); err != nil {
			http.Error(w, "Failed to write file", http.StatusInternalServerError)
			log.Printf("Failed to write file: %v", err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
		log.Printf("Item created: %v", item)
	}
}

// readHandler handles reading an item by ID.
func readHandler(dataDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ids, ok := r.URL.Query()["id"]
		if !ok || len(ids[0]) < 1 {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			log.Println("Missing id parameter")
			return
		}

		id := ids[0]
		filePath := filepath.Join(dataDir, id+".json")
		file, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			log.Printf("File not found: %v", err)
			return
		}

		var item Item
		if err := json.Unmarshal(file, &item); err != nil {
			http.Error(w, "Failed to unmarshal JSON", http.StatusInternalServerError)
			log.Printf("Failed to unmarshal JSON: %v", err)
			return
		}

		json.NewEncoder(w).Encode(item)
		log.Printf("Item read: %v", item)
	}
}

// updateHandler handles updating an existing item by ID.
func updateHandler(dataDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			log.Printf("Invalid input: %v", err)
			return
		}

		filePath := filepath.Join(dataDir, item.ID+".json")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			log.Printf("File not found: %v", err)
			return
		}

		file, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			log.Printf("Failed to marshal JSON: %v", err)
			return
		}

		if err := os.WriteFile(filePath, file, 0644); err != nil {
			http.Error(w, "Failed to write file", http.StatusInternalServerError)
			log.Printf("Failed to write file: %v", err)
			return
		}

		json.NewEncoder(w).Encode(item)
		log.Printf("Item updated: %v", item)
	}
}

// deleteHandler handles deleting an item by ID.
func deleteHandler(dataDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ids, ok := r.URL.Query()["id"]
		if !ok || len(ids[0]) < 1 {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			log.Println("Missing id parameter")
			return
		}

		id := ids[0]
		filePath := filepath.Join(dataDir, id+".json")
		if err := os.Remove(filePath); err != nil {
			http.Error(w, "Failed to delete file", http.StatusInternalServerError)
			log.Printf("Failed to delete file: %v", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		log.Printf("Item deleted: %s", id)
	}
}

// generateID generates a new UUID.
func generateID() string {
	return uuid.New().String()
}

// getEnv retrieves environment variables or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
