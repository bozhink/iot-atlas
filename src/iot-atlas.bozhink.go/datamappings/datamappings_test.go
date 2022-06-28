package datamappings

import (
	"reflect"
	"testing"
	"time"

	"../apimodels"
	"../datamodels"
)

func TestGetDTO(t *testing.T) {
	type args struct {
		e *apimodels.EventRequestModel
	}
	tests := []struct {
		name string
		args args
		want *datamodels.EventDataModel
	}{
		{
			name: "GetDTO with nil event should return nil",
			args: args{
				e: nil,
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with nil readings should return nil",
			args: args{
				e: &apimodels.EventRequestModel{},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with empty readings should return nil",
			args: args{
				e: &apimodels.EventRequestModel{
					Readings: make([]apimodels.ReadingRequestModel, 0),
				},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with single nil reading should return nil",
			args: args{
				e: &apimodels.EventRequestModel{
					Readings: make([]apimodels.ReadingRequestModel, 1),
				},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with valid readings with single reading with invalid sensor should return nil",
			args: args{
				e: &apimodels.EventRequestModel{
					Sender:  "sender",
					Event:   "event",
					Version: "version",
					Readings: []apimodels.ReadingRequestModel{
						{
							Sensor:      "",
							Humidity:    0,
							Temperature: 0,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with valid readings with single reading with low temperature should return nil",
			args: args{
				e: &apimodels.EventRequestModel{
					Sender:  "sender",
					Event:   "event",
					Version: "version",
					Readings: []apimodels.ReadingRequestModel{
						{
							Sensor:      "sender",
							Humidity:    0,
							Temperature: -300,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with valid readings with single reading with low humidity should return nil",
			args: args{
				e: &apimodels.EventRequestModel{
					Sender:  "sender",
					Event:   "event",
					Version: "version",
					Readings: []apimodels.ReadingRequestModel{
						{
							Sensor:      "sender",
							Humidity:    -10,
							Temperature: 0,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with valid readings with single reading with high humidity should return nil",
			args: args{
				e: &apimodels.EventRequestModel{
					Sender:  "sender",
					Event:   "event",
					Version: "version",
					Readings: []apimodels.ReadingRequestModel{
						{
							Sensor:      "sender",
							Humidity:    110,
							Temperature: 0,
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "GetDTO with valid event with valid readings with single valid reading should return valid object",
			args: args{
				e: &apimodels.EventRequestModel{
					Sender:  "sender",
					Event:   "event",
					Version: "version",
					Readings: []apimodels.ReadingRequestModel{
						{
							Sensor:      "Sensor",
							Humidity:    0,
							Temperature: 0,
						},
					},
				},
			},
			want: &datamodels.EventDataModel{
				Sender:  "sender",
				Event:   "event",
				Version: "version",
				Date:    time.Now(),
				Readings: []datamodels.ReadingDataModel{
					{
						Sensor:      "Sensor",
						Humidity:    0,
						Temperature: 0,
						HeatIndex:   0,
						DewPoint:    0,
						Pressure:    0,
						Altitude:    0,
						DP:          0,
						Ps:          6.11,
						Pa:          0,
						HI:          -8.78,
					},
				},
			},
		},
		{
			name: "GetDTO with valid event with valid readings with one valid reading and one invalid should return valid object with single reading",
			args: args{
				e: &apimodels.EventRequestModel{
					Sender:  "sender",
					Event:   "event",
					Version: "version",
					Readings: []apimodels.ReadingRequestModel{
						{
							Sensor:      "Sensor",
							Humidity:    0,
							Temperature: 0,
						},
						{
							Sensor: "",
						},
					},
				},
			},
			want: &datamodels.EventDataModel{
				Sender:  "sender",
				Event:   "event",
				Version: "version",
				Date:    time.Now(),
				Readings: []datamodels.ReadingDataModel{
					{
						Sensor:      "Sensor",
						Humidity:    0,
						Temperature: 0,
						HeatIndex:   0,
						DewPoint:    0,
						Pressure:    0,
						Altitude:    0,
						DP:          0,
						Ps:          6.11,
						Pa:          0,
						HI:          -8.78,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDTO(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
