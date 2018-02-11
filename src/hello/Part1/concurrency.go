package main

//并发访问对象 在go里，我们可以在对象内部保存一个函数类型的channel，涉及对象状态的操作都放入channel里，对象初始化的时候开启一条goroutinue，不停地执行匿名函数

import (
	"fmt"
	"strconv" //类型转换
	"time"
)

type Person struct {
	name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()

	return p
}

func (p *Person) backend() {
	//管道阻塞等待数据 ,相当于串行线程
	for f := range p.chF {
		f()
	}
}

func (p *Person) addSalary(sal float64) {
	p.chF <- func() { p.salary += sal }
}

func (p *Person) reduceSalary(sal float64) {
	p.chF <- func() { p.salary -= sal }
}

func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }

	return <-fChan
}

func (p *Person) String() string {
	return p.name + "- salary is:" + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	p := NewPerson("jack", 8888.0)
	fmt.Println(p)

	for i := 1; i <= 500; i++ {
		go func() {
			p.addSalary(2)
		}()
	}
	for i := 1; i <= 500; i++ {
		go func() {
			p.reduceSalary(1)
		}()
	}

	time.Sleep(3e9)
	fmt.Println("after changed:")
	fmt.Println(p)
}
