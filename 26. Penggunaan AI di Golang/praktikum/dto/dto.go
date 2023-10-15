package dto

type RequestRecomended struct {
	Budget  string `json:budget`
	Purpose string  `json:purpose`
	Merk    string  `json:merk`
}

type ResponseRecomended struct {
	Status         string
	Recommendation string
}
