package entities

type CVEntity struct {
	Id    int
	Title string
}
type InfoEntity struct {
	Id         int
	Txt        string
	Progress   int
	CvmodelId  int
	DeleteFlag bool
}
type CVEntities []CVEntity
type InfoEntities []InfoEntity
