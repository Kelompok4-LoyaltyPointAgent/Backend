package config

import (
	"os"
	"strings"
)

type serviceAccountKey struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

type FirebaseConfig struct {
	ServiceAccountKey serviceAccountKey
	Bucket            string
}

func LoadFirebaseConfig() FirebaseConfig {
	return FirebaseConfig{
		ServiceAccountKey: serviceAccountKey{
			Type:                    os.Getenv("TYPE"),
			ProjectID:               os.Getenv("PROJECT_ID"),
			PrivateKeyID:            os.Getenv("PRIVATE_KEY_ID"),
			PrivateKey:              strings.Replace(os.Getenv("PRIVATE_KEY"), `\n`, "\n", -1),
			ClientEmail:             os.Getenv("CLIENT_EMAIL"),
			ClientID:                os.Getenv("CLIENT_ID"),
			AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
			TokenURI:                "https://oauth2.googleapis.com/token",
			AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
			ClientX509CertURL:       "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-1irbg%40capstone-project-eede7.iam.gserviceaccount.com",
		},
		Bucket: os.Getenv("BUCKET"),
	}
}
