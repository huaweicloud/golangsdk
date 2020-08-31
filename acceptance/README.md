# Golangsdk Acceptance tests

The purpose of these acceptance tests is to validate that SDK features meet
the requirements of a contract - to consumers, other parts of the library, and
to a remote API.

> **Note:** Because every test will be run against a real API endpoint, you
> may incur bandwidth and service charges for all the resource usage. These
> tests *should* remove their remote products automatically. However, there may
> be certain cases where this does not happen; always double-check to make sure
> you have no stragglers left behind.

### Step 1. Set environment variables

Tests rely on environment variables for configuration, so you will need
to set them manually before running the suite.

#### Authentication

|Name|Description|
|---|---|
|`OS_USERNAME`|Your API username|
|`OS_PASSWORD`|Your API password|
|`OS_DOMAIN_NAME`|Your API account name|
|`OS_DOMAIN_ID`|Your API account ID|
|`OS_TENANT_NAME`|Your API project name|
|`OS_TENANT_ID`|Your API project ID|
|`OS_REGION_NAME`|The region you want your resources to reside in|
|`OS_AUTH_URL`|The identity URL you need to authenticate|

#### Compute

|Name|Description|
|---|---|
|`OS_IMAGE_ID`|The ID of the image your want your server to be based on|
|`OS_FLAVOR_ID`|The ID of the flavor you want your server to be based on|
|`OS_FLAVOR_ID_RESIZE`|The ID of the flavor you want your server to be resized to|
|`OS_POOL_NAME`|The Pool from where to obtain Floating IPs|
|`OS_NETWORK_NAME`|The internal/private network to launch instances on|
|`OS_EXTGW_ID`|The external/public network|

#### Shared file systems
|Name|Description|
|---|---|
|`OS_SHARE_NETWORK_ID`| The share network ID to use when creating shares|

### 2. Run the test suite

From the root directory, run:

```
./script/acceptancetest
```

Alternatively, add the following to your `.bashrc`:

```bash
golangsdktest() {
  if [[ -n $1 ]] && [[ -n $2 ]]; then
    pushd  $GOPATH/src/github.com/huaweicloud/golangsdk
    go test -v github.com/huaweicloud/golangsdk/acceptance/openstack/$1 -run "$2" | tee ~/golangsdk.log
    popd
fi
}
```

Then run either groups or individual tests by doing:

```shell
$ golangsdktest networking/v1 TestVpcList
$ golangsdktest networking/v1 TestVpcsCRUD
```

