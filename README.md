# Projeto

> go-sqs-worker

## Indíce

- [Informação](#informação)
- [Tecnologias](#tecnologias)
- [Instalação](#instalação)
- [Ambiente](#ambiente)
- [Arquitetura de pastas](#arquitetura-de-pastas)
- [Testes](#testes)

## Informação

Este worker é uma POC realizada para avaliar a performance e usabilidade da linguagem Golang e o sistema de mensageria SQS. O projeto utiliza arquitetura CLEAN e conexão com banco de dados MYSQL.

## Tecnologias

- [GoLang](https://golang.org/) - compilador da linguagem Go
- [Go Mod](https://github.com/golang/mod) - gerenciador de dependencias
- [Docker](https://hub.docker.com/) - gerador e manipulador de containers
- [Docker Compose](https://docs.docker.com/compose/install/) - ferramenta de definição e compartilhamento com containers

## Instalação

Clonando o projeto

```bash
cd $PROJECT_HOME
git clone https://github.com/SamuelVasconc/go-sqs-worker.git
```

Instalando dependências

```
$ go get
```

Removendo dependencias indesejadas

```bash
$ go mod tidy
```

Baixando as dependencias para a vendor local

```bash
$ go mod vendor
```

## Ambiente

Configurando as variáveis de ambiente

| Nome              | Descrição                                       | Valor Padrão | Obrigatório        |
| ----------------- | ----------------------------------------------- | ------------ | ------------------ |
| DBHOST            | Endereço do banco a ser acessado                |              | :white_check_mark: |
| DBNAME            | Nome do banco a ser acessado                    |              | :white_check_mark: |
| DBPASSWORD        | Senha do banco a ser acessado                   |              | :white_check_mark: |
| DBPORT            | Porta do banco a ser acessado                   |              | :white_check_mark: |
| DBUSER            | Usuário do banco a ser acessado                 |              | :white_check_mark: |
| CONNMAXLIFETIME   | Tempo de vida da conexão com o banco            |              | :white_check_mark: |
| MAXIDLECONNS      | Quantidade máxima de conexões ociosas           |              | :white_check_mark: |
| MAXOPENCONNS      | Quantidade máxima de conexões abertas           |              | :white_check_mark: |

## Arquitetura de pastas

### Diretórios

```bash
go-sqs-worker
       |-- cmd
           |-- worker.go
       |-- config
           |-- db
       |-- interfaces
       |-- migrations
       |-- models
       |-- repositories
           |-- mocks
       |-- usecases
       |-- utils
       |-- .gitignore
       |-- README.md
```

#### cmd

Está camada trata os arquivos de orquestração da API/Worker.

#### server.go

Aqui está o orquestrador do server, o arquivo principal que apenas chama as outras camadas.

#### config

Está camada trata as configurações gerais do sistema.

#### db

Está camada trata as conexões com o banco de dados.

#### interfaces

Está camada terá todos os contratos definidos nas interfaces de usecases e repositories.

#### usecases/mocks e repositories/mocks

Reúne todos os artefatos que geram algum mock para o sistema.

#### models

Está camada vai armazenar qualquer object struct. Exemplo: Cliente, Estudante, Livro.

#### repositories

Repository vai armazenar qualquer manipulador de banco de dados ou até mesmo chamado HTTP para outros serviços.

#### utils

Reúne utilitários para auxiliar nos processos comuns aos testes ou configurações do mesmo.

## Iniciando

Gerando container Mysql

```bash
# execute o comando abaixo para gerar o container apartir do arquivo docker-compose.yml na aplicação
$ sudo docker-compose up -d

# OU

# execute o comando abaixo para gerar o container apartir do docker
$ docker run --name container_mysql -e "MYSQL_PASSWORD=1234" -p 15432:5432  -d mysql
```

Buildando o projeto

```bash
# execute o comando abaixo para buildar a aplicação e garantir que nada está quebrado
$ go build
```

Executando o projeto

```bash
$ go run main.go or ./go-sqs-worker
```

## Testes

```bash
# Para execução dos testes automatizados executar o comando abaixo no terminal dentro da pasta da aplicação
$ go test -v -cover ./...

# Para gerar a interface mostrando todos os arquivos e as linhas "Covered", "Not Covered" e "Not Tracked":
$ go test ./... -coverprofile fmtcoverage.html fmt
$ go test ./... -coverprofile cover.out
$ go tool cover -html=cover.out -o cover.html
$ open 'cover.html' file
```
