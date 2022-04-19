# Balanced Brackets

<p>A bracket is considered to be any one of the following characters: <code>(</code>, <code>)</code>, <code>{</code>, <code>}</code>, <code>[</code>, or <code>]</code>. </p>

<p>Two brackets are considered to be a <em>matched pair</em> if the an opening bracket (i.e., <code>(</code>, <code>[</code>, or <code>{</code>) occurs to the left of a closing bracket (i.e., <code>)</code>, <code>]</code>, or <code>}</code>) <em>of the exact same type</em>. There are three types of matched pairs of brackets: <code>[]</code>, <code>{}</code>, and <code>()</code>.</p>

<p>A matching pair of brackets is <em>not balanced</em> if the set of brackets it encloses are not matched. For example, <code>{[(])}</code> is not balanced because the contents in between <code>{</code> and <code>}</code> are not balanced. The pair of square brackets encloses a single, unbalanced opening bracket, <code>(</code>, and the pair of parentheses encloses a single, unbalanced closing square bracket, <code>]</code>.</p>

<p>By this logic, we say a sequence of brackets is <em>balanced</em> if the following conditions are met:</p>

<ul>
<li>It contains no unmatched brackets.</li>
<li>The subset of brackets enclosed within the confines of a matched pair of brackets  is also a matched pair of brackets.</li>
</ul>

<p>Given strings of brackets, determine whether each sequence of brackets is balanced. If a string is balanced, return <code>YES</code>.  Otherwise, return <code>NO</code>.  </p>

<p><strong>Function Description</strong>  </p>

<p>Complete the function <em>isBalanced</em> in the editor below.   </p>

<p>isBalanced has the following parameter(s):</p>

<ul>
<li><em>string s</em>: a string of brackets   </li>
</ul>

<p><strong>Returns</strong>   </p>

<ul>
<li><em>string:</em> either <code>YES</code> or <code>NO</code>  </li>
</ul></div></div></div>

is the length of the sequence.  </li>
<li>All chracters in the sequences ∈ { <strong>{</strong>, <strong>}</strong>, <strong>(</strong>, <strong>)</strong>, <strong>[</strong>, <strong>]</strong> }.</li>

STDIN           Function
-----           --------
3               n = 3
{[()]}          first s = '{[()]}'
{[(])}          second s = '{[(])}'
{{[[(())]]}}    third s ='{{[[(())]]}}'
<li>The string <code>{[()]}</code> meets both criteria for being a balanced string. </li>
<li>The string <code>{[(])}</code> is not balanced because the brackets enclosed by the matched pair <code>{</code> and <code>}</code> are not balanced: <code>[(])</code>.      </li>
<li>The string <code>{{[[(())]]}}</code> meets both criteria for being a balanced string.   </li>
</ol></div></div></div></div></div></div>