package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// 0) connect key - val on pathsToUrls
	// 1) dig val by key
	// 2-1) if key exists
	//			redirect to val
	// 2-2) if key doesn't exist
	//			fallback

	// returning is http.HandlerFunc.
	//  => just func(w http.ResponseWriter, r *http.Request)
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// fetching url and result(ok) by path
		url, ok := pathsToUrls[path]
		if ok {
			// use http.Redirect (code: 307)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		} else {
			// use fallback's serveHttp
			fallback.ServeHTTP(w, r)
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
	// parse yaml
	// if error is happening, return err
	// after same => maphandler
	//  return mapHandler

	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	var pathsToUrls = make(map[string]string)
	for _, path := range pathUrls {
		pathsToUrls[path.Path] = path.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml: path`
	Url  string `yaml: url`
}
