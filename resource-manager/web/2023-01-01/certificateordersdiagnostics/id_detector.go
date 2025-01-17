package certificateordersdiagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DetectorId{}

// DetectorId is a struct representing the Resource ID for a Detector
type DetectorId struct {
	SubscriptionId       string
	ResourceGroupName    string
	CertificateOrderName string
	DetectorName         string
}

// NewDetectorID returns a new DetectorId struct
func NewDetectorID(subscriptionId string, resourceGroupName string, certificateOrderName string, detectorName string) DetectorId {
	return DetectorId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		CertificateOrderName: certificateOrderName,
		DetectorName:         detectorName,
	}
}

// ParseDetectorID parses 'input' into a DetectorId
func ParseDetectorID(input string) (*DetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(DetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DetectorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDetectorIDInsensitively parses 'input' case-insensitively into a DetectorId
// note: this method should only be used for API response data and not user input
func ParseDetectorIDInsensitively(input string) (*DetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(DetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DetectorId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DetectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.CertificateOrderName, ok = input.Parsed["certificateOrderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateOrderName", input)
	}

	if id.DetectorName, ok = input.Parsed["detectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectorName", input)
	}

	return nil
}

// ValidateDetectorID checks that 'input' can be parsed as a Detector ID
func ValidateDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Detector ID
func (id DetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CertificateRegistration/certificateOrders/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CertificateOrderName, id.DetectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Detector ID
func (id DetectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCertificateRegistration", "Microsoft.CertificateRegistration", "Microsoft.CertificateRegistration"),
		resourceids.StaticSegment("staticCertificateOrders", "certificateOrders", "certificateOrders"),
		resourceids.UserSpecifiedSegment("certificateOrderName", "certificateOrderValue"),
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorName", "detectorValue"),
	}
}

// String returns a human-readable description of this Detector ID
func (id DetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Certificate Order Name: %q", id.CertificateOrderName),
		fmt.Sprintf("Detector Name: %q", id.DetectorName),
	}
	return fmt.Sprintf("Detector (%s)", strings.Join(components, "\n"))
}
