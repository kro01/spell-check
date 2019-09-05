package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// tolerance limit 
var tol int 

func main(){
	
	tol = 2
	//readBG()

	readEN()

    bild()

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
    	fmt.Println()
    	fmt.Println()
    }
    
}