package documentdb

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/go-autorest/autorest"
)

// DatabaseAccountsClient is a minimal interface for azure DatabaseAccountsClient
type DatabaseAccountsClient interface {
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result documentdb.DatabaseAccountsListResult, err error)
	ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListKeysResult, err error)
}

type databaseAccountsClient struct {
	documentdb.DatabaseAccountsClient
}

var _ DatabaseAccountsClient = &databaseAccountsClient{}

// NewDatabaseAccountsClient creates a new DatabaseAccountsClient
func NewDatabaseAccountsClient(subscriptionID string, authorizer autorest.Authorizer) DatabaseAccountsClient {
	client := documentdb.NewDatabaseAccountsClient(subscriptionID)
	client.Authorizer = authorizer

	return &databaseAccountsClient{
		DatabaseAccountsClient: client,
	}
}
