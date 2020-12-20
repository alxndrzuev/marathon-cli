# marathon-cli
__Marathon Explorer__ - command line tool for getting information from a Mesos+Marathon cluster

## Usage
marathon_explorer [global options] command [command options] [arguments...]

## Commands
__applications__ - get full info about applications  
__help, h__ - Shows a list of commands or help for one command

## Global options:
__--url value, -u value__ - marathon url  
__--user value, -U value__ - marathon basic auth user  
__--password value, -p value__ - marathon basic auth password  
__--project value, -P value__ - filter information by project. For filter by multiple projects use multiple flags, e.g. `-P project1 -P project2`  
__--instances value, -i value__ - filter information by instances count (default: 0)  
__--image value, -I value__ - filter information by image substring  
__--help, -h__ - show help (default: false)  

## Example of usage
`cmd -url=http://mesoshost/marathon -U dev -p 123456 -image config-server applications`  
This command will show all applications whose image contains the substring "config-server".

Example of output:
```bigquery
+---------------+-----------------------------+--------------------------------------------------+-----------+-----+--------+
|     NAME      |           PROJECT           |                      IMAGE                       | INSTANCES | CPU | MEMORY |
+---------------+-----------------------------+--------------------------------------------------+-----------+-----+--------+
| config-server | project1                    | eco.binary.myregistry.ru/config-server:2.5.2     |         1 | 0.1 |  512.0 |
| config-server | project2                    | eco.binary.myregistry.ru/config-server:2.5.2     |         1 | 0.1 |  384.0 |
| config-server | project3                    | eco.binary.myregistry.ru/config-server:2.4.0     |         1 | 0.3 |  512.0 |
+---------------+-----------------------------+--------------------------------------------------+-----------+-----+--------+
```
