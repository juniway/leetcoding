[LeetCode 110] Balanced Binary Tree


class Solution {
public:
    bool isBalanced(TreeNode *root) {
		if (root == nullptr) return true;
 		return getHeight(root) == -1 ? false : true;
	}

	int getHeight(TreeNode *root) {
		if (root == nullptr) return 0;
		int left = getHeight(root->left);
		if(left == -1) return -1;   // no need to call right = getHeight(root->right) if left == -1
		int right = getHeight(root->right);
		if (right == -1) return -1;
		if (abs(left - right) > 1) return -1;
		return max(left, right) + 1;

	}
};

class Solution {
public:
	boolean isBalanced(TreeNode *root) {
        if(root == nullptr) return true;
        // 如果子树高度差大于1，则不平衡
        if(abs(depth(root->left) - depth(root->right)) > 1){
        	return false;
        }
        return isBalanced(root->left) && isBalanced(root->right);
    }

	int depth(TreeNode *root){
		if(root == null) return 0;
		return 1 + max(depth(root->left), depth(root->right));
	}
};
