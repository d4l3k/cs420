# 5.

Assume that there are two points such that $P_1, P_2 \in P$ and have the largest
distance between them of any two points $\in P$ and are not on the edge of the
convex hull. Since the shape is convex, any line between two points in the
convex hull is in the convex hull by definition. There must be a point $P_3$ on the
convex hull that is aligned with the line going through $P_1$ and $P_2$ due to
the convex structure. Since $P_3$ is on the convex hull, it must be further away
from one of $P_1, P_2$ and thus having a longer distance than that between $P_1$
and $P_2$. This is a contradiction and thus the two points with the largest
distance between them must be on the convex hull.
