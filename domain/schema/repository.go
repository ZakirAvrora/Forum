package schema

type Repository interface {
	UserRepository
	PostRepository
	CommentRepository
}

type UserRepository interface {
	ListUsers() ([]User, error)
	GetUserByEmail(email string) (User, error)
	CreateUser(user User) error
	DeleteUser(user User) error
}

type PostRepository interface {
	ListPosts() ([]Post, error)
	CreatePost(post Post) error
	UpdateLikes(post Post) error
	UpdateDislikes(post Post) error
}

type CommentRepository interface {
	ListComment() ([]Comment, error)
	CreateComment(comment Comment, post Post) error
	DeleteComment(comment Comment, post Post) error
}
