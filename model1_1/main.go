package main

import "fmt"

func main(){
	strs :=[]string{"I","am","stupid","and","weak"}
	for ix,_ :=range strs{
		if ix==2{
			strs[ix]="smart"
		}else if ix==4{
			strs[ix]="strong"
		}
	}
	fmt.Printf("%v",strs)
}
