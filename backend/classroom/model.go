package classroom

import (
	"backend/database"
	"backend/lesson"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClassroomModel struct{}

func (cm ClassroomModel) signupClassroom(body ClassroomSignup) (err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	classroomCollection := seDatabase.Collection("classroom")

	_ ,err = classroomCollection.InsertOne(ctx,body)

	if err != nil {
		log.Fatal(err)
	}

	return nil

}

func (cm ClassroomModel) updateClassroom(class_id string,body ClassroomSignup) (err error) {

	lm := lesson.LessonModel{}

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	classroomCollection := seDatabase.Collection("classroom")

	_ ,err = classroomCollection.ReplaceOne(
		ctx,
		bson.M{"classid": class_id},
		bson.M{
			"classid": body.ClassID,
			"classname":  body.ClassName,
			"classinstructor":  body.ClassInstructor,
			"classdetail":  body.ClassDetail,
			"classimage": body.ClassImage ,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	err = lm.UpdateLessonSubject(class_id,body.ClassID)
	
	if err != nil {
		log.Fatal(err)
	}

	return nil

}


func (cm ClassroomModel) GetClassroombyClassID(classId string) (classroom ClassroomInfo,err error){

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	classroomCollection := seDatabase.Collection("classroom")

	cursor := classroomCollection.FindOne(ctx,bson.M{"classid": classId}).Decode(&classroom);

	if cursor == nil {
		return classroom, err
	}

	return classroom,nil

}


func (cm ClassroomModel) getAllClassroom() (classroom []ClassroomInfo,err error){

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	opts := options.Find()
	opts.SetSort(bson.D{{"classname", 1}})

	seDatabase := db.Database("se_project")
	classroomCollection := seDatabase.Collection("classroom")

	cursor, err := classroomCollection.Find(ctx, bson.M{},opts)
	if err != nil {
		log.Fatal(err)
	} else if cursor == nil {
		return classroom,err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var c ClassroomInfo
		if err = cursor.Decode(&c); err != nil {
			log.Fatal(err)
			return classroom,err
		}
		
		classroom = append(classroom, c)
	}

	return classroom, nil

}


// Delete
func (cm ClassroomModel) deleteClassroombyClassId(class_id string)  error{

	db, ctx := database.DB()

	err := db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("classroom")

	_ , err = lessonCollection.DeleteOne(ctx, bson.M{"classid": class_id})

	if err != nil {
		log.Fatal(err)
	}

	return  nil

}