package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var (
	maxCPU   int
)

func init(){
	mcpu := os.Getenv("MAX_CPU_UTIL")
	parsedMcpu, err2 := strconv.ParseInt(mcpu,10,16)

	if  err2 != nil{
		maxCPU   = 30
	}else{
		maxCPU	 = int(parsedMcpu)
	}

}

func main() {
	var ports string
	var ips string
	flag.StringVar(&ports,"p","","List of ports to scan")
	flag.StringVar(&ips,"ip","","List of ips to scan")
	flag.Parse()
	os_name := runtime.GOOS
	runtime.SetCPUProfileRate(maxCPU)



	if len(ips) <= 0 {
		fmt.Println("[ERROR]: No Ips were found to scan....")
		os.Exit(1)
	}else{
		fmt.Printf("[OS]: %v\n",os_name)
		fmt.Printf("[ips]: %v\n",ips)
		fmt.Printf("[ports]: %v\n",ports)
	}
}