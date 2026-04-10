package board

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostID     string `gorm:"uniqueIndex;not null" json:"postId"`
	AuthorID   string `gorm:"not null;index" json:"-"`
	University string `gorm:"not null;index" json:"university"`
	Faculty    string `gorm:"not null;index" json:"faculty"`
	Department string `gorm:"index" json:"department"`
	Title      string `gorm:"not null" json:"title"`
	Content    string `gorm:"not null" json:"content"`
	Category   string `gorm:"not null;index" json:"category"`
	Upvote     int    `gorm:"default:0" json:"upvote"`
	Downvote   int    `gorm:"default:0" json:"downvote"`
	IsReported bool   `gorm:"default:false" json:"isReported"`
}

type Comment struct {
	gorm.Model
	CommentID string `gorm:"uniqueIndex;not null" json:"commentId"`
	PostID    string `gorm:"not null;index" json:"postId"`
	AuthorID  string `gorm:"not null;index" json:"-"`
	Content   string `gorm:"not null" json:"content"`
	Post      *Post  `gorm:"foreignKey:PostID;references:PostID" json:"-"`
}

type CommentReply struct {
	gorm.Model
	ReplyID   string   `gorm:"uniqueIndex;not null" json:"replyId"`
	CommentID string   `gorm:"not null;index" json:"commentId"`
	AuthorID  string   `gorm:"not null;index" json:"-"`
	Content   string   `gorm:"not null" json:"content"`
	Comment   *Comment `gorm:"foreignKey:CommentID;references:CommentID" json:"-"`
}

// DB models
type PostVote struct {
	gorm.Model
	PostID   string `gorm:"not null;uniqueIndex:idx_post_user_vote" json:"postId"`
	UserID   string `gorm:"not null;uniqueIndex:idx_post_user_vote" json:"userId"`
	VoteType string `gorm:"not null;check:vote_type IN ('up', 'down')" json:"voteType"`
}

type PostReport struct {
	gorm.Model
	PostID string `gorm:"not null;index;constraint:OnDelete:CASCADE" json:"postId"`
	UserID string `gorm:"not null" json:"userId"`
	Reason string `gorm:"not null" json:"reason"`
	Status string `gorm:"not null;default:'pending'" json:"status"`
	Post   *Post  `gorm:"foreignKey:PostID;references:PostID" json:"-"`
}

type ReportPostRequest struct {
	Reason string `json:"reason"`
}

type VotePostRequest struct {
	VoteType string `json:"voteType"`
}

type PostCreateRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	University string `json:"university"`
	Department string `json:"department"`
	Faculty    string `json:"faculty"`
}

type PostUpdateRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	University string `json:"university"`
	Department string `json:"department"`
	Faculty    string `json:"faculty"`
}

type CommentCreateRequest struct {
	AuthorID string `json:"userId"`
	Content  string `json:"content"`
	PostID   string `json:"postId"`
}

type CommentUpdateRequest struct {
	AuthorID string `json:"userId"`
	Content  string `json:"content"`
	PostID   string `json:"postId"`
}

type FeedItem struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
