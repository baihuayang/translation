package main
 
import (
    "fmt"
    "bufio"
    "os"
)

var inputReader *bufio.Reader
var input string
var err error


func main(){
    inputReader = bufio.NewReader(os.Stdin)
    fmt.Println("Please enter some input: ")
    input, err = inputReader.ReadString('S')
	
    if err == nil {
		count1 := 0  
		count2 := 1  
		count3 := 0  
		for _,v := range input{
			if(v != '\n' && v != '\r'){
				count1 += 1
			}
			if(v == 32 || v==10){
				count2 += 1
			}
			if(v == 10){
				count3 += 1
			}
		}
		fmt.Printf("The char count %d\n", count1)
        fmt.Printf("word count: %d\n", count2)
        fmt.Printf("The line length: %d\n", count3)
    }
}