# Go-ddd-clean-architect
A sample project layout for golang clean architect and domain driven design

## Prequiresites
```
#not have yet? following https://docs.docker.com/engine/install/ubuntu/
docker
#not have yet? following https://docs.docker.com/compose/install/linux/
docker-compose
#not have yet? following https://cmake.org/install/ or https://snapcraft.io/install/cmake/ubuntu
cmake
```

## Prepare

```
make prepare
```

## Build app
```
make build
```

## Start everything
```
make up
```

## generate data
because hardcoding for generate test data is not a really good ideas.
but you may want to do it yourself.
Therefor I created a stored procedure name `generate_test_data`
input: amount of record that you want to generate.

ex generate 100k records:
```
call app.generate_test_data(100000);
```

query data
```
select BIN_TO_UUID(PK, 1), 
        status, 
        delivery_start_time, 
        delivery_commit_time,
        BIN_TO_UUID(ref_shipment_id, 1),
        BIN_TO_UUID(original_shipment_id, 1),
        BIN_TO_UUID(ref_order_id, 1),
        BIN_TO_UUID(client_order_id, 1),
        client_order_code,
        partner_name,
        from_facility,
        to_facility,
        delivery_address,
        job_name,
        BIN_TO_UUID(job_id, 1),
        volume,
        weight,
        receiver_name,
        shipment_total from app.db_jobs limit 100;
```

## Sample requests

### Create index in opensearch
```
#create jobs index
curl --location 'localhost:9090/jobs/opensearch/create-index/jobs' \
--header 'Content-Type: application/json' \
--data '{}'
```

### Push sample data from database to specific index
```
#to jobs index
curl --location 'localhost:9090/jobs/opensearch/push-documents/jobs' \
--header 'Content-Type: application/json' \
--data '{}'
```

### Search from MySQL stored procedure
```
curl --location 'localhost:9090/jobs/search-by-db?term=job_name&pageIndex=0&pageAmount=50'
```

### Search from opensearch
```
curl --location 'localhost:9090/jobs/search?term=order_code&pageIndex=0&pageAmount=50&index=jobs'
```




## troubleshooting timeout when using command generate data
```
Go to Edit -> Preferences -> SQL Editor and set to a higher value this parameter: DBMS connection read time out (in seconds). For instance: 86400.

Close and reopen MySQL Workbench. Kill your previously query that probably is running and run the query again.
```


OpenSearch Node start then stop and you got error: 

```
2023-04-12 12:21:26 [1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
```

try this command:
```
sudo sysctl -w vm.max_map_count=262144
```


test migration
```
#issue of tidb
set global tidb_skip_isolation_level_check=1

#start migrate
migrate -source "file://infrastructure/migrations" -database "mysql://root@tcp(localhost:4000)/app" up 2
```