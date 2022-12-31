package domain

type Log struct {
	Service       string `bson:"service" json:"service"`
	ContainerName string `bson:"container_name" json:"container_name"`
	Time          string `bson:"time" json:"time"`
	RemoteIP      string `bson:"remote_ip" json:"remote_ip"`
	Host          string `bson:"host" json:"host"`
	Method        string `bson:"method" json:"method"`
	Uri           string `bson:"uri" json:"uri"`
	UserAgent     string `bson:"user_agent" json:"user_agent"`
	Status        string `bson:"status" json:"status"`
	Latency       string `bson:"latency" json:"latency"`
	LatencyHuman  string `bson:"latency_human" json:"latency_human"`
}
