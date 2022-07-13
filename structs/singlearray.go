package structs

type SingleArray[T any] struct {
	array []T
}

func NewSingleArray[T any]() SingleArray[T] {
	return SingleArray[T]{}
}

func (a SingleArray[T]) Size() int {
	return len(a.array)
}

func (a *SingleArray[T]) Add(item T) {
	a.grow()
	a.array[a.Size()-1] = item
}

func (a *SingleArray[T]) grow() {
	newSize := a.Size() + 1
	newAllocatedArray := make([]T, newSize)

	copy(newAllocatedArray, a.array)

	a.array = newAllocatedArray
}

func (a *SingleArray[T]) Remove(index int) T {
	res := a.array[index]

	newSize := a.Size() - 1
	newAllocatedArray := make([]T, newSize)

	copy(newAllocatedArray, a.array[:index])
	copy(newAllocatedArray[index:], a.array[index+1:])

	a.array = newAllocatedArray

	return res
}
