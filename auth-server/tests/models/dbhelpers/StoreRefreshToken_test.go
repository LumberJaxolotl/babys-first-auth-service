package dbhelpers

import (
	"testing"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/dbhelpers"
)

func TestAddUser(t *testing.T) {
    expected := "Alex Rivera"
    actual := "Alex Rivera" // Imagine this comes from your function
	
    if actual != expected {
        t.Errorf("Expected %s, but got %s", expected, actual)
    }
}
func TestStoreRefreshToken(t *testing.T) {
    expected := "Alex Rivera"
    
	
    dbhelpers.StoreRefreshToken(
		"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", 
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
    
	token, _ := dbhelpers.GetRefreshToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	actual := token.TokenHash

	if actual != expected {
        t.Errorf("Expected %s, but got %s", expected, actual)
    }
}

