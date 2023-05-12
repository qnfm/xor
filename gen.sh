for i in `seq 0 16 1024`
do
	tmp=$(($i+16))
	echo "fastxor.Block(buf3["$i":"$tmp"], buf1["$i":"$tmp"], buf2["$i":"$tmp"])">> url.txt
done