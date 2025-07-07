# LOL-CHAMPIONS-API

Just a REST API app for CRUD' ing informations related to League of Legends champions, written with Go. 

## Installation
Before using, you must have Go and Docker. Of course there will be no data in database but don't worry, you can fetch latest hero informations from LOL server with [this repo](https://github.com/fukaraca/JSON-to-mongoDB) 

`
go get -u github.com/fukaraca/lol-champions-api 
`

Dockerized MongoDB containers can be initialized from app folder with commands depending container status:

`docker compose up` or `docker compose start`


Since our containers are running: 

`go run main.go`

Ready to go...

## Quick Start
You can Bash it or use POSTMAN. (Portal column at the table points to related POSTMAN tab eg params, headers, body, raw, json etc)

### To retrieve hero list
If you want to check hero list:

|Method| URL                                       | Portal  |
|---|-------------------------------------------|---------|
|GET| http://localhost:8080/rest/champions/list |NA|

### Query Hero and Informations
You can query by two options. Sample query URL's given in the table: 
- Either query a hero with its name and take results of desired fields.For brevity, comma is used for seperation so repetition of `q=vals` is not needed. 
- Or make a conditional query by key, operand and value.


|Method| URL                                                                        | Portal |
|---|----------------------------------------------------------------------------|--------|
|GET| http://localhost:8080/rest/champions?name=Jinx&q=lore,stats.hp,spells.name | Params |
|GET| http://localhost:8080/rest/champions?key=stats.mp&op=gt&val=250            | Params |

- First query URL returns `Lore`,`Health Point`and `Spells` for hero named `Jinx`
- Second query URL returns hero names who has `gt`(greater than) 250 `Mana point`

Field's can be found at lib/model.go as json tagged. It is same as default LOL JSON return.
Operands for second query is same for MongoDB but with no $ sign. For example; `eq`, `gt`, `lt`...


### Create New Champion
You can use both raw JSON and Form-Data tab but for detailed creation JSON way must be picked.

| Method | URL                                              | Portal          |
|--------|--------------------------------------------------|-----------------|
| POST   | http://localhost:8080/rest/champions             | Body/form-data  |
| POST   | http://localhost:8080/rest/champions/createWJSON | Body/raw - JSON |

### Update Hero Informations
You can revoke, boost or nerf you favorite champion or change any field you want to.  Total of 4 Body/form-data inputs are required.
These are `name` indicates hero name to be changed,`key` is for field that will be modified,`op` is operand obviously and `val` is new value .


| Method | URL                                              | Portal          |
|--------|--------------------------------------------------|-----------------|
| PATCH  | http://localhost:8080/rest/champions             | Body/form-data  |


### Delete Champion
You can also delete any champion you want.

| Method | URL                                              | Portal          |
|--------|--------------------------------------------------|-----------------|
| DELETE | http://localhost:8080/rest/champions             | Body/form-data  |


