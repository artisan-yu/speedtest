package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:',.<>?/`~\"\\ "

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "用法: %s <大小MB>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "示例: %s 100\n", os.Args[0])
		os.Exit(1)
	}

	sizeMB, err := strconv.Atoi(os.Args[1])
	if err != nil || sizeMB <= 0 {
		fmt.Fprintf(os.Stderr, "错误: 请提供正整数作为大小参数\n")
		os.Exit(1)
	}

	totalBytes := int64(sizeMB) * 1024 * 1024
	filename := fmt.Sprintf("%dmb.txt", sizeMB)

	fmt.Printf("正在生成 %s (%d MB) ...\n", filename, sizeMB)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "创建文件失败: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 1024*1024)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	buf := make([]byte, 64*1024)
	var written int64

	for written < totalBytes {
		remaining := totalBytes - written
		batch := int64(len(buf))
		if remaining < batch {
			batch = remaining
		}

		for i := int64(0); i < batch; i++ {
			buf[i] = charset[rng.Intn(len(charset))]
		}

		n, err := writer.Write(buf[:batch])
		if err != nil {
			fmt.Fprintf(os.Stderr, "写入失败: %v\n", err)
			os.Exit(1)
		}
		written += int64(n)
	}

	if err := writer.Flush(); err != nil {
		fmt.Fprintf(os.Stderr, "刷新缓冲区失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("完成: %s 已生成 (%d 字节)\n", filename, written)
}
