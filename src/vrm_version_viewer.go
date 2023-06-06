package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/qmuntal/gltf"
)

type VRMJSONv0 struct {
	ExporterVersion string        `json:"exporterVersion"`
	SpecVersion     string        `json:"specVersion"`
	Meta            VRMJSONv0Meta `json:"meta"`
	// humanoid           json.RawMessage `json:"humanoid"`
	// firstPerson        json.RawMessage `json:"firstPerson"`
	// blendShapeMaster   json.RawMessage `json:"blendShapeMaster"`
	// secondaryAnimation json.RawMessage `json:"secondaryAnimation"`
	// materialProperties json.RawMessage `json:"materialProperties"`
}

type VRMJSONv0Meta struct {
	Title                string `json:"title"`
	Version              string `json:"version"`
	Author               string `json:"author"`
	ContactInformation   string `json:"contactInformation"`
	Reference            string `json:"reference"`
	Texture              int    `json:"texture"`
	AllowedUserName      string `json:"allowedUserName"`
	ViolentUssageName    string `json:"violentUssageName"`
	SexualUssageName     string `json:"sexualUssageName"`
	CommercialUssageName string `json:"commercialUssageName"`
	OtherPermissionUrl   string `json:"otherPermissionUrl"`
	LicenseName          string `json:"licenseName"`
	OtherLicenseUrl      string `json:"otherLicenseUrl"`
}

type VRMJSONv1 struct {
	SpecVersion string        `json:"specVersion"`
	Meta        VRMJSONv1Meta `json:"meta"`
	// Humanoid    json.RawMessage `json:"humanoid"`
	// FirstPerson json.RawMessage `json:"firstPerson"`
	// LookAt      json.RawMessage `json:"lookAt"`
	// Expressions json.RawMessage `json:"expressions"`
}

type VRMJSONv1Meta struct {
	Name                           string          `json:"name"`
	Version                        string          `json:"version"`
	Authors                        []string        `json:"authors"`
	CopyrightInformation           string          `json:"copyrightInformation"`
	ContactInformation             string          `json:"contactInformation"`
	References                     []string        `json:"references"`
	ThirdPartyLicenses             string          `json:"thirdPartyLicenses"`
	ThumbnailImage                 json.RawMessage `json:"thumbnailImage"`
	LicenseUrl                     string          `json:"licenseUrl"`
	AvatarPermission               string          `json:"avatarPermission"`
	AllowExcessivelyViolentUsage   bool            `json:"allowExcessivelyViolentUsage"`
	AllowExcessivelySexualUsage    bool            `json:"allowExcessivelySexualUsage"`
	CommercialUsage                string          `json:"commercialUsage"`
	AllowPoliticalOrReligiousUsage bool            `json:"allowPoliticalOrReligiousUsage"`
	AllowAntisocialOrHateUsage     bool            `json:"allowAntisocialOrHateUsage"`
	CreditNotation                 string          `json:"creditNotation"`
	AllowRedistribution            bool            `json:"allowRedistribution"`
	Modification                   string          `json:"modification"`
	OtherLicenseUrl                string          `json:"otherLicenseUrl"`
}

var version string
var rVRMFilename *regexp.Regexp = regexp.MustCompile(`(?i)(.+).vrm`)
const VRMv0ExtensionName = "VRM"
const VRMv1ExtensionName = "VRMC_vrm"

func init() {
	// v0
	gltf.RegisterExtension(VRMv0ExtensionName, VRMv0Unmarshal)
	// v1
	gltf.RegisterExtension(VRMv1ExtensionName, VRMv1Unmarshal)
}

func VRMv0Unmarshal(data []byte) (interface{}, error) {
	foo := new(VRMJSONv0)
	err := json.Unmarshal(data, foo)
	return foo, err
}

func VRMv1Unmarshal(data []byte) (interface{}, error) {
	foo := new(VRMJSONv1)
	err := json.Unmarshal(data, foo)
	return foo, err
}

