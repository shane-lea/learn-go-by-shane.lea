### 基准测试

- 什么是基准测试
  被测试得代码运行 b.N次所需要得时间 
- 如何编写基准测试
  - BenchmarkXxxxx[方法名称](b *testing.B)
  - testing.B 使得我们可以访问隐性命名 b.N [cryptically named]
  - go test -bench=. / go test -bench="." [Windows Powershell]
    - TIPS: 基准测试顺序运行 上面的测试如果不通过会中断线程,基准测试不会被执行
- 结果解析：运行一次要 154.5纳秒 一共运行7782970次
    ```markdown
    goos: darwin
    goarch: amd64
    pkg: learn-go-by-shane.lea/iteration/v1
    cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
    BenchmarkRepeat
    BenchmarkRepeat-8   	 7782970	       154.5 ns/op
    PASS
    ```