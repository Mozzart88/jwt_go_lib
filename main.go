package main

import (
	"fmt"
	"log"

	_ "crypto/sha256"

	"diliapi.com/token"
)

/*
sub: user_id string
user: user_name string
exp: timestamp_expired timestamp
role: user_role string
*/
func main() {
	var tok *token.Token
	var jwt string
	var err error
	var users = []map[string]any{
		{
			"sub":  "rfgy7u",
			"user": "tom",
			"role": "user",
		},
		{
			"sub":  "rf1y7u",
			"user": "bruce",
			"role": "admin",
		},
		{
			"sub":  "rfgY7u",
			"user": "jackie",
			"role": "root",
		},
	}

	for _, user := range users {
		tok = token.NewToken(nil, user, "12345678", 525600)
		if jwt, err = tok.GetJWT(); err != nil {
			log.Panicln(err.Error())
		}
		fmt.Println(jwt)
	}
}
