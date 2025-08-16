package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	} else {
		return slice[index]
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	newSlice := []int{}
	if len(values) != 0 {
		for i := range values {
			newSlice = append(newSlice, values[i])
		}
	}

	return append(newSlice, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {

		prefixSlice := slice[:index]
		subfixSlice := slice[index+1:]

		slice = append(prefixSlice, subfixSlice...)
	} else if index == len(slice) {
		slice = slice[:index]
	}

	return slice
}
