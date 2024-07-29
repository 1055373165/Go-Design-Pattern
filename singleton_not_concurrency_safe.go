package main

import "fmt"

// 保证这个类私有化，外界不能通过这个类直接创建出一个对象
type singleton struct {
}

func (s *singleton) SomeThing() {
	fmt.Println("单例的某个方法进行处理...")
}

// 保证有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向(golang中没有常指针概念，所以只能通过将这个指针私有化，不让外部模块直接访问)
var instance *singleton = new(singleton)

// 如果全部都是私有化，那么外部模块将永远无法访问到这个对象
// 需要提供一个方法对外暴露出这个对象
// 不能将 GetInstance 作为 singleton 的一个成员方法，因为外界显访问对象再访问函数
// 但是类和对象已经私有化了，外界无法访问，所有这个方法一定是一个全局普通函数。
func GetInstance() *singleton {
	return instance
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()
}
