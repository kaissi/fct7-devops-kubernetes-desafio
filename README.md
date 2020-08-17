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