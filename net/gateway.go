package net

import (
	"context"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/dedis/drand/protobuf/beacon"
	"github.com/dedis/drand/protobuf/dkg"
	"github.com/dedis/drand/protobuf/external"
)

// ADDRESSES and TLS
// https://github.com/denji/golang-tls
// How do we manage plain tcp and TLS connection with self-signed certificates
// or CA-signed certificates:
// (A) For non tls servers, when initiating connection just call with
// grpc.WithInsecure(). Same for listening ( see Golang gRPC API).
// (B) TLS communication using certificates
// How to differentiate (A) and (B) ?
// 	=> simple set of rules ? (xxx:443 | https | tls) == (B), rest is (A)
//
// For (B):
// Certificates is signed by a CA, so no options needed, simply
// 		crendentials.FromTLSCOnfig(&tls.Config{}) when connecting, or
// 		credentials.FromTLSConfig{&tls.Config{cert,private...}} for listening
// Certificates are given as a command line option "-cert xxx.crt" == (2) , otherwise (1)
// Since gRPC golang library does not allow us to access internal connections,
// every pair of communicating nodes is gonna have two active connections at the
// same time, one outgoing from each party.

// Peer is a simple interface that allows retrieving the address of a
// destination. It might further e enhanced with certificates properties and
// all.
type Peer interface {
	Address() string
	TLS() bool
}

// Gateway is the main interface to communicate to the external world. It
// acts as a listener to receive incoming requests and acts a client connecting
// to external particpants.
// The gateway fixes all functionalities offered by drand at a point in time.
type Gateway struct {
	Listener
	Client
}

// Service represents all the functionalities that a drand daemon must fullfill
// to participate
type Service interface {
	external.RandomnessServer
	beacon.BeaconServer
	dkg.DkgServer
}

// Client represents all methods that are callable to drand nodes
type Client interface {
	Public(p Peer, in *external.PublicRandRequest) (*external.PublicRandResponse, error)
	Setup(p Peer, in *dkg.DKGPacket) (*dkg.DKGResponse, error)
	NewBeacon(p Peer, in *beacon.BeaconPacket) (*beacon.BeaconResponse, error)
}

// grpcClient implements the Client functionalities using grpc connections
type grpcClient struct {
	sync.Mutex
	conns map[string]*grpc.ClientConn
}

func (g *grpcClient) Public(p Peer, in *external.PublicRandRequest) (*external.PublicRandResponse, error) {
	c, err := g.conn(p)
	if err != nil {
		return nil, err
	}
	client := external.NewRandomnessClient(c)
	return client.Public(context.Background(), in, nil)
}

func (g *grpcClient) Setup(p Peer, in *dkg.DKGPacket) (*dkg.DKGResponse, error) {
	c, err := g.conn(p)
	if err != nil {
		return nil, err
	}
	client := dkg.NewDkgClient(c)
	return client.Setup(context.Background(), in, nil)
}

func (g *grpcClient) NewBeacon(p Peer, in *beacon.BeaconPacket) (*beacon.BeaconResponse, error) {
	c, err := g.conn(p)
	if err != nil {
		return nil, err
	}
	client := beacon.NewBeaconClient(c)
	return client.NewBeacon(context.Background(), in, nil)
}

func (g *grpcClient) conn(p Peer) (*grpc.ClientConn, error) {
	g.Lock()
	defer g.Unlock()
	var err error
	c, ok := g.conns[p.Address()]
	if !ok {
		if !p.TLS() {
			c, err = grpc.Dial(p.Address(), grpc.WithInsecure())
			g.conns[p.Address()] = c
		} else {
			// TODO implement pool self signed certificates
		}
	}
	return c, err
}

type Listener interface {
	Start()
	Stop()
	RegisterDrandService(Service)
}

type grpcListener struct {
	service Service
	server  *grpc.Server
	lis     net.Listener
}

func NewGrpcListener(l net.Listener) Listener {
	return &grpcListener{
		server: grpc.NewServer(),
		lis:    l,
	}
}

func NewTCPGrpcListener(addr string) Listener {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("tcp listener: " + err.Error())
	}
	return NewGrpcListener(lis)
}

func (g *grpcListener) RegisterDrandService(s Service) {
	external.RegisterRandomnessServer(g.server, g.service)
	beacon.RegisterBeaconServer(g.server, g.service)
	dkg.RegisterDkgServer(g.server, g.service)
}

func (g *grpcListener) Start() {
	g.server.Serve(g.lis)
}

func (g *grpcListener) Stop() {
	g.server.GracefulStop()
}
