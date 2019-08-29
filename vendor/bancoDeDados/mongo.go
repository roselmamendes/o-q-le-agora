package bancoDeDados

import (
	"context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
	"artigos"
)

var db     *mongo.Database

func init() {
	GetDatabase()
// 	var feminismoNegro = artigos.Categoria{Nome: "Feminismo Negro", Prioridade: false, NaoLidos: 0}
// 	var prioridade = artigos.Categoria{Artigos:[]artigos.Artigo{{Url: "url2", Lido: false}}, Nome:"Prioridade", Prioridade: true, NaoLidos: 0}
// 	var prioridade2 = artigos.Categoria{Artigos: []artigos.Artigo{{Url: "url1", Lido: true}}, Nome:"Prioridade2", Prioridade: true, NaoLidos: 0}
	
// 	err3 := AdicionaCategoria(prioridade)
// 	mostraErro(err3, "Erro no metodo AdicionaCategoria")

// 	err4 := AdicionaCategoria(prioridade2)
// 	mostraErro(err4, "Erro no metodo AdicionaCategoria")

// 	err0 := AdicionaCategoria(feminismoNegro)
// 	mostraErro(err0, "Erro no metodo AdicionaCategoria")
	
// 	var artigo1 = artigos.Artigo{Url: "http://algumtexto.com", Lido: false}
// 	var artigo2 = artigos.Artigo{Url: "http://algumtexto.com", Lido: false}

// 	err1 := AdicionaArtigo("Feminismo Negro", artigo1)
// 	mostraErro(err1, "Erro no metodo AdicionaArtigo")

// 	err2 := AdicionaArtigo("Feminismo Negro", artigo2)
// 	mostraErro(err2, "Erro no metodo AdicionaArtigo")	
}

func mostraErro(erro string, mensagem string) {
	if erro != "" {
		fmt.Println(mensagem, erro)
	}
}

func GetDatabase() {
		clientOptions := options.Client().ApplyURI("mongodb://db:27017")
	    client, err := mongo.Connect(context.TODO(), clientOptions)

	    if err != nil {
	        log.Fatal(err)
	    }

	    err = client.Ping(context.TODO(), nil)

	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Println("Connected to MongoDB!")

	    db = client.Database("leagora")
}

func ObtemCategorias(q interface{}) (resultadosBusca []artigos.Categoria, searchErr string) {
    searchErr     = ""
    resultadosBusca = []artigos.Categoria{}

    collection := db.Collection("categoria")

    // Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
	    log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
	    
	    // create a value into which the single document can be decoded
	    var elem artigos.Categoria
	    err := cur.Decode(&elem)
	    if err != nil {
	        log.Fatal(err)
	    }

	    resultadosBusca = append(resultadosBusca, elem)
	}

	if err := cur.Err(); err != nil {
	   log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Categorias encontradas): %+v\n", resultadosBusca)
	return resultadosBusca, searchErr
}

func AdicionarCategoria(categoria artigos.Categoria)  string {
	searchErr := ""
	collection := db.Collection("categoria")
	categoriaAdicionada, err := collection.InsertOne(context.TODO(), categoria)
    if err != nil {
        searchErr = "Database Error"
    }
    fmt.Println("Inserted a Single Document: ", categoriaAdicionada.InsertedID)

    return searchErr

}

// func AdicionaArtigo(nomeCategoria string, artigo artigos.Artigo) string {
// 	searchErr := ""
// 	categoria, err := ObtemCategorias(bson.M{"nome": nomeCategoria})
// 	if err != "" {
// 		return err
// 	}
	
// 	if len(categoria) != 0 {
// 		searchErr = adicionaArtigoNoBanco(categoria[0], artigo)
// 	} else {
// 		searchErr = "AdicionaArtigo - Categoria " + nomeCategoria + " não encontrada"
// 	}
// 	return searchErr
// }

// func adicionaArtigoNoBanco(categoria artigos.Categoria, artigo artigos.Artigo) string {
// 	query := func(colecao *mgo.Collection) error {
// 		fn := colecao.Update(bson.M{"_id": categoria.ID}, bson.M{"$push": bson.M{"artigos": artigo}})
// 		return fn
// 	}

// 	adicionaArtigo := func() error {
// 		return withCollection("categoria", query)
// 	}
// 	err1 := adicionaArtigo()
// 	if err1 != nil {
// 		return "Database Error"
// 	}
// 	return ""
// }

func LimpaBanco() string {
	searchErr := ""
	collection := db.Collection("categoria")
	collection.DeleteMany(context.TODO(), bson.D{})
	return searchErr
}

// 	limpa := func() error {
// 		return withCollection("categoria", query)
// 	}
// 	err := limpa()
// 	if err != nil {
// 		searchErr = "Database Error"
// 	}
// 	return searchErr
// }
