package tls

import "net"

const (
   TLS_RSA_WITH_3DES_EDE_CBC_SHA           uint16 = 0x000a
   TLS_RSA_WITH_RC4_128_SHA                uint16 = 0x0005
)

var (
   FakeFFDHE2048 = uint16(0x0100)
   FakeFFDHE3072 = uint16(0x0101)
)

func Listen(network, laddr string, config *Config) (net.Listener, error) {
   return nil, nil
}

type Certificate int

func LoadX509KeyPair(certFile, keyFile string) (Certificate, error) {
   return 0, nil
}

type ClientHelloID struct {
   Client string
   Seed *PRNGSeed
}

type ClientHelloSpec int

type ClientSessionState int

type Config int

func (c *Config) Clone() *Config {
   return nil
}

type PRNGSeed int

type UConn int

func (uconn *UConn) ApplyPreset(p *ClientHelloSpec) error {
   return nil
}

type ClientSessionCache interface {
   Get(sessionKey string) (session *ClientSessionState, ok bool)
   Put(sessionKey string, cs *ClientSessionState)
}
