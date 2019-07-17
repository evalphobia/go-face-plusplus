package face

import "strings"

// DetectRequest has request parameters for Detect API.
type DetectRequest struct {
	ImageURL         string             `url:"image_url,omitempty"`
	ImageBase64      string             `url:"image_base64,omitempty"`
	ReturnLandmark   ReturnLandmarkType `url:"return_landmark,omitempty"`
	ReturnAttributes ReturnAttributes   `url:"-"`
	CalculateAll     bool               `url:"-"`
	FaceRectangle    string             `url:"face_rectangle,omitempty"`
}

type detectRequest struct {
	DetectRequest    `url:",squash"`
	ReturnAttributes string `url:"return_attributes,omitempty"`
	CalculateAll     int    `url:"calculate_all,omitempty"`
}

// ToParam converts and casts parameters to suitable format.
func (d DetectRequest) ToParam() detectRequest {
	calculateAll := 0
	if d.CalculateAll {
		calculateAll = 1
	}

	return detectRequest{
		DetectRequest:    d,
		ReturnAttributes: d.ReturnAttributes.Join(),
		CalculateAll:     calculateAll,
	}
}

// ReturnLandmarkType is a parameter for ReturnLandmark.
type ReturnLandmarkType int

const (
	ReturnLandmarkNO      = 0
	ReturnLandmarkYES     = 1 // 83-point landmarks
	ReturnLandmarkYES106p = 2 // 106-point landmarks
)

// ReturnAttributes containes multiple ReturnAttribute.
type ReturnAttributes []ReturnAttribute

func (a ReturnAttributes) Join() string {
	list := make([]string, len(a))
	for i, v := range a {
		list[i] = string(v)
	}
	return strings.Join(list, ",")
}

// ReturnAttribute is single parameter for return_attributes.
type ReturnAttribute string

const (
	AttributeGender      = "gender"
	AttributeAge         = "age"
	AttributeSmiling     = "smiling"
	AttributeHeadPose    = "headpose"
	AttributeFaceQuality = "facequality"
	AttributeBlur        = "blur"
	AttributeEyeStatus   = "eyestatus"
	AttributeEmotion     = "emotion"
	AttributeEthnicity   = "ethnicity"
	AttributeBeauty      = "beauty"
	AttributeMouthStatus = "mouthstatus"
	AttributeEyeGaze     = "eyegaze"
	AttributeSkinStatus  = "skinstatus"
)

// GetAllReturnAttributes returns all of return_attributes.
func GetAllReturnAttributes() []ReturnAttribute {
	return []ReturnAttribute{
		AttributeGender,
		AttributeAge,
		AttributeSmiling,
		AttributeHeadPose,
		AttributeFaceQuality,
		AttributeBlur,
		AttributeEyeStatus,
		AttributeEmotion,
		AttributeEthnicity,
		AttributeBeauty,
		AttributeMouthStatus,
		AttributeEyeGaze,
		AttributeSkinStatus,
	}
}
