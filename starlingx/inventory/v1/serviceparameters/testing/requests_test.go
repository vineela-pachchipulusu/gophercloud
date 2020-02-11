/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2020 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/serviceparameters"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListServiceParameters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceParameterListSuccessfully(t)

	pages := 0
	err := serviceparameters.List(client.ServiceClient(), serviceparameters.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := serviceparameters.ExtractServiceParameters(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 4 {
			t.Fatalf("Expected 4 serviceparameters, got %d", len(actual))
		}
		th.CheckDeepEquals(t, ServiceParameterHerp, actual[0])
		th.CheckDeepEquals(t, ServiceParameterDerp, actual[1])
		th.CheckDeepEquals(t, ServiceParameterMerp, actual[2])
		th.CheckDeepEquals(t, ServiceParameterBerp, actual[3])

		return true, nil
	})
	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllServiceParameters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceParameterListSuccessfully(t)

	allPages, err := serviceparameters.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := serviceparameters.ExtractServiceParameters(allPages)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, ServiceParameterHerp, actual[0])
	th.CheckDeepEquals(t, ServiceParameterDerp, actual[1])
	th.CheckDeepEquals(t, ServiceParameterMerp, actual[2])
	th.CheckDeepEquals(t, ServiceParameterBerp, actual[3])
}

func TestGetServiceParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	/* HandleServiceParameterGetSuccessfully returns the berpUUID structure */
	HandleServiceParameterGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := serviceparameters.Get(client, berpUUID).Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, ServiceParameterBerp, *actual)
}

func TestCreateServiceParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceParameterCreationSuccessfully(t, ServiceParameterCreateBody)

	serviceParamService := "bbq"
	serviceParamSection := "briquettes"
	resource := "bbq::briquettes::charcoal::mode"
	params := make(map[string]string)
	params["charcoal"] = "enabled"
	actual, err := serviceparameters.Create(client.ServiceClient(), serviceparameters.ServiceParameterOpts{
		Service:    &serviceParamService,
		Section:    &serviceParamSection,
		Parameters: &params,
		Resource:   &resource,
	}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, ServiceParameterBBQ, *actual)
}

/* Test Update mocks a patch for the BBQ service parameter */
func TestUpdateServiceParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceParameterUpdateSuccessfully(t)

	newValue := "disabled"

	actual, err := serviceparameters.Update(client.ServiceClient(),
		bbqUUID,
		serviceparameters.ServiceParameterPatchOpts{
			ParamValue: &newValue,
		}).Extract()
	if err != nil {
		t.Fatalf("Unexpected Update error: %v", err)
	}
	th.CheckDeepEquals(t, ServiceParameterBBQUpdated, *actual)
}

func TestDeleteServiceParameter(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceParameterDeletionSuccessfully(t)

	res := serviceparameters.Delete(client.ServiceClient(), "a5965fee-dc60-40dc-a234-edf87f1f9380")
	th.AssertNoErr(t, res.Err)
}
