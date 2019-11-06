# GoCV motion detection

## installation 
The procedure is really straight forward:
```bash
echo "START" \
&& go get -u -d gocv.io/x/gocv \
&& cd $GOPATH/src/gocv.io/x/gocv \
&& make install \
&& echo "STOP"
```
Notice two switches that we use with `get` command:
* **-u** update named packages and their dependencies
* **-d** just download package, but do not install it