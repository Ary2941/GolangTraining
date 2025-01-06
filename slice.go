package main
import ("fmt")


func main(){
	slice1 := make([]string,0,10)
	slice1 = append(slice1,"paquetá")
	slice1 = append(slice1,"ronaldo")
	slice1 = append(slice1,"neymar")
	slice1 = append(slice1,"kaká")

	fmt.Print(slice1)
	
}
