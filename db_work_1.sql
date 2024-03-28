show databases;
create database if not exists student_new_db;
use student_new_db;

-- 创建学生表，另外携带extra这一个额外的需要做修改的额外信息

create table if not exists student_info(
	stu_no BIGINT unsigned auto_increment,
	stu_name varchar(16) not null,
	stu_class INT unsigned not null,
	stu_extra INT,
	primary key (`stu_no`)
);

-- 修改stu_extra 字段
alter table student_info modify column stu_extra char(2);

alter table student_info
modify column stu_extra char(2) not null default 'm';

alter table student_info
drop stu_extra;

-- 增加constraint

ALTER TABLE student_info 
ADD (stu_desp VARCHAR(18));

ALTER TABLE  student_info 
ADD CONSTRAINT new_column_chk 
CHECK (stu_desp IS NOT NULL AND LENGTH(stu_desp) = 18 AND REGEXP_LIKE(stu_desp, '^[0-9]+$'));

ALTER TABLE  student_info 
ADD CONSTRAINT stu_constrain UNIQUE (stu_desp);

-- more constraint

alter table student_info 
add constraint stu_class_straint 
check (stu_class > 0);


alter table student_info 
add column stu_submis char(21) not null default '未审核';

alter table student_info 
add constraint stu_submisstrain
check (stu_submis in ('未审核','审核已通过','审核未通过'));

-- create index
create index stu_name_idx on student_info
(stu_name asc);

-- before DDL & DML, drop some unnecessary table
alter table student_info 
drop stu_desp;

alter table student_info 
drop stu_submis;

-- DDL and DML
insert into student (name, class_no, stu_id_card, gender)
values ('youmu', 1, 'cdcd123123', '女');

-- attempt to drop one item with foreignKey relationship
delete from class where no = 0;

-- attemp tp update one item with foreignKey relationship
update class set no = 50 where no = 0;

-- create view
create view stu_cls1_view (stu_name, stu_gender) as
select name, gender from student where class_no = 1;

select * from stu_cls1_view where stu_name like '%i%';

-- join query
-- 这里尝试Join学生和选课表，依据学生名字查出他所以选的课程
select title from lesson join student_lesson on student_lesson.lesson_id = lesson.id 
join student on student.`no` = student_lesson.student_no 
where student.name = "reimu";

-- update in view
update stu_cls1_view set stu_gender = "女" where stu_cls1_view.stu_gender = "男";
select * from stu_cls1_view scv ;

-- easy query
-- query all boys' detail
select `no` ,name ,class_no  from student s where gender = '男';

-- decs order
-- firstly, let's add some girls
-- refer to my Golang code

select * from student s  where gender = '女' order by class_no desc; 

-- query age and queue
-- first create age field emmmm...
select * from student s order by age asc;

-- query student who has specified stuno 
-- add first
select * from student s where cast(no as char) like '2002%';
select * from student s where cast(no as char) like '%01%';

-- add additional table
select * from submission s where `date` >= '2013-09-04' and status = '未审核';

select ID from submission s where status in ('未审核','审核未通过');

select ID,cause from submission s where `date` >= '2013-08-31' and `date` < '2013-09-02';
select ID,cause from submission s where `date` between  '2013-08-31' and '2013-09-02';

-- join and query count(student) (not need to be unique) of one teacher
select count(student_lesson.student_no) from student_lesson
join (select * from lesson where lesson.teacher = 'scarletborder') as new_sl
on student_lesson.lesson_id = new_sl.id;
-- as same as above, but unique student
select count(distinct  student_lesson.student_no) from student_lesson
join (select * from lesson where lesson.teacher = 'scarletborder') as new_sl
on student_lesson.lesson_id = new_sl.id;

select Submission.Cause as Cause, Student.Name as Name from submission join student 
on Student.No = Submission.Student_No where Submission.Status = '审核未通过' and 
Submission.Teacher = 'scarletborder';

-- statistic lessons' student
select  lesson.title ,count(distinct student_lesson.Student_No) as `number` from student_lesson 
join lesson on lesson.id = student_lesson.lesson_id  group by lesson.id ;

select lesson.title from student_lesson join lesson on lesson.id = student_lesson.lesson_id 
group by lesson.id having count(distinct student_lesson.student_no) >= 3
order by lesson.title desc ;

-- complex query
select Teacher from (select Teacher, student_no from submission where status = '未审核') as s 
group by Teacher order by count(student_no) desc limit 1;

select lesson.title ,lesson.teacher from lesson 
join (select student_lesson.lesson_id  from student_lesson group by student_lesson.lesson_id 
order by count(student_lesson.student_no) desc limit 2) as s on s.lesson_id = lesson.id;

-- 假设课程id = 1的 code 是<算法设计>
select student.name from student 
join ( select student_lesson.student_no from student_lesson
	join (select lesson.id from lesson where lesson.title = 'code') as l 
	group by student_lesson.student_no having count(student_lesson.lesson_id) = 1
) as `ssl` on student.`no`  = `ssl`.student_no

-- 选取选修了所有课程
SELECT student.name
FROM student
JOIN student_lesson ON student.`no` = student_lesson.student_no
JOIN lesson ON student_lesson.lesson_id = lesson.id
GROUP BY student.name
HAVING COUNT(DISTINCT student_lesson.lesson_id) = (SELECT COUNT(*) FROM lesson);



