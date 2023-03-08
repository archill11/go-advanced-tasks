package main

import (
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// При использовании шаблона проектирования Command мы пытаемся инкапсулировать какое-то действие
// или информацию в облегченный пакет, который должен быть обработан где-то в другом месте.
// Это похоже на шаблон стратегии, но, на самом деле, Command может запустить предварительно настроенную стратегию где-то в другом месте, так что они не совпадают.

// Цели
// Инкапсулировать некоторую информацию в коробку. Получатель откроет коробку и узнает ее содержимое.
// Делегируйте какое-то действие куда-то еще.

// Простая очередь
// Мы поместим некоторую информацию в реализацию Command, и у нас будет очередь.
// Мы создадим множество экземпляров типа, реализующего шаблон Command, и передадим их в очередь,
// которая будет хранить команды до тех пор, пока три из них не окажутся в очереди, после чего она их обработает.

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}
		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}
