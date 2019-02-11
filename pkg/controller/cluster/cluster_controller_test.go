package cluster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInit(t *testing.T) {
	a := assert.New(t)

	output := `You can now join any number of machines by running the following on each node
as root:

kubeadm join 192.168.0.200:6443 --token j04n3m.octy8zely83cy2ts --discovery-token-ca-cert-hash    sha256:84938d2a22203a8e56a787ec0c6ddad7bc7dbd52ebabc62fd5f4dbea72b14d1f

`

	token, hash := ParseKubeadmOutput(output)
	a.Equal("j04n3m.octy8zely83cy2ts", token)
	a.Equal("sha256:84938d2a22203a8e56a787ec0c6ddad7bc7dbd52ebabc62fd5f4dbea72b14d1f", hash)
}
