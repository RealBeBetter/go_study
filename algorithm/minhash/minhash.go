package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	prime = 1000000007 // 一个大的素数
	m     = 1000000    // 输出范围
)

// HashFunc 定义一个哈希函数类型
type HashFunc func(string) int

// GenerateHashFuncs 生成一组哈希函数
func GenerateHashFuncs(num int) []HashFunc {
	rand.Seed(24)
	hashFunctions := make([]HashFunc, num)
	for i := range hashFunctions {
		a := rand.Int()
		b := rand.Int()
		hashFunctions[i] = func(s string) int {
			var h int64
			for _, c := range s {
				h = ((h*int64(a) + int64(c)) * int64(b)) % int64(prime)
			}
			return int(h % int64(m))
		}
	}
	return hashFunctions
}

// MinHashSignatures 计算每个集合的MinHash签名
func MinHashSignatures(shingles []string, hashFuncs []HashFunc) []int {
	signatures := make([]int, len(hashFuncs))
	for i, hf := range hashFuncs {
		minHash := hf(shingles[0])
		for _, shingle := range shingles[1:] {
			hashValue := hf(shingle)
			if hashValue < minHash {
				minHash = hashValue
			}
		}
		signatures[i] = minHash
	}
	return signatures
}

// CalculateJaccardSimilarity 计算两个集合之间的Jaccard相似度
func CalculateJaccardSimilarity(signatures1, signatures2 []int) float64 {
	matchingCount := 0
	for i := range signatures1 {
		if signatures1[i] == signatures2[i] {
			matchingCount++
		}
	}
	return float64(matchingCount) / float64(len(signatures1))
}

// Shingle 将文本转换为shingles
func Shingle(text string, n int) []string {
	var shingles []string
	for i := 0; i <= len(text)-n; i++ {
		shingles = append(shingles, text[i:i+n])
	}
	return shingles
}

func main() {
	text1 := "The quick brown fox jumps over the lazy dog."
	text2 := "The quick brown dog jumps over the lazy cat."

	// 将文本转换为shingles
	shingles1 := Shingle(strings.ToLower(text1), 3)
	shingles2 := Shingle(strings.ToLower(text2), 3)

	// 生成哈希函数
	hashFunctions := GenerateHashFuncs(100)

	// 计算MinHash签名
	signatures1 := MinHashSignatures(shingles1, hashFunctions)
	signatures2 := MinHashSignatures(shingles2, hashFunctions)

	// 计算Jaccard相似度
	similarity := CalculateJaccardSimilarity(signatures1, signatures2)
	// 这里输出的 Jaccard 相似度控制在 0.7 左右
	fmt.Printf("Estimated Jaccard similarity: %.2f\n", similarity)
}
