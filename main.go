package main

import (
	"fmt"
)

func main() {
	elements := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	re1 := MyMergeSort2(elements)
	fmt.Print(re1)
}

// 遞歸作法
func MyMergeSort2(inputData []int) []int {

	if len(inputData) < 2 {
		return inputData
	}

	var middle = len(inputData) / 2
	// 經由上面的if判斷式，只剩一個element時，就不會再切分
	var a = MyMergeSort2(inputData[:middle])
	var b = MyMergeSort2(inputData[middle:])
	return MyMergeFunc(a, b)
}

// 迭代作法
func MyMergeSort(inputData []int) []int {
	l := len(inputData)
	fmt.Printf("len: %v\n", l)
	arrarr := make([][]int, 0, l)
	// 切分為一個一個只有單一element的slice
	for _, v := range inputData {
		arrarr = append(arrarr, []int{v})
	}
	fmt.Printf("arrarr: %v\n", arrarr)
	// 倆倆比較，合併，假設有8個元素，就會比較4次
	// 最後一個若沒有可以比較的，就不合併直接放入新的資料陣列
	// 8 變 4，9 變 5，所以Cap都先使用除以二再加一
	for l != 1 {
		newArray := make([][]int, 0, l/2+1)
		for i := 0; i <= l/2; i++ {
			fmt.Printf("i: %v\n", i)
			// 避免out of index
			// input Data是奇數個，len=9，9/2取4，2i剛好是最後一個element，index: 8，第2i+1會out of index
			// input Data是偶數個，len=8，8/2取4，2i就會out of index，最大index為7，會out of index
			if l%2 != 0 && 2*i+1 >= l {
				if len(arrarr) == 2 {
					fmt.Printf("最後一次, 直接break\n")
					break
				}
				fmt.Printf("l= %v,2i+1= %v, break\n", l, 2*i+1)
				// 沒有第2i+1個element比較時，就直接將第2i個放進去
				newArray = append(newArray, arrarr[2*i])
				break
			}
			if l%2 == 0 && 2*i+1 >= l {
				if len(arrarr) == 2 {
					fmt.Printf("最後一次, 直接break\n")
					break
				}
				fmt.Printf("l= %v,2i+1= %v, break\n", l, 2*i+1)
				// 沒有第2i+1個element比較時，就直接將第2i個放進去
				// 偶數個沒有2i
				break
			}

			margedArray := MyMergeFunc(arrarr[2*i], arrarr[2*i+1])
			newArray = append(newArray, margedArray)
			fmt.Printf("newArray: %v\n", newArray)
		}
		// 重新設定參數，run again
		arrarr = newArray
		fmt.Printf("reset arrarr: %v\n", arrarr)
		l = len(arrarr)
	}

	return arrarr[0]
}

// 將兩個排序好的array合併，只有比較第一個element
func MyMergeFunc(arr1 []int, arr2 []int) (result []int) {
	times := len(arr1) + len(arr2)
	for i := 0; i < times; i++ {
		fmt.Printf("arr1: %v arr2: %v\n", arr1, arr2)
		if arr1[0] < arr2[0] {
			// 新增至result
			result = append(result, arr1[0])
			// 移除至原本arr
			arr1 = arr1[1:]
		} else {
			// 新增至result
			result = append(result, arr2[0])
			// 移除至原本arr
			arr2 = arr2[1:]
		}

		// fmt.Printf("arr1: %v\n", arr1)
		// fmt.Printf("arr2: %v\n", arr2)

		// 移除玩array中的第一個資料後，避免下次 index out of range
		// 其中一個 out of range 就 break for loop
		if len(arr1) == 0 {
			result = append(result, arr2...)
			// fmt.Printf("arr1 len = 0, result: %v\n", result)
			break
		}
		if len(arr2) == 0 {
			result = append(result, arr1...)
			// fmt.Printf("arr2 len = 0, result: %v\n", result)
			break
		}

		fmt.Printf("temp result: %v\n", result)
	}

	return
}

// 遞歸作法example
func merge(array1 []int, array2 []int) []int {
	var result = make([]int, len(array1)+len(array2))
	var i = 0
	var j = 0

	// i j都小於len
	for i < len(array1) && j < len(array2) {

		// 比大小
		if array1[i] <= array2[j] {
			result[i+j] = array1[i]
			i++
		} else {
			result[i+j] = array2[j]
			j++
		}

	}

	// 各個array剩餘的element
	for i < len(array1) {
		result[i+j] = array1[i]
		i++
	}
	for j < len(array2) {
		result[i+j] = array2[j]
		j++
	}
	fmt.Printf("marge: %v & %v, result: %v\n", array1, array2, result)
	return result
}

func Mergesort(items []int) []int {
	fmt.Printf("all items: %v\n", items)
	// 傳入array的長度只有1或0時，就直接回傳
	if len(items) < 2 {
		fmt.Printf("items: %v\n", items)
		return items
	}

	var middle = len(items) / 2
	// 經由上面的if判斷式，只剩一個element時，就不會再切分
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)
}
