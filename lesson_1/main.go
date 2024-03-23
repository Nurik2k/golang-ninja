package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged out",
	"logged in",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestamp)
	}

	return out
}

func main() {
	rand.Seed(time.Now().Unix())

	users := generateUsers(1000)

	for _, user := range users {
		saveUserInfo(user)
	}
}

func saveUserInfo(user User) error {
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d\n", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	return err
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@ninja.go", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}

	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions))-1],
		}
	}

	return logs
}
