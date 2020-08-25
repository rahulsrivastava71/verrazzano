# Verrazzano
> **NOTE**: This is an early alpha release of Verrazzano. It is suitable for investigation and education usage. It is not suitable for production use. 

## Introduction
Verrazzano is an end-to-end Enterprise Container Platform for deploying cloud-native and traditional applications in multi-cloud and hybrid environments. It is made up of a curated set of open source components – many that you may already use and trust, and some that were written specifically to pull together all of the pieces to make this a cohesive and easy to use platform. 

Verrazzano Enterprise Container Platform includes the following capabilities:

- Hybrid and multi-cluster workload management
- Special handling for WebLogic, Coherence, and Helidon applications
- Multi-cluster infrastructure management
- Integrated and pre-wired application monitoring
- Integrated security
- DevOps and GitOps enablement

This repository contains installation scripts and example applications for use with Verrazzano.

> **NOTE**: This is an early alpha release of Verrazzano. Some features are still in development. 

## tl;dr
To install Verrazzano, follow these steps:  
1. Create an [Oracle Cloud Infrastructure Container Engine for Kubernetes (OKE)](https://docs.cloud.oracle.com/en-us/iaas/Content/ContEng/Concepts/contengoverview.htm) cluster.
2. Launch an [OCI Cloud Shell](https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/cloudshellgettingstarted.htm).
3. Set up a [kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/) file in the OCI Cloud Shell for the OKE cluster. See these detailed [instructions](https://docs.cloud.oracle.com/en-us/iaas/Content/ContEng/Tasks/contengdownloadkubeconfigfile.htm).
4. Clone this [repo](https://github.com/verrazzano/verrazzano) into the home directory of the OCI Cloud Shell.
   - `git clone https://github.com/verrazzano/verrazzano`
   - `cd verrazzano`
5. Execute the following commands in the OCI Cloud Shell:
   - `export CLUSTER_TYPE=OKE`
   - `export VERRAZZANO_KUBECONFIG=~/.kube/config`
   - `export KUBECONFIG=~/.kube/config`
   - `kubectl create secret docker-registry ocr --docker-username=<username> --docker-password=<password> --docker-server=container-registry.oracle.com`
   - `./install/1-install-istio.sh`
   - `./install/2a-install-system-components-magicdns.sh`
   - `./install/3-install-verrazzano.sh`
   - `./install/4-install-keycloak.sh`
6. (Optional) Install some example applications - see below for details.
7. In the case that Verrazzano needs to be uninstalled, execute the following command in the OCI Cloud Shell:
   - `./uninstall/uninstall-verrazzano.sh`

> **NOTE**: This alpha release of Verrazzano is intended for installation in a single OKE or Oracle Linux Cloud Native Environment (OLCNE) cluster. You should only install Verazzano in a cluster that can be safely deleted when your evaluation is complete.

## Deploying the example applications

To deploy the example applications, please see the following instructions:

* [Helidon Hello World](./examples/hello-helidon/README.md)
* [Bob's Books](./examples/bobs-books/README.md)

## More Information

More detailed [installation instructions](./install/README.md) can be found in the `install` directory.

More detailed [uninstall instructions](./uninstall/README.md) can be found in the `uninstall` directory
