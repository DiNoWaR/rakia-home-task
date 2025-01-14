package service

import (
	"encoding/json"
	. "github.com/dinowar/rakia-home-task/internal/pkg/domain/model"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

var dataPath = "./data/blog_data.json"

type PostsWrapper struct {
	Posts []PostAdapter `json:"posts"`
}

type PostAdapter struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type RepositoryService struct {
	db     map[string]*Post
	logger *LogService
	lock   sync.RWMutex
}

func NewRepositoryService(logger *LogService) *RepositoryService {
	db := make(map[string]*Post)
	initErr := initDB(db, logger.logger)
	if initErr != nil {
		logger.logger.Error(initErr.Error())
	}
	return &RepositoryService{db: db, logger: logger}
}

func initDB(db map[string]*Post, logger *zap.Logger) error {
	file, openErr := os.Open(dataPath)
	if os.IsNotExist(openErr) {
		logger.Error(openErr.Error())
		return nil
	}

	defer file.Close()

	bytes, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return readErr
	}

	var postsWrapper PostsWrapper
	unmarshalErr := json.Unmarshal(bytes, &postsWrapper)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	for _, post := range postsWrapper.Posts {
		postId := strconv.Itoa(post.Id)
		db[postId] = &Post{
			Id:      postId,
			Title:   post.Title,
			Content: post.Content,
			Author:  post.Author,
		}
	}

	return nil
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
