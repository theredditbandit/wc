# wc 


A ~slightly~ significantly worse implementation of GNU [wc](https://www.gnu.org/software/coreutils/manual/html_node/wc-invocation.html#wc-invocation) in go. 

# Images 

![comparison](https://github.com/theredditbandit/wc/assets/85390033/fc7742c0-acfa-4896-b1b6-2b5945536e0a)

![verbose](https://github.com/theredditbandit/wc/assets/85390033/272f55ec-5587-4684-a30f-d2ed47bb74d1)

# Strengths
 - identical results when compared with GNU `wc` over the test data.
 - `-v` flag to understand what the hell is actually being printed

# Limitations
 - cannot read from stdin if no file is passed so `cat somefile.txt | gwc -l` won't work.
 - messy codebase
 - no automated testing

# TODO (maybe/eventually/will never get to it)
 - [ ] rewrite using simpler techniques.
 - [ ] use charmcli tools and have nice coloured output
 - [ ] read from stdin if no file is provided
