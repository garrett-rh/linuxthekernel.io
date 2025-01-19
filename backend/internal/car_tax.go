package internal

type CarInfo struct {
	// yea i know you aren't supposed to use floats for money but im feeling lazy
	Value    float64 `json:"value"`
	Locality string  `json:"locality"`
	Taxes    float64 `json:"taxes"`
}

// https://www.arlingtonva.us/Government/Programs/Taxes/Vehicles/Vehicle-Tax-Relief
// see the above link for where these numbers come from
func (c *CarInfo) ArlingtonTaxCalculator() {
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
func (c *CarInfo) FairfaxTaxCalculator() {
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
func (c *CarInfo) AlexandriaTaxCalculator() {
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
