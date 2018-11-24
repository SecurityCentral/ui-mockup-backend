package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ui-mockup-backend"
)

type StandardsService struct {
	stdCollection *mgo.Collection
	certCollection *mgo.Collection
	userCollection *mgo.Collection
}

func NewStandardsService(session *mgo.Session, config *root.MongoConfig) *StandardsService {
	Stdcollection := session.DB(config.DbName).C("std")
	Certcollection := session.DB(config.DbName).C("cert")
	Usercollection := session.DB(config.DbName).C("userCert")
	Stdcollection.EnsureIndex(standardsIndex())
	Certcollection.EnsureIndex(certificationIndex())
	Usercollection.EnsureIndex(userCertIndex())
	return &StandardsService{stdCollection:Stdcollection, certCollection:Certcollection, userCollection:Usercollection}
}

func (p *StandardsService) CreateStandard(std *root.Standard) error {
	standard := newStandardModel(std)
	return p.stdCollection.Insert(&standard)
}

func (p *StandardsService) CreateCertification(u *root.Certification) error {
	return p.certCollection.Insert(&u)
}

func (p *StandardsService) GetCertificationForUser(userName string) (error, []root.UserControlModel) {
	certModel := root.UserCertModel{}
	err := p.userCollection.Find(bson.M{"username": userName}).One(&certModel)
	return err, certModel.Controls
}

func (p *StandardsService) GetStandardInfo(standardName string) (error, []root.Standard) {
	standardsModel := []root.Standard{}
	err := error(nil)
	if (standardName == "all") {
		err = p.stdCollection.Find(bson.M{}).Iter().All(&standardsModel)
	} else {
		err = p.stdCollection.Find(bson.M{"standardname": standardName}).Iter().All(&standardsModel)
	}
	return err, standardsModel
}

func (p *StandardsService) AddCertificationToUser(model root.UserCertModel) error{
	return p.userCollection.Insert(&model)
}

func (p *StandardsService) GetCertificationInfo(certificationName string) (error, []root.Certification) {
	certModel := []root.Certification{}
	cert := root.Certification{}
	err := error(nil)
	if (certificationName != "all"){
		err = p.certCollection.Find(bson.M{"certificationname": certificationName}).One(&cert)
		certModel = make([]root.Certification, 1)
		certModel[0] = cert
	} else {
		err = p.certCollection.Find(bson.M{}).Iter().All(&certModel)
	}
	return err, certModel
}