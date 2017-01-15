# 3.

We can take any two adjacent points on the edge of a convex hull and use them to
define a line that has all of the points on one side of the line since by
definition there are no points beyond the edge of a convex hull. To compute all
candidate lines we can take all of the corners of the convex hulls and use them
to define lines in $O(n)$. This satisfies the constraint that all points of P
are on one side.

We can compute the centeroid of P in $O(n)$ by taking the average of all the X
and Y coordinates. Since the centeroid is the average of all the points, the
distance from the centeroid is a good estimate of the average distance from all
points. We can then iterate through all of the candidate lines and compute the
distance from the centeroid returning the one with the smallest distance. Since
all steps are done in $O(n)$ the total runtime is $\in O(n)$.
