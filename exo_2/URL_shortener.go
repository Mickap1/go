package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}

// func buildMap(parsedYaml map[string]string) map[string]string { //Generated by copilot no idea what this does
// 	return nil
// }

// func parseYAML(yaml []byte) (map[string]string, error) { //Generated by copilot no idea what this does
// 	return nil, nil
// }

// func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) { //Copied from the exercise no idea what this does
// 	parsedYaml, err := parseYAML(yaml)
// 	if err != nil {
// 	  return nil, err
// 	}
// 	pathMap := buildMap(parsedYaml)
// 	return MapHandler(pathMap, fallback), nil
//   }

// func main() { //Online example of how to use mux
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", HomeHandler)
// 	r.HandleFunc("/products", ProductsHandler)
// 	r.HandleFunc("/articles", ArticlesHandler)
// 	http.Handle("/", r)
// }

// r := mux.NewRouter() New example of how to use mux with regular expressions
// r.HandleFunc("/products/{key}", ProductHandler)
// r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