func main() {
	fmt.Printf("VRM Version Viwer %s\n", version)

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
	asset := doc.Asset
	var keys []string
	for key := range extensions {
		keys = append(keys, key)
	}
	fmt.Printf("GLTF Extensions %v\n", keys)
	// fmt.Printf("GLTF Asset: %+v\n", doc.Asset)

	_, okv0 := extensions[VRMv0ExtensionName]
	if okv0 {
		vrm_data := extensions[VRMv0ExtensionName].(*VRMJSONv0)

		fmt.Printf("GLTF.Asset.Generator: %s\n", asset.Generator)
		fmt.Printf("VRM Version: %+s\n", vrm_data.SpecVersion)
		fmt.Printf("VRM exporterVersion: %s\n", vrm_data.ExporterVersion)

		var vrm_meta VRMJSONv0Meta = vrm_data.Meta
		fmt.Printf("VRM Meta\n")
		fmt.Printf(" - Name: %s\n", vrm_meta.Title)
		fmt.Printf(" - Version: %s\n", vrm_meta.Version)
		fmt.Printf(" - Author: %s\n", vrm_meta.Author)
		fmt.Printf(" - ContactInformation: %s\n", vrm_meta.ContactInformation)
		fmt.Printf(" - Reference: %s\n", vrm_meta.Reference)
		fmt.Printf(" - AllowedUserName: %s\n", vrm_meta.AllowedUserName)
		fmt.Printf(" - ViolentUssageName: %s\n", vrm_meta.ViolentUssageName)
		fmt.Printf(" - SexualUssageName: %s\n", vrm_meta.SexualUssageName)
		fmt.Printf(" - CommercialUssageName: %s\n", vrm_meta.CommercialUssageName)
		fmt.Printf(" - OtherPermissionUrl: %s\n", vrm_meta.OtherPermissionUrl)
		fmt.Printf(" - LicenseName: %s\n", vrm_meta.LicenseName)
		fmt.Printf(" - OtherLicenseUrl: %s\n", vrm_meta.OtherLicenseUrl)

	}

	_, okv1 := extensions[VRMv1ExtensionName]
	if okv1 {
		vrm_data := extensions[VRMv1ExtensionName].(*VRMJSONv1)
		fmt.Printf("GLTF.Asset.Generator: %s\n", asset.Generator)
		fmt.Printf("VRM Version: %+s\n", vrm_data.SpecVersion)
		// fmt.Printf("type: %T\n", vrm_data)

		var vrm_meta VRMJSONv1Meta = vrm_data.Meta
		fmt.Printf("VRM Meta\n")
		fmt.Printf(" - Name: %s\n", vrm_meta.Name)
		fmt.Printf(" - Version: %s\n", vrm_meta.Version)
		fmt.Printf(" - Authors:\n")
		for _, author := range vrm_meta.Authors {
			fmt.Printf("    - %s\n", author)
		}
		fmt.Printf(" - CopyrightInformation: %s\n", vrm_meta.CopyrightInformation)
		fmt.Printf(" - ContactInformation: %s\n", vrm_meta.ContactInformation)
		fmt.Printf(" - References:\n")
		for _, reference := range vrm_meta.References {
			fmt.Printf("    - %s\n", reference)
		}
		fmt.Printf(" - ThirdPartyLicenses: %s\n", vrm_meta.ThirdPartyLicenses)
		fmt.Printf(" - ThumbnailImage: %s\n", vrm_meta.ThumbnailImage)
		fmt.Printf(" - LicenseUrl: %s\n", vrm_meta.LicenseUrl)
		fmt.Printf(" - AvatarPermission: %s\n", vrm_meta.AvatarPermission)
		fmt.Printf(" - AllowExcessivelyViolentUsage: %t\n", vrm_meta.AllowExcessivelyViolentUsage)
		fmt.Printf(" - AllowExcessivelySexualUsage: %t\n", vrm_meta.AllowExcessivelySexualUsage)
		fmt.Printf(" - CommercialUsage: %s\n", vrm_meta.CommercialUsage)
		fmt.Printf(" - AllowPoliticalOrReligiousUsage: %t\n", vrm_meta.AllowPoliticalOrReligiousUsage)
		fmt.Printf(" - AllowAntisocialOrHateUsage: %t\n", vrm_meta.AllowAntisocialOrHateUsage)
		fmt.Printf(" - CreditNotation: %s\n", vrm_meta.CreditNotation)
		fmt.Printf(" - AllowRedistribution: %t\n", vrm_meta.AllowRedistribution)
		fmt.Printf(" - Modification: %s\n", vrm_meta.Modification)
		fmt.Printf(" - OtherLicenseUrl: %s\n", vrm_meta.OtherLicenseUrl)

	}
	fmt.Println()
}
