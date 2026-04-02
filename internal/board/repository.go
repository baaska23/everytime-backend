package board

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
	ListPosts()
	CreatePost()
	GetPostDetail()
	DeletePost()
	UpvotePost()
	DownvotePost()
	ReportPost()
	ListComments()
	AddComment()
	EditComment()
	DeleteComment()
}