import xxhash
import jumphash
import socket
import struct


class RedisClient:
    def __init__(self, replicas):
        self.conns = [[] for _ in replicas]
        self.writers = [[] for _ in replicas]
        self.readers = [[] for _ in replicas]
        for i, servers in enumerate(replicas):
            for server in servers:
                ip, port = server.split(":")
                conn = socket.socket()
                conn.connect((ip, int(port)))
                self.conns[i].append(conn)
                self.writers[i].append(conn.makefile("w"))
                self.readers[i].append(conn.makefile("r"))

    @staticmethod
    def find_shard(key, num_servers):
        h = xxhash.xxh64()
        h.update(key.encode("utf-8"))
        return jumphash.jump_consistent(h.intdigest(), num_servers)

    def set(self, key, value):
        shard = self.find_shard(key, len(self.conns[0]))
        result, err = None, None
        for i, replica_conns in enumerate(self.conns):
            conn = replica_conns[shard]
            writer = self.writers[i][shard]
            self.send_command(writer, "SET", key, value)
            result, err = self.read_simple_string(conn)
        return result, err

    def delete(self, key):
        shard = self.find_shard(key, len(self.conns[0]))
        result, err = None, None
        for i, replica_conns in enumerate(self.conns):
            conn = replica_conns[shard]
            writer = self.writers[i][shard]
            self.send_command(writer, "DELETE", key)
            result, err = self.read_simple_string(conn)
        return result, err
