package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() *auth.Client {
	firebaseJson, err := simpleApp4f689FirebaseAdminsdk4esrk11f23bc586JsonBytes()
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
