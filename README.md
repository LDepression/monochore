# monochore
基于etcd封装的微服务框架


## 服务注册

### 服务抽象与服务节点的抽象
```go
// 服务抽象
type Service struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes"`
}

// 服务节点的抽象
type Node struct {
	Id     string `json:"id"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}
```
### 服务注册的插件
现在我们是基于etcd实现的服务注册，未来可以添加consul之类的
我们认为实现了下面五种方法就实现了该插件
```go
// 服务注册插件的接口
type Registry interface {
	//插件的名字
	Name() string
	//初始化
	Init(ctx context.Context, opts ...Option) (err error)
	//服务注册
	Register(ctx context.Context, service *Service) (err error)
	//服务反注册
	Unregister(ctx context.Context, service *Service) (err error)
	//服务发现：通过服务的名字获取服务的位置信息（ip和port列表）
	GetService(ctx context.Context, name string) (service *Service, err error)
}
```
### 插件的管理者
```go
type PluginMgr struct {
	plugins map[string]Registry
	lock    sync.Mutex
}
```
管理者封装了注册插件以及初始化插件，也就是插件自己进行初始化。然后被添加到管理插件的map上去


### etcd作为插件

etcd作为插件，也就是说etcd要实现上面Registry的五个方法
```go
type EtcdRegistry struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service /注册服务异步化，所以使用chan

	value              atomic.Value                //服务发现的时候加锁太影响性能了，直接原子操作
	lock               sync.Mutex                  //这个锁只有穿透到etcd才会加锁
	registryServiceMap map[string]*RegisterService //存下要注册的所有service
}


type RegisterService struct {
	id          clientv3.LeaseID
	service     *registry.Service
	registered  bool // 有没有注册过
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
}
```
etcd要注册的服务必然是不止一个的，所以我们使用一个map再次管理要注册的服务
每次请求必然不可能每次都从etcd请求，这样会导致etcd压力过大，所以我们使用缓存。又因为加锁的话，会影响性能，所以直接使用atomic包下的Value

```go
var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		serviceCh:          make(chan *registry.Service, MaxServiceNum),
		registryServiceMap: make(map[string]*RegisterService, MaxServiceNum),
	}
)

func init() {

	allServiceInfo := &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaxServiceNum),
	}

	etcdRegistry.value.Store(allServiceInfo) //先缓存
	registry.RegisterPlugin(etcdRegistry)    //先初始化管理插件
	go etcdRegistry.run()                    //执行注册
}
```

