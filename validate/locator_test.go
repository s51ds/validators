package validate

import "testing"

func TestLocator(t *testing.T) {
	type args struct {
		locator string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok-1",
			args: args{
				locator: "JN76to",
			},
			want: true,
		},
		{
			name: "nok-1",
			args: args{
				locator: " JN76to",
			},
			want: false,
		},
		{
			name: "nok-2",
			args: args{
				locator: "JN7 6to",
			},
			want: false,
		},
		{
			name: "nok-3",
			args: args{
				locator: "76JNto",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Locator(tt.args.locator); got != tt.want {
				t.Errorf("Locator() = %v, want %v", got, tt.want)
			}
		})
	}
}
