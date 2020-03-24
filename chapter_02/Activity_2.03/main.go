package main
import (
  "fmt"
)

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	var nums2 []int
	for j:=0; j < len(nums); j++ {
		for i := 1; i < len(nums); i++ {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
		nums2 = append(nums2, nums[len(nums)-1-j])
	}
	fmt.Println("Before:", nums)
	fmt.Println("After:", nums2)
}
