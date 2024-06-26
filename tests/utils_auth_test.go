package tests

import (
	"context"
	"testing"

	"server.go/configs"
	"server.go/constants"
	auth_utils "server.go/utils/auth"
)

func TestGenearteJwtToken(t *testing.T) {
	configs.LoadEnv()
	ctx := context.Background()

	token, err := auth_utils.GenearteJwtToken(ctx, constants.EMAIL)
	if err != nil {
		t.Errorf("Failed to generate JWT token: %v", err)
	}
	if token == "" {
		t.Error("Generated JWT token is empty")
	}
}

func TestValidateJwtToken(t *testing.T) {
	configs.LoadEnv()
	ctx := context.Background()

	tokenString, err := auth_utils.GenearteJwtToken(ctx, constants.EMAIL)
	if err != nil {
		t.Errorf(err.Error())
	}

	userEmail, err := auth_utils.ValidateJwtToken(ctx, tokenString)

	if err != nil {
		t.Errorf(err.Error())
	}
	if userEmail == "" {
		t.Error("Failed to validate JWT token")
	}
	if userEmail != constants.EMAIL {
		t.Errorf("Expected email is %s", constants.EMAIL)
	}
}

func TestHashPassword(t *testing.T) {
	hashedPassword := auth_utils.HashPassword(constants.PASSWORD)

	compare, _ := auth_utils.VerifyPassword(constants.PASSWORD, hashedPassword)
	if !compare {
		t.Error("Failed to verify hashed password")
	}
}

func TestVerifyPassword(t *testing.T) {
	hashedPassword := auth_utils.HashPassword(constants.PASSWORD)

	compare, _ := auth_utils.VerifyPassword(constants.PASSWORD, hashedPassword)
	if !compare {
		t.Error("Failed to verify hashed password")
	}

	invalidPassword := "wrongpassword"

	compare, _ = auth_utils.VerifyPassword(invalidPassword, hashedPassword)
	if compare {
		t.Error("Incorrectly verified invalid password")
	}
}
