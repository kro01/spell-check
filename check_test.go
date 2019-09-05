package main

import (
	"testing"
	//"fmt"
)

func TestCheckSpell(t *testing.T) {

	tol = 2
	readBG()
    bild()

	var (
		words = []string{"гара","мара","аз","вас",}
		a int
	)
	for _,word := range words {
		for _,res := range checkSpell(word){
			a = dist(word,res)
			if a > 2{
				t.Errorf("Returnet result is wrong. Get %s, return: %s witch are on distance %d", word, res, a)
			}
		}
	}
}