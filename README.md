## Description

[Ginkgo](https://github.com/onsi/ginkgo) with [Powerfulseal](https://github.com/powerfulseal/powerfulseal) testing tools for Kubernetes. This repository aims to conduct some End-To-End testing for Kubernetes. For example, kill a pod and it should be rescheduled and become running again. 

Since Powerfulseal can not interact with on-premise kubernetes cluster for `nodeAction`, we introduce the script of [kvaps](https://github.com/kvaps/kubectl-node-shell) to use in Powerfulseal. This enable the actions to node by using SSH. Like reboot.

## Prerequisites

* Golang 1.12 up
* [Ginkgo](https://onsi.github.io/ginkgo/#getting-ginkgo)
* Docker

## Usage

### Start

```
sudo bash dev.sh --config-path /path/to/config -- start-test-env
```

This will bring up `Powerfulseal` container that we need to use.

- `config-path`: should be the target k8s cluster config
- Be sure above command is executed under Kubernetes_e2e folder

If you are using on-premise kubernetes cluster that uses a private registry:

```
sudo bash dev.sh --config-path /path/to/config --target-server $your_registry_ip -- start-test-env
```

If you are using minikube for testing:

```
sudo bash dev.sh --config-path /path/to/config --use-minikube -- start-test-env
```

- Be sure current user is right for minikube since it will mount `~/.minikube` folder into Powerfulseal container to use some certifications for minikube.

### Clear

```
sudo bash dev.sh -- clean-test-env
```

This will remove `Powerfulseal` container.

### Run test

You can write your testcases under `tests/testcases` folder and add into `tests/k8s_test.go`.

To run all testcases:

```
ginkgo -r -timeout 3600s
```

To run specific testcases:

```
ginkgo -r -timeout 3600s --focus="$testcaseName"
```

- `$testcaseName` will be the name of your Ginkgo context defined in `tests/k8s_test.go`. 

Example: 

```
ginkgo -r -timeout 3600s --focus="Testing k8s health"
```

To pass some arguments you want to use:

```
ginkgo -r -timeout 3600s -- -test1="parameter1" -test2="parameter2"
```
