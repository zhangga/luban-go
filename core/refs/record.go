package refs

type Record struct {
	AutoIndex int
	Data      interface{} //DBean
	Source    string
	Tags      []string
}
