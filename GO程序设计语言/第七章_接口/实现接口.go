package main

import "fmt"

// 动物
type Animal interface {
	eat(foodName string)
	action()
}

// 高级动物
type HighAnimal interface {
	speak()
	// 相当于继承了Animal接口，有eat和action两个方法
	Animal
}

type Dog struct {
	Name string
}

func (dog *Dog) eat(foodName string) {
	fmt.Printf("狗[%s] 吃了%s\n", dog.Name, foodName)
}

func (dog *Dog) action() {
	fmt.Printf("狗[%s] 走路\n", dog.Name)
}

type Person struct {
	Name string
}

func (p *Person) eat(foodName string) {
	fmt.Printf("人[%s] 吃了%s\n", p.Name, foodName)
}

func (p *Person) action() {
	fmt.Printf("人[%s] 走路\n", p.Name)
}

func (p *Person) speak() {
	fmt.Printf("人[%s] 说话\n", p.Name)
}

func main() {
	dog := Dog{"小黄狗"}
	man := Person{"吴亦凡"}

	dog.eat("骨头")
	man.eat("饭")
	man.speak()

	// 接口变量
	fmt.Println(">>> 定义接口变量")
	var animalInterface Animal
	var highAnimalInterface HighAnimal

	animalInterface = &dog
	highAnimalInterface = &man

	animalInterface.action()     // 狗[小黄狗] 走路
	highAnimalInterface.action() // 人[吴亦凡] 走路
	highAnimalInterface.speak()  // 人[吴亦凡] 说话

	fmt.Println(">>> 指向man")
	// 指向man
	animalInterface = &man
	animalInterface.action() // 人[吴亦凡] 走路

	fmt.Println(">>> 指向嵌入式接口变量")
	animalInterface = highAnimalInterface
	animalInterface.action() // 人[吴亦凡] 走路

	// 编译报错。原因是HighAnimal包含了Animal所有方法定义，而Animal反之
	//highAnimalInterface = animalInterface

}
