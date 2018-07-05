package testing

const (
	dehEndpoint = "/dedicated-hosts"
	HostID      = "011d21e2-fbc3-4e4a-9993-9ea223f73264"
)

var allocateRequest = `{
     "availability_zone": "eu-de-02",
     "name": "Test-1",
     "auto_placement": "off",
     "host_type": "h1",
     "quantity": 2
}`

var allocateResponse = `{
    "dedicated_host_ids": [
        "fb4733fd-70a3-44e1-a1cb-0311f028d7e5",
        "7408f985-047d-4313-b3c8-8e12bef01d12"
    ]
}`

var updateRequest = `{
"dedicated_host": 
     {
          "auto_placement": "off",
		   "name": "Test-2"
    }
}`

var getResponse = `
{
    "dedicated_host": {
        "allocated_at": "2018-06-13T07:44:55Z",
        "availability_zone": "eu-de-02",
        "csg_host": "pod01.eu-de-02",
        "name": "test-aj2",
        "available_memory": 270336,
        "released_at": "",
        "auto_placement": "off",
        "available_vcpus": 36,
        "dedicated_host_id": "66156a61-27c2-4169-936b-910dd9c73da3",
        "state": "available",
        "instance_total": 0,
        "host_properties": {           
            "host_type": "h1",
            "vcpus": 36,
            "memory": 270336,
            "cores": 12,
            "sockets": 2,
            "host_type_name": "High performance"
        },
        "csd_host": "fc-nova-compute010#8120665",
        "instance_uuids": [],
        "project_id": "17fbda95add24720a4038ba4b1c705ed"
    }
}
		`

var listResponse = `
{
    "dedicated_hosts": [ {
            "availability_zone": "eu-de-01",
            "name": "c2c-deh-test",
            "available_memory": 262144,
            "auto_placement": "off",
            "available_vcpus": 70,
            "dedicated_host_id": "671611d2-b45c-4648-9e78-06eb24522291",
            "state": "available",
            "instance_total": 2,
            "host_properties": {                
                "host_type": "general",
                "vcpus": 72,
                "memory": 270336,
                "cores": 12,
                "sockets": 2,
                "host_type_name": "General computing"
            },
            "instance_uuids": [
                "3de1ce75-2550-4a46-a689-dd33ca2b62d6",
                "885dc71d-905d-48b5-bae7-db66801dc175"
            ],
            "project_id": "17fbda95add24720a4038ba4b1c705ed"
        }]
}
			`

var listserverResponse = `
{
     "servers": [ {
            "status": "ACTIVE", 
            "flavor": {
                "id": "normal1"
            },
 			"addresses": {
                "0b98c646-617f-4d90-9ca5-385f0cd73ea7": [
                    {
                        "version": 4,
                        "addr": "192.168.3.133"
                    }
                ]
            },
            "id": "3de1ce75-2550-4a46-a689-dd33ca2b62d6",
            "user_id": "6d78fa8550ae45d6932a1fadfb1fa552",
            "name": "c2c-ecs-test-2",
            "tenant_id": "17fbda95add24720a4038ba4b1c705ed",
            "metadata": {
                "metering.image_id": "c0ea3ff1-432e-4650-8a1b-372a80b2d2be",
                "metering.imagetype": "gold",
                "metering.resourcespeccode": "deh.linux",
                "metering.cloudServiceType": "sys.service.type.ec2",
                "image_name": "Standard_CentOS_7_latest",
                "metering.resourcetype": "1",
                "os_bit": "64",
                "vpc_id": "0b98c646-617f-4d90-9ca5-385f0cd73ea7",
                "os_type": "Linux",
                "charging_mode": "0"
            }
        }]
}
			`
