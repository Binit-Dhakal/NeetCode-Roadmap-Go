package heappq

import (
	"container/heap"
)

type Post struct {
	TweetId int
	Date    int64
}

type Posts []Post

func (p Posts) Len() int { return len(p) }

func (p Posts) Less(i, j int) bool { return p[i].Date < p[j].Date }

func (p Posts) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *Posts) Push(x any) {
	*p = append(*p, x.(Post))
}

func (p *Posts) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

type Followees map[int]struct{}

type Twitter struct {
	posts map[int]Posts
	users map[int]Followees
	count int
}

func ConstructTwitter() Twitter {
	posts := make(map[int]Posts, 0)
	users := make(map[int]Followees, 0)
	return Twitter{
		posts: posts,
		users: users,
		count: 0,
	}
}

func (this *Twitter) AllUsers() (map[int]Followees, map[int]Posts) {
	return this.users, this.posts
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.users[userId]; !ok {
		this.users[userId] = Followees{}
	}
	post := Post{
		TweetId: tweetId,
		Date:    int64(this.count),
	}
	this.posts[userId] = append(this.posts[userId], post)
	this.count++
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = Followees{}
	}

	this.users[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if follows, ok := this.users[followerId]; ok {
		delete(follows, followeeId)
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	n := 10
	h := Posts{}

	heap.Init(&h)

	findPosts := func(userId int) {
		posts := this.posts[userId]
		for i := len(posts) - 1; i >= 0; i-- {
			post := posts[i]
			if len(h) < n {
				heap.Push(&h, post)
				continue
			}
			if post.Date > h[0].Date {
				heap.Pop(&h)
				heap.Push(&h, post)
			} else {
				// logic checks cuz posts are in descending order
				break
			}
		}
	}

	findPosts(userId)
	for followee := range this.users[userId] {
		findPosts(followee)
	}

	res := make([]int, len(h))

	for i := len(h) - 1; i >= 0; i-- {
		x := heap.Pop(&h).(Post)
		res[i] = x.TweetId
	}

	return res
}
