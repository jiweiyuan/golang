package main

import (
	"github.com/jiweiyuan/golang/gopl/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const temp = `{{.TotalCount}} issues:
{{range .Items}}-----------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo }} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main(){
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temp))
	result, err := github.SearchIssues(github.TEST_PARAMS)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
