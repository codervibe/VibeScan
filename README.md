#  震波 扫描工具

## 说明
* 使用 线程池 进行多线程全端口 扫描
~~~
Usage:   -i <ip>
         --ip <ip>
         -h --help
~~~
* IP地址 全端口扫描
~~~
go run main.go -i 192.168.1.1
~~~
或者
~~~
go run main.go --ip 192.168.1.1
~~~
#### 以下 是 单线程 和 普通 多线程扫描
* 将下面的两个方法 粘贴到代码中 就可以 再 main 方法中 进行调用
~~~
/*
*
单线程扫描 速度比较慢
*/
func singleThreadedPortScan(ip string) {

	for i := 21; i <= 120; i++ {
		var address = fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println(address, "是关闭的")
			continue
		}
		conn.Close()
		fmt.Println(address, "打开")
	}

}

/*
*
多线程端口扫描
*/
func multiThreadedPortScan(ip string) {
	var begin = time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var address = fmt.Sprintf("%s:%d", ip, i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println(address, "打开")
		}(i)
	}
	wg.Wait()
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时", elapseTime)
}

~~~