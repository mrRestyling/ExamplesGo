package main






func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)


    for i, num := range nums{
        final := target - num

        if q,ok := myMap[final]; ok{
            return []int{q,i}
        }
        myMap[num]= i
    }

return []int{}
}