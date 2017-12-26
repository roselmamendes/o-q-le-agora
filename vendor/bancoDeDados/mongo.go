package bancoDeDados

import (
	"gopkg.in/mgo.v2"
	"artigos"
)

//http://denis.papathanasiou.org/posts/2012.10.14.post.html

var (
    mgoSession     *mgo.Session
    databaseName = "leagora"
)

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
    c := session.DB(databaseName).C(collection)
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

	adiciona := func() error {
		 return withCollection("categoria", query)
	}
	err := adiciona()
	if err != nil {
		searchErr = "Database Error"
	}
	return searchErr
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
