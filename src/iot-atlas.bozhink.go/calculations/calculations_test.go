package calculations

import (
	"math"
	"testing"
)

const epsd float64 = 1e-8

func TestGetDewPoint(t *testing.T) {
	type args struct {
		t float64
		h float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Dew point on H = 0% should be 0",
			args: args{t: 10, h: 0},
			want: 0,
		},
		{
			name: "Dew point on H < 0% should be 0",
			args: args{t: 10, h: -20},
			want: 0,
		},
		{
			name: "Dew point on H > 100% should be 0",
			args: args{t: 10, h: 120},
			want: 0,
		},
		{
			name: "Dew point on H = 100% should be 10",
			args: args{t: 10, h: 100},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDewPoint(tt.args.t, tt.args.h); math.Abs(got-tt.want) > epsd {
				t.Errorf("GetDewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHeatIndex(t *testing.T) {
	type args struct {
		t float64
		h float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Heat index on t = 0 °C and H = 0% should be c1",
			args: args{t: 0, h: 0},
			want: c1d,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHeatIndex(tt.args.t, tt.args.h); math.Abs(got-tt.want) > epsd {
				t.Errorf("GetHeatIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
