package models

type CVModel struct {
	Id      int
	Title   string
	Content []Info
}
type Info struct {
	Id       int
	Text     string
	Progress int
}

// type CVModels []CVModel

func (cvModel *CVModel) Setter(id int, title string, content []Info) {
	cvModel.Id = id
	cvModel.Title = title
	cvModel.Content = content
}
func (cvModel CVModel) Getter() CVModel {
	return CVModel{cvModel.Id, cvModel.Title, cvModel.Content}
}

// func (cvModels *CVModels) Add(cvModel CVModel) {
// 	*cvModels = append(*cvModels, cvModel)
// }
