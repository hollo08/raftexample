package cache

import "flag"

type options struct {
	dataDir        string // data directory
	HttpAddress    string // http server address
	raftTCPAddress string // construct Raft Address
	bootstrap      bool   // start as master or not
	JoinAddress    string // peer address to join
}

func NewOptions() *options {
	opts := &options{}

	var httpAddress = flag.String("http", "127.0.0.1:6000", "Http address")
	var raftTCPAddress = flag.String("raft", "127.0.0.1:7000", "raft tcp address")
	var node = flag.String("node", "node1", "raft node name")
	var bootstrap = flag.Bool("bootstrap", true, "start as raft cluster")
	var joinAddress = flag.String("join", "", "join address for raft cluster")
	flag.Parse()
	opts.dataDir = "./" + *node
	opts.HttpAddress = *httpAddress
	opts.bootstrap = *bootstrap
	opts.raftTCPAddress = *raftTCPAddress
	opts.JoinAddress = *joinAddress
	return opts
}
