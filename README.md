# chatgptredisclient
_A distributed, fault-tolerant Redis Client POC in Go completely written by prompting ChatGPT._


* You can start a toy local cluster with 2 shards and 2 replicas each using `start_redis_cluster.py`. (Redis-server must be pre-installed)

* Then compile and run `myredisapp.go` to run some very basic commands against the cluster.

* To test fault-tolerance, try killing nodes and seeing if it still works.

* To validate scalability, you'll see your keyspace split between shards.

* The full transcript is available in `chat-transcript/full_transcript`

* An incomplete Python implementation that it copied from Go to Python in one prompt is in `incomplete_py_implementation.py`
