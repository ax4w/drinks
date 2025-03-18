package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type drink struct {
	Name         string
	Glass        string
	Instructions string
	Alcoholic    bool
	Ingredients  map[string]string
}

func getRandomDrink() (drink, bool) {
	var (
		response, err = http.Get("https://www.thecocktaildb.com/api/json/v1/1/random.php")
		r             = make(map[string]any)
		d             = drink{
			Ingredients: make(map[string]string),
		}
	)
	if err != nil {
		println(err.Error())
		return drink{}, false
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&r); err != nil {
		println(err.Error())
		return drink{}, false
	}
	if len(r["drinks"].([]interface{})) == 0 {
		return drink{}, false
	}
	r = r["drinks"].([]interface{})[0].(map[string]any)
	d.Name = r["strDrink"].(string)
	d.Glass = r["strGlass"].(string)
	d.Alcoholic = strings.ToLower(r["strAlcoholic"].(string)) == "alcoholic"
	d.Instructions = r["strInstructions"].(string)

	// api spec is 1 to 15
	for i := 1; i <= 15; i++ {
		var (
			ingredient  = r[fmt.Sprintf("strIngredient%d", i)]
			measurement = r[fmt.Sprintf("strMeasure%d", i)]
		)
		if ingredient == nil {
			break
		}
		if measurement == nil {
			d.Ingredients[ingredient.(string)] = ""
		} else {
			d.Ingredients[ingredient.(string)] = measurement.(string)
		}

	}
	return d, true

}
