#1. run server
#2. go build ./
#3. go run ./

curl -X POST -H "Content-Type:application/json" -d "{\"username\": \"isa\", \"points\": 25}" "localhost:8080/points"
curl -X POST -H "Content-Type:application/json" -d "{\"username\": \"yc\", \"points\": 100}" "localhost:8080/points"
curl -X POST -H "Content-Type:application/json" -d "{\"username\": \"cy\", \"points\":66}" "localhost:8080/points"

curl -X GET -H "Content-type: application/json" "localhost:8080/points/isa"

curl -X GET -H "Content-type: application/json" "localhost:8080/leaderboard"