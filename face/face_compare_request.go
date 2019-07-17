package face

// CompareRequest has request parameters for Compare API.
type CompareRequest struct {
	FaceToken1   string `url:"face_token1,omitempty"`
	ImageURL1    string `url:"image_url1,omitempty"`
	Image1Base64 string `url:"image_base64_1,omitempty"`

	FaceToken2   string `url:"face_token2,omitempty"`
	ImageURL2    string `url:"image_url2,omitempty"`
	Image2Base64 string `url:"image_base64_2,omitempty"`
}
