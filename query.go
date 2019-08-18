package main

import "fmt"

type A struct{
	word *Node
	next *A
	previous *A
}

var first,last *A

type query struct{
	first *A
	last *A
}
func (q *query) push(s *Node){
	//fmt.Printf("\n\n")
	if(q.first == nil){
		tmp := A{s,nil,nil}
		q.last = &tmp
		q.first = q.last
		//fmt.Println("push")
		//showQ()
	} else{
		//fmt.Println("push")
		tmp := A{s,first,nil}
		a := &tmp

		(*(q.first)).previous = a
		//fmt.Println("*first")
		//fmt.Println(*first)
		q.first = &tmp
		//fmt.Println("push insait  ")
		//showQ()
		//fmt.Printf("\n\n")
		
	}
}

func (q *query) pop() (bool , *Node){
	if(q.last == nil){
		return false, nil
	}else{
		//fmt.Println("pop insait ")
		//fmt.Println(*(q.last))
		//showQ()


		tmp := *(q.last)
		q.last = tmp.previous
		if(q.last != nil){
			(*(q.last)).next = q.last
			//showQ()
		}
		//tmp2 := A{"",last,nil}
		//*last.next = tmp2
		//fmt.Println("last")
		//last.next = last
		//fmt.Println(*last)
		/*
		tmp :=*last
		a := *(tmp.previous)
		b := *(a.previous)
		tmp2 := A{a.word, &b ,nil}
		b.next = &tmp2
		last = &tmp2
		*/
		return true , tmp.word
	}
}
func (q query)showQ(){
	a := q.first
	for ;;{
		if(a==nil){
			break
		} else{
			fmt.Println(*a)
			a = a.next
		}
	}
}
//var q query
/*func test(){
	//fmt.Printf("2\n")
	var q query
	q.push("1")
	//q.showQ()
	q.push("2")
	//q.showQ()
	q.push("3")
	//q.showQ()
	//q.push("4")
	a, b := q.pop()
	fmt.Printf("pop %t %s\n",a,b)
	//showQ()
	a, b = q.pop()
	fmt.Printf("pop %t %s\n",a,b)
	a, b = q.pop()
	fmt.Printf("pop %t %s\n",a,b)
	a, b = q.pop()
	fmt.Printf("pop %t %s\n",a,b)
}*/