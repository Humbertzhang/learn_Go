/*
创建切片方式:
1.  var x [] float64
2.  x := make([]float64, 5)
    x := make([]float64, 5, 10) 10代表潜在容量
3.  arr := [5]float64{1, 2, 3, 4, 5}
    x := arr[0:5]
    can omit low or hight or low&high
    x := arr[:5]
    x := arr[0:]
    x := arr[:]
*/
