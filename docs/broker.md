# Eventing RabbitMQ Broker

RabbitMQ _is a Messaging Broker_ - an intermediary for messaging. It gives your applications a common platform to send and receive messages, and your messages a safe place to live until received.

![Eventing RabbitMQ Broker](rabbitmq-knative-broker.png)

## Installation

### Install Eventing RabbitMQ Broker

Install the latest version of the [Operator based Knative RabbitMQ Broker](https://github.com/knative-extensions/eventing-rabbitmq/releases/):

```shell
kubectl apply --filename https://github.com/knative-extensions/eventing-rabbitmq/releases/latest/download/rabbitmq-broker.yaml
```

Or install a specific version, e.g., v0.25.0, run:

```shell
kubectl apply --filename https://github.com/knative-extensions/eventing-rabbitmq/releases/download/v0.25.0/rabbitmq-broker.yaml
```

Or install a nightly version:

```shell
kubectl apply -f https://storage.googleapis.com/knative-nightly/eventing-rabbitmq/latest/rabbitmq-broker.yaml
```

For development purposes or to use the latest from the repository, use [`ko`](https://github.com/google/ko) for installation from a local copy of the repository.

```
ko apply -f config/broker/
```

## Uninstall

### Remove eventing-rabbitmq components and resources

Use `kubectl delete --filename <installation-file>` to remove the components installed during [Installation](#install-eventing-rabbitmq-broker). For example:

```shell
kubectl delete --filename https://github.com/knative-extensions/eventing-rabbitmq/releases/download/v0.25.0/rabbitmq-broker.yaml
```

If `ko` was used to install, can also be used for uninstallation:

```
ko delete -f config/broker/
```

### Remove RabbitMQ Cluster and Topology Operators

To remove RabbitMQ cluster and topology operators, use similar `kubectl delete` commands with the files used for installation:

```
kubectl delete -f https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml
```

### Uninstall Knative Serving and Eventing

Follow the instructions [here](https://knative.dev/docs/install/uninstall/#uninstalling-optional-channel-messaging-layers) to uninstall Knative components.
