
package entities

import (
    "github.com/google/uuid"
    "time"
)

type Project struct {
    ID          uuid.UUID
    Name        string
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
