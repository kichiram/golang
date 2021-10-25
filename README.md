# golangのインストールと使い方について
EC2のインスタンス上で実施します。
#### 1. インストール
##### 1.1. yumコマンドでインストールします。
```
$ $ sudo yum install golang -y
--- 省略 ---
Installed:
  golang.x86_64 0:1.15.14-1.amzn2.0.1 
--- 省略 ---
Dependency Updated:
  glibc.x86_64 0:2.26-55.amzn2                              glibc-all-langpacks.x86_64 0:2.26-55.amzn2            
  glibc-common.x86_64 0:2.26-55.amzn2                       glibc-locale-source.x86_64 0:2.26-55.amzn2            
  glibc-minimal-langpack.x86_64 0:2.26-55.amzn2             libcrypt.x86_64 0:2.26-55.amzn2                       

Complete!
```
##### 1.3. インストールされたこと確認します。
```
$ yum list installed golang
Loaded plugins: extras_suggestions, langpacks, priorities, update-motd
Installed Packages
golang.x86_64                                   1.15.14-1.amzn2.0.1                                    @amzn2-core  
```
#### 2. ~/.bash_profileに環境変数GOPATHを追加
```
$ sed -i -e '$ a export GOPATH=$HOME/go' ~/.bash_profile
$ source ~/.bash_profile
$ echo $GOPATH
/home/ec2-user/go
```
#### 3. 動作確認用のプログラム作成
```
$ mkdir -p $HOME/testgo;cd $HOME/testgo
$ vi hello.go
package main
  
import "fmt"

func main() {
    fmt.Printf("Hello World\n")
}
```
#### 4. 動作確認
##### 4.1. go runで実行してみます。
```
$ go run hello.go 
Hello World
```
Hello Worldと表示されていれば成功です。
##### 4.2. ビルドして実行ファイル（バイナリ）を作成します。
```
$ go build hello.go 
$ ls hello
hello
```
##### 4.3. 実行ファイル（バイナリ）を実行してみます。
```
$ ./hello
Hello World
```
先ほどと同様にHello Worldと表示されていれば成功です。
