# rabbitmq

# Quick Start

## Step 1

```bash
make all
```

## Step 2

```bash
# Consumer 1
go run . consumer one

# Now in another terminal window
go run . consumer two

```

## Step 3

```bash
# In a separate terminal window 
go run . producer
```




```bash
kubectl exec --stdin --tty rabbitmq-7d6959db66-49xbx -- /bin/bash
```


```bash

docker run -d \
  -it \
  --name rabbitmq \
  --mount type=bind,source="$(pwd)"/config,target=/etc/rabbitmq/conf.d \
  -p 15672:15672 -p 5671:5671 \
  rabbitmq:3-management


docker run -d -v config/:/etc/rabbitmq/conf.d -p 15672:15672 -p 5671:5671 rabbitmq:3-management
  
docker exec -ti rabbitmq /bin/bash  
```