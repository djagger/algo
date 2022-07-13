package structs

const growFactor = 10

type FactorArray[T any] struct {
	VectorArray[T]
}

func (a *FactorArray[T]) grow() {
	newSize := a.Size()*growFactor + 1
	newAllocatedArray := make([]T, newSize)

	copy(newAllocatedArray, a.array)

	a.array = newAllocatedArray
}
