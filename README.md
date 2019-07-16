# TCP Implementation in GO using TUN interface

Author: Eliezer Abate


In the following project I implement TCP layer on top of a TUN IP interface. The 
project fully implements all the various states of the TCP finite state machine.
The project implements TCP three-way handshake, communication, and termination.

## FEATURES IN DEVELOPMENT

- TCP Tahoe-like Sliding Window Protocol (https://pdfs.semanticscholar.org/9fce/b5b39a1c342a63a5c8f41c7dd5e10e93156d.pdf)[see here]

## UNSUPORTED PLATFORMS

- gonet has not been tested on windows

## BUGS 

- TODO 3/14
