### 单飞模式

data, err := getData("key")
执行结果如下：
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
2021/07/14 16:01:15 get key from db
2021/07/14 16:01:15 receive data:  hello
可以看到10个请求都是走的db，因为cache中不存在该值		
                    
            
data, err := getDataInSingleflight("key")
执行结果如下：
2021/07/14 15:59:45 get key from db
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello
2021/07/14 15:59:45 receive data:  hello


可以看到只有一个请求进入的db，其他的请求也正常返回了值，从而保护了后端DB。