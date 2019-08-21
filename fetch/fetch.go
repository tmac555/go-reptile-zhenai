package fetch

import (
	"io/ioutil"
	"net/http"
)

func Fetchrequest(url string)([]byte,error){
	response, e := http.Get(url)
	if e!= nil {
		return nil, e
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)



}
