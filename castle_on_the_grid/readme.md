# Castle On The Grid


<p>You are given a square grid with some cells open (<strong>.</strong>) and some blocked (<strong>X</strong>).  Your playing piece can move along any row or column until it reaches the edge of the grid or a blocked cell.  Given a grid, a start and a goal, determine the minmum number of moves to get to the goal.  </p>

<p><strong>Example</strong>. </p>
<p>The grid is shown below:</p>

<pre><code>...
.X.
...
</code></pre>

<p>The starting position is ( 0, 0 ) which moves from (X,Y) = (0,0) to (0,1) to (1,1). Hence take 2 steps to reach the goal.</p>

<p><strong>Function Description</strong> <br>
Complete the <em>minimumMoves</em> function in the editor.   </p>

<p>minimumMoves has the following parameter(s):</p>

<ul>
<li><em>string grid[n]:</em> an array of strings that represent the rows of the grid  </li>
<li><em>int startX:</em> starting X coordinate    </li>
<li><em>int startY:</em> starting Y coordinate    </li>
<li><em>int goalX:</em> ending X coordinate    </li>
<li><em>int goalY:</em> ending Y coordinate    </li>
</ul>

<p><strong>Returns</strong>  </p>

<ul>
<li><em>int:</em> the minimum moves to reach the goal</li>
</ul><pre><code>STDIN       FUNCTION
-----       --------
3           grid[] size n = 3
.X.         grid = ['.X.','.X.', '...']
.X.
...
0 0 0 2     startX = 0, startY = 0, goalX = 0, goalY = 2
</code></pre></div></div></div>