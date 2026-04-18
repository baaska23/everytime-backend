package timetable

// Timetable (Future-Proof API Surface)

// Student Endpoints
// - GET /timetable
// get active timetable with slots, conflicts, credit summary
// - GET /timetable/semesters
// list available semesters for the user
// - POST /timetable/semesters
// create timetable semester plan
// - PATCH /timetable/semesters/:semesterId/activate
// set active semester timetable
// - DELETE /timetable/semesters/:semesterId
// archive/delete a semester plan

// Slot Endpoints
// - GET /timetable/slots
// list slots for active semester with filters: day, courseId, type, page, limit, sort
// - POST /timetable/slots
// add slot (manual or from course section)
// - GET /timetable/slots/:slotId
// slot detail
// - PUT /timetable/slots/:slotId
// update slot (time/room/instructor/section)
// - DELETE /timetable/slots/:slotId
// remove slot
// - POST /timetable/slots/:slotId/duplicate
// duplicate slot to another week pattern or semester
// - POST /timetable/slots/bulk
// bulk add slots from selected course sections
// - DELETE /timetable/slots/bulk
// bulk remove slots

// Validation and Planning
// - POST /timetable/validate
// validate timetable for conflicts, overlaps, max credits, invalid gaps
// - GET /timetable/conflicts
// list detected conflicts with severity and affected slots
// - GET /timetable/recommendations
// suggest alternatives for conflicted slots or better schedules

// Import and Export
// - POST /timetable/import
// import timetable from JSON/ICS/template
// - GET /timetable/export
// export timetable to JSON/ICS

// Production Conventions
// - strict ownership: all timetable data is user-owned only
// - university scoping for course/section references
// - conflict checks on create/update and bulk operations
// - transaction-safe bulk mutations
// - soft delete or archive for semester plans
// - consistent error shape and validation messages
// - optimistic concurrency (version/updatedAt) on slot updates

type Repository interface {
// split later into SemesterRepository, SlotRepository, ValidationRepository
}