package hdfs

import (
	"bytes"
	"fmt"
	"github.com/vladimirvivien/gowfs"
)

func HDFSConnectTest() {


	config := gowfs.Configuration{Addr: "127.0.0.1:9870", User: "root"}
	client, err := gowfs.NewFileSystem(config)
	if err != nil {
		panic(fmt.Sprintln("gowfs.NewFileSystem(config) err", err))
	}

	path := gowfs.Path{Name: "/test.txt"}
	/*
		Create函数接收如下参数。
		data:io.Reader,一个实现了io.Reader接口的struct
		Path:很简单,就是我们这里的path
		overwrite:是否覆盖,如果为false表示不覆盖，那么要求文件不能存在，否则报错
		blocksize:块大小
		replication:副本
		permission:权限
		buffersize:缓存大小
		contenttype:内容类型
		返回一个bool和error
	*/
	if _, err :=client.Create(
		bytes.NewBufferString("this is first test"), //如果不指定内容，就直接bytes.NewBufferString()即可
		path, //路径
		true,//不覆盖
		0,
		0,
		0666,
		0,
	); err != nil {
		fmt.Println("hdfs create test file error:", err)
	} else {
		fmt.Println("hdfs create test file success")  //创建文件成功, flag = true
	}

	/*
		Create函数接收如下参数。
		data:io.Reader,一个实现了io.Reader接口的struct
		Path:很简单,就是我们这里的path
		overwrite:是否覆盖,如果为false表示不覆盖，那么要求文件不能存在，否则报错
		blocksize:块大小
		replication:副本
		permission:权限
		buffersize:缓存大小
		contenttype:内容类型

		返回一个bool和error
	*/
	//if err != nil {
	//	fmt.Println("hdfs error:", err)
	//} else {
	//	fmt.Println("hdfs success")  //创建文件成功, flag = true
	//}


	//tcp        0      0 localhost:9000          0.0.0.0:*               LISTEN
	//tcp        0      0 localhost:27017         0.0.0.0:*               LISTEN
	//tcp        0      0 0.0.0.0:9866            0.0.0.0:*               LISTEN
	//tcp        0      0 0.0.0.0:9867            0.0.0.0:*               LISTEN
	//tcp        0      0 localhost:20043         0.0.0.0:*               LISTEN
	//tcp        0      0 0.0.0.0:9868            0.0.0.0:*               LISTEN
	//tcp        0      0 localhost:20044         0.0.0.0:*               LISTEN
	//tcp        0      0 localhost:37005         0.0.0.0:*               LISTEN
	//tcp        0      0 localhost:20013         0.0.0.0:*               LISTEN

	//file, _ := client.Open("/mobydick.txt")

	//buf := make([]byte, 59)
	//file.ReadAt(buf, 48847)
	//
	//fmt.Println(string(buf))

}
