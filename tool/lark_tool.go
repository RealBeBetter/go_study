/**
 * @author Real
 * @since 2023/11/26 1:11
 */
package main

import (
	"flag"
	"fmt"
	log2 "log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var log = log2.Default()

var pattern = regexp.MustCompile(`(!)(\[).*?(])(\(https?://[^#]*)#.*?(\))`)

func main() {
	var (
		dirFlag       string
		extFlag       string
		recursiveFlag bool
	)

	flag.StringVar(&dirFlag, "dir", "", "directory to scan (default: ~/Downloads if exists, otherwise working dir)")
	flag.StringVar(&extFlag, "ext", ".md", "file extension to process")
	flag.BoolVar(&recursiveFlag, "recursive", false, "scan sub-directories")
	flag.Parse()

	// Keep old behavior as a default, but allow -recursive to override.
	if !recursiveFlag {
		switch runtime.GOOS {
		case "darwin":
			fmt.Println("当前操作系统是: macOS, 默认不进行递归处理（可通过 -recursive 开启）")
		case "windows":
			recursiveFlag = true
			fmt.Println("当前操作系统是: Windows, 默认进行递归处理")
		default:
			fmt.Printf("当前操作系统是: %s, 默认不进行递归处理（可通过 -recursive 开启）\n", runtime.GOOS)
		}
	}

	scanDir := resolveScanDir(dirFlag)
	fmt.Printf("working dir: %s\n", mustGetwd())
	fmt.Printf("executable dir: %s\n", mustExeDir())
	fmt.Printf("最终扫描目录: %s\n", scanDir)

	traverseDirectory(scanDir, extFlag, recursiveFlag)
	// 运行时可以看到输出
	time.Sleep(1 * time.Second)
}

func resolveScanDir(dirFlag string) string {
	if strings.TrimSpace(dirFlag) != "" {
		d := strings.TrimSpace(dirFlag)
		if abs, err := filepath.Abs(d); err == nil {
			d = abs
		}
		return filepath.Clean(d)
	}

	// Prefer ~/Downloads if present (works for both source-run and packaged binary).
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		downloads := filepath.Join(home, "Downloads")
		if st, err := os.Stat(downloads); err == nil && st.IsDir() {
			if abs, err := filepath.Abs(downloads); err == nil {
				return abs
			}
			return downloads
		}
	}

	// Fallback: working directory.
	if wd, err := os.Getwd(); err == nil {
		if abs, err := filepath.Abs(wd); err == nil {
			return abs
		}
		return wd
	}

	// Last resort: executable dir.
	return mustExeDir()
}

func mustGetwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return "(unknown)"
	}
	if abs, err := filepath.Abs(wd); err == nil {
		return abs
	}
	return wd
}

func mustExeDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "(unknown)"
	}
	exe, _ = filepath.EvalSymlinks(exe)
	return filepath.Dir(exe)
}

func replaceTextInFile(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return
	}

	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "<br />", "\n\n")
	contentStr = removeHashFromLinks(contentStr)

	err = os.WriteFile(filePath, []byte(contentStr), 0600)
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

func traverseDirectory(path, fileExtension string, recursion bool) {
	fmt.Printf("正在扫描目录: %s\n", path)

	files, err := os.ReadDir(path)
	if err != nil {
		log.Printf("Error reading directory [%s]: %v", path, err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() && recursion {
			traverseDirectory(filePath, fileExtension, recursion)
		} else if strings.HasSuffix(file.Name(), fileExtension) {
			log.Printf("File: %s", file.Name())
			replaceTextInFile(filePath)
		}
	}
}
