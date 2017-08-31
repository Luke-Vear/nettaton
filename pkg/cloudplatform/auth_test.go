package cloudplatform

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

var (
	testPassword123 = "$2a$10$YZtNOffXRIekdOzILPokJuaX1Yn5qIi2bEY1kbPWAcTvdHl77dqca"
)

func TestGenPwHash(t *testing.T) {
	password := "世界世界世界世界"
	pwHash, _ := genPwHash(password)
	if err := bcrypt.CompareHashAndPassword([]byte(pwHash), []byte(password)); err != nil {
		t.Errorf("invalid password hash")
	}
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

func TestGenerateTokenStringParseJWT(t *testing.T) {
	name := "thisWillBeTheSub"
	user := NewUser(name)

	validToken, _ := generateTokenString(user.(*User))
	actualSubClaim, err := parseJWT("Bearer: "+validToken, "sub")
	if actualSubClaim != name {
		t.Errorf("actualSubClaim: %v, expectedSubClaim: %v", actualSubClaim, name)
	}

	expiredToken := "Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDEyNTg2NjgsIm5iZiI6MTUwMDcxODY2OCwic3ViIjoibmljayJ9.lNSCEYI9Ij7XLlOB4yOq8Ezd1pQeMojmuqeOa4f3LwY"
	_, err = parseJWT(expiredToken, "sub")
	if err.Error() != "signature is invalid" {
		t.Errorf("actualErr: %v, expected: signature is invalid", err.Error())
	}

	_, err = generateTokenString(user.(*User))
	if err != nil {
		t.Errorf("err: %v", err)
	}
}

func TestIDFromToken(t *testing.T) {
	name := "thisWillBeTheSub"
	user := NewUser(name)
	validToken, _ := generateTokenString(user.(*User))
	actualSub, _ := IDFromToken("Bearer: " + validToken)
	if actualSub != name {
		t.Errorf("actualSub: %v, expectedSub: %v", actualSub, name)
	}
}
