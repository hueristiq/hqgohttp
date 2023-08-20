package headers

const (
	// Authentication
	Authorization      = "Authorization"
	ProxyAuthenticate  = "Proxy-Authenticate"
	ProxyAuthorization = "Proxy-Authorization"
	WWWAuthenticate    = "WWW-Authenticate"

	// Caching
	Age           = "Age"
	CacheControl  = "Cache-Control"
	ClearSiteData = "Clear-Site-Data"
	Expires       = "Expires"
	Pragma        = "Pragma"
	Warning       = "Warning"

	// Client hints
	AcceptCH         = "Accept-CH"
	AcceptCHLifetime = "Accept-CH-Lifetime"
	ContentDPR       = "Content-DPR"
	DPR              = "DPR"
	EarlyData        = "Early-Data"
	SaveData         = "Save-Data"
	ViewportWidth    = "Viewport-Width"
	Width            = "Width"

	// Conditionals
	ETag              = "ETag"
	IfMatch           = "If-Match"
	IfModifiedSince   = "If-Modified-Since"
	IfNoneMatch       = "If-None-Match"
	IfUnmodifiedSince = "If-Unmodified-Since"
	LastModified      = "Last-Modified"
	Vary              = "Vary"

	// Connection management
	Connection      = "Connection"
	KeepAlive       = "Keep-Alive"
	ProxyConnection = "Proxy-Connection"

	// Content negotiation
	Accept         = "Accept"
	AcceptCharset  = "Accept-Charset"
	AcceptEncoding = "Accept-Encoding"
	AcceptLanguage = "Accept-Language"

	// Controls
	Cookie      = "Cookie"
	Expect      = "Expect"
	MaxForwards = "Max-Forwards"
	SetCookie   = "Set-Cookie"

	// CORS
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	AccessControlMaxAge           = "Access-Control-Max-Age"
	AccessControlRequestHeaders   = "Access-Control-Request-Headers"
	AccessControlRequestMethod    = "Access-Control-Request-Method"
	Origin                        = "Origin"
	TimingAllowOrigin             = "Timing-Allow-Origin"
	XPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

	// Do Not Track
	DNT = "DNT"
	Tk  = "Tk"

	// Downloads
	ContentDisposition = "Content-Disposition"

	// Message body information
	ContentEncoding = "Content-Encoding"
	ContentLanguage = "Content-Language"
	ContentLength   = "Content-Length"
	ContentLocation = "Content-Location"
	ContentType     = "Content-Type"

	// Proxies
	Forwarded       = "Forwarded"
	Via             = "Via"
	XForwardedFor   = "X-Forwarded-For"
	XForwardedHost  = "X-Forwarded-Host"
	XForwardedProto = "X-Forwarded-Proto"

	// Redirects
	Location = "Location"

	// Request context
	From           = "From"
	Host           = "Host"
	Referer        = "Referer"
	ReferrerPolicy = "Referrer-Policy"
	UserAgent      = "User-Agent"

	// Response context
	Allow  = "Allow"
	Server = "Server"

	// Range requests
	AcceptRanges = "Accept-Ranges"
	ContentRange = "Content-Range"
	IfRange      = "If-Range"
	Range        = "Range"

	// Security
	ContentSecurityPolicy           = "Content-Security-Policy"
	ContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	CrossOriginResourcePolicy       = "Cross-Origin-Resource-Policy"
	ExpectCT                        = "Expect-CT"
	FeaturePolicy                   = "Feature-Policy"
	PublicKeyPins                   = "Public-Key-Pins"
	PublicKeyPinsReportOnly         = "Public-Key-Pins-Report-Only"
	StrictTransportSecurity         = "Strict-Transport-Security"
	UpgradeInsecureRequests         = "Upgrade-Insecure-Requests"
	XContentTypeOptions             = "X-Content-Type-Options"
	XDownloadOptions                = "X-Download-Options"
	XFrameOptions                   = "X-Frame-Options"
	XPoweredBy                      = "X-Powered-By"
	XXSSProtection                  = "X-XSS-Protection"

	// Server-sent event
	LastEventID = "Last-Event-ID"
	NEL         = "NEL"
	PingFrom    = "Ping-From"
	PingTo      = "Ping-To"
	ReportTo    = "Report-To"

	// Transfer coding
	TE               = "TE"
	Trailer          = "Trailer"
	TransferEncoding = "Transfer-Encoding"

	// WebSockets
	SecWebSocketAccept     = "Sec-WebSocket-Accept"
	SecWebSocketExtensions = "Sec-WebSocket-Extensions"
	SecWebSocketKey        = "Sec-WebSocket-Key"
	SecWebSocketProtocol   = "Sec-WebSocket-Protocol"
	SecWebSocketVersion    = "Sec-WebSocket-Version"

	// Other
	AcceptPatch         = "Accept-Patch"
	AcceptPushPolicy    = "Accept-Push-Policy"
	AcceptSignature     = "Accept-Signature"
	AltSvc              = "Alt-Svc"
	Date                = "Date"
	Index               = "Index"
	LargeAllocation     = "Large-Allocation"
	Link                = "Link"
	PushPolicy          = "Push-Policy"
	RetryAfter          = "Retry-After"
	ServerTiming        = "Server-Timing"
	Signature           = "Signature"
	SignedHeaders       = "Signed-Headers"
	SourceMap           = "SourceMap"
	Upgrade             = "Upgrade"
	XDNSPrefetchControl = "X-DNS-Prefetch-Control"
	XPingback           = "X-Pingback"
	XRequestedWith      = "X-Requested-With"
	XRobotsTag          = "X-Robots-Tag"
	XUACompatible       = "X-UA-Compatible"
)
