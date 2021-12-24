package classroom

import (
	"fmt"
	"net/http"
	"reflect"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

type ClassroomController struct{}


// Post 
func (cc *ClassroomController) ClassroomSignup(c *fiber.Ctx) error{

	cm := ClassroomModel{}

	var request ClassroomSignup
	
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}


	clr, err := cm.GetClassroombyClassID(request.ClassID)

	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting class id.",
		})
	}

	if (clr != ClassroomInfo{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Class already taken",
		})
	}

	err = cm.signupClassroom(request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while sign up. %s", err),
		})
	}
	
	return nil

}

func (cc *ClassroomController) UpdateClassroom(c *fiber.Ctx) error{

	cm := ClassroomModel{}

	var request ClassroomSignup
	
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	updateclassroom := c.Query("classroom")


	err := cm.updateClassroom(updateclassroom,request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while sign up. %s", err),
		})
	}
	
	return nil

}

//Get
func (cc *ClassroomController) GetAllClassroom(c *fiber.Ctx) error {

	cm := ClassroomModel{}


	classroomInfo, err := cm.getAllClassroom()


	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	if  reflect.DeepEqual(classroomInfo, []ClassroomInfo{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "No Class room",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "get planner successfully.",
		"classroom": classroomInfo,
	})


}

// Delete
func (cc *ClassroomController) DeleteLesson(c *fiber.Ctx) error {

	cm := ClassroomModel{}

	lessonDelete := c.Query("classroom")

	err := cm.deleteClassroombyClassId(lessonDelete)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while delete classroom. %s", err),
		})
	}
	
	return nil


}