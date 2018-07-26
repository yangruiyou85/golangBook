package mock

type Retriever1 struct {
	Contents string
}

func (r Retriever1) Get(url string) string {
	return r.Contents
}
