package hqgohttp

// This file list HTTP header field names as constants.
//
// HTTP header fields are components of the header section of request
// and response messages in the HTTP protocol, which is the
// foundation of data communication for the World Wide Web.

const (
	// Authentication - These header fields are related to authentication and authorization.
	HeaderAuthorization      = "Authorization"
	HeaderProxyAuthenticate  = "Proxy-Authenticate"
	HeaderProxyAuthorization = "Proxy-Authorization"
	HeaderWWWAuthenticate    = "WWW-Authenticate"
	// Caching - These header fields are related to caching policy.
	HeaderAge           = "Age"
	HeaderCacheControl  = "Cache-Control"
	HeaderClearSiteData = "Clear-Site-Data"
	HeaderExpires       = "Expires"
	HeaderPragma        = "Pragma"
	HeaderWarning       = "Warning"
	// Client hints - These header fields are used for content adaptation.
	HeaderAcceptCH         = "Accept-CH"
	HeaderAcceptCHLifetime = "Accept-CH-Lifetime"
	HeaderContentDPR       = "Content-DPR"
	HeaderDPR              = "DPR"
	HeaderEarlyData        = "Early-Data"
	HeaderSaveData         = "Save-Data"
	HeaderViewportWidth    = "Viewport-Width"
	HeaderWidth            = "Width"
	// Conditionals - These header fields are used for conditional requests.
	HeaderETag              = "ETag"
	HeaderIfMatch           = "If-Match"
	HeaderIfModifiedSince   = "If-Modified-Since"
	HeaderIfNoneMatch       = "If-None-Match"
	HeaderIfUnmodifiedSince = "If-Unmodified-Since"
	HeaderLastModified      = "Last-Modified"
	HeaderVary              = "Vary"
	// Connection management - These header fields are related to connection management.
	HeaderConnection      = "Connection"
	HeaderKeepAlive       = "Keep-Alive"
	HeaderProxyConnection = "Proxy-Connection"
	// Content negotiation - These header fields are used for content negotiation.
	HeaderAccept         = "Accept"
	HeaderAcceptCharset  = "Accept-Charset"
	HeaderAcceptEncoding = "Accept-Encoding"
	HeaderAcceptLanguage = "Accept-Language"
	// Controls - These header fields are related to general controls.
	HeaderCookie      = "Cookie"
	HeaderExpect      = "Expect"
	HeaderMaxForwards = "Max-Forwards"
	HeaderSetCookie   = "Set-Cookie"
	// CORS (Cross-Origin Resource Sharing) - These header fields are related to CORS.
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderOrigin                        = "Origin"
	HeaderTimingAllowOrigin             = "Timing-Allow-Origin"
	HeaderXPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"
	// Do Not Track - These header fields are related to user tracking preference.
	HeaderDNT = "DNT"
	HeaderTk  = "Tk"
	// Downloads - This header field is related to content disposition.
	HeaderContentDisposition = "Content-Disposition"
	// Message body information - These header fields are related to the message body.
	HeaderContentEncoding = "Content-Encoding"
	HeaderContentLanguage = "Content-Language"
	HeaderContentLength   = "Content-Length"
	HeaderContentLocation = "Content-Location"
	HeaderContentType     = "Content-Type"
	// Proxies - These header fields are related to proxy servers.
	HeaderForwarded       = "Forwarded"
	HeaderVia             = "Via"
	HeaderXForwardedFor   = "X-Forwarded-For"
	HeaderXForwardedHost  = "X-Forwarded-Host"
	HeaderXForwardedProto = "X-Forwarded-Proto"
	// Redirects - This header field is related to HTTP redirection.
	HeaderLocation = "Location"
	// Request context - These header fields are related to the request context.
	HeaderFrom           = "From"
	HeaderHost           = "Host"
	HeaderReferer        = "Referer"
	HeaderReferrerPolicy = "Referrer-Policy"
	HeaderUserAgent      = "User-Agent"
	// Response context - These header fields are related to the response context.
	HeaderAllow  = "Allow"
	HeaderServer = "Server"
	// Range requests - These header fields are related to partial requests and responses.
	HeaderAcceptRanges = "Accept-Ranges"
	HeaderContentRange = "Content-Range"
	HeaderIfRange      = "If-Range"
	HeaderRange        = "Range"
	// Security - These header fields are related to security.
	HeaderContentSecurityPolicy           = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	HeaderCrossOriginResourcePolicy       = "Cross-Origin-Resource-Policy"
	HeaderExpectCT                        = "Expect-CT"
	HeaderFeaturePolicy                   = "Feature-Policy"
	HeaderPublicKeyPins                   = "Public-Key-Pins"
	HeaderPublicKeyPinsReportOnly         = "Public-Key-Pins-Report-Only"
	HeaderStrictTransportSecurity         = "Strict-Transport-Security"
	HeaderUpgradeInsecureRequests         = "Upgrade-Insecure-Requests"
	HeaderXContentTypeOptions             = "X-Content-Type-Options"
	HeaderXDownloadOptions                = "X-Download-Options"
	HeaderXFrameOptions                   = "X-Frame-Options"
	HeaderXPoweredBy                      = "X-Powered-By"
	HeaderXXSSProtection                  = "X-XSS-Protection"
	// Server-sent event - These header fields are related to server-sent events.
	HeaderLastEventID = "Last-Event-ID"
	HeaderNEL         = "NEL"
	HeaderPingFrom    = "Ping-From"
	HeaderPingTo      = "Ping-To"
	HeaderReportTo    = "Report-To"
	// Transfer coding - These header fields are related to transfer coding.
	HeaderTE               = "TE"
	HeaderTrailer          = "Trailer"
	HeaderTransferEncoding = "Transfer-Encoding"
	// WebSockets -These header fields are related to the WebSocket protocol.
	HeaderSecWebSocketAccept     = "Sec-WebSocket-Accept"
	HeaderSecWebSocketExtensions = "Sec-WebSocket-Extensions"
	HeaderSecWebSocketKey        = "Sec-WebSocket-Key"
	HeaderSecWebSocketProtocol   = "Sec-WebSocket-Protocol"
	HeaderSecWebSocketVersion    = "Sec-WebSocket-Version"
	// Other -These header fields do not fall into the above categories.
	HeaderAcceptPatch         = "Accept-Patch"
	HeaderAcceptPushPolicy    = "Accept-Push-Policy"
	HeaderAcceptSignature     = "Accept-Signature"
	HeaderAltSvc              = "Alt-Svc"
	HeaderDate                = "Date"
	HeaderIndex               = "Index"
	HeaderLargeAllocation     = "Large-Allocation"
	HeaderLink                = "Link"
	HeaderPushPolicy          = "Push-Policy"
	HeaderRetryAfter          = "Retry-After"
	HeaderXRatelimitRemaining = "X-Ratelimit-Remaining"
	HeaderServerTiming        = "Server-Timing"
	HeaderSignature           = "Signature"
	HeaderSignedHeaders       = "Signed-Headers"
	HeaderSourceMap           = "SourceMap"
	HeaderUpgrade             = "Upgrade"
	HeaderXDNSPrefetchControl = "X-DNS-Prefetch-Control"
	HeaderXPingback           = "X-Pingback"
	HeaderXRequestedWith      = "X-Requested-With"
	HeaderXRobotsTag          = "X-Robots-Tag"
	HeaderXUACompatible       = "X-UA-Compatible"
)
