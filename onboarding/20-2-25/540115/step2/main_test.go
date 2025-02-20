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
		name             string
		args             args
		wantUnion        []int
		wantIntersection []int
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
			wantUnion:        []int{1, 2, 3, 5, 6, 7, 8, 9},
			wantIntersection: []int{3, 5},
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
			wantUnion:        []int{4, 9, 12, 15, 18, 21, 24, 30},
			wantIntersection: []int{9, 15},
		},
		{
			name: "Example 3",
			args: args{
				arrs: [][]int{
					{10, 20, 30},
					{40, 50, 60},
				},
			},
			wantUnion:        []int{10, 20, 30, 40, 50, 60},
			wantIntersection: []int{},
		},
		{
			name: "Example 4",
			args: args{
				arrs: [][]int{
					{-5, -2, 0, 3},
					{-2, 1, 3, 4},
					{-2, 3, 5, 6},
				},
			},
			wantUnion:        []int{-5, -2, 0, 1, 3, 4, 5, 6},
			wantIntersection: []int{-2, 3},
		},
		{
			name: "Empty input",
			args: args{
				arrs: [][]int{},
			},
			wantUnion:        []int{},
			wantIntersection: []int{},
		},
		{
			name: "One empty array",
			args: args{
				arrs: [][]int{
					{1, 2, 3},
					{},
				},
			},
			wantUnion:        []int{1, 2, 3},
			wantIntersection: []int{},
		},
		{
			name: "All arrays empty",
			args: args{
				arrs: [][]int{
					{},
					{},
				},
			},
			wantUnion:        []int{},
			wantIntersection: []int{},
		},
		{
			name: "Duplicate elements within arrays",
			args: args{
				arrs: [][]int{
					{1, 2, 2, 3},
					{2, 3, 3, 4},
				},
			},
			wantUnion:        []int{1, 2, 3, 4},
			wantIntersection: []int{2, 3},
		},
		{
			name: "Large numbers",
			args: args{
				arrs: [][]int{
					{1000000, 2000000},
					{2000000, 3000000},
				},
			},
			wantUnion:        []int{1000000, 2000000, 3000000},
			wantIntersection: []int{2000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnion, gotIntersection := FindUnionAndIntersection(tt.args.arrs)
			if !reflect.DeepEqual(gotUnion, tt.wantUnion) {
				t.Errorf("FindUnionAndIntersection() gotUnion = %v, want %v", gotUnion, tt.wantUnion)
			}
			if !reflect.DeepEqual(gotIntersection, tt.wantIntersection) {
				t.Errorf("FindUnionAndIntersection() gotIntersection = %v, want %v", gotIntersection, tt.wantIntersection)
			}
		})
	}
}
