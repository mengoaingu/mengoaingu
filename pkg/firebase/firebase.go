package firebase

import (
	"backend/pkg/token"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() *auth.Client {
	firebaseJson, err := mengoaingu_firebaseJsonBytes()
	if err != nil {
		panic("FirebaseAdminsdk load error")
	}
	opt := option.WithCredentialsJSON(firebaseJson)
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}
	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}
	return auth
}

func SetupJWT() token.Maker {
	tokenMaker, err := token.NewJWTMaker("12345678901234567890123456789012")
	if err != nil {
		log.Fatal("Falied to setup JWT maker")
	}
	return tokenMaker
}
