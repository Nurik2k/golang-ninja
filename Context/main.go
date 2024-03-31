package main

import (
	"context"
	"fmt"
	"time"
)

// 1. context.Background() - на самом высоком уровне.
// 2. context.TODO - кода не уверены, какой контекст использовать
// 3. context.Value - стоит использовать как можно реже, и передвать только необезательные параметры
// 4. context.WithValue - передача необезательных параметров
// 5. context.WithTimeout - установка таймаута
// 6. context.WithCancel - установка отмены
// 7. ctx всегда передается первым аргументом в функции
// 8. context.WithDeadline - установка таймаута 2

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second * 2))
	defer cancel()

	parse(ctx)
}

func parse(ctx context.Context){
	ctx, _ = context.WithTimeout(ctx, time.Second * 3)

	for{
		select{
		case <- time.After(time.Second * 2):
			fmt.Println("parsing complete")
			return 
		case <- ctx.Done():
			fmt.Println("deadline exceeded")
			return 
		}
	}
}