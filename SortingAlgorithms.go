package main

import ( 
	 "fmt" 
    "math/rand" 
) 

func swap(first_address_value *int, second_address_value *int) {
	temp := *first_address_value;
	*first_address_value = *second_address_value;
	*second_address_value = temp;
}

// 2,7,4,1,5,3 
func SelectionSort(unsorted_array []int) [] int {
	array_size:= len(unsorted_array)
	for i:=0; i < array_size; i+=1 {
		
		for j:=i+1; j < array_size; j+=1 {
		
			if unsorted_array[i] > unsorted_array[j] {

				swap(&unsorted_array[i], &unsorted_array[j])
		
			} 
		}
	}
	return unsorted_array;
}

// 2,7,4,1,5,3 
func BubbleSort(unsorted_array []int) []int {

	flag:=1;
	 
	counter:=0;
	for flag > 0 {
		
		swap_flag := 0;
		
		for index:= 0; index < len(unsorted_array)-counter; index+=1 {
			
			next_iteration := index + 1;
			
			if next_iteration < len(unsorted_array)-counter && unsorted_array[index] > unsorted_array[next_iteration] {
				
				swap(&unsorted_array[index], &unsorted_array[next_iteration])

				swap_flag = 1;
			}
	   }

	   if swap_flag == 0 {
		flag = 0
	   } else {
		counter += 1;
	   }
    }

	return unsorted_array;	

}

// 2,7,4,1,5,3 
func InsertionSort(unsorted_array []int) []int {
	
	sorted_array:= []int{};

	for index:=0; index < len(unsorted_array); index+=1 {

		sorted_array_size:= len(sorted_array);

		if (sorted_array_size == 0) {
			sorted_array = append(sorted_array, unsorted_array[index])
		} else {
			found_index:=sorted_array_size;
			for sorted_index:=0; sorted_index < sorted_array_size; sorted_index +=1 {
				if (unsorted_array[index] <= sorted_array[sorted_index]) {
					found_index = sorted_index;
					break;
				} 
			} 
            if (found_index == sorted_array_size) {
				sorted_array = append(sorted_array, unsorted_array[index])
			} else {
				sorted_array = append(sorted_array[:found_index+1], sorted_array[found_index:]...)
				sorted_array[found_index] = unsorted_array[index]
			}
		}

	}

	return sorted_array;

}

// 2,7,4,1,5,3 
func InsertionInPlaceSort(unsorted_array []int) []int {

	for index:=0; index < len(unsorted_array); index+=1 {
		
		j:= index - 1;

		elem:=unsorted_array[index] 
		
		for j >= 0 {
			if (unsorted_array[j] > elem) {
				temp:= unsorted_array[j];
				unsorted_array[j] = elem;
				unsorted_array[j+1] = temp;
			}

			j = j - 1
		}
	}
	return unsorted_array
}

func MergeArray(sorted_array_left []int, sorted_array_right []int) []int {

	left_size:= len(sorted_array_left);
	right_size:= len(sorted_array_right);
	left_index:=0;
	right_index:=0;

	sorted_array:= []int{};

	for left_index < left_size && right_index < right_size {
		
		if sorted_array_left[left_index] <= sorted_array_right[right_index] {
		
			sorted_array = append(sorted_array, sorted_array_left[left_index])
		
			left_index += 1;

		} else   {
			
			sorted_array = append(sorted_array, sorted_array_right[right_index])
			
			right_index += 1;
		}

	} 

	for left_index < left_size {

		sorted_array = append(sorted_array, sorted_array_left[left_index])
		left_index += 1
	}

	for right_index < right_size {
		sorted_array = append(sorted_array, sorted_array_right[right_index]);
		right_index += 1
	}

	return sorted_array

}

func MergeSort(unsorted_array []int) []int {
	
	array_size := len(unsorted_array)
	
	if array_size < 2 {
		return unsorted_array
	}

	left_array := unsorted_array[:array_size/2]
	right_array := unsorted_array[array_size/2:]

	return MergeArray(MergeSort(left_array), MergeSort(right_array));

}

func PartitionData(unsorted_array []int, start_index int, end_index int) int {

	pivot_elem  := unsorted_array[end_index];

	pivot_index := start_index

	fmt.Println(unsorted_array);
	fmt.Println(pivot_elem);

	for index:=start_index; index < end_index; index+=1 {
		
		if unsorted_array[index] < pivot_elem {

			swap(&unsorted_array[pivot_index],&unsorted_array[index]);
	
			pivot_index += 1;
		}
	}

	swap(&unsorted_array[pivot_index],&unsorted_array[end_index]);
	

	return pivot_index
}

func RandomPartitionData(unsorted_array []int, start_index int, end_index int) int {
	
	RandomIntegerwithinRange := rand.Intn(end_index -start_index) + start_index;
	
	swap(&unsorted_array[RandomIntegerwithinRange], &unsorted_array[end_index]);
	
	return PartitionData(unsorted_array, start_index , end_index);
}

func InplaceQuickSort(pointer_array []int, start_index int, end_index int) {

	if start_index < end_index  {

		pivot_index := RandomPartitionData(pointer_array, start_index, end_index)	
		fmt.Println(pointer_array, start_index, end_index, pivot_index);
		InplaceQuickSort(pointer_array, start_index, pivot_index - 1)
		InplaceQuickSort(pointer_array, pivot_index+1, end_index)

	} 

}

func QuickSort(unsorted_array []int) []int {
	
	InplaceQuickSort(unsorted_array, 0, len(unsorted_array) - 1)
 
	return unsorted_array
}

func main() {
	
	i:=6
	
	unsorted_array := []int{2,7,4,1,5,3}; 

	sorted_array := []int{};

	fmt.Printf("%v", unsorted_array);

	switch i {
    case 1:
        sorted_array = SelectionSort(unsorted_array);
    case 2:
		sorted_array = BubbleSort(unsorted_array);
    case 3:
        sorted_array = InsertionSort(unsorted_array);
	case 4:
		sorted_array = InsertionInPlaceSort(unsorted_array);
	case 5:
		sorted_array = MergeSort(unsorted_array);	
	case 6:
		sorted_array = QuickSort(unsorted_array);		
    }

	//HEAP SORT TO BE IMPLEMENTED
	fmt.Printf("%v", sorted_array);

}
