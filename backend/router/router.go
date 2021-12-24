package router

import (
	"backend/blog"
	"backend/classroom"
	"backend/lesson"
	"backend/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func Router() *fiber.App {

	app := fiber.New()
	app.Use(cors.New())
	userController := user.UserController{}
	classroomController := classroom.ClassroomController{}
	lessonController := lesson.LessonController{}
	blogController := blog.BlogController{}

	//User
	app.Get("/user", userController.CheckStatus)
	app.Post("/auth/sign-up", userController.UserSignup)
	app.Post("/auth/login", userController.Login)

	//Classroom
	app.Post("/classroom/sign-up", classroomController.ClassroomSignup)
	app.Get("/classroom/getallclassroom", classroomController.GetAllClassroom)
	app.Post("/classroom/updateclassroom", classroomController.UpdateClassroom)
	app.Delete("/classroom/deleteclassroom", classroomController.DeleteLesson)

	//Lesson
	app.Post("/lesson/sign-up", lessonController.SignupLesson)
	app.Get("/lesson/getalllesson", lessonController.GetAllLesson)
	app.Get("/lesson/getlessonbyclassid", lessonController.GetLessonbyClassID)
	app.Delete("/lesson/deletelesson", lessonController.DeleteLesson)
	app.Delete("/lesson/deleteFile",lessonController.DeleteFile)

	//Blog
	app.Post("/blog/sign-up", blogController.SignupBlog)
	app.Get("/blog/getblogbylessonno", blogController.GetBlogByLessonNo)
	app.Get("/blog/getblogbyclasseoom", blogController.GetBlogByClassroom)
	app.Post("/blog/uploadFile", blogController.Uploadfile)
	app.Post("/blog/updateblog",blogController.UpdateBlog)


	return app

}

