---
page_title: "Troubleshooting"
subcategory: ""
description: |-
  Troubleshooting tips and techniques for diagnosing issues with the SMC Terraform provider.
---

# Troubleshooting

## Enabling Debug Mode

To help diagnose issues, you can enable debug logging for the Terraform provider. Export the following environment variables before running any Terraform CLI commands. This will generate a detailed log file at /tmp/tf.log:

```sh
export TF_LOG_PATH=/tmp/tf.log
export TF_LOG_PROVIDER=DEBUG
export TF_LOG_SDK=OFF
```

Review the log file for error messages, stack traces, or unexpected behavior. This information is useful for troubleshooting and when reporting issues.

on ubuntu, install it via:

## Capturing HTTP Traffic

Sometimes it is helpful to capture the API traffic between Terraform
and the SMC to diagnose communication or protocol issues.

### Capturing at the TCP Level

If you are using plain HTTP, you can use `tcpflow` to capture TCP
traffic. This is useful for low-level troubleshooting.

On Ubuntu, install `tcpflow` with:

```sh
sudo apt install tcpflow
```

Then, to capture traffic on port 8082:

```sh
sudo tcpflow -i any port 8082 -C | tee /tmp/tcp_flow.txt
```

### Capturing via an HTTP Proxy

You can also use an HTTP proxy tool such as Charles, Burp Suite,
mitmproxy, or http-breakout-proxy to inspect and debug HTTP requests
and responses.

On Linux, set the following environment variable to direct traffic
through your proxy:

```sh
export HTTP_PROXY="http://smcproxy:8080"
```

**Note:** If you are trying to contact the SMC on localhost, the HTTP
proxy won't work as expected. This is due to a limitation in the go
networking library used by the terraform SMC provider.

The workaround is to create an alias for localhost in `/etc/hosts`
(e.g., `127.0.0.1 mysmchost`) and use that alias in your provider
configuration.
