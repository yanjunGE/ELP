package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 发送文件到服务端
func SendFile(filePath string, fileSize int64, conn net.Conn) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var count int64
	for {
		buf := make([]byte, 2048)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Println("文件传输完成")
			//告诉服务端结束文件接收
			//conn.Write([]byte("ok"))
			return
		}
		//发送给服务端
		conn.Write(buf[:n])

		count += int64(n)
		sendPercent := float64(count) / float64(fileSize) * 100
		value := fmt.Sprintf("%.2f", sendPercent)
		//打印上传进度
		fmt.Println("文件上传：" + value + "%")
	}
}
func Handler(conn net.Conn) {
	buf := make([]byte, 2048)
	//读取客户端发送的内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	resFileName := string(buf[:n])

	//获取客户端ip+port
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + ": 文件被存储为为--" + resFileName)
	//告诉客户端已经接收到文件名
	conn.Write([]byte("ok"))
	//创建文件
	f, err := os.Create(resFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//循环接收客户端传递的文件内容
	buf = make([]byte, 2048)
	for {
		fmt.Print(buf[:n])
		fmt.Print("1")
		n, err4 := conn.Read(buf)
		if err4 != nil {
			if err4 == io.EOF {
				fmt.Println("文件接受完毕")
				//i = NO_ROUTE
				fmt.Println(addr + ": 协程结束")
				//runtime.Goexit()
				break

			} else {
				fmt.Println("conn.Read err", err4)
			}
		}
		/*fmt.Print(buf[:n])
		fmt.Print("1")*/
		//结束协程
		/*if string(buf[:n]) == "ok" {
			fmt.Println(addr + ": 协程结束")
			runtime.Goexit()
		}*/
		f.Write(buf[:n])
	}
	//defer conn.Close()
	fmt.Println("fclose")
	defer f.Close()
	fmt.Println("close")

}

func main() {
	fmt.Print("请输入文件的完整路径：")
	//创建切片，用于存储输入的路径
	var str string
	fmt.Scan(&str)
	//获取文件信息
	fileInfo, err := os.Stat(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建客户端连接
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//文件名称
	fileName := fileInfo.Name()
	//文件大小
	fileSize := fileInfo.Size()
	//发送文件名称到服务端
	conn.Write([]byte(fileName))
	buf := make([]byte, 2048)
	//读取服务端内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf[:n])
	if revData == "ok" {
		//发送文件数据
		SendFile(str, fileSize, conn)
	}
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		//阻塞等待客户端
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//创建协程
		defer conn.Close()
		Handler(conn)

	}
}
