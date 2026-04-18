package marketplace

// Marketplace (Future-Proof API Surface)

// Public Student Endpoints
// - GET /marketplace/listings
// list listings with filters: university, faculty, category, condition, minPrice, maxPrice, search, status, page, limit, sort
// - POST /marketplace/listings
// create own listing
// - GET /marketplace/listings/:listingId
// listing detail (view-safe fields)
// - PUT /marketplace/listings/:listingId
// update own listing
// - DELETE /marketplace/listings/:listingId
// soft-delete own listing
// - PATCH /marketplace/listings/:listingId/status
// set status: active, reserved, sold, archived
// - POST /marketplace/listings/:listingId/interest
// express interest and create interest record
// - DELETE /marketplace/listings/:listingId/interest
// withdraw interest
// - GET /marketplace/listings/:listingId/interests
// owner-only list of interested users
// - POST /marketplace/listings/:listingId/interests/:interestId/reveal-contact
// owner approves and reveals contact details to selected interested user
// - POST /marketplace/listings/:listingId/report
// report listing
// - POST /marketplace/listings/:listingId/favorite
// save listing
// - DELETE /marketplace/listings/:listingId/favorite
// unsave listing

// My Marketplace Endpoints
// - GET /marketplace/me/listings
// list my listings with status filters
// - GET /marketplace/me/interests
// list interests I sent
// - GET /marketplace/me/favorites
// list saved listings
// - GET /marketplace/me/contact-reveals
// list listings where my contact was revealed or I received seller contact

// Admin and Moderation Endpoints
// - GET /admin/marketplace/reports
// list reported listings
// - PUT /admin/marketplace/reports/:reportId
// resolve report (approve, dismiss)
// - DELETE /admin/marketplace/listings/:listingId
// moderator remove listing
// - PATCH /admin/marketplace/listings/:listingId/visibility
// hide/unhide listing

// Production Conventions
// - enforce university scope on every read/write
// - pagination + stable sorting on every list endpoint
// - ownership checks for update/delete/status/interests views
// - idempotent interest/favorite operations
// - soft delete for listings and moderation traceability
// - audit fields for contact reveal actions
// - consistent error shape: bad_request, unauthorized, forbidden, not_found, conflict
// - anti-abuse controls: rate limit interest/report endpoints

type Repository interface {
// split later into ListingRepository, InterestRepository, FavoriteRepository, ModerationRepository
}