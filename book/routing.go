package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	aloha := createAnyUser("user")

	if num, ok := aloha.(Person); ok {
		fmt.Println(num)
	} else {
		fmt.Println(num)
	}
}

func createAnyUser(role string) interface{} {
	if role == "admin" {
		return 22
	}
	return struct {
		Name string
		Age  int
	}{
		Name: "Helen",
		Age:  25,
	}
}

//func newPathResolver() *pathResolver {
//	return &pathResolver{make(map[string]http.HandlerFunc)}
//}
//
//type pathResolver struct {
//	handlers map[string]http.HandlerFunc
//}
//
//func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
//	p.handlers[path] = handler
//}
//
//func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	check := req.Method + " " + req.URL.Path
//
//	for pattern, handlerFunc := range p.handlers {
//
//		if ok, err := path.Match(pattern, check); ok && err == nil {
//			handlerFunc(res, req)
//
//			return
//
//		} else if err != nil {
//
//			fmt.Fprint(res, err)
//
//		}
//	}
//}
//
//func hello(res http.ResponseWriter, req *http.Request) {
//	query := req.URL.Query()
//	name := query.Get("name")
//	if name == "" {
//		name = "Inigo Montoya"
//	}
//	fmt.Fprint(res, "Hello, my name is ", name)
//}
//
//func goodbye(res http.ResponseWriter, req *http.Request) {
//	path := req.URL.Path
//	parts := strings.Split(path, "/")
//	name := parts[2]
//	if name == "" {
//		name = "Inigo Montoya"
//	}
//	fmt.Fprint(res, "Goodbye ", name)
//}
