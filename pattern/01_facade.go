package pattern

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Описание
// Фасад, с точки зрения архитектуры, - это передняя стена, которая скрывает комнаты и коридоры
// здания. Он защищает своих обитателей от холода и дождя и обеспечивает им уединение. Это упорядочивает
// и разделяет жилища.
// Шаблон проектирования фасада делает то же самое, но в нашем коде. Он защищает код от
// нежелательного доступа, заказывает некоторые вызовы и скрывает от пользователя область сложности.

// Цели
// Вы используете Facade, когда хотите скрыть сложность некоторых задач, особенно когда большинство
// из них совместно используют утилиты (такие как аутентификация в API). Библиотека - это форма фасада,
// где кто-то должен предоставить разработчику некоторые методы, позволяющие делать определенные вещи
// удобным способом. Таким образом, если разработчику необходимо использовать вашу библиотеку, ему не нужно знать все
// внутренние задачи, чтобы получить желаемый результат.
// Итак, вы используете шаблон проектирования фасада в следующих сценариях:
// Когда вы хотите уменьшить сложность некоторых частей нашего кода. Вы скрываете эту сложность за фасадом, предоставляя более простой в использовании метод.
// Когда вы хотите сгруппировать действия, которые взаимосвязаны, в одном месте.
// Когда вы хотите создать библиотеку, чтобы другие могли использовать ваши продукты, не беспокоясь о том, как все это работает.

// Пример
// В качестве примера мы собираемся предпринять первые шаги к написанию нашей собственной библиотеки, которая обращается к сервису Openweathermap.
// В случае, если вы не знакомы с сервисом OpenWeatherMap, это HTTP-сервис, который предоставляет вам оперативную информацию о погоде, а также исторические данные о ней.
// HTTP REST API очень прост в использовании и будет хорошим примером того, как создать шаблон фасада для сокрытия сложности сетевых подключений за сервисом REST.

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (*Weather, error)
	GetByGeoCoordinates(lat, lon float32) (*Weather, error)
}

// Eдиный тип для доступа к данным.
// Вся информация, полученная из сервиса OpenWeatherMap будет проходить через него.
type CurrentWeatherData struct {
	APIkey string
}

func (p *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weatherq=%f,%f&APPID=%s", lat, lon, c.APIkey),
	)
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&APPID=%s", city, countryCode, c.APIkey),
	)
}

func (o *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("status code was %d, aborting. Error message was: %s", resp.StatusCode, errMsg)
		return
	}
	weather, err = o.responseParser(resp.Body)
	resp.Body.Close()
	return
}

func getMockData() io.Reader {
	response := `{
		"coord": {"lon":-3.7,"lat":40.42},
		"weather" : [{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],
		"base": "stations",
		"main": {"temp":303.56,"pressure":1016.46,"humidity": 26.8,"temp_min":300.95,"temp_max":305.93},
		"wind": {"speed":3.17,"deg":151.001},
		"rain": {"3h":0.0075},
		"clouds": {"all":68},
		"dt": 1471295823,
		"sys": {"type":3,"id":1442829648,"message":0.0278,"country":"ES","sunrise":1471238808,"sunset":1471288232},
		"id": 3117735,
		"name": "Madrid",
		"cod": 200
	}`
	r := bytes.NewReader([]byte(response))
	return r
}

type Weather struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`

	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`

	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`

	Dt uint32 `json:"dt"`

	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}