package util

import "testing"

func TestValidSign(t *testing.T) {
	type args struct {
		param map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only string value",
			args: args{param: map[string]interface{}{"lastName": "chen", "firstName": "deo", "sign": "f11795d19e7c1c8fedf9b47ea91a7132"}},
			want: true,
		},
		{
			name: "have int value",
			args: args{param: map[string]interface{}{"lastName": "chen", "firstName": "deo", "age": 20, "sign": "30ec5fe0a1aca7ff46bad2250250f692"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidSign(tt.args.param); got != tt.want {
				t.Errorf("ValidSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
