nc 144.21.37.133 8080 | { while read fm; do if [[ "$fm" == "<password>" ]]; then etherwake -i br-lan 00:d0:7f:ff:25:56; fi done }
