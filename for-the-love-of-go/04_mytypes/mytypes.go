package mytypes

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}
