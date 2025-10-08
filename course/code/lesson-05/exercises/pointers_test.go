package exercises

import (
	"reflect"
	"strings"
	"testing"
)

// Test Pointer Basics
func TestDoubleValue(t *testing.T) {
	value := 21
	DoubleValue(&value)

	if value != 42 {
		t.Errorf("DoubleValue() failed: got %d, want 42", value)
	}
}

func TestSwapStrings(t *testing.T) {
	a := "hello"
	b := "world"

	SwapStrings(&a, &b)

	if a != "world" || b != "hello" {
		t.Errorf("SwapStrings() failed: got a=%s, b=%s, want a=world, b=hello", a, b)
	}
}

func TestGetLargerPointer(t *testing.T) {
	tests := []struct {
		a, b     int
		expected *int
	}{
		{5, 3, nil}, // Will be set to &a
		{2, 8, nil}, // Will be set to &b
		{7, 7, nil}, // Will be set to &a
	}

	for i, test := range tests {
		a, b := test.a, test.b
		result := GetLargerPointer(&a, &b)

		var expected *int
		if a >= b {
			expected = &a
		} else {
			expected = &b
		}

		if result != expected {
			t.Errorf("Test %d: GetLargerPointer(%d, %d) returned wrong pointer", i, test.a, test.b)
		}

		if *result != *expected {
			t.Errorf("Test %d: GetLargerPointer(%d, %d) = %d, want %d", i, test.a, test.b, *result, *expected)
		}
	}
}

// Test Function Parameters
func TestIncrementCounter(t *testing.T) {
	counter := 10
	result := IncrementCounter(&counter)

	if counter != 11 || result != 11 {
		t.Errorf("IncrementCounter() failed: counter=%d, result=%d, want both 11", counter, result)
	}
}

func TestAppendString(t *testing.T) {
	target := "Hello"
	AppendString(&target, " World")

	if target != "Hello World" {
		t.Errorf("AppendString() failed: got %s, want 'Hello World'", target)
	}
}

func TestFindMinMax(t *testing.T) {
	numbers := []int{5, 2, 8, 1, 9, 3}
	var min, max int

	FindMinMax(numbers, &min, &max)

	if min != 1 || max != 9 {
		t.Errorf("FindMinMax() failed: min=%d, max=%d, want min=1, max=9", min, max)
	}
}

// Test BankAccount methods
func TestBankAccountOperations(t *testing.T) {
	account := &BankAccount{
		AccountNumber: "12345",
		Balance:       100.0,
		Owner:         "John Doe",
	}

	// Test deposit
	account.Deposit(50.0)
	if account.Balance != 150.0 {
		t.Errorf("Deposit failed: balance=%f, want 150.0", account.Balance)
	}

	// Test successful withdrawal
	success := account.Withdraw(30.0)
	if !success || account.Balance != 120.0 {
		t.Errorf("Withdraw failed: success=%t, balance=%f, want true, 120.0", success, account.Balance)
	}

	// Test insufficient funds withdrawal
	success = account.Withdraw(200.0)
	if success || account.Balance != 120.0 {
		t.Errorf("Withdraw should fail: success=%t, balance=%f, want false, 120.0", success, account.Balance)
	}

	// Test account info
	info := account.GetAccountInfo()
	if !strings.Contains(info, "12345") || !strings.Contains(info, "John Doe") {
		t.Errorf("GetAccountInfo failed: %s", info)
	}
}

func TestBankAccountTransfer(t *testing.T) {
	account1 := &BankAccount{AccountNumber: "111", Balance: 100.0, Owner: "Alice"}
	account2 := &BankAccount{AccountNumber: "222", Balance: 50.0, Owner: "Bob"}

	// Test successful transfer
	success := account1.TransferTo(account2, 30.0)
	if !success || account1.Balance != 70.0 || account2.Balance != 80.0 {
		t.Errorf("Transfer failed: success=%t, balance1=%f, balance2=%f", success, account1.Balance, account2.Balance)
	}

	// Test insufficient funds transfer
	success = account1.TransferTo(account2, 100.0)
	if success || account1.Balance != 70.0 || account2.Balance != 80.0 {
		t.Errorf("Transfer should fail: success=%t, balance1=%f, balance2=%f", success, account1.Balance, account2.Balance)
	}
}

