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
	router.HandleFunc("/get_certification/{certificationName}", standardRouter.getCertificationHandler).Methods("GET")
	router.HandleFunc("/addCertificationToUser", standardRouter.addCertificationToUserHandler).Methods("PUT")
	router.HandleFunc("/getCertificationForUser/{userName}", standardRouter.getCertificationForUserHandler).Methods("GET")
	return router
}


func(sr *standardRouter) getCertificationForUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	err, std := sr.standardService.GetCertificationForUser(userName)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, std)
}

func(sr *standardRouter) addCertificationToUserHandler(w http.ResponseWriter, r *http.Request) {

	model := root.UserCertModel{}
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&model)
	}
	fmt.Println(model)
	sr.standardService.AddCertificationToUser(model)
	Json(w, http.StatusOK, model)
}

func(sr *standardRouter) loadCertificationHandler(w http.ResponseWriter, r *http.Request) {

	path := "/home/mukul/git/certifications/"
	//path := "/home/ec2-user/git/certifications/"
	filenames := []string{"fedramp-high.yaml", "fedramp-moderate.yaml", "fedramp-low.yaml", "fisma-high-impact.yaml", "fisma-mod-impact.yaml", "fisma-low-impact.yaml", "icd-503-high.yaml", "icd-503-moderate.yaml", "icd-503-low.yaml", "dhs-4300a.yaml"}
	certs  := []root.Certification{}

	for _, file := range filenames {
		_, cert := LoadCertification(path + file)
		sr.standardService.CreateCertification(&cert)
		certs = append(certs, cert)
	}

	Json(w, http.StatusOK, certs)
}

func LoadCertification(file string) (error, root.Certification){

	certYamlFile, err := ioutil.ReadFile(file)
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

func(sr *standardRouter) getCertificationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	certificationName := vars["certificationName"]

	err, cert := sr.standardService.GetCertificationInfo(certificationName)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, cert)
}



func LoadStandards() (error, []root.Standard){

	//path := "/home/mukul/git/standards/"
	path := "/home/ec2-user/git/standards/"
	filenames := []string{"nist-800-53-latest.yaml", "tsc-2017.yaml"}
	stds := []root.Standard{}
	var err error
	for _, file := range filenames{
		fullpa := path + file
		standardsYamlFile, err := ioutil.ReadFile(fullpa)
		if err != nil {
			log.Printf("standardsYamlFile.Get err   #%v ", err)
		}
		standardsJson, err := yaml.YAMLToJSON(standardsYamlFile)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		var standardsResult map[string]interface{}
		json.Unmarshal([]byte(standardsJson), &standardsResult)
		i := 0
		stdName := standardsResult["name"].(string)
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

				controlInfo := root.ControlInfo{ Family:family, Name:name, Description:desc }
				controls = append(controls, root.Controls{ ControlName: key , ControlInfo: controlInfo })
				i += 1
				standard := root.Standard{StandardName:stdName, Controls: controls}
				stds = append(stds, standard)
			}
		}
	}
	return err, stds

}


