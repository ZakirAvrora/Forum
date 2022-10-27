package store

import "Forum/domain/schema"

func (s Store) ListComment() ([]schema.Comment, error) {
	var comments []schema.Comment

	return comments, nil
}

func (s *Store) CreateComment(comment schema.Comment, post schema.Post) error {
	return nil
}

func (s *Store) DeleteComment(comment schema.Comment, post schema.Post) error {
	return nil
}
