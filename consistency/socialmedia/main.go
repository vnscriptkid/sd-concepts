package main

import (
	"fmt"
	"sync"
	"time"
)

// Post represents a social media post
type Post struct {
	Content string
	Time    time.Time
}

// Feed represents a user's feed
type Feed struct {
	Posts []Post
	mutex sync.Mutex
}

// AddPost adds a post to the feed
func (f *Feed) AddPost(post Post) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.Posts = append(f.Posts, post)
}

// GetPosts retrieves the posts from the feed
func (f *Feed) GetPosts() []Post {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.Posts
}

// Simulate eventual consistency with a delay
func propagatePost(followers []*Feed, post Post, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second) // Simulate network delay
	for _, feed := range followers {
		feed.AddPost(post)
	}
}

func main() {
	// User's own feed
	userFeed := &Feed{}
	// Followers' feeds
	follower1Feed := &Feed{}
	follower2Feed := &Feed{}
	followers := []*Feed{follower1Feed, follower2Feed}

	// User creates a post
	post := Post{Content: "Hello, world!", Time: time.Now()}
	userFeed.AddPost(post)

	var wg sync.WaitGroup
	wg.Add(1)
	// Propagate the post to followers after a delay
	go propagatePost(followers, post, &wg)

	// Immediately check followers' feeds (they won't have the new post yet)
	fmt.Println("Follower 1 Feed Before Propagation:", follower1Feed.GetPosts())
	fmt.Println("Follower 2 Feed Before Propagation:", follower2Feed.GetPosts())

	wg.Wait()

	// Check followers' feeds after propagation
	fmt.Println("Follower 1 Feed After Propagation:", follower1Feed.GetPosts())
	fmt.Println("Follower 2 Feed After Propagation:", follower2Feed.GetPosts())
}
