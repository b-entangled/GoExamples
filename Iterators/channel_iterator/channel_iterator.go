package channeliterator

// GetData :- Get Any Data in batches (For Example Data from API's or Databases)
func GetData(n int)[]int{
	var ret = make([]int, 0)
	for i:=1; i<=n; i++{
		ret = append(ret, i)
	}
	return ret
}

// Iterate :- Iterate over the data and send it to channel until end of data
func Iterate(n int) <-chan int{
	var ret = make(chan int)
	go func(n int){
		data := GetData(n)
		for _, d := range data{
			ret <- d
		}
		close(ret)
	}(n)
	return ret
}