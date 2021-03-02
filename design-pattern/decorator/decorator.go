package decorator

type Calculate interface {
	Cal() int
}

// OriCalculate: create base struct
type OriCalculate struct {
	num int
}

func NewOriCalculate(num int) *OriCalculate {
	return &OriCalculate{
		num: num,
	}
}

func (o *OriCalculate) Cal() int {
	return o.num
}

// MutCalculate: create mul base with struct
type MutCalculate struct {
	Calculate
	num int
}

func NewMutCalculate(C Calculate, num int) *MutCalculate {
	return &MutCalculate{
		Calculate: C,
		num:       num,
	}
}

func (m *MutCalculate) Cal() int {
	return m.num * m.Calculate.Cal()
}

// AddCalculate: create add base with struct
type AddCalculate struct {
	Calculate
	num int
}

func NewAddCalculate(C Calculate, num int) *AddCalculate {
	return &AddCalculate{num: num, Calculate: C}
}

func (a *AddCalculate) Cal() int {
	return a.num + a.Calculate.Cal()
}
