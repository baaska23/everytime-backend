package courses

//   Courses
//   - GET /courses — list courses
//   - GET /courses/:id — get course detail + reviews
//   - POST /courses/:id/reviews — submit course review
//   - PUT /courses/:id/reviews/:reviewId — update own review
//   - DELETE /courses/:id/reviews/:reviewId — delete own review
//   - GET /courses/professors — list professors
//   - GET /courses/professors/:id — professor detail + ratings
//   - POST /courses/professors/:id/ratings — rate professor


// Courses (Future-Proof API Surface)
// Public Student Endpoints
// - GET /courses
// list courses with filters: university, faculty, department, category, search, creditsMin, creditsMax, page, limit, sort
// - GET /courses/:courseId
// course detail with summary aggregates: reviewCount, avgDifficulty, avgWorkload, avgTeachingQuality
// - GET /courses/:courseId/reviews
// paginated course reviews with filters: rating, semester, year, page, limit, sort
// - POST /courses/:courseId/reviews
// create own review (one review per user per course per semester recommended)
// - GET /courses/:courseId/reviews/:reviewId
// single review detail
// - PUT /courses/:courseId/reviews/:reviewId
// update own review
// - DELETE /courses/:courseId/reviews/:reviewId
// soft-delete own review
// - POST /courses/:courseId/reviews/:reviewId/helpful
// mark review helpful
// - DELETE /courses/:courseId/reviews/:reviewId/helpful
// remove helpful vote
// - POST /courses/:courseId/reviews/:reviewId/report
// report abusive/inaccurate review

// Professor Endpoints
// - GET /courses/professors
// list professors with filters: university, faculty, department, search, page, limit, sort
// - GET /courses/professors/:professorId
// professor profile + aggregate ratings
// - GET /courses/professors/:professorId/ratings
// paginated ratings with filters: courseId, semester, year, page, limit, sort
// - POST /courses/professors/:professorId/ratings
// create own professor rating
// - GET /courses/professors/:professorId/ratings/:ratingId
// single rating detail
// - PUT /courses/professors/:professorId/ratings/:ratingId
// update own rating
// - DELETE /courses/professors/:professorId/ratings/:ratingId
// soft-delete own rating
// - POST /courses/professors/:professorId/ratings/:ratingId/helpful
// mark rating helpful
// - DELETE /courses/professors/:professorId/ratings/:ratingId/helpful
// remove helpful vote
// - POST /courses/professors/:professorId/ratings/:ratingId/report
// report abusive/inaccurate rating

// Metadata and Discovery
// - GET /courses/filters
// return faculties, departments, categories, semesters, sort options
// - GET /courses/trending
// trending courses by review volume and score windows
// - GET /courses/professors/top
// top professors by configured scoring rules

// Admin and Moderation Endpoints
// - GET /admin/courses/reviews/reports
// list reported course reviews
// - PUT /admin/courses/reviews/reports/:reportId
// resolve report (approve, dismiss)
// - GET /admin/courses/professors/ratings/reports
// list reported professor ratings
// - PUT /admin/courses/professors/ratings/reports/:reportId
// resolve report
// - DELETE /admin/courses/reviews/:reviewId
// moderator remove review
// - DELETE /admin/courses/professors/ratings/:ratingId
// moderator remove rating

// Production-Ready Conventions For All Endpoints
// - enforce university scope on every read/write
// - use page + limit + stable sort for all list endpoints
// - support cursor pagination later for high-scale feeds
// - use soft delete for user content
// - enforce ownership on update/delete
// - keep strict validation and conflict handling (duplicate review/rating)
// - return consistent error codes and response shape
// - include audit fields: createdAt, updatedAt, deletedAt, createdBy where appropriate

type Repository interface {
	
}