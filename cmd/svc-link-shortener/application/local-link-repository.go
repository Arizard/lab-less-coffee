package application

import (
	"crypto/sha256"
	"fmt"
	"github.com/arizard/lab-less-coffee/pkg/services"
	"sync"
	"time"
)

type LocalLinkRepository struct {
	links map[string]Link
	mu sync.Mutex
}

func NewLocalLinkRepository() *LocalLinkRepository {
	return &LocalLinkRepository{links: map[string]Link{}}
}


func (repo *LocalLinkRepository) Get(uid string, link *Link) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if savedLink, ok := repo.links[uid]; ok {
		link.Uid = savedLink.Uid
		link.Destination = savedLink.Destination
		link.DeleteCode = savedLink.DeleteCode

		return nil
	} else {
		return services.RepositoryRecordNotFound
	}
}

func (repo *LocalLinkRepository) Insert(newLink *Link) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	hash := sha256.Sum256([]byte(fmt.Sprintf("%s:%s", newLink.Destination, time.Now().String())))
	newLink.Uid = fmt.Sprintf("%x", hash[0:4])

	repo.links[newLink.Uid] = *newLink

	fmt.Printf("%d links saved in memory\n", len(repo.links))

	return nil
}

func (repo *LocalLinkRepository) Delete(uid string) error {
	panic("implement me")
}



