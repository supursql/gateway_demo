package load_balance

type LbType int

const (
	LbRandom LbType = iota
	LbRoundRobin
	LbWeightRoundRobin
	LbConsistentHash
)

func LoadBalanceFactory(lbType LbType) LoadBalance {
	switch lbType {
	case LbRandom:
		return &RandomBalance{}
	case LbConsistentHash:
		return NewConsistentHashBalance(10, nil)
	case LbRoundRobin:
		return &RoundRobinBalance{}
	case LbWeightRoundRobin:
		return &WeightRoundRobinBalance{}
	default:
		return &RandomBalance{}
	}
}

func LoadBalanceFactoryWithConf(lbType LbType, mConf LoadBalanceConf) LoadBalance {
	switch lbType {
	case LbRandom:
		lb := &RandomBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbConsistentHash:
		lb := NewConsistentHashBalance(10, nil)
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbRoundRobin:
		lb := &RoundRobinBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbWeightRoundRobin:
		lb := &WeightRoundRobinBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	default:
		lb := &RandomBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	}
}
