## Description

This shell file is from [kvaps](https://github.com/kvaps/kubectl-node-shell) and revised for:

```
registry="docker.io/library"
image="$registry/alpine"
...
{
    ...
    "image": "$image",
    "imagePullPolicy": "IfNotPresent",
    ...
}
...
```

And this should be copied into Powerfulseal container to ssh into target node in Kubernetes.

## Usage

Enter Powerfulseal container:

```
docker exec -it kubeclient bash
```

SSH into node in Kubernetes cluster:

```
kubectl node-shell master86
```

Just execute some commands:

```
kubectl node-shell master86 -- ls
```


