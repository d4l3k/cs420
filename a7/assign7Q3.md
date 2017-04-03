# Question 3

Proof via reduction.

ALG is "Given an n-vertex graph G, return any integer between χ(G) and χ(G) +
573".

To show that ALG is NP-Hard we must do a reduction from computing the chromatic
number to ALG.

Create a new graph with 574 copies for each vertex. Make each set of vertex
points fully connected, and then color the graph. Since the error has to be <
573, the output number we get is the (true number * 574) + (0, 573), divide that
by 574 and round down to get the true chromatic number. This transformation can
happen in $O(n)$ time since there are $574n$ vertices.

Since we can compute the chromatic number, which is NP-hard using ALG, ALG must
also be NP-hard.
