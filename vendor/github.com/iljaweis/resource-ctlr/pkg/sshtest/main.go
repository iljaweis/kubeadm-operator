package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

func main() {
	key, err := getPrivateKey("/Users/iw/go/src/github.com/iljaweis/resource-ctlr/test/.vagrant/machines/controller02/virtualbox/private_key")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", key)

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.10.2:22", config)

	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	//
	//
	//
	//sshConfig := &ssh.ClientConfig{
	//	User: "your_user_name",
	//	Auth: []ssh.AuthMethod{
	//		PublicKeyFile("/path/to/your/pub/certificate/key")
	//	},
	//}
}

func getPrivateKey(file string) (ssh.Signer, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read file %s")
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, errors.Wrapf(err, "could not parse ssh key from file %s")
	}
	return key, nil
}

/*
var hostKey ssh.PublicKey
// A public key may be used to authenticate against the remote
// server by using an unencrypted PEM-encoded private key file.
//
// If you have an encrypted private key, the crypto/x509 package
// can be used to decrypt it.
key, err := ioutil.ReadFile("/home/user/.ssh/id_rsa")
if err != nil {
    log.Fatalf("unable to read private key: %v", err)
}

// Create the Signer for this private key.
signer, err := ssh.ParsePrivateKey(key)
if err != nil {
    log.Fatalf("unable to parse private key: %v", err)
}

config := &ssh.ClientConfig{
    User: "user",
    Auth: []ssh.AuthMethod{
        // Use the PublicKeys method for remote authentication.
        ssh.PublicKeys(signer),
    },
    HostKeyCallback: ssh.FixedHostKey(hostKey),
}

// Connect to the remote server and perform the SSH handshake.
client, err := ssh.Dial("tcp", "host.com:22", config)
if err != nil {
    log.Fatalf("unable to connect: %v", err)
}
defer client.Close()
*/
