// Package swaggerdocs serves a Swagger UI instance entirely from embedded
// assets, with no CDN or network calls, so the docs work fully offline.
package swaggerdocs

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var distFS embed.FS

// Handler returns an http.Handler serving the Swagger UI static assets.
// Mount it under a path prefix, e.g. router.Mount("/docs/", swaggerdocs.Handler("/docs/")).
func Handler(prefix string) http.Handler {
	sub, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic(err)
	}
	return http.StripPrefix(prefix, http.FileServer(http.FS(sub)))
}

// SpecHandler serves the raw OpenAPI spec bytes with the given content type.
func SpecHandler(spec []byte, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(spec)
	}
}
