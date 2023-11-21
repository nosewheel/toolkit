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

// SliceCompare 判断两个切片是否相等, 去重后比较, 不考虑元素循序
func SliceCompare[T comparable](slice1, slice2 []T) bool {
	SliceRemoveDuplication(&slice1)
	SliceRemoveDuplication(&slice2)
	if len(slice1) != len(slice2) {
		return false
	}
	set1 := make(map[T]bool, len(slice1))
	for _, v := range slice1 {
		set1[v] = false
	}
	for _, v := range slice2 {
		if _, ok := set1[v]; ok {
			set1[v] = true
		} else {
			return false
		}
	}
	for _, v := range set1 {
		if !v {
			return false
		}
	}
	return true
}
