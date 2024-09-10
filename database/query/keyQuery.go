package query

import (
	"context"
	"fmt"
	"time"
	"viseh-api/database"
	"viseh-api/database/ent"
	"viseh-api/database/ent/key"
)

func SetKey(ctx context.Context, key string) error {

	_, err := database.Client.Key.
		Create().
		SetID(time.Now().Format("2006/02/01")).
		SetKey(key).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed querying user: %w", err)
	}
	return nil
}

func GetLatestKey(ctx context.Context) (*ent.Key, error) {

	k, err := database.Client.Key.
		Query().
		Where(key.ID(time.Now().Format("2006/02/01"))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return k, nil
}

func GetAllKeys(ctx context.Context) ([]*ent.Key, error) {

	k, err := database.Client.Key.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return k, nil
}
