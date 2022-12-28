# Raspberry Pi Kubernetes Cluster

<br>
<img src="doc/initial_cluster_setup.jpg" width="300">

# Hardware

-   2x Raspberry Pi Model B 8GB RAM + 16GB Micro SD
-   2x Raspberry Pi USB-C 5V 3A Power supply
-   1x Cluster Case with fans
-   1x TL-SG105 5 Port Network Switch

# Architecture

k3s has been chosen as the Kubernetes distrubution since it a lightweight distribution optimized for ARM, which is more suitable for Raspberry Pi's.

The cluster consists of 2 nodes.

-   Master node (cluster-master)
    -   Acts as the control plane of the Kubernetes cluster. Manages the worker nodes and the Pods in the cluster
-   Worker node (cluster-worker01)
    -   Worker node that runs the workloads. Runs Pods with the containerized applications

## Overview (WIP)

<img src="doc/cluster-diagram_WIP.png" >

# Cluster

## Deployment flow (WIP)

<img src="doc/deployment-flow_WIP.png" >

## Deployments

### Pi-Hole

Network service that will act as DNS server in my local network. Pi-Hole provides capabilities such as network-wide blocking of ads, telemetry and malware by rerouting network traffic.

### Prometheus + Grafana

Prometheus is Open-source monitoring and alerting solution that will be used for monitoring of the cluster's overall health by using metrics such as RAM and CPU usage from each node.

### Flux (CD)

A open-source set of continuous delivery solutions for Kubernetes that will be used to handle the deployments to the cluster. Flux will listen for changes to a deployment git repository and then sync the state between the repository, containing the manifests describing the cluster state, and the cluster. Flux supports a pull model, which is ideal for this cluster, since it runs on a private network and is not directly exposed to the internet. To deploy to the cluster, you just have to push a new manifest to <PROJECT_ROOT>/clusters/pi-cluster

### nordnet-fetch

My python script for fetching Nordnet account data such as transactions and performance graph data, and store the data in GCP BigQuery. The script will be deployed in the cluster as a CronJob that runs daily.

## Configuration Management
