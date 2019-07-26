# container-install

### Install
install container from url:
   location file:
  ```shell script
   container-install install 
          --docker true
          --host 172.16.213.131
          --host 172.16.213.132
          --user root
          --passwd admin
          --registry 127.0.0.1
          --registry 127.0.0.2
          --lib /var/lib/docker
          --pkg-url  /root/docker-19.0.3.tgz
```
   remote url:
  ```shell script
   container-install install 
          --docker true
          --host 172.16.213.131
          --host 172.16.213.132
          --user root
          --passwd admin
          --registry 127.0.0.1
          --registry 127.0.0.2
          --lib /var/lib/docker
          --pkg-url  https://download.docker.com/linux/static/stable/x86_64/docker-18.09.4.tgz
```
### UnInstall
uninstall container:
  ```shell script
   container-install uninstall 
          --docker true
          --host 172.16.213.131
          --host 172.16.213.132
          --user root
          --passwd admin
          --docker-lib /var/lib/docker
```

### Print Download Url
print download url for docker:
 ```shell script
  container-install print --docker true
```

the Newest version is v19.03.0.
ex:

```
cuisongliu@cuisongliu-PC:~$ container-install print --docker true
https://download.docker.com/linux/static/stable/x86_64/docker-17.03.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.03.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.03.2-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.06.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.06.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.06.2-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.09.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.09.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.12.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-17.12.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.03.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.03.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.0-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.1-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.2-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.06.3-ce.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.0.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.1.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.2.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.3.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.4.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.5.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.6.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.7.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-18.09.8.tgz
https://download.docker.com/linux/static/stable/x86_64/docker-19.03.0.tgz

```
