package pattern

import (
	"fmt"
	"io"
	"os"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Описание
// В шаблоне проектирования Visitor мы пытаемся отделить логику, необходимую для работы с конкретным объектом, от самого объекта.
// Таким образом, у нас могло бы быть много разных посетителей, которые делают что-то с определенными типами.
// Например, представьте, что у нас есть логгер, которое записывает данные в консоль.
// Мы могли бы сделать регистратор “доступным для посещения”, чтобы вы могли добавлять любой текст к каждому журналу.
// Мы могли бы написать шаблон Visitor, который добавляет дату, время и имя хоста к полю, хранящемуся в объекте.

// Цели
// Чтобы отделить алгоритм некоторого типа от его реализации в рамках какого-либо другого типа.
// Повысить гибкость некоторых типов, используя их практически без логики, чтобы можно было добавлять все новые функциональные возможности без изменения структуры объекта.
// Чтобы исправить структуру или поведение, которые нарушили бы принцип открытости/ закрытости в типе.

// Приложение для логгирования
// Мы собираемся разработать простое приложение для логгирования в качестве примера шаблона Visitor.
// Для этого конкретного примера мы создадим посетителя, который добавляет различную информацию к типам, которые он “посещает”.

type Visitable interface {
	Accept(Visitor)
}

type MessageA struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

// Посетитель - это тип, который будет действовать в пределах типа, доступного для посещения.
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

// Посетитель добавит к сообщению для печати текст “Посетил A” или “Посетил B” соответственно.
type MessageVisitor struct{}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

type MsgFieldVisitorPrinter struct{}

func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}
func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}
