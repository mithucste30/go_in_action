package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)
func Run(searchTerm string)  {
	// retrieve the list of feeds t search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// create a unbuffered channel to receive match results
	results := make(chan *Result)

	//setup a wait group so that we can process all the feeds
	var waitGroup sync.WaitGroup

	// set the number of goroutines we need to wait for while
	// they process the individual feeds
	waitGroup.Add(len(feeds))

	// launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// launch the goroutine to perform the search
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// launch a goroutine to monitor when all the work is done.

	go func() {
		waitGroup.Wait()

		// close the channel to signal to the display
		// function that we can exit the program
		close(results)
	}()

	// start displaying the results as they are available
	// return after the final result is displayed
	Display(results)
}
