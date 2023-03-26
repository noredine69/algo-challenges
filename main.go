package main

import "fmt"

func main() {
	values := []int{9, 1, 8, 2, 9, 2, 1, 7, 7}
	fmt.Println("Result : ", findLonelyInteger(values))

	changes := makeChanges(27, []int{25, 10, 5, 1})
	fmt.Printf("Result : %v \n", changes)

	changes = makeChanges(79, []int{50, 25, 10, 5, 1})
	fmt.Printf("Result : %v \n", changes)

	valuesBinSearch := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	index := binarySearch(valuesBinSearch, 8)
	fmt.Printf("Result : %d %d %d \n", 8, index, valuesBinSearch[index])

	index = binarySearchIterative(valuesBinSearch, 8)
	fmt.Printf("Result : %d %d %d \n", 8, index, valuesBinSearch[index])
	list := newLinkedList()
	list.walk()
	list.add(1, "A")
	list.add(2, "B")
	list.add(3, "C")
	list.add(4, "D")
	list.add(5, "E")
	list.walk()
	list.delete(5)
	list.walk()
	list.delete(3)
	list.walk()
	list.delete(1)
	list.walk()
	list.delete(2)
	list.delete(4)
	list.walk()
	fmt.Printf("Result : %v\n", list)
	list.add(6, "F")
	list.add(7, "G")
	list.add(8, "H")
	list.add(9, "I")
	list.add(10, "J")
	list.walk()

	h := newHashMap(5)
	h.walk()
	h.put("A", 1)
	h.put("AB", 2)
	h.put("ABC", 3)
	h.put("ABCD", 4)
	h.put("ABCDE", 5)
	h.put("ABCDEF", 6)
	h.put("ABCDEFG", 7)
	h.put("ABCDEFGH", 8)
	h.put("ABCDEFGHI", 9)
	h.put("ABCDEFGHIJ", 10)
	h.put("ABCDEFGHIJK", 11)
	h.walk()

	fmt.Printf("Search : %s %d \n", "ABCDEF", h.get("ABCDEF"))
	fmt.Printf("Search : %s %d \n", "A", h.get("A"))
	fmt.Printf("Search : %s %d \n", "ABCD", h.get("ABCD"))
	fmt.Printf("Search : %s %d \n", "ABCDEFGHIJK", h.get("ABCDEFGHIJK"))
	fmt.Printf("Search : %s %d \n", "[]", h.get(""))
	fmt.Printf("Search : %s %d \n", "eolknceolicx", h.get("eolknceolicx"))

	fmt.Printf("Fibonacci(%d) = %d \n", 7, fibonacci(7))
	fmt.Printf("Fibonacci(%d) = %d \n", 10, fibonacci(10))
}

func fibonacci(value int) int {
	mapFibo := make(map[int]int)
	return fibonacciRecursive(value, mapFibo)
}

func fibonacciRecursive(value int, mapFibo map[int]int) int {
	if value <= 0 {
		return 0
	}
	if value == 1 {
		return 1
	} else {
		valMap, exists := mapFibo[value]
		if exists {
			return valMap
		} else {
			mapFibo[value] = fibonacciRecursive(value-1, mapFibo) + fibonacciRecursive(value-2, mapFibo)
		}
	}
	return mapFibo[value]
}

func findLonelyInteger(values []int) int {
	result := 0
	for _, val := range values {
		result ^= val
	}
	return result
}

func makeChanges(valueUSD int, coins []int) map[int]int {
	changes := make(map[int]int)
	rest := valueUSD
	for _, coin := range coins {
		rest = makeChange(rest, coin, changes)
		if rest == 0 {
			break
		}
	}
	return changes
}

func makeChange(valueUSD int, coin int, changes map[int]int) int {
	changes[coin] = valueUSD / coin
	return valueUSD % coin
}

func binarySearch(values []int, value int) int {
	return binarySearchRecursive(values, value, 0, len(values)-1)
}

func binarySearchRecursive(values []int, value, left, right int) int {
	if left == right {
		fmt.Printf("left == right %v \n", left)
		return left
	}
	mid := left + (right-left)/2
	if value <= values[mid] {
		return binarySearchRecursive(values, value, left, mid)
	} else {
		return binarySearchRecursive(values, value, mid+1, right)
	}
}

func binarySearchIterative(values []int, value int) int {
	left := 0
	right := len(values) - 1
	mid := left + (right-left)/2
	for ok := true; ok; ok = left != right {
		mid = left + (right-left)/2
		if value <= values[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

type HashMap struct {
	mapList []*LinkedList
	size    int
}

func newHashMap(size int) HashMap {
	hashMap := HashMap{mapList: make([]*LinkedList, size), size: size}
	for index := 0; index < size; index++ {
		hashMap.mapList[index] = newLinkedList()
	}
	return hashMap
}

func (h *HashMap) walk() {
	fmt.Printf("HashMap walk ==> %d\n", len(h.mapList))
	for _, list := range h.mapList {
		list.walk()
	}
	fmt.Println("<== HashMap walk")
}

func (h *HashMap) hash(key string) int {
	return len(key) % h.size
}

func (h *HashMap) put(key string, value int) {
	hash := h.hash(key)
	list := h.mapList[hash]
	list.add(value, key)
	list.walk()
}

func (h *HashMap) get(key string) int {
	hash := h.hash(key)
	list := h.mapList[hash]
	if list == nil {
		return -2
	}
	return list.search(key)

}

type Node struct {
	value int
	key   string
	next  *Node
}
type LinkedList struct {
	first *Node
	last  *Node
}

func (list *LinkedList) walk() {
	fmt.Printf("LinkedList walk ==> %v %v\n", list.first, list.last)
	currentNode := list.first
	for ok := currentNode != nil; ok; ok = currentNode != nil {
		fmt.Printf("Current value %v %v\n", currentNode.value, currentNode.next)
		currentNode = currentNode.next
	}
	fmt.Println("<== LinkedList walk")
}

func (list *LinkedList) search(key string) int {
	currentNode := list.first
	for ok := currentNode != nil; ok; ok = currentNode != nil {
		if currentNode.key == key {
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return -1
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) add(value int, key string) {
	fmt.Printf("Add value %d\n", value)
	node := Node{value: value, key: key}
	if list.first == nil {
		list.first = &node
		list.last = &node
		return
	}
	list.last.next = &node
	list.last = &node
}

func (list *LinkedList) delete(value int) {
	fmt.Printf("Delete value %d %v %v\n", value, list.first, list.last)
	if list.first == nil {
		return
	}
	currentNode := list.first
	beforeNode := currentNode
	for ok := currentNode != nil; ok; ok = currentNode != nil {
		if currentNode.value == value {
			if currentNode == list.first {
				list.first = currentNode.next
				if currentNode == list.last {
					list.last = nil
				}
				return
			} else {
				beforeNode.next = currentNode.next
				if currentNode == list.last {
					list.last = beforeNode
				}
				return
			}
		}
		beforeNode = currentNode
		currentNode = currentNode.next
	}
}
