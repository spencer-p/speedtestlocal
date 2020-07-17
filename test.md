In order, we have
 1. Over the 5Ghz netgear router
 2. Over the AT&T router
 3. Over loopback (notice it's ms)
 4. Over AT&T but in the room next to it

```
$ ./speedtestlocal --step $((1*1024*1024)) --wb $((1*1024*1024)) --max $((50*1024*1024))
2020/07/16 17:49:31 Booting up
2020/07/16 17:49:34 GET /ws
2020/07/16 17:49:39 Delivered 52428800 bytes in 5.322686106s
2020/07/16 17:50:16 GET /
2020/07/16 17:50:18 GET /ws
2020/07/16 17:51:12 Delivered 52428800 bytes in 54.255956705s
2020/07/16 17:53:09 GET /
2020/07/16 17:53:09 GET /index.js
2020/07/16 17:53:12 GET /ws
2020/07/16 17:53:13 Delivered 52428800 bytes in 568.454714ms
2020/07/16 17:54:18 GET /
2020/07/16 17:54:18 GET /index.js
2020/07/16 17:54:22 GET /ws
2020/07/16 17:54:51 Delivered 52428800 bytes in 28.977690393s
```
