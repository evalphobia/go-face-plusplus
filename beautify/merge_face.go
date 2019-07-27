package beautify

// MergeFace executes Merge Face API.
func (s *BeautifyService) MergeFace(req MergeFaceRequest) (*MergeFaceResponse, error) {
	resp := MergeFaceResponse{}
	err := s.client.CallPOST("/imagepp/v1/mergeface", req.ToParam(), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
