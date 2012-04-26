package main

import "github.com/johanwestin/go-twitter"
// import "github.com/csmcanarney/gooauth"
import "fmt"
import "time"
import "math/rand"
import "log"

const kMaxDepth = 5
const kStart = "3ch0"

var api *twitter.Api
var r *rand.Rand
var done chan bool

func main() {
	api = twitter.NewApi()
	done = make(chan bool)
	r = rand.New(rand.NewSource( time.Now().Unix() ))
	go crawl(kStart, 0)
	<-done
}

func crawl(userName string, level int) {
	// Get the user's status
	text := (<-api.GetUser(userName)).GetStatus().GetText()

	for i := 0; i < level; i++ {
		fmt.Printf("     ")
	}

	fmt.Printf("%s: %s\n", userName, text)

	level++
	if level > kMaxDepth {
		done <- true
		return
	}

	// Get the user's friends
	log.Println(userName)
	friends := <-api.GetFriends(userName, 1)
	log.Println(friends)
	length := len(friends)
	log.Println(length)
	if length == 0 {
		done <- true
		return
	}
	
	
	
	rVal := r.Intn(length - 1)
	// Choose a random friend for the next user
	nextUser := friends[rVal].GetScreenName()

	go crawl(nextUser, level)
}
