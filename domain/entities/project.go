
package entities

import "time"

type Project struct {
    ID          int
    Name        string
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func NewProject(name, description string) *Project {
    return &Project{
        Name:        name,
        Description: description,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
}
