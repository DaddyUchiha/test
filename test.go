package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)


func main(){

	type cool struct{
		name string
		age int
		address string
	}

	var user cool
	user.name = "aman"
	user.age = 23
	user.address = "bhopal"

	p := cool{"seenu",22,"bairagarh"}

	json, _ := json.Marshal(&p)

	fmt.Println(string(json))

	asd := user

	crea,_ := os.Create("new.txt")

	wr := bufio.NewWriter(crea)

	_,err := wr.WriteString(asd.name)
	if err != nil{
		fmt.Println(err)
	}

	err = wr.Flush()
	if err != nil{
		fmt.Println(err)
	}
	present,err := os.Stat("new.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(present.Name())


	// aa := os.Remove("new.txt")
	// if aa != nil{
	// 	fmt.Println("error")
	// }else{
	// 	fmt.Println("Done deleting")
	// }

	b := "/snap/bin/brave"
	brave := exec.Command(b)

	err = brave.Start()

	get,_ := os.Getuid(), os.Getgid()

	fmt.Println(get)
	list()

	open("new.txt")

	fmt.Println("Hello")
}  

func list(li ...string){
	comm := exec.Command("ls","-la")

	commt, _ := comm.Output()

	fmt.Println(string(commt))


}

func open(file string){
	//var bug bufio.Scanner
	file2,_ := os.OpenFile("/media/unknown/New Volume/xssscript.js",os.O_APPEND,0644)

	scan := bufio.NewScanner(file2)

	for scan.Scan(){
		te := scan.Text()
		fmt.Println(te)
	}

	wri := bufio.NewWriter(file2)

	write := "Hello this is written by Go bufio.NewWriter"

	wrt := wri.Write([]byte(write))


} 
