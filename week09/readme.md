
总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

fix length，定长编码
定长编码粘包：因为编码后报文长度是固定的，每次从缓冲区，按定长取即可
定长编码应用：tcp报文力度 以太、ip、tcp头都是固定长度的

delimiter based，基于定界符
基于定界符粘包：在缓冲区找到界定符，就可以分隔两个包了。
基于定界符应用：http 报文里 \r\n 代表一个头部的结束, : 符号作为name和value的区分。

length field based frame decoder，基于长度解码
基于长度解码粘包：主要是需要根据特征位先判断用几个字节表示长度，然后就可以计算出报文长度。
基于长度解码应用：TLV 编码，TLV：TLV是指由数据的类型Tag，数据的长度Length，数据的值Value组成的结构体，几乎可以描任意数据类型，TLV的Value也可以是一个TLV结构，正因为这种嵌套的特性，可以让我们用来包装协议的实现。

