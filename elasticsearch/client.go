package elasticsearch

import (
	"encoding/json"
	"errors"

	es "github.com/elastic/go-elasticsearch/v7"
)

var client *es.Client

// 定义 es 操作

// {
// 	"name" : "elasticsearch-master-0",
// 	"cluster_name" : "elasticsearch",
// 	"cluster_uuid" : "9QUDdTSrQuOdbdacfIdUNA",
// 	"version" : {
// 	  "number" : "7.17.1",
// 	  "build_flavor" : "default",
// 	  "build_type" : "docker",
// 	  "build_hash" : "e5acb99f822233d62d6444ce45a4543dc1c8059a",
// 	  "build_date" : "2022-02-23T22:20:54.153567231Z",
// 	  "build_snapshot" : false,
// 	  "lucene_version" : "8.11.1",
// 	  "minimum_wire_compatibility_version" : "6.8.0",
// 	  "minimum_index_compatibility_version" : "6.0.0-beta1"
// 	},
// 	"tagline" : "You Know, for Search"
//   }

type ESInfo struct {
	Name        string    `json:"name"`
	ClusterName string    `json:"cluster_name"`
	ClusterUUID string    `json:""cluster_uuid`
	Version     ESVersion `json:"version"`
	Tagline     string    `json:"tagline"`
}

type ESVersion struct {
	Number                           string `json:"number"`
	BuildFlavor                      string `json:"build_flavor"`
	BuildType                        string `json:"build_type"`
	BuildHash                        string `json:"build_hash"`
	BuildDate                        string `json:"build_date"`
	BuildSnapshot                    bool   `json:"build_snapshot"`
	LuceneVersion                    string `json:"lucene_version"`
	MinimumWireCompatibilityVersion  string `json:"minimum_wire_compatibility_version"`
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
}

func (i ESInfo) Summary() string {
	return "ElasticSearch Cluster: " + i.ClusterName + ", Version: " + i.Version.Number
}

func Info() (*ESInfo, error) {
	resp, err := client.Info()
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	// Check response status
	if resp.IsError() {
		return nil, errors.New("Error: " + resp.String())
	}
	var r ESInfo
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, errors.New("Error parsing the response body: " + err.Error())
	}
	return &r, nil
}

// 初始化 es client
func Init(c *es.Client) {
	client = c
}
