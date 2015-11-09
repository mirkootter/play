package task

// Implementation details of given interfaces (needed for test code to run).

// e implements Data<int> interface for a single element.
type e int

func (_ e) IsCollection() bool                   { return false }
func (_ e) GetCollection() CollectionᐸDataᐸintᐳᐳ { panic("not collection") }
func (e e) GetElement() int                      { return int(e) }

// c implements Data<int> interface for a collection of Data<int>.
type c collectionᐸDataᐸintᐳᐳ

func (_ c) IsCollection() bool                   { return true }
func (c c) GetCollection() CollectionᐸDataᐸintᐳᐳ { return collectionᐸDataᐸintᐳᐳ(c) }
func (_ c) GetElement() int                      { panic("not element") }

// collectionᐸDataᐸintᐳᐳ implements Collection<E> where E is Data<int>.
type collectionᐸDataᐸintᐳᐳ []Dataᐸintᐳ

func (c collectionᐸDataᐸintᐳᐳ) IsEmpty() bool { return len(c) == 0 }
func (c collectionᐸDataᐸintᐳᐳ) Iterator() IteratorᐸDataᐸintᐳᐳ {
	return &collectionIteratorᐸDataᐸintᐳᐳ{C: c}
}
func (c collectionᐸDataᐸintᐳᐳ) Size() int { return len(c) }

// collectionIteratorᐸDataᐸintᐳᐳ implements Iterator<E> where E is Data<int> for Collection<E>.
type collectionIteratorᐸDataᐸintᐳᐳ struct {
	C     collectionᐸDataᐸintᐳᐳ // Collection being iterated.
	index int                   // Index of the next element.
}

func (it *collectionIteratorᐸDataᐸintᐳᐳ) HasNext() bool {
	return it.index < len(it.C)
}

func (it *collectionIteratorᐸDataᐸintᐳᐳ) Next() Dataᐸintᐳ {
	e := it.C[it.index]
	it.index++
	return e
}
