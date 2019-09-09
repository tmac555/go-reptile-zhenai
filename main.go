package main

import (
	"reptile/enginer"
	"reptile/itemsave"
	"reptile/schedular"
	"reptile/hongniang/parse"
)

func main() {
	saveperfile, err := itemsave.Saveperfile()
	if err!=nil{
		panic(err)
	}
	e:=enginer.Concurrent{
		Schedular:&schedular.Simpleschedular{},
		Workcount:10,
		Itemchan:saveperfile,
	}

	e.Run(enginer.Request{
		Url:"",//address url
		Parsefunc:parse.Parsecitylist,
	  },
	)
}
