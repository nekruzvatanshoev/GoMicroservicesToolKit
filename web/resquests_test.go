package web

import (
	"encoding/json"
	"reflect"
	"testing"
)



func TestGetRequest(t *testing.T) {
	t.Run("testing Marcadolibre API for US", func(t *testing.T) {
		var country Country

		url :=  "https://api.mercadolibre.com/countries/US"
		bytes, err := GetRequest(url)
		if err != nil {
			t.Log(err)
		}

		if err != nil {
			t.Log(err)
		}

		err = json.Unmarshal(bytes, &country)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("Id:",country.Id)
		t.Log("Name:",country.Name)
		t.Log("CurrencyID:",country.CurrencyId)
		t.Log("Decimal:",country.DecimalSeparator)
		t.Log("Thousands:",country.ThousandsSeparator)
		t.Log("Locale:",country.Locale)
		t.Log("Timezone:",country.TimeZone)
		t.Log("GeoInformation:",country.GeoInformation.Location.Latitude)
		t.Log("States",country.States)


		if reflect.DeepEqual(country, US) == true {
			t.Log("JSON responses match!")
		} else {
			t.Error("JSON responses don't match! Expected: ", US, " Received: ", country)
		}

	})

}
