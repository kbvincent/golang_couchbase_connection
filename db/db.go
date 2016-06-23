package db

import (
	"encoding/json"
	"fmt"
	"github.com/couchbase/gocb"
)

type DataStore interface{
	Counter(key string, delta, initial int64, expiry uint32) (uint64, gocb.Cas, error)
	Get(key string, valuePtr interface{}) (gocb.Cas, error)
	Upsert(key string, value interface{}, expiry uint32) (gocb.Cas, error)
	Replace(key string, value interface{},cas gocb.Cas, expiry uint32) (gocb.Cas, error)
	Insert(key string, value interface{}, expiry uint32) (gocb.Cas, error)
}

type Config struct{
	DataStoreUrl string
	DataStoreBucket string
	DataStoreBucketPW string
}

func GetConfig() Config{
	var c Config

	c = Config{
		DataStoreUrl:"couchbase://10.84.101.122/",
		DataStoreBucket:"kvincent",
		DataStoreBucketPW:"",
	}

	fmt.Println(c, "config Loaded")
	return c
}


type CouchbaseTranscoder struct{}
var conf = GetConfig()

func (t CouchbaseTranscoder) Decode(bytes []byte, flags uint32, out interface{}) error {
	err := json.Unmarshal(bytes, &out)
	if err != nil {
		return err
	}
	return nil
}

func (t CouchbaseTranscoder) Encode(value interface{}) ([]byte, uint32, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, 0, err
	}
	return bytes, 0, nil
}

func ConnectDb() *gocb.Bucket {

	myCluster, err := gocb.Connect(conf.DataStoreUrl)
	if err != nil{
		panic(fmt.Errorf("No DB connection - can't connect dieing now", err))
	}

	bucket, err := myCluster.OpenBucket(conf.DataStoreBucket, conf.DataStoreBucketPW)
	if err != nil {
		panic(fmt.Errorf("Couldnt connect to specific bucket %s \n", err))
		return nil
	}

	bucket.SetTranscoder(CouchbaseTranscoder{})
	fmt.Println("db connected")
	return bucket

}

