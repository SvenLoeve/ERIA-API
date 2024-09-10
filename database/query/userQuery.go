package query

import (
	"context"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
	"log"
	"viseh-api/database"
	"viseh-api/database/ent"
	"viseh-api/database/ent/user"
	visehmath "viseh-api/services/math"
	"viseh-api/types"
)

func GetUser(ctx context.Context, id int) (types.User, error) {

	u, err := database.Client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)

	if err != nil {
		return types.User{}, fmt.Errorf("failed querying user: %w", err)
	}
	return types.ParseUser(u), nil
}

// GetFullUser Returns user with password.
func GetFullUser(ctx context.Context, id int) (*ent.User, error) {
	u, err := database.Client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func GetEncryption(ctx context.Context, id int) (*ent.User, error) {

	u, err := database.Client.User.
		Query().
		Where(user.ID(id)).
		Select(user.FieldEncryptionKey).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func CreateUser(ctx context.Context, User ent.User) error {

	u, err := database.Client.User.
		Create().
		SetName(User.Name).
		SetEmail(User.Email).
		SetPassword(User.Password).
		SetPhoneNumber(User.PhoneNumber).
		SetEncryptionKey(User.EncryptionKey).
		SetRole("user").
		SetMedID(User.MedID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	log.Println("user created: ", u)
	return err
}

func GetUsers(ctx context.Context, pageNumber int, pageSize int) (types.Users, error) {
	userList, err := database.Client.User.
		Query().
		Offset((pageNumber - 1) * pageSize).
		Order(user.ByID()).
		Limit(pageSize).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	users := types.Users{}
	for i := 0; i < len(userList); i++ {
		users = append(users, types.ParseUser(userList[i]))
	}
	return users, nil
}

func UpdateUser(ctx context.Context, User types.User, admin bool) error {
	if !admin {
		dbUser, _ := GetUser(ctx, User.ID)
		User.Role = dbUser.Role
		User.MedID = dbUser.MedID
	}
	u, err := database.Client.User.
		Update().
		SetName(User.Name).
		SetEmail(User.Email).
		SetPhoneNumber(User.PhoneNumber).
		SetRole(user.Role(User.Role)).
		SetMedID(User.MedID).
		Where(user.ID(User.ID)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	log.Println("user updated. id:", u)
	return err
}

func SearchUsers(ctx context.Context, email string) (types.Users, error) {
	userList, err := database.Client.User.
		Query().
		Where(user.EmailContains(email)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed searching user: %w", err)
	}

	users := types.Users{}
	for i := 0; i < len(userList); i++ {
		users = append(users, types.ParseUser(userList[i]))
	}

	return users, nil
}

func PatchUser(ctx context.Context, User types.FullUser, admin bool) error {
	userTarget, err := GetFullUser(ctx, User.ID)

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

	var modifiedUser types.FullUser
	err = json.Unmarshal(modifiedPatch, &modifiedUser)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	if !admin {
		dbUser, _ := GetUser(ctx, modifiedUser.ID)
		modifiedUser.Role = dbUser.Role
		modifiedUser.MedID = dbUser.MedID
	}

	u, err := database.Client.User.
		Update().
		SetName(modifiedUser.Name).
		SetEmail(modifiedUser.Email).
		SetPhoneNumber(modifiedUser.PhoneNumber).
		SetRole(user.Role(modifiedUser.Role)).
		SetMedID(modifiedUser.MedID).
		SetPassword(modifiedUser.Password).
		Where(user.ID(modifiedUser.ID)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to patch user: %w", err)
	}
	log.Println("user updated. id:", u)
	return nil
}

func EmailExists(ctx context.Context, email string) (bool, error) {
	userList, err := database.Client.User.
		Query().
		Where(user.Email(email)).
		All(ctx)
	if err != nil {
		return true, fmt.Errorf("failed searching user: %w", err)
	}

	if len(userList) == 0 {
		return false, nil
	}
	return true, nil
}

func TotalUsersAndTotalPages(ctx context.Context, items int) (int, int, error) {
	userList, err := database.Client.User.
		Query().
		All(ctx)
	if err != nil {
		return 0, 0, fmt.Errorf("failed querying user: %w", err)
	}
	totalUsers := len(userList)

	totalPages, remainder := visehmath.DivideWithRemainder(totalUsers, items)
	if remainder != 0 {
		totalPages++
	}
	return totalUsers, totalPages, nil
}
func GetTotalUsers(ctx context.Context) (int, error) {
	userList, err := database.Client.User.
		Query().
		All(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed querying med user: %w", err)
	}
	totalUsers := len(userList)

	return totalUsers, nil
}
