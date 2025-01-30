package heappq_test

import (
	"reflect"
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
)

func TestDesignTwitter(t *testing.T) {
	t.Run("Own Example", func(t *testing.T) {
		obj := heappq.ConstructTwitter()
		obj.PostTweet(1, 100)
		obj.PostTweet(1, 102)
		_, post := obj.AllUsers()
		if post[1][0].TweetId != 100 {
			t.Error("tweetid 100, not saved for user 1")
		}

		obj.Follow(1, 2)
		obj.Follow(1, 3)
		obj.Follow(2, 3)
		user, _ := obj.AllUsers()
		if !reflect.DeepEqual(user[1], heappq.Followees{2: struct{}{}, 3: struct{}{}}) {
			t.Errorf("Got: %v, want: %v", user[1], []int{2, 3})
		}

		obj.Unfollow(1, 3)
		user, _ = obj.AllUsers()
		if !reflect.DeepEqual(user[1], heappq.Followees{2: struct{}{}}) {
			t.Errorf("Got: %v, want: %v", user[1], []int{2})
		}

		multiplePostsInsert := []int{300, 400, 500, 600, 700, 301, 401, 501, 601, 701, 801, 901, 101}
		for _, post := range multiplePostsInsert {
			obj.PostTweet(2, post)
		}

		got := obj.GetNewsFeed(1)
		want := []int{101, 901, 801, 701, 601, 501, 401, 301, 700, 600}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})

	t.Run("Leetcode Example", func(t *testing.T) {
		obj := heappq.ConstructTwitter()
		obj.PostTweet(1, 5)
		// users, posts := obj.AllUsers()

		got := obj.GetNewsFeed(1)
		want := []int{5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Double Follow", func(t *testing.T) {
		obj := heappq.ConstructTwitter()
		obj.PostTweet(2, 5)
		obj.Follow(1, 2)
		obj.Follow(1, 2)

		got := obj.GetNewsFeed(1)
		want := []int{5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})
}
