package crypto

import (
	"testing"
)

func TestEncryptMd5(t *testing.T) {
	resultMap := map[string]string{
		"admin": "21232f297a57a5a743894a0e4a801fc3",
		"test":  "098f6bcd4621d373cade4e832627b4f6",
	}

	for k, v := range resultMap {
		ret, err := Md5Encrypt(k)
		if len(ret) == 0 || err != nil {
			t.Error("Failed to encrypt Md5")
		}

		if ret != v {
			t.Errorf("Encrypt Md5 error: %s = %s， not %s", k, ret, v)
		}
	}
}
