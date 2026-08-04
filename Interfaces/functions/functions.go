package functions

import "sync"

type MyFunctions interface {
	Create() *Data
	UpdateName(name string)
	AddData(d Data)
	GetData() Data
	UpdateAge(age int)
}

type Data struct {
	Name string
	ID   string
	Age  int
}

type DataStore struct {
	mu   sync.Mutex
	data Data
}

func (t *DataStore) Create() *Data {
	return &Data{}
}

func (t *DataStore) AddData(d Data) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.data.Name = d.Name
	t.data.ID = d.ID
	t.data.Age = d.Age
}

func (t *DataStore) UpdateName(name string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.data.Name = name
}

func (t *DataStore) UpdateAge(age int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.data.Age = age
}

func (t *DataStore) GetData() Data {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.data
}