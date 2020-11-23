# Configuration Guide

Refer to this guide for creating the right kind of configuration perfectly apt for the kind of notifications you want to receive from whatsupstream.

## Description

Nobody wants their desktop to be submerged in a flood of notifications.

Hence, you can tweak whatsupstream to tell it exactly what kind of issues for what repositories you want to be notified about.

The configuration is simple and written in YAML.

Currently, whatsupstream only supports raising notifications of issue-related activity for repositories. 

So, whatsupstream's configuration is supposed to contain a list of issue-configs where each item of it tells condition for certain kind of notification.

For example, one element of it can be responsible for telling whatsupstream to look into [Kubernetes Repository](https://github.com/kubernetes/kubernetes) and notify whenever any issue with "good first issue" and "sig/node" labels is created.

Another element of it can be responsible for telling whatsupstream to look into [KinD repository](https://github.com/kubernetes-sigs/kind) and notify whenever any issue with "kind/documentation" label and created/raised by the user "yashvardhan-kukreja" comes up.

And whatsupstream polls github at a timely rate looking for any updates/notifications.

## Configuration Fields

Please create the configuration with the following fields and save it as a .yaml file.

### Fields on top of the config
| Field | Description |
|------|-------------|
| `issue-configs` | (object) List of issue-related conditions for different kinds of notifications. (Refer to the above description for examples) |
| `polling-rate` | (integer) Rate in seconds at which whatsupstream polls github for any changes. Default: 60 seconds |

<br>

### Fields under each element of `issue-configs`
| Field | Description |
|------|-------------|
| `repository-url` | (string) URL of the repository to which the issues belong. MANDATORY field. |
| `labels` | ([]string) list of labels which must be there in the issues you want to be notified about. Leaving this empty will make whatsupstream notify about all kinds of issues|
| `assignee` | (string) the user who must be assigned to the issues you want to be notified about. Optional. |
| `creator` | (string) the user who must be the creator of the issues you want to be notified about. Optional. |
| `closed` | (bool) True: Fetch both open and closed issues. False: Fetch only open issues. Optional. Default: False |
| `since` | (string) timestamp in the format "yyyy-mm-ddTHH:MM:SSZ". `since` denotes the timestamp for which any issues created after it, will be even considered for notifications. Optional. Default: timestamp of current_time - 24hrs. For example: "2020-02-21T22:19:38Z"  |
| `max-issues-count` | (int) Say, max-issues-count is 10. If the list of issues finally for notifications are more than 10, then only top 10 of them (as per their creation time) will be considered for notitification. This will save the user's desktop from flooding. Optional. Default: 10 |
| `silent-mode` | (bool) True: A silent notification without any sound will be raised in user's desktop. False: A notification with the default notification sound will be raised. Optional. Default: False |

## Examples

Full configuration with all fields:
```yaml
polling-rate: 30

issue-configs:
- repository-url: "https://github.com/kubernetes/kubernetes"
  labels:
  - "good first issue"
  - "sig/node"
  creator: "yashvardhan-kukreja"
  assignee: "alex123"
  closed: false
  since: "2020-01-02T15:04:05Z"
  max-issues-count: 20
  silent-mode: true
- repository-url: "https://github.com/aquasecurity/trivy"
  labels:
  - "good first issue"
  creator: "bob456"
  assignee: "alex123"
  closed: true
  since: "2019-01-05T15:04:05Z"
  max-issues-count: 30
  silent-mode: false
```

Configuration with some optional fields filled and some not:
```yaml
polling-rate: 15

issue-configs:
- repository-url: "https://github.com/kubernetes/kubernetes"
  labels:
  - "good first issue"
  max-issues-count: 10
  silent-mode: true
```

Configuration filled with no optional fields and only mandatory fields:
```yaml
issue-configs:
- repository-url: "https://github.com/kubernetes/kubernetes"
- repository-url: "https://github.com/moby/moby"
```