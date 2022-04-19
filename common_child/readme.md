# Common Child

<p>A string is said to be a child of a another string if it can be formed by deleting 0 or more characters from the other string.  Letters cannot be rearranged.  Given two strings of equal length, what's the longest string  that can be constructed such that it is a child of both?  </p>

<p><strong>Example</strong>   </p>

<p>These strings have two children with maximum length 3, <code>ABC</code> and <code>ABD</code>.  They can be formed by eliminating either the <code>D</code> or <code>C</code> from both strings.  Return 3</p>

<p><strong>Function Description</strong></p>

<p>Complete the <em>commonChild</em> function in the editor below.  </p>

<p>commonChild has the following parameter(s):</p>

<ul>
<li><em>string s1:</em>  a string</li>
<li><em>string s2:</em>  another string   </li>
</ul>

<p><strong>Returns</strong>   </p>

<ul>
<li><em>int:</em> the length of the longest string which is a common child of the input strings </li>
</ul></div></div></div>
<li>All characters are upper case in the range ascii[A-Z].</li>
</ul></div></div></div><p>The longest string that can be formed by deleting zero or more characters from , whose length is 2.</p>

<p><strong>Sample Input 1</strong></p>

<pre><code>AA
BB
</code></pre>

<p><strong>Sample Output 1</strong></p>

<pre><code>0
</code></pre>

<p><strong>Explanation 1</strong></p>

<p> have no characters in common and hence the output is 0.</p>

<p><strong>Sample Input 2</strong></p>

<pre><code>SHINCHAN
NOHARAAA
</code></pre>

<p><strong>Sample Output 2</strong></p>

<pre><code>3
</code></pre>

<p><strong>Explanation 2</strong></p>

<p>The longest string that can be formed between .</p>

<p><strong>Sample Input 3</strong></p>

<pre><code>ABCDEF
FBDAMN
</code></pre>

<p><strong>Sample Output 3</strong></p>

<pre><code>2
</code></pre>

<p><strong>Explanation 3</strong> <br>
 is the longest child of the given strings.</p></div></div></div></div></div></div>