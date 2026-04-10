package board

import (
	"context"
	"everytime-backend/internal/ads"
	"everytime-backend/internal/shared/apierror"
	"fmt"
)

type PostService struct {
	postRepo PostRepository
}

type CommentService struct {
	commentRepo CommentRepository
	postRepo    PostRepository
}

type FeedService struct {
	postRepo PostRepository
	adRepo   ads.AdRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func NewCommentService(commentRepo CommentRepository, postRepo PostRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func NewFeedService(postRepo PostRepository, adRepo ads.AdRepository) *FeedService {
	return &FeedService{
		postRepo: postRepo,
		adRepo:   adRepo,
	}
}

func (s *FeedService) GetFeed(ctx context.Context) ([]FeedItem, error) {
	posts, err := s.postRepo.ListPosts(ctx)
	if err != nil {
		return nil, err
	}

	ads, err := s.adRepo.ListActiveBanners(ctx)
	if err != nil {
		return nil, err
	}

	return buildFeed(posts, ads), nil
}

func (s *PostService) GetPost(ctx context.Context, id string) (*Post, error) {
	return s.postRepo.GetById(ctx, id)
}

func (s *PostService) ListPosts(ctx context.Context) ([]Post, error) {
	return s.postRepo.ListPosts(ctx)
}

func (s *PostService) UpdatePost(ctx context.Context, req PostUpdateRequest, postID string, userID string) (*Post, error) {
	post, err := s.postRepo.GetById(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.AuthorID != userID {
		return nil, fmt.Errorf("unauthorized: you don't own this post")
	}
	return s.postRepo.UpdatePost(ctx, req, postID)
}

func (s *PostService) CreatePost(ctx context.Context, req PostCreateRequest, userID string) (*Post, error) {
	// Validate required fields
	if req.Title == "" {
		return nil, apierror.BadRequest("title is required")
	}
	if req.Content == "" {
		return nil, apierror.BadRequest("content is required")
	}

	return s.postRepo.CreatePost(ctx, req)
}

func (s *PostService) DeletePost(ctx context.Context, postID string, userID string) error {
	post, err := s.postRepo.GetById(ctx, postID)

	if err != nil {
		return err
	}

	if post.AuthorID != userID {
		return fmt.Errorf("unauthorized: you dont own this post")
	}

	return s.postRepo.DeletePost(ctx, postID)
}

func (s *PostService) ReportPost(ctx context.Context, postID string, req ReportPostRequest, userID string) (*Post, error) {
	post, err := s.postRepo.GetById(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.AuthorID == userID {
		return nil, fmt.Errorf("cannot report your own post")
	}
	return s.postRepo.ReportPost(ctx, postID, req, userID)
}

func (s *PostService) UpvotePost(ctx context.Context, postID string, userID string) error {
	post, err := s.postRepo.GetById(ctx, postID)
	if err != nil {
		return err
	}
	if post.AuthorID == userID {
		return fmt.Errorf("cannot upvote your own post")
	}
	return s.postRepo.UpvotePost(ctx, postID)
}

func (s *PostService) DownvotePost(ctx context.Context, postID string, userID string) error {
	post, err := s.postRepo.GetById(ctx, postID)
	if err != nil {
		return err
	}
	if post.AuthorID == userID {
		return fmt.Errorf("cannot downvote your own post")
	}
	return s.postRepo.DownvotePost(ctx, postID)
}

func (s *CommentService) GetComment(ctx context.Context, id string) (*Comment, error) {
	return s.commentRepo.GetById(ctx, id)
}

func (s *CommentService) ListComments(ctx context.Context, id string) ([]Comment, error) {
	return s.commentRepo.ListComments(ctx, id)
}

func (s *CommentService) AddComment(ctx context.Context, req CommentCreateRequest) (*Comment, error) {
	post, err := s.postRepo.GetById(ctx, req.PostID)
	if err != nil {
		return nil, fmt.Errorf("post not found")
	}
	if post.IsReported {
		return nil, fmt.Errorf("cannot comment on a reported post")
	}
	return s.commentRepo.AddComment(ctx, req)
}

func (s *CommentService) DeleteComment(ctx context.Context, commentID string, userID string) error {
	comment, err := s.commentRepo.GetById(ctx, commentID)
	if err != nil {
		return err
	}
	if comment.AuthorID != userID {
		return fmt.Errorf("unauthorized: you don't own this comment")
	}
	return s.commentRepo.DeleteComment(ctx, commentID)
}

func (s *CommentService) UpdateComment(ctx context.Context, req CommentUpdateRequest, commentID string, userID string) (*Comment, error) {
	comment, err := s.commentRepo.GetById(ctx, commentID)
	if err != nil {
		return nil, err
	}
	if comment.AuthorID != userID {
		return nil, fmt.Errorf("unauthorized: you don't own this comment")
	}
	return s.commentRepo.UpdateComment(ctx, req, commentID)
}

func (s *CommentService) DownvoteComment(ctx context.Context, commentID string, userID string) error {
	comment, err := s.commentRepo.GetById(ctx, commentID)

	if err != nil {
		return fmt.Errorf("Cannot find comment")
	}

	if comment.AuthorID == userID {
		return fmt.Errorf("Cannot downvote your own comment")
	}
	return s.commentRepo.DownvoteComment(ctx, commentID)
}

func (s *CommentService) UpvoteComment(ctx context.Context, commentID string, userID string) error {
	comment, err := s.commentRepo.GetById(ctx, commentID)

	if err != nil {
		return fmt.Errorf("Cannot find comment")
	}

	if comment.AuthorID == userID {
		return fmt.Errorf("Cannot upvote your own comment")
	}
	return s.commentRepo.UpvoteComment(ctx, commentID)
}

func buildFeed(posts []Post, ads []ads.Ad) []FeedItem {
	const adInterval = 20

	feed := []FeedItem{}
	adIndex := 0

	for i, post := range posts {
		feed = append(feed, FeedItem{Type: "post", Payload: post})

		if (i+1)%adInterval == 0 && adIndex < len(ads) {
			feed = append(feed, FeedItem{Type: "ad", Payload: ads})
			adIndex++
		}
	}

	return feed
}
