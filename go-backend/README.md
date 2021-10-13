# My Device Store

Credit to https://github.com/eliben/code-for-blog/tree/master/2021/go-rest-servers/gorilla

```sh
# How to start the server
SERVERPORT=4300 go run server.go

# Get
curl http://localhost:4300/devices

# Post
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "iPad2","device_type": "iPad","owner": "Lucy","mac_addr": "AA-BC-CC-DD","ip_addr": "192.168.1.10","start_use_date": "May 6th 2018","is_commonly_used": true}' \
  http://localhost:4300/devices

# Delete
curl -X DELETE http://localhost:4300/devices/1

# Put
curl -H "Content-Type: application/json" -X PUT \
  --data '{"name": "iPad1","device_type": "iPad1","owner": "Lucy22","mac_addr": "AA-BC-CC-DD","ip_addr": "192.168.1.10","start_use_date": "May 6th 2018","is_commonly_used": true}' \
  http://localhost:4300/devices/0 

```