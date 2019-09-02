# spell-check
This is console spell-check for Bulgarian. It is work under windows.<br>
Start program in folder spell-check with go run spell-check.<br>
Write word on Bulgarien and it while return you closest match.<br><br>

Levenshtein distance between two words is the minimal number of operations required to convert one word to another. Allowed operations are to change a letter to another letter, to rub a letter and to add a letter.<br>
It make it with dynamic programming.<br>
Thise is in file  	LevenshteinDistance.go <br><br>

Program read the word from file words.txt. The code for tha is in file  	read.go<br><br>

The program can work by comparing all words with the one submitted by the user and return closest,but this is slow. It is fast with BK tree. Every node of the tree store 1 word from dictionary. We take first word for root node of the tree. Еvеry node have 1 or 0 childs on the same distance from him. For every ather word from dictionary we calculate it distance from the root. If the root not have child withe thise distance we make new word child of the root with thise distance. If root have childe with thise distance make he same withe new word and thise child. Thise is make in file  	buildTree.go<br><br>

For the purpos of searche we difene tolerance (in code is tol=2). User enter word. Program return all word withe distance equal or smallar then tol to the entered word.
It calculete distance between user word and root word let it be n. If distance are smaller or equal to tol it addet to return list. It make the same withe all child of root in interval n-tol n+tol.<bk>
Searche is implement in check.go<br><br>
