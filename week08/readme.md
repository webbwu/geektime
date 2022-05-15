1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 10 -t get,set -q
SET: 78186.08 requests per second
GET: 77279.75 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 20 -t get,set -q
SET: 78802.20 requests per second
GET: 77160.49 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 50 -t get,set -q
SET: 77639.75 requests per second
GET: 78431.38 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 100 -t get,set -q
SET: 78308.54 requests per second
GET: 78247.26 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 200 -t get,set -q
SET: 78740.16 requests per second
GET: 78003.12 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 1000 -t get,set -q
SET: 76745.97 requests per second
GET: 76277.65 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 5000 -t get,set -q
SET: 73099.41 requests per second
GET: 72046.11 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 10000 -t get,set -q
SET: 70821.53 requests per second
GET: 66666.66 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 20000 -t get,set -q
SET: 68917.99 requests per second
GET: 48192.77 requests per second

[root@iZ2zeidbsi5o2h57fw1zujZ bin]#  redis-benchmark -d 50000 -t get,set -q
SET: 45620.44 requests per second
GET: 32425.42 requests per second



由数据可以看出这台机器处理性能7.8W左右，从10-1K, 读写性能差不多，从5K开始性能急剧下降


2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。


| value 大小（byte） | key 序列化长度（byte） |
|----------------|-----------------|
| 10000          | 125             |
| 20000          | 242             |
| 50000          | 584             |
| 100000         | 1151            |
| 200000         | 2288            |
| 500000         | 5696            |

1. redis 的 key 随着 value 的增大会逐渐增大，在 value 大小不变的情况下，key 序列化后的长度也不会发生变化，不论同样长度的 value 的 key 有多少个。
2. 每当 value 大小翻倍增长时，key 也近乎成倍增长。