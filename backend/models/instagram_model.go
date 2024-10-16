package models

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Instagram struct {
	Users_to_monitor []string
	Last_activity    map[string]int
}

// For Adding a new user to monitor
func (Ig *Instagram) AddUser(username string) {
	Ig.Users_to_monitor = append(Ig.Users_to_monitor, username)
}

// For getting the number of posts done by user
func (Ig *Instagram) getPosts(profile_info_raw string) int {
	profile_info := ""

	profile_info += string(profile_info_raw[0])
	for i := 1; i < len(profile_info_raw)-1; i++ {
		if (profile_info_raw[i-1] >= '0' && profile_info_raw[i-1] <= '9') && (profile_info_raw[i+1] >= '0' && profile_info_raw[i+1] <= '9') && profile_info_raw[i] == ',' {
			continue
		}
		profile_info += string(profile_info_raw[i])
	}
	profile_info += string(profile_info_raw[len(profile_info_raw)-1])

	var info_numbers []int
	for i := 0; i < len(profile_info); i++ {
		num_string := ""
		for profile_info[i] >= '0' && profile_info[i] <= '9' && i < len(profile_info) {
			num_string += string(profile_info[i])
			i++
		}
		if num_string != "" {
			num, err := strconv.Atoi(num_string)
			if err != nil {
				return -1
			}
			info_numbers = append(info_numbers, num)
		}
	}

	return info_numbers[2]
}

// For getting the activity of the user, i.e. followers, following, posts
func (Ig *Instagram) getActivity(usr string) int {
	res, err := http.Get("https://www.instagram.com/" + usr + "/?hl=en")
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())
		return -1
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error %v\n", res.StatusCode)
		return -1
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body " + err.Error())
		return -1
	}

	content := string(data)
	for i := 0; i < len(content); i++ {
		if content[i] == '<' {
			if content[i:i+31] == "<meta property=\"og:description\"" {
				j := i
				for content[j:j+5] != "Posts" && j < i+100 {
					j++
				}
				return Ig.getPosts(content[i+41 : j+5])
			}
		}
	}
	return -1
}

// For checking and updating the posts of the user
func (Ig *Instagram) CheckAndUpdateActivity(usr string, last_act int) bool {
	// getting the lastest activity of the user
	latest_activity := Ig.getActivity(usr)

	// if unable to load the user or there's no activity
	if latest_activity == -1 || latest_activity == last_act {
		return false
	}

	// if there's a new activity
	Ig.Last_activity[usr] = latest_activity
	return true
}
