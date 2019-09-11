package urlshort

import (
	"fmt"
	yamlParsing "gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string) func(string) (string, bool) {
	return func(path string) (string, bool) {
		url, exists := pathsToUrls[path]
		fmt.Println(pathsToUrls[path])
		return url, exists
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
func YAMLHandler(yaml []byte) (func(string) (string, bool), error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap), nil
}

func parseYAML(yaml []byte) (parsedYaml []map[string]string, err error) {
	err = yamlParsing.Unmarshal(yaml, &parsedYaml)
	return parsedYaml, err
}

func buildMap(parsedYaml []map[string]string) map[string]string {
	mapping := make(map[string]string)
	for _, entry := range parsedYaml {
		key := entry["path"]
		mapping[key] = entry["url"]
	}
	return mapping
}

func HttpHandler(mapHandler func(string) (string, bool), fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, exists := mapHandler(r.URL.Path); exists {
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}