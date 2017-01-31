# Question 5: Kuba's March


## (a)

The basics of Kuba's march is partitioning the space from the middle outwards.
By repeatedly finding the furthest point from a line between two points you can
find a point on the convex hull in $O(n)$. Since you only add more lines when
you find a convex point you end up doing $O(h)$, $O(n)$ searches. Thus giving a
$O(nh)$ runtime.

## (b)

Graham's Scan computes the convex hull in $O(n\log n)$ and Kuba's March in
$O(nh)$. We can use Chan's algorithm with Graham's Scan and Kuba's March in
place of Jarvis march to have a runtime of $O(n\log h)$.

The first step is to partition the points into $h$ sets and use Graham's Scan to
compute the convex hull of each set. Since there are
$\left\lceil{n/h}\right\rceil$ subsets and
each invocation of Graham's Scan is $O(h\log h)$ we get a runtime of $O(n\log
h)$.

The second step combines all of the sub convex hulls into the final using Kuba's
March. Since the points of the subset convex hulls are sorted we can do binary
search on them in order to find the furthest point away from a line in each
iteration of Kuba's March. Finding that point is $O(\log h)$ on each convex hull
and there are $\left\lceil{n/h}\right\rceil$ hulls thus giving each iteration a
total runtime of $O(n/h\log h)$. Since Kuba's March requires $O(h)$ iterations
we get $O(n\log h)$ for the second step.

Since both steps have the same runtime, the final runtime is $O(n\log h)$.
