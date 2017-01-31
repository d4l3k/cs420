# Question 2

Proof by induction.

Let S be a set of n points in the plane. A Voroni diagram partitions the plane
into Voroni regions where $R_i$ is the set of points in the plane whose closest
site is $s_i$, that is, $R_i=\left\{p|d(p,s_i) \leq d(p,s_j)\right\}$ for all $j$.

## Base case

There is only a single point in the Voronoi diagram. Since there is only one
Voronoi region that encompasses all space it is trivially convex.

## Inductive step

The line formed by the edge of two Voronoi regions is perpendicular to the line
formed by the points of the two Voronoi regions $(p_1, p_2)$, since every point
along the line is equal distance from $p_1$ and $p_2$.

Assume there is a voronoi region that is convex.

When adding a new point to a voronoi diagram the new line formed partitions the
old convex region into a smaller new region. When a line intersects with a
convex hull, all of the edge angles are less than 180 degrees. Thus,
partitioning a convex shape into two with a single line creates two new convex
shapes. Therefore the new voronoi region is still convex.
