package services

import (
	"github.com/google/uuid"
	"github.com/meez25/boilerplateForumDDD/internal/forum"
)

type TopicService struct {
	topicRepository forum.TopicRepository
}

func NewTopicService(topicRepository forum.TopicRepository) *TopicService {
	return &TopicService{
		topicRepository: topicRepository,
	}
}

func (ts *TopicService) CreateTopic(title string, richContent string, authorID uuid.UUID) (forum.Topic, error) {
	forumTopic, err := forum.NewTopic(title, richContent, authorID)
	if err != nil {
		return forum.Topic{}, err
	}

	err = ts.topicRepository.Save(forumTopic)
	if err != nil {
		return forum.Topic{}, err
	}

	return forumTopic, nil

}

func (ts *TopicService) AddMessage(topicID string, richContent string, authorID uuid.UUID) error {
	topic, err := ts.topicRepository.FindByID(topicID)
	if err != nil {
		return err
	}

	err = topic.AddMessage(richContent, authorID)
	if err != nil {
		return err
	}

	return ts.topicRepository.Update(topic)
}

func (ts *TopicService) GetTopicByID(ID string) (forum.Topic, error) {
	return ts.topicRepository.FindByID(ID)
}

func (ts *TopicService) GetAllTopics() ([]forum.Topic, error) {
	return ts.topicRepository.FindAll()
}

func (ts *TopicService) UpdateMessage(topicID string, messageID string, richContent string, authorID uuid.UUID) error {
	topic, err := ts.topicRepository.FindByID(topicID)
	if err != nil {
		return err
	}

	err = topic.UpdateMessage(messageID, richContent, authorID)
	if err != nil {
		return err
	}

	return ts.topicRepository.Update(topic)
}

func (ts *TopicService) DeleteTopic(ID string) error {
	return ts.topicRepository.Delete(ID)
}

func (ts *TopicService) DeleteMessage(topicID string, messageID string) error {
	topic, err := ts.topicRepository.FindByID(topicID)
	if err != nil {
		return err
	}

	err = topic.DeleteMessage(messageID)
	if err != nil {
		return err
	}

	return ts.topicRepository.Update(topic)
}
