```docker
docker compose down --volumes --rmi all && docker compose build && docker compose up -d   && docker compose run --rm go-stress-test -url http://google.com -requests 10 -concurrency 2   
```