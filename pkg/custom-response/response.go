package response

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type Response struct {
// 	Code      int       `json:"code"`
// 	Info      string    `json:"info"`
// 	Status    bool      `json:"status"`
// 	Timestamp time.Time `json:"timestamp"`
// 	Message   string    `json:"message"`
// }

// func newResponse(code int, status bool, message, info string) *Response {
// 	return &Response{
// 		Code:      code,
// 		Info:      info,
// 		Status:    status,
// 		Timestamp: time.Now(),
// 		Message:   message,
// 	}
// }

// func sendResponse(c *gin.Context, code int, status bool, info, message string) {
// 	resp := newResponse(code, status, message, info)
// 	c.JSON(code, resp)
// }

// type ResponseMultiMessage struct {
// 	Code      int       `json:"code"`
// 	Info      string    `json:"info"`
// 	Status    bool      `json:"status"`
// 	Timestamp time.Time `json:"timestamp"`
// 	Message   []string  `json:"message"`
// }

// func newMultiMessageResponse(code int, status bool, message []string, info string) *ResponseMultiMessage {
// 	return &ResponseMultiMessage{
// 		Code:      code,
// 		Info:      info,
// 		Status:    status,
// 		Timestamp: time.Now(),
// 		Message:   message,
// 	}
// }

// func sendMultipleMessages(c *gin.Context, code int, status bool, info string, message []string) {
// 	resp := newMultiMessageResponse(code, status, message, info)
// 	c.JSON(code, resp)
// }

// // ------------------ Informational responses (100 – 199) ------------------

// // ------------------ Successful responses (200 – 299) ------------------

// // 200
// func Ok(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusOK, true, "OK", message)
// }

// // 201
// func Created(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusCreated, true, "CREATED", message)
// }

// // ------------------ Redirection messages (300 – 399) ------------------

// // ------------------ Client error responses (400 – 499) ------------------

// // 400 - Errors Array
// func ValidationBadRequest(c *gin.Context, message []string) {
// 	sendMultipleMessages(c, http.StatusBadRequest, false, "BAD_REQUEST", message)
// }

// // 400
// func BadRequest(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusBadRequest, false, "BAD_REQUEST", message)
// }

// // 401
// func Unauthorized(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusUnauthorized, false, "UNAUTHORIZED", message)
// }

// // 403
// func Forbidden(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusForbidden, false, "FORBIDDEN", message)
// }

// // 404
// func NotFound(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusNotFound, false, "NOT_FOUND", message)
// }

// // 405
// func MethodNotAllowed(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusMethodNotAllowed, false, "METHOD_NOT_ALLOWED", message)
// }

// // 409
// func Conflict(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusConflict, false, "CONFLICT", message)
// }

// // ------------------ Server error responses (500 – 599) ------------------

// // 500
// func InternalServerError(c *gin.Context, message string) {
// 	sendResponse(c, http.StatusInternalServerError, false, "INTERNAL_SERVER_ERROR", message)
// }
