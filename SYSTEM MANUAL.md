# SYSTEM MANUAL

## Installation

Test environments of the project

+ macOS Catalina, Intel Core i7
+ Golang 1.15
+ Docker 20.10.11
  + Mongo 5.0.5 (Deployed using mongo:latest image)
  + Hadoop 3.3.1 (Deployed using ubuntu:latest image)

## Configuration

+ Docker configuration file please refer to `configs/docker-compose.yaml`
+ MongoDB configuration file for each node please refer to `configs/mongo-docker`
+ Using `mongos -f mongos.yaml` to innitiate the mongo router node
  + Other node can directly use the mongo image and initate using `--config`
  + Mongo router node can only start with `mongos`
    + The walkaround is to install MongoDB from a ubuntu image instead of using mongo image
+ Configure the hdfs follow some classic tutorial
+ Using `go build` to build the project

> Some useful terminal commands
>
> + Show all containers ip address:
>
>   `docker inspect --format='{{.Name}} - {{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)`
>
>   > /config2 - 192.168.1.7
>   > /shard1_slave - 192.168.1.3
>   > /shard2_slave - 192.168.1.5
>   > /shard2_master - 192.168.1.4
>   > /config1 - 192.168.1.6
>   > /mongos1 - 192.168.1.8
>   > /hdfs - 192.168.1.9
>   > /shard1_master - 192.168.1.2

#### MongoDB Cluster Setup

+ entering shard1_master, shard2_master (using `mongosh`) and build two shards with replica mode
  + here s1m is the network alias name for shard1_master
  + `config={
       _id:'rs1',members:[{
       _id:0,host:'s1m:27017',priority:2},{
       _id:1,host:'s1s:27017',priority:1}]}`
  + `rs.initiate(config)`
+ entering config1 and build config server shard
+ entering router node and add shards
  + `sh.addShard("rs1/172.17.0.6:27017,172.17.0.5:27017") `
  + `sh.addShard("rs2/172.17.0.7:27017,172.17.0.3:27017") `

## Operation

+ Supported operation has been suggested in the CLI

+ Some useful commands are listed here:
  + `query article aid le 99 category eq technology;` 
  + query collections (user, read, article, beread, popular)
  + `show hdfs <path>`
  + `show collections`
  + `show shards`
  + `ping`
  + `set display_details true`: set if show picture and content from hdfs
  + `set sharding true`: shards initiate (zone assignment for each key should be configure in mongo router node)
  + `set logging true/false`