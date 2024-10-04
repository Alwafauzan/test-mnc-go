package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Activity represents an activity log entry.
type Activity struct {
	Timestamp time.Time `json:"timestamp"`
	Activity  string    `json:"activity"`
}

// HistoryService defines the methods available for the history service.
type HistoryService interface {
	LogActivity(filePath, activity string) error
}

type historyService struct{}

// NewHistoryService returns a new instance of HistoryService.
func NewHistoryService() HistoryService {
	return &historyService{}
}

func (s *historyService) LogActivity(filePath, activity string) error {
	activityLog := Activity{
		Timestamp: time.Now(),
		Activity:  activity,
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read history file: %v", err)
		return err
	}

	var activities []Activity
	if err := json.Unmarshal(data, &activities); err != nil {
		log.Printf("Failed to unmarshal history: %v", err)
		return err
	}

	activities = append(activities, activityLog)

	updatedData, err := json.Marshal(activities)
	if err != nil {
		log.Printf("Failed to marshal updated history: %v", err)
		return err
	}

	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		log.Printf("Failed to write updated history: %v", err)
		return err
	}

	fmt.Printf("Logged activity: %v\n", activityLog)
	return nil
}