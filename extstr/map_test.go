package extstr

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		ss      []string
		mapping func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"empty",
			args{
				[]string{},
				func(s string) string {
					return s + "1"
				},
			},
			[]string{},
		},
		{
			"map",
			args{
				[]string{"a", "b", "c"},
				func(s string) string {
					return s + "1"
				},
			},
			[]string{"a1", "b1", "c1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.ss, tt.args.mapping); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapx(t *testing.T) {
	type args struct {
		ss      []string
		mapping func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"empty",
			args{
				[]string{},
				func(s string) string {
					return s + "1"
				},
			},
			[]string{},
		},
		{
			"map",
			args{
				[]string{"a", "b", "c"},
				func(s string) string {
					return s + "1"
				},
			},
			[]string{"a1", "b1", "c1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.ss, tt.args.mapping); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
