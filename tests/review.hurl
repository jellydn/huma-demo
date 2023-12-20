POST http://0.0.0.0:8888/reviews
{
  "rating": 10
}
HTTP 422

POST http://localhost:8888/reviews
{
  "author": "IT Man",
  "rating": 5
}
HTTP 201

GET http://localhost:8888/reviews

HTTP/1.1 200

[Asserts]
jsonpath "$.reviews" count >= 1
