/*
In LevenshteinDistance.go is function dist() whitch callculate Levenshtein Distance
of two words. Levenshtein Distance is defined like minimal number of operation neded to change one word to ather.
Operation can be change one leter with other delete leter and add leter.
*/
package main

//import "fmt"

func MinOf(vars ...int) int {
    min := vars[0]

    for _, i := range vars {
        if min > i {
            min = i
        }
    }

    return min
}

func dist(sa string,ta string) int{

	s := []rune(sa)
	t := []rune(ta)

	if(len(s) == 1 || len(s) == 0){
		return len(t)
	}
	if(len(t) == 1 || len(t) == 0){
		return len(s)
	}

	var a[30][30] int
	for i := 1; i < len(s)+1; i++ {
		a[i][0] = i
	}
	for i := 1; i < len(t)+1; i++ {
		a[0][i] = i
	}
	var cost int
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if( s[i-1] == t[j-1] ){
				cost = 0
			} else{
				cost = 1
			}
			a[i][j] = MinOf(a[i-1][j]+1,a[i][j-1]+1,a[i-1][j-1]+cost)
		}
	}
	/*for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(t)+1; j++ {
			fmt.Printf("%d ",a[i][j])
		}
		fmt.Println()
	}*/
	//fmt.Printf("\n %d %d",len(s),len(t))
	return a[len(s)-1][len(t)-1]
}