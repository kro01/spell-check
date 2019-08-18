package main

import "fmt"

type Node struct{
    word string
    child [15]*Node    
}

var root Node


func bild() {

    root.word = words[0]
    d := 0 
    var( 
        tmp *Node
    )

    for i := range words{
        //fmt.Printf("\n %s",words[i])
        if i == 0 {continue}
        if words[i] == "" {continue}
        
        tmp = new(Node)
        (*tmp).word = words[i]

        //bfs
        wait := make([]*Node, 0)
        wait = append(wait, &root)
        d = dist(root.word, (*tmp).word)
        if d == 0{
            //fmt.Printf("problem %s  %d",words[i],i)
        }
        var tmp2 *Node
        p := 0
        for ;;{
            if len(wait) == 0 {
                fmt.Println("problem")
            }
            tmp2 = wait[0]
                // Discard top element
            wait = wait[1:]
                // Is empty ?
                //fmt.Println(x)
            //fmt.Printf("\n  %d   %s",p,(*tmp2).word)
            p++
            

            d = dist(((*tmp).word),((*tmp2).word))

            if ((*tmp2).child)[d] == nil {
                ((*tmp2).child)[d] = tmp
                //fmt.Printf("\n %d %s %s",d,(*((*tmp2).child[d])).word,(*tmp2).word)
                break
            } else{
                for j := range((*tmp2).child){
                    if( (*tmp2).child[j] != nil){
                        wait = append(wait,((*tmp2).child[j]))
                    }
                }
            }  
        } 
    }
}

func showBild(){
    //fmt.Println(root.word)
    wait := make([]*Node,0)
    wait = append(wait,&root)
    var tmp *Node
    p:=0
    for ;len(wait) != 0;{
        p++
        tmp = wait[0]
        wait = wait[1:]
        //fmt.Printf("\n %d %s",p,(*tmp).word)
        if (*tmp).word == ""{
            fmt.Println(p)
        }
        for i := range (*tmp).child{
            if (*tmp).child[i] != nil{
                wait = append(wait,(*tmp).child[i])
            }
        }
    }
    /*
    for i := range root.child{
        fmt.Printf("\n %d",i)
        //fmt.Println(root.child[i])
        if root.child[i] != nil{
            fmt.Printf(" %d  %s",i,(*(root.child[i])).word)
        }
    }*/
}