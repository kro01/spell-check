/*
This file read 3000 words from EnglisheWords.txt and put in array words
It do it by callyng function readEN. It use some functions from readBG.
*/
package main

import (
    "io/ioutil"
    "regexp"
)

func readEN(){
    readWordsEN()
    delNotWord()
    replace()
    countDuble()
    countDuble()
}

func checkErrEN(e error) {
    if e != nil {
        panic(e)
    }
}

//read words from  EnglishWords.txt
func readWordsEN(){
    dat, err := ioutil.ReadFile("EnglisheWords.txt")
    checkErrEN(err)

    spliter := regexp.MustCompile(`\n`)
    //words = make([]string,3020,3020)
    words = spliter.Split(string(dat), 3020)

}

func delNotWord(){
    re := regexp.MustCompile(`[^a-z]`)
    for i,word := range words{
        words[i] = re.ReplaceAllString(word,"")
    }
}