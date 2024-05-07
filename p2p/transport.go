package p2p

// we do tcp later
//Peer is an interface that represent theremote node
type Peer interface {
}

//Tansport is anything that handles teh communication that between the nodes in the netork
//This can be of the form TCP UDP web sockets
type Transport interface {
}
