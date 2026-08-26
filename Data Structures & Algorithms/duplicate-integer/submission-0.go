func hasDuplicate(nums []int) bool {
    hash_set := make(map[int]struct{})
    for _, val := range nums {
        if _, exists := hash_set[val]; exists {
            return true
        }
        hash_set[val] = struct{}{}
    }

    return false
}
