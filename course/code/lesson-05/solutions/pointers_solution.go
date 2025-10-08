package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// Exercise 1: Pointer Basics

func DoubleValue(ptr *int) {
	*ptr *= 2
}

func SwapStrings(a, b *string) {
	*a, *b = *b, *a
}

func GetLargerPointer(a, b *int) *int {
	if *a >= *b {
		return a
	}
	return b
}

// Exercise 2: Function Parameters with Pointers

func IncrementCounter(counter *int) int {
	*counter++
	return *counter
}

func AppendString(target *string, suffix string) {
	*target += suffix
}

func FindMinMax(numbers []int, min, max *int) {
	if len(numbers) == 0 {
		return
	}

	*min = numbers[0]
	*max = numbers[0]

	for _, num := range numbers[1:] {
		if num < *min {
			*min = num
		}
		if num > *max {
			*max = num
		}
	}
}

// Exercise 3: Struct Methods with Pointer Receivers

type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) bool {
	if ba.Balance >= amount {
		ba.Balance -= amount
		return true
	}
	return false
}

func (ba BankAccount) GetAccountInfo() string {
	return fmt.Sprintf("Account: %s, Owner: %s, Balance: $%.2f",
		ba.AccountNumber, ba.Owner, ba.Balance)
}

func (ba *BankAccount) TransferTo(target *BankAccount, amount float64) bool {
	if ba.Balance >= amount {
		ba.Balance -= amount
		target.Balance += amount
		return true
	}
	return false
}

// Exercise 4: Memory Management and Allocation

func AllocateInt(value int) *int {
	ptr := new(int)
	*ptr = value
	return ptr
}

func CreatePointerSlice(size int) []*int {
	ptrs := make([]*int, size)
	for i := 0; i < size; i++ {
		ptrs[i] = new(int)
		*ptrs[i] = i
	}
	return ptrs
}

func SafeDereference(ptr *int, defaultValue int) int {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

// Exercise 5: Linked List Implementation

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

func (ll *LinkedList) PrependNode(value int) {
	newNode := &ListNode{Value: value, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList) AppendNode(value int) {
	newNode := &ListNode{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList) RemoveValue(value int) bool {
	if ll.Head == nil {
		return false
	}

	// Remove from head
	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}

	// Remove from middle or end
	current := ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}

	return false
}

func (ll *LinkedList) FindNode(value int) *ListNode {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList) ToSlice() []int {
	result := make([]int, 0, ll.Size)
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Exercise 6: Performance Optimization

type LargeData struct {
	Numbers [1000]int
	Text    string
	Active  bool
}

func ProcessByValue(data LargeData) int {
	sum := 0
	for _, num := range data.Numbers {
		sum += num
	}
	return sum
}

func ProcessByPointer(data *LargeData) int {
	sum := 0
	for _, num := range data.Numbers {
		sum += num
	}
	return sum
}

func InitializeLargeData(data *LargeData, text string) {
	for i := 0; i < 1000; i++ {
		data.Numbers[i] = i
	}
	data.Text = text
	data.Active = true
}

// Exercise 7: Advanced Pointer Patterns

func ParseNameAge(input string, name *string, age *int) bool {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return false
	}

	parsedAge, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return false
	}

	*name = strings.TrimSpace(parts[0])
	*age = parsedAge
	return true
}

func ModifyThroughPointers(ptrs []*int, increment int) {
	for _, ptr := range ptrs {
		if ptr != nil {
			*ptr += increment
		}
	}
}

func CreateCounter(initialValue int) *int {
	counter := new(int)
	*counter = initialValue
	return counter
}

// Exercise 8: Pointer Safety and Best Practices

func SafeStringCopy(src *string, dst *string) bool {
	if src == nil || dst == nil {
		return false
	}
	*dst = *src
	return true
}

func SumValidPointers(ptrs []*int) int {
	sum := 0
	for _, ptr := range ptrs {
		if ptr != nil {
			sum += *ptr
		}
	}
	return sum
}

func DeepCopySlice(original []int) []int {
	if original == nil {
		return nil
	}

	copy := make([]int, len(original))
	for i, v := range original {
		copy[i] = v
	}
	return copy
}
