package root

type Controls struct {
	ControlName string `json:"controlName"`
	ControlInfo ControlInfo `json:"controlInfo"`
}

type ControlInfo struct {
	Family          string  `json:"family"`
	Name     		string  `json:"name"`
	Description     string  `json:"desc"`
}

type Standard struct {
	StandardName string `json:"standardName"`
	Controls[] Controls `json:"controls"`
}

type Certification struct {
	CertificationName string `json:"standardName"`
	StandardName string `json:"certificationName"`
	ControlName[] string `json:"controls"`
}

type UserCertModel struct {
	UserName string
	Controls []string
}

type StandardService interface {
	CreateStandard(u *Standard) error
	GetStandardInfo(standardname string) (error, []Standard)
	CreateCertification(u *Certification) error
	GetCertificationInfo(certificationName string) (error, []Certification)
	AddCertificationToUser(model UserCertModel) error
	GetCertificationForUser(userName string) (error, []string)
}

