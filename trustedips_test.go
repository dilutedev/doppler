package doppler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ip = "192.168.0.193/24"

func TestListTrustedIPs(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	ips, err := dp.ListTrustedIPs(projectName, config)
	assert.Nil(t, err)
	assert.NotNil(t, ips)

}

func TestAddIP(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	project, err := dp.AddIP(projectName, config, ip)
	assert.Nil(t, err)
	assert.NotNil(t, project)
}

func TestDeleteIP(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	res, err := dp.DeleteIP(projectName, config, ip)
	fmt.Println(res)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}
