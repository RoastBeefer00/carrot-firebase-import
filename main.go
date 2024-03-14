package main

import (
	"context"
	"encoding/json"
	// "fmt"
	"io"
	"log"
	"os"

	firebase "firebase.google.com/go"
	// "google.golang.org/api/option"
)

type Recipe []struct {
	Name       string   `json:"name"`
	Time       string   `json:"time"`
	Ingredents []string `json:"ingredients"`
	Steps      []string `json:"steps"`
}

func main() {
	//Open json file with recipes
	file, err := os.Open("db.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var recipes Recipe

	err = json.Unmarshal(content, &recipes)
	if err != nil {
		log.Fatal(err)
	}

	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "r-j-magenta-carrot-42069"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	for _, recipe := range recipes {
		_, _, err := client.Collection("recipes").Add(ctx, recipe)
		if err != nil {
			log.Fatalf("Failed adding alovelace: %v", err)
		}
	}
}
