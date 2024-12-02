package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:generate bash -c "mkdir -p codegen"
//go:generate go run generate.go
//go:generate go fmt codegen/user.gen.go

func main() {
	var resp map[string]interface{}
	in, _ := os.Open("../user.json")
	b, _ := io.ReadAll(in)
	json.Unmarshal(b, &resp)

	caser := cases.Title(language.English)

	data := struct {
		Name   string
		Fields map[string]interface{}
	}{
		"User",
		resp,
	}

	tpl, err := template.New("template.tpl").Funcs(template.FuncMap{
		"Title": caser.String,
		"TypeOf": func(v interface{}) string {
			if v == nil {
				return "string"
			}
			return strings.ToLower(reflect.TypeOf(v).String())
		},
	}).ParseFiles("template.tpl")

	if err != nil {
		fmt.Errorf("failed to parse template: %v", err)
	}

	out, _ := os.Create("codegen/user.gen.go")
	defer out.Close()

	if err = tpl.Execute(out, data); err != nil {
		fmt.Errorf("failed to execute template: %v", err)
	}
}