// Test Memory Management
func TestAllocateInt(t *testing.T) {
	ptr := AllocateInt(42)

	if ptr == nil {
		t.Error("AllocateInt returned nil")
		return
	}

	if *ptr != 42 {
		t.Errorf("AllocateInt failed: got %d, want 42", *ptr)
	}
}

func TestCreatePointerSlice(t *testing.T) {
	size := 5
	ptrs := CreatePointerSlice(size)

	if len(ptrs) != size {
		t.Errorf("CreatePointerSlice length: got %d, want %d", len(ptrs), size)
		return
	}

	for i, ptr := range ptrs {
		if ptr == nil {
			t.Errorf("Pointer at index %d is nil", i)
			continue
		}
		if *ptr != i {
			t.Errorf("Pointer at index %d: got %d, want %d", i, *ptr, i)
		}
	}
}

func TestSafeDereference(t *testing.T) {
	// Test with valid pointer
	value := 42
	result := SafeDereference(&value, 0)
	if result != 42 {
		t.Errorf("SafeDereference with valid pointer: got %d, want 42", result)
	}

	// Test with nil pointer
	result = SafeDereference(nil, 99)
	if result != 99 {
		t.Errorf("SafeDereference with nil pointer: got %d, want 99", result)
	}
}

// Test Linked List
func TestLinkedListOperations(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList returned nil")
	}

	// Test append
	ll.AppendNode(1)
	ll.AppendNode(2)
	ll.AppendNode(3)

	if ll.Size != 3 {
		t.Errorf("LinkedList size after appends: got %d, want 3", ll.Size)
	}

	// Test prepend
	ll.PrependNode(0)
	if ll.Size != 4 {
		t.Errorf("LinkedList size after prepend: got %d, want 4", ll.Size)
	}

	// Test find
	node := ll.FindNode(2)
	if node == nil || node.Value != 2 {
		t.Error("FindNode failed to find existing value")
	}

	node = ll.FindNode(99)
	if node != nil {
		t.Error("FindNode should return nil for non-existent value")
	}

	// Test to slice
	slice := ll.ToSlice()
	expected := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("ToSlice: got %v, want %v", slice, expected)
	}

	// Test remove
	success := ll.RemoveValue(2)
	if !success || ll.Size != 3 {
		t.Errorf("RemoveValue failed: success=%t, size=%d", success, ll.Size)
	}

	success = ll.RemoveValue(99)
	if success {
		t.Error("RemoveValue should return false for non-existent value")
	}
}

// Test Performance functions
func TestProcessByValue(t *testing.T) {
	data := LargeData{Text: "test", Active: true}
	for i := 0; i < 1000; i++ {
		data.Numbers[i] = i
	}

	result := ProcessByValue(data)
	expected := 499500 // Sum of 0 to 999

	if result != expected {
		t.Errorf("ProcessByValue: got %d, want %d", result, expected)
	}
}

func TestProcessByPointer(t *testing.T) {
	data := LargeData{Text: "test", Active: true}
	for i := 0; i < 1000; i++ {
		data.Numbers[i] = i
	}

	result := ProcessByPointer(&data)
	expected := 499500 // Sum of 0 to 999

	if result != expected {
		t.Errorf("ProcessByPointer: got %d, want %d", result, expected)
	}
}

func TestInitializeLargeData(t *testing.T) {
	var data LargeData
	InitializeLargeData(&data, "initialized")

	if data.Text != "initialized" || !data.Active {
		t.Errorf("InitializeLargeData failed: Text=%s, Active=%t", data.Text, data.Active)
	}

	for i := 0; i < 1000; i++ {
		if data.Numbers[i] != i {
			t.Errorf("InitializeLargeData: Numbers[%d]=%d, want %d", i, data.Numbers[i], i)
			break
		}
	}
}
