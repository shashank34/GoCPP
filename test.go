package main

import("fmt")

func hcf(number1 int, number2 int) int {


	fmt.Println(number1,number2)
	if  number2 == 0 { 
		return number1 
	}

	return hcf(number2, number1 % number2)   

}

func main() {

	hcf(8,10)
	// len:=4

	// count:=0;

	// array := [][]int{};


	// for index:=0; index < len; index += 1 {

	// 	back_iteration := index - 1;

	// 	next_iteration := index + 1;

	// 	if hcf(index, index) == 1 {
	// 		count = count + 1;
	// 		new_array3 := []int{index, index}
	// 		array = append(array, new_array3);	
	// 	}

	// 	if next_iteration < len && hcf(index, next_iteration) == 1 {
	// 		count = count + 1;
	// 		new_array := []int{index, next_iteration}
	// 		array = append(array, new_array);
	// 	}

	// 	for back_iteration >= 0 {

	// 		if hcf(back_iteration, index) == 1 {
	// 			count = count + 1
	// 			new_array2 := []int{index, back_iteration}
	// 		    array = append(array, new_array2);
	// 		}

	// 		back_iteration = back_iteration - 1;

	// 	}
	// }

	// fmt.Println(count);
	// fmt.Printf("%v", array);
}