package redispubsub

import (
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestRedisPubSub(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:25431",
		Password: "k9Gjj,dZ",
		DB:       0,
	})

	if client == nil {
		t.Fail()
	}

	getClientFunc := func() (*redis.Client, error) {
		return client, nil
	}

	redisNotify, err := NewRedisNotify("", getClientFunc)
	if err != nil {
		t.Fail()
	}

	pubsubkey := "pcmgr_test"

	done := make(chan bool)

	time.AfterFunc(time.Second, func() {
		pmq, err := redisNotify.GetPubChannel(pubsubkey, 1<<10)
		if err != nil {
			t.Fail()
		}
		pmq <- []byte("hello")
		done <- true
	})

	go func() {
		smq, err := redisNotify.GetSubChannel(pubsubkey, 0)
		if err != nil {
			t.Fail()
		}
		data := <-smq
		if string(data) != "hello" {
			t.Fail()
		}
	}()
	<-done
	time.Sleep(5 * time.Second)
	println("done")
}
