# Desafio FullCycle Stress Test

Esse desafio consiste em criar um container com uma aplicação que faça um stress test em um site de sua escolha.

## Tecnologias

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)

## Como executar
- Execute o comando `docker compose up -d` para iniciar o container.

## Como testar
- Para aprender mais sobre o CLI, você pode executar o seguinte comando:

```docker
 docker compose run --rm go-stress-test -h 
```
Resposta:

```
Usage of ./stress-test:
  -concurrency int
        Amount of concurrent requests to send
  -requests int
        Maximum amount of requests to send
  -url string
        URL to test
```
- Para testar é necessário passar os seguintes parâmetros: 
  - `-url`: URL do site que será testado.
  - `-requests`: Quantidade máxima de requisições a serem enviadas.
  - `-concurrency`: Quantidade de requisições concorrentes a serem enviadas.

- Para executar um teste, você pode executar o seguinte comando:

```docker
docker compose run --rm go-stress-test -url http://google.com -requests 137 -concurrency 2 
```

## Resultado
- Será exibido o resultado do teste, como por exemplo:
```json
{
    "requests": 137,
    "concurrency": 2,
    "total_duration": "30.162699173s",
    "average_duration": "220.165687ms",
    "status": {
        "200": 137
    },
    "errors": [],
    "errors_count": 0
}
```