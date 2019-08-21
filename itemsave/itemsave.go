package itemsave

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"reptile/enginer"
	"errors"
)

//存储信息
const host = ""//esurl
func Saveperfile() (chan enginer.Perfileitem,error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		return nil, err
	}
	out := make(chan enginer.Perfileitem)
	go func() {
		itemcount := 0
		for {
			result := <-out
			log.Printf("go item to #%d,%s", itemcount, result)
			itemcount++
			err := save(result, client)
			if err != nil {
				fmt.Println("save to err=", err)
			}
		}
	}()

	return out,nil

}
func save(item enginer.Perfileitem, client *elastic.Client) (err error) {
	if item.Type == "" {
		return errors.New("must supply type")
	}
	_, err = client.Index().
		Index("profile_item").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
