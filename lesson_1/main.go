package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
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

	t := time.Now()

	wg := &sync.WaitGroup{}

	users := make(chan User)
	go generateUsers(1000, users)
	for user := range users {
		wg.Add(1)
		go saveUserInfo(user, wg)
	}

	wg.Wait()

	fmt.Println("Time elapsed:", time.Since(t).String())
}

func saveUserInfo(user User, wg *sync.WaitGroup) error {

	defer wg.Done()

	fmt.Printf("WRITING FILE FOR USER ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d\n", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	if err != nil {
		return err
	}

	return nil
}

func generateUsers(count int, users chan User) {
	defer close(users)

	for i := 0; i < count; i++ {
		users <- User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@ninja.go", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
		//time.Sleep(time.Millisecond * 10)
	}
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}

	return logs
}
