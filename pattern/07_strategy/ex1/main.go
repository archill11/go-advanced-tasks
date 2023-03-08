package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Описание
// Шаблон стратегии использует различные алгоритмы для достижения некоторой конкретной функциональности.
// Эти алгоритмы скрыты за интерфейсом и, конечно же, они должны быть взаимозаменяемыми.
// Все алгоритмы достигают одной и той же функциональности по-разному.
// Например, у нас мог бы быть интерфейс сортировки и несколько алгоритмов сортировки.
// Результат тот же, некоторый список отсортирован, но мы могли бы использовать быструю сортировку, сортировку слиянием и так далее.

// Цели
// Предоставить несколько алгоритмов для достижения некоторой конкретной функциональности.
// Все типы обеспечивают одинаковую функциональность по-разному, но на клиента стратегии это не влияет.
// Шаблон стратегии на самом деле используется для множества сценариев, и многие программные инженерные решения поставляются с какой-то внутренней стратегией.

// Рендеринг изображений или текста
// Для этого примера мы собираемся печатать текст в консоли и рисовать объекты в файле.
// В этом случае у нас будут две стратегии: консольная и файловая. Но пользователю библиотеки не придется иметь дело со сложностью, стоящей за ними.
// Ключевая особенность заключается в том, что “вызывающий” не знает, как работает базовая библиотека, и он просто знает информацию,
// доступную по определенной стратегии.

// Наша стратегия определяет простой метод Print(), который возвращает ошибку.
// Типы, которые должны реализовать стратегию печати, будут называться ConsoleSquare и ImageSquare.
type OutputStrategy interface {
	Draw() error
}

type ConsoleSquare struct{}

func (t *ConsoleSquare) Draw() error {
	println("Square")
	return nil
}

type ImageSquare struct {
	DestinationFilePath string
}

func (t *ImageSquare) Draw() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}

	bgRectangle := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(bgRectangle, bgRectangle.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgRectangle, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("error opening image")
	}
	defer w.Close()

	if err = jpeg.Encode(w, bgRectangle, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}
	return nil
}

// go run main.go -output=console  || go run main.go -output=image
// var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")
var output = flag.String("output", "image", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()

	var activeStrategy OutputStrategy

	switch *output {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{"./tmp/image.jpg"}
	default:
		activeStrategy = &ConsoleSquare{}
	}

	err := activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
