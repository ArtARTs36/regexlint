package dot

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindString(t *testing.T) {
	cases := []struct {
		val         map[string]interface{}
		pointer     string
		expectedVal string
		expectedErr error
	}{
		{
			val: map[string]interface{}{
				"k": map[string]interface{}{
					"k1": "2",
				},
			},
			pointer:     "k.k1",
			expectedVal: "2",
		},
	}

	for _, tCase := range cases {
		gotVal, gotErr := FindString(tCase.val, tCase.pointer)
		if tCase.expectedErr != nil {
			assert.Equal(t, tCase.expectedErr, gotErr)
		} else {
			require.NoError(t, gotErr)

			assert.Equal(t, tCase.expectedVal, gotVal)
		}
	}
}
