package redispubsub

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
)

//暂只支持单节点模式的redis.Client(不支持集群redis.ClusterClient)

//redis client hook
type NewRedisClientFunc func() (*redis.Client, error)
type CloseRedisClientFunc func() error

//redis消息通知
type RedisNotify struct {
	addr           string
	redisOpenHook  NewRedisClientFunc
	redisCloseHook CloseRedisClientFunc
	stop           chan struct{}
}

func NewRedisNotify(addr string, funcs ...NewRedisClientFunc) (*RedisNotify, error) {
	var fn NewRedisClientFunc
	if len(funcs) > 0 {
		log.Println("using hook to create client")
		fn = funcs[0]
	} else {
		if addr == "" {
			return nil, errors.New("no redis conn create hook function and redis addr is null")
		}
		fn = func() (*redis.Client, error) {
			client := redis.NewClient(&redis.Options{
				Addr: "addr",
			})
			return client, nil
		}
	}
	return &RedisNotify{
		addr:          addr,
		redisOpenHook: fn,
		stop:          make(chan struct{}),
	}, nil
}

//获取redis发布通道
func (r *RedisNotify) GetPubChannel(pubKey string, channelsize int) (chan<- []byte, error) {

	if channelsize < 0 {
		channelsize = 0
	}

	mq := make(chan []byte, channelsize)
	client, err := r.redisOpenHook()
	if err != nil {
		return nil, err
	}
	//测试连通性
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case data, ok := <-mq:
				if !ok {
					panic("pub channel is closed")
				}
				err := client.Publish(pubKey, string(data)).Err()
				if err != nil {
					log.Println(err)
				}
			case <-r.stop:
				log.Println("redis notify is stop")
				close(mq)
				return
			}
		}
	}()
	return mq, nil
}

//获取redis订阅通道
func (r *RedisNotify) GetSubChannel(subKey string, channelsize int) (<-chan []byte, error) {

	if channelsize < 0 {
		channelsize = 0
	}
	mq := make(chan []byte, channelsize)
	client, err := r.redisOpenHook()
	if err != nil {
		return nil, err
	}
	//测试连通性
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	go func() {
		pubsub := client.Subscribe(subKey)
		defer pubsub.Close()

		for {
			select {
			case <-r.stop:
				log.Println("redis notify is stop")
				close(mq)
				return

			default:
				msg, err := pubsub.ReceiveMessage()
				if err != nil {
					log.Println(err.Error())
					if err.Error() == "redis: client is closed" {
						panic(err)
					} else {
						continue
					}
				}
				switch msg.Channel {
				case subKey:
					log.Printf("receive subscribe message: %s\n", msg.Payload)
					data := []byte(msg.Payload)
					mq <- data
				}
			}

		}
	}()
	return mq, nil
}

func (r *RedisNotify) SubscribeHooks(hooks map[string]func(string)) error {

	hooksLen := len(hooks)

	if hooksLen == 0 {
		return errors.New("hooks length must more than zero")
	}

	client, err := r.redisOpenHook()
	if err != nil {
		return err
	}
	//测试连通性
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}

	subKeys := make([]string, hooksLen)
	for k, _ := range hooks {
		subKeys = append(subKeys, k)
	}

	log.Printf("subkeys => %v\n", subKeys)

	pubsub := client.Subscribe(subKeys...)
	defer pubsub.Close()

	for {
		select {
		case <-r.stop:
			log.Println("redis notify is stop")
			return errors.New("redis notify is stop")

		default:
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				log.Println(err.Error())
				if err.Error() == "redis: client is closed" {
					panic(err)
				} else {
					continue
				}
			}

			if hookfunc, ok := hooks[msg.Channel]; ok {
				log.Printf("receive a hook notify : %s\n", msg.Channel)
				if hookfunc != nil {
					hookfunc(msg.Payload)
				}
			} else {
				log.Printf("receive a bad hook notify : %s\n", msg.Channel)
			}
		}
	}
}

//设置关闭回调
func (r *RedisNotify) SetRedisCloseHook(fn CloseRedisClientFunc) {
	r.redisCloseHook = fn
}

//关闭redis通知
func (r *RedisNotify) Close() {
	r.stop <- struct{}{}
	if r.redisCloseHook != nil {
		r.redisCloseHook()
	}
}
