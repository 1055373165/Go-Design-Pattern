package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Fetcher 负责从远端服务器拉取最新数据
// 并在拉取到最新数据后通知 Watcher
type Fetcher interface {
	Fetch()
	AddWatcher(Client) chan struct{}
	DeleteWatcher(Client)
}

type PaladinFetcher struct {
	mm     sync.Map
	update chan struct{} // 接收更新信号
}

func (pf *PaladinFetcher) AddWatcher(client Client) chan struct{} {
	if _, ok := pf.mm.Load(client.GetName()); ok {
		fmt.Printf("%s 已经注册过了，不要重复监听...", client.GetName())
		return nil
	}

	ch := make(chan struct{})
	pf.mm.Store(client.GetName(), ch)
	return ch
}

func (pf *PaladinFetcher) DeleteWatcher(client Client) {
	if _, ok := pf.mm.Load(client); ok {
		pf.mm.Delete(client.GetName())
		return
	}

	return
}

func (pf *PaladinFetcher) Fetch() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var newMap = map[string]string{
		"1": strconv.FormatInt(int64(r.Intn(1000)), 10),
		"2": strconv.FormatInt(int64(r.Intn(1000)), 10),
		"3": strconv.FormatInt(int64(r.Intn(1000)), 10),
	}

	fmt.Println("未更新前的 map 数据为:")
	printM()
	// 模拟更新本地存储
	update(newMap)
	fmt.Println("更新后的 map 数据为:")
	printM()
}

func (pf *PaladinFetcher) notify() {
	pf.mm.Range(func(key, value interface{}) bool {
		if singlechan, ok := value.(chan struct{}); ok {
			singlechan <- struct{}{}
		}
		// Range 方法的回调函数返回一个布尔值，表示是否继续遍历
		return true
	})
	fmt.Println("通知信号发送完毕")
	fmt.Println("=========================")
}

// client.go
type Client interface {
	ConnectToPaladinSDK()
	WatchPaldin()
	GetName() string
}

type AKSOClient struct {
	name    string
	fetcher Fetcher // 在同一个 go 文件中，因此使用组合模拟
}

func (ak *AKSOClient) ConnectToPaladinSDK() {
	fmt.Println("akso connected to paladin sdk")
}

func (ak *AKSOClient) WatchPaldin() {
	ch := ak.fetcher.AddWatcher(ak)
	for {
		select {
		case <-ch:
			fmt.Printf("%s收到更新信号\n", ak.name)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (ak *AKSOClient) GetName() string {
	return ak.name
}

type MemoClient struct {
	name    string
	fetcher Fetcher
}

func (ak *MemoClient) ConnectToPaladinSDK() {
	fmt.Println("memo connected to paladin sdk")
}

func (me *MemoClient) WatchPaldin() {
	ch := me.fetcher.AddWatcher(me)
	for {
		select {
		case <-ch:
			fmt.Printf("%s收到更新信号\n", me.name)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (me *MemoClient) GetName() string {
	return me.name
}

var database = map[string]string{
	"1": "123",
	"2": "234",
	"3": "345",
}

func update(newMap map[string]string) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	for k, v := range newMap {
		database[k] = v
	}
	fmt.Println("最新数据更新完毕...")
}

func printM() {
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	for k, v := range database {
		fmt.Printf("%s:%s\n", k, v)
	}
}

var rwmutex sync.RWMutex

func main() {
	paladinFetcher := &PaladinFetcher{
		update: make(chan struct{}),
	}

	akso := &AKSOClient{
		name:    "ASKO",
		fetcher: paladinFetcher,
	}
	go akso.WatchPaldin()

	memo := &MemoClient{
		name:    "Memo",
		fetcher: paladinFetcher,
	}
	go memo.WatchPaldin()

	// 模拟 paladin server 更新
	t := time.NewTicker(250 * time.Millisecond)
	go func() {
		for {
			select {
			case <-t.C:
				paladinFetcher.update <- struct{}{}
			}
		}
	}()

	for {
		select {
		case <-paladinFetcher.update:
			paladinFetcher.Fetch()
			paladinFetcher.notify()
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}
