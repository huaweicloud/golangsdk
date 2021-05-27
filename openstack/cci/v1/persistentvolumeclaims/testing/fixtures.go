package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/cci/v1/persistentvolumeclaims"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

const (
	expectedRequest = `
{
	"apiVersion": "v1",
	"kind": "PersistentVolumeClaim",
	"metadata": {
		"annotations": {
			"fsType": "ext4",
			"volumeID": "46d10433-35fb-459b-8482-2dd6fa192262"
		},
		"name": "cci-ssd-test-demo",
		"namespace": "terraform-test"
	},
	"spec": {
		"resources": {
			"requests": {
				"storage": "10Gi"
			}
		},
		"storageClassName": "ssd"
	}
}`

	expectedCreateResponse = `
{
	"kind": "PersistentVolumeClaim",
	"apiVersion": "v1",
	"metadata": {
		"name": "cci-ssd-test-demo",
		"namespace": "terraform-test",
		"selfLink": "/api/v1/namespaces/terraform-test/persistentvolumeclaims/cci-ssd-test-demo",
		"uid": "1ae26e9c-9e39-432d-9438-e6c669670629",
		"resourceVersion": "285666212",
		"creationTimestamp": "2021-05-25T03:45:45Z",
		"labels": {
			"failure-domain.beta.kubernetes.io/region": "cn-north-4",
			"failure-domain.beta.kubernetes.io/zone": "cn-north-4a"
		},
		"annotations": {
			"kubernetes.io/volumeId": "46d10433-35fb-459b-8482-2dd6fa192262"
		},
		"finalizers": [
			"kubernetes.io/pvc-protection"
		]
	},
	"spec": {
		"accessModes": [
			"ReadWriteMany"
		],
		"resources": {
			"requests": {
				"storage": "10Gi"
			}
		},
		"volumeName": "cci-evs-import-46d10433-35fb-459b-8482-2dd6fa192262",
		"storageClassName": "ssd",
		"volumeMode": "Filesystem"
	},
	"status": {
		"phase": "Pending"
	}
}`

	expectedListResponse = `
[
	{
		"persistentVolumeClaim": {
			"metadata": {
				"name": "cci-ssd-test-demo",
				"namespace": "terraform-test",
				"selfLink": "/api/v1/namespaces/terraform-test/persistentvolumeclaims/cci-ssd-test-demo",
				"uid": "acf521b6-1e60-4ecd-af4a-671005e6d4bd",
				"resourceVersion": "285728642",
				"creationTimestamp": "2021-05-25T06:26:29Z",
				"labels": {
					"failure-domain.beta.kubernetes.io/region": "cn-north-4",
					"failure-domain.beta.kubernetes.io/zone": "cn-north-4a"
				},
				"annotations": {
					"kubernetes.io/volumeId": "45571ae3-125a-4cfa-bd76-a76eedcde45a"
				},
				"finalizers": [
                    "kubernetes.io/pvc-protection"
                ]
			},
			"spec": {
				"accessModes": [
					"ReadWriteMany"
				],
				"resources": {
					"requests": {
						"storage": "10Gi"
					}
				},
				"volumeName": "cci-evs-import-45571ae3-125a-4cfa-bd76-a76eedcde45a",
				"storageClassName": "ssd",
				"volumeMode": "Filesystem"
			},
			"status": {
				"phase": "Pending"
			}
		},
		"persistentVolume": {
			"metadata": {
				"name": "cci-evs-import-45571ae3-125a-4cfa-bd76-a76eedcde45a",
				"selfLink": "/api/v1/persistentvolumes/cci-evs-import-45571ae3-125a-4cfa-bd76-a76eedcde45a",
				"uid": "e8dc63e5-014d-4e39-970c-5cc151e92144",
				"resourceVersion": "285728644",
				"creationTimestamp": "2021-05-25T06:26:29Z",
				"labels": {
					"failure-domain.beta.kubernetes.io/region": "cn-north-4",
					"failure-domain.beta.kubernetes.io/zone": "cn-north-4a",
					"tenant.kubernetes.io/domain-id": "0970d7b7d400f2470fbec00316a03560",
					"tenant.kubernetes.io/project-id": "0970dd7a1300f5672ff2c003c60ae115"
				},
				"annotations": {
					"kubernetes.io/createdby": "cci-apiserver",
					"pv.kubernetes.io/bound-by-cci": "yes",
					"pv.kubernetes.io/namespace": "terraform-test",
					"tenant.kubernetes.io/domain-id": "0970d7b7d400f2470fbec00316a03560",
					"tenant.kubernetes.io/project-id": "0970dd7a1300f5672ff2c003c60ae115"
				},
				"finalizers": [
					"kubernetes.io/pv-protection"
				]
			},
			"spec": {
				"capacity": {
					"storage": "10Gi"
				},
				"flexVolume": {
					"driver": "huawei.com/fuxivol",
					"fsType": "ext4",
					"options": {
						"fsType": "ext4",
                        "storage_class": "ssd",
						"volumeID": "45571ae3-125a-4cfa-bd76-a76eedcde45a"
					}
				},
				"accessModes": [
					"ReadWriteMany"
				],
				"claimRef": {
					"namespace": "terraform-test",
					"name": "cci-ssd-test-demo"
				},
				"persistentVolumeReclaimPolicy": "Delete",
				"storageClassName": "ssd",
				"volumeMode": "Filesystem"
			},
			"status": {
				"phase": "Available"
			}
		}
	}
]`
)

