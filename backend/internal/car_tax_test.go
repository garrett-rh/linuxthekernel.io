package internal

import (
	"testing"
)

func Test_carInfo_alexandriaTaxCalculator(t *testing.T) {
	type fields struct {
		Value    float64
		Locality string
		Taxes    float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "tax free",
			fields: fields{
				Value:    2999.0,
				Locality: "Alexandria City",
				Taxes:    0.0,
			},
			want: 0.0,
		},
		{
			name: "bound one",
			fields: fields{
				Value:    15000.0,
				Locality: "Alexandria City",
				Taxes:    0.0,
			},
			want: 383.76,
		},
		{
			name: "bound two",
			fields: fields{
				Value:    24999.0,
				Locality: "Alexandria City",
				Taxes:    0.0,
			},
			want: 1055.2867,
		},
		{
			name: "full rate",
			fields: fields{
				Value:    30000.0,
				Locality: "Alexandria City",
				Taxes:    0.0,
			},
			want: 1471.08,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarInfo{
				Value:    tt.fields.Value,
				Locality: tt.fields.Locality,
				Taxes:    tt.fields.Taxes,
			}
			c.AlexandriaTaxCalculator()
			if got := c.Taxes; got != tt.want {
				t.Errorf("CarInfo.alexandriaTaxCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_carInfo_arlingtonTaxCalculator(t *testing.T) {
	type fields struct {
		Value    float64
		Locality string
		Taxes    float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "tax free",
			fields: fields{
				Value:    2999.0,
				Locality: "Arlington County",
				Taxes:    0,
			},
			want: 0.0,
		},
		{
			name: "under 20000",
			fields: fields{
				Value:    19999.0,
				Locality: "Arlington County",
				Taxes:    0,
			},
			want: 645.962,
		},
		{
			name: "above 20000",
			fields: fields{
				Value:    28325.0,
				Locality: "Arlington County",
				Taxes:    0,
			},
			want: 1062.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarInfo{
				Value:    tt.fields.Value,
				Locality: tt.fields.Locality,
				Taxes:    tt.fields.Taxes,
			}
			c.ArlingtonTaxCalculator()
			if got := c.Taxes; got != tt.want {
				t.Errorf("CarInfo.arlingtonTaxCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_carInfo_fairfaxTaxCalculator(t *testing.T) {
	type fields struct {
		Value    float64
		Locality string
		Taxes    float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "below 20000",
			fields: fields{
				Value:    10000.0,
				Locality: "Fairfax County",
				Taxes:    0,
			},
			want: 228.5,
		},
		{
			name: "above 20000",
			fields: fields{
				Value:    30000.0,
				Locality: "Fairfax County",
				Taxes:    0,
			},
			want: 914.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CarInfo{
				Value:    tt.fields.Value,
				Locality: tt.fields.Locality,
				Taxes:    tt.fields.Taxes,
			}
			c.FairfaxTaxCalculator()
			if got := c.Taxes; got != tt.want {
				t.Errorf("CarInfo.fairfaxTaxCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
