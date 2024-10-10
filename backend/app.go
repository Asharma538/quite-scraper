package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quite-scraper/controllers"
	"quite-scraper/models"
	"sync"
)

func main() {
	firebaseController := controllers.FirebaseController{}
	firebaseController.Collection = "Quite-Scraper"
	firebaseController.Init()
	defer firebaseController.FirestoreClient.Close()

	Ig := models.Instagram{}
	Ig.Last_activity = make(map[string]int)
	Ig.Users_to_monitor = firebaseController.GetUsersToMonitor("Instagram")

	fmt.Println("Users to monitor: ", Ig.Users_to_monitor)

	http.HandleFunc("/getactivity", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var waitGroup sync.WaitGroup
		for _, user := range Ig.Users_to_monitor {
			waitGroup.Add(1)
			go func(u string) {
				defer waitGroup.Done()
				Ig.CheckAndUpdateActivity(u, Ig.Last_activity[u])
			}(user)
		}
		waitGroup.Wait()

		responseJson, err := json.Marshal(Ig.Last_activity)
		if err != nil {
			fmt.Print("Error:" + err.Error())
			return
		}
		w.Write(responseJson)
	})

	http.ListenAndServe(":8080", nil)
}
