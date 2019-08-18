package main


import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

var tol int // tolerance limit 

func main(){
	
	tol = 2
	readWords()
    replace()
    countDuble()
    //show()
    delDuble()
    //show()
    bild()
    //showBild()
   
    
    
    reader := bufio.NewReader(os.Stdin)
  	fmt.Println("Simple Spell Check")
 	fmt.Println("---------------------")

    for{	
    	fmt.Println("Enter text: ")
    	word, _ := reader.ReadString('\n')
   	 	// convert CRLF to LF
    	word = strings.Replace(word, "\n", "", -1)
    	fmt.Println("The result is:")
    	fmt.Println(checkSpell(word))
    	fmt.Println("\n")
    }
}