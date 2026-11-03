package fakedbhelpers

import (
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/dbhelpers"
	"github.com/LumberJaxolotl/babys-first-auth-service/models/lib"
	"github.com/golang-jwt/jwt/v5"
)

func init() {
	
	

}

// Seeds database with access, refresh, and email-verification tokens
func SeedTokens(){
	userIds := []string{
		"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		"b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12",
		"c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13",
	}
	
	// CAN BE USED WHEN ALL "STOREXXXTOKEN" FUNCTIONS HAVE BEEN IMPLIMENTED
	// for i, tokenType := range maps.Keys(lib.TOKEN_SIGNING_SECRETS){

	// changing 
	tokenTypesToBeSeeded := []lib.TokenType{"verification"}
	if slices.Contains(tokenTypesToBeSeeded, "verification") {
		user1Token, err := lib.GenerateAuthToken(userIds[1], "verification", time.Minute * 10 )
		if err != nil {
			log.Fatal(err)
		}
		dbhelpers.StoreVerificationToken(userIds[0], user1Token)
		
		user2Token, _ := lib.GenerateAuthToken(userIds[1], "verification", time.Minute * 10 )
		if err != nil {
			log.Fatal(err)
		}
		dbhelpers.StoreRefreshToken(userIds[1], user2Token)
		
		user3Token, _ := lib.GenerateAuthToken(userIds[2], "verification", time.Hour * 24 * 30 )
		dbhelpers.StoreRefreshToken(userIds[2], user3Token)
		if err != nil {
			log.Fatal(err)
		}
		
	} else if slices.Contains(tokenTypesToBeSeeded, "access") {
		// Todo impliment later
	} else if slices.Contains(tokenTypesToBeSeeded, "refresh") {
		// Todo impliment later
	}

	fmt.Println("DB has been Seeded with Auth Tokens")
}

func StoreRefreshToken(token *jwt.Token) error {
	return nil
}
