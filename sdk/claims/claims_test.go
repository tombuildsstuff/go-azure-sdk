// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MPL-2.0 License. See NOTICE.txt in the project root for license information.

package claims_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/auth"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/claims"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/internal/test"
)

func TestAccParseClaims_azureCli(t *testing.T) {
	ctx := context.Background()

	test.AccTest(t)

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.AzureCliAuthorizerOptions{
		Api:      env.MicrosoftGraph,
		TenantId: test.TenantId,
	}
	authorizer, err := auth.NewAzureCliAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewAzureCliAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token(context.Background(), &http.Request{})
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatalf("token.AccessToken was empty")
	}

	claims, err := claims.ParseClaims(token)
	if err != nil {
		t.Fatal(err)
	}
	checkClaims(t, claims)
}

func TestAccParseClaims_clientCertificate(t *testing.T) {
	ctx := context.Background()
	test.AccTest(t)

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	pfx := test.Base64DecodeCertificate(t, test.ClientCertificate)

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: []string{},
		ClientId:     test.ClientId,
		Pkcs12Data:   pfx,
		Pkcs12Path:   test.ClientCertificatePath,
		Pkcs12Pass:   test.ClientCertPassword,
	}
	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token(context.Background(), &http.Request{})
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatal("token.AccessToken was empty")
	}

	claims, err := claims.ParseClaims(token)
	if err != nil {
		t.Fatal(err)
	}
	checkClaims(t, claims)
}

func TestAccParseClaims_clientSecret(t *testing.T) {
	ctx := context.Background()
	test.AccTest(t)

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ClientSecretAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: []string{},
		ClientId:     test.ClientId,
		ClientSecret: test.ClientSecret,
	}
	authorizer, err := auth.NewClientSecretAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientSecretAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token(context.Background(), &http.Request{})
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatalf("token.AccessToken was empty")
	}

	claims, err := claims.ParseClaims(token)
	if err != nil {
		t.Fatal(err)
	}
	checkClaims(t, claims)
}

func checkClaims(t *testing.T, claims *claims.Claims) {
	if claims == nil {
		t.Fatal("claims was nil")
	}
	if claims.AppId == "" {
		t.Fatal("claims.AppId was empty")
	}
	if claims.Audience == "" {
		t.Fatal("claims.Audience was empty")
	}
	if claims.Issuer == "" {
		t.Fatal("claims.Issuer was empty")
	}
	if len(claims.Roles) == 0 && claims.Scopes == "" {
		t.Fatal("claims.Roles and claims.Scopes were empty")
	}
	if claims.Subject == "" {
		t.Fatal("claims.Subject was empty")
	}
	if claims.TenantId == "" {
		t.Fatal("claims.TenantId was empty")
	}
}
