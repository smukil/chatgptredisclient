PORTS=($(seq 6380 1 6383))

for i in ${PORTS[@]}; do
  redis-server --port $i &
done
