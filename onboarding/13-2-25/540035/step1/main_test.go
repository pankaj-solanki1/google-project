package main

import "testing"

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		want   bool
	}{
		{
			name:   "Target found at even index",
			arr:    []int{1, 4, 5, 7, 9, 10},
			target: 5,
			want:   false,
		},
		{
			name:   "Target not found",
			arr:    []int{1, 4, 7, 9, 10},
			target: 3,
			want:   false,
		},
		{
			name:   "Ignore large numbers",
			arr:    []int{50, 51, 52, 53},
			target: 51,
			want:   false,
		},
		{
			name:   "Negative target",
			arr:    []int{1, 2, 3},
			target: -1,
			want:   false,
		},
		{
			name:   "Even sum",
			arr:    []int{1, 1, 2},
			target: 1,
			want:   false,
		},
		{
			name:   "Zero as target when sum < 10",
			arr:    []int{1, 2, 3},
			target: 0,
			want:   true,
		},
		{
			name:   "Zero as target when sum >= 10",
			arr:    []int{1, 2, 7},
			target: 0,
			want:   false,
		},
		{
			name:   "Target found at odd index",
			arr:    []int{1, 2, 3},
			target: 2,
			want:   true,
		},
		{
			name:   "Target found multiple times, even index first",
			arr:    []int{2, 1, 2, 3},
			target: 2,
			want:   true,
		},
		{
			name:   "Target found multiple times, odd index first",
			arr:    []int{1, 2, 2, 3},
			target: 2,
			want:   false,
		},
		{
			name:   "Empty array",
			arr:    []int{},
			target: 5,
			want:   false,
		},
		{
			name:   "Empty array, zero target, sum < 10",
			arr:    []int{},
			target: 0,
			want:   true,
		},
		{
			name:   "large number exists and is not the target",
			arr:    []int{1, 2, 3, 51},
			target: 3,
			want:   true,
		},
		{
			name:   "large number exists and is the target",
			arr:    []int{1, 2, 3, 51},
			target: 51,
			want:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := LinearSearch(tc.arr, tc.target)
			if got != tc.want {
				t.Errorf("LinearSearch(%v, %d) = %v, want %v", tc.arr, tc.target, got, tc.want)
			}
		})
	}
}
