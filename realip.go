package realip

import (
	"context"
	"net/http"

	tomasenRealip "github.com/tomasen/realip"
)

// Key to use when setting the Real IP.
type ctxKeyRealIP int

// RealIPKey is the key that holds the unique IP in a request context.
const RealIPKey ctxKeyRealIP = 0

// RealIP is a middleware that sets a http.Request's RemoteAddr to the results
// of parsing either the X-Forwarded-For header or the X-Real-IP header.
// It also injects the IP into the context of each request.
func Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		realIP := tomasenRealip.FromRequest(r)

		if realIP != "" {
			ctx = context.WithValue(ctx, RealIPKey, realIP)
			r.RemoteAddr = realIP
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// GetRealIP returns a IP from the given context if one is present.
// Returns the empty string if a IP cannot be found.
func GetRealIP(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RealIPKey).(string); ok {
		return reqID
	}
	return ""
}
