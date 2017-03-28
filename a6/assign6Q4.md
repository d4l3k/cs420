# Question 4

We can reduce the PARTITION problem into a LIST SCHEDULING problem in polynomial
time.

In PARTITION we want to find two subsets of equal summed value. To transform
PARTITION into LIST SCHEDULING, take $m=2$ and $d=\frac{\sum_{i=1}^n x_i}{2}$,
and then assign $p=x$. This transformation runs in $O(n)$. If there is a subset
that is equal, LIST SCHEDULING will return that all jobs can be completed before
the deadline.

Since PARTITION is NP-hard, and can be transformed into LIST SCHEDULING in
polynomial time, LIST SCHEDULING must be NP-hard as well.
