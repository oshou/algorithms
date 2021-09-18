package binary_search

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr []int
		v   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{arr: []int{1, 2, 3}, v: 2},
			want: true,
		},
		{
			name: "case2",
			args: args{arr: []int{1, 2, 3}, v: 4},
			want: false,
		},
		{
			name: "case3",
			args: args{arr: []int{2, 3, 5, 7, 11, 13, 17}, v: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
