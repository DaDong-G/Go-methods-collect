package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main()  {
	Trace.Println("这是Trac")
	Info.Println("这是Info")
	Warning.Println("这是Warning")
	//Error.Println("这是Error")
}

var (
	Trace 			*log.Logger // 记录所有的日志
	Info  			*log.Logger // 记录重要的信息
	Warning 		*log.Logger // 需要注意的信息
	Error    		*log.Logger // 记录非常严重信息
)

func init() {
	file,err := os.OpenFile("./errors.txt",os.O_CREATE,0666)
	if err!= nil{
		log.Fatalln("Failed to open error log")
	}
	Trace = log.New(ioutil.Discard,"Trace:",log.Ldate | log.Ltime | log.Lshortfile)  //ioutil.Discard 不输出，不写入，可以用于占位
	Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile) //os.Stdout  不记录日志仅仅作为标准输出
	Warning = log.New(file,"Warning:",log.Ldate | log.Ltime | log.Lshortfile) // file  会将日志写入file,并输出
	Error = log.New(io.MultiWriter(file,os.Stdin),"Error:",log.Ldate | log.Ltime | log.Lshortfile) //写入，不输出，可以更改os.std。
}

