# hello-frog

## About this plugin
This plugin is a template and a functioning example for a basic JFrog CLI plugin. 
This README shows the expected structure of your plugin's README.

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install repo-cfg`

Installing a specific version:

`$ jfrog plugin install repo-cfg@version`

Uninstalling a plugin

`$ jfrog plugin uninstall repo-cfg`

## Usage
### Commands


goals 
1/ export repo definition inot 4 YAML files : local, remote, virtual, federated 
2/ generate YAML for repo creation
3/ merge YAML files 
4/ convert local to federated 

configuration files 
packageType (for loop)
    goal 1/

removeField
    goal 2/

renameField
    goal 2/

list repo to convert
    goal 3/

* hello
    - Arguments:
        - addressee - The name of the person you would like to greet.
    - Flags:
        - shout: Makes output uppercase **[Default: false]**
        - repeat: Greets multiple times **[Default: 1]**
    - Example:
    ```
  $ jfrog hello-frog hello world --shout --repeat=2
  
  NEW GREETING: HELLO WORLD!
  NEW GREETING: HELLO WORLD!
  ```

### Environment variables
* HELLO_FROG_GREET_PREFIX - Adds a prefix to every greet **[Default: New greeting: ]**

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).


