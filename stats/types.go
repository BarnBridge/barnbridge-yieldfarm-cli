package stats

type tokenBalances struct {
	Token            string `header:"token name"`
	TotalBalance     string `header:"total balance"`
	EffectiveBalance string `header:"effective balance"`
}

type poolSizes struct {
	Pool            string `header:"pool"`
	Size            string `header:"pool size"`
	MyShare         string `header:"my share"`
	PotentialReward string `header:"potential reward"`
}
