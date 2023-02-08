package main

//url health checker based on the urls in a slice
import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://twitter.com",
		"http://amazon.com",
		"http://golang.org",
	}
	for _, link := range links {
		// By default when you run a program, we create a single go Routine
		// checkLink(link)
		//child routines are not getting respect as that of the main routine
		go checkLink(link)

		//one core
		// creates each go routine for each iteration and run the function with it
		//Scheduler runs one routine until it finishes or makes a blocking call(like a http call), then it makes the other go routine.
		// By default, tries to use one core. When running on one CPU, they are not running truly at the same time.

		//If we change the confi to run in multiple cores, then each go routine will run in each core
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + "might be down")
		return
	}

	fmt.Println(link + " is up!")
}
