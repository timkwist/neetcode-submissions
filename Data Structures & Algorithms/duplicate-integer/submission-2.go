func hasDuplicate(nums []int) bool {
    seen := map[int]bool {}
    for i := 0; i < len(nums); i++ {
        _, ok := seen[nums[i]]
        if ok {
            return true
        }
        seen[nums[i]] = true
    }
    return false
}
