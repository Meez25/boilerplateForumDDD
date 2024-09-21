package persistence

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/forum"
)

var ErrTopicNotFound = errors.New("ErrTopicNotFound")

type TopicMemoryRepo struct {
	topics map[uuid.UUID]forum.Topic
}

func NewTopicMemoryRepo() *TopicMemoryRepo {
	return &TopicMemoryRepo{
		topics: make(map[uuid.UUID]forum.Topic),
	}
}

func (r *TopicMemoryRepo) Save(topic forum.Topic) error {
	r.topics[topic.ID] = topic
	return nil
}

func (r *TopicMemoryRepo) FindByID(ID string) (forum.Topic, error) {
	t, ok := r.topics[uuid.MustParse(ID)]
	if !ok {
		return forum.Topic{}, ErrTopicNotFound
	}

	return t, nil
}

func (r *TopicMemoryRepo) FindAll() ([]forum.Topic, error) {
	topics := make([]forum.Topic, 0, len(r.topics))
	for _, t := range r.topics {
		topics = append(topics, t)
	}

	return topics, nil
}

func (r *TopicMemoryRepo) Update(topic forum.Topic) error {
	r.topics[topic.ID] = topic
	return nil
}

func (r *TopicMemoryRepo) Delete(ID string) error {
	delete(r.topics, uuid.MustParse(ID))
	return nil
}
