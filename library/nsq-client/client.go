// nsq客户端，目前只实现了消息的发送，使用FAN-OUT模型
// mainChan为主chan
// 为发现的每个nsqd启动一个go routine，从mainCh 读取读取数据并且写入子routine的inCh(带缓冲默认128)
//
//	/----- nsq_publisher.publish(inCh)
//
// mainChan --- ------ nsq_publisher.publishAsync(inCh)
//
//	\----- nsq_publisher.publishMultiAsync(inCh)
package nsq_client

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jxncyjq/lib_stardust/core/system/thread"
	"github.com/nsqio/go-nsq"
	"math/rand"
	"sync"
	"time"
)

// Warning: 非标准nsq client实现，只实现了发送逻辑
type NSQClient interface {
	NodeAddress() []string
	Lookup() error
	CreatePublishers() error
	CreateConsumer() error
	StartProducers(useAsync bool, batchSize int) error
	Send([]byte)
	Shutdown()
}
type Config struct {
	LookupdHTTPAddrs    []string      `json:"lookupd_http_addrs" hcl:"lookupd_http_addrs"`
	TlsConfig           *tls.Config   `json:"tls_config" hcl:"tls_config"`
	AuthSecret          string        `json:"auth_secret" hcl:"auth_secret"`
	ReadTimeout         time.Duration `hcl:"read_timeout" min:"100ms" max:"5m" default:"60s"`
	WriteTimeout        time.Duration `hcl:"write_timeout" min:"100ms" max:"5m" default:"1s"`
	LookupdPollInterval time.Duration `hcl:"lookupd_poll_interval" min:"10ms" max:"5m" default:"60s"`
}
type nsqClient struct {
	sync.Pool
	sync.Mutex
	logger           hclog.Logger
	nodeAddress      []string
	publishers       map[string]*publisher
	mainChan         chan []byte
	topic            string
	cfg              *nsq.Config
	lookupdHTTPAddrs []string
	mainCtx          context.Context
	mainCancel       context.CancelFunc
	useAsync         bool
	batchSize        int
}

func (nc *nsqClient) NodeAddress() []string {
	return nc.nodeAddress
}

func NewNsqClient(cfg *Config, topic string, logger hclog.Logger) (NSQClient, error) {

	if len(topic) == 0 {
		return nil, errors.New(thread.GetCaller(3) + " topic required")
	}
	if len(cfg.LookupdHTTPAddrs) == 0 {
		return nil, errors.New(thread.GetCaller(3) + " lookupd http address required")
	}
	nsqConfig := nsq.NewConfig()
	nsqConfig.TlsConfig = cfg.TlsConfig
	nsqConfig.AuthSecret = cfg.AuthSecret
	nsqConfig.ReadTimeout = cfg.ReadTimeout
	nsqConfig.WriteTimeout = cfg.WriteTimeout
	nsqConfig.LookupdPollInterval = cfg.LookupdPollInterval

	return &nsqClient{
		lookupdHTTPAddrs: cfg.LookupdHTTPAddrs,
		cfg:              nsqConfig,
		mainChan:         make(chan []byte, 819200),
		topic:            topic,
		logger:           logger,
		publishers:       make(map[string]*publisher, 0),
	}, nil
}

// 通过lookupd查找nsqd的地址
func (nc *nsqClient) Lookup() error {
	nc.cfg.UserAgent = fmt.Sprintf("go-nsq/%s", nsq.VERSION)
	lookup := NewLookup(nsq.NewConfig(), nc.lookupdHTTPAddrs, nc.topic)
	nc.nodeAddress = lookup.LookupNsqdAddress()
	if len(nc.nodeAddress) == 0 {
		err := errors.New("lookup nsqd nodes is empty")
		return err
	}
	return nil
}

// 创建Publishers
func (nc *nsqClient) CreatePublishers() error {
	nc.Lock()
	defer nc.Unlock()

	mainCtx, mainCancel := context.WithCancel(context.Background())
	nc.mainCtx = mainCtx
	nc.mainCancel = mainCancel

	if err := nc.Lookup(); err != nil {
		return err
	}
	for _, addr := range nc.nodeAddress {
		//可能是重连，跳过存在的
		_, ok := nc.publishers[addr]
		if ok {
			continue
		}
		producer, err := nsq.NewProducer(addr, nc.cfg)
		if err != nil {
			nc.logger.Error("lookup", "addr", addr, "err", err.Error())
		}
		pub := NewPublisher(nc.topic, &Producer{producer}, nc.logger.Named("publisher"))
		nc.publishers[addr] = pub

	}

	if err := nc.startPublisher(); err != nil {
		return err
	}

	return nil
}

func (nc *nsqClient) Send(data []byte) {
	select {
	case nc.mainChan <- data:
	default:
		time.Sleep(time.Millisecond * 30)
	}
}

// 启动所有producer的处理逻辑
func (nc *nsqClient) StartProducers(useAsync bool, batchSize int) error {
	nc.Lock()
	defer nc.Unlock()

	nc.useAsync = useAsync
	nc.batchSize = batchSize

	go func() {
		for {
			select {
			case data := <-nc.mainChan:
				addr := nc.nodeAddress[rand.Int31n(int32(len(nc.publishers)))]
				pub := nc.publishers[addr]
				pub.Publish(data)
			default:
				time.Sleep(time.Millisecond * 30)
			}
		}
	}()

	nc.autoReconnect(time.Minute)

	return nil
}

func (nc *nsqClient) Shutdown() {
	nc.mainCancel()
}

func (nc *nsqClient) startPublisher() error {
	for addr, pub := range nc.publishers {
		ctx, err := pub.StarPublisher(nc.useAsync, nc.batchSize, nc.mainCtx)
		if nil != err {
			nc.logger.Error("start publisher", "err", err.Error())
			return err
		}
		go func(addr string) {
			select {
			case <-ctx.Done():
				nc.Lock()
				delete(nc.publishers, addr)
				nc.Unlock()
			}
		}(addr)
	}
	return nil
}

func (nc *nsqClient) autoReconnect(duration time.Duration) {
	//定时重连
	go func() {
		ticker := time.NewTicker(duration)
		for true {
			select {
			case <-ticker.C:
				if err := nc.CreatePublishers(); err != nil {
					nc.logger.Error("auto discovery nsqd", "err", err.Error())
				}
			}
		}
	}()
}
