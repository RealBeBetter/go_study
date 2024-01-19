package sort

import "testing"

func TestMergeSort_Sort(t *testing.T) {
	type Case struct {
		Arg  []int
		Want []int
	}

	cases := []Case{
		Case{Arg: []int{1, 4, 3, 2, 5}, Want: []int{1, 2, 3, 4, 5}},
	}

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
