package handlers

import (
	"encoding/json"
	"math"
	"net/http"
)

type carInfo struct {
	// yea i know you aren't supposed to use floats for money but im feeling lazy
	Value    float64 `json:"value"`
	Locality string  `json:"locality"`
	Taxes    float64 `json:"taxes"`
}

type locality struct {
	Name string `json:"name"`
}

var localityMapping = map[string]func(*carInfo){
	"Arlington": (*carInfo).arlingtonTaxCalculator,
}

func LocalitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	supportedLocalities := []locality{
		{Name: "Arlington"},
	}
	err := json.NewEncoder(w).Encode(supportedLocalities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CarTaxHandler(w http.ResponseWriter, r *http.Request) {
	var car carInfo
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if fn := localityMapping[car.Locality]; fn != nil {
		fn(&car)
	}
	err = json.NewEncoder(w).Encode(car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// https://www.arlingtonva.us/Government/Programs/Taxes/Vehicles/Vehicle-Tax-Relief
// see the above link for where these numbers come from
func (c *carInfo) arlingtonTaxCalculator() {
	// generic tax rate before relief
	const (
		taxRate                   = .05
		forgivenessRateUnder17000 = .76
		fullReliefAmount          = 3000.0
		partialReliefAmount       = 17000.0
	)
	remainingValue := c.Value

	// everything under $3000 is forgiven
	remainingValue -= fullReliefAmount
	if remainingValue <= 0 {
		c.Taxes = 0
		return
	}
	// from $3001 - $20000 the forgiveness rate is 24% on 5% of the car

	// if value is less than 17000 then we know we don't hit the next bracket
	if remainingValue-partialReliefAmount <= 0 {
		c.Taxes = (remainingValue * taxRate) * forgivenessRateUnder17000
	} else {
		remainingValue -= partialReliefAmount
		c.Taxes = (partialReliefAmount * taxRate) * forgivenessRateUnder17000

		// remaining amount is unforgiven
		c.Taxes += remainingValue * taxRate
	}

	// Round to the nearest whole number
	c.Taxes = math.Trunc(c.Taxes)
}
