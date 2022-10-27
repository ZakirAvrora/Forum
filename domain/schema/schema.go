package schema

import "time"

type Post struct {
	PostID       int
	PostUser     string
	PostTitle    string
	PostContent  string
	PostCreated  time.Time
	PostLikes    int
	PostDislikes int
	Category     []Category
}

type Comment struct {
	CommentID      int
	CommentUser    string
	CommentPost    int
	CommentContent string
	CommentCreated time.Time
}

type Category struct {
	CategoryName string
}

type User struct {
	UserEmail string
	UserName  string
	Passwd    string
}

type PostComments struct {
	Post     Post
	Comments []Comment
}
