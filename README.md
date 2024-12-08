# Protótipo de Rede Social - Devbook

Este é um protótipo de rede social desenvolvido em Go, onde os usuários podem criar posts com texto. O backend do projeto é implementado em Go e o banco de dados é executado dentro de um container Docker. O projeto está configurado para facilitar o desenvolvimento de uma aplicação completa com frontend e backend no mesmo repositório.

## Funcionalidades

- Criação de posts com texto.
- Armazenamento de posts em um banco de dados no Docker.

## Tecnologias Utilizadas

- **Go** (Golang) para a construção do backend.
- **Docker** para containerização do banco de dados.
- **PostgreSQL** (ou outro banco de dados, dependendo da configuração do seu projeto) para o armazenamento dos dados.

## Pré-requisitos

Antes de rodar o projeto, certifique-se de que você tem os seguintes pré-requisitos instalados:

- [Go](https://golang.org/doc/install) (versão 1.23 ou superior).
- [Docker](https://www.docker.com/get-started) para rodar o banco de dados no container.\

## Como Rodar o Projeto

### 1. Clone o Repositório

Primeiro, clone este repositório para a sua máquina local:

```bash
git clone https://github.com/Dan0Silva/devbook_application.git
cd devbook_application
```

### 2. Inicie o Banco de Dados com Docker

Aqui voce irá precisar rodar por conta própria, pois ainda não implementei uma solução facilitada, então execute os seguintes comandos:

```bash
docker pull mysql
docker run -e MYSQL_ROOT_PASSWORD=root --name devbook-database -d mysql
```

Isso irá criar um container de nome `devbook-database` a partir da imagem mysql. Por padrão, ele vem com o usuário root:root.
Você pode optar por criar um usuário novo com permissões apenas para o banco em especifico, ou pode continuar utilizando o usuário root. A configuração do banco e das tabelas estarão na pasta `sql`.

### 3. Instale as Dependências do Go

Dentro do diretório do projeto, instale as dependências do Go executando:

```bash
go mod tidy
```

Isso irá baixar todas as dependências necessárias para o projeto. Faça isso dentro das respectivas pastas de backend e frontend.

### 4. Configure as Variáveis de Ambiente

Certifique-se de configurar as variáveis de ambiente necessárias para a conexão com o banco de dados. Você pode criar um arquivo `.env` no diretório raiz com as configurações:

```bash
DB_USER=root
DB_PASSWORD=root
DB_NAME=database

DB_ADDRESS=0.0.0.0
DB_PORT=3306

API_PORT=5000
```

### 5. Rode o Backend

Para rodar o servidor backend, execute o seguinte comando:

```bash
go run .
```

Isso iniciará o servidor na porta padrão na porta especificada dentro do arquivo `.env`.

### 6. Teste a API

Agora você pode acessar a API que está rodando localmente. Por exemplo, para criar um usuário, você pode enviar uma requisição POST para `http://localhost:5000/users` com o corpo da requisição no formato JSON.

Exemplo de requisição:

```json
{
  "name": "João Silva",
  "nick": "joaosilva",
  "email": "joao.silva@example.com",
  "password": "senha123",
}
```
<!--
## Estrutura do Projeto

A estrutura do projeto é organizada da seguinte forma:

```
/src
  /controller    # Contém os handlers da API
  /model         # Contém os modelos de dados
  /repository    # Interage diretamente com o banco de dados
  /service       # Contém a lógica de negócios
/main.go         # Arquivo principal do projeto, onde o servidor é iniciado
/docker-compose.yml  # Configuração do Docker para o banco de dados
```
-->

## Como Contribuir

Sinta-se à vontade para abrir issues ou enviar pull requests com melhorias e correções de bugs.

1. Faça um fork deste repositório.
2. Crie uma nova branch: `git checkout -b minha-branch`.
3. Faça suas alterações.
4. Commit suas alterações: `git commit -am 'Adiciona nova funcionalidade'`.
5. Envie para o repositório remoto: `git push origin minha-branch`.
6. Crie um pull request.
