package lesson

import (
	"backend/database"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LessonModel struct{}

func (lm LessonModel) signupLesson(body LessonSignup) (err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")


	_, err = lessonCollection.InsertOne(ctx, body)

	if err != nil {
		log.Fatal(err)
	}

	return nil

}


func (lm LessonModel) updateLesson(lesson_no int,body LessonInfo) error{

	db, ctx := database.DB()

	err := db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	_ ,err = lessonCollection.ReplaceOne(
		ctx,
		bson.M{"lessonon": lesson_no},
		bson.M{
			"lesson_no": body.LessonNo,
			"lesson_name":  body.LessonName,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (lm LessonModel) UpdateLessonSubject(class_id string,lesson_subject string) error{

	db, ctx := database.DB()

	err := db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	_ ,err = lessonCollection.UpdateMany(
		ctx,
		bson.M{"lesssubject": class_id},
		bson.D{
			{"$set", bson.D{{"lesssubject", lesson_subject}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}




// Get
func (lm LessonModel) getLessonbyLessonNo(lesson_subject string,lesson_no int) (lesson LessonInfo, err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	
	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	

	cursor := lessonCollection.FindOne(ctx,bson.M{
		"$and": []bson.M{
			{"lessonno": lesson_no},
			{"lesssubject":lesson_subject},
		},
	}).Decode(&lesson)

	

	if cursor == nil {
		return lesson, err
	}

	return lesson, nil

}


func (lm LessonModel) getLessonbyClassID(classId string) (lesson []LessonInfo, err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	opts := options.Find()
	opts.SetSort(bson.D{{"lessonno", 1}})

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	cursor,err := lessonCollection.Find(ctx, bson.M{"lesssubject": classId}, opts)

	if err != nil {
		log.Fatal(err)
	} else if cursor == nil {
		return lesson, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var c LessonInfo
		if err = cursor.Decode(&c); err != nil {
			log.Fatal(err)
			return lesson, err
		}

		lesson = append(lesson, c)
	}

	return lesson, nil

}

func (lm LessonModel) getAllLesson() (lesson []LessonInfo, err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	cursor, err := lessonCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	} else if cursor == nil {
		return lesson, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var c LessonInfo
		if err = cursor.Decode(&c); err != nil {
			log.Fatal(err)
			return lesson, err
		}

		lesson = append(lesson, c)
	}

	return lesson, nil

}

// Delete
func (lm LessonModel) deleteLessonbyLessonNo(lesson_subject string,lesson_no int)  error{

	db, ctx := database.DB()

	err := db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	lessonBeforedelete,err := lm.getLessonbyLessonNo(lesson_subject,lesson_no)

	err = lm.Deletefile(lessonBeforedelete.LessonDetail)

	_ , err = lessonCollection.DeleteOne(ctx, bson.M{
		"$and": []bson.M{
			{"lessonno": lesson_no},
			{"lesssubject":lesson_subject},
		},
	})

	

	if err != nil {
		log.Fatal(err)
	}

	return  nil

}


func (lm LessonModel) Deletefile(fname string) error{

	filename := fname+".md"

	path := ("../frontend/content/blog/"+ filename)

	err := os.Remove(path)  // delete an entire directory

	if err != nil {
		return err
	}

	return nil


}
