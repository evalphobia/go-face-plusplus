go-face-plusplus
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22]

[1]: https://godoc.org/github.com/evalphobia/go-face-plusplus?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-face-plusplus
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-face-plusplus.svg
[6]: https://github.com/evalphobia/go-face-plusplus/releases/latest
[7]: https://travis-ci.org/evalphobia/go-face-plusplus.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-face-plusplus
[9]: https://coveralls.io/repos/evalphobia/go-face-plusplus/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-face-plusplus?branch=master
[11]: https://codecov.io/github/evalphobia/go-face-plusplus/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-face-plusplus?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-face-plusplus
[14]: https://goreportcard.com/report/github.com/evalphobia/go-face-plusplus
[15]: https://img.shields.io/github/downloads/evalphobia/go-face-plusplus/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-face-plusplus/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-face-plusplus.svg
[18]: https://github.com/evalphobia/go-face-plusplus/stargazers
[19]: https://codeclimate.com/github/evalphobia/go-face-plusplus/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/go-face-plusplus
[21]: https://bettercodehub.com/edge/badge/evalphobia/go-face-plusplus?branch=master
[22]: https://bettercodehub.com/


Unofficial golang library for Face++.


# Current Supported API List


- Face
    - Detect
    - Compare
- Beautify
    - Merge Face

# Quick Usage

## Face

### Detect

```go
package main

import (
	"fmt"

	"github.com/evalphobia/go-face-plusplus/config"
	"github.com/evalphobia/go-face-plusplus/face"
)

func getFacesFromURL(url string) ([]face.Faces, error) {
	conf := config.Config{
		APIKey:    "",
		APISecret: "",
		Debug:     true,
	}

	svc, err := face.New(conf)
	if err != nil {
		return nil, err
	}

	resp, err:= svc.DetectByURL(
		url,
		face.WithReturnLandmark(face.ReturnLandmarkYES),
        face.WithReturnAttributes(face.GetAllReturnAttributes()...))
	if err != nil {
		return nil, err
    }

    fmt.Printf("Face Count: %d\n", resp.FaceNum)
    return resp.Faces, nil
}
```


### Compare

```go
package main

import (
	"fmt"

	"github.com/evalphobia/go-face-plusplus/config"
	"github.com/evalphobia/go-face-plusplus/face"
)

func compareFacesFromURL(url1, url2 string) ([]face.Faces, []face.Faces, error) {
	conf := config.Config{
		APIKey:    "",
		APISecret: "",
		Debug:     true,
	}

	svc, err := face.New(conf)
	if err != nil {
		return nil, err
	}

	resp, err:= svc.CompareByURL(url1, url2)
	if err != nil {
		return nil, err
    }

    fmt.Printf("Compare Confidence: %f\n", resp.Confidence)
    return resp.Faces1, resp.Faces2, nil
}
```

## Beautify

### Merge Face

```go
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

```


# Environment variables

| Name | Description |
|:--|:--|
| `FACEPP_API_KEY` | API Key of Face++. |
| `FACEPP_API_SECRET` | API Secret of Face++. |
