package dsa

import (
	"errors"
	"fmt"
	"strings"
)

type DynamicArr struct {
	len   int
	cap   int
	Value []interface{}
}

func NewDynamicArr(len int) *DynamicArr {
	a := new(DynamicArr)
	a.cap = len
	a.Value = make([]interface{}, len)

	return a
}

func (d *DynamicArr) Get(pos int) interface{} {
	return d.Value[pos]
}
func (d *DynamicArr) Set(pos int, value interface{}) {
	d.Value[pos] = value
}
func (d DynamicArr) Size() int {
	return d.len
}
func (d DynamicArr) Capacity() int {
	return d.cap
}
func (d DynamicArr) IsEmpty() bool {
	return d.len == 0
}

func (d *DynamicArr) Add(value interface{}) {

	if d.len+1 >= d.cap {
		if d.cap == 0 {
			d.cap = 1
		} else {
			d.cap *= 2
		}

		newArr := make([]interface{}, d.cap)
		fmt.Println(d.len+1, d.cap, newArr)

		for i := 0; i < d.len; i++ {
			newArr[i] = d.Value[i]
		}
		d.Value = newArr
	}

	d.Value[d.len] = value
	d.len++
}

func (d *DynamicArr) RemoveAt(pos int) (interface{}, error) {
	if pos >= d.len && pos < 0 {
		return nil, errors.New("index out of bounds Exception")
	}

	data := d.Value[pos]

	newArr := make([]interface{}, d.len-1)

	for i, j := 0, 0; i < d.len; i, j = i+1, j+1 {
		if i == pos {
			j--
		} else {
			newArr[j] = d.Value[i]
		}
	}

	d.Value = newArr
	d.cap = d.len - 1
	d.len = d.cap

	return data, nil
}

func (d *DynamicArr) Remove(value interface{}) bool {
	for i := 0; i < d.len; i++ {
		if d.Value[i] == value {
			d.RemoveAt(i)
			return true
		}
	}
	return false
}

func (d *DynamicArr) IndexOf(value interface{}) int {
	for i := 0; i < d.len; i++ {
		if d.Value[i] == value {

			return i
		}
	}
	return -1
}

func (d *DynamicArr) Contains(value interface{}) bool {
	return d.IndexOf(value) != -1
}

func (d *DynamicArr) String() string {
	if d.len == 0 {
		return "[]"
	}

	var sb strings.Builder

	sb.WriteString("[")

	for i := 0; i < d.len-1; i++ {
		sb.WriteString(fmt.Sprintf("%v", (d.Value[i])) + ", ")
	}

	sb.WriteString(fmt.Sprintf("%v", d.Value[d.len-1]) + "]")

	return sb.String()

}
