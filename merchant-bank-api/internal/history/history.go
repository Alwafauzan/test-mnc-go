package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Activity represents a log entry with a timestamp and a description of the activity.
type Activity struct {
	Timestamp time.Time `json:"timestamp"`
	Activity  string    `json:"activity"`
}

// LogActivity logs a new activity to the specified file.
func LogActivity(filePath, activity string) error {
	activityLog := Activity{
		Timestamp: time.Now(),
		Activity:  activity,
	}

	// Read the existing activities from the file.
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file doesn't exist, create an empty array of activities.
			data = []byte("[]")
		} else {
			log.Printf("Failed to read history file: %v", err)
			return err
		}
	}

	var activities []Activity
	if err := json.Unmarshal(data, &activities); err != nil {
		log.Printf("Failed to unmarshal history: %v", err)
		return err
	}

	// Append the new activity to the list of activities.
	activities = append(activities, activityLog)

	// Marshal the updated list of activities back to JSON.
	updatedData, err := json.MarshalIndent(activities, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal updated history: %v", err)
		return err
	}

	// Write the updated JSON back to the file.
	if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
		log.Printf("Failed to write updated history: %v", err)
		return err
	}

	fmt.Printf("Logged activity: %v\n", activityLog)
	return nil
}