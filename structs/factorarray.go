package structs

type FactorArray[T any] struct {
	VectorArray[T]
	factor int
}

func NewFactorArray[T any](factor int) FactorArray[T] {
	return FactorArray[T]{
		factor: factor,
	}
}

func (a *FactorArray[T]) Add(item T) {
	if a.Size() == len(a.array) {
		a.grow()
	}

	a.array[a.Size()] = item

	a.size++
}

func (a *FactorArray[T]) grow() {
	newSize := a.Size()*a.factor + 1
	newAllocatedArray := make([]T, newSize)

	copy(newAllocatedArray, a.array)

	a.array = newAllocatedArray
}
