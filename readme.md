# P2P Network Simulator

A REST API application to simulate a p2p network with a tree topology. Following are the endpoints to interact with the program using HTTP calls
1. The first one is where a node can request to join the p2p network. In this step we just assign the node to the best-fitting parent (the node with the most free capacity).
2. With the second endpoint, the node can communicate to the service that it is leaving the network. In this case, we want to reorder the current node tree (not all the network) to build the solution where the tree has the fewest number of depth levels.
3. The last endpoint will reflect the status of the network, returning a list of encoded strings.

see more details about the implementation on [wiki](https://github.com/Uzama/p2p-network-simulator/wiki) 

### How to start 

- Make sure you have installed latest version of docker and docker engine is up and running.
- Clone the service locally and navigate to project root directory.
- To run the service, type ```docker compose up``` and enter.
- Make sure service is up and running. 
- Now you can send request to the service at ```localhost:8080```.

### API Reference

#### Join

```
  POST /join
```

 - Request body
```json
    {
        "id": 1,
        "capacity": 2,
    }
```

- Response 
```json
    {
        "message":"successfully joined",
        "error":false,
        "data":1
    }
```

#### Leave

```
  DELETE /leave/1
```

- Response 
```json
    {
        "message":"successfully left",
        "error":false,
        "data":1
    }
```

#### Trace

```
  GET /trace
```

- Response 
```json
    {
        "message":"trace received",
        "error":false,
        "data":["1(0/2)"]
    }  
```