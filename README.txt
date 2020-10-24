Build with "go build ./main.go" (or skip this step by using "go run ./main.go")

Command line args:
-h
-profile [non-negative integer indicating number of times to repeat the profile, default 0 indicating no profiling]
-url [URL string in the form host:port/path (default "cloudflaregeneralswe2021.mcarth.workers.dev:80/links")]

Example usage:
./main.exe -profile=5 -url="www.google.com:80"



TEST RESULTS:


against /links API

PS C:\Users\mcart\Development\cloudflaresystemsswe2021> ./main.exe -profile=100       
Requests attempted: 100

Fastest request roundtrip: 38 ms
Slowest request roundtrip: 238 ms
Mean request roundtrip time: 68 ms
Median request roundtrip: 66.5 ms

Percentage successful requests:  100

No errors or non-2xx responses encountered

Smallest request response: 905 (bytes)
Largest request response: 925 (bytes)


against homepage

PS C:\Users\mcart\Development\cloudflaresystemsswe2021> ./main.exe -profile=50 -url="cloudflaregeneralswe2021.mcarth.workers.dev:80/" 
Requests attempted: 50

Fastest request roundtrip: 53 ms
Slowest request roundtrip: 159 ms
Mean request roundtrip time: 106 ms
Median request roundtrip: 106 ms

Percentage successful requests:  100

No errors or non-2xx responses encountered

Smallest request response: 2923 (bytes)
Largest request response: 3137 (bytes)


against google

PS C:\Users\mcart\Development\cloudflaresystemsswe2021> ./main.exe -profile=100 -url="www.google.com:80"
Requests attempted: 100

Fastest request roundtrip: 104 ms
Slowest request roundtrip: 181 ms
Mean request roundtrip time: 120 ms
Median request roundtrip: 118 ms

Percentage successful requests:  100

No errors or non-2xx responses encountered

Smallest request response: 48134 (bytes)
Largest request response: 48279 (bytes)


against google (small response size)

PS C:\Users\mcart\Development\cloudflaresystemsswe2021> ./main.exe -profile=50 -url="google.com:80"
Requests attempted: 50 

Fastest request roundtrip: 25 ms
Slowest request roundtrip: 70 ms
Mean request roundtrip time: 31 ms
Median request roundtrip: 30.5 ms

Percentage successful requests:  0

Error or non-2xx code responses encountered:
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently
Non 2xx http response code: 301 Moved Permanently

Smallest request response: 547 (bytes)
Largest request response: 547 (bytes)