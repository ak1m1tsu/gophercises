package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(data)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseJson(data)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func parseJson(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	if err := json.Unmarshal(data, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	if err := yaml.Unmarshal(data, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}
