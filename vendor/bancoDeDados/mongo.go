package bancoDeDados

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"artigos"
	"fmt"
)

//http://denis.papathanasiou.org/posts/2012.10.14.post.html

var (
    mgoSession     *mgo.Session
    DatabaseName = "leagora"
)

func init() {
	var feminismoNegro = artigos.Categoria{Nome: "Feminismo Negro", Prioridade: false, NaoLidos: 0}
	var prioridade = artigos.Categoria{Artigos:[]artigos.Artigo{{Url: "url2", Lido: false}}, Nome:"Prioridade", Prioridade: true, NaoLidos: 0}
	var prioridade2 = artigos.Categoria{Artigos: []artigos.Artigo{{Url: "url1", Lido: true}}, Nome:"Prioridade2", Prioridade: true, NaoLidos: 0}
	
	err3 := AdicionaCategoria(prioridade)
	mostraErro(err3, "Erro no metodo AdicionaCategoria")

	err4 := AdicionaCategoria(prioridade2)
	mostraErro(err4, "Erro no metodo AdicionaCategoria")

	err0 := AdicionaCategoria(feminismoNegro)
	mostraErro(err0, "Erro no metodo AdicionaCategoria")
	
	var artigo1 = artigos.Artigo{Url: "http://algumtexto.com", Lido: false}
	var artigo2 = artigos.Artigo{Url: "http://algumtexto.com", Lido: false}

	err1 := AdicionaArtigo("Feminismo Negro", artigo1)
	mostraErro(err1, "Erro no metodo AdicionaArtigo")

	err2 := AdicionaArtigo("Feminismo Negro", artigo2)
	mostraErro(err2, "Erro no metodo AdicionaArtigo")	
}

func mostraErro(erro string, mensagem string) {
	if erro != "" {
		fmt.Println(mensagem, erro)
	}
}

func getSession () *mgo.Session {
	if mgoSession == nil {
		var err error
		dialInfo, err0 := mgo.ParseURL("mongodb://172.19.0.2:27017")
		if err0 != nil {
			panic(err0)
		}
		dialInfo.Direct = true
		dialInfo.FailFast = true
		mgoSession, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			panic(err)
		}
	
	}
	return mgoSession.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
    session := getSession()
    defer session.Close()
    c := session.DB(DatabaseName).C(collection)
    return s(c)
}

func ObtemCategorias(q interface{}) (resultadosBusca []artigos.Categoria, searchErr string) {
    searchErr     = ""
    resultadosBusca = []artigos.Categoria{}
    query := func(colecao *mgo.Collection) error {
	    fn := colecao.Find(q).All(&resultadosBusca)
        return fn
    }
    search := func() error {
        return withCollection("categoria", query)
    }
    err := search()
    if err != nil {
        searchErr = "Database Error"
    }
    return
}

func AdicionaCategoria(categoria artigos.Categoria) string {
	searchErr := ""
	query := func(colecao *mgo.Collection) error {
		fn := colecao.Insert(&categoria)
		return fn
	}

	adicionaCategoria := func() error {
		 return withCollection("categoria", query)
	}
	err := adicionaCategoria()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
}

func AdicionaArtigo(nomeCategoria string, artigo artigos.Artigo) string {
	searchErr := ""
	categoria, err := ObtemCategorias(bson.M{"nome": nomeCategoria})
	if err != "" {
		return err
	}
	
	if len(categoria) != 0 {
		searchErr = adicionaArtigoNoBanco(categoria[0], artigo)
	} else {
		searchErr = "AdicionaArtigo - Categoria " + nomeCategoria + " não encontrada"
	}
	return searchErr
}

func adicionaArtigoNoBanco(categoria artigos.Categoria, artigo artigos.Artigo) string {
	query := func(colecao *mgo.Collection) error {
		fn := colecao.Update(bson.M{"_id": categoria.ID}, bson.M{"$push": bson.M{"artigos": artigo}})
		return fn
	}

	adicionaArtigo := func() error {
		return withCollection("categoria", query)
	}
	err1 := adicionaArtigo()
	if err1 != nil {
		return "Database Error"
	}
	return ""
}

func LimpaBanco() string {
	searchErr := ""
	query := func(colecao *mgo.Collection) error {
		fn := colecao.Remove(nil)
		return fn
	}

	limpa := func() error {
		return withCollection("categoria", query)
	}
	err := limpa()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
}
