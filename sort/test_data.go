package sort

var tables = []struct {
	desc   string
	in     []int
	expect []int
}{
	{
		"empty list",
		[]int{},
		[]int{},
	},
	{
		"single item",
		[]int{1234},
		[]int{1234},
	},
	{
		"odd number of items",
		[]int{25, 17, 20, 21, 100, 18, 22, 12, 5},
		[]int{5, 12, 17, 18, 20, 21, 22, 25, 100},
	},
	{
		"even number of items",
		[]int{25, 17, 20, 100, 18, 22, 12, 5},
		[]int{5, 12, 17, 18, 20, 22, 25, 100},
	},
	{
		"sorted list",
		[]int{10, 20, 25, 100, 150, 178, 200},
		[]int{10, 20, 25, 100, 150, 178, 200},
	},
	{
		"reverse sorted list",
		[]int{200, 178, 170, 100, 90, 87, 24},
		[]int{24, 87, 90, 100, 170, 178, 200},
	},
	{
		"duplicate item",
		[]int{25, 10, 45, 10, 19, 90, 100},
		[]int{10, 10, 19, 25, 45, 90, 100},
	},
	{
		"all duplicate items",
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	},
}
