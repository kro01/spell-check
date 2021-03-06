/*
This file read 1000 words from words.txt and put in array words
It do it by callyng function readBG.
*/
package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"
    "runtime"
)
	
var words []string

func readBG(){
    //openFile()
    readWords(openFile())
    replace()
    countDuble()
    delDuble()
}

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

//open file words.txt from the folder of sourse code
func openFile() string{
    _, filePathe, _, _ := runtime.Caller(0)
    //fmt.Println(filePathe)
    re := regexp.MustCompile(`/[^/]+\.go\z`)
    filePathe = string(re.ReplaceAll([]byte(filePathe), []byte("/words.txt")))
    //fmt.Println(filePathe)
    re = regexp.MustCompile(`/`)
    filePathe = string(re.ReplaceAll([]byte(filePathe), []byte("\\")))
    code, err := ioutil.ReadFile(filePathe)
    checkErr(err)
    //fmt.Println(string(code))
    return string(code)
}
//take off not used words from the file words.txt
func readWords(dat string){
    //dat, err := ioutil.ReadFile("words.txt")
    //checkErr(err)

    spliter := regexp.MustCompile(`\n`)
    dat2 := spliter.Split(string(dat), 996)

    line := regexp.MustCompile(`\t`)
    words = make([]string,996,996)
    for i:= range dat2{
        tmp := line.Split(dat2[i],5)
        //fmt.Println(tmp[1])
        words[i] = strings.ToLower(tmp[1])
    }
}

//delete emty space from the end of the word
func replace(){
    re := regexp.MustCompile(` *$`)
    for i := range words{
        words[i] = re.ReplaceAllString(words[i],"")
        //words[i] = strings.ToLower(words[i])
    }
}

//countDuble() replaces repeated words with an empty string
func countDuble(){
    var dub map[string]int
    dub = make(map[string]int)
    for i := range words{
        _, ok := dub[words[i]]
        if ok {
            words[i] = ""
        } else{
            dub[words[i]] = 1
        }
    }
}

//delDuble() change takeoff emty strings 
func delDuble(){
    i := 0
    for j := range words{
        if words[j] == ""{
        }else{
            words[i] = words[j]
            i++
        }
    }
    for j:=i ;j<len(words);j++{
        words[j] = ""
    }
    //fmt.Printf("\n %d  %s \n",i,words[913])
    words = words[0:i]
    //fmt.Printf("\n %s  %s %s \n",words[912],words[913],words[910])
}

//show() display words on the screan
func show(){
    fmt.Printf("\n\n")
    for i := range words{
        fmt.Printf("%s,",words[i])
    }
}
