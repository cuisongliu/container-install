# docker-install

### Install
install docker from url:
  ```shell script
   docker-install install 
          --host 172.16.213.131
          --user root
          --passwd admin
          --registry 127.0.0.1
          --docker-lib /var/lib/docker
          --pkg-url  /root/docker-19.0.3.tgz
```

### UnInstall
uninstall docker:
  ```shell script
   docker-install uninstall 
          --host 172.16.213.131
          --user root
          --passwd admin
          --docker-lib /var/lib/docker
```

### Print Download Url
print download url for docker:
 ```shell script
  docker-install print
```

the Newest version is v19.03.0.
