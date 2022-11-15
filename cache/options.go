package cache

import (
	"flag"
	_ "runtime"
)

type options struct {
	DataDir        string // data directory
	HttpAddress    string // http server address
	RaftTCPAddress string // construct Raft Address
	Bootstrap      bool   // start as master or not
	JoinAddress    string // peer address to join
}

func NewOptions() *options {
	opts := &options{}
	var httpAddress = flag.String("http", "127.0.0.1:6000", "Http address")
	var raftTCPAddress = flag.String("raft", "127.0.0.1:7000", "raft tcp address")
	var node = flag.String("node", "g://raft/node1", "raft node name")
	var bootstrap = flag.Bool("bootstrap", true, "start as raft cluster")
	var joinAddress = flag.String("join", "", "join address for raft cluster")
	flag.Parse()
	opts.DataDir = *node
	opts.HttpAddress = *httpAddress
	opts.Bootstrap = *bootstrap
	opts.RaftTCPAddress = *raftTCPAddress
	opts.JoinAddress = *joinAddress
	return opts
}
