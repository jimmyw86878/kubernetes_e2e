#!/bin/bash
set -e
script_start=`date +%s`
base_dir=$(pwd "$0")

config_path=""
target_server=""
use_minikube="false"

#main function
function usage {
    cat <<EOF
Usage: $0  [options] -- COMMAND

Options:
    --target-server         Target Kubernetes cluster
    --config-path           Give Kubernetes config
    --use-minikube          Using minikube for your cluster
      
Commands:
    start-test-env          To set up test environment
    clean-test-env          To clean test environment
EOF
}

while [ "$#" -gt 0 ]; do
    case "$1" in
    (--target-server)
        target_server=$2
        shift 2
        ;;
    (--config-path)
        config_path=$2
        shift 2
        ;;
    (--use-minikube)
        use_minikube="true"
        shift 1
        ;;
    (--)
        shift
        break
        ;;
    (*)
        usage
        exit 3
        ;;
esac
done

case "$1" in

(start-test-env)
        echo "start-test-env"
        if [ -z "$config_path" ] ; then
            echo "Please assign the config path by using option --config-path"
            exit 1
        fi
        if [ "$use_minikube" = "true" ]; then
            home_dir=$(echo ~)
            docker run -t -d -v $config_path:/root/.kube/config -v $home_dir/.minikube:$home_dir/.minikube -v $base_dir/yamls:/home/testyaml --privileged --name kubeclient --network host --entrypoint /bin/bash bloomberg/powerfulseal:3.1.0    
        else
            docker run -t -d -v $config_path:/root/.kube/config -v $base_dir/yamls:/home/testyaml --privileged --name kubeclient --network host --entrypoint /bin/bash bloomberg/powerfulseal:3.1.0
        fi
        cp -f $base_dir/scripts/kubectl-node_shell $base_dir/kubectl-node_shell
        if [ -z "$target_server" ] ; then
            echo "Target server is empty. Use docker.io alpine."
        else
            sed -i 's/docker.io\/library/'$target_server:5000'/g' $base_dir/kubectl-node_shell
        fi
        docker cp $base_dir/kubectl-node_shell kubeclient:/usr/local/bin/
        ;;
(clean-test-env)
        echo "clean-test-env"
        docker rm -f kubeclient
        ;;
(*)     usage
        exit 0
        ;;
esac
