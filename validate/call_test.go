package validate

import "testing"

func TestCallSign(t *testing.T) {
	type args struct {
		callSign string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty-FALSE",
			args: args{
				callSign: "",
			},
			want: false,
		},
		{
			name: "space-FALSE",
			args: args{
				callSign: " ",
			},
			want: false,
		},

		{
			name: "spaces-FALSE",
			args: args{
				callSign: "     ",
			},
			want: false,
		},

		{
			name: "TooShort-1-FALSE",
			args: args{
				callSign: "ab",
			},
			want: false,
		},
		{
			name: "a /-FALSE",
			args: args{
				callSign: "a /",
			},
			want: false,
		},
		{
			name: "abc-FALSE",
			args: args{
				callSign: "abc",
			},
			want: false,
		},
		{
			name: "E2E",
			args: args{
				callSign: "E2E",
			},
			want: true,
		},
		{
			name: "s59 abc-FALSE",
			args: args{
				callSign: "s59 abc",
			},
			want: false,
		},
		{
			name: "9A/S51DS",
			args: args{
				callSign: "9A/S51DS",
			},
			want: true,
		},
		{
			name: "9A/S51DS/P",
			args: args{
				callSign: "9A/S51DS/P",
			},
			want: true,
		},
		{
			name: "9A/S51DS--FALSE",
			args: args{
				callSign: "9A/S51DS//P",
			},
			want: false,
		},
		{
			name: "9A/S51DS-1--FALSE",
			args: args{
				callSign: "9A/S51DS/",
			},
			want: false,
		},
		{
			name: "S59ABC0--FALSE",
			args: args{
				callSign: "S59ABC0",
			},
			want: false,
		},
		{
			name: "S59ABC9--FALSE",
			args: args{
				callSign: "S59ABC9",
			},
			want: false,
		},
		{
			name: "S53O",
			args: args{
				callSign: "S53O",
			},
			want: true,
		},
		{
			name: "13MEK",
			args: args{
				callSign: "13MEK",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallSign(tt.args.callSign); got != tt.want {
				t.Errorf("CallSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
