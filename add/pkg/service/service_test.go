package service

import (
	"context"
	"testing"
)

var svc = new(basicAddService)
var ctx = context.Background()

func Test_basicAddService_Add(t *testing.T) {
	type args struct {
		ctx     context.Context
		numbers []float64
	}
	tests := []struct {
		name    string
		b       *basicAddService
		args    args
		wantRs  float64
		wantErr bool
	}{
		{"no numbers", svc, args{ctx, []float64{}}, 0, false},
		{"some int numbers", svc, args{ctx, []float64{1, 1}}, 2, false},
		{"some float numbers", svc, args{ctx, []float64{1.1, 1.1}}, 2.2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRs, err := tt.b.Add(tt.args.ctx, tt.args.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicAddService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRs != tt.wantRs {
				t.Errorf("basicAddService.Add() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}
