vector<vector<int>> levelOrder(TreeNode* root) {
    vector<vector<int>> ans;
    
    if (!root) {
        return ans;
    }
    
    std::queue<TreeNode*> q;
    q.push(root);
    
    while(!q.empty()) {
        vector<int> current;
        int len = q.size();
        
        for(int i=0; i<len; ++i) {
            TreeNode* node = q.front();
            q.pop();
            current.push_back(node->val);
            
            if (node->left) {
                q.push(node->left);
            }
            
            if (node->right) {
                q.push(node->right);
            }
        }
        
        ans.push_back(current);
    }
    
    return ans;
}
