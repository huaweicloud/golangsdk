package testing

const CreateRequest = `
{ 
  "name": "test-cassandra-01", 
  "datastore": { 
    "type": "GeminiDB-Cassandra", 
    "version": "3.11", 
    "storage_engine": "rocksDB" 
  }, 
  "region": "aaa", 
  "availability_zone": "bbb", 
  "vpc_id": "674e9b42-cd8d-4d25-a2e6-5abcc565b961", 
  "subnet_id": "f1df08c5-71d1-406a-aff0-de435a51007b", 
  "security_group_id": "7aa51dbf-5b63-40db-9724-dad3c4828b58", 
  "password": "Test@123", 
  "mode": "Cluster", 
  "flavor": [ 
    { 
      "num": "3", 
      "size": 500,
      "storage": "ULTRAHIGH",
      "spec_code": "nosql.cassandra.4xlarge.4" 
    } 
  ], 
  "backup_strategy": { 
    "start_time": "08:15-09:15", 
    "keep_days": "8"
  },
  "enterprise_project_id": "0" 
}     `

const CreateResponse = `
{ 
  "id": "39b6a1a278844ac48119d86512e0000bin06", 
  "name": "test-cassandra-01", 
  "datastore": { 
    "type": "GeminiDB-Cassandra", 
    "version": "3.11", 
    "storage_engine": "rocksDB" 
  },
  "status": "creating",
  "region": "aaa", 
  "availability_zone": "bbb", 
  "vpc_id": "674e9b42-cd8d-4d25-a2e6-5abcc565b961", 
  "subnet_id": "f1df08c5-71d1-406a-aff0-de435a51007b", 
  "security_group_id": "7aa51dbf-5b63-40db-9724-dad3c4828b58", 
  "mode": "Cluster", 
  "flavor": [ 
    { 
      "num": "3", 
      "size": "500",
      "storage": "ULTRAHIGH",
      "spec_code": "nosql.cassandra.4xlarge.4" 
    } 
  ], 
  "backup_strategy": { 
    "start_time": "08:15-09:15", 
    "keep_days": "8"
  } ,
  "job_id": "c010abd0-48cf-4fa8-8cbc-090f093eaa2f",
  "enterprise_project_id": "0" 
}
    `
const AllInstancesResponse = `
{  
    "instances": [ 
        { 
            "id": "8436a91546294036b75931e879882200in06", 
            "name": "GeminiDB-efa6", 
            "status": "normal", 
            "port": "8635", 
            "mode": " Cluster", 
            "region": "aaa", 
            "datastore": { 
                "type": "GeminiDB-Cassandra", 
                "version": "3.11" 
            }, 
            "engine": " rocksDB ", 
            "created": "2019-01-17T07:05:52", 
            "updated": "2019-01-17T07:05:47", 
            "db_user_name": "rwuser", 
            "vpc_id": "674e9b42-cd8d-4d25-a2e6-5abcc565b961", 
            "subnet_id": "f1df08c5-71d1-406a-aff0-de435a51007b", 
            "security_group_id": "7aa51dbf-5b63-40db-9724-dad3c4828b58", 
            "backup_strategy": { 
                "start_time": "16:00-17:00", 
                "keep_days": 7 
            }, 
            "pay_mode": "0", 
            "maintenance_window": "02:00-06:00", 
            "groups": [ 
                { 
                    "id": "0b0ff12541794e1084f6827e424be2d6gr06",
                    "status": "creating",
                    "volume": { 
                        "size": "100", 
                        "used": "0.003" 
                    }, 
                    "nodes": [ 
                        { 
                            "id": "233eaac9c6f245c0bb9c2d21eea12d1bno06", 
                            "name": "GeminiDB-efa6_priam_node_2", 
                            "status": "normal", 
                            "private_ip": "192.168.0.174", 
                            "spec_code": "nosql.cassandra.xlarge.4", 
                            "availability_zone": "bbb" 
                        }, 
                        { 
                            "id": "d57d76d6320a4a7b86db82c317550c4ano02", 
                            "name": "GeminiDB-efa6_priam_node_1", 
                            "status": "normal", 
                            "private_ip": "192.168.0.39", 
                            "spec_code": "nosql.cassandra.xlarge.4", 
                            "availability_zone": "bbb" 
                        }, 
                        { 
                            "id": "f46b0a1cf4d9400e9fd7af17f8742d37no02", 
                            "name": "GeminiDB-efa6_prima_node_3", 
                            "status": "normal", 
                            "private_ip": "192.168.0.176", 
                            "spec_code": "nosql.cassandra.xlarge.4", 
                            "availability_zone": "bbb" 
                        } 
                    ] 
                }
            ], 
            "enterprise_project_id": "0", 
            "time_zone": "", 
            "actions": [ 
              "CREATE" 
             ] 
        }, 
        { 
            "id": "1236a91546294036b75931e879882200in06", 
            "name": "GeminiDB-efa7", 
            "status": "normal", 
            "port": "8635", 
            "mode": " Cluster ", 
            "region": "aaa", 
            "datastore": { 
                "type": "GeminiDB-Cassandra", 
                "version": "3.11" 
            }, 
            "engine": " rocksDB ", 
            "created": "2019-01-17T07:05:52", 
            "updated": "2019-01-17T07:05:47", 
            "db_user_name": "rwuser", 
            "vpc_id": "674e9b42-cd8d-4d25-a2e6-5abcc565b961", 
            "subnet_id": "f1df08c5-71d1-406a-aff0-de435a51007b", 
            "security_group_id": "7aa51dbf-5b63-40db-9724-dad3c4828b58", 
            "backup_strategy": { 
                "start_time": "16:00-17:00", 
                "keep_days": 7 
            }, 
            "pay_mode": "0", 
            "maintenance_window": "02:00-06:00", 
            "groups": [ 
                { 
                    "id": "0b0ff12541794e1084f6827e424be2d1gr07",
                    "status": "creating",
                    "volume": { 
                        "size": "100", 
                        "used": "0.003" 
                    }, 
                    "nodes": [ 
                        { 
                            "id": "233eaac9c6f245c0bb9c2d21eea12d1bno06", 
                            "name": "GeminiDB-efa7_priam_node_2", 
                            "status": "normal", 
                            "private_ip": "192.168.0.174", 
                            "spec_code": "nosql.cassandra.xlarge.4", 
                            "availability_zone": "bbb" 
                        }, 
                        { 
                            "id": "d57d76d6320a4a7b86db82c317550c4ano02", 
                            "name": "GeminiDB-efa7_priam_node_1", 
                            "status": "normal", 
                            "private_ip": "192.168.0.39", 
                            "spec_code": "nosql.cassandra.xlarge.4", 
                            "availability_zone": "bbb" 
                        }, 
                        { 
                            "id": "f46b0a1cf4d9400e9fd7af17f8742d37no02", 
                            "name": "GeminiDB-efa7_prima_node_3", 
                            "status": "normal", 
                            "private_ip": "192.168.0.176", 
                            "spec_code": "nosql.cassandra.xlarge.4", 
                            "availability_zone": "bbb" 
                        } 
                    ] 
                } 
            ], 
            "enterprise_project_id": "0", 
            "time_zone": "", 
            "actions": [ 
              "CREATE" 
             ] 
        } 
    ], 
    "total_count": 2 
}
`
