# chaika - The log shipper

* Currently, just support ship log over UDP to GELF (graylog), more log format and adapter will be support soon
* Chaika use consul to register and deregister log info
* Checkout the log4js appender for NodeJS: https://github.com/duythinht/log4js-chaika-appender

    /chaika -help
    Usage of ./chaika:
      -consul-host string
            Consul hostname (default "localhost")
      -consul-port int
            Consul port (default 8500)
      -graylog-host string
            Default graylog host or domain (default "localhost")
      -graylog-port int
            Default graylog port (default 12201)
      -p int
            Port for agent run on (default 2435)
