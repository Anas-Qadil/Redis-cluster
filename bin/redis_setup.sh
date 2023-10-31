#!/bin/bash

# Create a Redis cluster with three nodes

# Set the cluster name
CLUSTER_NAME=redis-cluster

# Set the node IDs and names
NODE_IDS=(1 2 3)
NODE_NAMES=("node-1" "node-2" "node-3")

# Create the data directory for each node
for node_id in "${NODE_IDS[@]}"; do
  mkdir -p data/node-${node_id}
done

# Start a Redis container for each node
for node_id in "${NODE_IDS[@]}"; do
  docker run -d --name redis-cluster-node-${node_id} --network redis-cluster-network -v data/node-${node_id}:/data redis:latest
done

# Wait for all of the nodes to be ready
until redis-cli -c -h 172.17.0.1:7000 cluster info; do
  sleep 1
done

# Create the Redis cluster
redis-cli -c -h 172.17.0.1:7000 cluster create --cluster-replicas 1 ${NODE_NAMES[@]}

# Connect to the Redis cluster
redis-cli -c -h 172.17.0.1:7000
