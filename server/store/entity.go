package store

import (
	"context"
	"fmt"

	// sqlite necessary
	_ "github.com/mattn/go-sqlite3"
)

// CreateEntity creates a Entity entry in the db
func (gs *GeneralStore) createEntity(ctx context.Context, entity interface{}, createEntitySQL string) error {
	return gs.namedExec(ctx, createEntitySQL, entity)
}

func (gs *GeneralStore) deleteEntity(ctx context.Context, entityTable table, entityID string) error {
	deleteEntitySQL := fmt.Sprintf(`DELETE FROM %v WHERE id=$1`, entityTable)
	return gs.exec(ctx, deleteEntitySQL, entityID)
}

func (gs *GeneralStore) findEntity(ctx context.Context, entity interface{}, entityTable table, entityID string) error {
	findEntitySQL := fmt.Sprintf(`SELECT * FROM %v WHERE id=$1`, entityTable)
	return gs.get(ctx, entity, findEntitySQL, entityID)
}

// findAllOfEntity finds all entries in the db for the given entity
func (gs *GeneralStore) findAllEntity(ctx context.Context, entity interface{}, entityTable table) error {
	findAllEntitySQL := fmt.Sprintf(`SELECT * FROM %v ORDER BY id ASC`, entityTable)
	return gs.getSet(ctx, entity, findAllEntitySQL)
}

func (gs *GeneralStore) updateEntity(ctx context.Context, entity interface{}, updateEntitySQL string) error {
	return gs.namedExec(ctx, updateEntitySQL, entity)
}
