package collector

type Formatter struct{}

type HasFormatter interface {}

func (* Formatter) format(){}
