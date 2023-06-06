package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/qmuntal/gltf"
)

type VRMJSONv0 struct {
	ExporterVersion    string          `json:"exporterVersion"`
	SpecVersion        string          `json:"specVersion"`
	Meta               json.RawMessage `json:"meta"`
	humanoid           json.RawMessage `json:"humanoid"`
	firstPerson        json.RawMessage `json:"firstPerson"`
	blendShapeMaster   json.RawMessage `json:"blendShapeMaster"`
	secondaryAnimation json.RawMessage `json:"secondaryAnimation"`
	materialProperties json.RawMessage `json:"materialProperties"`
}
type VRMJSON struct {
	SpecVersion string          `json:"specVersion"`
	Meta        json.RawMessage `json:"meta"`
}

var rVRMFilename *regexp.Regexp = regexp.MustCompile(`(?i)(.+).vrm`)

func init() {
	gltf.RegisterExtension("VRM", VRMv0Unmarshal)
	gltf.RegisterExtension("VRMC_vrm", VRMv1Unmarshal)
}

func VRMv0Unmarshal(data []byte) (interface{}, error) {
	foo := new(VRMJSONv0)
	err := json.Unmarshal(data, foo)
	return foo, err
}
func VRMv1Unmarshal(data []byte) (interface{}, error) {
	foo := new(VRMJSON)
	err := json.Unmarshal(data, foo)
	return foo, err
}
func main() {
	fmt.Println("VRM Version Viwer")

	if len(os.Args) != 2 {
		fmt.Println("missing args")
		os.Exit(1)
	}

	var vrm_filename string = os.Args[1]
	fmt.Printf("VRM filename: %s\n", vrm_filename)

	if !rVRMFilename.MatchString(vrm_filename) {
		fmt.Println("missing VRM file(name)")
		os.Exit(1)
	}

	doc, err := gltf.Open(vrm_filename)
	if err != nil {
		panic(err)
	}
	extensions := doc.Extensions
	var keys []string
	for key := range extensions {
		keys = append(keys, key)
	}
	fmt.Printf("extensions %v\n", keys)

	_, okv0 := extensions["VRM"]
	if okv0 {
		vrm_data := extensions["VRM"].(*VRMJSONv0)
		// fmt.Printf("%+v\n", vrm_data)
		// fmt.Printf("type: %T\n", vrm_data)

		fmt.Printf("VRM Version: %+s\n", vrm_data.SpecVersion)
		fmt.Printf("VRM exporterVersion: %s\n", vrm_data.ExporterVersion)

	}

	_, okv1 := extensions["VRMC_vrm"]
	if okv1 {
		vrm_data := extensions["VRMC_vrm"].(*VRMJSON)
		fmt.Printf("VRM Version: %+s\n", vrm_data.SpecVersion)
		// fmt.Printf("type: %T\n", vrm_data)

	}
	fmt.Println()
}
