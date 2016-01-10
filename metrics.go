package main

// gaugeDefinition is used to programmatically create prometheus.Gauge metrics
type gaugeDefinition struct {
	id   int
	name string
	desc string
	key  string
}

// counterVecDefinition is used to programmatically create prometheus.CounterVec metrics
type counterVecDefinition struct {
	id       int
	name     string
	desc     string
	label    string
	labelMap map[string]string
}

// Mappings used in CounterVec definitions
// Maps PowerDNS stats key name to Prometheus label value
var (
	rTimeLabelMap = map[string]string{
		"answers0-1": "0-1ms",
		"answers1-10": "1-10ms",
		"answers10-100": "10-100ms",
		"answers100-1000": "100-1000ms",
		"answers-slow": ">1000ms",
		"packetcache-hits": "0ms",
	}

	rCodeLabelMap = map[string]string{
		"servfail-answers": "servfail",
		"nxdomain-answers": "nxdomain",
		"noerror-answers": "noerror",
	}

	exceptionsLabelMap = map[string]string{
		"resource-limits": "resource-limit",
		"over-capacity-drops": "over-capacity-drop",
		"unreachables": "ns-unreachable",
		"outgoing-timeouts": "outgoing-timeout",
	}

	qTypeLabelMap = map[string]string{
		"recursing-questions": "recurse",
		"tcp-queries": "tcp",
		"answers10-100": "10-100ms",
		"answers100-1000": "100-1000ms",
		"answers-slow": ">1000ms",
		"packetcache-hits": "0ms",
	}
)

// PowerDNS recursor metrics definitions
var (
	recursorGaugeDefs = []gaugeDefinition{
		gaugeDefinition{1, "latency_avg_microseconds", "Question-to-answer latency average in microseconds.", "qa-latency"},
		gaugeDefinition{2, "concurrent_queries", "Number of concurrent queries.", "concurrent-queries"},
		gaugeDefinition{3, "cache_size", "Number of entries in the cache.", "cache-entries"},
	}
	
	recursorCounterVecDefs = []counterVecDefinition{
		counterVecDefinition{
			1, "incoming_queries_total", "Total number of incoming queries by network.", "net",
			map[string]string{"questions": "udp", "tcp-questions": "tcp"},
		},
		counterVecDefinition{
			2, "outgoing_queries_total", "Total number of outgoing queries by network.", "net",
			map[string]string{"all-outqueries": "udp", "tcp-outqueries": "tcp"},
		},
		counterVecDefinition{
			3, "cache_lookups_total", "Total number of cache lookups by result.", "result",
			map[string]string{"cache-hits": "hit", "cache-misses": "miss"},
		},
		counterVecDefinition{4, "answers_rcodes_total", "Total number of answers by response code.", "rcode", rCodeLabelMap},
		counterVecDefinition{5, "answers_rtime_total", "Total number of answers by response time.", "responsetime", rTimeLabelMap},
		counterVecDefinition{6, "exceptions_total", "Total number of exceptions by type.", "type", exceptionsLabelMap},
	}
)

// PowerDNS authoritative server metrics definitions
var (
	authoritativeGaugeDefs = []gaugeDefinition{
		gaugeDefinition{1, "latency_avg_microseconds", "Question-to-answer latency average in microseconds.", "latency"},
		gaugeDefinition{2, "packet_cache_size", "Number of entries in the packet cache.", "packetcache-size"},
		gaugeDefinition{3, "signature_cache_size", "Number of entries in the signature cache.", "signature-cache-size"},
		gaugeDefinition{4, "key_cache_size", "Number of entries in the key cache.", "key-cache-size"},
		gaugeDefinition{5, "metadata_cache_size", "Number of entries in the metadata cache.", "meta-cache-size"},
		gaugeDefinition{6, "qsize", "Number of packets waiting for database attention.", "qsize-q"},
	}
	authoritativeCounterVecDefs = []counterVecDefinition{
		counterVecDefinition{
			1, "queries_total", "Total number of queries by network.", "net",
			map[string]string{"tcp-queries": "tcp", "udp-queries": "udp"},
		},
		counterVecDefinition{
			2, "answers_total", "Total number of answers by network.", "net",
			map[string]string{"tcp-answers": "tcp", "udp-answers": "udp"},
		},
		counterVecDefinition{
			3, "recursive_queries_total", "Total number of recursive queries by status.", "status",
			map[string]string{"rd-queries": "requested", "recursing-questions": "processed", "recursing-answers": "answered", "recursion-unanswered": "unanswered"},
		},
		counterVecDefinition{
			4, "update_queries_total", "Total number of DNS update queries by status.", "status",
			map[string]string{"dnsupdate-answers": "answered", "dnsupdate-changes": "applied", "dnsupdate-queries": "requested", "dnsupdate-refused": "refused"},
		},
		counterVecDefinition{
			5, "packet_cache_lookups_total", "Total number of packet-cache lookups by result.", "result",
			map[string]string{"packetcache-hit": "hit", "packetcache-miss": "miss"},
		},
		counterVecDefinition{
			6, "query_cache_lookups_total", "Total number of query-cache lookups by result.", "result",
			map[string]string{"query-cache-hit": "hit", "query-cache-miss": "miss"},
		},
		counterVecDefinition{
			7, "exceptions_total", "Total number of exceptions by type.", "type",
			map[string]string{"servfail-packets": "servfail", "timedout-questions": "timeout", "udp-recvbuf-errors": "recvbuf-error", "udp-sndbuf-errors": "sndbuf-error"},
		},
	}
)

// PowerDNS Dnsdist metrics definitions
var (
	dnsdistGaugeDefs = []gaugeDefinition{}
	dnsdistCounterVecDefs = []counterVecDefinition{}
)
