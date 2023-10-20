package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

type RaceResultConfig struct {
	Key      string         `json:"key"`
	Contests map[int]string `json:"contests"`
	Server   string         `json:"server"`
}

type RaceResultList struct {
	Data [][]interface{} `json:"data"`
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("Please pass a race id as argument\n")
		os.Exit(1)
	}
	raceId := os.Args[1]

	configJsonBytes := fetchJsonAsString(fmt.Sprintf("https://my.raceresult.com/%s/RRPublish/data/config?page=results&noVisitor=1", raceId))
	var raceResultConfig RaceResultConfig
	unmarshalJson(configJsonBytes, &raceResultConfig)

	listJsonBytes := fetchJsonAsString(fmt.Sprintf("https://%s/%s/RRPublish/data/list?key=%s&listname=Ergebnislisten%%7CErgebnisliste%%20MW&page=results&contest=1&r=group&name=%%231_10%%20km%%0C%%231_M%%C3%%A4nnlich&f=%%0C", raceResultConfig.Server, raceId, raceResultConfig.Key))
	var raceResultList RaceResultList
	unmarshalJson(listJsonBytes, &raceResultList)

	for _, starter := range raceResultList.Data {
		line := ""
		for _, detail := range starter {
			field := ""
			if reflect.TypeOf(detail).String() == "string" {
				field = detail.(string)
			}
			line = line + field + ";"
		}
		fmt.Println(line)
	}
}

func fetchJsonAsString(url string) string {
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		fmt.Printf("Error during http request: %s\n", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	return string(body)
}

func unmarshalJson(jsonString string, target any) {
	err := json.Unmarshal([]byte(jsonString), &target)
	if err != nil {
		fmt.Printf("Error during unmarshalling json: %s\n", err.Error())
		os.Exit(1)
	}
}
