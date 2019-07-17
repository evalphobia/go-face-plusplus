package face

// Attributes contains response of return_attributes.
type Attributes struct {
	Gender       StringValue    `json:"gender"`
	Age          IntValue       `json:"age"`
	Smile        ThresholdValue `json:"smile"`
	Glass        StringValue    `json:"glass"`
	HeadPose     `json:"headpose"`
	Blur         `json:"blur"`
	EyeStatus    `json:"eyestatus"`
	Emotion      `json:"emotion"`
	FaceQueality ThresholdValue `json:"facequality"`
	Ethnicity    StringValue    `json:"ethnicity"`
	MouthStatus  `json:"mouthstatus"`
	EyeGaze      `json:"eyegaze"`
	SkinStatus   `json:"skinstatus"`
	Beauty       `json:"beauty"`
}

type StringValue struct {
	Value string `json:"value"`
}

type IntValue struct {
	Value int `json:"value"`
}

type ThresholdValue struct {
	Threshold float64 `json:"threshold"`
	Value     float64 `json:"value"`
}

type HeadPose struct {
	YawAngle   float64 `json:"yaw_angle"`
	PitchAngle float64 `json:"pitch_angle"`
	RollAngle  float64 `json:"roll_angle"`
}

type Blur struct {
	Blurness     ThresholdValue `json:"blurness"`
	MotionBlur   ThresholdValue `json:"motionblur"`
	GaussianBlur ThresholdValue `json:"gaussianblur"`
}

type EyeStatus struct {
	LeftEyeStatus  eyeStatus `json:"left_eye_status"`
	RightEyeStatus eyeStatus `json:"right_eye_status"`
}

type eyeStatus struct {
	NormalGlassEyeOpen  float64 `json:"normal_glass_eye_open"`
	NormalGlassEyeClose float64 `json:"normal_glass_eye_close"`
	NoGlassEyeOpen      float64 `json:"no_glass_eye_open"`
	NoGlassEyeClose     float64 `json:"no_glass_eye_close"`
	Occlusion           float64 `json:"occlusion"`
	DarkGlasses         float64 `json:"dark_glasses"`
}

type Emotion struct {
	Anger     float64 `json:"anger"`
	Disgust   float64 `json:"disgust"`
	Fear      float64 `json:"fear"`
	Happiness float64 `json:"happiness"`
	Neutral   float64 `json:"neutral"`
	Sadness   float64 `json:"sadness"`
	Surprise  float64 `json:"surprise"`
}

type MouthStatus struct {
	Open                     float64 `json:"open"`
	Close                    float64 `json:"close"`
	SurgicalMaskOrRespirator float64 `json:"surgical_mask_or_respirator"`
	OtherOcclusion           float64 `json:"other_occlusion"`
}

type EyeGaze struct {
	LeftEyeGaze  eyeGaze `json:"left_eye_gaze"`
	RightEyeGaze eyeGaze `json:"right_eye_gaze"`
}

type eyeGaze struct {
	PositionXCoordinate float64 `json:"position_x_coordinate"`
	PositionYCoordinate float64 `json:"position_y_coordinate"`
	VectorXComponent    float64 `json:"vector_x_component"`
	VectorYComponent    float64 `json:"vector_y_component"`
	VectorZComponent    float64 `json:"vector_z_component"`
}

type SkinStatus struct {
	Acne       float64 `json:"acne"`
	Health     float64 `json:"health"`
	Stain      float64 `json:"stain"`
	DarkCircle float64 `json:"dark_circle"`
}

type Beauty struct {
	FemaleScore float64 `json:"female_score"`
	MaleScore   float64 `json:"male_score"`
}
