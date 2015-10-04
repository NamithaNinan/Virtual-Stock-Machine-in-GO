package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"net/rpc/jsonrpc"
	//"net/url"
	//"net/http"
	//"encoding/json"
	"os"
)

type Stock struct {
	name string
	cost int
}

var ssp string
//var budget float32
func main() {

	client, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	c:=jsonrpc.NewClient(client)
	var response int
	var result2[] string
	k:=0

	var result3[10] string
	n:=len(os.Args)
	if(n==2){
		
		c.Call("Third.Portfolio",os.Args[1],&response)
	}else{
	result:=strings.Split(os.Args[1],",")
 
	for i := range result {
		result2=strings.Split(result[i],":")
		result3[k]=result2[0]
	
	k=k+1
	
	result3[k]=result2[1]
	
	k=k+1
	
	
	}
	
  budget:=os.Args[2]
  
  err=c.Call("Abc.Urlopen",result3,&response)
  c.Call("First.Budgetcal",budget,&response)
	if err!=nil{
	fmt.Println(err)
	}
	}
}
 
 