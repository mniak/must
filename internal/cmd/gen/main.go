package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

//go:embed *.tmpl
var embedFS embed.FS

const BaseName = "Do"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		fatal(1, "Missing argument MAX")
	}

	max, err := strconv.Atoi(os.Args[1])
	if err != nil || max < 0 {
		printUsage()
		fatal(2, "The value of MAX (%s) is not a valid positive integer", os.Args[1])
	}

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"IntRange": func(min, max int) []int {
				result := make([]int, 0)
				for i := min; i <= max; i++ {
					result = append(result, i)
				}
				return result
			},
			"FuncName": func(i int) string {
				return fmt.Sprintf("%s%d", BaseName, i)
			},
		}).
		ParseFS(embedFS, "*.tmpl")
	handle(err)

	fmust, err := os.Create("must.go")
	handle(err)
	defer fmust.Close()

	data := map[string]any{
		"Max":      max,
		"BaseName": BaseName,
	}
	handle(tmpl.ExecuteTemplate(fmust, "must.go.tmpl", data))

	fmusttest, err := os.Create("must_test.go")
	handle(err)
	defer fmust.Close()
	handle(tmpl.ExecuteTemplate(fmusttest, "must_test.go.tmpl", data))
}
