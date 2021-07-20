package main

func removeDuplicates(s string, k int) string {
	str := make([]byte, 0)
	stack := make([]int, 0)

	str = append(str, byte('#'))
	stack = append(stack, 1)

	// range over the string
	for i := 0; i < len(s); i++ {

		// if the byte is == byte before it
		if s[i] == str[len(str)-1] {
			// increment the top of the stack
			stack[len(stack)-1]++
		} else {
			// otherwise we append 1 to the stack
			stack = append(stack, 1)
			// and append the last char, for the next loop
			str = append(str, s[i])

			// fmt.Printf("%c",str)
		}

		// if the end of the stack is equal to k
		if stack[len(stack)-1] == k {
			// pop off the last elem
			stack = stack[:len(stack)-1]
			// pop off the last string
			str = str[:len(str)-1]

		}
	}

	res := make([]byte, 0)
	// fmt.Printf("%c", str)

	// range over the string stack
	for i, v := range str {
		// range over the number stack
		for j := 0; j < stack[i]; j++ {
			// write characters
			res = append(res, v)
		}
	}

	// convert to string, except for the first char
	// "every character following the first 1:"
	return string(res[1:])
}
