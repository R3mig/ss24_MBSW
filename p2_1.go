package main

import (
	"fmt"
	"reflect"
)

// Example  linked list
type node[T any] struct {
	val  T
	next *node[T]
}

func mkNode[T any](v T) *node[T] {
	return &node[T]{val: v, next: nil}
}

func insert[T any](n *node[T], v T) {
	n.next = mkNode[T](v)
}

// n > 0
func mkList[T any](n int, v T) *node[T] {

	head := mkNode(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert(current, v)
		current = current.next

	}
	return head

}

func len[T any](n *node[T]) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

// Generic translation

type node_G struct {
	val  interface{}
	next *node_G
}

func mkNode_G(v interface{}) *node_G {
	return &node_G{val: v, next: nil}
}

func insert_G(n *node_G, v interface{}) {
	n.next = mkNode_G(v)
}

func mkList_G(n int, v interface{}) *node_G {

	head := mkNode_G(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_G(current, v)
		current = current.next
	}

	return head
}

func len_G(n *node_G) int {
	i := 0

	for n != nil {
		i++
		n = n.next
	}
	return i
}

// Monomorphization

type node_int struct {
	val  int
	next *node_int
}

func mkNode_int(v int) *node_int {
	return &node_int{val: v, next: nil}
}

func insert_int(n *node_int, v int) {
	n.next = mkNode_int(v)
}

func mkList_int(n int, v int) *node_int {

	head := mkNode_int(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_int(current, v)
		current = current.next

	}
	return head
}

func len_int(n *node_int) int {
	i := 0
	for n != nil {
		i++
		n = n.next
	}
	return i
}

// Tests

func testNode() {
	n := mkList[int](5, 1)

	fmt.Printf("\n Length: (%d)", len(n))
	i := 0
	for n != nil {
		fmt.Printf("\n %d: [ %v ]", i, n.val)
		n = n.next
		i++
	}
	fmt.Print("\n\n")
}

func testNode_G(l int, v interface{}) {
	n := mkList_G(l, v)

	fmt.Printf("\n Length: (%d)", len_G(n))
	fmt.Printf("\n Type: %v <-- %v", reflect.TypeOf(n.val), reflect.TypeOf(n))
	i := 0
	for n != nil {
		fmt.Printf("\n %d: [ %v ]", i, n.val)
		n = n.next
		i++
	}
	fmt.Print("\n\n")

}

func testNode_int(l int, v int) {
	n := mkList_int(l, v)

	fmt.Printf("\n Length: (%d)", len_int(n))
	fmt.Printf("\n Type: %v <-- %v", reflect.TypeOf(n.val), reflect.TypeOf(n))
	i := 0
	for n != nil {
		fmt.Printf("\n %d: [ %d ]", i, n.val)
		n = n.next
		i++
	}
	fmt.Print("\n\n")

}

// Example 2
func sum[T int | float32](xs []T) T {
	var x T
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x

}

func sum_G(xs interface{}) interface{} {
	switch xs := xs.(type) {
	case []int:
		var sum int
		for _, v := range xs {
			sum += v
		}
		return sum
	case []float32:
		var sum float32
		for _, v := range xs {
			sum += v
		}
		return sum
	default:
		return nil
	}
}

func sum_float32(xs []float32) float32 {
	var x float32
	x = 0
	for _, v := range xs {
		x = x + v
	}
	return x
}

func testSum_G_F(arr []float32) {
	for i, v := range arr {
		if i > 0 {
			print("+")
		}
		fmt.Printf(" %.1f ", v)
	}
	fmt.Printf(" = %.1f\n", sum_G(arr))

}
func testSum_G_I(arr []int) {
	for i, v := range arr {
		if i > 0 {
			print("+")
		}
		fmt.Printf(" %d ", v)
	}
	fmt.Printf(" = %d\n", sum_G(arr))

}

func testSum_f(arr []float32) {
	for i, v := range arr {
		if i > 0 {
			print("+")
		}
		fmt.Printf(" %.1f ", v)
	}
	fmt.Printf(" = %.1f\n", sum_float32(arr))

}

// Example 3
func swap[T any](x *T, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

func swap_G(x interface{}, y interface{}) {
	tmp := x
	x = y
	y = tmp
}

func swap_string(x *string, y *string) {
	tmp := *x
	*x = *y
	*y = tmp
}

func testSwap_G(x string, y string) {
	println("Generic swap:")
	fmt.Printf("Before swap: %s %s \n", x, y)
	swap_G(&x, &y)
	fmt.Printf("After swap: %s %s \n", x, y)
}

func testSwap_String(x string, y string) {
	println("Monomorph swap:")
	fmt.Printf("Before swap: %s %s \n", x, y)
	swap_string(&x, &y)
	fmt.Printf("After swap: %s %s \n", x, y)
}

func main() {
	l := 5
	test_i := 1337
	test_f := 1.337
	test_s := "LEET"
	test_arrF := []float32{1.1, 2.3, 9.9}
	test_arrI := []int{3, 3, 3, 1}
	s1 := "World!"
	s2 := "Hello"

	testNode()

	testNode_G(l, test_i)
	testNode_G(l, test_f)
	testNode_G(l, test_s)

	testNode_int(l, test_i)

	testSum_G_I(test_arrI)
	testSum_G_F(test_arrF)
	testSum_f(test_arrF)

	testSwap_G(s1, s2)
	testSwap_String(s1, s2)

}
