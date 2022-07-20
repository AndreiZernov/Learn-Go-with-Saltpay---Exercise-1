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

