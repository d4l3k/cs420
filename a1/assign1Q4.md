# 4.

Pick two points that are most definitely not in the convex hull. This can be
done by calculating the centeroids of both of the convex hulls and then finding
the point in P closest to the centeroid of Q and the point in Q closest to the
centeroid of P. Since these are the inner points they will not be in the final
convex hull. We will use these two points to determine where to join P and Q
when doing the second part of Graham's scan. Since the second part of Graham's
scan is O(n), the total runtime is $O(n+m)$.


## old/wrong
Pick one point that will be part of the convex hull. Similarly to Graham Scan,
you can pick the coordinate with the smallest Y coordinate, then compute the
angle from that coordinate to that of every other coordinate. This can be done
in $O(m+n)$ time since there are $m+n$ points with a constant amount of work for
each point.


