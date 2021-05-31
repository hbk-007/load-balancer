package models

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	URL          *url.URL
	IsDead       bool
	lock         *sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	Servers        []*Server
	ServerCount    int64
	LastServerUsed int64
}

func NewServerpool() *ServerPool { return new(ServerPool) }

func (serverPool *ServerPool) RegisterServer(server *Server) {
	serverPool.Servers = append(serverPool.Servers, server)
	serverPool.ServerCount++
}

func GetNewServer(path string) *Server {
	server := new(Server)
	server.URL, _ = url.Parse(path)
	server.IsDead = false
	server.ReverseProxy = httputil.NewSingleHostReverseProxy(server.URL)
	server.lock = new(sync.RWMutex)
	return server
}

func RoundRobinScheduler(pool *ServerPool) *Server {
	var counts int
	for i := pool.GetNextAvailableServer(); ; i++ {
		if counts > int(pool.ServerCount) {
			return nil
		}
		fmt.Println("i is: ", i)
		idx := (i) % pool.ServerCount
		fmt.Println("next available server is: ", idx)
		if !pool.Servers[idx].GetHealth() {
			pool.UpdateLastServerUsed(idx)
			return pool.Servers[idx]

		}
		counts++
	}
}

func (pool *ServerPool) GetNextAvailableServer() int64 {
	return ((pool.LastServerUsed + 1) % pool.ServerCount)
}

func (pool *ServerPool) UpdateLastServerUsed(idx int64) {
	pool.LastServerUsed = idx
	fmt.Println("laserver used is updated to : ", pool.LastServerUsed)
}

func (server *Server) SetHealth(status bool) {
	server.lock.Lock()
	server.IsDead = status
	server.lock.Unlock()
}

func (server *Server) GetHealth() bool {
	server.lock.RLock()
	health := server.IsDead
	server.lock.RUnlock()
	return health
}
