package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"ui-mockup-backend"
	"ui-mockup-backend/mongo"

	"github.com/ghodss/yaml"
)

/*
type Standards struct {
	ControlName string
	ControlInfo ControlInfo
}


type ControlInfo struct {
	Family          string  `json:"family"`
	Name     		string  `json:"name"`
	Description     string  `json:"desc"`
}

type Certification struct {
	CertificationName string
	StandardsList[]	Standards
}
*/

func temp() {

	//a := App{}
	//a.Initialize()
	//a.Run()
	LoadStandards()


	/*
	standardsYamlFile, err := ioutil.ReadFile("/Users/gauravbang/Documents/meng/security-central/standards/nist-800-53-latest.yaml")
	if err != nil {
		log.Printf("standardsYamlFile.Get err   #%v ", err)
	}
	standardsJson, err := yaml.YAMLToJSON(standardsYamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	var standardsResult map[string]interface{}
	json.Unmarshal([]byte(standardsJson), &standardsResult)


	for key, value := range standardsResult {
		// Each value is an interface{} type, that is type asserted as a string

		var desc, family, name string
		for k, v := range value.(map[string]interface{}) {
			if k == "family" {
				family = v.(string)
			}
			if k == "name" {
				name = v.(string)
			}
			if k == "description" {
				desc = v.(string)
			}
		}

		controlInfo := ControlInfo{ Family:family, Name:name, Description:desc }
		standard := Standards{ControlInfo: controlInfo, ControlName:key}
		// TODO: insert every standard into DB
		fmt.Println(standard)
		break // TODO: remove after test
	}
	*/

	// TODO: Certifications
	/*
	certificationYamlFile, err := ioutil.ReadFile("/Users/gauravbang/Documents/meng/security-central/certifications/fedramp-low.yaml")
	if err != nil {
		log.Printf("certificationYamlFile.Get err   #%v ", err)
	}
	certificatesJson, err := yaml.YAMLToJSON(certificationYamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	var certificatesResult map[string]interface{}
	json.Unmarshal([]byte(certificatesJson), &certificatesResult)

	fmt.Println(certificatesResult)
	for key, value := range certificatesResult {
		fmt.Println(key)
		fmt.Println(value)
	}
	*/

}


func LoadStandards() (error, string){

	print("Loading Standards....\n")

	standardsYamlFile, err := ioutil.ReadFile("/home/mukul/git/standards/nist-800-53-latest.yaml")
	if err != nil {
		log.Printf("standardsYamlFile.Get err   #%v ", err)
	}
	standardsJson, err := yaml.YAMLToJSON(standardsYamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err, "nist-800-53-latest"
	}

	var standardsResult map[string]interface{}
	json.Unmarshal([]byte(standardsJson), &standardsResult)

	//var controls[] root.Controls
	i := 0
	for key, value := range standardsResult {
		// Each value is an interface{} type, that is type asserted as a string
		controls := []root.Controls{}

		var desc, family, name string
		for k, v := range value.(map[string]interface{}) {
			if k == "family" {
				family = v.(string)
			}
			if k == "name" {
				name = v.(string)
			}
			if k == "description" {
				desc = v.(string)
			}
		}

		//controlInfo := ControlInfo{ Family:family, Name:name, Description:desc }
		//standard := Standards{ControlInfo: controlInfo, ControlName:key}
		//controlInfo := root.Controls{ Family:family, Name:name, Description:desc }
		controlInfo := root.ControlInfo{ Family:family, Name:name, Description:desc }
		//print(controlInfo)
		//print(key)
		//controls[i] = root.Controls{ ControlName: key , ControlInfo: controlInfo }
		controls = append(controls, root.Controls{ ControlName: key , ControlInfo: controlInfo })
		//fmt.Println(controls)
		i += 1
		// todo: Replace with standard name from file name
		standard := root.Standard{StandardName:"nist-800-53-latest", Controls: controls}
		//print(standard.StandardName)
		// TODO: insert every standard into DB
		//var standardService root.StandardService
		var standardService mongo.StandardsService
		standardService.CreateStandard(&standard)
		fmt.Println(standard)
		break // TODO: remove after test
	}
	// todo: Replace with standard name from file name
	return err, "nist-800-53-latest"

}

