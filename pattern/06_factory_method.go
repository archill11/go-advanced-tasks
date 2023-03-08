package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Шаблон Factory method (или просто Factory), вероятно, является вторым по известности и
// использованию шаблоном проектирования в отрасли. Его цель состоит в том, чтобы абстрагировать пользователя от знаний
// о структуре, которые ему нужны для достижения определенной цели, такой как извлечение некоторого значения,
// возможно, из веб-службы или базы данных. Пользователю нужен только интерфейс, который предоставляет ему
// это значение. Делегируя это решение фабрике, эта фабрика может предоставить интерфейс, который
// соответствует потребностям пользователя. Это также облегчает процесс понижения рейтинга или обновления
// реализации базового типа, если это необходимо.

// Описание
// При использовании шаблона проектирования Factory method мы получаем дополнительный уровень инкапсуляции, чтобы наша программа могла развиваться в контролируемой среде.
// С помощью фабричного метода мы делегируем создание семейств объектов другому пакету или объекту, чтобы абстрагироваться от знания пула возможных объектов, которые мы могли бы использовать.
// Представьте, что вы хотите организовать свой отпуск с помощью туристического агентства.
// Вы не занимаетесь отелями и путешествиями, а просто сообщаете агентству интересующий вас пункт назначения, чтобы они предоставили вам все необходимое.
// Туристическое агентство представляет собой фабрику поездок.

// Цели
// Делегирование создания новых экземпляров структур другой части программы.
// Работа на уровне интерфейса, а не с конкретными реализациями.
// Группировка семейств объектов для получения создателя объектов семейства.

// Пример – фабрика способов оплаты для магазина.
// В нашем примере мы собираемся внедрить фабрику платежных методов, которая предоставит нам различные способы оплаты в магазине.
// Вначале у нас будет два способа оплаты – наличными и кредитной картой.
// У нас также будет интерфейс с методом Pay, который должна реализовать каждая структура, которая хочет использоваться в качестве способа оплаты.

type PaymentMethod interface {
	Pay(amount float32) string // общий метод для каждого способа оплаты под названием Pay.
}

// Мы должны определить идентифицированные способы оплаты фабрики как константы,
// чтобы мы могли вызывать и проверять возможные способы оплаты из-за пределов пакета.
type payment int

const (
	Cash      payment = 1
	DebitCard payment = 2
)

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

type DebitCardPM struct{}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}

func GetPaymentMethod(m payment) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized\n ", m)
	}
}
