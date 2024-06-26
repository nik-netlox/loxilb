#!/bin/bash

source ../common.sh

echo "#########################################"
echo "Spawning all hosts"
echo "#########################################"

spawn_docker_host --dock-type loxilb --dock-name llb1 
spawn_docker_host --dock-type host --dock-name ep1
spawn_docker_host --dock-type host --dock-name ep2
spawn_docker_host --dock-type host --dock-name c1
spawn_docker_host --dock-type host --dock-name br1

echo "#########################################"
echo "Connecting and configuring  hosts"
echo "#########################################"


connect_docker_hosts c1 br1
connect_docker_hosts llb1 br1
connect_docker_hosts ep1  br1
connect_docker_hosts ep2  br1

config_docker_host --host1 llb1 --host2 br1 --ptype phy --addr 10.2.100.85/16
config_docker_host --host1 ep1 --host2 br1 --ptype phy --addr 10.2.96.58/16 --gw 10.2.100.85
config_docker_host --host1 ep2 --host2 br1 --ptype phy --addr 10.2.89.253/16 --gw 10.2.100.85
config_docker_host --host1 c1 --host2 br1 --ptype phy --addr 10.2.9.19/16 --gw 10.2.100.85

sleep 1

create_docker_host_cnbridge --host1 br1 --host2 llb1
create_docker_host_cnbridge --host1 br1 --host2 ep1
create_docker_host_cnbridge --host1 br1 --host2 ep2
create_docker_host_cnbridge --host1 br1 --host2 c1

##Create LB rule
$dexec llb1 loxicmd create lb 10.2.100.85 --sctp=38412:38412 --endpoints=10.2.96.58:1,10.2.89.253:1 --mode=fullnat

$dexec llb1 bash -c 'for i in /proc/sys/net/ipv4/conf/*/rp_filter; do echo 0 > "$i"; done'

sleep 10
