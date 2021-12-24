package lesson

import (
	"fmt"
	"net/http"

	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type LessonController struct{}

func (lc *LessonController) SignupLesson(c *fiber.Ctx) error {


	lm := LessonModel{}

	var request LessonSignup

	request.LessSubject = c.Query("classroom")

	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	


	lr, err := lm.getLessonbyLessonNo(request.LessSubject,request.LessonNo)

	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting class id.",
		})
	}

	if (lr != LessonInfo{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Lesson already taken",
		})
	}

	err = lm.signupLesson(request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while sign up lesson. %s", err),
		})
	}
	
	return nil



}

func (lc *LessonController) UpdateLesson(c *fiber.Ctx) error{

	lm := LessonModel{}

	var request LessonInfo
	
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	updatelesson ,err := strconv.Atoi(c.Query("lesson"))


	err = lm.updateLesson(updatelesson,request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while sign up. %s", err),
		})
	}
	
	return nil

}


func (lc *LessonController) GetAllLesson(c *fiber.Ctx) error {

	lm := LessonModel{}

	lessonInfo, err := lm.getAllLesson()


	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	if  reflect.DeepEqual(lessonInfo, []LessonInfo{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "No Lesson",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get planner successfully.",
		"lesson": lessonInfo,
	})


}

func (lc *LessonController) GetLessonbyClassID(c *fiber.Ctx) error{

	lm := LessonModel{}

	findLesson := c.Query("classroom")

	lessonInfo, err := lm.getLessonbyClassID(findLesson)

	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	if  reflect.DeepEqual(lessonInfo, []LessonInfo{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "No Lesson",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get planner successfully.",
		"lesson": lessonInfo,
	})


}


func (lc *LessonController) DeleteLesson(c *fiber.Ctx) error {

	lm := LessonModel{}


	classroomLesson := c.Query("classroom")
	lessonDelete,err := strconv.Atoi(c.Query("lesson"))

	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while covert string to int. %s", err),
		})
	}

	err = lm.deleteLessonbyLessonNo(classroomLesson,lessonDelete)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while delete lesson. %s", err),
		})
	}


	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error Delete File",
		})
	
	}
	
	return nil




}


func (lc *LessonController) DeleteFile(c *fiber.Ctx) error {

	lm := LessonModel{}


	oldLessonFile := c.Query("filename")


	
	err := lm.Deletefile(oldLessonFile)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while delete file. %s", err),
		})
	}
	
	return nil




}


