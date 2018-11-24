package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)


var root = ""

func getHash(path string) string {
	f, _ := os.Open(path)
	hash := md5.New()
	io.Copy(hash, f)
	return hex.EncodeToString(hash.Sum(nil))
}

func deduplicate() []string {
	tempMap := make(map[string]byte)
	result := []string{}
	f, _ := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return err
		}
		if filepath.Ext(path) == ".jpg" {
			l := len(tempMap)
			tempMap[getHash(path)] = 0
			if len(tempMap) == l {
				result = append(result, path)
				log.Println(path)
				writer.WriteString(path)
				writer.WriteString("\n")
			}
		}
		return err

	})
	return result
}

func test() {
	// getHash("helloworld")
	// fmt.Println(getHash("helloworld"))
	// fmt.Println(filepath.Ext(root))
	f, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	writer := bufio.NewWriter(f)
	writer.WriteString("hello")
	writer.WriteString("hello")
	writer.Flush()

}

func main() {
	start := time.Now()
	result := deduplicate()
	log.Println("发现重复文件共：", len(result), "个")
	// test()
	end := time.Now()
	log.Println(end.Sub(start))
}
