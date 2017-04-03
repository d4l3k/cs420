# Question 4

I looked up a definition of k-competitive algorithms online.

For a page replacement algorithm to be k-competitive, given $k$ pages and cost
function $c$ such that for all sequences of $p=\{p_1,p_2,\ldots, p_n\}$,
$c_{ALG}(p_i) = k*c_{OPT}(p_i)$ must hold.

We assume that both the FIFO and OPT algorithms start with the same pages in
memory.

We can partition $p$ into a number of phases such that FIFO has at most k fault
on $P(0)$ and exactly $k$ faults on $P(i), i\geq 1$. To show that FIFO is
k-competitive, we must show that for each phase, OPT has at least one page
fault.

For the case of $P(0)$, since both FIFO and OPT both start with the same pages
in memory, the first fault for FIFO occurs at the same time as OPT.

For $k$ page faults to occur in a phase with FIFO, there must be at least $k$
items that are different from the request immediately before the phase starts.
Assuming this is true, since there are $k$ new items in this phase, OPT must
have at least one page fault.

Assume that FIFO faults on a page twice in a phase. This implies that has been
evicted once during the phase. Since the algorithm must see $k$ new items to
remove one that's been added, this cannot happen with only $k$ page faults.
Thus, during this phase, all $k$ page faults must be distinct, and OPT must
fault at least once.

Therefore, FIFO is k-competitive since for every $k$ faults of FIFO, there must
be at least one by OPT.
