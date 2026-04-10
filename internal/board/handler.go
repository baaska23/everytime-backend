package board

import (
    "everytime-backend/internal/shared/apierror"
    "net/http"

    "github.com/gin-gonic/gin"
)

type PostHandler struct {
    postService *PostService
}

type CommentHandler struct {
    commentService *CommentService
}

func NewPostHandler(postService *PostService) *PostHandler {
    return &PostHandler{postService: postService}
}

func NewCommentHandler(commentService *CommentService) *CommentHandler {
    return &CommentHandler{commentService: commentService}
}

// GET /board/posts
func (h *PostHandler) ListPosts(c *gin.Context) {
    posts, err := h.postService.ListPosts(c.Request.Context())
    if err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": posts})
}

// POST /board/posts
func (h *PostHandler) CreatePost(c *gin.Context) {
    var req PostCreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        apierror.WriteGin(c, apierror.BadRequest("invalid request body"))
        return
    }

    userID := c.GetString("userID") // from auth middleware
    if userID == "" {
        apierror.WriteGin(c, apierror.Unauthorized("missing user"))
        return
    }

    post, err := h.postService.CreatePost(c.Request.Context(), req, userID)
    if err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusCreated, gin.H{"data": post})
}

// GET /board/posts/:id
func (h *PostHandler) GetPost(c *gin.Context) {
    id := c.Param("id")

    post, err := h.postService.GetPost(c.Request.Context(), id)
    if err != nil {
        apierror.WriteGin(c, apierror.NotFound("post not found"))
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /board/posts/:id
func (h *PostHandler) DeletePost(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID")

    if err := h.postService.DeletePost(c.Request.Context(), id, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

// POST /board/posts/:id/upvote
func (h *PostHandler) UpvotePost(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID")

    if err := h.postService.UpvotePost(c.Request.Context(), id, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "upvoted"})
}

// POST /board/posts/:id/downvote
func (h *PostHandler) DownvotePost(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID")

    if err := h.postService.DownvotePost(c.Request.Context(), id, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "downvoted"})
}

// POST /board/posts/:id/report
func (h *PostHandler) ReportPost(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID")

    var req ReportPostRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        apierror.WriteGin(c, apierror.BadRequest("invalid request body"))
        return
    }

    if _, err := h.postService.ReportPost(c.Request.Context(), id, req, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "reported"})
}

// GET /board/posts/:id/comments
func (h *CommentHandler) ListComments(c *gin.Context) {
    postID := c.Param("id")

    comments, err := h.commentService.ListComments(c.Request.Context(), postID)
    if err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": comments})
}

// POST /board/posts/:id/comments
func (h *CommentHandler) AddComment(c *gin.Context) {
    var req CommentCreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        apierror.WriteGin(c, apierror.BadRequest("invalid request body"))
        return
    }
    req.PostID = c.Param("id")
    req.AuthorID = c.GetString("userID")

    comment, err := h.commentService.AddComment(c.Request.Context(), req)
    if err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusCreated, gin.H{"data": comment})
}

// DELETE /board/comments/:id
func (h *CommentHandler) DeleteComment(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID")

    if err := h.commentService.DeleteComment(c.Request.Context(), id, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}

func (h *CommentHandler) UpdateComment(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID")

    var req CommentUpdateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        apierror.WriteGin(c, apierror.BadRequest("invalid request body"))
        return
    }

    comment, err := h.commentService.UpdateComment(c.Request.Context(), req, id, userID)
    if err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": comment}) // 👈 was "comment deleted"
}

func (h *CommentHandler) UpvoteComment(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID") // 👈 missing

    if err := h.commentService.UpvoteComment(c.Request.Context(), id, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "upvoted"})
}

func (h *CommentHandler) DownvoteComment(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetString("userID") // 👈 missing

    if err := h.commentService.DownvoteComment(c.Request.Context(), id, userID); err != nil {
        apierror.WriteGin(c, err)
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "downvoted"})
}