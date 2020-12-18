package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v \n", err.Error())
		return
	}
	/**
	func NewReader(rd io.Reader) *Reader
	NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
	*/
	inputReader := bufio.NewReader(os.Stdin)
	for {
		/**
		func (b *Reader) ReadString(delim byte) (line string, err error)
		ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
		如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
		当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
		*/
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err : %v \n", err)
			break
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}

		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed, err : %v \n", err)
			break
		}
	}
}
