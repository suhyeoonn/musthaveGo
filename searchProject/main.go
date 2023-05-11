package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {
	lineNo int
	line   string
}

// 파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요함. ex) ex26.1 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------")
	}

}

func FindWordInAllFiles(word, path string) (findInfos []FindInfo) {
	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일을 찾을 수 없음 :", err)
		return
	}
	fmt.Println("filelist", filelist)
	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
		filepath.Walk(path, FindWordInFile(word, filename, ch))

	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return
}
func GetFileList(path string) ([]string, error) {
	filelist := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			matched, _ := filepath.Match(pattern, info.Name())
			if matched {
				filelist = append(filelist, path)
			}
		}
		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return filelist, nil
	//return filepath.Glob(path)
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없음", filename)
		ch <- findInfo
		return
	}
	defer file.Close() // ⚠️ 연 파일은 꼭 닫기

	scanner := bufio.NewScanner(file) // 스캐너 생성해서 한 줄씩 읽기
	line := 0
	for scanner.Scan() {
		line++
		if strings.Contains(scanner.Text(), word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo: line, line: scanner.Text()})
		}
	}

	ch <- findInfo
}
