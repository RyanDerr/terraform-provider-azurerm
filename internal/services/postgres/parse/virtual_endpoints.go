// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type VirtualEndpointsId struct {
	SubscriptionId      string
	ResourceGroup       string
	FlexibleServerName  string
	VirtualendpointName string
}

func NewVirtualEndpointsID(subscriptionId, resourceGroup, flexibleServerName, virtualendpointName string) VirtualEndpointsId {
	return VirtualEndpointsId{
		SubscriptionId:      subscriptionId,
		ResourceGroup:       resourceGroup,
		FlexibleServerName:  flexibleServerName,
		VirtualendpointName: virtualendpointName,
	}
}

func (id VirtualEndpointsId) String() string {
	segments := []string{
		fmt.Sprintf("Virtualendpoint Name %q", id.VirtualendpointName),
		fmt.Sprintf("Flexible Server Name %q", id.FlexibleServerName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Virtual Endpoints", segmentsStr)
}

func (id VirtualEndpointsId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/flexibleServers/%s/virtualendpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.FlexibleServerName, id.VirtualendpointName)
}

// VirtualEndpointsID parses a VirtualEndpoints ID into an VirtualEndpointsId struct
func VirtualEndpointsID(input string) (*VirtualEndpointsId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, fmt.Errorf("parsing %q as an VirtualEndpoints ID: %+v", input, err)
	}

	resourceId := VirtualEndpointsId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.FlexibleServerName, err = id.PopSegment("flexibleServers"); err != nil {
		return nil, err
	}
	if resourceId.VirtualendpointName, err = id.PopSegment("virtualendpoints"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