var (
	createOpts = &persistentvolumeclaims.CreateOpts{
		ApiVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: persistentvolumeclaims.Metadata{
			Name:      "cci-ssd-test-demo",
			Namespace: "terraform-test",
			Annotations: &persistentvolumeclaims.Annotations{
				FsType:   "ext4",
				VolumeID: "46d10433-35fb-459b-8482-2dd6fa192262",
			},
		},
		Spec: persistentvolumeclaims.Spec{
			Resources: persistentvolumeclaims.ResourceRequirement{
				Requests: &persistentvolumeclaims.ResourceName{
					Storage: "10Gi",
				},
			},
			StorageClassName: "ssd",
		},
	}

	expectedCreateResponseData = &persistentvolumeclaims.PersistentVolumeClaim{
		Kind:       "PersistentVolumeClaim",
		ApiVersion: "v1",
		Metadata: persistentvolumeclaims.MetaResp{
			Name:              "cci-ssd-test-demo",
			Namespace:         "terraform-test",
			SelfLink:          "/api/v1/namespaces/terraform-test/persistentvolumeclaims/cci-ssd-test-demo",
			UID:               "1ae26e9c-9e39-432d-9438-e6c669670629",
			ResourceVersion:   "285666212",
			CreationTimestamp: "2021-05-25T03:45:45Z",
			Annotations: map[string]string{
				"kubernetes.io/volumeId": "46d10433-35fb-459b-8482-2dd6fa192262",
			},
			Labels: map[string]string{
				"failure-domain.beta.kubernetes.io/region": "cn-north-4",
				"failure-domain.beta.kubernetes.io/zone":   "cn-north-4a",
			},
			Finalizers: []string{
				"kubernetes.io/pvc-protection",
			},
		},
		Spec: persistentvolumeclaims.SpecResp{
			AccessModes: []string{
				"ReadWriteMany",
			},
			Resources: persistentvolumeclaims.ResourceRequirement{
				Requests: &persistentvolumeclaims.ResourceName{
					Storage: "10Gi",
				},
			},
			VolumeName:       "cci-evs-import-46d10433-35fb-459b-8482-2dd6fa192262",
			StorageClassName: "ssd",
			VolumeMode:       "Filesystem",
		},
		Status: persistentvolumeclaims.Status{
			Phase: "Pending",
		},
	}

	listOpts = persistentvolumeclaims.ListOpts{
		StorageType: "bs",
	}

	expectedListResponseData = []persistentvolumeclaims.ListResp{
		{
			PersistentVolumeClaim: persistentvolumeclaims.PersistentVolumeClaim{
				Metadata: persistentvolumeclaims.MetaResp{
					Name:              "cci-ssd-test-demo",
					Namespace:         "terraform-test",
					UID:               "acf521b6-1e60-4ecd-af4a-671005e6d4bd",
					ResourceVersion:   "285728642",
					SelfLink:          "/api/v1/namespaces/terraform-test/persistentvolumeclaims/cci-ssd-test-demo",
					CreationTimestamp: "2021-05-25T06:26:29Z",
					Annotations: map[string]string{
						"kubernetes.io/volumeId": "45571ae3-125a-4cfa-bd76-a76eedcde45a",
					},
					Labels: map[string]string{
						"failure-domain.beta.kubernetes.io/region": "cn-north-4",
						"failure-domain.beta.kubernetes.io/zone":   "cn-north-4a",
					},
					Finalizers: []string{
						"kubernetes.io/pvc-protection",
					},
				},
				Spec: persistentvolumeclaims.SpecResp{
					AccessModes: []string{
						"ReadWriteMany",
					},
					Resources: persistentvolumeclaims.ResourceRequirement{
						Requests: &persistentvolumeclaims.ResourceName{
							Storage: "10Gi",
						},
					},
					VolumeName:       "cci-evs-import-45571ae3-125a-4cfa-bd76-a76eedcde45a",
					StorageClassName: "ssd",
					VolumeMode:       "Filesystem",
				},
				Status: persistentvolumeclaims.Status{
					Phase: "Pending",
				},
			},
			PersistentVolume: persistentvolumeclaims.PersistentVolumeClaim{
				Metadata: persistentvolumeclaims.MetaResp{
					Name:              "cci-evs-import-45571ae3-125a-4cfa-bd76-a76eedcde45a",
					UID:               "e8dc63e5-014d-4e39-970c-5cc151e92144",
					SelfLink:          "/api/v1/persistentvolumes/cci-evs-import-45571ae3-125a-4cfa-bd76-a76eedcde45a",
					ResourceVersion:   "285728644",
					CreationTimestamp: "2021-05-25T06:26:29Z",
					Annotations: map[string]string{
						"kubernetes.io/createdby":         "cci-apiserver",
						"pv.kubernetes.io/bound-by-cci":   "yes",
						"pv.kubernetes.io/namespace":      "terraform-test",
						"tenant.kubernetes.io/domain-id":  "0970d7b7d400f2470fbec00316a03560",
						"tenant.kubernetes.io/project-id": "0970dd7a1300f5672ff2c003c60ae115",
					},
					Labels: map[string]string{
						"failure-domain.beta.kubernetes.io/region": "cn-north-4",
						"failure-domain.beta.kubernetes.io/zone":   "cn-north-4a",
						"tenant.kubernetes.io/domain-id":           "0970d7b7d400f2470fbec00316a03560",
						"tenant.kubernetes.io/project-id":          "0970dd7a1300f5672ff2c003c60ae115",
					},
					Finalizers: []string{
						"kubernetes.io/pv-protection",
					},
				},
				Spec: persistentvolumeclaims.SpecResp{
					Capacity: persistentvolumeclaims.ResourceName{
						Storage: "10Gi",
					},
					FlexVolume: persistentvolumeclaims.FlexVolume{
						Driver: "huawei.com/fuxivol",
						FsType: "ext4",
						Options: persistentvolumeclaims.Options{
							FsType:   "ext4",
							VolumeID: "45571ae3-125a-4cfa-bd76-a76eedcde45a",
						},
					},
					AccessModes: []string{
						"ReadWriteMany",
					},
					ClaimRef: persistentvolumeclaims.ClaimRef{
						Namespace: "terraform-test",
						Name:      "cci-ssd-test-demo",
					},
					PersistentVolumeReclaimPolicy: "Delete",
					StorageClassName:              "ssd",
					VolumeMode:                    "Filesystem",
				},
				Status: persistentvolumeclaims.Status{
					Phase: "Available",
				},
			},
		},
	}
)

func handlePersistentVolumeClaimCreate(t *testing.T) {
	th.Mux.HandleFunc("/namespaces/terraform-test/extended-persistentvolumeclaims",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = fmt.Fprint(w, expectedCreateResponse)
		})
}

func handlePersistentVolumeClaimList(t *testing.T) {
	th.Mux.HandleFunc("/namespaces/terraform-test/extended-persistentvolumeclaims",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, expectedListResponse)
		})
}
