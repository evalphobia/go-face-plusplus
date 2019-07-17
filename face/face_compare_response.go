package face

import (
	"github.com/evalphobia/go-face-plusplus/client"
)

// CompareResponse has response from Compare API.
type CompareResponse struct {
	client.BaseResponse
	ImageID1 string `json:"image_id1"`
	Faces1   []Face `json:"faces1"`

	ImageID2 string `json:"image_id2"`
	Faces2   []Face `json:"faces2"`

	Confidence float64 `json:"confidence"`
	Thresholds `json:"threshold"`
}

type Thresholds struct {
	E3 float64 `json:"1e-3"`
	E4 float64 `json:"1e-4"`
	E5 float64 `json:"1e-5"`
}
