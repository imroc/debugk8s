# debugk8s

debug tools for kubernetes

## debug-unhealthy

found unhealthy pod and ping it, check network connectivity.

``` bash 
$ ./debug-unhealthy --kubeconfig ~/.kube/config
2019/06/17 22:25:26 Starting Event controller
unhealthy ip: 172.16.0.221
ping 172.16.0.221
sent 3, received 3, loss 0.000000%
```
