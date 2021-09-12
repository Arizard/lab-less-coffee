package infra

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
)

func TestJSONUserRepositoryCreate(t *testing.T) {
	repo := NewJSONUserRepository(fmt.Sprintf("data/user.%d.json", rand.Int()))

	newUser := entity.User{
		Login: "aoldman",
		UID:   entity.NewUserUID(),
	}

	user, err := repo.Create(newUser)

	if err != nil {
		t.Error(err)
	}

	if user.Login != "aoldman" || user.UID != newUser.UID {
		t.Errorf("expected %+v, got %+v", newUser, user)
	}
}

func TestJSONUserRepositoryGet(t *testing.T) {
	repo := NewJSONUserRepository(fmt.Sprintf("data/user.%d.json", rand.Int()))
	uid := entity.NewUserUID()

	newUser := entity.User{
		Login: "aoldman01",
		UID:   uid,
	}

	_, err := repo.Create(newUser)

	if err != nil {
		t.Error(err)
	}

	user, err := repo.Get(uid)

	if err != nil {
		t.Error(err)
	}

	if user.Login != "aoldman01" || user.UID != newUser.UID {
		t.Errorf("expected %+v, got %+v", newUser, user)
	}

}
