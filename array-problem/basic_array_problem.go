package arrayproblem

func LargestElem(arr []int) int {
	l := arr[0]

	for _, v := range arr {
		if l < v {
			l = v
		}
	}

	return l
}

func MaximuxConsOnes(arr []int) int {
	f := 0
	count := 0
	tempC := 0
	for f < len(arr) {
		if arr[f] == 1 {
			tempC++
		} else {
			if count < tempC {
				count = tempC
			}
			tempC = 0
		}
		f++
	}

	return  count
}

func MissingNum(arr []int) int {
	xor1 := 1
	for i := 2; i <= len(arr)+1; i++ {
		xor1 = xor1 ^ i
	}

	xor2 := arr[0]
	for i := 1; i < len(arr); i++ {
		xor2 = xor2 ^ arr[i]
	}

	return xor1 ^ xor2
}

func SecoundElem(arr []int) int {
	l := arr[0]
	sl := 0

	for _, v := range arr {
		if l < v {
			sl = l
			l = v

		} else if sl < v && sl < l {
			sl = v
		}
	}

	return sl
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func RemoveDuplicateFromSorted(arr []int) []int {
	f := 0
	s := f + 1
	for s < len(arr) {
		if arr[f] == arr[s] {
			s++
		} else {
			f++
			arr[f] = arr[s]
		}
	}
	return arr[:f+1]
}


func FindAppareOnce(arr[]int)int {
	n := 0 
	for _, v := range arr {
		n = n^v
	}

	return  n
}