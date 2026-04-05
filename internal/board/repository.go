package board

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"github.com/google/uuid"
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
	// ListPosts()
	// GetPostDetail()
	// DeletePost()
	// UpvotePost()
	// DownvotePost()
	// ReportPost()
	// ListComments()
	// AddComment()
	// EditComment()
	// DeleteComment()
}

type repositoryImpl struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) CreatePost(ctx context.Context, req PostCreateRequest) (*Post, error){
	post := &Post{
		PostID: uuid.NewString(),
		Title: req.Title,
		Content: req.Content,
	}

	if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
		return nil, fmt.Errorf("Cannot create post")
	}

	return post, nil
}