package beautify

import (
	"encoding/base64"
	"fmt"

	"github.com/evalphobia/go-face-plusplus/client"
)

// MergeFaceRequest has request parameters for Merge Face API.
type MergeFaceRequest struct {
	TemplateURL    string `url:"template_url,omitempty"`
	TemplateBase64 string `url:"template_base64,omitempty"`
	// TemplateRectangle is the bounding box of the face within the template image.
	// Four integers, segmented by commas, "<Top>,<Left>,<Width>,<Height>"
	// e.g. "70,80,100,100".
	TemplateRectangle string `url:"-"`

	TemplateRectangleX      int `url:"-"` // left
	TemplateRectangleY      int `url:"-"` // top
	TemplateRectangleHeight int `url:"-"`
	TemplateRectangleWidth  int `url:"-"`

	MergeURL       string `url:"merge_url,omitempty"`
	MergeBase64    string `url:"merge_base64,omitempty"`
	MergeRectangle string `url:"-"`

	MergeRectangleX      int `url:"-"` // left
	MergeRectangleY      int `url:"-"` // top
	MergeRectangleHeight int `url:"-"`
	MergeRectangleWidth  int `url:"-"`

	MergeRate int `url:"merge_rate,omitempty"`
}

func (r MergeFaceRequest) HasTemplateRectangleInt() bool {
	return r.TemplateRectangleHeight != 0 && r.TemplateRectangleWidth != 0
}

func (r MergeFaceRequest) HasMergeRectangleInt() bool {
	return r.MergeRectangleHeight != 0 && r.MergeRectangleWidth != 0
}

// ToParam converts and casts parameters to suitable format.
func (r MergeFaceRequest) ToParam() mergeFaceRequest {
	tRect := r.TemplateRectangle
	if tRect == "" && r.HasTemplateRectangleInt() {
		tRect = fmt.Sprintf("%d,%d,%d,%d", r.TemplateRectangleY, r.TemplateRectangleX, r.TemplateRectangleWidth, r.TemplateRectangleHeight)
	}
	mRect := r.MergeRectangle
	if mRect == "" && r.HasMergeRectangleInt() {
		mRect = fmt.Sprintf("%d,%d,%d,%d", r.MergeRectangleY, r.MergeRectangleX, r.MergeRectangleWidth, r.MergeRectangleHeight)
	}

	return mergeFaceRequest{
		MergeFaceRequest:  r,
		TemplateRectangle: tRect,
		MergeRectangle:    mRect,
	}
}

type mergeFaceRequest struct {
	MergeFaceRequest  `url:",squash"`
	TemplateRectangle string `url:"template_rectangle,omitempty"`
	MergeRectangle    string `url:"merge_rectangle,omitempty"`
}

// MergeFaceResponse has response from Merge Face API.
type MergeFaceResponse struct {
	client.BaseResponse
	Result string `json:"result"` // base64 image
}

func (r MergeFaceResponse) GetResultImage() ([]byte, error) {
	return base64.StdEncoding.DecodeString(r.Result)
}
