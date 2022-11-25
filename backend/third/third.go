package third

import (
	"fmt"
	"sort"
)

var (
	u = trace("var")
)

func init() {
	fmt.Println("third init")
}

func trace(s string) string {
	fmt.Println("third :", s)
	return s
}

type Sequerce []int

func (s Sequerce) Len() int {
	return len(s)
}

func (s Sequerce) Less(i, j int) bool {
	fmt.Println("i, j less", i, j, s[i] < s[j])
	return s[i] < s[j]
}

func (s Sequerce) Swap(i, j int) {
	fmt.Println("i, j swap", i, j)
	s[i], s[j] = s[j], s[i]
}

type Stringer interface {
	String() string
}

type CString struct {
	index string
}

func (c *CString) String() string {
	return fmt.Sprintf("%v", c.index)
}

type IntSlice []int

func (s *IntSlice) Add(data []int) {
	slice := *s
	slice = append(*s, data...)
	*s = slice
}

func Run(s []int) {
	sort.Sort(Sequerce(s))

	//var value interface{}
	var value1 Stringer = &CString{}
	var value interface{}
	value = value1
	switch value.(type) {
	case string:
		fmt.Println("string")
	case Stringer:
		fmt.Println("Stringer")
	case *CString:
		fmt.Println("CString")
	default:
		fmt.Println("default")
	}

	if str, ok := value.(string); ok {
		fmt.Println("is string", str, ok)
		//} else if str, ok := value.(Stringer); ok {
		//	fmt.Println("is Stringer", str, ok)
	} else if str, ok := value.(*int); ok {
		fmt.Println("is int", str, ok)
	} else if str, ok := value.(*CString); ok {
		fmt.Println("is CString", str, ok)
	} else {
		fmt.Println("is not string", value, ok)
	}

	sli := IntSlice{1, 1, 1, 1, 1, 1}
	fmt.Println("cap ", cap(sli))
	old := sli
	d1 := []int{4} //, 5, 6, 7, 8, 9, 10, 11, 12}
	sli.Add(d1)
	sli[0] = 5
	fmt.Println(sli)
	fmt.Println(old)
}
