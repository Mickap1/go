package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

//Create an interface with short and long in url

type UrlsWrapper struct {
	Urls []Url `yaml:"urls"`
}

type Url struct {
	Short string
	Long  string
}

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

func buildMap(urls UrlsWrapper) map[string]string {
	pathMap := make(map[string]string)
	for key, value := range urls.Urls {
		pathMap[value.Short] = value.Long
		// fmt.Println(key, value)
	}
	return pathMap
}

func parseYAML(yamlData []byte) (UrlsWrapper, error) {
	//List of url structs
	var urls UrlsWrapper

	// fmt.Print(string(yamlData))
	err := yaml.Unmarshal(yamlData, &urls)
	if err != nil {
		fmt.Print("An error occured")
		return urls, err
	}
	// //Display the content of urls DEBUG
	// for i := 0; i < len(urls.Urls); i++ {
	// 	fmt.Println(urls.Urls[i].Short, urls.Urls[i].Long)
	// }
	return urls, nil
}

func main() {
	yamlFile, err_i := os.ReadFile("yaml.yaml")
	if err_i != nil {
		panic(err_i)
	}
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
