package proxy

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"fmt"
	// "github.com/gladmo/film-info/tools"
)

// to do, proxy service
// pre host set a proxy list
// can set use times pre hour

var PoolLength = 10
var UseTimes int64 = 135
var OtherProxyPool = "https://proxy.2pm.me"

const (
	GET_PRO = "https://proxy.2pm.me/get"
	DEL_PRO = "https://proxy.2pm.me/delete?proxy=%s"
	ALL_PRO = "https://proxy.2pm.me/get_all"
	REF_PRO = "https://proxy.2pm.me/refresh"
)

// 限制单ip每小时150次

type Proxy struct {
	Ip      string
	Endtime time.Time
	Times   int64
}

type ProxyPool struct {
	Proxys map[string]Proxy
	mu     sync.Mutex
}

var Proxys ProxyPool = ProxyPool{Proxys: make(map[string]Proxy)}

func (_ *Proxy) GetProxy() (result Proxy) {
	Proxys.mu.Lock()

	// delete times is zero
	for k, v := range Proxys.Proxys {
		if v.Times <= 0 {
			fmt.Println("This is used ", UseTimes, ", ", v.Ip)
			delete(Proxys.Proxys, k)
		}
	}

	if len(Proxys.Proxys) < PoolLength {
		count := PoolLength - len(Proxys.Proxys)
		for i := 0; i < count; i++ {

			// is repeat
			ok := true
			ip := ""
			for ok {
				ip = string(new(Proxy).ScrapyOne())
				_, ok = Proxys.Proxys[ip]
			}

			Proxys.Proxys[ip] = Proxy{
				Ip:      ip,
				Endtime: time.Now(),
				Times:   UseTimes,
			}
		}
	}

	i := 0
	keys := make([]string, len(Proxys.Proxys))
	for k := range Proxys.Proxys {
		keys[i] = k
		i++
	}

	key := keys[rand.Intn(len(keys))]
	result = Proxys.Proxys[key]

	result.Times--

	Proxys.Proxys[key] = result

	Proxys.mu.Unlock()

	return result
}

var repeat int = 0

func (_ *Proxy) ScrapyOne() (ip []byte) {

	res, err := http.Get(GET_PRO)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		repeat++

		// error 100 times, panic it
		if repeat >= 100 {
			panic("proxy server error!")
		}
		return new(Proxy).ScrapyOne()
	}

	ip, _ = ioutil.ReadAll(res.Body)

	repeat = 0

	return
}

func (_ *Proxy) DeleteOne(ip string) {

	// delete from ProxyPool
	if _, ok := Proxys.Proxys[ip]; ok {
		delete(Proxys.Proxys, ip)
	}

	res, _ := http.Get(fmt.Sprintf(DEL_PRO, ip))
	defer res.Body.Close()

	// r, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(fmt.Sprintf(DEL_PRO, ip))
	// fmt.Println(string(r))
}
