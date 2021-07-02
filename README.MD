# To run prometheus with docker:
`sudo docker run --net="host" -p 9090:9090 -v /etc/prometheus/:/etc/prometheus prom/prometheus`
# To run alert manager with docker:
`docker run -v ~/go/src/tezos/missedEventsNotifier/config:/configs --name alertmanager -p 127.0.0.1:9093:9093 quay.io/prometheus/alertmanager --config.file=/configs/alertmanager.yml`