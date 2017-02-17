# Question 3

We can solve this by creating a flow network from a to b. If we set all the
capacities of the edges to 1, we get something close to the vertex disjoint
paths, but not quite. There is an edge case where there are two edges going into
and out of a vertex, in which case there will be two paths using that vertex. To
solve this we have to split every vertex into two, one with all the incoming
edges, the other with all the outgoing edges. We then link these new vertexes
with a single edge of capacity 1. Thus, there can only one path that will go
through that vertex. We can then solve it using something like the
Ford-Fulkerson algorithm to find all disjoint paths.

Since there is a capacity of 1 on every vertex, that means that there must not
be any two paths that share vertices, thus all paths must be disjoint.
