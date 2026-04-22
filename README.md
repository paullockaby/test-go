> [!IMPORTANT]
> If you are viewing this repository in GitHub, know that GitHub is [a mirror of the original repository](https://code.lockaby.org/public/test-go).

# test-go

This is a testing ground for building a Go project with all of the bells and whistles of a build pipeline using GitHub Actions.

![GitHub License](https://img.shields.io/github/license/paullockaby/test-go)
![GitHub Release](https://img.shields.io/github/v/release/paullockaby/test-go)

[![Mastodon Follow](https://img.shields.io/mastodon/follow/106882571030731815?domain=https%3A%2F%2Funcontrollablegas.com)](https://uncontrollablegas.com/@paul)

## Table of contents

* [Introduction](#introduction)
* [Quick start](#quick-start)
* [Usage](#usage)
* [Known issues and limitations](#known-issues-and-limitations)
* [Getting help](#getting-help)
* [Contributing](#contributing)
* [License](#license)
* [Acknowledgments](#acknowledgments)

## Introduction

TODO

## Quick Start

First you will want to generate an SSH deploy key for the repository, like this:

```
ssh-keygen -t ed25519 -C commitizen -f deploy_key -P ""
```

This will create two files: `deploy_key` and `deploy_key.pub`. **DO NOT COMMIT THESE.** Save them off somewhere for future reference.

Once you have the SSH keys, take the `deploy_key.pub` file and add it as a deploy key to the repo and make sure that it has write access. Take the private key, aka the non `.pub` file, and create a secret called `DEPLOY_KEY` in the repo. This will be used by commitizen to push commits with changelog updates to the `main` branch during a release.

Next we need to set a branch ruleset for the repository. Rulesets for branches only work on public repos. Or, if you have paid for a pro account or team plan or enterprise plan, then it will work on private repos, too.

The rule that we're going to create is will look like this:

* Ruleset Name: default
* Enforcement status: Active
* Bypass list: deploy keys
* Target branches: Default
* Rules: uncheck everything except these:
  * Restrict deletions
  * Require a pull request before merging
    * Required approvals: your call but if you're a solo developer set it to zero
    * Dismiss stale pull request approvals when new commits are pushed
    * Require review from Code Owners
  * Require status checks to pass
    * Require branches to be up to date before merging
  * Block force pushes

So to sum up:

1. Create the branch rule as described above.
2. Create the deploy key as described above.
3. Tweak the workflows in `.github/workflows` to your liking such that they correctly run lints and tests and builds.
4. Try making a pull request and watch the tests run.
5. Merge the pull request and try creating a release.

What is next? In order to use this repository more effectively there are things that you should do locally. These include:

1. Install pre-commit hooks into your local copy of this repository. You can do that by running this: `make pre-commit`. This should push your development "left" and keep from wasting CI time on failed linting.
2. Make changes only in branches. The branch rule created above will not accept pushes to `main`.
3. Write commit messages only using convential commit messages. Make sure that the SUBJECT LINE of the pull request describes what the pull request does because that is the only thing that will appear in the changelog.I
4. When you are ready to make a release, go to GitHub Actions page and choose "Create Releases" and then click "Run workflow" and run the workflow against the `main` branch. You will see, based on the commit messages in your pull requests, an update to CHANGELOG.md, a new tag, a GitHub Release, plus whatever you put into the `release-build.yaml` action.

## Usage

See [Quick Start](#quick-start).

## Known issues and limitations

There are no known issues or limitations at this time.

## Getting help

Please use GitHub Issues to raise bugs or feature requests or to otherwise communicate with the project. For more details on how to get help, please read the [contributing](#contributing) section.

## Contributing

This project welcomes contributions! Please be cognizant of the [code of conduct](CODE_OF_CONDUCT.md) when interacting with or contributing to this project. You may contribute in many ways:

* By filing bug reports. Please use GitHub Issues to submit any bugs that you may encounter.
* By filing feature requests. Be aware that we may not implement every feature request but we will evaluate them and provide feedback.
* By submitting pull requests. Please use GitHub to submit pull requests.
* By writing documentation. If you see something that could be explained better or is not explained at all, please submit a pull request to update the documentation. Alternatively, just submit an issue describing what is unclear and how it could be more clear and we will endeavor to update the documentation.

If you choose to submit a pull request please follow these guidelines:

* Please provide a clean, concise title for your pull request and a clear description of what you are changing so that it may be evaluated more effectively.
* Limit any pull request to a single change or the minimum number of changes necessary to achieve the feature or bug fix. Many smaller pull requests are preferred over fewer, larger pull requests.
* No pull request will be reviewed unless and until the linter and the tests pass. You can the linter and the tests locally and we encourage you to do so.
* Not every change can or may be accepted. If you are uncertain whether your pull request would be accepted then please open an issue before beginning work to discuss what you would like to do.
* This project is licensed using the Apache License and that by submitting code you accept that your code will be licensed the same.

Again, please use GitHub Issues and GitHub Pull Requests to communicate with the project. It is the fastest and most effective way to be heard.

If you have security feedback you can reach out to [contact@paullockaby.com](mailto:contact@paullockaby.com) to raise your security finding in a confidential manner so that we may provide a fix when the vulnerability is made public. If you are not sure that your feedback is security related please err on the side of caution and send the email. The worst that will happen is you will be asked to create a GitHub Issue.

## License

Copyright &copy; 2024 Paul Lockaby. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at [http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

## Acknowledgements

I took the format for this README file from [Michael Hucka's READMINE repository](https://github.com/mhucka/readmine).
