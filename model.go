package main

type pulls struct {
	title     string
	user      string
	state     string
	createdAt string
}

type convertedPulls struct {
	title     string
	user      string
	state     string
	createdAt string
}

func (*pulls) convertPulls() convertedPulls {
	return convertedPulls{}
}
