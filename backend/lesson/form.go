package lesson


type LessonSignup struct {
	LessSubject  string `json:"lesson_subject"`
	LessonNo     int    `json:"lesson_no"`
	LessonName   string `json:"lesson_name"`
	LessonDetail string `json:"lesson_detail"`
}

type LessonInfo struct {
	LessSubject  string `json:"lesson_subject"`
	LessonNo   int    `json:"lesson_no"`
	LessonName string `json:"lesson_name"`
	LessonDetail string    `json:"lesson_detail"`
}

type LessonRequest struct {
	LessSubject  string `json:"lesson_subject"`
	LessonNo     int       `json:"lesson_no"`
	LessonName   string    `json:"lesson_name"`
	LessonDetail string    `json:"lesson_detail"`
	CreateDate   string `json:"create_date"`
}