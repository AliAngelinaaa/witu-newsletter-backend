package auth

import (
    "context"
    "log"
    "os"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/idtoken"
)

var (
    oauth2Config *oauth2.Config
    googleClientID string
	googleClientSecret string
)

func init() {
    googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
    googleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
    if googleClientID == "" || googleClientSecret == "" {
        log.Fatal("GOOGLE_CLIENT_ID or GOOGLE_CLIENT_SECRET environment variable not set")
    }

    oauth2Config = &oauth2.Config{
        ClientID:     googleClientID,
        ClientSecret: googleClientSecret,
        RedirectURL:  "http://localhost:8080/auth/callback", // Your callback URL
        Scopes:       []string{"openid", "profile", "email"},
        Endpoint:     google.Endpoint,
    }
}

// VerifyToken validates the Google ID token and returns the payload.
func VerifyToken(idToken string) (*idtoken.Payload, error) {
    ctx := context.Background()
    payload, err := idtoken.Validate(ctx, idToken, googleClientID)
    if err != nil {
        return nil, err
    }
    return payload, nil
}
