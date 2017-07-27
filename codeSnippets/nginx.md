#Modulized Nginx Config

## nginx.conf

```
user  nginx;
worker_processes  1;

error_log  /dev/stderr warn;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /dev/stdout  main;
    sendfile        on;
    keepalive_timeout  65;

    include /etc/nginx/upstreams/*.conf;

    server {
        listen 80;
        include /etc/nginx/conf.d/*.conf;
    }

}
```

## foo.http.conf

```
location /socket {
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header Host $http_host;
      proxy_set_header X-NginX-Proxy true;

      proxy_pass http://socketio;
      proxy_redirect off;

      proxy_http_version 1.1;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
}
```

## foo.upstream.conf
```
upstream socketio {
    server 127.0.0.1:3000;
}
```
