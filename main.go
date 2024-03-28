package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Student struct {
	No        uint16 `gorm:"primaryKey"` // 学号
	Name      string
	ClassNo   uint16
	StuIDCard string // 身份证号
	Gender    string `gorm:"check:gender IN ('男','女')"`
	Age       uint
}

type Class struct {
	ID uint16 `gorm:"primaryKey"`
	No uint16
	// has many
	Students []Student `gorm:"foreignKey:ClassNo;references:No"`
}

type Lesson struct {
	ID      uint
	Title   string
	Teacher string
}

type Student_Lesson struct {
	ID        uint
	StudentNo uint16
	LessonID  uint
}

type Submission struct {
	ID        uint
	Status    string `gorm:"check:status IN ('审核已通过','审核未通过','未审核')"`
	Date      time.Time
	Cause     string
	StudentNo uint16
	Student   Student `gorm:"foreignKey:StudentNo;references:No"`
	Teacher   string
}

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/student_new_db?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// db.AutoMigrate(&Student{}, &Class{}, &Lesson{}, &Student_Lesson{})

	// Insert_example(db)
	// Create_external_link(db)

	// Create_choose_lessson(db)
	// Add_additional_students(db)
	// Add_age_field(db)
	// Add_specified_students(db)
	// Add_submission_items(db)
	// Additional_submission_fields(db)
	// join_query_demo1(db)
}
