package heap

type Heap interface {
	DataStructure

	Insert(val int)
	Extract() int
}

type DataStructure interface {
	DataStructurePrinter

	Size() int
	IsEmpty() bool
}

type DataStructurePrinter interface {
	Println()
}
