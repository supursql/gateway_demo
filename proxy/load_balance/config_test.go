package load_balance

import "testing"

func TestNewLoadBalanceObserver(t *testing.T) {
	moduleConf, err := NewLoadBalanceZkConf("rs_server", "/rs_server", []string{"127.0.0.1:2181"}, map[string]string{"127.0.0.1:2003": "20"})
	if err != nil {
		panic(err)
	}
	LoadBalanceObserver := NewLoadBalanceObserver(moduleConf)
	moduleConf.Attach(LoadBalanceObserver)
	moduleConf.UpdateConf([]string{"122.11.11"})
}
