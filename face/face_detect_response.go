package face

import (
	"github.com/evalphobia/go-face-plusplus/client"
)

// DetectResponse has response from Detect API.
type DetectResponse struct {
	client.BaseResponse
	ImageID string `json:"image_id"`
	Faces   []Face `json:"faces"`
	FaceNum int    `json:"face_num"`
}

type Face struct {
	FaceToken     string              `json:"face_token"`
	Landmark      map[string]Landmark `json:"landmark"`
	FaceRectangle `json:"face_rectangle"`
	Attributes    `json:"attributes"`
}

type FaceRectangle struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ref: https://console.faceplusplus.com/documents/6329308
type Landmark struct {
	X int `json:"x"`
	Y int `json:"y"`
}
