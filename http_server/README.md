# 検証で利用するHTTP Serverのセットアップ
EC2のインスタンス上で実施します。
#### 1. 検証用のプログラム準備
##### 1.1. ダウンロードします。
ログ出力とprometheus用のメトリクスを生成する検証用プログラムをダウンロードします。
```
$ cd
$ wget https://raw.githubusercontent.com/kichiram/golang/main/testgo/test_httpserver.go
$ ls test_httpserver.go 
test_httpserver.go
```
##### 1.2. ビルドに必要なprometheusのライブラリをダウンロードします。
```
$ go get github.com/prometheus/client_golang/prometheus
```
##### 1.3. 検証用のプログラムをビルドします。
```
$ go build test_httpserver.go
$ ls test_httpserver
test_httpserver
```
#### 2. 常駐プロセス化
daemon（常駐プロセス）にして管理しやすいようにします。
##### 2.1. ファイル整理
ファイルを/usr/local/bin配下に移動します。
```
$ sudo mv ~/test_httpserver /usr/local/bin/
$ ls /usr/local/bin/test_httpserver
/usr/local/bin/test_httpserver
```
##### 3.2. daemonの設定ファイル作成
```
$ sudo vi /usr/lib/systemd/system/test_httpserver.service
[Unit]
Description=test HTTP Srver.
Documentation=https://github.com/kichiram/golang/http_server
After=network-online.target

[Service]
Type=simple
ExecStart=/bin/sh -c "/usr/local/bin/test_httpserver >> /var/log/test_httpserver.log 2>&1"

[Install]
WantedBy=multi-user.target
```
##### 3.3. daemonの自動起動設定と起動
```
$ sudo systemctl enable test_httpserver.service
Created symlink from /etc/systemd/system/multi-user.target.wants/test_httpserver.service to /usr/lib/systemd/syste
m/test_httpserver.service.
$ sudo systemctl start test_httpserver.service
```
##### 3.4. 起動確認
```
$ sudo systemctl status test_httpserver.service
● test_httpserver.service - test HTTP Srver.
   Loaded: loaded (/usr/lib/systemd/system/test_httpserver.service; enabled; vendor preset: disabled)
   Active: active (running) since Tue 2021-10-26 09:27:25 UTC; 6s ago
     Docs: https://github.com/kichiram/golang/http_server
 Main PID: 10324 (sh)
   CGroup: /system.slice/test_httpserver.service
           ├─10324 /bin/sh -c /usr/local/bin/test_httpserver >> /var/log/test_httpserver.log 2>&1
           └─10325 /usr/local/bin/test_httpserver

Oct 26 09:27:25 ip-172-31-39-113.ap-northeast-1.compute.internal systemd[1]: Started test HTTP Srver..
--- 省略 ---
```
active (running)と表示されていれば成功です。
#### 4. 動作確認
##### 4.1. hello出力
```
http://<ホスト名>:8080/hello
```
「hello」と出力されれば成功です。
##### 4.2. world出力
```
http://<ホスト名>:8080/world
```
「world」と出力されれば成功です。
##### 4.3. prometheusメトリクス出力
```
http://<ホスト名>:8081/metrics
```
下記のように今回検証で利用するメトリクスの「a_http_request_count_total」が出力されていれば成功です。
```
# HELP a_http_request_count_total Test Counter
# TYPE a_http_request_count_total counter
a_http_request_count_total{testlabel="Hello"} 1
a_http_request_count_total{testlabel="world"} 1
--- 省略 ---
```
なお、このメトリクスはhello、worldにリクエストされたカウントを計測します。
##### 4.4. ログ出力
```
$ tail -5 /var/log/test_httpserver.log 
2021/10/26 09:27:57 Hello!
2021/10/26 09:27:57 Hello!
2021/10/26 09:28:26 World
2021/10/26 09:28:26 World
2021/10/26 09:28:26 World
```
hello、worldにリクエストされた際のログが出力されていれば成功です。
