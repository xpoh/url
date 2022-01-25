docker stop prometheus
docker rm prometheus

docker run \
    -d \
    -p 9090:9090 \
    -v ~/projects/GB/opt_go_app/lesson-1/scripts/config:/etc/prometheus \
    --name prometheus \
    prom/prometheus

