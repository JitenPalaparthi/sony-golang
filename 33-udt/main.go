package main

func main() {
	type RS = struct {
		count   int
		errors  []string
		limit   int
		offset  int
		records []struct {
			id   int
			name string
		}
	}

	//var rs1 RS

	// json.Ummarshal(bytes,&rs1)
}

type Result struct {
	count  int
	errors []string
	//records []Record

	records []struct {
		id   int
		name string
	}
}

type Record struct {
	id   int
	name string
}

/*{
	"result":{
		"errors":["no data","issue with data"],
		"count":100
		 "records":[
			{"id":101,
		"name":"item1"},
		{"id":102,
		"name":"item1"},
		{"id":103,
		"name":"item1"},
		{"id":104,
		"name":"item1"}
		 ]
	}
}*/
