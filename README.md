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
