package main

import "testing"

func Test_numDifferentIntegers(t *testing.T) {
	type args struct {
		word string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "non-leading zeros number",
			args: args{
				word: "a123bc34d8ef34",
			},
			want: 3,
		},
		{
			name: "leading zeros number 01",
			args: args{
				word: "A1b01c001",
			},
			want: 1,
		},
		{
			name: "leading zeros number 02",
			args: args{
				word: "A1b01c101",
			},
			want: 2,
		},
		{
			name: "leading zeros number 03",
			args: args{
				word: "A1b01c101d000",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDifferentIntegers(tt.args.word); got != tt.want {
				t.Errorf("NumDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
