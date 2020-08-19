package engine

import (
	"fmt"
	"time"

	"github.com/PaulSnow/factom2d/common/globals"
)

func printSummary(summary *int, value int, listenTo *int, wsapiNode *int) {

	if *listenTo < 0 || *listenTo >= len(fnodes) {
		return
	}
	for {
		PrintOneStatus(*listenTo, *wsapiNode)
		time.Sleep(2 * time.Second)
	}
}

func GetSystemStatus(listenTo int, wsapiNode int) string {
	fnodes := GetFnodes()
	prt := "===SummaryStart===\n\n"
	prt = prt + fmt.Sprintf("%sTime: %d %s Elapsed time:%s\n", prt, time.Now().Unix(), time.Now().Format("2006-01-02 15:04:05"), time.Since(globals.StartTime).String())

	for _, f := range fnodes {
		prt = prt + fmt.Sprintf("%s\n", f.State.A_Instance.Status())
	}
	prt = prt + "===SummaryEnd===\n"
	return prt
}

var out string // previous status

func PrintOneStatus(listenTo int, wsapiNode int) {
	prt := GetSystemStatus(listenTo, wsapiNode)
	if prt != out {
		fmt.Println(prt)
		out = prt
	}

}

func SystemFaults(f *FactomNode) string {
	dbheight := f.State.LLeaderHeight
	pl := f.State.ProcessLists.Get(dbheight)
	if pl == nil {
		return ""
	}
	if len(pl.System.List) == 0 {
		str := fmt.Sprintf("%5s %13s %6s Length: 0\n", "", "System List", f.State.FactomNodeName)
		return str
	}
	str := fmt.Sprintf("%5s %s\n", "", "System List")
	for _, ff := range pl.System.List {
		if ff != nil {
			str = fmt.Sprintf("%s%8s%s\n", str, "", ff.String())
		}
	}
	str = str + "\n"
	return str
}

func FaultSummary() string {

	return ""
}
