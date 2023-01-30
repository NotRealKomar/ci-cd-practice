package request

import "golang.org/x/exp/slices"

var store []Request = make([]Request, 0)

func GetMany() []Request {
	return store[:]
}

func GetOne(id string) *Request {
	foundIndex := slices.IndexFunc(store, func(element Request) bool {
		return element.Id == id
	})

	if foundIndex != -1 {
		return &store[foundIndex]
	}

	return nil
}

func UpdateOne(id string, title string) bool {
	request := GetOne(id)

	if request == nil {
		return false
	}

	request.Title = title

	return true
}

func Insert(title string) bool {
	store = append(store, create(title))

	return true
}

func DeleteOne(id string) bool {
	foundIndex := slices.IndexFunc(store, func(element Request) bool {
		return element.Id == id
	})

	if foundIndex != -1 {
		store = append(store[:foundIndex], store[foundIndex+1:]...)

		return true
	}

	return false
}
