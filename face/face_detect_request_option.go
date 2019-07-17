package face

// DetectRequestOption is used for adding optional request parameters.
type DetectRequestOption interface {
	Apply(DetectRequest) DetectRequest
}

// WithReturnLandmark sets ReturnLandmark.
func WithReturnLandmark(landmarkType ReturnLandmarkType) DetectRequestOption {
	return withReturnLandmark{landmarkType}
}

type withReturnLandmark struct{ typ ReturnLandmarkType }

func (w withReturnLandmark) Apply(r DetectRequest) DetectRequest {
	r.ReturnLandmark = w.typ
	return r
}

// WithReturnAttributes sets ReturnAttributes.
func WithReturnAttributes(attrs ...ReturnAttribute) DetectRequestOption {
	return withReturnAttributes{attrs}
}

type withReturnAttributes struct{ attrs []ReturnAttribute }

func (w withReturnAttributes) Apply(r DetectRequest) DetectRequest {
	r.ReturnAttributes = ReturnAttributes(w.attrs)
	return r
}
