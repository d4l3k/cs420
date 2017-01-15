# 2.
> (from Problem 4 in 3.2.3. Exercises from “Computational Geometry in C” by J. O’Rourke) Affine hulls. An affine combination of points $p_1, \ldots, p_n$ is a sum of the form $\alpha_1p_1 + \ldots +\alpha_np_n$, with $\alpha_1+\ldots+\alpha_n = 1$. Note this differs from the definition of a convex combination in that the condition $\alpha_i ≥ 0$ is dropped. What is the affine hull (the set of possible affine combinations) of two points in the plane? Three points? n > 3 points? What is the affine hull of two points in three dimensional space? Three points? Four points? n > 4 points?


The affine hull contains the convex hull of the points since it is a subset of
the convex hull definition. The affine hull of two points is the line formed by
those two points. The affine hull of three points assuming they are not in a
line is the entire coordinate space in $R^2$ and a plane in $R^3$. The affine
hull of four points in $R^3$ is the entire $R^3$ coordinate space assuming they
are not in a plane or line. Having more than 3 points in $R^2$ creates multiple
duplicate solutions for every $R^2$ coordinate as does having more than 4
points in $R^3$.
