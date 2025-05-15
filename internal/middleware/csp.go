package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type key string

const NonceKey key = "nonces"

type Nonces struct {
	InlineScript string
	InlineStyle  string
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func GetNonces(ctx context.Context) Nonces {
	nonceSet := ctx.Value(NonceKey)
	if nonceSet == nil {
		log.Fatal("error getting nonce set - is nil")
	}

	nonces, ok := nonceSet.(Nonces)
	if !ok {
		log.Fatal("error getting nonce set - not ok")
	}

	return nonces
}

func GetInlineScriptNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)

	return nonceSet.InlineScript
}

func GetInlineStyleNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)

	return nonceSet.InlineStyle
}

func CSPMiddleware(
	h http.Handler,
	htmxHash string,
	hyperscriptHash string,
	echartsHash string,
) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nonceSet := Nonces{
			InlineScript: generateRandomString(16),
			InlineStyle:  generateRandomString(16),
		}

		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)

		// create a csp with hashes and the inline script and style nonce
		csp := []string{
			"default-src 'self'",
			"base-uri 'none'",
			// "require-trusted-types-for 'script'",
			fmt.Sprintf(
				"script-src '%s' '%s' '%s' 'nonce-%s' 'unsafe-inline'",
				htmxHash,
				hyperscriptHash,
				echartsHash,
				nonceSet.InlineScript,
			),
			fmt.Sprintf(
				"style-src 'self' 'nonce-%s' 'unsafe-inline'",
				nonceSet.InlineStyle,
			),
			"img-src 'self' data:",
		}

		w.Header().Set("Content-Security-Policy", strings.Join(csp, "; "))
		// w.Header().Set("Content-Security-PolicyReport-Only", strings.Join(csp, "; "))

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
