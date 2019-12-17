package stack

func swapInterface(addr *interface{}, new interface{}) interface{}

func casInterface(addr *interface{}, old, new interface{}) bool

func loadInterface(addr *interface{}) interface{}

func storeInterface(addr *interface{}, new interface{})
