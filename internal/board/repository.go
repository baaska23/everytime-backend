package board

import (
	"context"
	"errors"

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

// Future-proof contract types (V2). Keep current interfaces below for backward compatibility.
type ErrorCode string

const (
	ErrCodeNotFound    ErrorCode = "not_found"
	ErrCodeConflict    ErrorCode = "conflict"
	ErrCodeForbidden   ErrorCode = "forbidden"
	ErrCodeInvalid     ErrorCode = "invalid_argument"
	ErrCodeUnauthorized ErrorCode = "unauthorized"
	ErrCodeInternal    ErrorCode = "internal"
)

type DomainError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *DomainError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Code)
}

func (e *DomainError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

var (
	ErrNotFound     = &DomainError{Code: ErrCodeNotFound, Message: "resource not found", Err: errors.New("resource not found")}
	ErrConflict     = &DomainError{Code: ErrCodeConflict, Message: "resource conflict", Err: errors.New("resource conflict")}
	ErrForbidden    = &DomainError{Code: ErrCodeForbidden, Message: "operation forbidden", Err: errors.New("operation forbidden")}
	ErrInvalidInput = &DomainError{Code: ErrCodeInvalid, Message: "invalid input", Err: errors.New("invalid input")}
)

type Pagination struct {
	Page  int
	Limit int
}

type CursorPagination struct {
	Cursor string
	Limit  int
}

type SortOption struct {
	Field string
	Desc  bool
}

type PagedResult[T any] struct {
	Items []T
	Total int64
	Page  int
	Limit int
}

type CursorPageResult[T any] struct {
	Items      []T
	NextCursor string
	HasMore    bool
}

type BoardScope struct {
	University string
	Faculty    string
	Department string
}

type PostListFilter struct {
	Pagination
	CursorPagination
	Scope           BoardScope
	Category        string
	Search          string
	IncludeReported bool
	Sort            SortOption
}

type CommentListFilter struct {
	Pagination
	CursorPagination
	PostID string
	Sort   SortOption
}

type CreatePostInput struct {
	UserID   string
	Title    string
	Content  string
	Category string
	Scope    BoardScope
}

type UpdatePostInput struct {
	Title    *string
	Content  *string
	Category *string
	Scope    *BoardScope
}

type CreateCommentInput struct {
	PostID  string
	UserID  string
	Content string
}

type UpdateCommentInput struct {
	Content *string
}

type VoteType string

const (
	VoteUp   VoteType = "up"
	VoteDown VoteType = "down"
)

type PostVoteInput struct {
	PostID   string
	UserID   string
	VoteType VoteType
}

type CommentVoteInput struct {
	CommentID string
	UserID    string
	VoteType  VoteType
}

type CreatePostReportInput struct {
	PostID  string
	UserID  string
	Reason  string
	Details string
}

type TxRunner interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type PostReadRepositoryV2 interface {
	GetByID(ctx context.Context, postID string) (*Post, error)
	List(ctx context.Context, filter PostListFilter) (PagedResult[Post], error)
	ListByCursor(ctx context.Context, filter PostListFilter) (CursorPageResult[Post], error)
}

type PostWriteRepositoryV2 interface {
	Create(ctx context.Context, req CreatePostInput) (*Post, error)
	Update(ctx context.Context, postID string, req UpdatePostInput) (*Post, error)
	Delete(ctx context.Context, postID string) error
}

type PostVoteRepositoryV2 interface {
	Vote(ctx context.Context, req PostVoteInput) error
}

type PostReportRepositoryV2 interface {
	Report(ctx context.Context, req CreatePostReportInput) (*PostReport, error)
}

type CommentReadRepositoryV2 interface {
	GetByID(ctx context.Context, commentID string) (*Comment, error)
	List(ctx context.Context, filter CommentListFilter) (PagedResult[Comment], error)
	ListByCursor(ctx context.Context, filter CommentListFilter) (CursorPageResult[Comment], error)
}

type CommentWriteRepositoryV2 interface {
	Create(ctx context.Context, req CreateCommentInput) (*Comment, error)
	Update(ctx context.Context, commentID string, req UpdateCommentInput) (*Comment, error)
	Delete(ctx context.Context, commentID string) error
	DeleteByPostID(ctx context.Context, postID string) error
}

type CommentVoteRepositoryV2 interface {
	Vote(ctx context.Context, req CommentVoteInput) error
}

type PostRepositoryV2 interface {
	TxRunner
	PostReadRepositoryV2
	PostWriteRepositoryV2
	PostVoteRepositoryV2
	PostReportRepositoryV2
}

type CommentRepositoryV2 interface {
	TxRunner
	CommentReadRepositoryV2
	CommentWriteRepositoryV2
	CommentVoteRepositoryV2
}

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