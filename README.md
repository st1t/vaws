# Vaws

The vaws command was created to simplify the display of AWS resources.  
This repository is a Go version of the command that was created in the following repository.  
https://github.com/st1t/vaws

## Install

```shell
$ git clone git@github.com:st1t/vaws.git
$ cd vaws/
$ make install
```

Download the appropriate one for your CPU architecture from the following site.  
If you are using a Mac, you may get a developer validation error.  
In that case, right-click in the Finder and select Open.  
https://github.com/st1t/vaws/releases

## Usage

```bash
$ vaws -h
The vaws command was created to simplify the display of AWS resources.

Usage:
  vaws [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  ec2         Show EC2 instances.
  elb         Show ELB.
  help        Help about any command
  rds         Show RDS instances.
  sg          Show Security Group
  subnet      Show subnet
  vpc         Show VPC

Flags:
  -p, --aws-profile string   -p my-aws
  -h, --help                 help for vaws
  -s, --sort-position int    -s 1 (default 1)
  -v, --version              version for vaws

Use "vaws [command] --help" for more information about a command.
$
```

### EC2

```shell
$ vaws ec2 -p my-aws
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| NAME  |         ID          |   TYPE   |  PRIVATE IP   |   PUBLIC IP   |  STATE  |            SECURITY GROUP             |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| app01 | i-06d4c29e4e5ccadc4 | t2.micro | 172.31.35.175 | 54.238.30.226 | running | launch-wizard-2(sg-0f0b4c4642ffb5ef2) |
| app02 | i-06723a6629e542c50 | t3.small | 172.31.21.33  | 18.179.33.182 | running | launch-wizard-3(sg-08d35fef29987e75e) |
| web01 | i-0abee92626b0a28a7 | t3.nano  | 172.31.18.8   | 35.73.127.100 | running | launch-wizard-1(sg-0d642190887707fd0) |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
```

If you want to sort by a specific column, use the S option.  
The following command sorts by SecurityGroup column.
```shell
$ vaws ec2 -p my-aws -s 7
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| NAME  |         ID          |   TYPE   |  PRIVATE IP   |   PUBLIC IP   |  STATE  |            SECURITY GROUP             |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
| web01 | i-0abee92626b0a28a7 | t3.nano  | 172.31.18.8   | 35.73.127.100 | running | launch-wizard-1(sg-0d642190887707fd0) |
| app01 | i-06d4c29e4e5ccadc4 | t2.micro | 172.31.35.175 | 54.238.30.226 | running | launch-wizard-2(sg-0f0b4c4642ffb5ef2) |
| app02 | i-06723a6629e542c50 | t3.small | 172.31.21.33  | 18.179.33.182 | running | launch-wizard-3(sg-08d35fef29987e75e) |
+-------+---------------------+----------+---------------+---------------+---------+---------------------------------------+
```

## RDS

```shell
$ vaws rds -p my-aws
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
|     CLUSTER     |  STATUS   |        INSTANCES         |                            WRITE-ENDPOINT                             |                              READ-ENDPOINT                               |
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
| test-cluster-01 | available | test-cluster-instance-01 | test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com | test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com |
| test-cluster-01 | available | test-cluster-instance-02 | test-cluster-01.cluster-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com | test-cluster-01.cluster-ro-cb8aaaaaaaaa.ap-northeast-1.rds.amazonaws.com |
+-----------------+-----------+--------------------------+-----------------------------------------------------------------------+--------------------------------------------------------------------------+
```

## SecurityGroup

```shell
$ vaws sg -p my-aws
+-----------------+---------+----------------------+------+----------------------+-----------------------+
|      NAME       |  TYPE   |          ID          | PORT |        SOURCE        |          VPC          |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
| default         | inbound | sg-0d642190887707fd0 |   -1 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-1 | inbound | sg-0d642190887707fd0 |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 8.8.8.8/32           | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | 0.0.0.0/0            | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   22 | sg-0d642190887707fd0 | vpc-0f9999c7db8c44b21 |
| launch-wizard-2 | inbound | sg-08d35fef29987e75e |   53 | pl-61a12345          | vpc-0f9999c7db8c44b21 |
+-----------------+---------+----------------------+------+----------------------+-----------------------+
```

## VPC

```shell
$ vaws vpc
+--------------+--------------+-------------+
|     NAME     |      ID      |    CIDR     |
+--------------+--------------+-------------+
| hoge service | vpc-123ZZZZZ | 10.0.0.0/16 |
| default vpc  | vpc-123XXXXX | 10.1.0.0/16 |
+--------------+--------------+-------------+
```

## Subnet

```shell
+----------------+-----------------+-------------+--------------+-----------------+-----------+---------------+--------------------+
|      NAME      |    SUBNET ID    |    CIDR     |     VPC      |       AZ        |   AZ ID   | MAP PUBLIC IP | AVAILABLE IP COUNT |
+----------------+-----------------+-------------+--------------+-----------------+-----------+---------------+--------------------+
| test-subnet-01 | subnet-yyyyyyyy | 10.1.0.0/24 | vpc-12345678 | ap-northeast-1a | apne1-az4 | false         |                250 |
| test-subnet-02 | subnet-xxxxxxxx | 10.2.0.0/24 | vpc-12345678 | ap-northeast-1a | apne1-az4 | true          |                250 |
| test-subnet-03 | subnet-zzzzzzzz | 10.3.0.0/24 | vpc-12345678 | ap-northeast-1a | apne1-az4 | false         |                250 |
+----------------+-----------------+-------------+--------------+-----------------+-----------+---------------+--------------------+
```

## ELB

```shell
$ vaws elb
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
|    LB     |    TYPE     |     SCHEME      |     VPC      |                      SUBNET                       |   SECURITY GROUP    | IP TYPE |                  DNS NAME                  |
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
| test-lb01 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | sg-084d3a6xxxxxxxxx | ipv4    | test-lb01.ap-northeast-1.elb.amazonaws.com |
| test-lb02 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | sg-084d3a6xxxxxxxxx | ipv4    | test-lb02.ap-northeast-1.elb.amazonaws.com |
| test-lb03 | application | internet-facing | vpc-xxxxxxxx | subnet-1234567e3xxxxxxxx,subnet-1234567e3zzzzzzzz | none                | ipv4    | test-lb03.ap-northeast-1.elb.amazonaws.com |
+-----------+-------------+-----------------+--------------+---------------------------------------------------+---------------------+---------+--------------------------------------------+
```

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
