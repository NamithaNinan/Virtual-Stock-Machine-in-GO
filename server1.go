package main

import (
	"log"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/http"
	"strings"
	"encoding/json"
	"strconv"
	
)


type Abc int
type First int
type Third int
var generic map[string]interface{}
var askvalue[2] float64
var firstperc,secperc,totamount int
var number[2] int
var amount[2] float64
var symbol[2] string
var k int =0
var ide int =0
var tradeno int
var trade_id int=0
var trade string
var oldsymbol[5][5] string
var oldaskvalue[5][5] float64
var newaskvalue[5][3] float64



func stockcalculate(){

amount[0]=float64(totamount)*float64(firstperc)/100
amount[1]=float64(totamount)*float64(secperc)/100

number[0]=int(amount[0]/askvalue[0])
number[1]=int(amount[1]/askvalue[1])

var numberstr[2] string
var askvaluestr[2] string
var askvaluei[2] int
numberstr[0]=strconv.Itoa(number[0])
numberstr[1]=strconv.Itoa(number[1])
askvaluei[0]=int(askvalue[0])
askvaluei[1]=int(askvalue[1])
askvaluestr[0]=strconv.Itoa(askvaluei[0])
askvaluestr[1]=strconv.Itoa(askvaluei[1])

var st string
var unvested[2] float64
unvested[0]=amount[0]-(float64(number[0])*askvalue[0])
unvested[1]=amount[1]-(float64(number[1])*askvalue[1])
fmt.Println("trade id : ",trade_id)

	 
st="stocks is "+symbol[0]+":"+numberstr[0]+":$"+askvaluestr[0]+" , "+symbol[1]+":"+numberstr[1]+"$"+askvaluestr[1]
fmt.Println(st)
fmt.Println("unvested amounts are ",unvested[0],unvested[1])

tradeids:=strconv.Itoa(trade_id)
oldsymbol[trade_id][0]=tradeids
oldaskvalue[trade_id][0]=float64(trade_id)
oldsymbol[trade_id][1]=symbol[0]
oldsymbol[trade_id][2]=symbol[1]
oldsymbol[trade_id][3]=numberstr[0]
oldsymbol[trade_id][4]=numberstr[1]
oldaskvalue[trade_id][1]=askvalue[0]
oldaskvalue[trade_id][2]=askvalue[1]
oldaskvalue[trade_id][3]=unvested[0]
oldaskvalue[trade_id][4]=unvested[1]

trade_id++
ide++
}

func parseMap(aMap map[string]interface{}){

    for key, val := range aMap {
        switch concreteVal := val.(type) {
        case map[string]interface{}:
           
			
			parseMap(val.(map[string]interface{}))
        case []interface{}:
            
			
			parseArray(val.([]interface{}))
		default:
            
			if(key=="Ask"){
				if(k>1){k=0}
				//fmt.Println("value is procured to be",concreteVal)
				//i:=val.(concreteVal)
				//iAreaId, ok := concreteVal.(float64) 
				iAreaId:=concreteVal.(string)
				askarr,_:=strconv.ParseFloat(iAreaId,64)
				//println(err)
				askvalue[k]=askarr
				
				k++
				//fmt.Println("integer value of ask is ",askvalue[0],askvalue[1])
				}
			//fmt.Println()
						
		}
    }
}


func parseArray(anArray []interface{}) {
    for i,val := range anArray {
        switch concreteVal := val.(type) {
        case map[string]interface{}:
           
            parseMap(val.(map[string]interface{}))
			
        case []interface{}:
          
            parseArray(val.([]interface{}))
			
        default:
            fmt.Println("Index", i, ":", concreteVal)
			
		
    }
}
}

func parseMap2(aMap map[string]interface{},num int){

    for key, val := range aMap {
        switch concreteVal := val.(type) {
        case map[string]interface{}:
           
			
			parseMap2(val.(map[string]interface{}),num)
        case []interface{}:
            
			
			parseArray2(val.([]interface{}),num)
		default:
            
			if(key=="Ask"){
				if(k>1){k=0}
				//fmt.Println("value is procured to be",concreteVal)
				//i:=val.(concreteVal)
				//iAreaId, ok := concreteVal.(float64) 
				iAreaId:=concreteVal.(string)
				askarr,_:=strconv.ParseFloat(iAreaId,64)
				//println(err)
				askvalue[k]=askarr
				newaskvalue[num][0]=float64(num)
				newaskvalue[num][1]=askvalue[0]
				newaskvalue[num][2]=askvalue[1]
				k++
				//fmt.Println("integer value of ask is ",askvalue[0],askvalue[1])
				}
			//fmt.Println()
						
		}
    }
}


