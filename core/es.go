package core

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"server/global"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	//host := "http://127.0.0.1:9200

	client, err := elastic.NewClient(
		elastic.SetURL(global.Config.Es.Url()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.Es.User, global.Config.Es.Password),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	return client
}
