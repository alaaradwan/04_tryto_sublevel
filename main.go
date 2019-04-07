package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//---------------------------------------------------------------
// create strunct for object to be saved in db
//---------------------------------------------------------------
type Data struct {
	Id    string
	Name  string
	Group []string
}

//---------------------------------------------------------------
// create function that create bson
//---------------------------------------------------------------
func Martualtobson(Data Data) (value []byte, err error) {
	value, err = json.Marshal(Data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return value, nil
}

func main() {
	//---------------------------------------------------------------
	// create db database or open if exist
	//---------------------------------------------------------------
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		fmt.Println("err ...")
		log.Fatal(err)
	}

	//---------------------------------------------------------------
	// insert into db
	//---------------------------------------------------------------
	values := []Data{
		{Id: "0", Name: "ali", Group: []string{"A", "B"}},
		{Id: "1", Name: "ola", Group: []string{"B", "c"}},
		{Id: "2", Name: "aya", Group: []string{"C", "D"}},
		{Id: "3", Name: "mi", Group: []string{"D", "E"}},
		{Id: "4", Name: "noha", Group: []string{"E", "F"}},
		{Id: "5", Name: "nour", Group: []string{"F", "G"}},
	}
	d, _ := Martualtobson(values[0])
	d1, _ := Martualtobson(values[1])
	d2, _ := Martualtobson(values[2])
	d3, _ := Martualtobson(values[3])
	d4, _ := Martualtobson(values[4])
	d5, _ := Martualtobson(values[5])
	//---------------------------------------------------------------
	// the ket structure is "id-name"
	//---------------------------------------------------------------
	err = db.Put([]byte("0-ali"), d, nil)
	err = db.Put([]byte("1-ola"), d1, nil)
	err = db.Put([]byte("2-aya"), d2, nil)
	err = db.Put([]byte("3-mi"), d3, nil)
	err = db.Put([]byte("4-noha"), d4, nil)
	err = db.Put([]byte("5-nour"), d5, nil)

	//---------------------------------------------------------------
	//get by key
	//---------------------------------------------------------------
	data, err := db.Get([]byte("0-ali"), nil)
	fmt.Println("---------------------------------- ")
	fmt.Println("get by key :- ")
	fmt.Println("---------------------------------- ")
	fmt.Println(string(data))
	//---------------------------------------------------------------
	//iterate (get all)
	//---------------------------------------------------------------
	fmt.Println("---------------------------------- ")
	fmt.Println("get all data in database :- ")
	fmt.Println("---------------------------------- ")
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Println("your key :- " + string(key) + " || your data :- " + string(value))

	}
	//---------------------------------------------------------------
	//iterate by id (get by id)
	//---------------------------------------------------------------
	//iter := db.NewIterator(&util.Range{Start: []byte("key"), Limit: []byte("key4")}, nil) //use this to limit the number of data
	fmt.Println("---------------------------------- ")
	fmt.Println("get all data in database by id :- ")
	fmt.Println("---------------------------------- ")
	iter = db.NewIterator(util.BytesPrefix([]byte("0")), nil) // get all key starts by k
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Println("your key :- " + string(key) + " || your data :- " + string(value))

	}
	//---------------------------------------------------------------
	//update db by key
	//---------------------------------------------------------------
	fmt.Println("---------------------------------- ")
	fmt.Println("update by key  :- update <0-ali> key put him in group F and A ")
	fmt.Println("---------------------------------- ")
	newdata := Data{Id: "0", Name: "ali", Group: []string{"F", "A"}}
	newjsion, _ := Martualtobson(newdata)
	err = db.Put([]byte("0-ali"), newjsion, nil)
	//get the updated opject
	data, err = db.Get([]byte("0-ali"), nil)
	fmt.Println("get by key :- " + string(data))

	//----------------------------------------------------------------------------
	// delete by key from db
	//----------------------------------------------------------------------------
	fmt.Println("---------------------------------- ")
	fmt.Println("delete by key  :- delete <1-ola> ")
	fmt.Println("---------------------------------- ")
	err = db.Delete([]byte("1-ola"), nil)
	if err != nil {
		fmt.Println("err..")
		log.Fatal(err)
	} else {
		fmt.Println("deleted >>>")
	}
	// test if the opject sill exists
	fmt.Println("find 1-ola key ... ")
	reobj, err := db.Get([]byte("1-ola"), nil) // get the data inserted
	if err != nil {
		fmt.Println("err.. Not founded")
		log.Fatal(err)
	} else {
		fmt.Println(string(reobj))
	}
	//---------------------------------------------------------------
	//release the iterator
	//---------------------------------------------------------------

	iter.Release()
	err = iter.Error()

	db.Close()
}
