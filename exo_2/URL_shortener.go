package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type MyHandler struct{}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestedPath := r.URL.Path

		if destination, ok := pathsToUrls[requestedPath]; ok {
			http.Redirect(w, r, destination, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) { //Copied from the exercise no idea what this does
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(parsedYaml map[string]string) map[string]string {
	pathMap := make(map[string]string)
	for key, value := range parsedYaml {
		pathMap[key] = value
	}
	return pathMap
}

func parseYAML(yamlData []byte) (map[string]string, error) {
	var data map[string]string
	// Var_name map[]
	err := yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return nil, err
	}

	// DEBUG
	for short, long := range data {
		fmt.Printf("Short: %s, Long: %s\n", short, long)
	}

	return data, nil
}

func main() {
	yamlFile, err_i := os.ReadFile("yaml.yaml")
	if err_i != nil {
		panic(err_i)
	}
	fmt.Print(string(yamlFile))
	parsedYaml, err := parseYAML(yamlFile)
	if err != nil {
		panic(err)
	}
	pathMap := buildMap(parsedYaml)

	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	myHandler := MapHandler(pathMap, fallbackHandler)

	http.ListenAndServe(":8080", myHandler)
}
