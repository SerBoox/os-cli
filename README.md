# OpenStack Client (os-cli)
> Console client via OpenStack REST API 

#### 1) Create Instance
##### Command:
``` bash
 ./os-cli create --auth-host http://10.0.2.15/identity/v3/auth/tokens \
  --login admin \
  --pass pass1 \
  --instName test-2 \
  --imageRef 81311795-d489-43d6-b12c-e982db3824a7 \
  --flavorRef 42

```
##### Stdout:
``` bash
+--------------------------------------+---------+--------------+
|             INSTANCE ID              | STATUS  |  ADMINPASS   |
+--------------------------------------+---------+--------------+
| e5e50439-ec0a-4f4a-9d68-daacef3e3dba | success | fHqV9oz4ZueR |
+--------------------------------------+---------+--------------+


```

#### 2) List Instances
##### Command:
``` bash
 ./os-cli show --auth-host http://10.0.2.15/identity/v3/auth/tokens \
  --login admin \
  --pass pass1

```
##### Stdout:
``` bash
+---------------+--------------------------------------+-------------+----------+--------+-------------------+
| INSTANCE NAME |               IMAGE ID               | IP ADDRESS  | FLAVORID | STATUS | AVAILABILITY ZONE |
+---------------+--------------------------------------+-------------+----------+--------+-------------------+
| test-2        | 81311795-d489-43d6-b12c-e982db3824a7 | 2001:db8::9 | 42       | ACTIVE | nova              |
+---------------+--------------------------------------+-------------+----------+--------+-------------------+
| test-2        | 81311795-d489-43d6-b12c-e982db3824a7 | 10.0.2.228  | 42       | ACTIVE | nova              |
+---------------+--------------------------------------+-------------+----------+--------+-------------------+
| test-1        |                                      | 2001:db8::3 | 42       | ACTIVE | nova              |
+---------------+--------------------------------------+-------------+----------+--------+-------------------+
| test-1        |                                      | 10.0.2.231  | 42       | ACTIVE | nova              |
+---------------+--------------------------------------+-------------+----------+--------+-------------------+

```
* You can use ```./os-cli --help``` 
* The application support only DevStack Ocata
