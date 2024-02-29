/**
 * @author Real
 * @since 2023/11/26 1:11
 */
package main

import (
	"io/ioutil"
	log2 "log"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var log = log2.Default()

var pattern = regexp.MustCompile(`(!)(\[).*?(])(\(https?://[^#]*)#.*?(\))`)

func main() {
	traverseDirectory("C:\\Users\\Real\\Documents\\yuque\\docs", ".md")
	// 运行时可以看到输出
	time.Sleep(1 * time.Second)
}

func replaceTextInFile(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return
	}

	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "<br />", "\n\n")
	contentStr = removeHashFromLinks(contentStr)

	err = ioutil.WriteFile(filePath, []byte(contentStr), 0600)
	if err != nil {
		log.Printf("Error writing to file %s: %v", filePath, err)
	}
}

func removeHashFromLinks(text string) string {
	result := pattern.ReplaceAllStringFunc(text, func(match string) string {
		groups := pattern.FindStringSubmatch(match)
		if len(groups) == 6 {
			return groups[1] + groups[2] + groups[3] + groups[4] + groups[5]
		}
		return match
	})
	return result
}

func traverseDirectory(path, fileExtension string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Error reading directory [%s]: %v", path, err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			traverseDirectory(filePath, fileExtension)
		} else if strings.HasSuffix(file.Name(), fileExtension) {
			log.Printf("File: %s", file.Name())
			replaceTextInFile(filePath)
		}
	}
}
