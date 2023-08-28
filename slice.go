package toolkit

// PaginateSlice paginate a slice with page and pageSize.
// Parameter page start with number 1.
func PaginateSlice[T any](slice []T, page int, pageSize int) []T {
	if pageSize <= 0 {
		return []T{}
	}
	if page <= 0 {
		page = 1
	}
	start := (page - 1) * pageSize
	if start > len(slice) {
		return []T{}
	}
	end := start + pageSize
	if end > len(slice) {
		end = len(slice)
	}
	return slice[start:end]
}

// SliceRemoveDuplication ...
func SliceRemoveDuplication[T comparable](slice *[]T) {
	if len(*slice) == 0 {
		return
	}
	set := make(map[T]struct{}, len(*slice))
	j := 0
	for _, v := range *slice {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
