package main

import (
	"fmt"
)

func main() {
	data := []int{22, 99, 27, 175, 300, 201, 1, 22, 56, 88, 3, 176, 22, 9, 73, 13, 11, 22, 143, 88}

	data = MergeSortRecursive(data)
	fmt.Printf("sorted data: %v\n", data)

	l, r := RepeatValueBinarySearch(data, 88)
	// fmt.Printf("target index: %v\n", BinarySearch(data, 22))
	fmt.Printf("left index: %v\nright index: %v\n", l, r)
}

// 遞歸作法
func MergeSortRecursive(inputData []int) []int {

	if len(inputData) < 2 {
		return inputData
	}

	var middle = len(inputData) / 2
	// 經由上面的if判斷式，只剩一個element時，就不會再切分
	var a = MergeSortRecursive(inputData[:middle])
	var b = MergeSortRecursive(inputData[middle:])
	return MyMergeFunc(a, b)
}

// 迭代作法
func MergeSortIterate(inputData []int) []int {
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
func MergeExample(array1 []int, array2 []int) []int {
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

func MergeSortRecursiveExample(items []int) []int {
	fmt.Printf("all items: %v\n", items)
	// 傳入array的長度只有1或0時，就直接回傳
	if len(items) < 2 {
		fmt.Printf("items: %v\n", items)
		return items
	}

	var middle = len(items) / 2
	// 經由上面的if判斷式，只剩一個element時，就不會再切分
	var a = MergeSortRecursiveExample(items[:middle])
	var b = MergeSortRecursiveExample(items[middle:])
	return MergeExample(a, b)
}

func BinarySearch(data []int, key int) (index int) {
	fmt.Printf("BinarySearch sorted data: %v, target data: %v\n", data, key)
	low := 0
	upper := len(data) - 1

	// 閉區間
	// 等於的時候要在比較一次該數值
	for low <= upper {
		// 無條件捨去，只取商數
		// low: 0, upper: 11, targetIndex會等於5
		targetIndex := (low + upper) / 2
		fmt.Printf("low: %v, upper: %v, targetIndex: %v\n", low, upper, targetIndex)

		if low == upper {
			v := data[targetIndex]
			if v == key {
				return targetIndex
			}
			// 找不到正確資料，回傳相差最小的資料
			// 比較此時的target index的值，如果比target小，那就要找target index + 1 中的值再比一次，比較大則相反
			switch {
			case v == key:
				return targetIndex
			case v < key:
				// avoid out of slice index
				if targetIndex+1 > len(data)-1 {
					return targetIndex
				}
				// 如果 v < key，就要再找一個比key大的，形成 v < key < v2，再比大小
				v2 := data[targetIndex+1]
				if key-v > v2-key {
					// v2和key相差的值最小，回傳v2的index值
					return targetIndex + 1
				} else {
					return targetIndex
				}
			case v > key:
				// avoid out of slice index
				if targetIndex-1 < 0 {
					return targetIndex
				}
				// 如果 v > key，就要再找一個比key小的，形成 v2 < key < v，再比大小
				v2 := data[targetIndex-1]
				if v-key > key-v2 {
					// v2和key相差的值最小，回傳v2的index值
					return targetIndex - 1
				} else {
					return targetIndex
				}
			}
		}

		// targetIndex +- 1，target Index的值已經比較過，就直接排除掉。
		switch {
		case data[targetIndex] == key:
			// 重複值 todo
			return targetIndex
		case data[targetIndex] > key:
			upper = targetIndex - 1
			fmt.Printf("data[targetIndex] %v > targetData %v , upper: %v\n", data[targetIndex], key, upper)
		case data[targetIndex] < key:
			low = targetIndex + 1
			fmt.Printf("data[targetIndex] %v < targetData %v , low: %v\n", data[targetIndex], key, low)
		}
		fmt.Printf("for loop process over low: %v, upper: %v\n\n", low, upper)
	}
	// fmt.Printf("final low: %v, upper: %v\n", low, upper)
	// fmt.Print("data not found\n")

	// return low
	return -1
}

func RepeatValueBinarySearch(data []int, key int) (leftIndex int, rightIndex int) {
	low := 0
	upper := len(data) - 1
	fmt.Printf("BinarySearch sorted data: %v, target data: %v, len: %v\n", data, key, len(data))

	// 先找到重複的target value最小的index
	for low <= upper {
		mi := (low + upper) / 2
		switch {
		case data[mi] < key:
			low = mi + 1
		// 等於的時候也要往左移動upper index，因為最終要找的是>=target data中最小的index，直到不等於key值，再加一回去就是index值最小的重複值
		case data[mi] >= key:
			upper = mi - 1
		}
	}
	// leftIndex = upper + 1
	leftIndex = low

	low2 := 0
	upper2 := len(data) - 1

	// 找到不等於重複的target value最小的index
	for low2 <= upper2 {
		mi := (low2 + upper2) / 2
		switch {
		// 等於的時候也要往右移動index(加一)，直到不等於key值，再減一回去就是index值最小的重複值
		case data[mi] <= key:
			low2 = mi + 1
		case data[mi] > key:
			upper2 = mi - 1
		}
	}
	// rightIndex = low2 - 1
	rightIndex = upper2

	// 再找到大於target value最小的index
	return
}
