package utils

// Convert slice to slice of pointers
func ToPointers[T any](arr []T) []*T {
	result := make([]*T, len(arr))
	for a := range len(arr) {
		result[a] = &arr[a]
	}

	return result
}
