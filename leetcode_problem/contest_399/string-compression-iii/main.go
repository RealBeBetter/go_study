package string_compression_iii

import "strconv"

func compressedString(word string) string {
	comp := ""
	for i := 0; i < len(word); i++ {
		cnt := 1
		for i+1 < len(word) && word[i] == word[i+1] {
			if cnt == 9 {
				break
			}
			cnt++
			i++
		}
		comp += strconv.Itoa(cnt) + string(word[i])
	}

	return comp
}
