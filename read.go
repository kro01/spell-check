package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strings"
)
	
var words []string

func checkErr(e error) {
    if e != nil {
        panic(e)
    }
}

func readWords(){
    dat, err := ioutil.ReadFile("words.txt")
    checkErr(err)

    spliter := regexp.MustCompile(`\n`)
    dat2 := spliter.Split(string(dat), 996)

    line := regexp.MustCompile(`\t`)
    words = make([]string,996,996)
    for i:= range dat2{
        tmp := line.Split(dat2[i],5)
        //fmt.Println(tmp[1])
        words[i] = tmp[1]
    }
}

func replace(){
    re := regexp.MustCompile(` *$`)
    for i := range words{
        words[i] = re.ReplaceAllString(words[i],"")
        words[i] = strings.ToLower(words[i])
    }
}

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

func show(){
    fmt.Printf("\n\n")
    for i := range words{
        fmt.Printf("%s,",words[i])
    }
}





    /*b := []byte(" ")
    //fmt.Println(b[0])
    c := 0
    for i:= range words{
        for j:= range words[i]{
            if(words[i][j] == b[0]){
                fmt.Println(words[i])
                c++
                break
            }
        }
    }
    fmt.Println(c)
    max := 0
    w := ""
    j := 0
    for i:= range words{
        if(len(words[i]) > max){
            max = len(words[i])
            w = words[i]
            j = i
        }
    }
    a := []rune(words[j])
    //a = []rune("asdasd")
    for i:=0; i < len(a); i++{
        fmt.Printf(" %c ", a[i])
    } 
    fmt.Println(words[j][0])
    fmt.Println(max)
    fmt.Println("bigest word is")
    fmt.Println(w) */
