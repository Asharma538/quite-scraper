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
	fmt.Println("Connecting to firebase...")
	firebaseController := controllers.FirebaseController{}
	firebaseController.Collection = "Quite-Scraper"
	firebaseController.Init()
	defer firebaseController.FirestoreClient.Close()
	fmt.Println("Connected!")

	fmt.Println("Starting Instagram monitoring...")
	Ig := models.Instagram{}
	Ig.Last_activity = make(map[string]int)
	Ig.Users_to_monitor = firebaseController.GetUsersToMonitor("Instagram")
	fmt.Println("Users to monitor:",Ig.Users_to_monitor);

	var waitGroup sync.WaitGroup
	for _, user := range Ig.Users_to_monitor {
		waitGroup.Add(1)
		go func(u string) {
			defer waitGroup.Done()
			Ig.CheckAndUpdateActivity(u, Ig.Last_activity[u])
		}(user)
	}
	waitGroup.Wait()

	fmt.Println("Starting server...")
	http.HandleFunc("/getactivity", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var users_with_updates []string;
		var mutex sync.Mutex

		for _, user := range Ig.Users_to_monitor {
			waitGroup.Add(1)
			go func(u string) {
				defer waitGroup.Done()
				if Ig.CheckAndUpdateActivity(u, Ig.Last_activity[u]) {
					mutex.Lock()
					users_with_updates = append(users_with_updates, u);
					mutex.Unlock()
				}
			}(user)
		}
		waitGroup.Wait()
		
		responseJson, err := json.Marshal(map[string][]string{"users_with_updates": users_with_updates})
		if err != nil {
			fmt.Print("Error:" + err.Error())
			return
		}
		w.Write(responseJson)
	})

	http.ListenAndServe(":8080", nil)
}
