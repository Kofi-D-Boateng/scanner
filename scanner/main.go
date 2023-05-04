package scanner

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Kofi-D-Boateng/scanner/models"
)

func StartPortScan(items *[]models.OpenPort,ips *string,ports *string){
	ipList := strings.Split(*ips, ",")
	var waitgroup sync.WaitGroup
	waitgroup.Add(len(ipList))
	for _,ip := range ipList{
		go scanPorts(ip,*ports,items,&waitgroup)
	}
	waitgroup.Wait()
}

// func StartVulnerabilityScan()

func scanPorts(ip string,ports string,items *[]models.OpenPort, wg *sync.WaitGroup){
	defer wg.Done()
	portRange := strings.Split(ports,":")
	start,_ := strconv.Atoi(portRange[0])
	end,_ := strconv.Atoi(portRange[1])
	for port := start; port < end;port++{
		addr := strings.Join([]string{ip,strconv.Itoa(port)},":")
		conn,err := net.Dial("tcp",addr)
		if err == nil{
			var openPort models.OpenPort
			openPort.Name = ip
			openPort.Number = port
			openPort.Message = "OPEN"
			if port == 22{
				openPort.Mitigation = "Close Telnet or Set Rule to allow same subnet mask/specified IP to ssh into system"
				openPort.Level = 3
			}
			openPort.Mitigation = "Close if not necessary to be open"
			openPort.Level = 1
			openPort.DateScanned = time.Now().Local().String()
			openPort.DateFound = time.Now().Local().String()
			conn.Close()
			*items = append(*items, openPort)
		}else{
			fmt.Printf("[%v]: %v\n",time.Now().Local().String(),err)
		}
	}
}