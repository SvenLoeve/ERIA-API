package query

import (
	"context"
	"fmt"
	"viseh-api/database"
	"viseh-api/database/ent"
	"viseh-api/database/ent/user"
)

func GetNfcWithId(ctx context.Context, uid int) (*ent.User, error) {

	u, err := database.Client.User.
		Query().
		Where(user.ID(uid)).
		Select(user.FieldData, user.FieldEncryptionKey).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func GetNfcWithMedID(ctx context.Context, medID string) (*ent.User, error) {

	u, err := database.Client.User.
		Query().
		Where(user.MedID(medID)).
		Select(user.FieldData, user.FieldEncryptionKey).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func UpdateNfc(ctx context.Context, uid int, data []byte) error {

	_, err := database.Client.User.
		UpdateOneID(uid).
		SetData(data).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user: %w", err)
	}
	return nil
}
