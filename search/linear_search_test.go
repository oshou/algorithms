package linear_search

import "testing"

func Test_linearSearch(t *testing.T) {
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
			args: args{arr: []int{1, 2, 3}, v: 3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linearSearch(tt.args.arr, tt.args.v); got != tt.want {
				t.Errorf("linearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
