package server

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"ui-mockup-backend"
)

type standardRouter struct {
	standardService root.StandardService
	auth *authHelper
}

func NewStandardRouter(u root.StandardService, router *mux.Router, a *authHelper) *mux.Router {
	standardRouter := standardRouter{u,a}
	//router.HandleFunc("/load_standards", a.validate(standardRouter.loadStandardHandler)).Methods("GET")
	router.HandleFunc("/load_standards", standardRouter.loadStandardHandler).Methods("GET")
	router.HandleFunc("/get_standard/{standardName}", standardRouter.getStandardHandler).Methods("GET")
	router.HandleFunc("/load_certifications", standardRouter.loadCertificationHandler).Methods("GET")
	return router
}

func(sr *standardRouter) loadCertificationHandler(w http.ResponseWriter, r *http.Request) {
	err, cert := LoadCertification()

	sr.standardService.CreateCertification(&cert)

	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, cert)
}

func LoadCertification() (error, root.Certification){

	certYamlFile, err := ioutil.ReadFile("/home/mukul/git/certifications/fedramp-low.yaml")
	if err != nil {
		log.Printf("certYamlFile.Get err   #%v ", err)
	}

	certJson, err := yaml.YAMLToJSON(certYamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}


	var certResult map[string]interface{}
	json.Unmarshal([]byte(certJson), &certResult)
	cert := root.Certification{}
	cert.CertificationName = certResult["name"].(string)
	ctrls := []string{}

	for key, value := range certResult["standards"].(map[string]interface{}) {

		cert.StandardName = key

		for k, _ := range value.(map[string]interface{}) {

			ctrls = append(ctrls, k)
		}
		cert.ControlName = ctrls
	}

	return err, cert

}


func(sr *standardRouter) loadStandardHandler(w http.ResponseWriter, r *http.Request) {
	err, stds := LoadStandards()

	for stdI := range stds{
		sr.standardService.CreateStandard(&stds[stdI])
	}

	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, stds)
}

func(sr *standardRouter) getStandardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	standardName := vars["standardName"]

	err, std := sr.standardService.GetStandardInfo(standardName)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, std)
}


func LoadStandards() (error, []root.Standard){

	//print("LOADING STANDARDS")

	standardsYamlFile, err := ioutil.ReadFile("/home/mukul/git/standards/nist-800-53-latest.yaml")
	if err != nil {
		log.Printf("standardsYamlFile.Get err   #%v ", err)
	}
	standardsJson, err := yaml.YAMLToJSON(standardsYamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		//return err, "nist-800-53-latest"
	}

	//print(standardsJson)

	var standardsResult map[string]interface{}
	json.Unmarshal([]byte(standardsJson), &standardsResult)

	//var controls[] root.Controls
	stds := []root.Standard{}
	i := 0
	for key, value := range standardsResult {
		// Each value is an interface{} type, that is type asserted as a string
		controls := []root.Controls{}
		var desc, family, name string
		vt := reflect.TypeOf(value).Kind()
		if (vt != reflect.String){
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
			//controls[i] = rmongooot.Controls{ ControlName: key , ControlInfo: controlInfo }
			controls = append(controls, root.Controls{ ControlName: key , ControlInfo: controlInfo })
			i += 1
			// todo: Replace with standard name from file name
			standard := root.Standard{StandardName:"nist-800-53-latest", Controls: controls}
			//fmt.Print(standard)
			// TODO: insert every standard into DB
			//standardService := new(mongo.StandardsService)
			stds = append(stds, standard)
			//fmt.Println(standard)
			//break // TODO: remove after test

		}
		}
	// todo: Replace with standard name from file name
	return err, stds

}


