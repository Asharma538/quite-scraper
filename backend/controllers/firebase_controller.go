package controllers

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

type FirebaseController struct {
	AuthClient      *auth.Client
	FirestoreClient *firestore.Client
	Collection      string
}

func (controller *FirebaseController) Init() {
	opt := option.WithCredentialsFile("./config/secret_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	
	if err != nil {
		fmt.Printf("Error initializing app: %v\n", err)
	}

	controller.AuthClient, err = app.Auth(context.Background())
	if err != nil {
		fmt.Printf("Error getting Auth client: %v\n", err)
	}

	controller.FirestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		fmt.Printf("Error initializing firestore client %v\n", err)
	}
}

func (controller *FirebaseController) GetUsersToMonitor(platform string) []string {
	// Requesting the document of the Platform
	snapshot, err := controller.FirestoreClient.Collection(controller.Collection).Doc(platform).Get(context.Background())
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