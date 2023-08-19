package doppler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var revoke_tokens = []RevokeTokenParam{
	{
		Token: "dp.st.dev.SwvJ4gUBRWxL3yehQxA3lCs3Q3Ks8cQefI6CBmlJyVT",
	},
	{
		Token: "dp.st.dev.2o0CsuhrjEcBcb4q2x4CGQTREV7OGWOfMySNDYlLUMy",
	},
}

func TestRevokeToken(t *testing.T) {
	dp, err := NewFromEnv()
	assert.Nil(t, err)
	assert.NotNil(t, dp)

	err = dp.RevokeTokens(revoke_tokens)
	assert.Nil(t, err)

}
