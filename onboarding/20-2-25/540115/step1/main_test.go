package main

import (
	"reflect"
	"testing"
)

func TestFindUnionAndIntersection(t *testing.T) {
	type args struct {
		arrs [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "Example 1",
			args: args{
				arrs: [][]int{
					{1, 3, 5, 7},
					{2, 3, 5, 8},
					{3, 5, 6, 9},
				},
			},
			want:  []int{1, 2, 3, 5, 6, 7, 8, 9},
			want1: []int{3, 5},
		},
		{
			name: "Example 2",
			args: args{
				arrs: [][]int{
					{4, 9, 12, 15},
					{9, 15, 18, 21},
					{9, 15, 24, 30},
				},
			},
			want:  []int{4, 9, 12, 15, 18, 21, 24, 30},
			want1: []int{9, 15},
		},
		{
			name: "Example with no intersection",
			args: args{
				arrs: [][]int{
					{10, 20, 30},
					{40, 50, 60},
				},
			},
			want:  []int{10, 20, 30, 40, 50, 60},
			want1: []int{},
		},
		{
			name: "Empty input",
			args: args{
				arrs: [][]int{},
			},
			want:  []int{},
			want1: []int{},
		},
		{
			name: "One empty array",
			args: args{
				arrs: [][]int{{1, 2, 3}, {}},
			},
			want:  []int{1, 2, 3},
			want1: []int{},
		},
		{
			name: "Duplicate values in input arrays",
			args: args{
				arrs: [][]int{{1, 2, 2, 3}, {2, 3, 3, 4}},
			},
			want:  []int{1, 2, 3, 4},
			want1: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindUnionAndIntersection(tt.args.arrs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUnionAndIntersection() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindUnionAndIntersection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