func parseArray2(anArray []interface{},num int) {
    for i,val := range anArray {
        switch concreteVal := val.(type) {
        case map[string]interface{}:
           
            parseMap2(val.(map[string]interface{}),num)
			
        case []interface{}:
          
            parseArray2(val.([]interface{}),num)
			
        default:
            fmt.Println("Index", i, ":", concreteVal)
			
		
    }
}
}

func Random( result []string)int{
i:=1
j:=0
var percentage[5] string
for i<=3{

		//k:=0
		perc:=strings.Split(result[i],"%")
		percentage[j]=perc[0]
		j=j+1
		i=i+2
	}
	
		firstperc,_=strconv.Atoi(percentage[0])
		secperc,_=strconv.Atoi(percentage[1])
		
		
return 0
}


//the budget is converted to an integer value
func (r *First)Budgetcal (budget string,reply *int) error{

i,err := strconv.Atoi(budget)
if err != nil {
		log.Fatal("there is an error")
		}
		totamount=i

stockcalculate()
*reply=20
return nil
}

//the url is opened and decoded and sent to be parsed
func (t *Abc)Urlopen (result3 []string,reply *int) error{

var link string
symbol[0]=result3[0]
symbol[1]=result3[2]
link ="https://query.yahooapis.com/v1/public/yql?q=select%20symbol%2CAsk%20from%20yahoo.finance.quotes%20where%20symbol%20in%20(%22"+result3[0]+"%22%2C%22"+result3[2]+"%22)&format=json&env=http%3A%2F%2Fdatatables.org%2Falltables.env&callback="

u, err := http.Get(link)
	if err != nil {
		log.Fatal("there is an error")
		}
	   err = json.NewDecoder(u.Body).Decode(&generic)
	   parseMap(generic)
   if err != nil {
   log.Fatal(err)
   }
 
   Random(result3)
   *reply=10
	return nil
}


func(g *Third)Portfolio(tradenos string,reply *int) error{
tradeno,_=strconv.Atoi(tradenos)
l:=0
for l<len(oldaskvalue){
	if(tradeno==int(oldaskvalue[l][0])){
	//fmt.Println("hi")
	parseMap2(generic,tradeno)
	//a:=strconv.Itoa(int(oldaskvalue[tradeno][3]))
	//b:=strconv.Itoa(int(oldaskvalue[tradeno][4]))
	c:=strconv.Itoa(int(newaskvalue[tradeno][1]))
	d:=strconv.Itoa(int(newaskvalue[tradeno][2]))
	fmt.Println()
	fmt.Println("__________________PORTFOLIO_____________________")
	
	if(newaskvalue[tradeno][1]>=oldaskvalue[tradeno][1]&&newaskvalue[tradeno][2]>=oldaskvalue[tradeno][2]){
		str1:=oldsymbol[tradeno][1]+":"+oldsymbol[tradeno][3]+":$+"+c+" , "+oldsymbol[tradeno][2]+":"+oldsymbol[tradeno][4]+":$+"+d
		fmt.Println(str1)
	}else if(newaskvalue[tradeno][1]>=oldaskvalue[tradeno][1]&&newaskvalue[tradeno][2]<oldaskvalue[tradeno][2]){
		str1:=oldsymbol[tradeno][1]+":"+oldsymbol[tradeno][3]+":$+"+c+" , "+oldsymbol[tradeno][2]+":"+oldsymbol[tradeno][4]+":$-"+d
		fmt.Println(str1)
	}else if(newaskvalue[tradeno][1]<oldaskvalue[tradeno][1]&&newaskvalue[tradeno][2]>=oldaskvalue[tradeno][2]){
		str1:=oldsymbol[tradeno][1]+":"+oldsymbol[tradeno][3]+":$-"+c+" , "+oldsymbol[tradeno][2]+":"+oldsymbol[tradeno][4]+":$+"+d
		fmt.Println(str1)
	}else {
		str1:=oldsymbol[tradeno][1]+":"+oldsymbol[tradeno][3]+":$-"+c+" , "+oldsymbol[tradeno][2]+":"+oldsymbol[tradeno][4]+":$-"+d
		fmt.Println(str1)
	}
	fmt.Println("unvested amounts are ",oldaskvalue[tradeno][3],oldaskvalue[tradeno][4])
	fmt.Println("new value",newaskvalue[tradeno][1],newaskvalue[tradeno][2])
	
	return nil 
		}
	}
*reply=10
return nil

}
func main() {
	ab := new(Abc)
	cd := new(First)
	gh := new(Third)
	server := rpc.NewServer()
	server.Register(ab)
	server.Register(cd)
	server.Register(gh)
	
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}