package board

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Board
//   - GET /board/posts — list posts (filter by faculty/department)
//   - POST /board/posts — create anonymous post
//   - GET /board/posts/:id — get post detail
//   - DELETE /board/posts/:id — delete own post
//   - POST /board/posts/:id/upvote — upvote
//   - POST /board/posts/:id/downvote — downvote
//   - POST /board/posts/:id/report — report post
//   - GET /board/posts/:id/comments — list comments
//   - POST /board/posts/:id/comments — add comment
//   - DELETE /board/comments/:id — delete own comment

type Repository interface {
	CreatePost(context.Context, PostCreateRequest) (*Post, error)
	ListPosts(context.Context) ([]Post, error)
	GetById(context.Context, string) (*Post, error)
	DeletePost(context.Context, string) error
	// UpvotePost()
	// DownvotePost()
	// ReportPost()
	// ListComments()
	// AddComment()
	// EditComment()
	// DeleteComment()
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetById(ctx context.Context, id string) (*Post, error) {
	var post Post

	if err := r.db.WithContext(ctx).First(&post, "post_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("Cannot find post with id")
	}

	return &post, nil
}

func (r *repositoryImpl) DeletePost(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("post_id = ?", id).Delete(&Post{})

	if result.Error != nil {
		return fmt.Errorf("delete post")
	}
	return nil
}

func (r *repositoryImpl) CreatePost(ctx context.Context, req PostCreateRequest) (*Post, error) {
	post := &Post{
		PostID:  uuid.NewString(),
		Title:   req.Title,
		Content: req.Content,
	}

	if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
		return nil, fmt.Errorf("Cannot create post")
	}

	return post, nil
}

func (r *repositoryImpl) ListPosts(ctx context.Context) ([]Post, error) {
	var posts []Post

	if err := r.db.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("Cannot list posts")
	}

	return posts, nil
}
