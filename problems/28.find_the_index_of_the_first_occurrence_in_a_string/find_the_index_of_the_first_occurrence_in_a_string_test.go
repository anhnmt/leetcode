package find_the_index_of_the_first_occurrence_in_a_string

import (
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "exam 1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "exam 2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr1(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "exam 1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "exam 2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr1(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
