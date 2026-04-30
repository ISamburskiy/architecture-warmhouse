package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"strings"
)

// TemperatureResponse структура ответа API
type TemperatureResponse struct {
	Value       float64   `json:"value"`
	//Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	//Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	//SensorType  string    `json:"sensor_type"`
	//Description string    `json:"description"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	// Если location не указан, определяем по sensorID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
	case "2":
			location = "Bedroom"
	case "3":
			location = "Kitchen"
	default:
		location = "Unknown"
	}
	}

	// Если sensorID не указан, определяем по location
	if sensorID == "" {
		switch location {
		case "Living Room":
		sensorID = "1"
	case "Bedroom":
		sensorID = "2"
	case "Kitchen":
	sensorID = "3"
	default:
	sensorID = "0"
	}
	}

	// Генерируем случайную температуру от -5 до 35 °C
	temperature := rand.Float64()*40 - 5

	response := TemperatureResponse{
		Location:    location,
		SensorID:   sensorID,
		Value: temperature,
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func temperatureByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем sensorID из пути (например, /temperature/1)
	parts := strings.Split(r.URL.Path, "/")
	sensorID := parts[len(parts)-1]

	// Переиспользуем логику из temperatureHandler
	location := ""
	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}

	temperature := rand.Float64()*40 - 5
	response := TemperatureResponse{
		Location: location,
		SensorID: sensorID,
		Value:    temperature,
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/temperature", temperatureHandler)
	http.HandleFunc("/temperature/", temperatureByIDHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("Temperature API starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
