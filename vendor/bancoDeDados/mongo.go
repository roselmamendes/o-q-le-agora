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

func mostraErro(erro string, mensagem string) {
	if erro != "" {
		fmt.Println(mensagem, erro)
	}
}

func ConfiguraBanco(connectionString string) {
	clientOptions := options.Client().ApplyURI(connectionString)
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
