//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//|------------------------------------------------------------------
//| Cellgo Framework Boot file
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-8-8

package cellgo

import (
	ctcpip "github.com/mrkt/cellgo/tcpip"
	"github.com/mrkt/tcpip"
)

type Boot struct {
}

//ON&GC Event
func (b *Boot) GCEvent() {
	for _, v := range Events {
		go v.EventON()
		go v.EventGC()
	}
}

//RunTcp Tcp
func (b *Boot) RunSocketIO() {
	a_socketIO := new(ctcpip.SocketIO)
	b_socketIO := new(tcpip.SocketIO)
	a_socketIO.SetB(b_socketIO)
	a_socketIO.RunSocketIO()
}
