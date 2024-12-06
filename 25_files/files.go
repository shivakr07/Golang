package main

import (
	"fmt"
	"os"
)

func main() {

	// f, err := os.Open("hello.txt")
	// // keep in mind this if you want to run directly with this path then you have to come inside the directory first [cd 25_files] otherwise change the path
	// // in go errors are get return not throw so Open method is returns two things
	// if err != nil { // means err
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat() //return file info, err if path is wrong
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name is :", fileInfo.Name())
	// fmt.Println("is folder ? :", fileInfo.IsDir())
	// fmt.Println("file size :", fileInfo.Size()) //in byte
	// fmt.Println("file permission :", fileInfo.Mode())
	// fmt.Println("file modified at :", fileInfo.ModTime())
	// ---------------------------------------------------
	// READ FILE ...........
	// f, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// // this will be called when main will end
	// // keep things near to their relevent otherwise in between if some erro happens then that part you'll never reach tht you want to execute

	// // we read and store in buffer[arr of bytes]
	// buf := make([]byte, 10) //you can give size as well

	// d, err := f.Read(buf) //d-> len of buffer
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }
	// ------------------------------------------------------
	// FOR READING ...............
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))
	//but this method of reading the file is not good because
	//readfile method loaded the entire content of file in our memory in one go
	// so if file is small then okay but big [image, video ..] then it is problem
	// -----------------------------------------------------------------------
	// if file is big then we can do something streaming kind of thing as we do in nodejs
	//-----------------------------------------------------------
	// WE CAN READ FOLDER TO CHECK WHAT IT CONTAINS FILES, FOLDERS

	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1) //it will return slice of records[files and folders] of size mentioned
	// //if you keep n <= 0 it will return all

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// -------------------------------------------------------
	// HOW TO CREATE & WRITE IN THE FILE .....................

	// f, err := os.Create("newfile2.text")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("hi!!, go")
	// // f.WriteString(" nice to see you")
	// //it is working in append mode [not re writing]

	// //we can add data in other way also : although strings are also bytes
	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	//--------------------------------------------
	// we will read one file and write in other file
	// but we will utilize / mimic streamind kind of thing

	//READ & WRITE [after creating the file] TO ANOTHER FILE [STREAMIN FASHION]

	// sourceFile, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("testwritefile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// //for streaming fashion we have bufio
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)
	// //this buffer have its own size : check

	// for {
	// 	b, err := reader.ReadByte()
	// 	//readbyte returns byte[byte that it is reading and error]
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		//break if end of file to break infinite loop
	// 		break
	// 	}

	// 	e := writer.WriteByte(b) //it returns err only
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }
	// // we flush the remaining data from the writer
	// writer.Flush()

	// fmt.Println("written to new file successfully")

	//above method is for streaming fashion
	// -----------------------------------------------------------
	//IF YOU WANT TO COPY CONTENT ONLY : checkout
	// --------------------------------------------------------
	// HOW TO DELETE A FILE ...

	// sourceFile, err := os.Open("random.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	err := os.Remove("random.txt") //it returns err
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")
}
