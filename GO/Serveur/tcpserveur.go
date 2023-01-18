package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const NO_ROUTE int = 99999

func getLength(n int, length []int, last []int, matrice [][]int, s int) ([]int, []int) {
	//func getLength(n int, matrice [][]int, s int) ([]int, []int) {
	//n := len(matrice)
	//last := []int{9999, 99999, 99999}
	//length := []int{0, 99999, 99999}
	for i := 0; i < n; i++ {
		if matrice[s][i] != NO_ROUTE {
			if length[s]+matrice[s][i] < length[i] {
				last[i] = s
				length[i] = length[s] + matrice[s][i]
			}
		}
	}
	return length, last
}
func allPass(m []bool) bool {
	for i := 0; i < len(m); i++ {
		if m[i] == false {
			return false
		}
	}
	return true
}

func dijistra(Nb_sommets int, tab [][]int, src int) ([]int, []int) {
	//remplacer -1 dans le matrice par NO_ROUTE
	//Nb_sommets := int(len(tab))
	for i := 0; i < Nb_sommets; i++ {
		for j := 0; j < Nb_sommets; j++ {
			if tab[i][j] == -1 {
				tab[i][j] = NO_ROUTE
			}
		}
	}
	//le length de route est infini au d'abord sauf src est 0
	length_de_route := make([]int, Nb_sommets)
	route_pass := make([]int, Nb_sommets)
	for i := 0; i < Nb_sommets; i++ {
		length_de_route[i] = NO_ROUTE
		route_pass[i] = NO_ROUTE
	}
	length_de_route[src] = 0
	//判断哪些点被经过了。noter si les sommets sont passe
	sommets_pass := make([]bool, Nb_sommets)
	for i := 0; i < Nb_sommets; i++ {
		sommets_pass[i] = false
	}
	var next_src int
	var min int
	for allPass(sommets_pass) == false {
		min = NO_ROUTE
		for i := 0; i < Nb_sommets; i++ {
			if min > length_de_route[i] && sommets_pass[i] == false {
				min = length_de_route[i]
				next_src = i

			}
		}
		print("source:", next_src)
		length_de_route, route_pass = getLength(Nb_sommets, length_de_route, route_pass, tab, next_src)
		//length_de_route, route_pass = getLength(Nb_sommets, tab, next_src)
		sommets_pass[next_src] = true
	}
	fmt.Print(sommets_pass)
	return length_de_route, route_pass
}

func openfile(fN string) ([][]int, int, int) {
	//var n int
	var src int
	src = 0
	file, err := os.Open(fN) // For read access.
	/*var arr2 = make([][]int, n)
	for j := 0; j < n; j++ {
		arr2[j] = make([]int, n)
	}*/

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	str1, _ := reader.ReadString('\n')
	arr0 := strings.Fields(str1)
	fmt.Print("arr0:", arr0)
	var n = len(arr0)
	fmt.Print("le nombre de sommet est", len(arr0))
	var arr2 = make([][]int, n)

	for j := 0; j < n; j++ {
		arr2[j] = make([]int, n)
	}
	for {
		fmt.Println("Merci d'entre le source de", fN, ":")
		//fmt.Scanln(&src)
		//src = 0
		fmt.Print(src)
		if src < 0 || src >= n {
			fmt.Println("source n'existe pas, merci de entrer un nombre entre 0 et", n-1)
			os.Exit(1)
		} else {
			break
		}
	}
	for j := 0; j < n; j++ {
		//arr2[i] = make([]int, n)
		arr2[0][j], _ = strconv.Atoi(arr0[j])
		//print("i:", i, "j:", j, "arr:", arr2[i][j], "\n")
		fmt.Println(arr2)
	}
	//var arr2 = make([][]int, n)
	/*for j := 0; j < n; j++ {
		arr2[j] = make([]int, n)
	}*/
	for i := 1; i < n; i++ {
		str, err := reader.ReadString('\n')
		arr := strings.Fields(str)
		//var n = len(arr)
		//fmt.Print("le nombre de sommet est", len(arr))
		/*if len(arr) != n {
			fmt.Print("nombre de sommet incorrect")
			os.Exit(1)
		}
		var arr2 = make([][]int, n)
		for j := 0; j < n; j++ {
			arr2[i] = make([]int, n)
		}*/
		for j := 0; j < n; j++ {
			//arr2[i] = make([]int, n)
			arr2[i][j], _ = strconv.Atoi(arr[j])
			//print("i:", i, "j:", j, "arr:", arr2[i][j], "\n")
			fmt.Println(arr2)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	//fmt.Print("arr:", arr2)
	file.Close()
	return arr2, n, src

	//return arr2, n, src

}

func Child(fileName string) {
	fmt.Print("begin child")
	tab, lenth, source := openfile(fileName)
	cost, road := dijistra(lenth, tab, source)
	fmt.Print("road :", road)
	fmt.Print("cost:", cost)
	file, err := os.OpenFile("res"+fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	defer file.Close()
	for i := 0; i < lenth; i++ {
		str = "the costfrom" + strconv.Itoa(source) + " to " + strconv.Itoa(i) + " est " + strconv.Itoa(cost[i]) + "\n"
		file.WriteString(str)
	}

	for i := 0; i < lenth; i++ {
		str = "\nthe road from " + strconv.Itoa(source) + " to " + strconv.Itoa(i) + " est \n" + strconv.Itoa(i) + "<-"
		file.WriteString(str)
		for q := i; road[q] != NO_ROUTE; {
			if road[q] == source {
				str2 := strconv.Itoa(road[q])
				q = road[q]
				file.WriteString(str2)
			} else {
				str2 := strconv.Itoa(road[q]) + "<-"
				q = road[q]
				file.WriteString(str2)
			}
		}
	}

}

func SendFile(resFileName string, fileSize int64, conn net.Conn) {
	f, err := os.Open(resFileName)
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
	fileName := string(buf[:n])
	//获取客户端ip+port
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + ": 客户端传输的文件名为--" + fileName)
	//告诉客户端已经接收到文件名
	conn.Write([]byte("ok"))
	//创建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//循环接收客户端传递的文件内容
	buf = make([]byte, 2048)
	//for {
	fmt.Print(buf[:n])
	fmt.Print("1")
	n, err4 := conn.Read(buf)
	if err4 != nil {
		if err4 == io.EOF {
			fmt.Println("文件接受完毕")
			//i = NO_ROUTE
			fmt.Println(addr + ": 协程结束")
			//runtime.Goexit()
			//break
			fmt.Println(addr + ": 协程结束")
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
	//}
	//defer conn.Close()
	fmt.Println("fclose")
	defer f.Close()
	fmt.Println("close")
	Child(fileName)
	resFileName := "res" + fileName
	fileInfo, err := os.Stat(resFileName)
	fileSize := fileInfo.Size()
	conn2, err := net.Dial("tcp", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn2.Close()
	conn2.Write([]byte(resFileName))
	buf2 := make([]byte, 2048)
	//读取服务端内容
	n2, err := conn2.Read(buf2)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf2[:n2])
	if revData == "ok" {
		//发送文件数据
		SendFile(resFileName, fileSize, conn2)
	}
	//SendFile(resFileName, fileSize, conn2)
}

func main() {
	//创建tcp监听
	listen, err := net.Listen("tcp", ":8000")
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
		go Handler(conn)

	}

}
