package main

//url health checker based on the urls in a slice
import (
	"fmt"
	"net/http"
	"time"
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
	//creating a new channel
	c := make(chan string)

	for _, link := range links {
		// By default when you run a program, we create a single go Routine
		// checkLink(link)
		//child routines are not getting respect as that of the main routine
		go checkLink(link, c)

		//one core
		// creates each go routine for each iteration and run the function with it
		//Scheduler runs one routine until it finishes or makes a blocking call(like a http call), then it makes the other go routine.
		// By default, tries to use one core. When running on one CPU, they are not running truly at the same time.

		//If we change the confi to run in multiple cores, then each go routine will run in each core
	}

	//infinite loop

	for l := range c {
		// go checkLink(l, c)
		//anonymous function for sleep time - Function literals
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		//blocking code which when gets a data, exits the code - passing this channel value is a blocking operation
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + "might be down")
		c <- link
		return
	}
	c <- link
	fmt.Println(link + " is up!")
}
