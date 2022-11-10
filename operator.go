package main

import "fmt"

func main() {
	fmt.Println("运算符")
	/**
	  Go 语言内置的运算符有：
	  算术运算符
	  关系运算符
	  逻辑运算符
	  位运算符
	  赋值运算符
	  其他运算符
	*/
	//算数运算符+ - * / % ++ --
	var d int = 1
	var b int = 2
	fmt.Println("相加+", d+b)
	fmt.Println("相减-", d-b)
	fmt.Println("相乘+", d*b)
	fmt.Println("相除/", d/b)
	fmt.Println("取余%", d%b)
	d++
	fmt.Println("自增+", d)
	b--
	fmt.Println("自减-", b)

	//关系运算符 == != > < >= <=

	//逻辑运算符 &&并 ||或 ！非

	//位运算符 位运算符可以对整数在内存中的二进制位进行操作
	// &：与运算，全真为真；  |：或运算，全假才假；  ^:异或运算，相同为假，不同为真

	//赋值运算符
	/**
	运算符	描述	实例
	=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
	+=	相加后再赋值	C += A 等于 C = C + A
	-=	相减后再赋值	C -= A 等于 C = C - A
	*=	相乘后再赋值	C *= A 等于 C = C * A
	/=	相除后再赋值	C /= A 等于 C = C / A
	%=	求余后再赋值	C %= A 等于 C = C % A
	<<=	左移后赋值	C <<= 2 等于 C = C << 2
	>>=	右移后赋值	C >>= 2 等于 C = C >> 2
	&=	按位与后赋值	C &= 2 等于 C = C & 2
	^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
	|=	按位或后赋值	C |= 2 等于 C = C | 2
	*/

	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200
	c <<= 2
	fmt.Printf("第 6 行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)

	//其他运算符
	/**

	运算符	描述	实例
	&	返回变量存储地址	&a; 将给出变量的实际地址。
	*	指针变量。	*a; 是一个指针变量
	*/
}
