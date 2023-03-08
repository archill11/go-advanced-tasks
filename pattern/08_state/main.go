package main

import (
	"fmt"
	"math/rand"
	"os"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Описание
// Выключатель света имеет два состояния – включено и выключено.
// Одно «состояние» может переходить в другое и наоборот. Способ, которым работает шаблон состояния, аналогичен.
// У нас есть интерфейс состояния и реализация каждого состояния, которого мы хотим достичь.
// Также обычно существует контекст, который содержит перекрестную информацию между состояниями.

// Цели
// Иметь тип, который изменяет свое собственное поведение, когда изменились какие-то внутренние вещи.
// Модельные сложные графики и конвейеры можно легко модернизировать, добавив больше состояний и перенаправив их выходные состояния.

// Небольшая игра в угадай число
// Мы собираемся разработать очень простую игру. Эта игра представляет собой игру в угадывание чисел.
// Идея проста – нам нужно будет угадать какое-то число от 0 до 10, и у нас есть всего несколько попыток, иначе мы проиграем.
// Мы оставим игроку право выбора уровня сложности, спросив, сколько попыток у пользователя есть до проигрыша.

// Нам нужен интерфейс для представления различных состояний и игровой контекст для хранения информации между состояниями.
// Для этой игры контекст должен сохранять количество повторных попыток, выиграл пользователь или нет, секретный номер, который нужно угадать, и текущее состояние.
// Состояние будет иметь метод executeState, который принимает один из этих контекстов и возвращает true, если игра закончена, или false, если нет
type GameState interface {
	executeState(*GameContext) bool
}

type GameContext struct {
	SecretNumber int
	Retries      int
	Won          bool
	Next         GameState
}

// Игрок должен быть в состоянии ввести желаемое количество повторных попыток.
// Это будет достигнуто с помощью StartState.
// Кроме того, структура StartState должна подготовить игру, установив контексту его начальное значение перед игроком:
type StartState struct{}

func (s *StartState) executeState(c *GameContext) bool {
	c.Next = &AskState{}
	// rand.Seed(time.Now().UnixNano())
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	c.SecretNumber = rand.Intn(10)
	fmt.Println("Introduce a number of retries to set the difficulty:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)
	return true
}

// Ask состояние начинается с сообщения для игрока, в котором его просят ввести новый номер.
// В следующих трех строках мы создаем локальную переменную для хранения номера, который введет игрок.
type AskState struct{}

func (a *AskState) executeState(c *GameContext) bool {
	fmt.Printf("Introduce a number between 0 and 10, you have %d tries left\n", c.Retries)

	var n int
	fmt.Fscanf(os.Stdin, "%d", &n)
	c.Retries = c.Retries - 1

	if n == c.SecretNumber {
		c.Won = true
		c.Next = &FinishState{}
	}
	if c.Retries == 0 {
		c.Next = &FinishState{}
	}
	return true
}

type FinishState struct{}

func (f *FinishState) executeState(c *GameContext) bool {
	if c.Won {
		println("Congrats, you won")
	} else {
		fmt.Printf("You loose. The correct number was: %d\n", c.SecretNumber)
	}
	// В этом случае игру можно считать законченной, поэтому мы возвращаем значение false, чтобы сказать, что игра не должна продолжаться.
	return false
}

func main() {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}
	for game.Next.executeState(&game) {
	}
}
