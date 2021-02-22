# Camunda Cloud Kubernetes Zeebe Operator

This repository contains the source code for the Zeebe Operator for Camunda Cloud. 
This Kubernetes Operator allows you to create `ZeebeCluster` Kubernetes resources in your own Kubernetes Cluster that will map to Zeebe Clusters provisioned inside Camunda Cloud. 

The main objective of this operator is to give you native control of your Zeebe Clusters from inside your Kubernetes environment, allowing other applications to have a local object to query to connect to these remote resources. By having a Zeebe Cluster resource type you can manage remote Zeebe Clusters locally and leverage the declarative nature of Kubernetes Resources to manage your Zeebe Clusters inside Camunda Cloud. 



