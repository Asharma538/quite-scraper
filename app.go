package main

import (
	"fmt"
	"time"
)

func main(){
	initFirebase();
	Ig := Instagram{};
	Ig.last_activity = make(map[string]int);
	Ig.addUser("living__motivation");
	
	for {
		for _,user := range Ig.users_to_monitor {
			if Ig.checkAndUpdateActivity(user,Ig.last_activity[user]){
				fmt.Println("New activity detected for user "+user);
			}
			fmt.Println(user,Ig.last_activity[user]);
		}
		time.Sleep(5*time.Minute);
	}
}