package face

// CompareByBase64 executes Compare API by given base64 image data.
func (s *FaceService) CompareByBase64(base64Str1, base64Str2 string) (*CompareResponse, error) {
	return s.Compare(CompareRequest{
		Image1Base64: base64Str1,
		Image2Base64: base64Str2,
	})
}

// CompareByURL executes Compare API by given url.
func (s *FaceService) CompareByURL(url1, url2 string) (*CompareResponse, error) {
	return s.Compare(CompareRequest{
		ImageURL1: url1,
		ImageURL2: url2,
	})
}

// Compare executes Compare API.
func (s *FaceService) Compare(req CompareRequest) (*CompareResponse, error) {
	resp := CompareResponse{}
	err := s.client.CallPOST("/facepp/v3/compare", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
