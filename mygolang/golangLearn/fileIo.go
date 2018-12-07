//https://studygolang.com/articles/10552
package main
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"io"
)

func CreateFileByIoutil(){
	path := "./file/test1.txt"
	fd := []byte("hello world 和罗咯 我忍了到")
	//filemode表示文件存储的权限读写执行
	err := ioutil.WriteFile(path,fd,0777)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("sava finish\n")
}

func CreateFileByOS(){
	//若文件已经存在，则往文件中追加内容
	path := "./file/test2"
	f,err := os.Create(path)
	if err !=nil{
		panic(err)
	}
	defer f.Close()
	//使用字节写入
	d := []byte{115,111,109,101,10}
	n2,err := f.Write(d)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("worte %d bytes\n",n2)


	//使用支付串写入
	n3,err := f.WriteString("hello world 和罗咯我忍了到")
	fmt.Printf("wrote %d bytes\n",n3)

	//以同步方式打开
	f.Sync()
	
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes \n",n4)

	//将流中的缓存字符串发出去
	w.Flush()


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
	//b长度不够会出现 hello world 和罗�� 
	defer fd.Close()
	b := make([]byte,200)
	s,err := fd.Read(b)
	fmt.Printf("%d bytes: %s \n",s,string(b))

	//Seek 游标
	os2,err := fd.Seek(6,0)
	if err != nil{
		log.Fatal(err)
	}
	b2 := make([]byte,200)
	s2,err := fd.Read(b2)
	fmt.Printf("%d bytes @%d: %s \n",s2,os2,b2)

    o3, err := fd.Seek(6, 0)
    if err != nil{
		log.Fatal(err)
	}
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(fd, b3, 2)
    if err != nil{
		log.Fatal(err)
	}
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
}


func main(){
	CreateFileByOS()
	// CreateFileByIoutil()
	// ReadFileByIoutil()
	//ReadFileByOpenFile()
	//ReadFileByOpen()
}