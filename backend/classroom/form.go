package classroom

type ClassroomSignup struct {
	ClassID      string    `json:"class_id"`
	ClassName   string    `json:"class_name"`
	ClassInstructor string `json:"class_instructor"`
	ClassDetail string `json:"class_detail"`
	ClassImage string `json:"class_image"`
}

type ClassroomInfo struct {
	ClassID      string    `json:"class_id"`
	ClassName   string    `json:"class_name"`
	ClassInstructor string `json:"class_instructor"`
	ClassDetail string `json:"class_detail"`
	ClassImage string `json:"class_image"`
}

