package gofaker

import "testing"

func TestRegexify(t *testing.T) {
	type args struct {
		expression string
		input      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Regexify(tt.args.expression, tt.args.input); got != tt.want {
				t.Errorf("Regexify() = %v, want %v", got, tt.want)
			}
		})
	}
}
