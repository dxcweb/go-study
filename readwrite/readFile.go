package readwrite

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile 读文科文件
func ReadFile() {
	filename := "./temp/testwritefile.txt"
	readWithBufioReadString(filename)
}

func readWithBufioReadString(filename string) error {
	fileObj, err := os.Open(filename)
	defer fileObj.Close()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(fileObj)
	result, err := reader.ReadString(byte('@'))
	fmt.Println(result)
	return err
}
