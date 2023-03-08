package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// Здесь у нас есть наша функция or, которая принимает переменный фрагмент каналов и возвращает один канал.
// Поскольку это рекурсивная функция, мы должны настроить критерии завершения.
// Первый заключается в том, что если вариационный фрагмент пуст, мы просто возвращаем нулевой канал.
// Это согласуется с идеей передачи без каналов; мы бы не ожидали, что составной канал будет что‐либо делать.
// Наш второй критерий завершения гласит, что если наш вариационный фрагмент содержит только один элемент, мы просто возвращаем этот элемент.
// Вот основной текст функции, и где происходит рекурсия.
// Мы создайте подпрограмму, чтобы мы могли ждать сообщений на наших каналах без блокировки.
// Из-за того, как мы выполняем рекурсию, каждый рекурсивный вызов or будет иметь, по крайней мере, два канала.
// В качестве оптимизации, позволяющей ограничить количество маршрутов, мы размещаем здесь особый случай для звонков по двум каналам или только с ними.
// Здесь мы рекурсивно создаем or-канал из всех каналов в нашем фрагменте после третьего индекса, а затем выбираем из этого.
// Это рекуррентное соотношение разрушит остальную часть среза на or-каналы, чтобы сформировать дерево, из которого первый сигналнал вернется.
// Мы также проходим в канале orDone, так что, когда маршруты вверх по дереву выходят, маршруты вниз по дереву также выходят.
// Это довольно лаконичная функция, которая позволяет вам объединить любое количество каналов
// вместе в один канал, который закроется, как только любой из его составляющих каналов будет закрыт или записан в него.
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}