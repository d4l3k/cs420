# Question 2

Let the candidate set be all vertices in G. Pick a vertex $v$ in the candidate set,
add it to the independent set, remove all neighbors of $v$ from the candidate
set. Repeat until the candidate set is empty. This can be done in $O(n)$ which
makes it very efficient.

The tricky part is showing that this has the correct size.

This algorithm will pick at least one vertex from the set of a vertex and
it's neighbors of degree $d$. If it picks the central vertex with degree $d$,
that means of that local set $\frac{1}{d+1} * (d+1)$ is included. In the case of
a star, the maximum independent set will be all of the edge vertices with size
$d$. The worst case as above is picking the middle, which is $\frac{1}{d}$ the
optimal size. Thus in the limit.


