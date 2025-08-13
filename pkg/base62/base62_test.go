package base62

import (
	"testing"
)

func TestInt2String(t *testing.T) {
	tests := []struct {
		name string
		args uint64
		want string
	}{
		{name: "0", args: 0, want: "0"},
		{name: "1", args: 1, want: "1"},
		{name: "62", args: 62, want: "10"},
		{name: "6347", args: 6347, want: "1En"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.args); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantSeq uint64
	}{
		{name: "0", args: args{s: "0"}, wantSeq: 0},
		{name: "1", args: args{s: "1"}, wantSeq: 1},
		{name: "10", args: args{s: "10"}, wantSeq: 62},
		{name: "1En", args: args{s: "1En"}, wantSeq: 6347},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSeq := String2Int(tt.args.s); gotSeq != tt.wantSeq {
				t.Errorf("String2Int() = %v, want %v", gotSeq, tt.wantSeq)
			}
		})
	}
}
