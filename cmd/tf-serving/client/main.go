package main

import (
	tfclient "github.com/drankou/deep-odds/pkg/tf-serving/client"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading .env file. #%v", err)
	}

	client, err := tfclient.NewPredictionClient(os.Getenv("TF_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	mockInput := []float32{84, 2, 2, 80, 84, 77, 44, 10, 2, 3, 6, 8, 1, 1, 2, 0, 0, 0, 0, 6.5, 1.3, 5.5, 3.25, 4.2, 1.9}
	prediction, err := client.Predict(os.Getenv("MODEL_NAME"), mockInput)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Prediction: %+v", prediction)
}
