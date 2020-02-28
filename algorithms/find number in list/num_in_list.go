package main


func NumInList(list []int, num int) bool {

	n := len(list)

	for i := 0; i < n; i++ {

		if (list[i] == num) {
			return true
		}
	}

	return false

}
