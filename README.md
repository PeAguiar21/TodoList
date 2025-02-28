# ToDoList Backend

Este é o backend de um aplicativo de lista de tarefas (ToDoList) desenvolvido em Golang e MongoDB. O projeto fornece uma API RESTful para gerenciar tarefas, permitindo a criação, leitura, atualização e exclusão de itens (CRUD).

## Funcionalidades

- **Criação de Tarefas**: Adicionar novas tarefas à lista.
- **Leitura de Tarefas**: Listar todas as tarefas ou buscar uma tarefa específica por ID.
- **Atualização de Tarefas**: Atualizar o status ou detalhes de uma tarefa existente.
- **Exclusão de Tarefas**: Remover uma tarefa da lista.
- **Validação de Dados**: Validação básica dos dados recebidos nas requisições.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação utilizada para desenvolver o backend.
- **MongoDB**: Banco de dados NoSQL para armazenamento das tarefas.
- **Gin**: Biblioteca para roteamento e manipulação de requisições HTTP.
- **MongoDB Go Driver**: Driver oficial para conectar e interagir com o MongoDB a partir do Golang.
- **Godotenv**: Utilização de variáveis de ambiente para configurações sensíveis, como a conexão com o banco de dados.

## Pré-requisitos

- Go 1.16 ou superior
- MongoDB 4.4 ou superior
- Git (opcional, para clonar o repositório)

## Como Executar o Projeto

1. **Clone o repositório** (se necessário):

   ```bash
   git clone https://github.com/seu-usuario/todolist-backend.git
   cd todolist-backend
