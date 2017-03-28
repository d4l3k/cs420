# Question 3

Assume we have a black box algorithm that computes whether or not a undirected
graph G and an integer k, contains a clique of size k as well as an independent
set of size k.

We know that computing the size of the maximal independent set in a graph is
NP-complete. We can use this algorithm to compute the size, by running it on the
graph n times where $k=\{1,\ldots,n\}$. If the algorithm is in polynomial time,
running it $n$ times is still in polynomial time. Likewise, if our algorithm is
not in polynomial time, running it $n$ times is still not in polynomial time.
Since finding the maximal independent set is NP-complete, our black box
algorithm must also be NP-complete.

Computing whether there is a clique of size k is also NP-complete. Thus, since
both parts are NP-complete the overall algorithm is NP-complete.
