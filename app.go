package main

import (
	"fmt"
	"time"
	"quite-scraper/controllers"
	"quite-scraper/models"
)

func main() {
	firebaseController := controllers.FirebaseController{};
	firebaseController.Collection = "Quite-Scraper"
	firebaseController.Init()
	defer firebaseController.FirestoreClient.Close()

	Ig := models.Instagram{}
	Ig.Last_activity = make(map[string]int)
	Ig.Users_to_monitor = firebaseController.GetUsersToMonitor("Instagram")

	for {
		for _, user := range Ig.Users_to_monitor {
			if Ig.CheckAndUpdateActivity(user, Ig.Last_activity[user]) {
				fmt.Println("New activity detected for user " + user)
			}
			fmt.Println(user, Ig.Last_activity[user])
		}
		fmt.Print("\n\n\n\n");
		time.Sleep(5 * time.Minute)
	}
}
