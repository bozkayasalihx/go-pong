run:
	go run .

kill: 
	kill -9 $(ps aux | grep "go run ." | grep -v grep | awk '{print $2}')
