package maths

// 380+ms
//class Solution {
//private static final int N = 9, M0 = 205;
//private final int[] pow10 = new int[M0], pinv = new int[M0], l = new int[N], len = new int[1 << N];
//private int n, ans, p, r, B;
//private final Map<Integer, List<Integer>> c = new HashMap<>();
//private final Map<Integer, List<Node>> d = new HashMap<>();
//private final Map<Integer, Integer> hash = new HashMap<>();
//
//public int treeOfInfiniteSouls(int[] a, int p, int r) {
//n = a.length;
//ans = 0;
//B = (n + 2) / 3;
//this.p = p;
//this.r = r;
//if (p == 2 || p == 5)
//return p == 2 && r == 1 || p == 5 && r == 4 ? (int) (C((n - 1) * 2, n - 1) / n * fac(n)) : 0;
//pow10[0] = 1 % p;
//for (int i = 1; i < M0; ++i) pow10[i] = (int) ((long) pow10[i - 1] * 10 % p);
//for (int i = 0; i < M0; ++i) pinv[i] = (int) inv(pow10[i], p);
//for (int i = 0; i < n; ++i) l[i] = log10(a[i]);
//for (int i = 1; i < (1 << n); ++i) {
//len[i] = (Integer.bitCount(i) * 2 - 1) * 2;
//for (int j = 0; j < n; ++j) if ((i & (1 << j)) > 0) len[i] += l[j];
//}
//for (int i = 0; i < n; ++i)
//c.computeIfAbsent(1 << i, key -> new ArrayList<>()).add((int) (((long) a[i] * 10 + pow10[l[i] + 1] + 9) % p));
//for (int i = 1; i < (1 << n); ++i)
//if (Integer.bitCount(i) <= B * 2) {  //component below u
//int t = pow10[len[i] - 1] + 9;
//for (int j = (i - 1) & i; j != 0; j = (j - 1) & i)
//if (n == 9 || Integer.bitCount(i) < Math.max((n + 1) / 2, 2) || Math.max(Integer.bitCount(j), Integer.bitCount(i - j)) <= Math.min(B, (n - 1) / 2))
//for (Integer v1 : c.getOrDefault(j, new ArrayList<>())) {
//long t1 = (long) v1 * pow10[len[i - j] + 1] + t;
//for (Integer v2 : c.getOrDefault(i - j, new ArrayList<>()))
//c.computeIfAbsent(i, key -> new ArrayList<>()).add((int) (((long) v2 * 10 + t1) % p));
//}
//}
//d.computeIfAbsent(1 << n, key -> new ArrayList<>()).add(new Node(0, 0, 0)); //component above u
//BiFunction<Integer, Integer, Boolean> yes = (x, y) -> true;
//BiFunction<Integer, Integer, Boolean> no = (x, y) -> false;
//if (n == 9) {
//calc(4, yes, no, (j) -> Integer.bitCount(j) != 6); //remove size 6
//calc(5, (i, j) -> j != (1 << n) || Integer.bitCount(i - j) >= 2, //remove size 5
//no, (j) -> Integer.bitCount(j) != 5);
//calc(6, (i, j) -> j != (1 << n) || Integer.bitCount(i - j) == 3, //remove size 4
//(i, j) -> j == (1 << n) && Integer.bitCount(i - j) == 4,
//(j) -> Integer.bitCount(j) != 4);
//} else {
//calc(n / 2 + 1, yes, (i, j) -> n % 2 == 0 && Integer.bitCount(j) == 1 && Integer.bitCount(i - j) == n / 2,
//(j) -> Integer.bitCount(j) < (n + 1) / 2 || Integer.bitCount(j) > B * 2);
//}
//return ans;
//}
//
//private void calc(int t0, BiFunction<Integer, Integer, Boolean> f1, BiFunction<Integer, Integer, Boolean> f2, Function<Integer, Boolean> f3) {
//for (int i = (1 << n) + 1; i < (1 << (n + 1)); ++i) d.getOrDefault(i, new ArrayList<>()).clear();
//for (int i = 1 << n; i < (1 << (n + 1)); ++i)
//if (Integer.bitCount(i) <= t0) {
//boolean _f2;
//for (int j = (i - 1) & i; (j >> n) > 0; j = (j - 1) & i)
//if ((_f2 = f2.apply(i, j)) || f1.apply(i, j))
//for (Node t : d.getOrDefault(j, new ArrayList<>())) {
//int l1 = len[j - (1 << n)] - t.l + 2 * (j > (1 << n) ? 1 : 0);
//for (Integer v2 : c.getOrDefault(i - j, new ArrayList<>())) {
//d.computeIfAbsent(i, key -> new ArrayList<>()).add(new Node((t.v1 + pow10[l1]) % p,
//(int) (((long) t.v2 * pow10[len[i - j] + 1] + (long) v2 * 10 + 9) % p),
//t.l + len[i - j] + 1));
//if (!_f2)
//d.computeIfAbsent(i, key -> new ArrayList<>()).add(new Node((int) ((t.v1 + pow10[l1 + len[i - j]] + (long) v2 * pow10[l1]) % p),
//(int) (((long) t.v2 * 10 + 9) % p), t.l + 1));
//}
//}
//int j = (1 << (n + 1)) - 1 - i;
//hash.clear();
//if (f3.apply(j)) continue;
//for (Integer v : c.getOrDefault(j, new ArrayList<>())) hash.put(v, hash.getOrDefault(v, 0) + 1);
//for (Node t : d.getOrDefault(i, new ArrayList<>()))
//ans += hash.getOrDefault((int) ((((long) r - t.v2 - (long) t.v1 * pow10[len[j] + t.l]) % p + p) * pinv[t.l] % p), 0);
//}
//}
//
//private static class Node {
//int v1, v2, l;
//
//public Node(int v1, int v2, int l) {
//this.v1 = v1;
//this.v2 = v2;
//this.l = l;
//}
//}
//
//private int log10(int n) {
//return n < 10 ? 1 : log10(n / 10) + 1;
//}
//
//private long inv(long a, long b) {
//x = 0;
//y = 0;
//exgcd(a, b);
//return (x % b + b) % b;
//}
//
//private long x, y;
//
//private long exgcd(long a, long b) {
//if (b == 0) {
//x = 1;
//y = 0;
//return a;
//}
//long d = exgcd(b, a % b);
//long tmp = x;
//x = y;
//y = tmp - a / b * y;
//return d;
//}
//
//private long fac(int n) {
//long res = 1;
//while (n > 0) res *= n--;
//return res;
//}
//
//private long C(int n, int m) {
//long res = 1L;
//for (int i = 1; i <= m; i++) res = res * (n - i + 1) / i;
//return res;
//}
//}
