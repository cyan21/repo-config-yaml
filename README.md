# repo-config

## About this plugin
This plugin aims to manage Artifactory repositories via YAML files. 
This README shows the expected structure of your plugin's README.

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install repo-config`

Installing a specific version:

`$ jfrog plugin install repo-config@version`

Uninstalling a plugin

`$ jfrog plugin uninstall repo-config`

## Usage
### Commands

* export : 
    - Arguments:
        - repoType - all, local, remote, virtual. Wil generate 1 YAML file per repository type.
    - Flags:
        - pkgType: Filter repositories based on one specific package type
    - Example:
    ```
  $ jfrog repo-config export all --pkgType=docker 
  
  ```

### Environment variables
None.

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).


