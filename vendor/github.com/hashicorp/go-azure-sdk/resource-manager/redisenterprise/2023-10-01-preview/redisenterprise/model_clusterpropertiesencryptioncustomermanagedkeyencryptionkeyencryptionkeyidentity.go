package redisenterprise

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyEncryptionKeyIdentity struct {
	IdentityType                   *CmkIdentityType `json:"identityType,omitempty"`
	UserAssignedIdentityResourceId *string          `json:"userAssignedIdentityResourceId,omitempty"`
}
