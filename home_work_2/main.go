package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

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
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func main() {
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	users := generateUsers(100)
	saveUserInfos(users)

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfos(users []User) {
	var wg sync.WaitGroup
	workerCount := 10 // Number of workers in the pool

	userChannel := make(chan User, len(users))
	wg.Add(workerCount)

	// Start worker pool
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for user := range userChannel {
				saveUserInfo(user)
			}
		}()
	}

	// Send users to the channel for processing
	for _, user := range users {
		userChannel <- user
	}
	close(userChannel)

	wg.Wait()
}

func saveUserInfo(user User) {
	fmt.Printf("WRITING FILE FOR UID %d\n", user.id)

	filename := fmt.Sprintf("users/uid%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.getActivityInfo())
	time.Sleep(time.Second)
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@company.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
		fmt.Printf("generated user %d\n", i+1)
		time.Sleep(time.Millisecond * 100)
	}

	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions))],
			timestamp: time.Now(),
		}
	}

	return logs
}

// Преимущество использования паттерна Worker Pool заключается в том, что он позволяет 
//обрабатывать задачи параллельно с ограниченным количеством горутин (воркеров). 
//Это позволяет управлять нагрузкой на систему и повысить эффективность использования ресурсов процессора.

// Прирост в скорости зависит от множества факторов, таких как количество пользователей, 
//количество логов для каждого пользователя, количество воркеров в пуле и характеристики системы, 
//на которой выполняется код. Однако, в целом, параллельное выполнение операций в пуле воркеров должно 
//значительно сократить время выполнения по сравнению с последовательным выполнением операций.