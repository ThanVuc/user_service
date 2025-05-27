package response

type HTTPStatusCode int

const (
	OK                     HTTPStatusCode = 200
	CREATED                HTTPStatusCode = 201
	ACCEPTED               HTTPStatusCode = 202
	NO_CONTENT             HTTPStatusCode = 204
	BAD_REQUEST            HTTPStatusCode = 400
	UNAUTHORIZED           HTTPStatusCode = 401
	FORBIDDEN              HTTPStatusCode = 403
	NOT_FOUND              HTTPStatusCode = 404
	METHOD_NOT_ALLOWED     HTTPStatusCode = 405
	NOT_ACCEPTABLE         HTTPStatusCode = 406
	CONFLICT               HTTPStatusCode = 409
	UNSUPPORTED_MEDIA_TYPE HTTPStatusCode = 415
	INTERNAL_SERVER_ERROR  HTTPStatusCode = 500
	SERVICE_UNAVAILABLE    HTTPStatusCode = 503
)

var MSG = map[HTTPStatusCode]string{
	OK:                     "success",
	CREATED:                "created",
	ACCEPTED:               "accepted",
	NO_CONTENT:             "no content",
	BAD_REQUEST:            "bad request",
	UNAUTHORIZED:           "unauthorized",
	FORBIDDEN:              "forbidden",
	NOT_FOUND:              "not found",
	METHOD_NOT_ALLOWED:     "method not allowed",
	NOT_ACCEPTABLE:         "not acceptable",
	CONFLICT:               "conflict",
	UNSUPPORTED_MEDIA_TYPE: "unsupported media type",
	INTERNAL_SERVER_ERROR:  "internal server error",
	SERVICE_UNAVAILABLE:    "service unavailable",
}
