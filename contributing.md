## Contribuindo com o projeto

### O projeto usa

- Docker, Docker Compose

- Golang: Macaron, goquery, gin, mgo


Execute o programa: `docker-compose up --build`

Acesse o projeto em 'localhost:4000'.

Executar testes: `docker-compose run -e CONNECTIONSTRING=mongodb://db:27017 web go test`