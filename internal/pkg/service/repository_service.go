package service

import (
	. "github.com/dinowar/rakia-home-task/internal/pkg/domain/model"
	"sync"
)

type RepositoryService struct {
	db   map[string]*Post
	lock sync.RWMutex
}

func NewRepositoryService() *RepositoryService {
	return &RepositoryService{db: make(map[string]*Post)}
}

func (rep *RepositoryService) CreatePost(post *Post) *Post {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	rep.db[post.Id] = post
	return post
}

func (rep *RepositoryService) GetPost(id string) *Post {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	return rep.db[id]
}

func (rep *RepositoryService) GetPosts(author string) []*Post {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	var posts []*Post
	for _, post := range rep.db {
		if author != "" && post.Author == author {
			posts = append(posts, post)
		} else {
			posts = append(posts, post)
		}
	}
	return posts
}

func (rep *RepositoryService) UpdatePost(id string, title, content string) *Post {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	post, exists := rep.db[id]
	if !exists {
		return nil
	}
	if post.Title != "" {
		post.Title = title
	}
	if post.Content != "" {
		post.Content = content
	}
	rep.db[id] = post
	return post
}

func (rep *RepositoryService) DeletePost(id string) *Post {
	rep.lock.Lock()
	defer rep.lock.Unlock()
	post, exists := rep.db[id]
	if !exists {
		return nil
	}
	delete(rep.db, id)
	return post
}
