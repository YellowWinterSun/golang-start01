package main

import "fmt"

/**
	定义结构体
 */
type Student struct {
	id int32
	name string
	rank int8
	classId int32
}

type SchoolClass struct {
	id int32
	title string
	stus []Student
}

func main() {
	/*
		基本声明
	 */
	stu1 := Student{
		1,
		"hdy",
		1,
		1,
	}
	stu2 := Student{2, "azure", 2, 1}
	// 声明一个默认初始值的结构体对象
	stu3 := Student{name: "null"}
	// 声明一个带有某些初始值的结构体对象
	stu4 := Student{
		id: 4,
		name: "meiji",
	}

	fmt.Println(stu1, stu2, stu3, stu4)

	/*
		指针用法
	 */
	p1 := &stu1
	p2 := &stu2

	p1.rank = 5
	p2.rank = p1.rank
	p1.rank = 6

	fmt.Println(stu1, stu2)

	/*
		嵌套struct
	 */
	class1 := SchoolClass{
		id: 1,
		title: "宇宙第一班",
		//stus: make([]Student, 4),
	}
	//class1.stus[0] = stu1
	//class1.stus[1] = stu2
	//class1.stus[2] = stu3
	//class1.stus[3] = stu4
	class1.stus = append(class1.stus, stu1)
	class1.stus = append(class1.stus, stu2)
	class1.stus = append(class1.stus, stu3)
	class1.stus = append(class1.stus, stu4)

	fmt.Println("班级", class1)
	fmt.Println("班级人数", len(class1.stus))

	for index := 0; index < len(class1.stus); index++ {
		fmt.Println(class1.stus[index].name)
	}

}
