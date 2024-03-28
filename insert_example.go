package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func Insert_example(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Class{})

	// 示例插入
	var Name_List = []string{"reimu", "marisa", "sakuya", "sanse", "reisen"}
	var Gender_List = []string{"男", "女"}
	for idx, name := range Name_List {
		user := Student{Name: name, ClassNo: uint16(idx % 2),
			StuIDCard: fmt.Sprintf("%x", md5.Sum([]byte(name))),
			Gender:    Gender_List[idx%2],
		}
		db.Create(&user)
	}
}

func Create_external_link(db *gorm.DB) {
	db.AutoMigrate(&Class{})
	var res []Student
	db.Model(&Student{}).Where("class_no = ?", 0).Find(&res)
	class1 := Class{No: 0, Students: res}
	db.Create(&class1)

	db.Model(&Student{}).Where("class_no = ?", 1).Find(&res)
	class2 := Class{No: 1, Students: res}
	db.Create(&class2)
}

func Create_choose_lessson(db *gorm.DB) {
	db.AutoMigrate(&Student_Lesson{}, &Lesson{})
	db.Create([]Lesson{{ID: 1, Title: "code", Teacher: "Epicmo"},
		{ID: 2, Title: "sleep", Teacher: "scarletborder"},
		{ID: 3, Title: "security", Teacher: "wisdomgo"}})

	db.Create([]Student_Lesson{{StudentNo: 1, LessonID: 1}, {StudentNo: 1, LessonID: 3},
		{StudentNo: 2, LessonID: 2}, {StudentNo: 2, LessonID: 1},
		{StudentNo: 3, LessonID: 3}, {StudentNo: 3, LessonID: 1}, {StudentNo: 3, LessonID: 2},
		{StudentNo: 4, LessonID: 1}, {StudentNo: 5, LessonID: 2}})
}

func Add_additional_students(db *gorm.DB) {
	var Name_List = []string{"remilia", "yuyuko", "ran", "chan", "hongmeiling", "cirno", "rumia", "flandre", "mystia"}
	var Gender_List = []string{"男", "女"}
	for _, name := range Name_List {
		user := Student{Name: name, ClassNo: uint16(rand.Intn(2)),
			StuIDCard: fmt.Sprintf("%x", md5.Sum([]byte(name))),
			Gender:    Gender_List[rand.Intn(2)],
		}
		db.Create(&user)
	}
}

func Add_age_field(db *gorm.DB) {
	var res []Student
	db.AutoMigrate(&Student{})
	db.Model(&Student{}).Find(&res)
	for _, stu := range res {
		stu.Age = uint(rand.Intn(10) + 20)
		db.Save(&stu)
	}
}

func Add_specified_students(db *gorm.DB) {
	db.Create([]Student{
		{No: 20010, Name: "nazerin", ClassNo: 0, StuIDCard: "abc123123", Gender: "女", Age: 24},
		{No: 20023, Name: "Alice", ClassNo: 1, StuIDCard: "bef233233", Gender: "女", Age: 22},
		{No: 20020, Name: "lilywhite", ClassNo: 0, StuIDCard: "dnf456456", Gender: "女", Age: 29},
	})
}
func generateRandomDate(start, end time.Time) time.Time {
	// 计算时间范围
	delta := end.Sub(start)

	// 生成一个位于这个范围内的随机时长
	sec := rand.Int63n(int64(delta.Seconds()))

	// 将这个随机时长加到开始日期上
	return start.Add(time.Second * time.Duration(sec))
}
func Add_submission_items(db *gorm.DB) {
	db.AutoMigrate(&Submission{})
	var status_list = []string{"未审核", "审核未通过", "审核已通过"}
	var Cause_list = []string{
		"家庭紧急情况",
		"个人健康问题",
		"医疗检查",
		"牙医预约",
		"家庭成员照顾",
		"精神健康日",
		"参加亲属婚礼",
		"孩子学校活动",
		"交通问题",
		"房屋维修",
		"紧急事务处理",
		"参加葬礼",
		"法律事务",
		"工作疲劳",
		"教育培训",
		"季节性流感",
		"旅行休假",
		"重要家庭活动",
		"宠物医疗紧急情况",
		"远程工作调整",
	}
	startDate := time.Date(2013, time.August, 15, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2013, time.September, 20, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 100; i++ {
		sub := Submission{Status: status_list[rand.Intn(3)], Date: generateRandomDate(startDate, endDate),
			Cause: Cause_list[rand.Intn(len(Cause_list))]}
		db.Create(&sub)
	}
}

func Additional_submission_fields(db *gorm.DB) {
	db.AutoMigrate(&Submission{})
	var student_list []Student
	db.Model(&Student{}).Find(&student_list)
	var lesson_list []Lesson

	var submission_list []Submission
	db.Model(&Submission{}).Find(&submission_list)
	db.Model(&Lesson{}).Select("Teacher").Find(&lesson_list)

	for _, sub := range submission_list {
		sub.Student = student_list[rand.Intn(len(student_list))]
		sub.Teacher = lesson_list[rand.Intn(len(lesson_list))].Teacher
		db.Save(&sub)
	}
}

type result struct {
	Name  string `json:"Name"`
	Cause string `json:"Cause"`
}

func join_query_demo1(db *gorm.DB) {
	var res []result
	db.Model(&Submission{}).Select("Submission.Cause as Cause, Student.Name as Name").
		Joins("JOIN Student ON Student.No = Submission.Student_No").
		Where("Submission.Status = ?", "审核未通过").
		Where("Submission.Teacher = ?", "scarletborder").
		Find(&res)

	fmt.Println(res)
}
