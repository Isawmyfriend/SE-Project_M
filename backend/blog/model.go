package blog

import (
	"backend/database"
	"backend/lesson"

	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogModel struct{}

func (bm BlogModel) signupBlog(lesson_subject string,lesson_no int,body BlogCreate) (err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")


	_, err = lessonCollection.UpdateOne(
		ctx,
		bson.M{"$and": []bson.M{

			{"lessonno": lesson_no},
			{"lesssubject":lesson_subject},
	
		}},
		bson.M{"$set": bson.M{"lessonname":body.LessonName,"lessondetail": body.BlogDetail,"createdate": body.BlogCreateData}})

	if err != nil {
		log.Fatal(err)
	}

	return nil

}



// Get 
func (bm BlogModel) getBlogByLessonNo(lesson_subject string,lesson_no int) (blog lesson.LessonRequest,err error){

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	cursor := lessonCollection.FindOne(ctx, bson.M{
		"$and": []bson.M{
			{"lessonno": lesson_no},
			{"lesssubject":lesson_subject},
		},
	}).Decode(&blog)


	if cursor == nil {
		return blog, err
	}

	return blog, nil

}


func (bm BlogModel) getBlogByClassroom(lesson_subject string) (blog []lesson.LessonRequest,err error){

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)
	

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	opts := options.Find()
	opts.SetSort(bson.D{{"lessonno", 1}})

	cursor,err := lessonCollection.Find(ctx, bson.M{"lesssubject": lesson_subject}, opts)

	if err != nil {
		log.Fatal(err)
	} else if cursor == nil {
		return blog, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var c lesson.LessonRequest
		if err = cursor.Decode(&c); err != nil {
			log.Fatal(err)
			return blog, err
		}

		blog = append(blog, c)
	}

	return blog, nil

}



func (bm BlogModel) fileUploadBlog(lesson_subject string,lesson_no int, lesson_detail string) (err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	loc, err := time.LoadLocation("Asia/Bangkok")
    if err != nil {
        panic(err)
    }

	
	BlogCreateData := time.Now().In(loc).Format("2006-01-02 15:04:05")




	_, err = lessonCollection.UpdateOne(
		ctx,
		bson.M{"$and": []bson.M{

			{"lessonno": lesson_no},
			{"lesssubject":lesson_subject},
	
		}},
		bson.M{"$set": bson.M{"lessondetail": lesson_detail,"createdate": BlogCreateData}})

	if err != nil {
		log.Fatal(err)
	}

	return nil

}


func (bm BlogModel) updateBlog(lesson_subject string,lesson_no int,body BlogUpdate) (err error) {

	db, ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	

	seDatabase := db.Database("se_project")
	lessonCollection := seDatabase.Collection("lesson")

	_, err = lessonCollection.UpdateOne(
		ctx,
		bson.M{"$and": []bson.M{

			{"lessonno": lesson_no},
			{"lesssubject":lesson_subject},
	
		}},
		bson.M{"$set": bson.M{"lessonno": body.LessonNo,"lessonname": body.LessonName}})

	if err != nil {
		log.Fatal(err)
	}

	return nil

}


func (bm BlogModel) Deletefile(fname string) error{

	filename := fname+".md"

	path := ("../frontend/content/blog/"+ filename)

	err := os.Remove(path)  // delete an entire directory

	if err != nil {
		return err
	}

	return nil


}
