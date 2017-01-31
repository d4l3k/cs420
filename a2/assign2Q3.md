# Question 3

We need to show that the closest pair of points of different colors in S are on
an edge of a triangle in the Delaunay triangulation.

Proof by contradiction.

Assume that the two closest points $p, q$ of different colors are not linked by
an edge between them. These points must form a triangle on the side of the other
point of the pair since by triangulation there must be triangles on every side.
Since there are only two colors, the triangles the closest points are part of
must have pairs of opposite colors and thus the vertexes must be further than
the distance between $p$ and $q$. Thus, the other closest point would be
included in the circumcircle of one of the triangles.

This violates the definition of the Delaunay triangulation and thus the closest
pair of points of different colors in S are on the edge of a triangle in the
Delaunay triangulation.
