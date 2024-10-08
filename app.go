package main

import (
	"fmt"
	"time"
)

func main() {
	firebaseController := FirebaseController{}
	firebaseController.collection = "Quite-Scraper"
	firebaseController.init()
	defer firebaseController.firestoreClient.Close()

	Ig := Instagram{}
	Ig.last_activity = make(map[string]int)
	Ig.users_to_monitor = firebaseController.getUsersToMonitor("Instagram")

	for {
		for _, user := range Ig.users_to_monitor {
			if Ig.checkAndUpdateActivity(user, Ig.last_activity[user]) {
				fmt.Println("New activity detected for user " + user)
			}
			fmt.Println(user, Ig.last_activity[user])
		}
		fmt.Print("\n\n\n\n");
		time.Sleep(5 * time.Minute)
	}
}
