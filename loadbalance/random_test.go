/**
 * @Author: lenovo
 * @Description:
 * @File:  random_test.go
 * @Version: 1.0.0
 * @Date: 2023/05/27 16:57
 */

package loadbalance

import (
	"context"
	"etcd/registry"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	balance := &RandomBalance{}
	var nodes []*registry.Node

	weights := [3]int{50, 100, 150}
	for i := 0; i < 4; i++ {
		node := &registry.Node{
			IP:     fmt.Sprintf("127.0.0.%d", i),
			Port:   8080,
			Weight: weights[i%3],
		}
		fmt.Println(node)
		nodes = append(nodes, node)
	}
	countStat := make(map[string]int)
	for i := 0; i < 1000; i++ {
		node, err := balance.Select(context.Background(), nodes)
		assert.NoError(t, err)
		countStat[node.IP]++
	}
	for key, val := range countStat {
		t.Logf("ip:%s count:%d", key, val)
	}
}
