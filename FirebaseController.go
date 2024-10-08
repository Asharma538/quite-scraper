package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

type FirebaseController struct {
	authClient      *auth.Client
	firestoreClient *firestore.Client
	collection      string
}

func (controller *FirebaseController) init() {
	opt := option.WithCredentialsFile("./secret_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		fmt.Printf("Error initializing app: %v\n", err)
	}

	controller.authClient, err = app.Auth(context.Background())
	if err != nil {
		fmt.Printf("Error getting Auth client: %v\n", err)
	}

	controller.firestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		fmt.Printf("Error initializing firestore client %v\n", err)
	}
}

func (controller *FirebaseController) getUsersToMonitor(platform string) []string {
	// Requesting the document of the Platform
	snapshot, err := controller.firestoreClient.Collection(controller.collection).Doc(platform).Get(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	// Getting the usernames array from the document
	usernamesInterface := snapshot.Data()["usernames"].([]interface{})
	usernames := make([]string, len(usernamesInterface))

	for idx, user := range usernamesInterface {
		usernames[idx] = user.(string)
	}

	return usernames
}