# fct7-devops-kubernetes-desafio
Code Education | Full Cycle Turma 7 | DevOps | Kubernetes - Projeto Prático - Utilizando K8s

## 1 - Servidor Web - Nginx

Pasta [./nginx](./nginx)

Para executar o container no Kubernetes:

```bash
kubectl apply -f nginx/deployment.yaml
```
Abrir o endereço http://\<ip-externo-nginx-service\>:8080 \*

Resultado:
> Code.education Rocks

\* Obs: a porta para acessar o cluster onde a aplicação está rodando pode ser diferente de 8080

---

## 2 - Configuração do MySQL

Pasta [./mysql](./mysql)

Para criar a senha do usuário 'root' do MySQL:
```bash
kubectl create secret generic mysql-pass --from-literal=password='a1b2c3d4'
```
Para executar o container no Kubernetes:

```bash
kubectl apply -f mysql/deployment.yaml
```
O serviço do MySQL não é exposto, portanto, para acessar o banco através do console, executar:

```bash
kubectl exec -it service/mysql-service /bin/bash
```

Dentro do container, executar:
```bash
mysql -u root -p
```

---

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

### 3.1 Para executar o container standalone:

```bash
docker run --rm --detach \
    --name desafio-k8s \
    --publish 8000:8000 \
    kaissi/devops-kubernetes-desafio-go:latest
```
Abrir o endereço http://localhost:8000

Resultado:
> <b>Code.education Rocks!</b>

### 3.2 - Para executar o container no Kubernetes

```bash
kubectl apply -f go/deployment.yaml
```

Abrir o endereço http://\<ip-externo-nginx-service\>:8000 \*

Resultado:
> <b>Code.education Rocks!</b>

\* Obs: a porta para acessar o cluster onde a aplicação está rodando pode ser diferente de 8000

### 3.3 - Compilando Docker Multi-Plataforma (386, amd64, arm64 e armv7)

Leia mais sobre [buildx][] para saber como funciona e como habilitá-lo no Docker.

O comando `make help` mostra as opções para compilar o projeto Go.

Para executar o build padrão, gerando apenas para a plataforma local, executar:  
`make build`

Internamente é executado:  
```bash
docker build \
    -t kaissi/devops-kubernetes-desafio-go:latest-amd64 \
    -f ./go/docker/Dockerfile \
    ./go
```

Para executar o buildx, gerando imagem Docker para múltiplas plataformas, fazer:  
`make build-all`

Caso dê algum erro para compilar uma plataforma diferente, executar o seguinte comando para a(s) plataforma(s) com erro:  
```bash
docker run --rm --privileged aptman/qus -s -- -p aarch64
````

Para remover o suporte a uma ou mais plataformas específicas, executar:  
```bash
docker run --rm --privileged aptman/qus -- -r aarch64
```

Para fazer o build para uma plataforma específica, fazer:  
`make build-armv7`

Internamente está sendo executado:  
```bash
docker buildx build \
    --pull \
    --load \
    -t kaissi/devops-kubernetes-desafio-go:latest-armv7 \
    --platform linux/arm/v7 \
    -f ./go/docker/Dockerfile.multi-arch \
    ./go
```

Para que ao executar `docker pull kaissi/devops-kubernetes-desafio-go:latest`, o Docker consiga escolher a plataforma automaticamente, o manifest é criado da seguinte forma:

```bash
docker manifest create \
    kaissi/devops-kubernetes-desafio-go:latest \
    kaissi/devops-kubernetes-desafio-go:latest-386 \
    kaissi/devops-kubernetes-desafio-go:latest-amd64 \
    kaissi/devops-kubernetes-desafio-go:latest-arm64 \
    kaissi/devops-kubernetes-desafio-go:latest-armv7

docker manifest annotate \
    kaissi/devops-kubernetes-desafio-go:latest \
    kaissi/devops-kubernetes-desafio-go:latest-386 --arch 386 --os linux

docker manifest annotate \
    kaissi/devops-kubernetes-desafio-go:latest \
    kaissi/devops-kubernetes-desafio-go:latest-amd64 --arch amd64 --os linux

docker manifest annotate \
    kaissi/devops-kubernetes-desafio-go:latest \
    kaissi/devops-kubernetes-desafio-go:latest-arm64 --arch arm64 --variant v8 --os linux

docker manifest annotate \
    kaissi/devops-kubernetes-desafio-go:latest \
    kaissi/devops-kubernetes-desafio-go:latest-armv7 --arch arm --variant v7 --os linux
```

Agora, para verificar o manifesto gerado, execute:  
```bash
docker manifest inspect kaissi/devops-kubernetes-desafio-go:latest
```

Para enviar o manifest para o Docker Hub:
```bash
docker manifest push kaissi/devops-kubernetes-desafio-go:latest
```

Para executar todos os passos para múltiplas plataformas - gerar o build de todas as plataformas, fazer o push para o Docker Hub de todas as imagens, e gerar e fazer o push do manifest para estas plataformas, executar:  
`make release`

[buildx]: <https://docs.docker.com/buildx/working-with-buildx/>