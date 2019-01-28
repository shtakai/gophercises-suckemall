package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	// 0) connect key - val on pathsToUrls
	// 1) dig val by key
	// 2-1) if key exists
	//			redirect to val
	// 2-2) if key doesn't exist
	//			fallback
	_ = pathsToUrls
	mux := http.NewServeMux()
	for _, key := range pathsToUrls {
		redirectHandler := http.RedirectHandler(pathsToUrls[key], http.StatusTemporaryRedirect)
		mux.Handle(key, redirectHandler)
	}
	mux.Handle("/*", fallback)

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		val := pathsToUrls[path]
		if val != nil {

		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
