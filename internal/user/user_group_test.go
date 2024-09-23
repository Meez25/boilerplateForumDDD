package user

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewUserGroup(t *testing.T) {
	owner := User{ID: uuid.New(), FirstName: "Owner"}

	t.Run("Valid UserGroup", func(t *testing.T) {
		ug, err := NewUserGroup("Test Group", "Test Description", owner)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if ug.ID == uuid.Nil {
			t.Error("Expected non-nil UUID, got nil")
		}
		if ug.Name != "Test Group" {
			t.Errorf("Expected name 'Test Group', got '%s'", ug.Name)
		}
		if ug.Description != "Test Description" {
			t.Errorf("Expected description 'Test Description', got '%s'", ug.Description)
		}
		if !reflect.DeepEqual(ug.Owner, owner) {
			t.Errorf("Expected owner %v, got %v", owner, ug.Owner)
		}
	})

	t.Run("Empty Name", func(t *testing.T) {
		_, err := NewUserGroup("", "Test Description", owner)
		if err != ErrEmptyName {
			t.Errorf("Expected ErrEmptyName, got %v", err)
		}
	})
}

func TestUserGroupMethods(t *testing.T) {
	owner := User{ID: uuid.New(), FirstName: "Owner"}
	member1 := User{ID: uuid.New(), FirstName: "Member1"}
	member2 := User{ID: uuid.New(), FirstName: "Member2"}

	ug, _ := NewUserGroup("Test Group", "Test Description", owner)

	t.Run("AddMember", func(t *testing.T) {
		ug.AddMember(member1)
		if len(ug.Members) != 1 {
			t.Errorf("Expected 1 member, got %d", len(ug.Members))
		}
		if !reflect.DeepEqual(ug.Members[0], member1) {
			t.Errorf("Expected member %v, got %v", member1, ug.Members[0])
		}
	})

	t.Run("IsMember", func(t *testing.T) {
		if !ug.IsMember(member1) {
			t.Error("Expected member1 to be a member")
		}
		if ug.IsMember(member2) {
			t.Error("Expected member2 to not be a member")
		}
	})

	t.Run("RemoveMember", func(t *testing.T) {
		ug.RemoveMember(member1)
		if len(ug.Members) != 0 {
			t.Errorf("Expected 0 members, got %d", len(ug.Members))
		}
		if ug.IsMember(member1) {
			t.Error("Expected member1 to be removed")
		}
	})

	t.Run("IsOwner", func(t *testing.T) {
		if !ug.IsOwner(owner) {
			t.Error("Expected owner to be the owner")
		}
		if ug.IsOwner(member1) {
			t.Error("Expected member1 to not be the owner")
		}
	})

	t.Run("ChangeOwner", func(t *testing.T) {
		ug.ChangeOwner(member1)
		if !reflect.DeepEqual(ug.Owner, member1) {
			t.Errorf("Expected owner to be %v, got %v", member1, ug.Owner)
		}
		if !ug.IsOwner(member1) {
			t.Error("Expected member1 to be the new owner")
		}
	})

	t.Run("ChangeName", func(t *testing.T) {
		ug.ChangeName("New Name")
		if ug.Name != "New Name" {
			t.Errorf("Expected name 'New Name', got '%s'", ug.Name)
		}
	})

	t.Run("ChangeDescription", func(t *testing.T) {
		ug.ChangeDescription("New Description")
		if ug.Description != "New Description" {
			t.Errorf("Expected description 'New Description', got '%s'", ug.Description)
		}
	})
}

