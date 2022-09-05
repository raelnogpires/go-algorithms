package challenges

func Is_palindrome_recursive(word string, low_index int, high_index int) bool {
	if word == "" {
		return false
	}

	if word[low_index] != word[high_index] {
		return false
	}

	return Is_palindrome_recursive(word, low_index+1, high_index-1)
}
