package main

type RoutingFile struct {
	Routing struct {
		DomainStrategy string `json:"domainStrategy"`
		DomainMatcher  string `json:"domainMatcher"`
		Balancers      []struct {
			Tag      string   `json:"tag"`
			Selector []string `json:"selector"`
		} `json:"balancers"`
		Rules []struct {
			InboundTag  []string `json:"inboundTag"`
			Domain      []any    `json:"domain,omitempty"`
			BalancerTag string   `json:"balancerTag,omitempty"`
			Type        string   `json:"type"`
			OutboundTag string   `json:"outboundTag,omitempty"`
			IP          []string `json:"ip,omitempty"`
		} `json:"rules"`
	} `json:"routing"`
	BurstObservatory struct {
		SubjectSelector []string `json:"subjectSelector"`
		PingConfig      struct {
			Destination  string `json:"destination"`
			Interval     string `json:"interval"`
			Connectivity string `json:"connectivity"`
			Timeout      string `json:"timeout"`
			Sampling     int    `json:"sampling"`
		} `json:"pingConfig"`
	} `json:"burstObservatory"`
}

type OutboundConfigBase struct {
	Tag      string `json:"tag"`
	Protocol string `json:"protocol"`
}

type OutboundsFile struct {
	Outbounds string `json:"outbounds"`
}
