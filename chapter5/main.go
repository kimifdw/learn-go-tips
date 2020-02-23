package main

import (
	"fmt"
)

func main() {
	// 打印数组
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])

	//var y[5] float64
	//y[0] = 98
	//y[1] = 93
	//y[2] = 77
	//y[3] = 82
	//y[4] = 83
	y := [5]float64{
		98, 93, 77, 82, 83,
	} // 初始化数组的第2种写法，注意最后的逗号

	tmp := y[0]
	for i, value := range y {
		if tmp > value {
			tmp = value
		} else {
			i = i + 1
		}
	}
	fmt.Println("tmp=", tmp) // 找出y中的最小数字

	var total float64 = 0
	//for i := 0; i < len(y); i++ {
	//	total += y[i]
	//}
	for _, value := range y {
		total += value
	} // for循环的第2种写法，_表示索引
	fmt.Println(total / float64(len(y))) // 需要将int强制转换为float64

	// slices。数组的一部分，长度可变
	//s := make([]float64, 5, 10) // 初始化slices，长度为5,最大容量为10
	s1 := []float64{1, 2, 3, 4, 5} // 第二种初始化slices的方法
	z := s1[0:5]                   // 获取索引为0到5的slices内容
	fmt.Println(z)

	s2 := append(s1, 7, 8) // s1的slices在初始化时不能设置默认长度
	fmt.Println("append=", s2)

	s3 := make([]float64, 3)
	copy(s3, s1) // 第一个参数为目标,copy的类型必须一致
	fmt.Println("copy=", s3)

	// maps.
	m := make(map[string]int) // 定义map对象，[key]value
	m["key"] = 10
	// 返回m[key]的值，如果获取值返回成功，则打印信息
	if name, ok := m["key"]; ok {
		fmt.Println(name, ok) // name表示获取的值，ok表示值返回是否成功
	}
	delete(m, "key") // 根据key值删除值

}
