package handlers

import (
	"encoding/json"
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
	"Arlington County": (*carInfo).arlingtonTaxCalculator,
	"Fairfax County":   (*carInfo).fairfaxTaxCalculator,
	"Alexandria City":  (*carInfo).alexandriaTaxCalculator,
}

func LocalitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	supportedLocalities := []locality{
		{Name: "Arlington County"},
		{Name: "Fairfax County"},
		{Name: "Alexandria City"},
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
}

// https://www.fairfaxcounty.gov/taxes/vehicles/vehicle-tax-subsidy
func (c *carInfo) fairfaxTaxCalculator() {
	const (
		taxRate          = .0457
		reliefUpperBound = 20000.0
		reliefRate       = .5
	)
	if c.Value <= reliefUpperBound {
		c.Taxes = (c.Value * taxRate) * reliefRate
	} else {
		c.Taxes = (c.Value * taxRate) - ((reliefUpperBound * taxRate) * reliefRate)
	}
}

// https://www.alexandriava.gov/CarTax
func (c *carInfo) alexandriaTaxCalculator() {
	const (
		taxRate              = .0533
		fullReliefAmount     = 5000.0
		boundOneRelief       = 20000.0
		boundOneReliefRate   = .48
		boundTwoRelief       = 25000.0
		boundTwoReliefRate   = .74
		boundThreeReliefRate = .88
	)
	if c.Value <= fullReliefAmount {
		c.Taxes = 0
	} else if c.Value <= boundOneRelief {
		c.Taxes = (c.Value * taxRate) * boundOneReliefRate
	} else if c.Value <= boundTwoRelief {
		// full rate for all above 20000
		c.Taxes = (c.Value - boundOneRelief) * taxRate
		// partial rate for first 20000
		c.Taxes += (boundOneRelief * taxRate) * boundTwoReliefRate
	} else {
		// full rate for all above 20000
		c.Taxes = (c.Value - boundOneRelief) * taxRate
		//partial rate for first 20000
		c.Taxes += (boundOneRelief * taxRate) * boundThreeReliefRate
	}
}
