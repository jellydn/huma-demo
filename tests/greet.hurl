GET http://0.0.0.0:8888/greeting/world

HTTP 200
[Asserts]
jsonpath "$.message" == "Hello, world!"
