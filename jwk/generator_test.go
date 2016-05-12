package jwk

import (
	"github.com/ory-am/hydra/pkg"
	"github.com/square/go-jose"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator(t *testing.T) {
	for _, c := range []struct {
		g     KeyGenerator
		check func(*jose.JsonWebKeySet)
	}{
		{
			g: &RS256Generator{},
			check: func(ks *jose.JsonWebKeySet) {
				assert.Len(t, ks, 2)
			},
		},
	} {
		keys, err := c.g.Generate("foo")
		pkg.AssertError(t, false, err)
		if err != nil {
			c.check(keys)
		}
	}
}
