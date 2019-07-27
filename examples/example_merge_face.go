package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/evalphobia/go-face-plusplus/beautify"
	"github.com/evalphobia/go-face-plusplus/config"
)

func main() {
	var templateURL string
	var mergeURL string
	flag.StringVar(&templateURL, "url1", "", "set temaplte image url")
	flag.StringVar(&mergeURL, "url2", "", "set merge image url")
	flag.Parse()

	conf := config.Config{
		APIKey:    "",
		APISecret: "",
		// Debug:     true,  // big base64 response data  will show up in stdout.
	}

	svc, err := beautify.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.MergeFace(beautify.MergeFaceRequest{
		TemplateURL: templateURL,
		MergeURL:    mergeURL,
	})
	if err != nil {
		panic(err)
	}

	img, err := resp.GetResultImage()
	if err != nil {
		panic(err)
	}

	fileName := "merge_face_example.jpg"
	if err := ioutil.WriteFile(fileName, img, 0644); err != nil {
		panic(err)
	}
	fmt.Printf("Success merge face! see: %s\n", fileName)
}
