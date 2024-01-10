// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/auth"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/claims"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/internal/test"
)

func TestCachedAuthorizer(t *testing.T) {
	tokenPattern := regexp.MustCompile("^[a-zA-Z0-9_-]+[.][a-zA-Z0-9_-]+[.][a-zA-Z0-9_-]+")
	req := &http.Request{}

	authorizer, err := auth.NewCachedAuthorizer(&test.TestAuthorizer{})
	if err != nil {
		t.Fatalf("received error for NewCachedAuthorizer(): %+v", err)
	}

	// Retrieve the first access tokens
	token, err := authorizer.Token(context.Background(), req)
	if err != nil {
		t.Fatalf("received error for CachedAuthorizer.Token(): %+v", err)
	}
	if !tokenPattern.MatchString(token.AccessToken) {
		t.Fatalf("unexpected access token received: %q", token.AccessToken)
	}
	auxTokens, err := authorizer.AuxiliaryTokens(context.Background(), req)
	if err != nil {
		t.Fatalf("received error for CachedAuthorizer.AuxiliaryTokens(): %+v", err)
	}
	for i, auxToken := range auxTokens {
		if !tokenPattern.MatchString(auxToken.AccessToken) {
			t.Fatalf("unexpected auxiliary access token received at %d: %q", i, token.AccessToken)
		}
	}

	// Parse the claims and compare the IssuedAt and Expiry times
	tokenClaims, err := claims.ParseClaims(token)
	if err != nil {
		t.Fatalf("received error for claims.ParseClaims(): %+v", err)
	}
	if tokenClaims.IssuedAt != test.TestTokenIssued.Unix() {
		t.Fatalf("unexpected `iat` claim for access token, expected: %d, received: %d", test.TestTokenIssued.Unix(), tokenClaims.IssuedAt)
	}
	if tokenClaims.Expires != test.TestTokenExpiry.Unix() {
		t.Fatalf("unexpected `exp` claim for access token, expected: %d, received: %d", test.TestTokenExpiry.Unix(), tokenClaims.Expires)
	}
	for i, auxToken := range auxTokens {
		auxTokenClaims, err := claims.ParseClaims(auxToken)
		if err != nil {
			t.Fatalf("received error for claims.ParseClaims(): %+v", err)
		}
		if auxTokenClaims.IssuedAt != test.TestTokenIssued.Unix() {
			t.Fatalf("unexpected `iat` claim for auxiliary access token at %d, expected: %d, received: %d", i, test.TestTokenIssued.Unix(), auxTokenClaims.IssuedAt)
		}
		if auxTokenClaims.Expires != test.TestTokenExpiry.Unix() {
			t.Fatalf("unexpected `exp` claim for auxiliary access token at %d, expected: %d, received: %d", i, test.TestTokenExpiry.Unix(), auxTokenClaims.Expires)
		}
	}

	// Wait for 5 seconds and advance the issued/expiry times for the testAuthorizer
	time.Sleep(5 * time.Second)
	earlierTestTokenIssued := test.TestTokenIssued
	earlierTestTokenExpiry := test.TestTokenExpiry
	test.TestTokenIssued = time.Now()
	test.TestTokenExpiry = time.Now().Add(3599 * time.Second)

	// Retrieve a second token, this should be retrieved from the cache
	token, err = authorizer.Token(context.Background(), req)
	if err != nil {
		t.Fatalf("received error for CachedAuthorizer.Token(): %+v", err)
	}
	if !tokenPattern.MatchString(token.AccessToken) {
		t.Fatalf("unexpected access token received: %q", token.AccessToken)
	}
	auxTokens, err = authorizer.AuxiliaryTokens(context.Background(), req)
	if err != nil {
		t.Fatalf("received error for CachedAuthorizer.AuxiliaryTokens(): %+v", err)
	}
	for i, auxToken := range auxTokens {
		if !tokenPattern.MatchString(auxToken.AccessToken) {
			t.Fatalf("unexpected auxiliary access token received at %d: %q", i, token.AccessToken)
		}
	}

	// Parse the claims for the second token, ensure the IssuedAt and Expiry times _have not_ changed
	tokenClaims, err = claims.ParseClaims(token)
	if err != nil {
		t.Fatalf("received error for claims.ParseClaims(): %+v", err)
	}
	if tokenClaims.IssuedAt != earlierTestTokenIssued.Unix() {
		t.Fatalf("unexpected `iat` claim for access token, expected: %d, received: %d", earlierTestTokenIssued.Unix(), tokenClaims.IssuedAt)
	}
	if tokenClaims.Expires != earlierTestTokenExpiry.Unix() {
		t.Fatalf("unexpected `exp` claim for access token, expected: %d, received: %d", earlierTestTokenExpiry.Unix(), tokenClaims.Expires)
	}
	for i, auxToken := range auxTokens {
		auxTokenClaims, err := claims.ParseClaims(auxToken)
		if err != nil {
			t.Fatalf("received error for claims.ParseClaims(): %+v", err)
		}
		if auxTokenClaims.IssuedAt != earlierTestTokenIssued.Unix() {
			t.Fatalf("unexpected `iat` claim for auxiliary access token at %d, expected: %d, received: %d", i, earlierTestTokenIssued.Unix(), auxTokenClaims.IssuedAt)
		}
		if auxTokenClaims.Expires != earlierTestTokenExpiry.Unix() {
			t.Fatalf("unexpected `exp` claim for auxiliary access token at %d, expected: %d, received: %d", i, earlierTestTokenExpiry.Unix(), auxTokenClaims.Expires)
		}
	}

	// Invalidate the access tokens
	if err = authorizer.InvalidateCachedTokens(); err != nil {
		t.Fatalf("received error for CachedAuthorizer.ExpireTokens(): %+v", err)
	}

	// Retrieve a third token, which should be re-acquired from the testAuthorizer
	token, err = authorizer.Token(context.Background(), req)
	if err != nil {
		t.Fatalf("received error for CachedAuthorizer.Token(): %+v", err)
	}
	if !tokenPattern.MatchString(token.AccessToken) {
		t.Fatalf("unexpected access token received: %q", token.AccessToken)
	}
	auxTokens, err = authorizer.AuxiliaryTokens(context.Background(), req)
	if err != nil {
		t.Fatalf("received error for CachedAuthorizer.AuxiliaryTokens(): %+v", err)
	}
	for i, auxToken := range auxTokens {
		if !tokenPattern.MatchString(auxToken.AccessToken) {
			t.Fatalf("unexpected auxiliary access token received at %d: %q", i, token.AccessToken)
		}
	}

	// Parse the claims for the third token, ensure the IssuedAt and Expiry times _have_ changed
	tokenClaims, err = claims.ParseClaims(token)
	if err != nil {
		t.Fatalf("received error for claims.ParseClaims(): %+v", err)
	}
	if tokenClaims.IssuedAt != test.TestTokenIssued.Unix() {
		t.Fatalf("unexpected `iat` claim for access token, expected: %d, received: %d", test.TestTokenIssued.Unix(), tokenClaims.IssuedAt)
	}
	if tokenClaims.Expires != test.TestTokenExpiry.Unix() {
		t.Fatalf("unexpected `exp` claim for access token, expected: %d, received: %d", test.TestTokenExpiry.Unix(), tokenClaims.Expires)
	}
	for i, auxToken := range auxTokens {
		auxTokenClaims, err := claims.ParseClaims(auxToken)
		if err != nil {
			t.Fatalf("received error for claims.ParseClaims(): %+v", err)
		}
		if auxTokenClaims.IssuedAt != test.TestTokenIssued.Unix() {
			t.Fatalf("unexpected `iat` claim for auxiliary access token at %d, expected: %d, received: %d", i, test.TestTokenIssued.Unix(), auxTokenClaims.IssuedAt)
		}
		if auxTokenClaims.Expires != test.TestTokenExpiry.Unix() {
			t.Fatalf("unexpected `exp` claim for auxiliary access token at %d, expected: %d, received: %d", i, test.TestTokenExpiry.Unix(), auxTokenClaims.Expires)
		}
	}
}
