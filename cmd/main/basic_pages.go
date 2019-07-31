package main

import (
	"bytes"
	"html/template"
	"net/http"
	"time"
)

const (
	// Do not work with hardcoded templates unless it is a must!
	indexPageTemplate = `<!doctype html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Welcome {{ .Guest }}</title>
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/foundation/6.5.3/css/foundation.min.css">
	</head>
	<body>
		<div class="grid-x">
			<div class="cell"><h1>Hello From Index</h1></div>
			<div class="cell">{{ .CurrentTime }}</div>
		</div>
	</body>
</html>`
)

type indexTemplateArgs struct {
	Guest       string
	CurrentTime string
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	t := template.New("index")
	currentTime := time.Now()
	params := indexTemplateArgs{
		Guest:       r.RemoteAddr,
		CurrentTime: currentTime.Local().Format("02-01-2006 15:04:05 PM MST"),
	}

	t, err := t.Parse(indexPageTemplate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	buff := make([]byte, 0)
	out := bytes.NewBuffer(buff)
	err = t.Execute(out, params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out.Bytes())
}
