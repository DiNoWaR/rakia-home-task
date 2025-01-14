package service

import (
	"sync"
	"testing"

	. "github.com/dinowar/rakia-home-task/internal/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestNewRepositoryService(t *testing.T) {
	db := map[string]*Post{
		"1": {Id: "1", Title: "Post 1", Content: "Content 1", Author: "Author 1"},
	}

	service := &RepositoryService{
		db:     db,
		logger: nil,
		lock:   sync.RWMutex{},
	}

	assert.NotNil(t, service)
	assert.Equal(t, 1, len(service.db))
}

func TestRepositoryService_CreatePost(t *testing.T) {
	db := make(map[string]*Post)
	repo := &RepositoryService{
		db:     db,
		logger: nil,
	}

	post := &Post{Id: "1", Title: "New Post", Content: "New Content", Author: "Author"}
	repo.CreatePost(post)

	assert.Equal(t, 1, len(repo.db))
	assert.Equal(t, post, repo.db["1"])
}

func TestRepositoryService_GetPost(t *testing.T) {
	db := map[string]*Post{
		"1": {Id: "1", Title: "Post 1", Content: "Content 1", Author: "Author 1"},
	}
	repo := &RepositoryService{
		db: db,
	}

	post := repo.GetPost("1")
	assert.NotNil(t, post)
	assert.Equal(t, "Post 1", post.Title)

	post = repo.GetPost("2")
	assert.Nil(t, post)
}

func TestRepositoryService_UpdatePost(t *testing.T) {
	db := map[string]*Post{
		"1": {Id: "1", Title: "Post 1", Content: "Content 1", Author: "Author 1"},
	}
	repo := &RepositoryService{
		db: db,
	}

	updated := repo.UpdatePost("1", "Updated Title", "Updated Content")
	assert.NotNil(t, updated)
	assert.Equal(t, "Updated Title", updated.Title)
	assert.Equal(t, "Updated Content", updated.Content)

	nonExistent := repo.UpdatePost("2", "Title", "Content")
	assert.Nil(t, nonExistent)
}

func TestRepositoryService_DeletePost(t *testing.T) {
	db := map[string]*Post{
		"1": {Id: "1", Title: "Post 1", Content: "Content 1", Author: "Author 1"},
	}
	repo := &RepositoryService{
		db: db,
	}

	deleted := repo.DeletePost("1")
	assert.NotNil(t, deleted)
	assert.Equal(t, "Post 1", deleted.Title)
	assert.Equal(t, 0, len(repo.db))

	nonExistent := repo.DeletePost("2")
	assert.Nil(t, nonExistent)
}
