package client

import (
	"net"

	"github.com/zollidan/tornet/bitfield"
	"github.com/zollidan/tornet/peers"
)

// A Client is a TCP connection with a peer
type Client struct {
	Conn     net.Conn
	Choked   bool
	Bitfield bitfield.Bitfield
	peer     peers.Peer
	infoHash [20]byte
	peerID   [20]byte
}