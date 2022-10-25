package main

import (
	"context"
	"fmt"
	"github.com/tupizz/go-foundation/42-creating-router-golang/handlers"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	Method  string
	Regex   *regexp.Regexp
	Handler http.HandlerFunc
}

func NewRouter(method, path string, handler http.HandlerFunc) Router {
	return Router{
		Method:  method,
		Regex:   regexp.MustCompile(fmt.Sprintf("^%s$", path)),
		Handler: handler,
	}
}

var routes []Router = []Router{
	NewRouter("GET", "/", handlers.Home),
	NewRouter("GET", "/contact", handlers.Contact),
	NewRouter("GET", "/api/widgets", handlers.ApiGetWidgets),
	NewRouter("POST", "/api/widgets", handlers.ApiCreateWidget),
	NewRouter("POST", "/api/widgets/([^/]+)", handlers.ApiUpdateWidget),
	NewRouter("POST", "/api/widgets/([^/]+)/parts", handlers.ApiCreateWidgetPart),
	NewRouter("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", handlers.ApiUpdateWidgetPart),
	NewRouter("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/delete", handlers.ApiDeleteWidgetPart),
	NewRouter("GET", "/([^/]+)", handlers.Widget),
	NewRouter("GET", "/([^/]+)/admin", handlers.WidgetAdmin),
	NewRouter("POST", "/([^/]+)/image", handlers.WidgetImage),
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.Regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}

			ctx := context.WithValue(r.Context(), handlers.CtxKey{}, matches[1:])
			route.Handler(w, r.WithContext(ctx))
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}

func main() {
	router := http.HandlerFunc(Serve)
	log.Fatal(http.ListenAndServe(":7070", router))
}
