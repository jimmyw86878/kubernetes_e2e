#!/bin/bash
set -e
script_start=`date +%s`
base_dir=$(pwd "$0")

config_path=""

#main function
function usage {
    cat <<EOF
Usage: $0  [options] -- COMMAND

Options:
    --config-path         Give Kubernetes config
    --testcase            To run specific testcase
      
Commands:
    start-test-env        To set up test environment
EOF
}

while [ "$#" -gt 0 ]; do
    case "$1" in
    (--config-path)
        config_path=$2
        shift 2
        ;;
    (--testcase)
        export GINKGO_TESTCASE=$2
        shift 2
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
        docker run -t -d -v $config_path:/root/.kube/config -v $base_dir/yamls:/home/testyaml --privileged --name kubeclient --network host --entrypoint /bin/bash bloomberg/powerfulseal:3.1.0
        ;;
(clean-test-env)
        echo "clean-test-env"
        docker rm -f kubeclient
        ;;
(run-test)
        echo "run test"
        export ACK_GINKGO_DEPRECATIONS=1.16.2
        ginkgo -r -timeout 3600s
        ;;
(*)     usage
        exit 0
        ;;
esac
