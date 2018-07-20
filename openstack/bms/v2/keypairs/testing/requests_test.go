package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/bms/v2/keypairs"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/os-keypairs", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListOutput)
	})

	actual, err := keypairs.List(client.ServiceClient(), keypairs.ListOpts{})
	th.AssertNoErr(t, err)
	expected := []keypairs.KeyPair{
		{

			Name:        "c2c-firstkey",
			PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyUHvhG56/beGw/23/zYGQJ9YUmsMiRhigdbuhpS7QgPMG1LMlW32Homqph166GXpfRvo66vwO+vTrQu9xLR0Z4oYntKDGtc9pF5SRE7nsSjxmtrs2GjJB+dBsk0WgSxVUZP0jV59ecJJhWz5IvtjDJ7UkuwmDv27GLDVnuADS4uAeXXhUKKHnCgkYXLgOsSbp52e9oq2ulMNCZ3RWtFLHE/phShPYjDvZ/8grG2WKkhsf65cR71CIOaOfDbf6AfOyUr0xFLeGg+elSE/g4IHe4yCZodAjGlvE78jkBdEIXb6wmr0nWY033KiunMyWX2ERey5rcQ1XI4YuUgP2ApCd Generated-by-Nova\n",
			Fingerprint: "68:17:18:4c:4f:a7:05:68:71:01:3c:f1:db:4c:38:4f",
		},
		{

			Name:        "c2c-secondkey",
			PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD2iZGC1Gr6YWjjobcqA5SP150HA4FzCbDk1r91KKr7GHpeVSsnaXM+e/Eh6VB04ahnF4bllz3fgXfeFgTbVfMbIOcomqU1KmtZpcXOjCVrHo4I6dpAisZ8yO6mBjbbl440Xgocs2UqnWVTlW2vf0O0IFmPODJLN7P1r1r2Vfd0gnZpGN5/J8HvzsQLtdbmttl/ylxkbrq/20bIlY1VF3FXNO7KeREJTwDgdo3xRTFBXMkbSAj9b7dBlN4nhB0lXl8lnAGAA+nkJ69Av3UE0yaG2jYJcW9yjqLuIH3GMFGXSdqLMitdGqfR0o8RhHW40xZM7wMOXUbShy2r4u7uGUXv Generated-by-Nova\n",
			Fingerprint: "a7:c0:d6:6b:60:32:15:fe:f7:74:37:18:85:7a:fa:31",
		},
	}
	th.AssertDeepEquals(t, expected, actual)
}
