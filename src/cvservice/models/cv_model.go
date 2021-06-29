package models

type CVModel struct {
	Id      int
	Title   string
	Content []string
}

type CVModels []CVModel

func (cvModel *CVModel) Setter(id int, title string, content []string) {
	cvModel.Id = id
	cvModel.Title = title
	cvModel.Content = content
}
func (cvModel CVModel) Getter() CVModel {
	return CVModel{cvModel.Id, cvModel.Title, cvModel.Content}
}

func (cvModels *CVModels) Add(cvModel CVModel) {
	*cvModels = append(*cvModels, cvModel)
}
