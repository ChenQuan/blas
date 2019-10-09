package goblas

import "testing"

func TestSnrm2(t *testing.T) {
	type args struct {
		N    int
		x    []float32
		incX int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1",
			args: args{
				N:    2,
				x:    []float32{1, 2},
				incX: 1,
			},
			want: 2.236068,
		},
		{
			name: "2",
			args: args{
				N:    2,
				x:    []float32{2, 2},
				incX: 1,
			},
			want: Sqrt(2 * 2 * 2),
		}, {
			name: "3",
			args: args{
				N:    2,
				x:    []float32{2, 4},
				incX: 1,
			},
			want: Sqrt(2*2 + 4*4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snrm2(tt.args.N, tt.args.x, tt.args.incX); got != tt.want {
				t.Errorf("Snrm2() = %v, want %v", got, tt.want)
			}
		})
	}
}
