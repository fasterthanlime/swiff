package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RangelReale/swfinfo"

	"encoding/json"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("swiff", "Your friendly neighborhood swf sniffer")
	urlParam = app.Arg("url", "The url of the swf file to sniff").Required().String()
)

type SniffResult struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
	Info    *SWFInfo `json:"info"`
}

type SWFInfo struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	result := &SniffResult{}

	info, err := operate(*urlParam)
	if err != nil {
		result.Success = false
		result.Errors = []string{err.Error()}
	} else {
		result.Success = true
		result.Errors = []string{}
		result.Info = info
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(jsonBytes)
}

func operate(url string) (*SWFInfo, error) {
	swf := &swfinfo.SWF{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Server responded with HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	err = swf.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal("Invalid command-line arguments", err)
	}

	return &SWFInfo{
		Width:  int(swf.FrameSize.Xmax.Pixels() - swf.FrameSize.Xmin.Pixels()),
		Height: int(swf.FrameSize.Ymax.Pixels() - swf.FrameSize.Ymin.Pixels()),
	}, nil
}
