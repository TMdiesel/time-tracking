
package interfaces

import "time-tracker/domain/entities"

type IProjectRepository interface {
    Save(project *entities.Project) error
}
