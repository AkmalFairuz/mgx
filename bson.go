package mgx

import (
	"go.mongodb.org/mongo-driver/bson"
)

type M bson.M
type A bson.A
type D bson.D
type E bson.E

type toBSON interface {
	bson() (string, any)
}

// Set is a wrapper around mongo's $set operator.
type Set bson.M

func (s Set) bson() (string, any) {
	return "$set", bson.M(s)
}

// Unset is a wrapper around mongo's $unset operator.
type Unset bson.M

func (u Unset) bson() (string, any) {
	return "$unset", bson.M(u)
}

// Inc is a wrapper around mongo's $inc operator.
type Inc bson.M

func (i Inc) bson() (string, any) {
	return "$inc", bson.M(i)
}

func processUpdate(update any) any {
	switch u := update.(type) {
	case toBSON:
		k, v := u.bson()
		return bson.M{k: v}
	case []toBSON:
		m := bson.M{}
		for _, v := range u {
			k, v2 := v.bson()
			m[k] = v2
		}
	case IDImpl:
		return bson.M{u.Field(): u.ID()}
	case bson.M:
		return u
	}
	return update
}

func processFilter(filter any) any {
	switch f := filter.(type) {
	case IDImpl:
		return bson.M{f.Field(): f.ID()}
	}
	return filter
}

// IDImpl is an interface for generating IDs.
// Example case: if you want to generate an auto-incrementing ID, you can implement this interface.
type IDImpl interface {
	GenerateID(collectionName string) (any, error)
	ID() any
	Field() string
}

func processInsert(collectionName string, doc any) (any, error) {
	if m, ok := doc.(M); ok {
		for i, v := range m {
			if id, ok := v.(IDImpl); ok {
				generatedID, err := id.GenerateID(collectionName)
				if err != nil {
					return nil, err
				}
				m[id.Field()] = generatedID
				delete(m, i)
			}
		}
	}
	return doc, nil
}
