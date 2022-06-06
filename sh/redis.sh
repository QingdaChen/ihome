docker run -p 6379:6379 --name redis -v /home/go/redis/data:/data \
-v /home/go/redis/conf/redis.conf:/etc/redis/redis.conf \
-d redis redis-server /etc/redis/redis.conf
