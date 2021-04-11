// Repository interface
// provides methods to act with postgres person table

package repository

// PersonModel represent the user model
type PersonModel struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	//ProjectName string `json:"project_name"`
	ProjectEnv  string `json:"project_env"`
	CompanyID   string `json:"company_id"`
	TeamID      string `json:"team_id"`
	GroupID     string `json:"group_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

// Repository represent the repositories
type Repository interface {
	Up() error
	Close()
	Drop() error
	Truncate() error

	Create(user *PersonModel) (string, error)
	Update(user *PersonModel) error
	Delete(id string) error
	// Find look for something
	Find() ([]*PersonModel, error)
	FindByID(id string) (*PersonModel, error)
}
