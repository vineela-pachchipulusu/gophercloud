/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2024 Wind River Systems, Inc. */

package testing

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/starlingx/inventory/v1/networkAddressPools"
	"github.com/gophercloud/gophercloud/testhelper/client"

	th "github.com/gophercloud/gophercloud/testhelper"
	"testing"
)

func TestListNetworkAddressPools(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkAddressPoolListSuccessfully(t)

	pages := 0
	err := networkAddressPools.List(client.ServiceClient(), networkAddressPools.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++
		actual, err := networkAddressPools.ExtractNetworkAddressPools(page)
		if err != nil {
			return false, err
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 network address pools, got %d", len(actual))
		}
		th.CheckDeepEquals(t, NetworkAddressPoolHerp, actual[0])
		th.CheckDeepEquals(t, NetworkAddressPoolDerp, actual[1])

		return true, nil
	})

	th.AssertNoErr(t, err)

	if pages != 1 {
		t.Errorf("Expected 1 page, saw %d", pages)
	}
}

func TestListAllNetworkAddressPools(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkAddressPoolListSuccessfully(t)

	network_address_pools, err := networkAddressPools.ListNetworkAddressPools(client.ServiceClient())
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, NetworkAddressPoolHerp, network_address_pools[0])
	th.CheckDeepEquals(t, NetworkAddressPoolDerp, network_address_pools[1])
}

func TestGetNetworkAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkAddressPoolGetSuccessfully(t)

	client := client.ServiceClient()
	actual, err := networkAddressPools.Get(client, "11111111-a6e5-425e-9317-995da88d6694").Extract()
	if err != nil {
		t.Fatalf("Unexpected Get error: %v", err)
	}

	th.CheckDeepEquals(t, NetworkAddressPoolHerp, *actual)
}

func TestCreateNetworkAddressPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkAddressPoolCreationSuccessfully(t, SingleNetworkAddressPoolBody)

	network_name := "oam"
	addresspool_name := "oam-ipv4"
	network_uuid := "11111111-0000-425e-9317-995da88d6694"
	addresspool_uuid := "11111111-1111-425e-9317-995da88d6694"

	actual, err := networkAddressPools.Create(client.ServiceClient(), networkAddressPools.NetworkAddressPoolOpts{
		NetworkUUID:     &network_uuid,
		AddressPoolUUID: &addresspool_uuid,
		NetworkName:     &network_name,
		AddressPoolName: &addresspool_name,
	}).Extract()

	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, NetworkAddressPoolHerp, *actual)
}

func TestDeleteNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNetworkAddressPoolDeletionSuccessfully(t)

	res := networkAddressPools.Delete(client.ServiceClient(), "11111111-a6e5-425e-9317-995da88d6694")
	th.AssertNoErr(t, res.Err)
}
