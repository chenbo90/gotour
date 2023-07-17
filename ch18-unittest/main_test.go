package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	//	y预先定义一组斐波那契数列作为测试用例
	fsMap := map[int]int{}
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	//fsMap[8] = 22 //21

	for k,v:=range fsMap{
		fsv := Fibonacci(k)
		if v == fsv{
			t.Logf("结果正确：n为%d,值为%d",k,fsv)
		}else{
			t.Errorf("结果错误：n为%d，期望值:%d,但是计算结果为：%d",k,v,fsv)
		}
	}


	//测试	go test -v ./ch18-unittest
	/**
	localhost:gotour chenbo$ go test -v ./ch18-unittest/
	=== RUN   TestFibonacci
	    main_test.go:23: 结果正确：n为0,值为0
	    main_test.go:23: 结果正确：n为1,值为1
	    main_test.go:23: 结果正确：n为3,值为2
	    main_test.go:23: 结果正确：n为5,值为5
	    main_test.go:23: 结果正确：n为6,值为8
	    main_test.go:23: 结果正确：n为2,值为1
	    main_test.go:23: 结果正确：n为4,值为3
	    main_test.go:23: 结果正确：n为7,值为13
	    main_test.go:25: 结果错误：n为8，期望值:22,但是计算结果为：21
	--- FAIL: TestFibonacci (0.00s)
	FAIL
	FAIL    demo/ch18-unittest      0.084s
	FAIL
	localhost:gotour chenbo$ go test -v ./ch18-unittest/
	*/

	/**
	1：含有单元测试代码的go文件必须以_test.go结尾
	2：单元测试文件名_test.go前面部分最好是被测的函数所在的go文件的文件名
	3：单元测试的函数必须以Test开头，是可导出的、公开的函数
	4：测试函数的签名必须接收一个指向testing.T类型的指针，并且不能返回任何值
	5：函数名最好是Test+要测试的函数名
	 */

	//单元测试的覆盖率
	//go test -v --coverprofile=ch18.cover ./ch18-unittest
	//go tool cover -html=ch18.cover -o=ch18.html



}

/**
2:基准测试
1：基准测试函数必须是以Benchmark开头，必须是可导出的
2：函数的签名必须接收一个指向testing.B类型的指针，且不能返回任何值
3：最后的for循环很重要，被测试的代码要放到循环里
4：b.N是基准测试框架提供的，表示循环次数，因需要反复调用测试代码才可评估性能

go test -bench=../ch18-unittest/
*/
func BenchmarkFibonacci(b *testing.B){
	for i:=0; i<b.N;i++  {
		Fibonacci(i)
	}
}
