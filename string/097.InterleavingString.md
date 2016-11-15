[LeetCode 97] Interleaving String

Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
For example,
Given:
s1 = "aabcc",
s2 = "dbbca",

When s3 = "aadbbcbcac", return true.
When s3 = "aadbbbaccc", return false.

Difficulty: Hard

class Solution {
public:
    bool isInterleave(string s1, string s2, string s3) {
        int n1 = s1.length(), n2 = s2.length(), n3 = s3.length();
        if(n1+n2 != n3) return false;
        bool dp[n1 + 1][n2 + 1];
        dp[0][0] = true;
        
        for (int i = 0; i < n2; ++i)
            dp[0][i+1] = dp[0][i] && s2[i] == s3[i];
        for (int i = 0; i < n1; ++i)
            dp[i+1][0] = dp[i][0] && s1[i] == s3[i];
        
        for(int i = 0; i < n1; ++i){
        	for(int j = 0; j < n2; ++j){
                dp[i+1][j+1] = (dp[i][j+1] && s1[i] == s3[i+j+1]) | (dp[i+1][j] && s2[j] == s3[i+j+1]);
        	}
        }
        return dp[n1][n2];
    }
};