package main

func countKConstraintSubstrings(s string, k int) int {
	count := 0
	for start := 0; start < len(s); start++ {
		zeros, ones := 0, 0
		for end := start; end < len(s); end++ {
			if s[end] == '0' {
				zeros++
			} else {
				ones++
			}
			if zeros > k && ones > k {
				break
			}
			count++
		}
	}
	return count
}
