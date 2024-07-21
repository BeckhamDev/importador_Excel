# Projeto Go: Importação de Dados do Excel para Banco de Dados

Este projeto se trata de uma API desenvolvida usando a linguagem Golang e tem como principal função salvar os dados de uma arquivo excel, com formato e colunas previamente conhecidos, em um banco de dados PostgreSQL. Além disso a API conta com uma lógica de autenticação e criação de usuários, juntamente com criação de tokens JWT para segurança das requisições.

## Funcionalidades

- Importação de dados de arquivos Excel (.xlsx).
- Salvamento dos dados em um banco de dados PostgreSQL.
- Login e criação de usuário.
- Buscas personalizadas.

## Requisitos

- Go 1.16 ou superior
- PostgreSQL

## Execução do Projeto

### Banco de Dados

1. Os dados a respeito da conexão com o banco de dados a ser usado no teste devem ser inseridos dentro do arquivo ".env", o projeto conta com um arquivo env.exemple que contém as variáveis que devem ser usadas. Vale ressaltar que o arquivo ".env" deve ser criado dentro do diretório `/cmd/api/`.

2. Após preencher as variáveis os scripts presentes no arquivo `/sql/reconfile.sql` devem ser executados para a criação das tabelas necessárias para a execução do projeto.

### Executando a API

1. Antes de executar a API ainda na raiz do projeto é necessário que se execute o comando `go mod tidy` para garantir que todas as dependencias necessárias sejam instaladas corretamente no projeto.

2. Para executar a API basta se encaminhar para o diretório `/cmd/api/` usando o terminal e executar o comando `go run main.go`.

## Rotas Desprotegidas (Não precisam de um token):

- **POST /createUser:** Rota de criação de usuário. A rota precisa que no corpo de sua requisição sejam passados os seguintes parâmetros: `name, email, password`

- **POST /login:** Rota de autenticação de usuário que irá devolver um token JWT que deve ser usado para realizar requisições nas demais rotas.
  A rota precisa que no corpo de sua requisição sejam passados os seguintes parâmetros: `email, password`

## Rotas Protegidas (É necessário que um token JWT válido seja passado como Bearer Token no momento da requisição):

- **POST /app/saveData:** Rota que importa os dados do excel para o banco de dados. É necessário que um arquivo no formato ".xlsx" seja passado na requisição para que seus dados sejam salvos.

- **GET /app/customerByName/{customerName}:** Retorna todos os dados presentes no banco a partir de um Customer Name.

- **GET /app/customersByMeterCategory/{meterCategory}:** Retorna todos os dados presentes no banco a partir de uma Meter Category.

- **GET /app/customersByMeterRegion/{meterRegion}:** Retorna todos os dados presentes no banco a partir de uma Meter Region.

- **GET /app/customersByResourceGroup/{resourceGroup}:** Retorna todos os dados presentes no banco a partir de um Resource Group.
