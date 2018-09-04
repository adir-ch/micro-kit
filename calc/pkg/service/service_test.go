package service

import (
	"reflect"
	"testing"
)

func Test_splitExpr(t *testing.T) {
	type args struct {
		expr      string
		seperator string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"simple1", args{"1.1 + 2.2", "+"}, []float64{1.1, 2.2}},
		{"simple2", args{"1.1 + 2", "+"}, []float64{1.1, 2.0}},
		{"simple3", args{"1 + 2.2", "+"}, []float64{1.0, 2.2}},
		{"no input", args{"", "+"}, []float64{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitExpr(tt.args.expr, tt.args.seperator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
