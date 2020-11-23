<!-- PROJECT LOGO -->
<br />
<p align="center">
  <img src="./assets/icon.png" height=300>
  <h3 align="center"> :rocket: Whatsupstream :rocket: </h3>

  <p align="center">
    A tool which will keep you up-to-date with the activity (issues) of your favorite Open Source repositories without any hassle or even opening browser.<br>
    Just tweak it according to your wants and it will notify on your desktop, if any issue meeting your criteria (like certain labels/creator) of any OSS repository got created. :sunglasses:
    <br />
    <br />
    ·
    <a href="https://github.com/yashvardhan-kukreja/nail-hacktoberfest/issues/new?assignees=&labels=bug&template=bug_report.md&title=">Report Bug</a>
    ·
    <a href="https://github.com/yashvardhan-kukreja/nail-hacktoberfest/issues/new?assignees=&labels=feature&template=feature_request.md&title=">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Project Setup](#project-setup)
* [Whatsupstream in action](#whatsupstream-in-action)
* [Contributing](#contributing)
* [License](#license)

<!-- ABOUT THE PROJECT -->
## About Whatsupstream

Whatsupstream is a tool (CLI actually) which runs as a background process and tracks the activity associated with any Open Source repositories you tell it. (for now, it tracks issues)

Say, you want to be notified as soon as possible whenever an issue with `good first issue` and `sig/node` labels is created in  [Kubernetes](https://github.com/kubernetes/kubernetes)

Just tell Whatsupstream (configure it, it's easy :P) about it and this tool will simply run behind the scenes and raise a desktop notification on your computer as soon it sees such kind of issue.

Benefit? 

You can pick that issue up instantly before anyone picks it up :P, work on it and be a proud OSS contributor :D

PS: you can configure it for countless repositories as per your liking.

<!-- GETTING STARTED -->
## Getting Started

Here are instructions to setup and use whatsupstream.

### Prerequisites

Just have `make` installed in your computer, rest all will be handled automatically for you.

Behind the scenes, this project entirely runs on GoLang but if you don't have it, no worries, `make` will set it up for you :)

### Project Setup

1. Clone this repo
   * With HTTPS
   ```sh
   git clone https://github.com/yashvardhan-kukreja/whatsupstream.git
   ```
   * With SSH
   ```sh
   git clone git@github.com:yashvardhan-kukreja/whatsupstream.git
   ```
2. Hop into the repository directory
```sh
cd whatsupstream
```
3. And install whatsupstream

    * For installing inside `/usr/local/bin/`
    ```sh
    make install
    ```
    * If you want the whatsupstream binary some other place:
    ```sh
    make install INSTALL_DIR=/path/to/other/dir
    ```
4. Setup your whatsupstream configuration in YAML by referring to [Whatsupstream Configuration Guide](./docs/configuration-guide.md) and save it in `$HOME/.whatsupstream/config.yaml` (preferable) or any other place.
4. Run it as a background process :rocket:

    * If your whatsupstream config is at `$HOME/.whatsupstream/config.yaml`
    ```sh
    make notify
    ```
    * If your whatsupstream config is at some other place, say, `/path/to/other/place/whatsupstream.yaml`
    ```sh
    make notify /path/to/other/place/whatsupstream.yaml
    ```
OR

5. Just run it as a foreground process with:
```sh
whatsupstream notify --config /path/to/config.yaml
```
Whatsupstream runs as a background process. So, to stop all the instances of whatsupstream, just run:
```sh
make stop
```

<!-- USAGE EXAMPLES -->
## Whatsupstream in action

<img src="./assets/usage_1.png" alt="Whatsupstream" width="200"/>

<!-- CONTRIBUTING -->
## Contributing

Refer to the [Developer's Guide](./docs/developer-guide.md) for a detailed explanation of developing over, testing and contributing to whatsupstream.

<!-- LICENSE -->
## License

Distributed under the Apache License Version 2.0. See [LICENSE](https://github.com/yashvardhan-kukreja/whatsupstream/blob/master/LICENSE) for more information.
