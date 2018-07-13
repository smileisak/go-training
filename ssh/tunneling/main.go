package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// Endpoint is a struct with Host and Port
type Endpoint struct {
	Host string
	Port int
}

// Strings is a Endpoint method to prints host and port.
func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

// SSHtunnel is a struct for SSH tunnel configuration
type SSHtunnel struct {
	Local  *Endpoint
	Server *Endpoint
	Remote *Endpoint
	Config *ssh.ClientConfig
}

// Start Method to start a local server and forward connection to the remote one.
func (tunnel *SSHtunnel) Start() error {
	listener, err := net.Listen("tcp", tunnel.Local.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go tunnel.forward(conn)
	}
}

func (tunnel *SSHtunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.String(), tunnel.Config)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.String())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}

// SSHAgent for ssh authentication
func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

func main() {
	// Create a local connexion Endpoint
	localEndpoint := &Endpoint{
		Host: "localhost",
		Port: 9000,
	}
	// Create a server connexion Endpoint to pass through
	serverEndpoint := &Endpoint{
		Host: "bastion.com",
		Port: 22,
	}
	// Create a remote connexion Endpoint
	remoteEndpoint := &Endpoint{
		Host: "remot.server.com",
		Port: 8443,
	}
	// Configure ssh
	sshConfig := &ssh.ClientConfig{
		User: "user",
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Auth: []ssh.AuthMethod{
			SSHAgent(),
		},
	}

	tunnel := &SSHtunnel{
		Config: sshConfig,
		Local:  localEndpoint,
		Server: serverEndpoint,
		Remote: remoteEndpoint,
	}
	log.Printf("[*] Opening SSH Tunnel: %s:%d --> %s:%d --> %s:%d\n",
		localEndpoint.Host, localEndpoint.Port,
		serverEndpoint.Host, serverEndpoint.Port,
		remoteEndpoint.Host, remoteEndpoint.Port)

	tunnel.Start()
}
