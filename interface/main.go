package main

import "fmt"

type notifier interface {
	call()
}

type nokia struct {
	osv string
}

// 值接收者，不改变原来值
func (n nokia) call() {
	n.osv = "N11"
	fmt.Println("i am nokia." + n.osv)
}

type iphone struct {
	osv string
}

// 指针接收者，引用相同内存，改变原来值
func (i *iphone) call() {
	i.osv = "ios13"
	fmt.Println("I am iphone." + i.osv)
}

func notify(n notifier) {
	n.call()
}

func main() {
	p1 := nokia{
		osv: "N12",
	}
	notify(p1)
	fmt.Println(p1.osv)

	p2 := iphone{
		osv: "ios14",
	}
	notify(&p2)
	fmt.Println(p2.osv)
}
