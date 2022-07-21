# Learn Go with Saltpay - Exercise 1

[Github Link](https://github.com/saltpay/learn-go-with-salt/blob/master/book/exercise1.md) 



## Part 26

Run the command below to generate the benchmarking data.
```
ab -n 10000 -c 100 -S -H "Authorization: Bearer SUPER_SECRET_API_KEY_1"  http://localhost:8080/add\?num\=42\&num\=5\&num\=32
```

```
Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /add?num=42&num=5&num=32
Document Length:        0 bytes

Concurrency Level:      100
Time taken for tests:   0.591 seconds
Complete requests:      10000
Failed requests:        0
Non-2xx responses:      10000
Total transferred:      910000 bytes
HTML transferred:       0 bytes
Requests per second:    16926.57 [#/sec] (mean)
Time per request:       5.908 [ms] (mean)
Time per request:       0.059 [ms] (mean, across all concurrent requests)
Transfer rate:          1504.22 [Kbytes/sec] received

Connection Times (ms)
              min   avg   max
Connect:        0     3    5
Processing:     1     3    7
Waiting:        0     3    6
Total:          4     6   10

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      6
  75%      6
  80%      6
  90%      7
  95%      7
  98%      8
  99%      8
 100%     10 (longest request)

```


## Part 29

Findings: 
### UUIDs Optimization:
1. First I used the uuid generator package through the `uuidgen` module. ```exec.Command("uuidgen").Output()```.
    Was not able to finish the generation of the 1 billion uuids as it took too long.
2. Replaced the ```uuidgen``` with ```github.com/google/uuid``` which increase the speed of the generation of 100000 uuids from 17 seconds to 4 seconds
3. Change the iteration over the WriteFile. Move the os.OpenFile and bufio.NewWriter functions outside the loop. It allows to reduce the time for 100000 uuids to less the 1 second.
    But 1 millions still take 57 seconds. But it takes more than 1 hour to generate 1 billion uuids.
4. After further optimization (create slice, append all uuids and then join them together and only than write to the file).
   Generating 1 million uuids takes less than 1 second. 10 millions = 9 seconds, 100 millions = 118 seconds. 
   While 1 billion uuids generating, I interrupt the script after 25 minutes. Still not acceptable. 
5. I was trying also use strings.Builder. Not much of improvements over previous method.
6. I tried another uuid generator ```github.com/pborman/uuid```. It gave few seconds improvements over previous package.

### UUID Search Optimization:
1. For search, I replace brute force for loop to generic binary search algorithm with ```sort.Search```. It did give only more time delay as the slice need to be sorted first.
    But still not acceptable performance. Request still was taken too long. 
2. For search the uuid in the file, I used the `grep` command. It did not increase the performance, the delay was too big, I interrupt the request.
    ```exec.Command("grep", "uuid", "uuid.txt").Output()```
3. Response time for the 100 millions uuids search is 10 seconds. 