package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Шаблон проектирования Builder – повторное использование алгоритма для создания множества реализаций интерфейса

// Шаблон Builder помогает нам создавать сложные объекты без непосредственного создания их структуры или написания логики, которая им требуется.
// Представьте себе объект, который мог бы содержать десятки полей, которые сами по себе являются более сложными структурами.
// Теперь представьте, что у вас есть много объектов с этими характеристиками, и у вас могло бы быть больше.
// Мы не хотим писать логику для создания всех этих объектов в пакете, который просто должен использовать объекты.

// Описание
// Создание экземпляра может быть таким же простым, как предоставление открывающих и закрывающих фигурных скобок {} и оставление экземпляра с нулевыми значениями,
// или таким же сложным, как объект, которому необходимо выполнить некоторые вызовы API, проверить состояния и создать объекты для своих полей.
// У вас также может быть объект, состоящий из множества объектов.
// В то же время вы могли бы использовать одну и ту же технику для создания многих типов объектов.
// Например, при сборке автомобиля вы будете использовать почти ту же технику, что и при сборке автобуса, за исключением того, что они будут разных размеров и количества посадочных мест,
// так почему бы нам не использовать процесс сборки повторно?
// Вот тут-то и приходит на помощь шаблон Builder.

// Цели
// Абстрагирование создания сложных объектов таким образом, чтобы создание объекта было отделено от пользователя объекта.
// Создавайте объект шаг за шагом, заполняя его поля и создавая внедренные объекты.
// Повторно используйте алгоритм создания объекта между многими объектами.

// Пример – производство транспортных средств
// Шаблон проектирования Builder обычно описывается как взаимосвязь между
// директором, несколькими разработчиками и продуктом, который они создают. Продолжая наш пример с
// автомобилем, мы создадим конструктор транспортных средств. Процесс (широко описываемый как алгоритм)
// создания транспортного средства (продукта) более или менее одинаков для любого вида транспортного средства – выберите
// тип транспортного средства, соберите конструкцию, установите колеса и сиденья. Если вы подумаете
// об этом, вы могли бы построить автомобиль и мотоцикл (два конструктора) по этому описанию, так что мы
// повторно используем описание для создания автомобилей на производстве. В нашем примере директор представлен
// типом Директор по производству

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Переменная директора по производству - это тот, кто отвечает за прием строителей.
// У него есть метод Construct, который будет использовать конструктор, и будет воспроизводить требуемые шаги.
type ManufacturingDirector struct {
	builder BuildProcess
}

// Метод Set Builder позволит нам изменить конструктор, используемый в директоре производства
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4 * 2
	return b
}
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
