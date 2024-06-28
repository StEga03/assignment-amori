package constant

// HandlerID represent of type aliasing for router handler identifier.
type HandlerID string

// MiddlewareID represent of type aliasing for middleware identifier.
type MiddlewareID string

// HTTPHandlerID represent list of handler id that used by the router,
// HandlerID must be unique.
const (
	// Auth HandlerID
	HTTPHandlerIDAuthCodeSend   HandlerID = "auth-code-send"
	HTTPHandlerIDAuthCodeVerify HandlerID = "auth-code-verify"
	HTTPHandlerIDTokenRefresh   HandlerID = "auth-token-refresh"

	// Users HandlerID.
	HTTPHandlerIDChannelCreate        HandlerID = "channel-create"
	HTTPHandlerIDChannelMessageCreate HandlerID = "channel-message-create"
	HTTPHandlerIDChannelMessageGet    HandlerID = "channel-message-get"

	// Uploads HandlerID.
	HTTPHandlerIDUploadMessageSource HandlerID = "upload-message-source"
)

// HTTP Header.
const (
	// Header Key.
	HTTPHeaderContentType    = "Content-Type"
	HTTPHeaderAccept         = "Accept"
	HTTPHeaderAcceptLanguage = "Accept-Language"
	HTTPHeaderAuthorization  = "Authorization"
	HTTPHeaderClientID       = "Client-Id"
	HTTPHeaderClientToken    = "Client-Token"
	HTTPHeaderUserID         = "User-ID"
	HTTPHeaderRequestID      = "X-Request-ID"
	HTTPHeaderTimestamp      = "timestamp"
	HTTPHeaderSignature      = "Signature"
	HTTPHeaderBearer         = "Bearer"

	// Header Value.
	HTTPContentTypeJSON        = "application/json"
	HTTPContentTypePostURLForm = "application/x-www-form-urlencoded"
	HTTPContentTypeImageJPEG   = "image/jpeg"

	// Type of response.
	HTTPDefaultResponseWriter MiddlewareID = "http-default-response-writer"
	HTTPPlainResponseWriter   MiddlewareID = "http-plain-response-writer"
)