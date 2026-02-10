package collector

/*
* A Collector is a time-deliver service running that us used for data collection and aggregations.
* A collector can be a Active or Passive. (1) Actively initiating actions that create or capture data or
* (2) passively collecting data from an existing source without initiate events to generate that data. For example, a service
* that pings is an active collector vs a collection that log BMP data which is done passively.
* A collector can be both active and passive. A way to thing of that signature is that an a collection that is both active
* and passive has both a input and out:http.ResponseWriter, r *http.Requestput signature or have input and output channel that are not control/management channels.
* Collectors can be stateless or stateful. They can also be on. We want to differential Collector which are support to be relatively
* long-lived from Middleware. Middleware can sit in-between data pipelines aor intermediate data following to a passive collector.
*
*  Remeber that DOME is serving as a proxy for the collectors. So all requests are sent to the DOME/collector endpoint.
*  We write a  for collectors. So DOME is the centralized way of managing everything.
 */

type Collector struct {
	Image   string
	Port    int
	Name    string
	URL     string
	Catalog map[string]int
}

func (c *Collector) collect() {}

func (c *Collector) run() {}

func (c *Collector) initialize() {}

func (c *Collector) scheduled_collect() {}

// start a colelciton event associated to an id
// scheduled the execution of the events (discrete collections)
// continuous collection start an event-loop that start a continous collection.
// Collect exposes a set of apis for getting accessing the underlying data
//
