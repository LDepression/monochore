/**
 * @Author: lenovo
 * @Description:
 * @File:  loadbalance
 * @Version: 1.0.0
 * @Date: 2023/05/27 16:53
 */

package loadbalance

import (
	"context"
	"errors"
	"etcd/registry"
)

var (
	ErrNotHaveNodes = errors.New("not have nodes")
	DefaultWright   = 100
)

type LoadBalance interface {
	Name() string
	Select(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error)
}
