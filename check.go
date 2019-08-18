package main

//import "fmt"
func checkSpell(word string) ([]string){
	var (
		//wait = query{nil,nil}
		//a bool
		tmp *Node
		d int
		result []string
	)

	//wait.push(root)
	wait :=make([]*Node, 0)
	//wait = append(wait,nil)
	wait = append(wait,&root)
	for ;len(wait) != 0;{
		//a,tmp = wait.pop()
		/*if len(wait) == 0 {
                fmt.Println("end")
                break
        }*/
        tmp = wait[0]
        //fmt.Println(word)
        //fmt.Println((*tmp).word)
		wait = wait[1:]

		d = dist((*tmp).word,word)

		if d < tol{
			result = append(result,(*tmp).word)
			//fmt.Println(tmp.word)
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