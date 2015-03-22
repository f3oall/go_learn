package main

import "net/http"

//Route  structure contains fields wich define route.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a slice of Route structures
type Routes []Route

var routes = Routes{
	//Clients
	Route{
		"ClientIndex",
		"GET",
		"/clients",
		index(&allCls),
	},
	Route{
		"ClientShow",
		"GET",
		"/clients/{recordID}",
		show(&allCls),
	},
	Route{
		"ClientCreate",
		"POST",
		"/clients",
		create(&allCls),
	},
	Route{
		"ClientEdit",
		"PUT",
		"/clients/{recordID}",
		update(&allCls),
	},
	Route{
		"ClientDelete",
		"DELETE",
		"/clients/{recordID}",
		delete(&allCls),
	}, /*
		//Products
		Route{
			"ProductIndex",
			"GET",
			"/products",
			index(&allPrds),
		},
		Route{
			"ProductShow",
			"GET",
			"/products/{recordID}",
			show(&allPrds),
		},
		Route{
			"ProductCreate",
			"POST",
			"/products",
			create(&allPrds),
		},
		Route{
			"ProductEdit",
			"PUT",
			"/products/{recordID}",
			update(&allPrds),
		},
		Route{
			"ProductDelete",
			"DELETE",
			"/products/{recordID}",
			delete(&allPrds),
		},
		//Orders
		Route{
			"OrderIndex",
			"GET",
			"/orders",
			index(&allOrds),
		},
		Route{
			"OrderShow",
			"GET",
			"/orders/{recordID}",
			show(&allOrds),
		},
		Route{
			"OrderCreate",
			"POST",
			"/orders",
			create(&allOrds),
		},
		Route{
			"OrderEdit",
			"PUT",
			"/orders/{recordID}",
			update(&allOrds),
		},
		Route{
			"OrderDelete",
			"DELETE",
			"/orders/{recordID}",
			delete(&allOrds),
		},*/
}
