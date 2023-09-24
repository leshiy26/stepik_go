package main

func main1(){


var a [3]int = [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [...]int{1, 2, 3} // ... вместо указания длинны массива
d := [3]int{1: 12}a := [5]int{1, 2, 3, 4, 5}
fmt.Println(a) // [1 2 3 4 5]

//итерация по всем элементам (range получает индекс и КОПИЮ элемента массива)
a := [5]int{1, 2, 3, 4, 5}
for idx, elem := range a {
    fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
}

}
func main(){
	workArray := [10]int
	i:=0
	for idx := range workArray{
		fmt.Scan(&i)
		workArray[idx] = i
	}	
	fmt.Print(workArray)
}