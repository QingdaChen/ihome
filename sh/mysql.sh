sudo docker run -p 3306:3306 --name mysql \
-v /home/go/mysql/log:/var/log/mysql \
-v /home/go/mysql/data:/var/lib/mysql \
-v /home/go/mysql/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=220108 \
-d mysql
