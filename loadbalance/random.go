/**
 * @Author: lenovo
 * @Description:
 * @File:  loadbalance
 * @Version: 1.0.0
 * @Date: 2023/05/27 16:50
 */

package loadbalance

import (
	"context"
	"etcd/registry"
	"math/rand"
)

type RandomBalance struct{}

func (r *RandomBalance) Name() string {
	return "random"
}
func (r *RandomBalance) Select(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error) {
	if len(nodes) == 0 {
		err = ErrNotHaveNodes
		return
	}
	var totalWight int
	for _, val := range nodes {
		if val.Weight == 0 {
			val.Weight = DefaultWright
		}
		totalWight += val.Weight
	}
	curWeight := rand.Intn(totalWight)
	curIndex := -1
	for index, node := range nodes {
		curWeight -= node.Weight
		if curWeight < 0 {
			curIndex = index
			break
		}
	}

	if curIndex == -1 {
		err = ErrNotHaveNodes
		return

	}
	node = nodes[curIndex]
	return
}
