package face

import (
	"encoding/base64"
	"io/ioutil"
)

// DetectByBase64 executes Detect API by given base64 image data.
func (s *FaceService) DetectByBase64(base64Str string, opts ...DetectRequestOption) (*DetectResponse, error) {
	return s.Detect(DetectRequest{
		ImageBase64: base64Str,
	}, opts...)
}

// DetectFromFile executes Detect API from local file (with base64).
func (s *FaceService) DetectFromFile(path string, opts ...DetectRequestOption) (*DetectResponse, error) {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return s.DetectByBase64(base64.StdEncoding.EncodeToString(byt), opts...)
}

// DetectByURL executes Detect API by given url.
func (s *FaceService) DetectByURL(url string, opts ...DetectRequestOption) (*DetectResponse, error) {
	return s.Detect(DetectRequest{
		ImageURL: url,
	}, opts...)
}

// Detect executes Detect API.
func (s *FaceService) Detect(req DetectRequest, opts ...DetectRequestOption) (*DetectResponse, error) {
	for _, opt := range opts {
		req = opt.Apply(req)
	}
	resp := DetectResponse{}
	err := s.client.CallPOST("/facepp/v3/detect", req.ToParam(), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
