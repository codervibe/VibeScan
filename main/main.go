package main

import (
	"fmt"
	"github.com/loveleshsharma/gohive"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	args := os.Args //获取用户输入的所有参数
	fmt.Println("len(args):", len(args))
	// 如果 a 包含 b
	for i := range args {
		if (len(args) == 1) || (len(args) > 1 && args[i] == "-h") || (len(args) > 1 && args[i] == "--help") {
			fmt.Println("Usage:   -i <ip> \n" +
				"\t -h --help \n" +
				"\t -t --thread \n")
			break
		}
		if (len(args) > 1 && args[i] == "-i") || (len(args) > 1 && args[i] == "--ip") {
			fmt.Println("Usage: gohive -i  <ip> \n" +
				"-h --help \n" +
				"-t --thread \n")
			//args的类型是[]string
			ip := args[2]
			fmt.Println("扫描的 ip:", ip)
			线程池端口扫描(ip, 7000)
			return
		}
		if (len(args) > 1 && args[i] == "-i") || (len(args) > 1 && args[i] == "--ip") {
			fmt.Println("Usage: gohive -i  <ip> \n" +
				"-h --help \n" +
				"-t --thread \n")
			//args的类型是[]string
			ip := args[2]
			fmt.Println("扫描的 ip:", ip)
			线程池端口扫描(ip, 7000)
			return
		}

	}

}

// ------------------------------------------
// 使用线程池 完成端口扫描
var wg sync.WaitGroup

// 地址管道 100容量
var addressChan = make(chan string, 100)

func worker() {
	defer wg.Done()
	for {
		address, ok := <-addressChan
		if !ok {
			break
		}
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Println(address, "开启")
	}
}

func 线程池端口扫描(ip string, thread int) {
	var begin = time.Now()
	var pool_size = 70000
	if thread != 0 {
		pool_size = thread
	}
	var pool = gohive.NewFixedSizePool(pool_size)
	go func() {
		for port := 1; port <= 65535; port++ {
			var address = fmt.Sprintf("%s:%d", ip, port)
			addressChan <- address
		}
		close(addressChan)
	}()
	for work := 0; work < pool_size; work++ {
		wg.Add(1)
		pool.Submit(worker)
	}
	wg.Wait()
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时", elapseTime)

}
func 参数处理() {

}
