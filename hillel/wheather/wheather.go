// wheather Пакет погоды

package wheather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WheatherNow Текуцая погода
var WheatherNow Wheather

func init() {
	WheatherNow = New()
}

var postalCode string = "121552"
var country string = "RU"

// Wheather Погода
// type Wheather struct {
// 	RelativeHumidity    string `json:"rh"`
// 	Pressure            string `json:"pres"`
// 	City                string `json:"city_name"`
// 	Temperature         string `json:"temp"`
// 	TemperatureApparent string `json:"app_temp"`
// }

// Image Wheather
type Image struct {
	Icon        string `json:"icon"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Wheather Погода
type Wheather struct {
	City                string  `json:"city_name"`
	Temperature         float64 `json:"temp"`
	TemperatureApparent float64 `json:"app_temp"`
	Pressure            float64 `json:"pres"`
	RelativeHumidity    float64 `json:"rh"`
	Image               `json:"weather"`
}

type wheatherAPI struct {
	Data []Wheather `json:"data"`
}

// New wheather
func New() Wheather {
	// return getWheather("https://api.weatherbit.io/v2.0/current?postal_code=121552&key=9eea2cdd9d39423bb8544131ae377e30")
	return getWheather("https://api.weatherbit.io/v2.0/current?lat=55.740&lon=37.422&key=9eea2cdd9d39423bb8544131ae377e30")
	// return getWheather("https://api.weatherbit.io/v2.0/current?lat=55.740948&lon=37.422611&key=9eea2cdd9d39423bb8544131ae377e30")
	// 55.740948, 37.422611

}

// getWheather Возвращает объект погоды
func getWheather(url string) Wheather {
	body := getData(url)

	var wapi wheatherAPI
	json.Unmarshal(body, &wapi)
	str := fmt.Sprint(string(body))
	fmt.Println(str)
	var wht Wheather = wapi.Data[0]
	wht.Pressure = wht.Pressure * 0.750063755419211
	// https://www.weatherbit.io/static/img/icons/c01n.png
	return wht
}

func getData(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body) // was "res", now "res.Body"
	if err != nil {
		panic(err)
	}
	// res.Close() // was absent
	return body

}
