package user_test

import (
	"testing"
	_ "time"

	userpkg "github.com/itsjustabavaar/oh-my-gosh/internal/user"
	_ "github.com/itsjustabavaar/oh-my-gosh/internal/util/dbutil"
)

func TestUserScenario(t *testing.T) {
	username, password := "testuser2", "1234"

	err := userpkg.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	user, err := userpkg.Login(username, password)
	if err != nil {
		t.Fatalf("Error logging in: %v", err)
	}

	if user.Username != username {
		t.Fatalf("Username does not match")
	}

	err = userpkg.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestUserNotFoundLogin(t *testing.T) {
	username, password := "testuser2", "1234"

	err := userpkg.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	err = userpkg.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}

	_, err = userpkg.Login(username, password)
	if err != nil {
		if err.Error() != "user not found" {
			t.Fatal("unexpected error")
		}
	}
}

func TestUserAlreadyExists(t *testing.T) {
	username, password := "testuser2", "1234"
	err := userpkg.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	err = userpkg.AddUser(username, password)
	if err != nil {
		if err.Error() != "user already exists" {
			t.Fatal("unexpected error")
		}
	}
	err = userpkg.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestUserLoginInvalidPassword(t *testing.T) {
	username, password := "testuser2", "1234"
	err := userpkg.AddUser(username, password)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	_, err = userpkg.Login(username, "1235")
	if err != nil {
		if err.Error() != "invalid password" {
			t.Fatal("unexpected error")
		}
	}
	err = userpkg.DeleteUser(username)
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}
