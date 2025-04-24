#!/bin/bash
until etcdctl endpoint health; do
  echo "等待 etcd 集群就绪..."
  sleep 1
done
echo "etcd 集群已就绪!"