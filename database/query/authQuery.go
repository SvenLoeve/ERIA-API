package query

import (
	"context"
	"fmt"
	"viseh-api/database"
	"viseh-api/database/ent"
	"viseh-api/database/ent/meduser"
	"viseh-api/database/ent/user"
)

func GetLoginUser(ctx context.Context, email string) (*ent.User, error) {

	u, err := database.Client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func GetLoginMedUser(ctx context.Context, email string) (*ent.MedUser, error) {

	u, err := database.Client.MedUser.
		Query().
		Where(meduser.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}
