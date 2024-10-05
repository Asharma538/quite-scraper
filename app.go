package main

import (
	"fmt"
	"time"
	"github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
)

func main(){
	url := launcher.NewUserMode().MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage("https://www.instagram.com/buffettbillions/")

	page.MustWaitStable()
	page.MustWaitLoad()

	time.Sleep(2 * time.Second);

	// content, err := page.HTML()
    // if err != nil {
    //     fmt.Printf("Error getting content: %v\n", err)
    //     return
    // }

	posts := page.MustElements("a[hrefs*='/p/]")
	for _,post := range posts {
		href,err := post.Attribute("href")
		if err == nil {
			fmt.Println("Found a post at: https://instagram.com"+*href);
		}
	}

}