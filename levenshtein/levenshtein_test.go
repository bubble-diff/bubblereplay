package levenshtein

import "testing"

func Test_levenshtein(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "horse->ros",
			args: args{"horse", "ros"},
			want: 3,
		},
		{
			name: "intention->execution",
			args: args{"intention", "execution"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levenshtein(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compute("{\"ping\": \"pong\"}", "{\"ping\": \"bong\"}")
	}
}
