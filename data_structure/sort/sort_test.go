package sort

import "testing"

type Case struct {
	Arg  []int
	Want []int
}

var cases []Case

func init() {
	cases = []Case{
		{Arg: []int{1, 4, 3, 2, 5}, Want: []int{1, 2, 3, 4, 5}},
		{Arg: []int{1, 2, 3, 4, 8, 7, 9, 5, 6}, Want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
}

func TestMergeSort_Sort(t *testing.T) {
	mergeSort := MergeSort{}
	for _, curCase := range cases {
		mergeSort.Sort(curCase.Arg)
		want := curCase.Want
		for index, value := range curCase.Arg {
			if want[index] != value {
				t.Fatal("sort result error")
			}
		}
	}
}

func TestQuickSort_Sort(t *testing.T) {
	quickSort := QuickSort{}
	for _, curCase := range cases {
		quickSort.Sort(curCase.Arg)
		want := curCase.Want
		for index, value := range curCase.Arg {
			if want[index] != value {
				t.Fatal("sort result error")
			}
		}
	}
}
