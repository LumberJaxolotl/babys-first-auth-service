package dbhelpers_test

import (
	"testing"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/dbhelpers"
)

// TODO write one big test for  SetUserToEmailVerified,

// Tests StoreVerificationToken, GetEmailVerificationToken, DoTokensMatch
func TestVerificationTokenStoreAndGetFunctions(t *testing.T) {
    expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
    var actual any;
	
    dbhelpers.StoreVerificationToken(
		"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", 
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	)
    
	token, err := dbhelpers.GetEmailVerificationToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	if err != nil {
		actual = err
	}
	


	if actual != expected {
        t.Errorf("Expected %s, but got %s", expected, actual)
    }
}

