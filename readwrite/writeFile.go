package readwrite

import (
	"bufio"
	"io/ioutil"
	"os"
)

// WriteFile 写文件
func WriteFile() {
	filename := "./temp/testwritefile.txt"
	content := "你好，世界\ndxcweb3"
	// writeWithIoutil(filename, content)
	// writeWithFileWriteString(filename, content)
	// writeWithFileWrite(filename, content)
	// writeWithFileWriteAt(filename, content, 1)
	writeWithBufio(filename, content)
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func writeWithIoutil(filename, content string) error {
	contents := []byte(content)
	err := ioutil.WriteFile(filename, contents, 0644)
	return err
}

// const (
// 	O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
// 	O_WRONLY int = syscall.O_WRONLY // 只写打开文件
// 	O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
// 	O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
// 	O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
// 	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
// 	O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
// 	O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
// )

//使用os.OpenFile() 打开文件对象，并使用WriteString方法进行文件写入操作
func writeWithFileWriteString(filename, content string) error {
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer fileObj.Close()
	if err != nil {
		return err
	}
	_, err = fileObj.WriteString(content)
	return err
	// contents := []byte(content)
	// if _, err := fileObj.Write(contents); err == nil {
	// 	fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.", content)
	// }
}

//使用os.OpenFile() 打开文件对象，并使用WriteString方法进行文件写入操作
func writeWithFileWrite(filename, content string) error {
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer fileObj.Close()
	if err != nil {
		return err
	}
	contents := []byte(content)
	_, err = fileObj.Write(contents)
	return err
}

//使用os.OpenFile() 打开文件对象，并使用WriteString方法进行文件写入操作
func writeWithFileWriteAt(filename, content string, offset int64) error {
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	defer fileObj.Close()
	if err != nil {
		return err
	}
	b := []byte(content)
	_, err = fileObj.WriteAt(b, offset)
	return err
}

//使用bufio包中Writer对象的相关方法进行数据的写入
func writeWithBufio(filename, content string) error {
	fileObj, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fileObj.Close()
	if err != nil {
		return err
	}
	writeObj := bufio.NewWriterSize(fileObj, 4096)

	if _, err = writeObj.WriteString(content); err != nil {
		return err
	}
	if err := writeObj.Flush(); err != nil {
		return err
	}
	return nil
}
