package main

import (
	"fmt"
	"net/http"
	"strings"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	claims, _ := oauth.GetMyAccessTokenClaims(r.Header.Get("Authorization"))
	if len(claims.Subject) > 0 {
		fmt.Fprint(w, "Hello user client! ")
		secGrps, _ := oauth.GetUserInfo(r.Header.Get("Authorization"))
		fmt.Println(strings.Join(secGrps, " "))
		fmt.Fprintf(w, strings.Join(secGrps, " "))
	} else {
		fmt.Fprintf(w, "Hello app client!")
	}
}
