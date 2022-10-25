package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

type CtxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(CtxKey{}).([]string)
	return fields[index]
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home\n")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "contact\n")
}

func ApiGetWidgets(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiGetWidgets\n")
}

func ApiCreateWidget(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiCreateWidget\n")
}

func ApiUpdateWidget(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	fmt.Fprintf(w, "apiUpdateWidget %s\n", slug)
}

func ApiCreateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	fmt.Fprintf(w, "apiCreateWidgetPart %s\n", slug)
}

func ApiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	id, _ := strconv.Atoi(getField(r, 1))
	fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}

func ApiDeleteWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	id, _ := strconv.Atoi(getField(r, 1))
	fmt.Fprintf(w, "apiDeleteWidgetPart %s %d\n", slug, id)
}

func Widget(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	fmt.Fprintf(w, "widget %s\n", slug)
}

func WidgetAdmin(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	fmt.Fprintf(w, "widgetAdmin %s\n", slug)
}

func WidgetImage(w http.ResponseWriter, r *http.Request) {
	slug := getField(r, 0)
	fmt.Fprintf(w, "widgetImage %s\n", slug)
}
