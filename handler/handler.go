package handler

import (
	"assignment-3/model"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

var wstat model.WeatherStatus

const htmlPath = "file/index.html"
const jsonPath = "file/status.json"

func MainHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	// read from json file and write to webData
	file, _ := ioutil.ReadFile(jsonPath)
	json.Unmarshal(file, &wstat)
	templates, _ := template.ParseFiles(htmlPath)
	context := model.CompiledWeatherStatus{
		Status: model.Status{
			Water: wstat.Status.Water,
			Wind:  wstat.Status.Wind,
		},
		StatusCompiled: wstat.Status.CheckStatus(),
	}
	templates.Execute(w, context)
}
