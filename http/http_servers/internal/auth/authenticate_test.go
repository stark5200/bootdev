package auth

import (
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	secret := "testsecret"
    userID := uuid.New()
    expiresIn := time.Minute

    token, err := MakeJWT(userID, secret, expiresIn)
    if err != nil {
        t.Fatalf("MakeJWT failed: %v", err)
    }

    validatedID, err := ValidateJWT(token, secret)
    if err != nil {
        t.Fatalf("ValidateJWT failed: %v", err)
    }
    if validatedID != userID {
        t.Errorf("Expected userID %v, got %v", userID, validatedID)
    }
}

func TestExpiredJWT(t *testing.T) {
    secret := "testsecret"
    userID := uuid.New()
    expiresIn := -time.Minute // Already expired

    token, err := MakeJWT(userID, secret, expiresIn)
    if err != nil {
        t.Fatalf("MakeJWT failed: %v", err)
    }

    _, err = ValidateJWT(token, secret)
    if err == nil {
        t.Error("Expected error for expired JWT, got nil")
    }
}

func TestWrongSecretJWT(t *testing.T) {
    secret := "testsecret"
    wrongSecret := "wrongsecret"
    userID := uuid.New()
    expiresIn := time.Minute

    token, err := MakeJWT(userID, secret, expiresIn)
    if err != nil {
        t.Fatalf("MakeJWT failed: %v", err)
    }

    _, err = ValidateJWT(token, wrongSecret)
    if err == nil {
        t.Error("Expected error for JWT signed with wrong secret, got nil")
    }
}