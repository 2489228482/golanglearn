//https://studygolang.com/articles/10552
package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func MakeFileByIoutil(){
	path := "./file/test1.txt"
	fd := []byte("hello world 和罗咯 我忍了到")
	//filemode表示文件存储的权限读写执行
	err := ioutil.WriteFile(path,fd,0777)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("sava finish\n")
}


func ReadFileByIoutil(){
	path := "./file/test1.txt"
	f,err :=ioutil.ReadFile(path)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(f)
	s := string(f)
	fmt.Println(s+"\n")
}

func ReadFileByOpenFile(){
	path := "./file/test1.txt"
	// os.O_RDONLY: 以只读的方式打开
	// os.O_WRONLY: 以只写的方式打开
	// os.O_RDWR : 以读写的方式打开
	// os.O_NONBLOCK: 打开时不阻塞
	// os.O_APPEND: 以追加的方式打开
	// os.O_CREAT: 创建并打开一个新文件
	// os.O_TRUNC: 打开一个文件并截断它的长度为零（必须有写权限）
	// os.O_EXCL: 如果指定的文件存在，返回错误
	// os.O_SHLOCK: 自动获取共享锁
	// os.O_EXLOCK: 自动获取独立锁
	// os.O_DIRECT: 消除或减少缓存效果
	// os.O_FSYNC : 同步写入
	// os.O_NOFOLLOW: 不追踪软链接
	f,err := os.OpenFile(path,os.O_RDWR|os.O_CREATE,0755)
	if err != nil{
		log.Fatal(err)

	}
	defer f.Close()

}

func ReadFileByOpen(){
	path := "./file/test1.txt"
	fd,err := os.Open(path)
	if err != nil{
		log.Fatal(err)
	}
	b := make([]byte,20)
	s,err := fd.Read(b)
	fmt.Println(string(s))
	os.Remove(path)

}


func main(){
	ReadFileByOpen()
	//ReadFileByOpenFile()
	//MakeFileByIoutil()
	//ReadFileByIoutil()
}