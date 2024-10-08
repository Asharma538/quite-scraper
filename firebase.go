package main

import (
	"context"
	"fmt"
	
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	
	"google.golang.org/api/option"
)

func initFirebase() *auth.Client {
	opt := option.WithCredentialsFile("./secret_key.json");
	app, err := firebase.NewApp(context.Background(), nil, opt);
	
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err);
	}
	
	client, err := app.Auth(context.Background());
	if err != nil {
		fmt.Printf("error getting Auth client: %v\n", err);
	}
	return client;
}