package assignment2

// countduplicates takes an array of strings as input,
// returns a map of respective strings and their
// number of occurrences in the list
func countduplicates(arr []string) map[string]int {
	arritem := make(map[string]int) // map of items in list, and their number

	for _, i := range arr { // loop through whole array
		arritem[i]++ // add occurrences to the counter map
	}

	return arritem
}
