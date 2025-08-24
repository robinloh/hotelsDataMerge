package utils

import (
	"testing"
)

func TestTrimSpacesInSlices(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Success - Trim spaces from strings with leading/trailing spaces",
			args: args{
				s: []string{"  hello  ", " world ", "  test  "},
			},
			want: []string{"hello", "world", "test"},
		},
		{
			name: "Success - Trim spaces from strings with no spaces",
			args: args{
				s: []string{"hello", "world", "test"},
			},
			want: []string{"hello", "world", "test"},
		},
		{
			name: "Success - Trim spaces from strings with only spaces",
			args: args{
				s: []string{"   ", "  ", " "},
			},
			want: []string{"", "", ""},
		},
		{
			name: "Success - Empty slice",
			args: args{
				s: []string{},
			},
			want: []string{},
		},
		{
			name: "Success - Single string with spaces",
			args: args{
				s: []string{"  single  "},
			},
			want: []string{"single"},
		},
		{
			name: "Success - Mixed content",
			args: args{
				s: []string{"  no spaces", "with spaces  ", "  both sides  ", ""},
			},
			want: []string{"no spaces", "with spaces", "both sides", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSpacesInSlices(tt.args.s); len(got) != len(tt.want) {
				t.Errorf("TrimSpacesInSlices() length = %d, want %d", len(got), len(tt.want))
			} else {
				for i, str := range got {
					if str != tt.want[i] {
						t.Errorf("TrimSpacesInSlices()[%d] = %q, want %q", i, str, tt.want[i])
					}
				}
			}
		})
	}
}

func TestTrimSpacesInString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success - Trim leading and trailing spaces",
			args: args{
				s: "  hello world  ",
			},
			want: "hello world",
		},
		{
			name: "Success - Trim only leading spaces",
			args: args{
				s: "  hello world",
			},
			want: "hello world",
		},
		{
			name: "Success - Trim only trailing spaces",
			args: args{
				s: "hello world  ",
			},
			want: "hello world",
		},
		{
			name: "Success - No spaces to trim",
			args: args{
				s: "hello world",
			},
			want: "hello world",
		},
		{
			name: "Success - Only spaces",
			args: args{
				s: "   ",
			},
			want: "",
		},
		{
			name: "Success - Empty string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "Success - Single character with spaces",
			args: args{
				s: "  a  ",
			},
			want: "a",
		},
		{
			name: "Success - Mixed whitespace characters",
			args: args{
				s: " \t\n hello \t\n ",
			},
			want: "\t\n hello \t\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSpacesInString(tt.args.s); got != tt.want {
				t.Errorf("TrimSpacesInString() = %v, want %v", got, tt.want)
			}
		})
	}
}
