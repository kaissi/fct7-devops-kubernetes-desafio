# fct7-devops-kubernetes-desafio
Code Education | Full Cycle Turma 7 | DevOps | Kubernetes - Projeto Prático - Utilizando K8s

## 1 - Servidor Web - Nginx

Pasta [./nginx](./nginx)

## 2 - Configuração do MySQL

Pasta [./mysql](./mysql)

## 3 - Desafio Go

https://hub.docker.com/r/kaissi/devops-kubernetes-desafio-go

Pasta [./go](./go)

Para executar os testes unitários:
```bash
go test ./go/github.com/kaissi/fct7-devops-kubernetes-desafio/ -v
```

Para gerar a imagem Docker do desafio K8s:

```bash
docker build ./go -t kaissi/devops-kubernetes-desafio-go:latest
```

Para executar o container:

```bash
docker run --rm --detach \
    --name desafio-k8s \
    --publish 8000:8000 \
    kaissi/devops-kubernetes-desafio-go:latest
```
Abrir o endereço http://localhost:8000

Resultado:
> Code.education Rocks!