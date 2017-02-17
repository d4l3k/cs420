# Question 4

## (a) Give an example of a matrix M that is not rearrangeable even through every column and every row contains at least one entry that equals 1.

$$
\begin{bmatrix}
1 & 0 & 0 \
1 & 0 & 0 \
0 & 0 & 1
\end{bmatrix}


## (b) Describe an efficient algorithm that determines if M is rearrangeable.


Check if every row has a distinct 1, then check every column.
$$
