![Golang](http://www.qureet.com/blog/wp-content/uploads/2013/11/jumbo_gopher-4bf98fbc72cc188289ba2b458d4ce680.png)
# O que lê agora

Salve os links que quer ler por categoria/tópico/etc.

Quais categorias você quer ler primeiro.

## Contribuindo com o projeto

### O projeto usa

- Node

- Golang: Macaron, goquery, gin

1. Construa a imagem: `docker build -t o-q-le-agora .`

Execute o programa: `docker run --rm -p 4000:4000 -v "${PWD}":/go/src/app -w /go/src/app -it o-q-le-agora gin -p 4000 run main.go`

Executar testes: `docker run --rm -v "${PWD}":/go/src/app -w /go/src/app -it o-q-le-agora go test`
