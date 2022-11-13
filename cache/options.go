package cache

type options struct {
	dataDir        string // data directory
	httpAddress    string // http server address
	raftTCPAddress string // construct Raft Address
	bootstrap      bool   // start as master or not
	joinAddress    string // peer address to join
}
