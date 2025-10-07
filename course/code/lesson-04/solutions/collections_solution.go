package solutions

import "sort"

// Exercise 1: Array Practice

func GetFirstFivePrimes() [5]int {
	return [5]int{2, 3, 5, 7, 11}
}

func FindMaxInArray(arr [10]int) (max int, index int) {
	max = arr[0]
	index = 0
	
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	
	return max, index
}

// Exercise 2: Slice Manipulation

func RemoveDuplicates(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	
	seen := make(map[int]bool)
	result := make([]int, 0, len(slice))
	
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	
	return result
}

func ReverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func MergeSortedSlices(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))
	i, j := 0, 0
	
	// Merge while both slices have elements
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] <= slice2[j] {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}
	
	// Add remaining elements from slice1
	for i < len(slice1) {
		result = append(result, slice1[i])
		i++
	}
	
	// Add remaining elements from slice2
	for j < len(slice2) {
		result = append(result, slice2[j])
		j++
	}
	
	return result
}

// Exercise 3: Map Operations

func CountCharacters(s string) map[rune]int {
	counts := make(map[rune]int)
	
	for _, char := range s {
		counts[char]++
	}
	
	return counts
}

func InvertMap(m map[string]int) map[int]string {
	inverted := make(map[int]string)
	
	for key, value := range m {
		inverted[value] = key
	}
	
	return inverted
}

func MergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	
	// Copy all entries from map1
	for key, value := range map1 {
		result[key] = value
	}
	
	// Copy all entries from map2 (overwrites duplicates)
	for key, value := range map2 {
		result[key] = value
	}
	
	return result
}

// Exercise 4: Advanced Collection Operations

func FindIntersection(slice1, slice2 []int) []int {
	set1 := make(map[int]bool)
	for _, v := range slice1 {
		set1[v] = true
	}
	
	var result []int
	seen := make(map[int]bool)
	
	for _, v := range slice2 {
		if set1[v] && !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	
	sort.Ints(result)
	return result
}

func GroupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)
	
	for _, word := range words {
		length := len(word)
		groups[length] = append(groups[length], word)
	}
	
	return groups
}

// Exercise 5: Real-World Scenario - Inventory Management

type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

type Inventory struct {
	products map[string]*Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]*Product),
	}
}

func (inv *Inventory) AddProduct(product *Product) {
	inv.products[product.ID] = product
}

func (inv *Inventory) GetProduct(id string) (*Product, bool) {
	product, exists := inv.products[id]
	return product, exists
}

func (inv *Inventory) UpdateStock(id string, newStock int) bool {
	product, exists := inv.products[id]
	if !exists {
		return false
	}
	
	product.Stock = newStock
	return true
}

func (inv *Inventory) GetLowStockProducts(threshold int) []*Product {
	var lowStock []*Product
	
	for _, product := range inv.products {
		if product.Stock < threshold {
			lowStock = append(lowStock, product)
		}
	}
	
	return lowStock
}

func (inv *Inventory) GetTotalValue() float64 {
	var total float64
	
	for _, product := range inv.products {
		total += product.Price * float64(product.Stock)
	}
	
	return total
}

// Exercise 6: Memory Efficiency Challenge

func EfficientAppend(initialSize int, values []int) []int {
	// Pre-allocate slice with proper capacity
	result := make([]int, 0, initialSize)
	
	// Append all values
	result = append(result, values...)
	
	return result
}

func ProcessInChunks(data []int, chunkSize int, processor func([]int) int) []int {
	var results []int
	
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		
		chunk := data[i:end]
		result := processor(chunk)
		results = append(results, result)
	}
	
	return results
}
