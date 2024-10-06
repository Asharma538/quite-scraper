package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var users_to_monitor[] string;
var last_activity (map[string]int);

func getPosts(profile_info string) int {
	var info_numbers[] int;
	for i:=0 ; i<len(profile_info) ; i++ {
		num_string := ""
		for profile_info[i]>='0' && profile_info[i]<='9' && i<len(profile_info) {
			num_string += string(profile_info[i]);
			i++;
		}
		if num_string!="" {
			num,err := strconv.Atoi(num_string);
			if err!=nil {
				return -1;
			}
			info_numbers = append(info_numbers, num);
		}
	}
	return info_numbers[2];
}

func getActivity(usr string) int {
	res,err := http.Get("https://www.instagram.com/"+usr+"/?hl=en");
	if err != nil {
		fmt.Printf("Error %s\n",err.Error());
		return -1;
	}
	defer res.Body.Close();

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error %v\n",res.StatusCode);
		return -1
	}
	data, err := io.ReadAll(res.Body);
	if err != nil {
		fmt.Println("Error reading body "+err.Error());
		return -1;
	}


	content := string(data);
	for i:=0; i<len(content); i++ {
		if (content[i]=='<'){
			if (content[i:i+31]=="<meta property=\"og:description\"") {
				j:=i;
				for content[j:j+5]!="Posts" && j<i+100 {
					j++;
				}
				return getPosts(content[i+41:j+5]);
			}
		}
	}
	return -1;
}

func checkAndUpdateActivity(usr string,last_act int) bool {
	// getting the lastest activity of the user
	latest_activity := getActivity(usr);
	
	// if unable to load the user or there's no activity
	if latest_activity == -1 || latest_activity==last_act{
		return false;
	}

	// if there's a new activity
	last_activity[usr] = latest_activity;
	return true;
}

func main(){
	last_activity = make(map[string]int);

	users_to_monitor = append(users_to_monitor, "zuck");
	users_to_monitor = append(users_to_monitor, "thisisbillgates");
	users_to_monitor = append(users_to_monitor, "living__motivation");

	for {
		for _,user := range users_to_monitor {
			if checkAndUpdateActivity(user,last_activity[user]){
				fmt.Println("New activity detected for user "+user);
			}
			fmt.Println(user,last_activity[user]);
		}
		time.Sleep(5*time.Minute);
	}
}