# Go Vanity URLs for normal vps

Go Vanity URLs is a simple app that allows you to set custom import paths for your Go packages. 

Note: **this go vanity urls is forked from GoogleCloudPlatform/govanityurls. GoogleCloudPlatform/govanityurls only works on google app engine. go vanity urls here could work on normal vps**

## Quickstart

Install go vanity urls:

```
$ go get github.com/bigwhite/govanityurls
```

Edit `vanity.yaml` to add any number of git repos. E.g., `tonybai.com/gowechat` will
serve the [https://github.com/bigwhite/gowechat](https://github.com/bigwhite/gowechat) repo.

```
/gowechat:
  repo: https://github.com/bigwhite/gowechat

/experiments:
  repo: https://github.com/bigwhite/experiments
```

You can add as many rules as you wish.

Before run the app, point your custom domain to the vps ip where govanityurl deployed.

Run the app:

```
$ cd $(go env GOPATH)/src/github.com/bigwhite/govanityurls
$ govanityurls -host tonybai.com
```

That's it! You can use `go get` to get the package from your custom domain.

```
$ go get tonybai.com/gowechat
```
