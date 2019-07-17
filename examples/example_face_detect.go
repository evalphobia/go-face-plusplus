package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/go-face-plusplus/config"
	"github.com/evalphobia/go-face-plusplus/face"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "set image url")
	flag.Parse()

	conf := config.Config{
		APIKey:    "",
		APISecret: "",
		Debug:     true,
	}

	svc, err := face.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.DetectByURL(
		url,
		face.WithReturnLandmark(face.ReturnLandmarkYES),
		face.WithReturnAttributes(face.GetAllReturnAttributes()...))
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%+v]\n", resp)
}
