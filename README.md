# Workaround for issue with system call from multiple threads in parallel

Demo of issue.

> go build test_concurrency.go

> ./test_concurrency.go


Running the functions sequentially:

Current directory from change_into_directory_1():
*** This is file 1 ***

Current directory from change_into_directory_2():
*** This is file 2 ***


--------------- Running the functions in parallel ---------------

Current directory from change_into_directory_2(): 
Current directory from change_into_directory_1():
*** This is file 2 ***

<<< ERROR getting file 1 from directory 1>>>


--------------------------------------------------------------------------------------------------------
O/S system calls to change the directory in thread_1 & thread_2 [functions change_into_directory_1() & change_into_directory_2()]
are not remaining indenpendent of one another. They seem to be sharing system state resources underneath, implying the code is 
not re-entrant.

EXPECTED BEHAVIOUR: That thread_2 does not reflect the current O/S directory state of thread_1 and operates independently of thread_1
including at the O/S level, so that the functions which are run in parallel remain fully re-entrant from the point-of-view of
the caller.

JAVA CASE
=========
Java copes with a similar  use case. To run the java:

> cd JAVA_TEST

> javac *.java

> java Concurrency_Test 

It works perfectly every time with a very similar multi-threaded use case with Java (OpenJDK 1.8). So my question is why doesn't GoLang behave similary (i.e.*reasonably*) like Java does...???

GoLang Case with working directory (Dir) member of cmd struct modified to exactly mirror correctly working Java equivalent
===========================================================================================================================

Running the functions sequentially:

Current directory from change_into_directory_1(): /home/blissnd/GoLang_Multi_Thread_System_Call
*** This is file 1 ***

Current directory from change_into_directory_2(): /home/blissnd
<<< ERROR getting file 2 from directory 2 >>>


--------------- Running the functions in parallel ---------------

Current directory from change_into_directory_2(): /home
<<< ERROR getting file 2 from directory 2 >>>

Current directory from change_into_directory_1(): /
<<< ERROR getting file 1 from directory 1 >>>


I.E. Does not behave according to [https://golang.org/pkg/os/exec/#Cmd]

