# wc 


A ~slightly~ significantly worse implementation of GNU [wc](https://www.gnu.org/software/coreutils/manual/html_node/wc-invocation.html#wc-invocation) in go. 

# Performance comparison
All the files mentioned below are present in the [`tests/`](https://github.com/theredditbandit/wc/tree/master/tests) directory

## GNU wc 
| file               	| size 	| time to execute 	| CPU usage 	| user CPU time 	| system CPU time 	|
|--------------------	|------	|-----------------	|-----------	|---------------	|-----------------	|
| bee-movie.txt      	| 88k  	| 0.002s          	| 76%       	| 0.00s         	| 0.00s           	|
| test.txt           	| 336k 	| 0.003s          	| 90%       	| 0.00s         	| 0.00s           	|
| 1mboftxt           	| 1.0M 	| 0.005s          	| 95%       	| 0.01s         	| 0.00s           	|
| les-miserables.txt 	| 3.3M 	| 0.020s          	| 97%       	| 0.02s         	| 0.00s           	|

## gwc (mine)
| file               	| size 	| time to execute 	| CPU usage 	| user CPU time 	| system CPU time 	|
|--------------------	|------	|-----------------	|-----------	|---------------	|-----------------	|
| bee-movie.txt      	| 88k  	| 0.004s          	| 120%      	| 0.01s         	| 0.00s           	|
| test.txt           	| 336k 	| 0.009s          	| 111%      	| 0.01s         	| 0.00s           	|
| 1mboftxt           	| 1.0M 	| 0.015s          	| 108%      	| 0.02s         	| 0.00s           	|
| les-miserables.txt 	| 3.3M 	| 0.051s          	| 99%       	| 0.04s         	| 0.01s           	|

## Conclusion 
On average , my program is **163 times slower** and consumes **26% more CPU** to provide performance that is worse than the `wc` tool that ships with most GNU/linux.

# Strengths
 - identical results when compared with GNU `wc` over the test data.
 - `-v` flag to understand what the hell is actually being printed

# Limitations
 - cannot read from stdin if no file is passed so `cat somefile.txt | gwc -l` won't work.
 - messy codebase
 - no automated testing
 - slower


# Images 

![comparison](https://github.com/theredditbandit/wc/assets/85390033/fc7742c0-acfa-4896-b1b6-2b5945536e0a)

![verbose](https://github.com/theredditbandit/wc/assets/85390033/272f55ec-5587-4684-a30f-d2ed47bb74d1)



# TODO (maybe/eventually/will never get to it)
 - [ ] rewrite using simpler techniques.
 - [ ] use charmcli tools and have nice coloured output
 - [ ] read from stdin if no file is provided
