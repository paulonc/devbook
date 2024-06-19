# DevBook

## 🚀 Descrição

DevBook é uma rede social simples desenvolvida utilizando Go (Golang) para o backend e PostgreSQL como banco de dados. A aplicação permite que os usuários se cadastrem, façam login, postem mensagens, sigam e deixem de seguir outros usuários, curtam postagens, e vejam publicações próprias e de quem seguem. A configuração e execução da aplicação são facilitadas pelo uso de Docker e Makefile.

## 🛠️ Tecnologias Utilizadas
- **Backend**:
    - Go
    - PostgreSQL
    - Migrate: Utilizado para gerenciamento de migrações de banco de dados
    - JWT (JSON Web Tokens): Utilizado para autenticação e autorização

- **Containerização e Automação**:
    - Docker: Para containerização da aplicação
    - Docker Compose: Para orquestração dos containers
    - Makefile: Para automação de tarefas comuns

Este projeto serve como um exemplo simples para compreender como essas tecnologias podem ser combinadas para criar uma aplicação web funcional.

## 📋 Funcionalidades
- **Cadastro e Login de Usuários**: Permite que novos usuários se cadastrem e usuários existentes façam login.
- **Postagem de Mensagens**: Usuários podem postar mensagens para compartilhar com seus seguidores.
- **Seguir e Deixar de Seguir**: Possibilidade de seguir e deixar de seguir outros usuários.
- **Curtir Postagens**: Dar like nas postagens de outros usuários.
- **Visualizar Publicações**: Veja suas próprias publicações e as das pessoas que você segue.

## 🔗 Principais Endpoints
- **Cadastro de Usuário**: `POST /users`
- **Login de Usuário**: `POST /login`
- **Postar Mensagem**: `POST /publications`
- **Seguir Usuário**: `POST /users/{id}/follow`
- **Deixar de Seguir Usuário**: `POST /users/{id}/unfollow`
- **Curtir Postagem**: `POST /publications/{publicationId}/like`
- **Ver Publicações**: `GET /publications`

## 📝 Licença
Este projeto está licenciado sob a [MIT License](LICENSE).

---
## 👥 By
- [Paulo Victor](https://github.com/paulonc)