package structs

const defaultVector = 2

type VectorArray[T any] struct {
	array  []T
	size   int
	vector int
}

func NewVectorArray[T any](vector int) VectorArray[T] {
	if vector == 0 {
		vector = defaultVector
	}

	return VectorArray[T]{
		array:  make([]T, vector),
		vector: vector,
	}
}

func (a VectorArray[T]) Size() int {
	return a.size
}

func (a *VectorArray[T]) Add(item T) {
	if a.Size() == len(a.array) {
		a.grow()
	}

	a.array[a.Size()] = item

	a.size++
}

func (a *VectorArray[T]) grow() {
	newSize := a.Size() + a.vector
	newAllocatedArray := make([]T, newSize)

	copy(newAllocatedArray, a.array)

	a.array = newAllocatedArray
}

func (a *VectorArray[T]) Remove(index int) T {
	res := a.array[index]

	copy(a.array[index:], a.array[index+1:])

	a.size--

	return res
}
