package solid

import "fmt"

type Journal struct {
	entries []string
}

func (j *Journal) Add(record string) {
	j.entries = append(j.entries, record)
}

func (j *Journal) Remove(record string) {
	// remove record
	fmt.Println("record removed: ", record)
}

func (j *Journal) Save(record string) {
	// Now save to can be file, database
	// So this resposiblilty should be seprated in another package
	fmt.Println("Save: ", record)
}

func Srp() {
	j := &Journal{}
	j.Add("new data")
}
