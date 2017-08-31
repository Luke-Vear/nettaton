package cloudplatform

import (
	"testing"
)

var (
	testPassword123 = "$2a$10$YZtNOffXRIekdOzILPokJuaX1Yn5qIi2bEY1kbPWAcTvdHl77dqca"
)

func init() {
	//password := "patsPassword123"
	//authUser.(User).
}

func TestGenPwHash(t *testing.T) {

}

func TestLogin(t *testing.T) {
	user := NewUser("namesAreNotImportant")
	user.(*User).HashedPassword = testPassword123

	_, err := login(user.(*User), "testPassword123")
	if err != nil {
		t.Errorf("err: %v, expected nil err", err)
	}

	expectErr := "crypto/bcrypt: hashedPassword is not the hash of the given password"
	_, actualErr := login(user.(*User), "")
	if actualErr.Error() != expectErr {
		t.Errorf("actualErr: %v, expected: %v", actualErr.Error(), expectErr)
	}
}

func TestGenerateTokenString(t *testing.T) {}

func TestParseJWT(t *testing.T) {
	expiredToken := "Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDEyNTg2NjgsIm5iZiI6MTUwMDcxODY2OCwic3ViIjoibmljayJ9.lNSCEYI9Ij7XLlOB4yOq8Ezd1pQeMojmuqeOa4f3LwY"
	_, err := parseJWT(expiredToken, "sub")
	if err.Error() != "signature is invalid" {
		t.Errorf("actualErr: %v, expected: signature is invalid", err.Error())
	}

}

func TestIDFromToken(t *testing.T) {

}
