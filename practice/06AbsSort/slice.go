/*
创建切片方式:
1.  var x [] float64
    生成一个不指定大小的slice

2.  x := make([]float64, 5)
    x := make([]float64, 5, 10) 10代表潜在容量
    Init一个指定大小的slice

3.  arr := [5]float64{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3}
    生成一个指派数据的slice

4.  x := arr[0:5]
    can omit low or hight or low&high
    x := arr[:5]
    x := arr[0:]
    x := arr[:]
    根据之前slice生成另一个slice
*/
