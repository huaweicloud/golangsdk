// +build acceptance

package v3

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/groups"
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/users"
)

func TestGroupUsersCRUD(t *testing.T) {

	client, err := clients.NewIdentityV3Client()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v", err)
	}

	// Create Group in the default domain
	createGroupOpts := groups.CreateOpts{
		Name:     "testgroup",
		DomainID: "default",
		Extra: map[string]interface{}{
			"email": "testgroup@example.com",
		},
	}
	group, err := CreateGroup(t, client, &createGroupOpts)
	if err != nil {
		t.Fatalf("Unable to create group: %v", err)
	}
	defer DeleteGroup(t, client, group.ID)

	tools.PrintResource(t, group)
	tools.PrintResource(t, group.Extra)

	// Create a test user
	createUserOpts := users.CreateOpts{
		Password: "foobar",
		DomainID: "default",
		Options: map[users.Option]interface{}{
			users.IgnorePasswordExpiry: true,
			users.MultiFactorAuthRules: []interface{}{
				[]string{"password", "totp"},
				[]string{"password", "custom-auth-method"},
			},
		},
		Extra: map[string]interface{}{
			"email": "jsmith@example.com",
		},
	}
	user, err := CreateUser(t, client, &createUserOpts)
	if err != nil {
		t.Fatalf("Unable to create user: %v", err)
	}
	defer DeleteUser(t, client, user.ID)

	tools.PrintResource(t, user)
	tools.PrintResource(t, user.Extra)

	err = users.AddUserToGroup(client, group.ID, user.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to add user to group: %v", err)
	}

	err = users.CheckGroupUser(client, group.ID, user.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Didn't successfully add user to group")
	}

	err = users.DeleteGroupUser(client, group.ID, user.ID).ExtractErr()
	if err != nil {
		t.Fatalf("Unable to remove user from group: %v", err)
	}
}
