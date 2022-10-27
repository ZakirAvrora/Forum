package store

import "Forum/domain/schema"

func (s *Store) ListPosts() ([]schema.Post, error) {
	var posts []schema.Post
	//rows, err := s.db.Query("SELECT ")

	return posts, nil
}

func (s *Store) CreatePost(p schema.Post) error {
	stmt, err := s.db.Prepare("INSERT INTO Posts(post_title, post_user, post_content, post_likes, post_dislikes, post_createdAt) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.PostTitle, p.PostUser, p.PostContent,
		p.PostLikes, p.PostDislikes, p.PostCreated)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateLikes(p schema.Post) error {
	stmt, err := s.db.Prepare("UPDATE Posts SET post_likes = post_likes+1 WHERE post_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.PostID)
	return nil
}

func (s *Store) UpdateDislikes(p schema.Post) error {
	stmt, err := s.db.Prepare("UPDATE Posts SET post_dislikes = post_dislikes+1 WHERE post_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.PostID)
	return nil
}
