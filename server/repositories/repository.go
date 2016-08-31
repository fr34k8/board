package repositories

import (
	"github.com/mtti/board/server/models"
)

type Repository interface {

	// Set repository's parent. If the repository needs to call any of the
	// other public methods of this interface, it should do so through its
	// parent.
	SetParent(parent Repository)

	// Load a card by its primary key ID
	LoadCardByID(id int) (*models.Card, error)

	// Load a card by its name, ie. "KEY-123"
	LoadCardByName(name string) (*models.Card, error)

	// Load a project by its key
	LoadProjectByKey(key string) (*models.Project, error)
	
}
