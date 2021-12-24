package blog

import "time"

type BlogCreate struct {
	LessonName string `json:"lesson_name"`
	BlogDetail     string    `json:"blog_detail"`
	BlogCreateData time.Time `json:"create_date"`
}

type BlogUpdate struct {
	LessonNo     int       `json:"lesson_no"`
	LessonName string `json:"lesson_name"`
	
}