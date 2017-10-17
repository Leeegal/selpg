# Selpg

Use golang to develop the  [selpg](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)（Command Line Interface）

# How to use

  - you may at least input three args there are:
    
    -sNumber , is the num of start page 
    
    -eNumber , is the num of end page 
    
    -inputFile , is the file you want to print
  - otehr args may be:
    
    -l ,the num of the length of every page (default: 72)
    
    -f ,means every page is devided by the '\f'
    
    -dlpx ,means printing the output by the printer lpx
    
# Design And Test
Because I didn't know how to write this code at the begginning, so I just imitated the source code [selpg.c](https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c), So I didn't use so many functions, which may let you see some awful. So some details of the code will be not described here.

- We can see the files in the project: 

> -the selpg.go is the source code.

> -the outputx.txt is the file to output the pages

> -the testx.txt is the "inputFile" I use

> -the errorx.txt is the stderr that may be printed

#Here are some command(I forgot to use the $GOPATH/bin)

![1.png](http://upload-images.jianshu.io/upload_images/1039936-1e3a5d09eb117810.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![2.png](http://upload-images.jianshu.io/upload_images/1039936-d0a00844817faba6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![3.png](http://upload-images.jianshu.io/upload_images/1039936-a8719c1ec79a8fd2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![4.png](http://upload-images.jianshu.io/upload_images/1039936-ed08c38bdf24b932.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![5.png](http://upload-images.jianshu.io/upload_images/1039936-7f577a57b763c225.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![6.png](http://upload-images.jianshu.io/upload_images/1039936-2579164b0f8b8df4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![7.png](http://upload-images.jianshu.io/upload_images/1039936-a41c6af2769ee05d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![8.png](http://upload-images.jianshu.io/upload_images/1039936-e8ef3fa6f9eaece3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![9.png](http://upload-images.jianshu.io/upload_images/1039936-82cbd00dc7117a48.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![9_2.png](http://upload-images.jianshu.io/upload_images/1039936-d5e9859081cffb23.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
