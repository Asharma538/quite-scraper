package main

import (
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

	http.HandleFunc("/getActivity", func(w http.ResponseWriter, r *http.Request) {
		users := ""
		var waitGroup sync.WaitGroup
		var mu = &sync.Mutex{}

		for _, user := range Ig.Users_to_monitor {
			waitGroup.Add(1)
			go func(u string) {
				defer waitGroup.Done()
				if Ig.CheckAndUpdateActivity(u, Ig.Last_activity[u]) {
					mu.Lock()
					users += u + ":yes,"
					mu.Unlock()
				} else {
					mu.Lock()
					users += u + ":no,"
					mu.Unlock()
				}
			}(user)
		}

		waitGroup.Wait()
		fmt.Fprint(w, users)
	})

	http.ListenAndServe(":8080", nil)
}
