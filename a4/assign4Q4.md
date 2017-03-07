# Question 4


If $x < y -1$ or $x = y + 1$, player Row wins.

if $x > y + 1$ or $x = y - 1$, player Col wins.

If $x = y$ no one wins.

This encourages both Row and Col to pick small numbers, but they can't be too
small.


## When n = 2

### Row

|   | 1  |  2 |
|---|---|---|
| 1 | 0  | -1 |
| 2 | 1  |  0 |

From these numbers, the optimal strategy is to always play 2, since you can't
lose to the other player, either both people will lose, or you'll win.


## When n = 3

### Row

|   |  1 |  2 |  3 |
|---|---|---|---|
| 1 |  0 | -1 |  1 |
| 2 |  1 |  0 | -1 |
| 3 | -1 |  1 |  0 |

For n=3, there's no optimal single strategy. If you only play a single number,
Col will win every time. If you only play two of the numbers, Col will win 2/3
the time. The smallest optimal strategy is to randomize equally between all
three options.

## When n > 3

### Row

|   |  1 |  2 |  3 |  4 |
|---|---|---|---|---|
| 1 |  0 | -1 |  1 |  1 |
| 2 |  1 |  0 | -1 |  1 |
| 3 | -1 |  1 |  0 | -1 |
| 4 | -1 | -1 |  1 |  0 |

Beyond, n = 3, there is no benefit to picking $x > 3$ since it's heavily skewed
in the favor of the other player. Thus, you want to randomize between 1-3
equally.
