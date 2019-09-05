/*
bildTree.go bild BK-tree from array words.
*/
package main

import "fmt"


type Node struct{
    word string
    child [15]*Node    
}

//the root of the tree
var root Node

//bild BK-tree from array words.
func bild() {

    root.word = words[0]
    d := 0 
    var( 
        tmp ,tmp2 *Node
    )

    for i := range words{
        if i == 0 {continue}
        if words[i] == "" {continue}
        
        tmp = new(Node)
        (*tmp).word = words[i]

        d = dist(root.word, (*tmp).word)
        if d == 0{
            //fmt.Printf("problem %s  %d",words[i],i)
        }
        tmp2 = &root

        for ;;{
        
            d = dist(((*tmp).word),((*tmp2).word))

            if ((*tmp2).child)[d] == nil {
                ((*tmp2).child)[d] = tmp
                break
            } else{
                tmp2 = ((*tmp2).child)[d]
            }  
        } 
    }
}


// show builded BK-tree
func showBild(){
    fmt.Println(root.word)
    wait := make([]*Node,0)
    wait = append(wait,&root)
    var tmp *Node
    p:=0
    for ;len(wait) != 0;{
        p++
        tmp = wait[0]
        wait = wait[1:]
        fmt.Printf("\n %d %s",p,(*tmp).word)
        if (*tmp).word == ""{
            fmt.Println(p)
        }
        for i := range (*tmp).child{
            if (*tmp).child[i] != nil{
                wait = append(wait,(*tmp).child[i])
            }
        }
    }
}