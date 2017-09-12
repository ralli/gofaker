package gofaker

import "testing"

func TestRegexify(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"remove leading slash and caret", args{"/^hase"}, "hase"},
		{"remove trailing dollar and slash", args{"hase$/"}, "hase"},
		{"all {2} become {2,2}", args{"{2}"}, "{2,2}"},
		{"all ? become {0,1}", args{"?"}, "{0,1}"},
		{"all [12]{1,1} become 1 or 2", args{"[12]{1,1}"}, "2"},
		{"all [12]{1,3}", args{"[12]{1,3}"}, "211"},
		{"all (aaa|bbb){1,1} become aaa or bbb", args{"(aaa|bbb){1,1}"}, "bbb"},
		{"all A{1,1} becomes A", args{"A{1,1}"}, "A"},
		{"all \\d{3} becomes \\d\\d\\d", args{"\\d{3}"}, "780"},
		{"all [A-Z,0-9] within [] becomes something in the range from A-Z", args{"[A-Z0-9]"}, "H"},
	}

	f, _ := NewFaker("en")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f.Reset()
			if got := f.Regexify(tt.args.expression); got != tt.want {
				t.Errorf("Regexify() = %v, want %v", got, tt.want)
			}
		})
	}
}
