package main

import (
	"fmt"
	"time"
)

// распространенный способ использования шаблона Command - делегировать информацию вместо выполнения другому объекту.

type Command interface {
	Info() string // команда, которая извлекает информацию имплементатора
}

type TimePassed struct {
	start time.Time // хранит время, когда программа была запущена
}

func (t *TimePassed) Info() string {
	// возвращает строковое представление времени, прошедшего с тех пор,
	// как мы присвоили значение до момента, когда мы вызвали метод Info()
	return time.Since(t.start).String()
}

type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}

// В этом случае мы извлекаем некоторую информацию, используя шаблон команды.
// Один тип хранит информацию о времени, в то время как другой ничего не хранит, и он просто возвращает ту же простую строку.
func main() {
	timeCommand := &TimePassed{time.Now()}

	helloCommand := &HelloMessage{}

	time.Sleep(time.Second)

	fmt.Println(timeCommand.Info())
	fmt.Println(helloCommand.Info())
}
