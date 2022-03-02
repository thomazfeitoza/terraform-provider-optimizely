package project

type ProjectClient interface {
	GetProject(projectId string) (Project, error)
}
