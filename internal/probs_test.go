package probs

import (
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name      string
		testSlice []int
		want      int
	}{
		{"Test1", []int{1, 2}, 3},
		{"Test2", []int{-1, 1}, 0},
		{"Test3", []int{0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.testSlice); got != tt.want {
				t.Errorf("Sum(%v) = %d; want %d", tt.testSlice, got, tt.want)
			}
		})
	}
}

func TestProbOne(t *testing.T) {
	tests := []struct {
		name   string
		counts int
		maxOf  int
		want   int
	}{
		{"Test1", 0, 10, 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProbOne(tt.counts, tt.maxOf); got != tt.want {
				t.Errorf("ProbOne(%d, %d) = %d; want %d", tt.counts, tt.maxOf, got, tt.want)
			}
		})
	}
}

func TestProbTwo(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  int
	}{
		{"Test1", 5, 2},
		{"Test2", 9, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProbTwo(tt.limit); got != tt.want {
				t.Errorf("ProbTwo(%d) = %d; want %d", tt.limit, got, tt.want)
			}
		})
	}
}
