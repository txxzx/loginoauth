package loginoauth

/**
    @date: 2023/7/2
**/
import (
	"testing"

	"github.com/txxzx/loginoauth/facebook"
	"github.com/txxzx/loginoauth/google"
)

func TestQqAuth(t *testing.T) {
	success, err := QqAuth("", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestSinaAuth(t *testing.T) {
	success, err := SinaAuth("", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestWxAuth(t *testing.T) {
	success, err := WxAuth("", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestXiaomiAuth(t *testing.T) {
	success, err := XiaomiAuth("", "", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestGoogleVerify(t *testing.T) {
	result, err := google.GoogleVerify("", "")
	t.Logf("Success: %v, Err: %v", result, err)
}

func TestFacebookVerify(t *testing.T) {
	result, err := facebook.FacebookVerify("", "", "")
	t.Logf("Success: %v, Err: %v", result, err)
}
