/*
Package tags provides functionality to operate with server tags that have been
provisioned by the OpenStack Comupte service.

Examples

	serverID := "47b6b7b7-568d-40e4-868c-d5c41735532e"

	// list all tags
	tags, err := ListTags(client, serverID).Extract()
	if err != nil {
		panic(err)
	}

	// put tags
	err := PutTags(client, serverID, "tag1", "tag2")
	if err != nil {
		panic(err)
	}

	// clean tags
	err := CleanTags(client, serverID)
	if err != nil {
		panic(err)
	}

	// check tag existance
	err := CheckTag(client, serverID, "tag1")

	// add a tag to server
	err := AddTag(client, serverID, "tagN")
	if err != nil {
		panic(err)
	}

	// delete a tag from server tag list
	err := DeleteTag(client, serverID, "tagN")
	if err != nil {
		panic(err)
	}

*/
package tags
