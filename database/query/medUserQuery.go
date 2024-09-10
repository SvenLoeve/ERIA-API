package query

import (
	"context"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"log"
	"viseh-api/database"
	"viseh-api/database/ent"
	"viseh-api/database/ent/meduser"
	visehmath "viseh-api/services/math"
	"viseh-api/types"
)

func GetMedUser(ctx context.Context, id int) (types.MedUser, error) {

	mu, err := database.Client.MedUser.
		Query().
		Where(meduser.ID(id)).
		Only(ctx)
	if err != nil {
		return types.MedUser{}, fmt.Errorf("failed querying user: %w", err)
	}
	return types.ParseMedUser(mu), nil
}

func CreateMedUser(ctx context.Context, MedUser ent.MedUser) error {
	_, err := database.Client.MedUser.
		Create().
		SetName(MedUser.Name).
		SetEmail(MedUser.Email).
		SetPassword(MedUser.Password).
		SetPhoneNumber(MedUser.PhoneNumber).
		SetRole("ambulance").
		SetOrganisation(MedUser.Organisation).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return err
}
func GetMedUsers(ctx context.Context, pageNumber int, pageSize int) (types.MedUsers, error) {
	userList, err := database.Client.MedUser.
		Query().
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		Order(meduser.ByID()).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	medUsers := types.MedUsers{}
	for i := 0; i < len(userList); i++ {
		medUsers = append(medUsers, types.ParseMedUser(userList[i]))
	}
	return medUsers, nil
}

func UpdateMedUser(ctx context.Context, MedUser types.MedUser) error {
	u, err := database.Client.MedUser.
		Update().
		SetName(MedUser.Name).
		SetEmail(MedUser.Email).
		SetPhoneNumber(MedUser.PhoneNumber).
		SetRole(meduser.Role(MedUser.Role)).
		SetOrganisation(MedUser.Organisation).
		Where(meduser.ID(MedUser.ID)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update med user: %w", err)
	}
	log.Println("med user updated. id:", u)
	return err
}
func SearchMedUsers(ctx context.Context, email string) (types.MedUsers, error) {
	userList, err := database.Client.MedUser.
		Query().
		Where(meduser.EmailContains(email)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed searching med user: %w", err)
	}

	users := types.MedUsers{}
	for i := 0; i < len(userList); i++ {
		users = append(users, types.ParseMedUser(userList[i]))
	}

	return users, nil
}
func MedEmailExists(ctx context.Context, email string) (bool, error) {
	userList, err := database.Client.MedUser.
		Query().
		Where(meduser.Email(email)).
		All(ctx)
	if err != nil {
		return true, fmt.Errorf("failed searching user: %w", err)
	}

	if len(userList) == 0 {
		return false, nil
	}
	return true, nil
}

func PatchMedUser(ctx context.Context, User types.FullMedUser) error {
	userTarget, err := GetFullMedUser(ctx, User.ID)
	if err != nil {
		return fmt.Errorf("failed to get full user: %w", err)
	}
	jsonUser1, err := json.Marshal(User)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	jsonUser2, err := json.Marshal(userTarget)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	modifiedPatch, err := jsonpatch.MergePatch(jsonUser2, jsonUser1)
	if err != nil {
		return fmt.Errorf("error merge patching JSON: %w", err)
	}

	var modifiedUser types.FullMedUser
	err = json.Unmarshal(modifiedPatch, &modifiedUser)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	_, err = database.Client.MedUser.
		Update().
		SetName(modifiedUser.Name).
		SetEmail(modifiedUser.Email).
		SetPhoneNumber(modifiedUser.PhoneNumber).
		SetRole(meduser.Role(modifiedUser.Role)).
		SetOrganisation(modifiedUser.Organisation).
		SetPassword(modifiedUser.Password).
		Where(meduser.ID(modifiedUser.ID)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to patch med user: %w", err)
	}
	return nil
}

// GetFullMedUser Returns user with password.
func GetFullMedUser(ctx context.Context, id int) (*ent.MedUser, error) {
	user, err := database.Client.MedUser.
		Query().
		Where(meduser.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return user, nil
}

func TotalMedUsersAndTotalPages(ctx context.Context, items int) (int, int, error) {
	userList, err := database.Client.MedUser.
		Query().
		All(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("failed querying med user: %w", err)
	}
	totalUsers := len(userList)

	totalPages, remainder := visehmath.DivideWithRemainder(totalUsers, items)
	if remainder != 0 {
		totalPages++
	}
	return totalUsers, totalPages, nil
}
func GetTotalMedUsers(ctx context.Context) (int, error) {
	userList, err := database.Client.MedUser.
		Query().
		All(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed querying med user: %w", err)
	}
	totalUsers := len(userList)

	return totalUsers, nil
}
