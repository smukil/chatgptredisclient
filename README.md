# chatgptredisclient
A distributed, fault-tolerant Redis Client POC completely written by prompting ChatGPT.


You can start a toy local cluster with 2 shards and 2 replicas each using `start_redis_cluster.py`.
Then compile and run `myredisapp.go` to run some very basic commands against the cluster.

To test fault-tolerance, try killing nodes and seeing if it still works.

To validate scalability, you'll see your keyspace split between shards.

The full transcript is available in `chat-transcript/full_transcript`
