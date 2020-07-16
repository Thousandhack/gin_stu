package main

import "fmt"

//
type Monkey struct {
	Name string
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "会爬树。。")
}


type LittleMonkey struct {
	// LittleMonkey 继承 Monkey  的 climbing()
	Monkey

}

func (l *LittleMonkey) playing() {
	fmt.Println(l.Name, "通过学习,会玩游戏。。")
}


type Fly interface {
	flying()
}



func main() {
	// 创建一个实例
	monkey := LittleMonkey{
		Monkey{Name: "test01"}}

	monkey.climbing()
	monkey.playing()     // 不破坏原来继承的情况下，再加自己的特性的方法
}