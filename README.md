# Learn Go with Saltpay - Exercise 1

[Github Link](https://github.com/saltpay/learn-go-with-salt/blob/master/book/exercise1.md) 

# Exercise 1

Each part builds on top of the previous one. After completing each part, tag the commit as `track-1-part-1` etc.

If you make any assumptions about the feature (e.g. negative numbers are not supported), mention it in the readme.


## Part 1

- [x] Write a function to add two integers.

  <em>e.g. `add(5, 4)` should return `9`</em>

## Part 2

- [x] Modify the function to add a variable number of integers. 

  <em>e.g. `add(5, 4, 3, 2, -10)` should return `4`</em>

## Part 3

- [x] Create a program `add` that will take any number of integer arguments and print out the sum. Ignore any non-integers in the list.

  <em>Example usage of the program: `add 4 5 32 100 867543` should output `867684`</em>

## Part 4

- [x] Modify the output of the program so numbers larger than 9999 are formatted in groups of thousands. 

  <em>e.g. 9999 will be shown as `9999`, 10000 will be shown as `10,000`, and 1234567890 will be shown as `1,234,567,890`</em>

## Part 5

- [x] Change the program so that when it is called with no arguments (`add`), it read the numbers from a file called `input.txt` in the same directory. The file will have one number per line. Ignore any non-integers in the file.

  <em>Note that the old behaviour of the program (taking a list of integers as arguments) still works.</em>

  Example of `input.txt`:

  ```
  4
  5
  32
  100
  867543
  ```

## Part 6

- [x] Change the program to specify the input file name as an argument. 
  - [ ] Part 6.1 Find the way to get the file from anywhere in the file system

  <em>e.g. `add --input-file data/input.txt`. The file can be anywhere on the file system.</em>

## Part 7

- [x] Change the program so it can also read the numbers from a CSV file. 

  <em>e.g. `add --input-file data/input2.csv` where `input2.csv` looks like this:</em>
  ```
  4,5,32,100,867543
  ```

  <em>You can assume the file is a CSV if the extension is .csv</em>

## Part 8

- [x] Change the program so it does not rely on the .csv extension. 

  <em>It should auto-detect if the file is a CSV or newline-separated file. Therefore `add --input-file input.txt` should work regardless of whether the numbers in `input.txt` are separated by commas or newlines.</em>

## Part 9

- [x] Change the program so it accepts multiple files as input, each of which can be any supported format.

  <em>e.g. `add --input-file one.txt --input-file data/two.txt`. The output is the sum of all the numbers from all the specified files.</em>

## Part 10

- [x] Write a CI pipeline in Github Actions to run all your tests when you push code to the trunk.

## Part 11

- [x] Change the program so it discards any duplicate numbers. e.g. `add 2 2 2 3 4 4 4` will output `9` (2+3+4).

## Part 12

- [x] Create a new program `math` so when it is invoked as `math --web-server`, it starts a web server with an endpoint `POST /add?num=4&num=5&num=32` that returns the response `41` as text.

  <em>(Ignore why it is a POST method, for now)</em>

## Part 13

- [x] Change the math web server, so it can also read a form-urlencoded body from the request.

  <em>Example request will still return `41` as text:</em>
  
  ```
  POST /add
  
  num=4&num=5&num=32
  ```


## Part 14

- [x] Change the math web server, so it can also accept numbers as a JSON array.

  Example request:
  
  ```
  POST /add
  
  {
      "nums": [4, 5, 32]
  }
  ```

  <em>Research the request headers that a client will send that will allow different forms of the request body to be accepted by the server.</em>


## Part 15

- [x] Add a naive authentication to the endpoint. Only requests containing a header `Authorization: Bearer SUPER_SECRET_API_KEY` (that exact API key) should be allowed in.

  <em>Research and implement the behaviour for when:</em>

  1. The header is not present
  2. The API key is incorrect

## Part 16

- [x] Instead of a hard-coded API key (should have never committed the API key to source control anyway), read in a list of API keys from an environment variable. Any request with one of those API keys should be allowed; all other requests should be rejected.

## Part 17

Add a new endpoint `GET /fibonacci/:n` which will return the nth number in the Fibonacci sequence (0,1,1,2,3,5,8,13 etc.) where 0 is the 1st number, and 13 is the 8th number.

## Part 18

Extend the API authentication done earlier to also protect the fibonacci endpoint.

## Part 19

For each request that comes in, log out the following info to STDOUT.

- timestamp
- HTTP method
- path and query params (`GET /fibonacci/8`, `POST /add?num=9&num=12`)
- first 8 chars of the API key
- request body size
- response code
- time taken to send the response

e.g. `2022-04-12T13:43:33.000Z GET /fibonacci/8 GYHR65NH 0 200 132ms`

## Part 20

For all endpoints, support a query param `?flakiness=:p`, where 0 <= p <= 1

Based on that probability, return a response code of 500 for that request.

e.g. if a request comes in as `?flakiness=1`, it's a 100% probability, so send a 500.
For `?flakiness=0.2`, the probability should be 1 in 5 of returning a 500.

## Part 21

Extend the flakiness param to specify the flaky response code.

e.g. `?flakiness=0.2,404` should return a 404 with a probability of 20%

## Part 22

Extend the flakiness param to simulate a slow server by specifying a delay.

e.g. `?flakiness=0.33,500,3s` will introduce a delay of 3 seconds before sending a 500, all with a probability of 33%

Support `s` and `ms` as units for the delay

## Part 23

Build a program `fiboclient` that calls the fibonacci endpoint (running in a separate process) for the value `10`.

## Part 24

Extend the `fiboclient` program to take in a list of integers as input, call the `/fibonacci` endpoint with that number and print the results as `fibo {input}: {result}` pairs.

e.g. `./fiboclient 1 2 4 6` will print

```
fibo 1: 0
fibo 2: 1
fibo 4: 2
fibo 6: 5
```

## Part 25

Change all endpoints so if a query parameter `?format=thousands` is passed, it formats the numbers by thousands (like Part 4), otherwise the numbers are not formatted.


## Part 26

Run the web server and simulate load to the `/add endpoint` from multiple concurrent users.

Try running 1000 requests with a concurrency of 10, then 10000 requests with a concurrency of 100.

Hint: use a tool such as [ab](https://httpd.apache.org/docs/2.4/programs/ab.html)


## Part 27

Extend the request logging from Part 19 so it writes to a file `access_log` as well as to STDOUT. The log file should be appended to and data should not be lost between application restarts.

Use a logging library if you want.

## Part 28

Write a script to generate a file `authorised_api_access_keys` filled with a 1000 UUIDs, one on each line.

Your web server should now use this file as the source for authorised API keys instead of reading them from the environment (from Part 16)

## Part 29

Use the UUID generating script to create a file with 1 _billion_ entries. The file would be approx. 5GB in size, so ensure you have enough disk space available.

If you change the web server to use this file as the source for authorised API keys, and test it with an API key near the end of the file, does the request still succeed? How long does it take? How does your application's memory usage look?
