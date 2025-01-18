# isucon_ansible

ISUCONをプレイするにあたって競技用インスタンスに対して諸々のデプロイを自動化するためのAnsibleファイルを集めたものです。

## 事前条件

- ローカルにansibleが使える環境が整っていること
- ローカルから競技用インスタンスにSSHができること
- 競技用サーバからGithubへアクセスするための公開鍵・秘密鍵が作成済みであること
    - `~/.ssh/isucon_id_rsa`, `~/.ssh/isucon_id_rsa.pub`として保存されていることを期待している
- 資材を管理するようのGithubレポジトリが作成済みであること

## Quick Usage

```
# レポジトリのダウンロード
$ git clone https://github.com/itohdak/isucon_ansible.git
$ cd isucon_ansible

# /etc/hostsファイル修正
$ sudo vim /etc/hosts
```
<details>
<summary>/etc/hostsの設定内容例</summary>

```
$ cat vim /etc/hosts
<競技用サーバ1のIP> s1
<競技用サーバ2のIP> s2
<競技用サーバ3のIP> s3
<監視用サーバのIP>  pprotein
```
</details>

```
# ~/.ssh/configの設定
$ vim ~/.ssh/config
```
<details>
<summary>~/.ssh/configの設定内容例</summary>

```
$ cat ~/.ssh/config
Host s1
  HostName s1
  User isucon
  Port 22
  IdentityFile /home/itohdak/.ssh/isucon_id_rsa
  LocalForward 8443 localhost:443

Host s2
  HostName s2
  User isucon
  Port 22
  IdentityFile /home/itohdak/.ssh/isucon_id_rsa

Host s3
  HostName s3
  User isucon
  Port 22
  IdentityFile /home/itohdak/.ssh/isucon_id_rsa

Host pprotein
  HostName pprotein
  User isucon
  Port 22
  IdentityFile /home/itohdak/.ssh/isucon_id_rsa
```
</details>

```
# 実行
$ ansible-playbook playbooks/deploy_all.yaml
```

## 実行内容

`deploy_all.yaml`では、以下の内容を実行します。
- 汎用パッケージのインストール
    - 以下のパッケージをインストールします
        - graphvix (pprotein用)
        - gv (pprotein用)
        - netdata
- /etc/hostsの設定
    - 競技用インスタンスや監視用サーバ間でホスト名で通信できるようにします
    - (s1|s2|s3|pprotein).maca.jpというホスト名がそれぞれのサーバに振られます
- 競技用インスタンスの資材をGitレポジトリで管理するための設定
    - 公開鍵・秘密鍵の配置
    - `.gitconfig`の配置
    - レポジトリのクローン
- 監視用スタックの構築
    - slp, alp, pprotein, pprotein-agentのダウンロードとインストール
    - ansibleの設定
        - ブラウザからansibleが実行できるようにansibleの設定を監視用インスタンス上にも展開する


## その他

pproteinについては、以下を参照してください。
- https://github.com/kaz/pprotein