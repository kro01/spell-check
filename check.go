package main

import "strings"

func checkSpell(word string) ([]string){
	var (
		tmp *Node
		d int
		result []string
	)
	wait :=make([]*Node, 0)
	wait = append(wait,&root)
	for ;len(wait) != 0;{
        tmp = wait[0]
		wait = wait[1:]

		d = dist((*tmp).word,strings.ToLower(word))

		if d < tol{
			result = append(result,(*tmp).word)
			
			for i := 0;i<=d+tol;i++{
				if (*tmp).child[i] != nil{
					wait = append(wait,(*tmp).child[i])
				}
			} 
		} else{
			for i := d-tol; i <= d+tol; i++{
				if (*tmp).child[i] != nil{
					wait = append(wait,(*tmp).child[i])
				}
			} 
		}
	}
	return result
}