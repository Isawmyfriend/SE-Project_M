package blog

import (
	"backend/classroom"
	"backend/lesson"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type BlogController struct{}

func (bc BlogController) SignupBlog(c *fiber.Ctx) error {

	bm := BlogModel{}

	var request BlogCreate

	request.BlogCreateData = time.Now()
	
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	lessonSubject := c.Query("classroom")
	lessonAddBlog ,err := strconv.Atoi(c.Query("lesson"))

	err = bm.signupBlog(lessonSubject,lessonAddBlog,request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while At Blog. %s", err),
		})
	}
	
	return nil

}


func (bc BlogController) GetBlogByLessonNo(c *fiber.Ctx) error {

	bm := BlogModel{}
	cm := classroom.ClassroomModel{}

	lessonSubject := c.Query("classroom")
	getBlog ,err := strconv.Atoi(c.Query("lesson"))


	blogInfo, err := bm.getBlogByLessonNo(lessonSubject,getBlog)

	classroomName, err := cm.GetClassroombyClassID(lessonSubject)

	cname := classroomName.ClassName






	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	if  reflect.DeepEqual(blogInfo, []lesson.LessonRequest{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "No Class room",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"classroom": cname,
		"blog": blogInfo,
	})
}


func (bc BlogController) GetBlogByClassroom(c *fiber.Ctx) error {

	bm := BlogModel{}
	cm := classroom.ClassroomModel{}

	lessonSubject := c.Query("classroom")



	blogInfo, err := bm.getBlogByClassroom(lessonSubject)
	classroomName, err := cm.GetClassroombyClassID(lessonSubject)

	cname := classroomName.ClassName


	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	if  reflect.DeepEqual(blogInfo, []lesson.LessonRequest{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "No Class room",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"classroom": cname,
		"blog": blogInfo,
	})
}

func (bc BlogController) Uploadfile(c *fiber.Ctx) error{

	bm := BlogModel{}

	lessonSubject := c.Query("classroom")
	getBlog := c.Query("lesson")
	getBlogInt,err := strconv.Atoi(getBlog)

	fname :=  uuid.New().String()
	filename := fname+".md"
	


	file, err := c.FormFile("document")


	if err == nil {

		err = bm.fileUploadBlog(lessonSubject,getBlogInt,fname)

		c.SaveFile(file, fmt.Sprintf("../frontend/content/blog/%s", filename))
		return nil

	}

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while At Blog. %s", err),
		})
	}
	
	return nil

	


}


func (bc BlogController) UpdateBlog(c *fiber.Ctx) error {

	bm := BlogModel{}

	var request BlogUpdate
	
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	lessonSubject := c.Query("classroom")
	lessonAddBlog ,err := strconv.Atoi(c.Query("lesson"))

	err = bm.updateBlog(lessonSubject,lessonAddBlog,request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while At Blog. %s", err),
		})
	}
	
	return nil

}

func (bc BlogController) Deletefile(c *fiber.Ctx) error{

	bm := BlogModel{}

	var filename string

	if err := c.BodyParser(filename); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}
	
	err := bm.Deletefile(filename)

	if err != nil {
		return err
	}

	return nil


}