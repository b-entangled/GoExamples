package funciterator

// GetData :- Get Any Data in batches (For Example Data from API's or Databases)
func GetData(n int) []int {
	var ret = make([]int, 0)
	for i := 1; i <= n; i++ {
		ret = append(ret, i)
	}
	return ret
}

// Iterator :- Iterate over the data and call callback function to process data one by one
func Iterator(iter []int, callback func(int)) {
	for _, val := range iter {
		callback(val)
	}
}
