package main
import "fmt"

func cetakBaris(n int){
	if n < 2{
		return 
	}
	if n%2 == 0{
		fmt.Print(n)
		if n > 2{
			fmt.Print(", ")
		}
	}
	cetakBaris(n-1)
}

func main(){
	var masukan int
	fmt.Scan(&masukan)
	cetakBaris(masukan)
	fmt.Println()
}