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

type PostRepository interface {
	CreatePost(context.Context, PostCreateRequest) (*Post, error)
	UpdatePost(context.Context, PostUpdateRequest, string) (*Post, error)
	DeletePost(context.Context, string) error
	ListPosts(context.Context) ([]Post, error)
	GetById(context.Context, string) (*Post, error)
	UpvotePost(context.Context, string) error
	DownvotePost(context.Context, string) error
	ReportPost(context.Context, string, ReportPostRequest, string) (*Post, error)
}

type CommentRepository interface {
	GetById(context.Context, string) (*Comment, error)
	ListComments(context.Context, string) ([]Comment, error)
	AddComment(context.Context, CommentCreateRequest) (*Comment, error)
	UpdateComment(context.Context, CommentUpdateRequest, string) (*Comment, error)
	DeleteComment(context.Context, string) error
	UpvoteComment(context.Context, string) error
	DownvoteComment(context.Context, string) error
}

type postRepositoryImpl struct {
	db *gorm.DB
}

type commentRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepositoryImpl{db: db}
}

func (r *commentRepositoryImpl) DeleteComment(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("comment_id = ?", id).Delete(&Comment{})

	if result.Error != nil {
		return fmt.Errorf("delete comment")
	}
	return nil
}

func (r *commentRepositoryImpl) UpdateComment(ctx context.Context, req CommentUpdateRequest, id string) (*Comment, error) {
	comment, err := r.GetById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("Cannot find post")
	}

	comment.Content = req.Content

	if err := r.db.WithContext(ctx).Where("comment_id =?", id).Updates(comment).Error; err != nil {
		return nil, fmt.Errorf("Cannot update comment")
	}

	return comment, nil
}

func (r *postRepositoryImpl) UpdatePost(ctx context.Context, req PostUpdateRequest, id string) (*Post, error) {
	post, err := r.GetById(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("Cannot find post")
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Category = req.Category
	post.University = req.University
	post.Department = req.Department
	post.Faculty = req.Faculty

	if err := r.db.WithContext(ctx).Where("post_id = ?", id).Updates(post).Error; err != nil {
		return nil, fmt.Errorf("Cannot update post")
	}

	return post, nil
}

func (r *commentRepositoryImpl) AddComment(ctx context.Context, req CommentCreateRequest) (*Comment, error) {
	comment := &Comment{
		AuthorID: req.AuthorID,
		Content:  req.Content,
		PostID:   req.PostID,
	}

	if err := r.db.WithContext(ctx).Create(&comment).Error; err != nil {
		return nil, fmt.Errorf("Cannot create comment")
	}

	return comment, nil
}

func (r *commentRepositoryImpl) ListComments(ctx context.Context, id string) ([]Comment, error) {
	var comments []Comment

	if err := r.db.WithContext(ctx).Find(&comments, "post_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("Cannot list comments")
	}

	return comments, nil
}

func (r *postRepositoryImpl) ReportPost(ctx context.Context, id string, req ReportPostRequest, userID string) (*Post, error) {
	post, err := r.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("report post: %w", err)
	}

	report := &PostReport{
		PostID: post.PostID,
		UserID: userID,
		Reason: req.Reason,
	}

	if err := r.db.WithContext(ctx).Create(report).Error; err != nil {
		return nil, fmt.Errorf("report post %s: %w", id, err)
	}

	return post, nil
}

func (r *postRepositoryImpl) GetById(ctx context.Context, id string) (*Post, error) {
	var post Post

	if err := r.db.WithContext(ctx).First(&post, "post_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("Cannot find post with id")
	}

	return &post, nil
}

func (r *commentRepositoryImpl) GetById(ctx context.Context, id string) (*Comment, error) {
	var comment Comment

	if err := r.db.WithContext(ctx).First(&comment, "comment_id =?", id).Error; err != nil {
		return nil, fmt.Errorf("Cannot find comment with id")
	}

	return &comment, nil
}

func (r *postRepositoryImpl) DeletePost(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("post_id = ?", id).Delete(&Post{})

	if result.Error != nil {
		return fmt.Errorf("delete post")
	}
	return nil
}

func (r *postRepositoryImpl) UpvotePost(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Model(&Post{}).
		Where("post_id = ?", id).
		UpdateColumn("upvote", gorm.Expr("upvote + 1"))
	if result.Error != nil {
		return fmt.Errorf("upvote post %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("post %s not found", id)
	}
	return nil
}

func (r *postRepositoryImpl) DownvotePost(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Model(&Post{}).
		Where("post_id = ?", id).
		UpdateColumn("downvote", gorm.Expr("downvote + 1"))
	if result.Error != nil {
		return fmt.Errorf("downvote post %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("post %s not found", id)
	}
	return nil
}

func (r *commentRepositoryImpl) UpvoteComment(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Model(&Comment{}).
		Where("comment_id = ?", id).
		UpdateColumn("upvote", gorm.Expr("upvote + 1"))
	if result.Error != nil {
		return fmt.Errorf("upvote comment %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("comment %s not found", id)
	}
	return nil
}

func (r *commentRepositoryImpl) DownvoteComment(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Model(&Comment{}).
		Where("comment_id = ?", id).
		UpdateColumn("downvote", gorm.Expr("downvote + 1"))
	if result.Error != nil {
		return fmt.Errorf("downvote comment %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("comment %s not found", id)
	}
	return nil
}

func (r *postRepositoryImpl) CreatePost(ctx context.Context, req PostCreateRequest) (*Post, error) {
	post := &Post{
		PostID:     uuid.NewString(),
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		University: req.University,
		Department: req.Department,
		Faculty:    req.Faculty,
	}

	if err := r.db.WithContext(ctx).Create(post).Error; err != nil {
		return nil, fmt.Errorf("Cannot create post")
	}

	return post, nil
}

func (r *postRepositoryImpl) ListPosts(ctx context.Context) ([]Post, error) {
	var posts []Post

	if err := r.db.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("Cannot list posts")
	}

	return posts, nil
}
